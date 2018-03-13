[![Build Status](https://www.travis-ci.org/baldguysoftware/pf2s3.svg?branch=master)](https://www.travis-ci.org/baldguysoftware/pf2s3)

# Postfix To S3

Ok, so I've been using this on my Raspberry Pi network for a while now, figure
I'd put it here for public consumption. This is intended to be run as a Postfix
delivery command. For example, if you have a address "upload@foo.com", then if
you configure that alias to be `upload@foo.com |"path/to/pf2s3 -b BUCKET-NAME
-p PATH-IN-BUCKET " and have the needed aws credentials set  `.aws/creds` file
for the user it will run as, then Postfix will send the message over standard in
and the executable will upload it.

For the name of the object in s3 it will extract the Message-Id header it finds
in the email. It will also extract the sender and recipient address and set
them as the values of the tags "sender" and "recipient", respectively. 

You will need to create the bucket prior to running this for the first time.
Included in the repo is a simple test message.


