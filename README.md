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

## Usage

```sh
➜  todo-view git:(master) ✗ bin/todo-view parse by-user

Todo's by user:

briandowns      README.md        this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4
briandowns      bin/todo-view    this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4
briandowns      command/show.go  this is an example todo format         2016-05-13 18:54:00 +0000 UTC   4 
```
