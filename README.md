# Hael

Hael is a CLI for giving you a nice and sweet message to remind you of how amazing you are. It is also designed to be used with Nix, meaning you can use it while leaving no trace of its existence on your computer.

## Usage

```
nix run github:isabelroses/hael
```

```
Usage: hael [FLAGS]

Flags:
  -h, --help    Print this help message
  -r, --random  Print a random euphoric message
  -p, --pickup  Print a pick-me-up message

Environment Variables:
  HAEL_GENDER  'fem', 'masc', 'all' or 'genderless'

Examples:
  hael -r
  hael --pickup

Help:
  https://github.com/isabelroses/hael
```

If you don't have to worry about someone finding out, you can also install it globally. Locate the appropriate binary for your computer from the [latest release](https://github.com/isabelroses/hael/releases/latest/).

## Configuration

You can set the `HAEL_GENDER` environment variable to either `fem`, `masc`, or `both` to configure the responses.

```bash
export HAEL_GENDER="fem"
```

