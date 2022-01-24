# Cyderes/Fishtech golang assessment

  

## iphelper: IP enrichment serverless function for AWS lambda

###  Andrew Burt

Command to run in the cloud:
`curl https://ia3xx10f46.execute-api.us-west-2.amazonaws.com/staging/iphelper?ip=1.1.1.1`
(replace `ip=` with whatever IP you like)

All tests run, though one test requires a private API key to be exported to the environment before it will return successfully (tests the 3rd-party API lookup).

The API right now uses the ipdata service to do an ip lookup and I just enriched it with a few arbitrary items from their lookup, though it could as easily been any others, or all of them. The JSON I returned includes the IP address itself, the country name it's from, the time zone offset from GMT/UTC, and a security lookup, which checks whether the IP address is a proxy or not. 

I worked on the project Saturday and Sunday, and completed most of the requirements, but not Terraform. The function and API Gateway run successfully on Lambda, reliably deploy, I have a couple of tests running, make the external API call successfully, validate the IP address coming in, set up structs for the incoming and outgoing JSON, and pass the errors through Amazon API Gateway's proxy response container.

Uncompleted: Terraform implementation, and an additional enrichment component. I spent a large part of Sunday working on getting Terraform deployment working, but though I think I had nearly everything correct, I was not able to get it to work (it kept hanging when I went to apply it, and the terraform debug log results indicated a timeout between amazon servers that I wasn't able to diagnose in time). This was my first time using AWS, and my first time using Terraform, so probably due to lack of familiarity with the config syntax. The function and the API Gateway run fine when deployed manually using the AWS command-line tools. Unfortunately, an additional enrichment component (an amazon Dynamodb result caching and lookup) I had started to work on also was left unfinished due to getting stuck on Terraform for so long. 

I included the nonfunctional terraform files. There is a 3rd-party API key that needs to be passed in on the terraform pland and apply commands (e.g. `terraform apply -var="KEY=PrivacyMatters"`)
