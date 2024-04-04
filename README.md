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
  ccwc [flags] [path to file]

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
- [SBCS vs MBCS](https://learn.microsoft.com/en-us/cpp/c-runtime-library/single-byte-and-multibyte-character-sets?view=msvc-170)
- [Glossary of globalization concepts](https://learn.microsoft.com/en-us/globalization/reference/glossary)

 
Created during 6 evening sessions (mostly weekends) after putting the kid to bed
Built on VS Code running in WSL: Ubuntu  
Tested against wc (GNU coreutils) 8.32  

Personal notes:
- This was my first project in Go  
- This was my first non-school project published to a public repository  
- I wanted to set up WSL on my Windows PC after reading [DHH's recent endorsement](https://world.hey.com/dhh/vscode-wsl-makes-windows-awesome-for-web-development-9bc4d528). This project provided an opportunity to kick the tires a bit
- Through the process, I had a few questions that I would like to explore in the future:
  - What should I know about software licenses?
  - How do I make a manual entry (```man ccwc```) for an app?
  - How would I create a test against the built app?
    
