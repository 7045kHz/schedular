# schedular

## SQL Server Tables

```sql
CREATE TABLE [dbo].[JOB_INVENTORY](
    [Id] [int] IDENTITY(1,1) NOT NULL,
    [Name] [nvarchar](50) NOT NULL,
    [Verbose] [int] NOT NULL,
    [Execution_Server] [nvarchar](120) NOT NULL,
    [Enabled] [int] NOT NULL,
    [Job_Definition] [nvarchar](max) NOT NULL,
    [Days] [nchar](50) NOT NULL,
    [Hour] [int] NOT NULL,
    [Minute] [int] NOT NULL,
    [Created_On] [datetime] NOT NULL DEFAULT CURRENT_TIMESTAMP,
    [Last_Updated] [datetime] NOT NULL DEFAULT CURRENT_TIMESTAMP,
    [Last_Run] [datetime] NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

ALTER TABLE [dbo].[JOB_INVENTORY]
    ADD CONSTRAINT [Job_Definition record should be formatted as JSON]
                   CHECK (ISJSON(Job_Definition)=1)
GO

CREATE TRIGGER [dbo].[JOB_INVENTORY_AU]
ON  [dbo].[JOB_INVENTORY]
AFTER UPDATE
AS
BEGIN
    UPDATE X 
    SET Last_Updated = CURRENT_TIMESTAMP
    FROM [dbo].[JOB_INVENTORY] X
    INNER JOIN inserted
    AS i
    on X.Id = i.Id
END ;
```

## Insert into table

```sql
insert into [OSDISCOVERY].[dbo].[JOB_INVENTORY] (
    [Name]
    ,[Execution_Server]
    ,[Enabled]
    ,[Job_Definition]
    ,[Days]
    ,[Hour]
    ,[Minute]
    ,[Created_On]
    ,[Last_Updated] ) values ('Windows PowerShell Test','P330.RSYSLAB.COM', 1,'', '[0,1,2,3,4,5,6]', 11, 30, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
```

## JSON Job Definition

Note: Exec is different between BAT script and powershell script.

### PowerShell File Example

```sql
DECLARE @json NVARCHAR(4000) = N'{ 
    "Name": "Windows PowerShell Test",
    "Engine" : "POWERSHELL",
    "Exec" : "powershell.exe",
    "Verbose": 1,
    "Enabled": 1,
    "Env" : ["MY_VAR=ABC"],
    "Args" : ["--NoProfile","-NonInteractive",".\\TEST3.ps1"]
}';

-- select isjson(@json) ;
update [OSDISCOVERY].[dbo].[JOB_INVENTORY] 
set [Job_Definition] = @json
where Id = 1;
```

### Windows BAT File Example

```sql
DECLARE @json NVARCHAR(4000) = N'{ 
    "Name": "Windows Batch File Test 1",
    "Engine" : "CMD",
    "Exec" : ".\\TEST1.BAT",
    "Verbose": 1,
    "Env" : ["MY_VAR=TEST1"],
    "Args" : [""]
}';

-- select isjson(@json) ;
update [OSDISCOVERY].[dbo].[JOB_INVENTORY] 
set [Job_Definition] = @json
where Id = 2;
```
