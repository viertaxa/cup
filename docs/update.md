# Updating Records
To update records, simply run `./cup update` on Linux or Mac, or `.\cup.exe update` on Windows.

## Defaults

### Config File Location

The default configuration file path for CUP is `./cupconf.yaml`. If you wish to run CUP with a minimum of flags, stick to this location.

To override this, specify the `--config </path/to/config.{yaml,toml,json}>` flag

### Log Level

The default log level is whatever was set during configuration.

To override this, specify the `--loglevel <level>` flag with your choice of the following: `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`

## Update Process

CUP will iterate on each domain you have configured, updating each host listed within it.

CUP will only update the content of an existing DNS record. It will not create one, or update any other setting, such as TTL.
