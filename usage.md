# Tff Manual

tff (ten flying fingers). Most commands need root permissions (sudo). https://github.com/guettli/tff

```text
tff [command] [global flags] [command flags]
```

### Commands

* [tff combos](#tff-combos)
* [tff create-events-from-csv](#tff-create-events-from-csv)
* [tff csv](#tff-csv)
* [tff help](#tff-help)
* [tff print](#tff-print)
* [tff reply-combo-log](#tff-reply-combo-log)
* [tff validate](#tff-validate)

# Commands

## `tff combos`

Conntect to one or several evdev devices and modify the events according to your configuration. Needs root permissions.

```text
tff combos [flags] combos.yaml [device1 [device2 ...]]
```

### Command Flags

```text
  -d, --debug   Print debug output
  -h, --help    help for combos
```

## `tff create-events-from-csv`

Read events from csv file, and emit the events. This does not rewrite the events like 'replay-combo-log' does.

```text
tff create-events-from-csv events.csv
```

### Command Flags

```text
  -h, --help   help for create-events-from-csv
```

## `tff csv`

Conntect to one evdev device and print the events in csv format. Needs root permissions.

```text
tff csv [device]
```

### Command Flags

```text
  -h, --help   help for csv
```

## `tff help`

Help about any command

```text
tff help [command] [flags]
```

### Command Flags

```text
  -h, --help   help for help
```

## `tff print`

Conntect to one evdev device and print the events. Needs root permissions.

```text
tff print [device]
```

### Command Flags

```text
  -h, --help   help for print
```

## `tff reply-combo-log`

Replay a combo log. Emit the events from the given log. This is useful for debugging. Needs root permissions.

```text
tff reply-combo-log combos.yaml combos.log
```

### Command Flags

```text
  -h, --help   help for reply-combo-log
```

## `tff validate`

Validate the config file.

```text
tff validate combos.yaml
```

### Command Flags

```text
  -h, --help   help for validate
```
