package tarasmal

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var tempFile string = "/tmp/"

func dataA() *schema.Resource {
	return &schema.Resource{
		ReadContext: resAGet,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resA() *schema.Resource {
	return &schema.Resource{
		CreateContext: resACreate,
		ReadContext:   resAGet,
		UpdateContext: resAUpdate,
		DeleteContext: resADelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resACreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var context string = d.Get("name").(string)
	var param string = d.Get("file").(string)
	file, err := os.OpenFile(tempFile+param, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return diag.FromErr(err)
	}
	defer file.Close()
	_, err = file.Write([]byte(context))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(tempFile + param)
	return diags
}

func resAUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resACreate(ctx, d, m)
}

func resADelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var param string = d.Get("file").(string)
	err := os.Remove(tempFile + param)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resAGet(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var param string = d.Get("file").(string)
	if _, err := os.Stat(tempFile + param); err == nil {
		d.SetId(tempFile + param)

		//name
		file, err := os.Open(tempFile + param)
		if err != nil {
			return diag.FromErr(err)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		d.Set("name", string(data))
	} else {
		d.SetId("")
	}
	return diags
}
