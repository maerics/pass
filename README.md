# pass

Password hashing and key generation utilities.

## Usage
```
Password hashing and key generation utilities.

Usage:
  pass [command]

Available Commands:
  argon2      Perform password key derivation using "argon2".
  help        Help about any command
  pbkdf2      Perform password key derivation using "pbkdf2".
  scrypt      Perform password key derivation using "scrypt".

Flags:
  -h, --help   help for pass

Use "pass [command] --help" for more information about a command.
```

## Examples
```sh
$ echo -n secret | pass argon2
dd29d0fa6c527596259c538bc8b089a97158a9d61fcde9798c00074ffc76248f
$ echo -n secret | pass scrypt --salt=deadbeef --length=64
2d436feef50b02740db2ab5069db54da21b20b600c45d2a6e29cbb26beb1cf0137cd7170dc609fbf64f62450d932a6a635ebc5f9cd4ec74169893cf597c4ac3e
```
