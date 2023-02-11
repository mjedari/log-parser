## Log Parser

This app is a simple app to parse large log file *(for instance in 100 millions lines)* and persist them in database.
The important part would be we should read lines by chunk and store them in chunk async.

The architecture we use is clean architecture. Tests done is this machine:

- OS: Mac os
- CPU: M1 pro
- Memory: 16GB

I used `bufio.scan()` but in this senario we can just set one buffer size.

```go
func (p *LogParser) parseChunk() {
number := 1
for p.scanner.Scan() {
line := p.scanner.Text()
// process line
number++
}
fmt.Println("Finished: #", number)
if err := p.scanner.Err(); err != nil {
fmt.Println(err)
os.Exit(1)
}
}
```

The problem here is that out put is:

```markdown
Finished: # 100000001
Duration:  3.663172167s
```

And this is the problem. without buffer it returns :

```markdown
Finished: # 100000001
Duration:  4.818584666s
```

I decided to use reader: