# pass

Password hashing and key derivation utilities.

## Usage
```
Password hashing and key derivation utilities.

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
$ echo -n secret | pass scrypt --salt=deadbeef --length=40
2d436feef50b02740db2ab5069db54da21b20b600c45d2a6e29cbb26beb1cf0137cd7170dc609fbf
$ echo -n secret | pass bcrypt
$2a$10$0U47Dz5Eu.clOcUkj8LlEeBuawib.7qu7EQ9W8U7NEVR7o9quYQki
$ echo -n secret | pass bcrypt --verify='$2a$10$0U47Dz5Eu.clOcUkj8LlEeBuawib.7qu7EQ9W8U7NEVR7o9quYQki'
OK: verify string matches given password
$ echo -n secret | pass bcrypt --verify='foobar'
ERROR: verify string is not the hash of the given password
```
