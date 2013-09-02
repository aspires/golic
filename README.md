# golic

An open source license generator for your projects.

## Installation

```bash
$ go get github.com/subosito/golic
```

## Usage

```
$ golic
Usage:
  golic [OPTIONS] LICENSE

Help Options:
  -h, --help              Show this help message

License Options:
  -y, --year=YEAR         License year
  -c, --copyright=NAME    Copyright name
  -u, --url=URL           URL
  -e, --email=EMAIL       E-Mail

General Options:
  -o, --output=FILE       Output file
  -l, --list              List supported licenses
```

The golic's [LICENSE](./LICENSE) file is created using command like:

```bash
$ golic -c "Alif Rachmawadi" -e "subosito@gmail.com" -o LICENSE MIT
```

## Todo

- Submit to http://mit-license.org/ when license type is MIT
- Support multiple authors and organization (via `-a`, `--authors`)
- Store configuration on `$HOME/.golicrc`, no need to write same info when generate a license
- Append to existing file, eg: `README.md`
- Build downloadable cross platforms binary
- Any other idea? :)

