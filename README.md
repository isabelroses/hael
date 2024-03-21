## Trans rights are human rights

This is a cli app is here to give you a nice and sweet message to remind you of how amazing you are.
It is also designed to leave no trace of its existence on your computer, so you can use it without worrying about someone finding out.

### Usage

For a 0 trace usage, you can use it with nix.
```bash
export HAEL_GENDER="fem" # this can be "fem", "masc" or "both"

nix run github:isabelroses/hael

rm -rf ~/.cache/nix # to remove the cache, rember 0 trace
```

If you don't have to worry about someone finding out, you can also install it globally.
You can find your appopriate release in the [releases tab](https://github.com/isabelroses/hael/releases).

```bash
export HAEL_GENDER="masc" # see the previous note

hael
```

### Help
```
Usage: hael [FLAGS]

Flags:
  -h, --help    Print this help message
  -r, --random  Print a random euphoric message
  -p, --pickup  Print a pick-me-up message

Environment Variables:
  HAEL_GENDER  'fem', 'masc' or 'both'

Examples:
  hael -r
  hael --pickup

Help:
  https://github.com/isabelroses/hael
```
