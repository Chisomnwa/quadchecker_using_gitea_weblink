# Quadchecker

## Overview
**Quadchecker** is a Go program that identifies which rectangle pattern function generated a given output. It depends on the **Quad** project, which provides functions (`QuadA`, `QuadB`, `QuadC`, `QuadD`, `QuadE`) that print rectangles with specific patterns.  

Quadchecker reads a rectangle (from standard input), compares it against all possible Quad outputs, and outputs which Quad function(s) match the input.

---

## Repository Structure

```text
.
├── go.mod # Go module definition
└── main.go # Main program to check rectangle patterns
```

```yaml
## How It Works

1. Input a rectangle via **stdin** (e.g., paste the output of a Quad function).  
2. The program calculates the **width and height** of the rectangle.  
3. Each Quad function (`QuadA` to `QuadE`) is executed with the calculated dimensions, and the outputs are captured.  
4. Quadchecker compares the input with all possible outputs.  
5. If a match is found, it prints the function name(s) and dimensions in the format:

[quadA] [width] [height] || [quadC] [width] [height]
```
If no match is found, it prints: 
Not a quad function

```yaml
## Usage Example

```bash
go run main.go
```

***Paste a rectangle output, e.g.:***

```lua
o---o
|   |
o---o
```

***output:***

```css
[quadA] [5] [3]
```

## Learning Outcomes

Handling stdin input in Go.

Capturing stdout from functions programmatically.

Using maps and string comparison for pattern matching.

Understanding modular code and package dependencies (quad/piscine).

Applying algorithmic thinking to compare structured outputs.