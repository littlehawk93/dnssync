# DNS Sync
A simple CLI tool for synchronizing your public IP with your DNS provider

Most home users have dynamic IP addresses issued from their ISPs, meaning they don't have any control over when the public IP to their home network changes. This is troublesome for those of us who have domains pointing to our home. dnssync is a simple CLI tool I wrote to solve this problem for myself. The number of DNS providers it supports is limited, but can be expanded if requested.

#### Currently Supported DNS Providers
- CloudFlare
- NameSilo

#### Examples

To use dnssync, simply run this command:

```sh
dnssync update -c "config.yaml" -p cloudflare -d home.example.com
```

This will automatically grab your current public IP address and update any DNS records for the provided domain with your new IP. By default, if the DNS record already matches your IP, no action is taken. This can be overriden using the `force` flag:

```sh
dnssync update -c "config.yaml" -p cloudflare -f -d home.example.com
```

More than one domain can be specified if you have multiple domains pointing back to your home that you want to update:

```sh
dnssync update -c "config.yaml" -p namesilo -d home.example.com -d home.mydomain.io
```

#### Configuration

A simple YAML config file is all that is needed to set up dnssync. Since this config file will contain API credentails for whichever DNS provider you are using, it is recommended that you make read access to this file as limited as possible.

```yaml
# Cloudflare Configuration
cloudflare:
    zone_id: <your zone id>
    token: <your API token>

# NameSilo Configuration
namesilo:
    key: <your API key>
```

### Roadmap

Currently, no other DNS providers are being considered to add to this tool, but that is only because I have no need for them personally. If others find this tool useful and want to suggest additional features, feel free to let me know!