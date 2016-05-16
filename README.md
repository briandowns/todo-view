# todo-view

todo-view is a small tool to extract information from TODO comments from source code.  The TODO comments need to be formatted in a specific way for `todo-view` to find them, reason about them, and display them.

## Format

```
TODO(<user>) <message> <timestamp> <weight>

Example:

TODO(briandowns) this is an example todo format 2016-05-13T18:54 4
```

## Weighting

```sh
todo-view weights:

1       Extremely important
2       Somewhat important
3       Moderately important
4       Slightly important
5       Not at all important
```

## Sort By

* User
* File
* Timestamp
* Weight

## Usage

By issuing the parse sub-command, you can view the todo data from all source files under the root that the command was run in.  If the `-d` flag issued, the output will be displayed in decending order rather than the default ascending order.

### Parse By User

```sh
➜  $ todo-view parse by-user

Todo's by user:

briandowns      README.md        this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4
briandowns      bin/todo-view    this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4
briandowns      command/show.go  this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4 
```

### Export CSV

Sorting in CSV isn't supported currently and I'm not sure it will be since whatever it's being imported to will most likely have that capability.

```sh
➜  $ todo-view export csv

todo-view export: cvs

briandowns,README.md,this is an example todo format,2016-05-13 18:54:00 +0000 UTC,4
briandowns,bin/todo-view,this is an example todo format,2016-05-13 18:54:00 +0000 UTC,4
briandowns,command/parse.go,have this controlled by a CLI flag,2016-05-13 16:14:00 +0000 UTC,2
briandowns,command/show.go,this is an example todo format,2016-05-13 18:54:00 +0000 UTC,4
```

## Development

```sh
$ mkdir bin
$ make
```
