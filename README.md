# Cyderes/Fishtech golang assessment

  

## iphelper: IP enrichment serverless function for AWS lambda

###  Andrew Burt

Command to run in the cloud:
`curl https://ia3xx10f46.execute-api.us-west-2.amazonaws.com/staging/iphelper?ip=1.1.1.1`
(replace `ip=` with whatever IP you like)

All tests run, though one requires a private API key to be exported before it will return successfully.

I worked on this Saturday and Sunday, and completed most of the requirements. The function and API Gateway run successfully on Lambda, reliably deploy, I have a couple of tests running, make the external API call 

Uncompleted: Terraform implementation, and an additional enrichment component. I spent a large part of Sunday working on getting Terraform deployment working, but though I think I had nearly everything correct, I was not able to get it to work (it kept hanging when I went to apply it, and the terraform debug log results indicated a timeout between amazon servers that I wasn't able to diagnose in time). This was my first time using AWS, and my first time using Terraform, so probably due to lack of familiarity with the config syntax. The function and the API Gateway run fine when deployed manually using the AWS command-line tools. Unfortunately, an additional enrichment component (an amazon Dynamodb result caching and lookup) I had started to work on also was left unfinished due to getting stuck on Terraform for so long. 

I included the nonfunctional terraform files. There is a 3rd-party API key that needs to be passed in on the terraform pland and apply commands (e.g. `terraform apply -var="KEY=PrivacyMatters"`)
