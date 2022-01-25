# sync-rds-logs

### Sync RDS logs to a local directory

```
You can install it with:
go get -v -u github.com/ChaosHour/aws-rds-logs/cmd/sync-rds-logs


The usage is as follows:

Usage:
  sync-rds-logs <dbInstanceIdentifier> <directory>

Example:
  sync-rds-logs pg-test-stg .
```

```
If you want to download logs from a different region:

export AWS_DEFAULT_REGION="us-west-2"

sync-rds-logs pg-test-stg .
INFO[0000] Found 73 log files for DBInstanceIdentifier=pg-test-stg
INFO[0000] Downloading error/postgresql.log.2021-10-18-18 to 'postgresql.log.2021-10-18-18'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-18
INFO[0000] Downloading error/postgresql.log.2021-10-18-19 to 'postgresql.log.2021-10-18-19'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-19
INFO[0000] Downloading error/postgresql.log.2021-10-18-20 to 'postgresql.log.2021-10-18-20'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-20
INFO[0000] Downloading error/postgresql.log.2021-10-18-21 to 'postgresql.log.2021-10-18-21'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-21
INFO[0000] Downloading error/postgresql.log.2021-10-18-22 to 'postgresql.log.2021-10-18-22'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-22
INFO[0000] Downloading error/postgresql.log.2021-10-18-23 to 'postgresql.log.2021-10-18-23'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-18-23
INFO[0000] Downloading error/postgresql.log.2021-10-19-00 to 'postgresql.log.2021-10-19-00'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-19-00
INFO[0000] Downloading error/postgresql.log.2021-10-19-01 to 'postgresql.log.2021-10-19-01'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-19-01
INFO[0000] Downloading error/postgresql.log.2021-10-19-02 to 'postgresql.log.2021-10-19-02'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-19-02
INFO[0000] Downloading error/postgresql.log.2021-10-19-03 to 'postgresql.log.2021-10-19-03'
INFO[0000] Successfully downloaded error/postgresql.log.2021-10-19-03
INFO[0000] Downloading error/postgresql.log.2021-10-19-04 to 'postgresql.log.2021-10-19-04'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-04
INFO[0001] Downloading error/postgresql.log.2021-10-19-05 to 'postgresql.log.2021-10-19-05'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-05
INFO[0001] Downloading error/postgresql.log.2021-10-19-06 to 'postgresql.log.2021-10-19-06'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-06
INFO[0001] Downloading error/postgresql.log.2021-10-19-07 to 'postgresql.log.2021-10-19-07'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-07
INFO[0001] Downloading error/postgresql.log.2021-10-19-08 to 'postgresql.log.2021-10-19-08'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-08
INFO[0001] Downloading error/postgresql.log.2021-10-19-09 to 'postgresql.log.2021-10-19-09'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-09
INFO[0001] Downloading error/postgresql.log.2021-10-19-10 to 'postgresql.log.2021-10-19-10'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-10
INFO[0001] Downloading error/postgresql.log.2021-10-19-11 to 'postgresql.log.2021-10-19-11'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-11
INFO[0001] Downloading error/postgresql.log.2021-10-19-12 to 'postgresql.log.2021-10-19-12'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-12
INFO[0001] Downloading error/postgresql.log.2021-10-19-13 to 'postgresql.log.2021-10-19-13'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-13
INFO[0001] Downloading error/postgresql.log.2021-10-19-14 to 'postgresql.log.2021-10-19-14'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-14
INFO[0001] Downloading error/postgresql.log.2021-10-19-15 to 'postgresql.log.2021-10-19-15'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-15
INFO[0001] Downloading error/postgresql.log.2021-10-19-16 to 'postgresql.log.2021-10-19-16'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-16
INFO[0001] Downloading error/postgresql.log.2021-10-19-17 to 'postgresql.log.2021-10-19-17'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-17
INFO[0001] Downloading error/postgresql.log.2021-10-19-18 to 'postgresql.log.2021-10-19-18'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-18
INFO[0001] Downloading error/postgresql.log.2021-10-19-19 to 'postgresql.log.2021-10-19-19'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-19
INFO[0001] Downloading error/postgresql.log.2021-10-19-20 to 'postgresql.log.2021-10-19-20'
INFO[0001] Successfully downloaded error/postgresql.log.2021-10-19-20
INFO[0001] Downloading error/postgresql.log.2021-10-19-21 to 'postgresql.log.2021-10-19-21'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-19-21
INFO[0002] Downloading error/postgresql.log.2021-10-19-22 to 'postgresql.log.2021-10-19-22'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-19-22
INFO[0002] Downloading error/postgresql.log.2021-10-19-23 to 'postgresql.log.2021-10-19-23'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-19-23
INFO[0002] Downloading error/postgresql.log.2021-10-20-00 to 'postgresql.log.2021-10-20-00'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-00
INFO[0002] Downloading error/postgresql.log.2021-10-20-01 to 'postgresql.log.2021-10-20-01'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-01
INFO[0002] Downloading error/postgresql.log.2021-10-20-02 to 'postgresql.log.2021-10-20-02'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-02
INFO[0002] Downloading error/postgresql.log.2021-10-20-03 to 'postgresql.log.2021-10-20-03'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-03
INFO[0002] Downloading error/postgresql.log.2021-10-20-04 to 'postgresql.log.2021-10-20-04'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-04
INFO[0002] Downloading error/postgresql.log.2021-10-20-05 to 'postgresql.log.2021-10-20-05'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-05
INFO[0002] Downloading error/postgresql.log.2021-10-20-06 to 'postgresql.log.2021-10-20-06'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-06
INFO[0002] Downloading error/postgresql.log.2021-10-20-07 to 'postgresql.log.2021-10-20-07'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-07
INFO[0002] Downloading error/postgresql.log.2021-10-20-08 to 'postgresql.log.2021-10-20-08'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-08
INFO[0002] Downloading error/postgresql.log.2021-10-20-09 to 'postgresql.log.2021-10-20-09'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-09
INFO[0002] Downloading error/postgresql.log.2021-10-20-10 to 'postgresql.log.2021-10-20-10'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-10
INFO[0002] Downloading error/postgresql.log.2021-10-20-11 to 'postgresql.log.2021-10-20-11'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-11
INFO[0002] Downloading error/postgresql.log.2021-10-20-12 to 'postgresql.log.2021-10-20-12'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-12
INFO[0002] Downloading error/postgresql.log.2021-10-20-13 to 'postgresql.log.2021-10-20-13'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-13
INFO[0002] Downloading error/postgresql.log.2021-10-20-14 to 'postgresql.log.2021-10-20-14'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-14
INFO[0002] Downloading error/postgresql.log.2021-10-20-15 to 'postgresql.log.2021-10-20-15'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-15
INFO[0002] Downloading error/postgresql.log.2021-10-20-16 to 'postgresql.log.2021-10-20-16'
INFO[0002] Successfully downloaded error/postgresql.log.2021-10-20-16
INFO[0002] Downloading error/postgresql.log.2021-10-20-17 to 'postgresql.log.2021-10-20-17'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-17
INFO[0003] Downloading error/postgresql.log.2021-10-20-18 to 'postgresql.log.2021-10-20-18'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-18
INFO[0003] Downloading error/postgresql.log.2021-10-20-19 to 'postgresql.log.2021-10-20-19'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-19
INFO[0003] Downloading error/postgresql.log.2021-10-20-20 to 'postgresql.log.2021-10-20-20'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-20
INFO[0003] Downloading error/postgresql.log.2021-10-20-21 to 'postgresql.log.2021-10-20-21'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-21
INFO[0003] Downloading error/postgresql.log.2021-10-20-22 to 'postgresql.log.2021-10-20-22'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-22
INFO[0003] Downloading error/postgresql.log.2021-10-20-23 to 'postgresql.log.2021-10-20-23'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-20-23
INFO[0003] Downloading error/postgresql.log.2021-10-21-00 to 'postgresql.log.2021-10-21-00'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-00
INFO[0003] Downloading error/postgresql.log.2021-10-21-01 to 'postgresql.log.2021-10-21-01'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-01
INFO[0003] Downloading error/postgresql.log.2021-10-21-02 to 'postgresql.log.2021-10-21-02'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-02
INFO[0003] Downloading error/postgresql.log.2021-10-21-03 to 'postgresql.log.2021-10-21-03'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-03
INFO[0003] Downloading error/postgresql.log.2021-10-21-04 to 'postgresql.log.2021-10-21-04'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-04
INFO[0003] Downloading error/postgresql.log.2021-10-21-05 to 'postgresql.log.2021-10-21-05'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-05
INFO[0003] Downloading error/postgresql.log.2021-10-21-06 to 'postgresql.log.2021-10-21-06'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-06
INFO[0003] Downloading error/postgresql.log.2021-10-21-07 to 'postgresql.log.2021-10-21-07'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-07
INFO[0003] Downloading error/postgresql.log.2021-10-21-08 to 'postgresql.log.2021-10-21-08'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-08
INFO[0003] Downloading error/postgresql.log.2021-10-21-09 to 'postgresql.log.2021-10-21-09'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-09
INFO[0003] Downloading error/postgresql.log.2021-10-21-10 to 'postgresql.log.2021-10-21-10'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-10
INFO[0003] Downloading error/postgresql.log.2021-10-21-11 to 'postgresql.log.2021-10-21-11'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-11
INFO[0003] Downloading error/postgresql.log.2021-10-21-12 to 'postgresql.log.2021-10-21-12'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-12
INFO[0003] Downloading error/postgresql.log.2021-10-21-13 to 'postgresql.log.2021-10-21-13'
INFO[0003] Successfully downloaded error/postgresql.log.2021-10-21-13
INFO[0003] Downloading error/postgresql.log.2021-10-21-14 to 'postgresql.log.2021-10-21-14'
INFO[0004] Successfully downloaded error/postgresql.log.2021-10-21-14
INFO[0004] Downloading error/postgresql.log.2021-10-21-15 to 'postgresql.log.2021-10-21-15'
INFO[0004] Successfully downloaded error/postgresql.log.2021-10-21-15
INFO[0004] Downloading error/postgresql.log.2021-10-21-16 to 'postgresql.log.2021-10-21-16'
INFO[0004] Successfully downloaded error/postgresql.log.2021-10-21-16
INFO[0004] Downloading error/postgresql.log.2021-10-21-17 to 'postgresql.log.2021-10-21-17'
INFO[0004] Successfully downloaded error/postgresql.log.2021-10-21-17
INFO[0004] Downloading error/postgresql.log.2021-10-21-18 to 'postgresql.log.2021-10-21-18'
INFO[0004] Successfully downloaded error/postgresql.log.2021-10-21-18

```

### Thank you! [Jonathan Stacks](https://github.com/jonstacks/aws)