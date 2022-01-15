# aws-rds-logs

###
Usage
```BASH
Create a go dir in your home directory then use go get to install the aws-rds-logs package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-rds-logs
```


```BASH 
aws-rds-logs -h
Usage of aws-rds-logs:
  -d string
    	Specify your db-instance-identifier to view it's logs.
  -l string
    	Log type to use Example: slow, error. (default "error")
  -r string
    	Select a aws region to view your clusters. (default "us-west-1")
```

```BASH

aws-rds-logs
dbRegion = us-west-1
Executing command:
aws rds describe-db-instances --region us-west-1 | jq -r '.DBInstances[] | { (.DBInstanceIdentifier):(.Endpoint.Address + ":" + (.Endpoint.Port|tostring))}' | sed -e 's/[{}]//g' -e 's/\"//g' -e '/^$/d'

Command Successfully Executed
  pg-test-stg:       pg-test-stg.xxxxx.us-west-1.rds.amazonaws.com:5432
  aurora-test-stg-0: aurora-test-stg-0.xxxxx.us-west-1.rds.amazonaws.com:3306
  aurora-test-stg-1: aurora-test-stg-1.xxxxxx.us-west-1.rds.amazonaws.com:3306


aws-rds-logs -d aurora-test-stg-0
dbInstanceIdentifier = aurora-test-stg-0
logType = error
Executing command:
aws rds describe-db-log-files --db-instance-identifier aurora-test-stg-0 | jq '.DescribeDBLogFiles[] | select(.LogFileName|test("error.")) | .LogFileName'

Command Successfully Executed
error/mysql-error-running.log
error/mysql-error-running.log.2021-07-19.20
error/mysql-error-running.log.2021-07-27.09
error/mysql-error-running.log.2021-07-27.10
error/mysql-error-running.log.2021-08-09.10
error/mysql-error.log"


aws-rds-logs -d aurora-test-stg-0 -l slow
dbInstanceIdentifier = aurora-test-stg-0
logType = slow
Executing command:
aws rds describe-db-log-files --db-instance-identifier aurora-test-stg-0 | jq '.DescribeDBLogFiles[] | select(.LogFileName|test("slow.")) | .LogFileName'

Command Successfully Executed
slowquery/mysql-slowquery.log
slowquery/mysql-slowquery.log.2021-08-17.19
slowquery/mysql-slowquery.log.2021-08-17.20
slowquery/mysql-slowquery.log.2021-08-17.21
slowquery/mysql-slowquery.log.2021-08-17.22
slowquery/mysql-slowquery.log.2021-08-17.23



````
