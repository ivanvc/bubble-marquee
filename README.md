# Marquee Bubble

<a href="https://pkg.go.dev/github.com/ivanvc/bubble-marquee?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>

A simple [BubbleTea] bubble component to generate marquees.

![Made with VHS](https://vhs.charm.sh/vhs-3kYqJAkAcOZpAlmmLFRE36.gif)

[BubbleTea]: https://github.com/charmbracelet/bubbletea

## Usage

First, initialize a new marquee, then set the text to display, and optionally
the width. If width is not defined, it's assumed from the text's length.

```go
func NewMyModel() MyModel {
    m := marquee.New()
    m.SetText("Hello World")
    m.SetWidth(100) // optional
    return MyModel{marquee: m}
}
```

It supports two scrolling directions, left and right.

```go
m.ScrollDirection = marquee.Right
```

You can also specify the scrolling speed (defaults to 250ms).

```go
m.ScrollSpeed = 50 * time.Millisecond
```

And you can make it continuous:

```go
m.SetContinuous(true)
```

Then, call scroll in your `Init()`.

```go
func(m MyModel) Init() tea.Cmd {
  return m.marquee.Scroll
}
```

That's it. Check the documentation and the [example](example/) directory for
more detailed usage.

## License

See [LICENSE](LICENSE).
