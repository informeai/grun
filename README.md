# grun
Execute commands from formulas
### Build
```
make build
```
### Usage
Initialize repository local.
```
grun init
```
*Generate formula*
```
grun generate -n [FORMULA_NAME]
```
*Add command to formula*
```
grun add -f test -n NAME_COMMAND -d 'DESCRIPTION' -a 'COMMAND_TO_EXECUTE'
grun add -f my_formula -n custom_echo -d 'this custom echo' -a 'echo custom' #eg.
```

*Run command*
```
grun run FORMULA:NAME_COMMAND
grun run my_formula:custom_echo #eg.
```