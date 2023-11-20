Output example.gif
Set Width 930
Set Height 670
Set Theme "Dracula"
Set FontFamily "Source Code Pro"
Set LetterSpacing 1

Hide
Type "export PS1='> '; export PATH=$PWD/bin:$PATH; mkdir bin"
Enter
Type "go build -o bin/example example/main.go"
Enter
Type "clear"
Enter
Type "example"
Enter
Show

Sleep 10s

Hide
Type "q"
Type "rm bin/example"
Enter
