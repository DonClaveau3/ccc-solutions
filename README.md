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
- This was my first project in Go. I wanted to try a new language and Go is popular. Some things that stood out:
  - nil instead of null
  - types go after names in declarations (like Typescript)
  - File.Stat() is great for byte counts on files, but can't be trusted for stdin. Seems to be a timing thing
  - unused variables prevent compiling. this was irritating when methods returned multiple values, but I was only interested in one of them. it also led to a fair amount of clutter related to error handling. I suspect this is either a safety feature of the language or I just didn't look hard enough for alternative approaches
  - there seem to be strong conventions about where to clone go projects relative to GOPATH. this felt odd as I'm used to keeping all projects in /dev or /src
  - I initially intended a monorepo that would hold multiple coding challenge projects. I abandoned this approach after the go module seemed to want to be at the git root.
- This was my first non-school project published to a public repository. I created this account several years ago with a desire to contribute to open source projects on the side. That hasn't happened because life. Career development energy mostly goes to watching Pluralsight courses related to work projects. It was refreshing to tackle a challenge originating outside of work as it allowed me to study concepts I would not have otherwise.   
- I wanted to set up WSL on my Windows PC after reading [DHH's recent endorsement](https://world.hey.com/dhh/vscode-wsl-makes-windows-awesome-for-web-development-9bc4d528). This project provided an opportunity to kick the tires a bit
- Through the process, I had a few questions that I would like to explore in the future:
  - What should I know about software licenses?
  - How do I make a manual entry (```man ccwc```) for an app?
  - How could I create tests at the interface level?
  - How would I create a performance benchmarking tool to test my solution against others?
  - What could I learn from reviewing other people's solutions?
- There were some strange formatting behaviors in the wc tool. For example, sometimes the output would begin with 2 spaces and sometimes with 3 spaces. Why is that? Why are there 2 spaces between newline count and word count in the default output, but only 1 space between every other part? For the sake of this project, I imitated what I saw. But it would be interesting to investigate the source
- I noticed that the challenge doesn't include all features of wc. Notably missing are the -L flag and multiple file inputs.
    
