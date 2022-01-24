output "iphelper_rest_api_endpoint" {
  value = "${aws_api_gateway_deployment.iphelper_rest_api_deployment.invoke_url}/iphelper"
}