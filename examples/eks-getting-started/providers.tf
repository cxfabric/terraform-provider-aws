terraform {
  required_version = ">= 0.12"
  backend "s3" { // Configure backend S3 bucket for statefile. Bucket must pre-exist
    bucket = "cxfab-terraform-state"
    key    = "dev-aws-eks/terraform.tfstate" // create folder object 'project_2' for state file
    region = "us-east-1" // Variables cannot be used here
  }

}

provider "aws" {
  region = var.aws_region
}

data "aws_availability_zones" "available" {}

# Not required: currently used in conjunction with using
# icanhazip.com to determine local workstation external IP
# to open EC2 Security Group access to the Kubernetes cluster.
# See workstation-external-ip.tf for additional information.
provider "http" {}
