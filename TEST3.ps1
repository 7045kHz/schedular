Start-Sleep -Seconds 61
Get-ADUser Administrator -Property MemberOf,Name | Select-Object -ExpandProperty MemberOf   