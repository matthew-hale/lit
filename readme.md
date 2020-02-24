# lit
Lit is a simple program that parses markdown-formatted text, extracting 
text from code blocks and saving them to standalone files. Files are 
specified by name in the code block, and multiple code blocks with the 
same file name are combined into the same file.

## Usage
Markdown-formatted text input looks like this:

    # Title
    This is a description of my script. Here's where I started:

    ```test.sh
    #!/bin/sh
    
    echo this is code in my script

    ```

    After that, I did this thing:

    ```test.sh
    if [ -n "$thing" ]; then
        echo "$thing"
    fi

    ```

    But I encountered these issues:

    * a
    * b
    * c

    So I fixed it with this:

    ```test.sh
    if [ -z "$var" ]; then
        exit 1
    fi
    ```

Using `lit` via stdin:

```
$ ls
file.md
$ cat file.md | lit
test.sh written
test.sh written
test.sh written
$ ls
file.md     test.sh
$ cat test.sh
#!/bin/sh

echo this is code in my script

if [ -n "$thing" ]; then
    echo "$thing"
fi

if [ -z "$var" ]; then
    exit 1
fi
```

Using `lit` via file input flag:

```
$ lit -i file.md
test.sh written
test.sh written
test.sh written
$ ls
file.md     test.sh
$ cat test.sh
#!/bin/sh

echo this is code in my script

if [ -n "$thing" ]; then
    echo "$thing"
fi

if [ -z "$var" ]; then
    exit 1
fi
```

Overwriting an existing `test.sh` file:

```
$ ls
file.md     test.sh
$ cat test.sh
this file contains this line of text
$ lit -i file.md -f
test.sh written
test.sh written
test.sh written
$ ls
file.md     test.sh
$ cat test.sh
#!/bin/sh

echo this is code in my script

if [ -n "$thing" ]; then
    echo "$thing"
fi

if [ -z "$var" ]; then
    exit 1
fi
```

Spitting out `test.sh` into a different directory:

```
$ ls
file.md     directory
$ lit -i file.md -o directory
test.sh written
test.sh written
test.sh written
$ ls
file.md     directory
$ ls directory
test.sh
$ cat directory/test.sh
#!/bin/sh

echo this is code in my script

if [ -n "$thing" ]; then
    echo "$thing"
fi

if [ -z "$var" ]; then
    exit 1
fi
```

## Installation
Currently, you can install lit by cloning the repository and using `go 
install`:

```
$ cd lit/lit
$ go install
```

## Background
One workflow that I really like is something akin to a homegrown 
"literate programming," where I write out my thoughts like a narrative 
prose, and intersperse them with code snippets. The problem with doing 
this is that I don't have a convenient way to extract those code 
snippets into a resultant code file; I have to either manually 
copy/paste, or do some wonky scripting. Now, I can do everything like I 
normally do in a markdown file, but at the end I can run it through lit 
and have the code files, too.

## Roadmap

* ~~Add flags for append/overwrite (currently only appends)~~ - complete
* ~~Add output directory flag (currently spits everything out to current 
directory)~~ - complete
* ~~Add input file flag (currently only accepts stdin)~~ - complete
* Add a flag that will, instead of spitting out separate script blocks, 
take the first named script block as a file name and spit out a script 
file containing all of the content of the file, with the markdown text 
included as comments, and the script block delimiters removed. Will 
also necessitate a comment sign parameter, for file types whose 
comments don't start with a '#'.
