data "aws_iam_policy_document" "trust_policy_document" {
    statement {
        actions = ["sts:AssumeRole"]
        effect = "Allow"

        principals {
            type        = "Service"
            identifiers = ["lambda.amazonaws.com"]
        }
  }
}

data "aws_iam_policy_document" "privilege_policy_document" {
    statement {
        actions = [
                "dynamodb:PutItem",
                "dynamodb:GetItem"
                ]
        effect = "Allow"
        resources = ["*"]
  }
}

resource "aws_iam_policy" "privilege_policy" {
  name        = "dynamodb-item-crud-role"
  policy = "${data.aws_iam_policy_document.privilege_policy_document.json}"
}

resource "aws_iam_role" "lambda_iphelper_executor" {
  name = "lambda-iphelper-executor"
  assume_role_policy = "${data.aws_iam_policy_document.trust_policy_document.json}"
}

resource "aws_iam_role_policy_attachment" "lambda_iphelper_executor_attach_trust_policy" {
  role       = "${aws_iam_role.lambda_iphelper_executor.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_iphelper_executor_attach_privilege_policy" {
  role       = "${aws_iam_role.lambda_iphelper_executor.name}"
  policy_arn = "${aws_iam_policy.privilege_policy.arn}"
}

