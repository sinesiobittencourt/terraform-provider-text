# terraform-provider-text
Allows for the storing off arbitrary text in the state file

# Installation
No special configuration is needed. 
```
provider text {}
```

# Usage
At its simplest, just hardcode some text.
```
resource text important_note {
    content = ""
}
```
then if someone changes it then it will show up in the plan:
```
resource text test_sh {
    content = "Next person that runs this need to remember to do the thing"
}
```

# The Real Usage
The use case that prompted this provider is passing configuration and configuration scripts to AWS instances.  Initially we used the cloudinit template data source but this rapidly grew very large, to the point of hitting the character limit for the instance user_data field.

The solution we came up with was to drop user_data and, instead, push the configuration scripts to s3 in a provisioner, run Ansible to deploy them and then remove them from s3 in another provisioner.

This solution has a fairly serious downside though.  The configuration scripts are no longer processed by Terraform and so if they change, no change in the infrastructure is generated.  Fine if you know about it because you can run the Ansible manually but when that arcane knowledge is lost, it is very easy for new instances to end up configured different from existing hosts.

What we can do now is simply add a text resource for each config file.  The actual deployment still happens via s3 but if the script files change then it flags up in the plan.

```
resource text test_sh {
    content = file("${path.module}/test.sh")
}
```