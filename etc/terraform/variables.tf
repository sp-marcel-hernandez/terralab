variable "aws_region" {
  description = "AWS region to host your network"
  default     = "us-east-1"
}

variable "aws_endpoint" {
  description = "The AWS endpoint, which can be overwritten to use with Localstack"
  default     = "http://localstack:4566"
}
