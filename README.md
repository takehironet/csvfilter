# csvfilter
A simple CSV filter to pick columns up from stdin

## How to use

```
csvfilter -i <input delimiter> -o <output delimiter> -c <columns list>
```

Example:

```
echo 'A,B,C,D,E' | csvfilter -i "," -o "," -c "1 3 5 2 4 1"
```

You get `"A","C","E","B","D","A"`.

## License

BSD 3-Clause "New" or "Revised" License. See `LICENSE` file.
