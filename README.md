rcheck is designed to simplify the process of verifying SPF (Sender Policy Framework) and DMARC (Domain-based Message Authentication, Reporting, and Conformance) records for any domain. With rcheck, you can quickly identify whether a domain has SPF and DMARC records configured, helping you ensure proper email authentication and security practices.

# Features

rcheck offers a straightforward command-line interface, making it easy to use for both beginners and experienced users.
- Verify if a domain has an SPF record configured.
- Display the SPF record if found.
- Highlight if the SPF record is missing.

# Install rcheck

You can use the following command to install rcheck:

```bash
go install -v github.com/0xhnl/rcheck
```

# Usage

You can use the following command to scan the list of targets.

```bash
$ rcheck dummy.txt
SPF record found for api.hackerone.com: "70gn9hp69jzpn3nkp42r8n9jwwtd1d70"
"v=spf1 -all"

DMARC record not found for api.hackerone.com
--------------------------
SPF record not found for gslink.hackerone.com
DMARC record not found for gslink.hackerone.com
--------------------------
SPF record found for support.hackerone.com: "v=spf1 -all"

DMARC record not found for support.hackerone.com
--------------------------
SPF record not found for go.hackerone.com
DMARC record not found for go.hackerone.com
--------------------------
SPF record not found for 3d.hackerone.com
DMARC record not found for 3d.hackerone.com
--------------------------
SPF record not found for _dmarc.hackerone.com
DMARC record not found for _dmarc.hackerone.com
--------------------------
SPF record not found for resources.hackerone.com
DMARC record not found for resources.hackerone.com
--------------------------
SPF record not found for info.hackerone.com
DMARC record not found for info.hackerone.com
--------------------------
SPF record found for fwdkim1.hackerone.com: spfmx1.domainkey.freshemail.io.
"v=spf1 include:sendgrid.net include:fdspfus.freshemail.io include:fdspfeuc.freshemail.io include:fdspfind.freshemail.io include:fdspfaus.freshemail.io ~all"

DMARC record not found for fwdkim1.hackerone.com
--------------------------
SPF record not found for b.ns.hackerone.com
DMARC record not found for b.ns.hackerone.com
```
