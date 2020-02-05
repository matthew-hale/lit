# lit
A program to extract code blocks from text into standalone code files.

Code blocks can be classed with a filename, like this:

```file.sh
```file.sh
#!/bin/sh

echo this is code written within a code block
```
```

Multiple code blocks classed with the same filename are combined into 
an output file of the same name.
