# lit
A program to extract code blocks from text into standalone code files.

## Usage
Code blocks in a markdown file can be classed with a filename, like 
this:

    ```test.sh
    #!/bin/sh
    
    echo this is code written within a code block
    ```

Multiple code blocks classed with the same filename are combined into 
an output file of the same name.

Run lit (currently only takes stdin; filepath support is TODO):

```
$ cat file.md | lit
```

## Installation
Currently, you can install lit by cloning the repository and using `go 
install`:

```
$ cd lit/lit
$ go install
```

## Examples
Say I'm explaining a particular operation in a shell language, and I 
want to demonstrate with a code block, like so:

```file.sh
echo this is a code block with my demonstration in it
# This code will dump to a file named file.sh
```

If I want to explain another completely irrelevant piece of code, but
still want to generate a script file with this code in it, I can make a
second code block with a different name:

```file2.sh
echo this is a second code block, which will dump to file2.sh
```

I can revisit the concept I talked about in the first code block by
just making another code block with the same file name; multiple code
blocks with the same file name all dump to a single file.

```file.sh
# This comment gets appended to the original file.sh file
```

Finally, say I've got a small snipped of code that's really only in a
code block because I didn't want to inline it; I can just present it
like normal, without giving it a file name, and its contents will be
skipped by lit.

Example:

```
# This unnamed code block will not appear in lit's parser.
echo Goodbye, world
```

This entire readme.md file can be processed by lit, and will spit out 2 
files total:

+ file.sh
+ file2.sh

It won't spit out test.sh, because it only parses real fenced code 
blocks that start at the beginning of the line. The test.sh example was 
just that - an example of what it should look like in your file.

## Background
One workflow that I really like is something akin to a homegrown 
"literate programming," where I write out my thoughts like a narrative 
prose, and intersperse them with code snippets. The problem with doing 
this is that I don't have a convenient way to extract those code 
snippets into a resultant code file; I have to either manually 
copy/paste, or do sone wonky scripting. Now, I can do everything like I 
normally do in a markdown file, but at the end I can run it through lit 
and have the code files, too.

## Roadmap

+ Add output directory flag (currently spits everything out to current 
directory)
+ Add input file flag (currently only accepts stdin)
