# download-logs

```
Create a go dir in your home directory then use go get to install.
package.

mkdir -p  go/{src,bin,pkg}

go get -v -u github.com/ChaosHour/aws-rds-logs/cmd/download-logs
```

### Usage
```
download-logs -h
Usage of download-logs:
  -d string
    	db-instance-identifier Example: aurora-test-copy-stg-0
  -l string
    	log-file-name Example: slowquery/mysql-slowquery.log.2021-09-08.20
  -r string
    	--region Example: us-west-1

```

```
download-logs -d aurora-test-copy-stg-0 -l slowquery/mysql-slowquery.log.2021-09-16.16
Executing command:
[aws rds download-db-log-file-portion --db-instance-identifier aurora-test-copy-stg-0 --log-file-name slowquery/mysql-slowquery.log.2021-09-16.16 --output text]

/rdsdbbin/oscar/bin/mysqld, Version: 5.7.12-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
/rdsdbbin/oscar/bin/mysqld, Version: 5.7.12-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
/rdsdbbin/oscar/bin/mysqld, Version: 5.7.12-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
# Time: 2021-09-16T15:45:20.188927Z
# User@Host: checkxxx[checkxxx] @  [10.213.15.234]  Id: 1250245
# Query_time: 0.110489  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
SET timestamp=1631807120;
show slave status;
# Time: 2021-09-16T15:48:23.996399Z
# User@Host: checkxxx[checkxxx] @  [10.213.15.234]  Id: 1250290
# Query_time: 0.174385  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
SET timestamp=1631807303;
show slave status;
# Time: 2021-09-16T15:49:25.244196Z
# User@Host: checkxxx[checkxxx] @  [10.213.15.234]  Id: 1250302
# Query_time: 0.102401  Lock_time: 0.000000 Rows_sent: 0  Rows_examined: 0
SET timestamp=1631807365;
show slave status;
/rdsdbbin/oscar/bin/mysqld, Version: 5.7.12-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
/rdsdbbin/oscar/bin/mysqld, Version: 5.7.12-log (MySQL Community Server (GPL)). started with:
Tcp port: 3306  Unix socket: /tmp/mysql.sock
Time                 Id Command    Argument
```
