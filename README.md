# Go: Panic on Nil Map Access

This repository demonstrates a common error in Go: panicking due to accessing a key in a nil map.  The `bug.go` file contains the problematic code, while `bugSolution.go` offers a safe alternative.

## Problem

In Go, attempting to access a key in an uninitialized (nil) map will cause a runtime panic. This is because the map doesn't exist, so there is no data to retrieve.

## Solution

The solution involves checking if the map is nil before attempting to access it.  If it's nil, you can either handle it gracefully (e.g., return a default value) or initialize the map.