resource "aws_dynamodb_table" "Ranking" {
  name = "Ranking"

  billing_mode = "PAY_PER_REQUEST"

  attribute {
    name = "PlayerID"
    type = "S"
  }

  attribute {
    name = "Score"
    type = "N"
  }

  hash_key  = "PlayerID"
  range_key = "Score"

  lifecycle {
    create_before_destroy = true
    ignore_changes        = [read_capacity, write_capacity]
  }

  tags = {
    Environment = var.environment
  }
}
