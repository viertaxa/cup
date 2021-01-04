# Configuration
To configure CUP, simply run `./cup configure` on Linux or Mac, or `.\cup.exe configure` on Windows. You will then be prompted for several bits of information that are required to get CUP running.

Alternately, advanced users may manually create or modify a YAML, TOML, or JSON file with the appropriate values.

## Updating a configuration

To update a configuration, just run the `configure` command again. You will be guided through the configuration you had previously set.

## Creating a configuration

Need to create a new configuration? Run the `configure` command and we'll guide you through setting up CUP. At the end, we'll save your configuration file to disk.

## Defaults

### Config File Location

The default configuration file path for CUP is `./cupconf.yaml`. If you wish to run CUP with a minimum of flags, stick to this location.

To override this, specify the `--config </path/to/config.{yaml,toml,json}>` flag

### Log Level

The default log level is INFO. We suggest leaving this as is, especially for `./cup configure`

To override this, specify the `--loglevel <level>` flag with your choice of the following: `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`

## Options

As you are guided through the configuration, you will encounter the following configuration options:

### Overall settings

- `Level` - This sets the log level as defined above. The same options apply, but we suggest just hitting the ENTER/RETURN key to accept the default value of `info`
- `API Base` - This sets the API base that we use for Cloudflare. Unless you really *really* know what you're doing, just hit the ENTER/RETURN key to accept the default.
- `API Key` - This is your Cloudflare API Key used to query and update the Cloudflare API. Enter it here. The default will not work and is just an example.

### IP determination Services

After setting up the overall settings, you will be guided through the default IP determination services.

If you don't have a reason to change them, it's best to just hit ENTER/RETURN and accept the defaults. There are three services entered by default, ipify, SeeIP, and WhatIsMyIPAddress. This should provide more than enough fault tolerance and validation that one site is not wrong.

- `Service Name` - The name of the service you will be using
- `API Base` - The API URL that we query to get your public IP

### Domains and Hosts

After setting up the IP determination services, you will be guided through setting up the domains and hosts you wish to update with your public IP when you run the `update` command.

You **must** enter your own values here, the defaults will not work.

- `Domain` - The domain you wish to update records on

Each *domain* represents the domain name you purchased and added to Cloudflare. For example, `example.com`, `example.co.uk`, `example.io`. do not enter hosts here, such as `www.example.com`, or `dynamic.example.com`. You'll configure the hosts later.

- `Hosts` - A comma separated list of each host record you wish to update the IP address of

There are several options here:

- To update the root domain, use the `@` keyword.
- To update a subdomain, enter just the subdomain. For example to update `www.example.com`, you would enter `www`
- To update a subdomain on a subdomain, just chain together the subdomains separated by `.`s. For example to update `www.foo.dynamic.example.com`, you would enter `www.foo.dynamic`
- to update a wildcard record, use the `*` keyword. For example to update your wildcard record `*.dynamic.example.com` you would enter `*.dynamic`