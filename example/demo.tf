
terraform {
  required_providers {
    demo = {
      version = "0.1"
      source = "hashicorp.com/malanyuk/demo"
    }
  }
}

resource "demo_a" "all" {
  name = "data"
  file = "a"
}

resource "demo_a" "aok" {
  name = "o"
  file = "b"
}


data "demo_a" "all" {
  depends_on = [demo_a.all]
  file = "a"
}

output "data" {
  value = data.demo_a.all
}

output "res" {
  value = demo_a.all
}