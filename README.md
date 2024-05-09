# sloglint

[![checks](https://github.com/go-simpler/sloglint/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/sloglint/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/sloglint.svg)](https://pkg.go.dev/go-simpler.org/sloglint)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/sloglint)](https://goreportcard.com/report/go-simpler.org/sloglint)
[![codecov](https://codecov.io/gh/go-simpler/sloglint/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/sloglint)

A Go linter that ensures consistent code style when using `log/slog`.

## 📌 About

The `log/slog` API allows two different types of arguments: key-value pairs and attributes.
While people may have different opinions about which one is better, most seem to agree on one thing: it should be consistent.
With `sloglint` you can enforce various rules for `log/slog` based on your preferred code style.

## 🚀 Features

* Enforce not mixing key-value pairs and attributes (default)
* Enforce using either key-value pairs only or attributes only (optional)
* Enforce not using global loggers (optional)
* Enforce using methods that accept a context (optional)
* Enforce using static log messages (optional)
* Enforce using constants instead of raw keys (optional)
* Enforce a single key naming convention (optional)
* Enforce putting arguments on separate lines (optional)
* Prevent usage of reserved keys in key-value pairs (optional)

## 📦 Install

`sloglint` is integrated into [`golangci-lint`][1], and this is the recommended way to use it.

To enable the linter, add the following lines to `.golangci.yml`:

```yaml
linters:
  enable:
    - sloglint
```

Alternatively, you can download a prebuilt binary from the [Releases][2] page to use `sloglint` standalone.

## 📋 Usage

Run `golangci-lint` with `sloglint` enabled.
See the list of [available options][3] to configure the linter.

When using `sloglint` standalone, pass the options as flags of the same name.

### No mixed arguments

The `no-mixed-args` option causes `sloglint` to report mixing key-values pairs and attributes within a single function call:

```go
slog.Info("a user has logged in", "user_id", 42, slog.String("ip_address", "192.0.2.0")) // sloglint: key-value pairs and attributes should not be mixed
```

It is enabled by default.

### Key-value pairs only

The `kv-only` option causes `sloglint` to report any use of attributes:

```go
slog.Info("a user has logged in", slog.Int("user_id", 42)) // sloglint: attributes should not be used
```

### Attributes only

In contrast, the `attr-only` option causes `sloglint` to report any use of key-value pairs:

```go
slog.Info("a user has logged in", "user_id", 42) // sloglint: key-value pairs should not be used
```

### No global

Some projects prefer to pass loggers as explicit dependencies.
The `no-global` option causes `sloglint` to report the usage of global loggers.

```go
slog.Info("a user has logged in", "user_id", 42) // sloglint: global logger should not be used
```

Possible values are `all` (report all global loggers) and `default` (report only the default `slog` logger).

### Context only

Some `slog.Handler` implementations make use of the given `context.Context` (e.g. to access context values).
For them to work properly, you need to pass a context to all logger calls.
The `context-only` option causes `sloglint` to report the use of methods without a context:

```go
slog.Info("a user has logged in") // sloglint: InfoContext should be used instead
```

Possible values are `all` (report all contextless calls) and `scope` (report only if a context exists in the scope of the outermost function).

### Static messages

To get the most out of structured logging, you may want to require log messages to be static.
The `static-msg` option causes `sloglint` to report non-static messages:

```go
slog.Info(fmt.Sprintf("a user with id %d has logged in", 42)) // sloglint: message should be a string literal or a constant
```

The report can be fixed by moving dynamic values to arguments:

```go
slog.Info("a user has logged in", "user_id", 42)
```

### No raw keys

To prevent typos, you may want to forbid the use of raw keys altogether.
The `no-raw-keys` option causes `sloglint` to report the use of strings as keys
(including `slog.Attr` calls, e.g. `slog.Int("user_id", 42)`):

```go
slog.Info("a user has logged in", "user_id", 42) // sloglint: raw keys should not be used
```

This report can be fixed by using either constants...

```go
const UserId = "user_id"

slog.Info("a user has logged in", UserId, 42)
```

...or custom `slog.Attr` constructors:

```go
func UserId(value int) slog.Attr { return slog.Int("user_id", value) }

slog.Info("a user has logged in", UserId(42))
```

> [!TIP]
> Such helpers can be automatically generated for you by the [`sloggen`][4] tool. Give it a try too!

### Key naming convention

To ensure consistency in logs, you may want to enforce a single key naming convention.
The `key-naming-case` option causes `sloglint` to report keys written in a case other than the given one:

```go
slog.Info("a user has logged in", "user-id", 42) // sloglint: keys should be written in snake_case
```

Possible values are `snake`, `kebab`, `camel`, or `pascal`.

### Arguments on separate lines

To improve code readability, you may want to put arguments on separate lines, especially when using key-value pairs.
The `args-on-sep-lines` option causes `sloglint` to report 2+ arguments on the same line:

```go
slog.Info("a user has logged in", "user_id", 42, "ip_address", "192.0.2.0") // sloglint: arguments should be put on separate lines
```

This report can be fixed by reformatting the code:

```go
slog.Info("a user has logged in",
    "user_id", 42,
    "ip_address", "192.0.2.0",
)
```

### Arguments with forbidden keys

To ensure users don't use reserved words in key-value pairs, you may want to enforce rules to avoid misuse.
The `forbidden-key` option causes `sloglint` to report usages of forbidden keys. These keys could be typically controlled by slog handlers and its usage should be preserved to avoid key duplication.

```go
slog.Info("a user has logged in", "forbidden_key", 49) // sloglint: "forbidden_key" key is forbidden and should not be used
```

Keys to forbid can either be comma separated or passed as separate `forbidden-key` options. 

To avoid key collisions and duplication when using `slog.JSONHandler` or `slog.TextHandler`,
you might want to use `--forbidden-key=source,msg,level,time` to avoid using reserved keys.

[1]: https://golangci-lint.run
[2]: https://github.com/go-simpler/sloglint/releases
[3]: https://golangci-lint.run/usage/linters/#sloglint
[4]: https://github.com/go-simpler/sloggen
