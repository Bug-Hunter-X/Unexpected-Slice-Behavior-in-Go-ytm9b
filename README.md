# Unexpected Slice Behavior in Go

This repository demonstrates a common pitfall in Go programming involving slices.  Slices are references, not copies, which means modifications to the underlying array affect all slices pointing to it.  This example showcases unexpected behavior when a function returns a slice created from another slice.

## Problem
The `myFunc` function doubles the values of a slice.  However, changes to the original slice after the function call *do not* affect the returned slice, which is unexpected for programmers accustomed to pass-by-value semantics.

## Solution
To create a true copy of the slice, use the `copy` function or create a new slice with `append`.  This ensures modifications to the original slice don't affect the copied slice.