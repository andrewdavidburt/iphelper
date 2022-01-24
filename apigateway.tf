resource "aws_api_gateway_rest_api" "iphelper_rest_api" {
    name        = "iphelper-api"
}

resource "aws_api_gateway_resource" "iphelper_rest_api_resource" {
    rest_api_id = "${aws_api_gateway_rest_api.iphelper_rest_api.id}"
    parent_id   = "${aws_api_gateway_rest_api.iphelper_rest_api.root_resource_id}"
    path_part   = "iphelper"
}

resource "aws_api_gateway_method" "iphelper_rest_api_method" {
    rest_api_id   = "${aws_api_gateway_rest_api.iphelper_rest_api.id}"
    resource_id   = "${aws_api_gateway_resource.iphelper_rest_api_resource.id}"
    http_method   = "ANY"
    authorization = "NONE"
}

resource "aws_api_gateway_integration" "iphelper_rest_api_integration" {
    rest_api_id               = "${aws_api_gateway_rest_api.iphelper_rest_api.id}"
    resource_id               = "${aws_api_gateway_resource.iphelper_rest_api_resource.id}"
    http_method               = "${aws_api_gateway_method.iphelper_rest_api_method.http_method}"
    type                      = "AWS_PROXY"
    integration_http_method   = "POST"
    uri                       = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/${aws_lambda_function.iphelper_function.arn}/invocations"
}

resource "aws_lambda_permission" "iphelper_rest_api_integration" {
    statement_id  = "AllowIphelperInvoke"
    action        = "lambda:InvokeFunction"
    function_name = "${aws_lambda_function.iphelper_function.function_name}"
    principal     = "apigateway.amazonaws.com"
    source_arn    = "${aws_api_gateway_rest_api.iphelper_rest_api.execution_arn}/*/*/iphelper"
}


resource "aws_api_gateway_deployment" "iphelper_rest_api_deployment" {
    depends_on = [
        aws_api_gateway_integration.iphelper_rest_api_integration
    ]
    rest_api_id = "${aws_api_gateway_rest_api.iphelper_rest_api.id}"
    stage_name  = "staging"
}