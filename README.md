### gee - Search from the command line
---
gee is a small command line tool I wrote to become more familier with Go. It is a command line tool that can be used to quickly search for something using multiple search engines from the command line.

It can be configured using a engines.json file in the operating system conventional config file directory.

You can find out the storage directory using the `-dir` flag: `gee -dir`

An example would look like this:

```json
[
  {
    "command": "g",
    "searchString": "https://www.google.com/search?q=[QUERY]"
  },
  {
    "command": "b",
    "searchString": "https://search.brave.com/search?q=[QUERY]"
  },
  {
    "command": "ddg",
    "searchString": "https://duckduckgo.com/?q=[QUERY]"
  },
  {
    "command": "wiki",
    "searchString": "https://en.wikipedia.org/w/index.php?search=[QUERY]"
  },
  {
    "command": "yt",
    "searchString": "https://www.youtube.com/results?search_query=[QUERY]"
  }
]


```

Place the `[QUERY]` placeholder where you want the search query to go.

This allows me to run `gee -yt -b -d -wiki "Never gonna give you up"` for example, which will search the specified query with all the specified search engines in the default browser.

> [!NOTE]
> Keep in mind that this project is not really intended to be useful, just a fun exercise for me.
