# Domain Records Scanner

rcheck is designed to simplify the process of verifying SPF (Sender Policy Framework) and DMARC (Domain-based Message Authentication, Reporting, and Conformance) records for any domain. With rcheck, you can quickly identify whether a domain has SPF and DMARC records configured, helping you ensure proper email authentication and security practices.

# Features

rcheck offers a straightforward command-line interface, making it easy to use for both beginners and experienced users.
- Verify if a domain has an SPF record configured.
- Display the SPF record if found.
- Highlight if the SPF record is missing.

# Install rcheck

One shot command to install rcheck.

```bash
go install -v github.com/0xhnl/rcheck
```
