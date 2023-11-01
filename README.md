# cli-todo-app
Small cli to do list built with Go

# Installation

In order to install this program, you need to have golang installed, you can get it 
[here](https://go.dev/doc/install)

To install all the files, in our cli type:
```console
go get github.com/ssmf/cli-todo-app
```

Please note that the exe file must be in **the same directory** as the data.json file

# Dependencies

- [Goolkit](https://github.com/gookit/color)

# Usage
In order to use the program, we just type the name of it in our cli and pass a flag to it:
```console
./main.exe -help
```

The program works primarly on using flags, passing in either none or 1 argument.

The flag works by itself (meaning, not parsing any value) if the value we want to define is bool (e.g.. if the task is done), We pass in a value if The value is not bool and we want to specify it (e.g. -name [value])

here is a list of all flags available:

- -add -> Specifies if user wants to add a new task
- -rm -> Specifies if user wants to remove an existing task
- -name [value] -> Defines task name
- -done -> Defines if task is done, if not used, value is false
- -list -> Displays current list of todos
- -help -> Displays all commands and their roles

*(note that you can also get this list by parsing the -help flag)*

here are examples of using the flags:
```console
./main.exe -help
./main.exe -add -name "Read a book"
./main.exe -rm -name "Read a book"
./main.exe -list
./main.exe -add -name "Write some code" -done
./main.exe -name "Write some code" -done
```

If the value doesn't contain any spaces *(e.g. read instead of "read a book")* we do not need to use quotation marks