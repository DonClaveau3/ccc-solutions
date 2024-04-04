# Build your own wc tool
A solution to John Crickett's Coding Challenge
https://codingchallenges.fyi/challenges/challenge-wc

Install:  
```
git clone git@github.com:DonClaveau3/ccwc.git
cd ccwc
go install
```

Examples:  

print newline, word, and byte counts for test.txt  
```
$ ccwc test.txt
  7145  58164 342190 test.txt
```  

print newline, word, and byte counts for stdin  
```
$ cat test.txt | ccwc
  7145  58164 342190
```

usage info  
```
$ ccwc --help
ccwc is a CLI app that imitates the behavior of wc.
This application is a solution to a coding challenge found here:
https://codingchallenges.fyi/challenges/challenge-wc

Usage:
  ccwc [path to file] [flags]

Flags:
  -c, --bytes   print the byte counts
  -m, --chars   print the character counts
  -h, --help    help for ccwc
  -l, --lines   print the newline counts
  -w, --words   print the word counts
```

References:  
- [Installing WSL](https://learn.microsoft.com/en-us/windows/wsl/install)
- [Connecting to GitHub w/ SSH](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/testing-your-ssh-connection)
- [Cobra](https://github.com/spf13/cobra) framework for creating CLIs in Go
- [Converting int64 to string](https://golangdocs.com/golang-int64-to-string-conversion)
- [Efficient line count algorithm by JimB](https://stackoverflow.com/a/24563853)
- [golang unit testing](https://golangdocs.com/golang-unit-testing)
- [Counting characters](https://golangdoc.github.io/pkg/1.8/bytes/index.html#example_Count)

 
Created using VS Code running in WSL: Ubuntu  
Tested against wc (GNU coreutils) 8.32  
