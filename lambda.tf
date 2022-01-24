resource "aws_lambda_function" "iphelper_function" {
  filename         = "./main.zip"
  function_name    = "iphelper"
  role             = "${aws_iam_role.lambda_iphelper_executor.arn}"
  handler          = "main"
  source_code_hash = "${filebase64sha256("./main.go")}"
  runtime          = "go1.x"

}