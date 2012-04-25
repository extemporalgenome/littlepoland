// Copyright 2012 Kevin Gillette. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package littlepoland

import (
	"fmt"
	"strconv"
)

type ErrInput struct {
	Message string
	ArgNum  int
	Arg     string
}

func (e ErrInput) Error() string {
	return fmt.Sprintf("argument [%d] %q: %s", e.ArgNum, e.Arg, e.Message)
}

func Run(args []string) ([]float64, error) {
	stack := make([]float64, 0, len(args))
	for i, str := range args {
		b := str[0]
		if b >= '0' && b <= '9' {
			if n, err := strconv.ParseFloat(str, 64); err != nil {
				return nil, err
			} else {
				stack = append(stack, n)
			}
			continue
		}
		l := len(stack)
		if l < 2 {
			return nil, ErrInput{"stack underrun", i, str}
		}
		m, n := stack[l-2], stack[l-1]
		var o float64
		switch str {
		case "+":
			o = m + n
		case "-":
			o = m - n
		case "*":
			o = m * n
		case "/":
			o = m / n
		default:
			return nil, ErrInput{"unrecognized operator", i, str}
		}
		stack[l-2] = o
		stack = stack[:l-1]
	}
	return stack, nil
}
