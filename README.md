# litaf
## Literals As Files - manipulate found literal string line in file, using line in other files

### Usage of litaf:

  -b    new line before
  
  -f1 string
        file with search line (default "file1.tmp")
        
  -f2 string
        file with action line (default "file2.tmp")
        
  -f3 string
        file with lines (default "data.tmp")
        
  -n    new line after
  
  -r    replace line
  
### Main use case:

echo, & other cmds, parse & remove double quote characters from strings, when sending them to stdout.

other cmds use regexp as input, to parse & exclude any special characters from search string.

this Go cmdline app is great for:
if you don't want to use regexp, & just use lines in files as literal strings for manipulation, especially to preserve double quotes.

### HowTo Make:
```shell
go build
```
### Example Usage:

See these 2 gists of mine where I use this app on linux to modify code:

forkware.yaml

forkware.sh

https://gist.github.com/roelds
