# lit
A program to extract code blocks from text into standalone code files.

## Usage
Code blocks can be classed with a filename, like this (read the raw 
readme.md for a better look):

    ```test.sh
    #!/bin/sh
    
    echo this is code written within a code block
    ```

Multiple code blocks classed with the same filename are combined into 
an output file of the same name.

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
echo this is a second code block, which will dump to a separate file
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

This entire readme.md file can be processed by lit, and will spit out 3 
files total:

+ test.sh
+ file.sh
+ file2.sh
