# gocov-merger
`gocov-merger` merges many golang coverprofiles into one coverprofile.

## Installation
```
go get github.com/amobe/gocov-merger
```

## Usage
From the command line, `gocov-merger` can merge many coverprofiles. By default, it prints the merge result to `stdout`.
```
gocov-merger [options] PATH ...
```
The result can be use for gocov, gocov-xml, and gocov-html

### Available options
```
  -q         quite mode
  -o file    specific output file
  -h         help page
```

## Example
```
cd project_a && go test -coverprofile=cover.out ./...
cd project_b && go test -coverprofile=cover.out ./...

gocov-merger -o merged.out project_a/cover.out project_b/cover.out ...
> Merged 2 reports to merged.out
```

## Related tools and services

[gocov](https://github.com/axw/gocov):
Coverage reporting tool for The Go Programming Language.

[gocov-html](https://github.com/matm/gocov-html):
A simple helper tool for generating HTML output from gocov.

[gocov-xml](https://github.com/AlekSi/gocov-xml):
A simple helper tool for generating XML output in Cobertura format for CIs like Jenkins and others from gocov.
