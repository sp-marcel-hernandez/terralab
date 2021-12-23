resource "aws_dynamodb_table" "sample_table" {
  name = "sample_table"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "ID"
    type = "S"
  }

  hash_key  = "ID"

  lifecycle {
    create_before_destroy = true
    ignore_changes        = [read_capacity, write_capacity]
  }

  tags = {
    Foo = "Bar"
  }
}
