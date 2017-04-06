# todo-view

todo-view is a small tool to extract information from TODO comments from source code.  The TODO comments need to be formatted in a specific way for `todo-view` to find them, reason about them, and display them.

## Format

```
TODO(<user>) <message> <timestamp> <weight>

Example:

TODO(briandowns) this is an example todo format 2016-05-13T18:54 4
```

## Priorities

```sh
todo-view priorities:

1       Extremely important
2       Somewhat important
3       Moderately important
4       Slightly important
5       Not at all important
```

## Web

```sh
todo-view web port 5000

todo-view web:

Serving on http://localhost:5000
```

## Sorting

* User
* File
* Timestamp
* Priority

## Usage

By issuing the parse sub-command, you can view the todo data from all source files under the root that the command was run in.  If the `-d` flag issued, the output will be displayed in decending order rather than the default ascending order.

### Sort By User

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

### Export to Jira Import Format

This is just CSV but formatted for import into Jira.  Jira requires a header to know how to match columns.

```sh
$ todo-view export jira

todo-view export: jira

Summary,Assignee,Reporter,Priority
this is an example todo format,briandowns,briandowns,4
this is an example todo format,briandowns,briandowns,4
have this controlled by a CLI flag,briandowns,briandowns,2
this is an example todo format,briandowns,briandowns,4
```

### Export to Jira Table Comment Notation

A table can be added to a Jira comment with this notation.  Output includes headers.

```sh
$ todo-view export jira-table

todo-view export: jira-table

||Summary||Assignee||Reporter||Priority||
|this is an example todo format|briandowns|briandowns|4|
|this is an example todo format|briandowns|briandowns|4|
```

## Development

```sh
$ mkdir bin
$ make
```
