# Domain Records Scanner

rcheck is designed to simplify the process of verifying SPF (Sender Policy Framework) and DMARC (Domain-based Message Authentication, Reporting, and Conformance) records for any domain. With rcheck, you can quickly identify whether a domain has SPF and DMARC records configured, helping you ensure proper email authentication and security practices.

# Features

rcheck offers a straightforward command-line interface, making it easy to use for both beginners and experienced users.
- Verify if a domain has an SPF record configured.
- Display the SPF record if found.
- Highlight if the SPF record is missing.

# Install rcheck

You can use the following command to install rcheck:

```bash
go install -v github.com/0xhnl/rcheck@latest
```

# Usage

You can use the following command to scan the list of targets.

```bash
$ rcheck -i dummy.txt
Not Found SPF record for b.ns.hackerone.com
Not Found DMARC record for b.ns.hackerone.com
--------------------------
Not Found SPF record for gslink.hackerone.com
Not Found DMARC record for gslink.hackerone.com
--------------------------
Not Found SPF record for mta-sts.managed.hackerone.com
Not Found DMARC record for mta-sts.managed.hackerone.com
--------------------------
Found SPF record for support.hackerone.com
Not Found DMARC record for support.hackerone.com
--------------------------
Not Found SPF record for a.ns.hackerone.com
Not Found DMARC record for a.ns.hackerone.com
--------------------------
Found SPF record for zendesk3.hackerone.com
Not Found DMARC record for zendesk3.hackerone.com
--------------------------
Not Found SPF record for docs.hackerone.com
Not Found DMARC record for docs.hackerone.com
--------------------------
Not Found SPF record for mta-sts.forwarding.hackerone.com
Not Found DMARC record for mta-sts.forwarding.hackerone.com
--------------------------
Not Found SPF record for events.hackerone.com
Not Found DMARC record for events.hackerone.com
--------------------------
Not Found SPF record for design.hackerone.com
Not Found DMARC record for design.hackerone.com
--------------------------
Found SPF record for fwdkim1.hackerone.com
Not Found DMARC record for fwdkim1.hackerone.com
--------------------------
Found SPF record for hackerone.com
Found DMARC record for hackerone.com
--------------------------
Not Found SPF record for _dmarc.hackerone.com
Not Found DMARC record for _dmarc.hackerone.com
--------------------------
Not Found SPF record for mta-sts.hackerone.com
Not Found DMARC record for mta-sts.hackerone.com
--------------------------
Found SPF record for zendesk1.hackerone.com
Not Found DMARC record for zendesk1.hackerone.com
--------------------------
Not Found SPF record for 3d.hackerone.com
Not Found DMARC record for 3d.hackerone.com
--------------------------
Found SPF record for zendesk4.hackerone.com
Not Found DMARC record for zendesk4.hackerone.com
--------------------------
Found SPF record for api.hackerone.com
Not Found DMARC record for api.hackerone.com
--------------------------
Not Found SPF record for resources.hackerone.com
Not Found DMARC record for resources.hackerone.com
--------------------------
Found SPF record for www.hackerone.com
Not Found DMARC record for www.hackerone.com
--------------------------
Found SPF record for zendesk2.hackerone.com
Not Found DMARC record for zendesk2.hackerone.com
--------------------------
Not Found SPF record for links.hackerone.com
Not Found DMARC record for links.hackerone.com
--------------------------
Not Found SPF record for go.hackerone.com
Not Found DMARC record for go.hackerone.com
--------------------------
Not Found SPF record for info.hackerone.com
Not Found DMARC record for info.hackerone.com
--------------------------
```
