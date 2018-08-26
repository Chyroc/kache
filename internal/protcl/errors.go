/*
 * MIT License
 *
 * Copyright (c)  2018 Kasun Vithanage
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package protcl

import "fmt"

// 协议错误

// 转 int 错误
type ErrCastFailedToInt struct {
	Val interface{}
}

func (e *ErrCastFailedToInt) Error() string {
	return fmt.Sprintf("%s: error casting %v to int", ERR, e.Val)
}

// 错误的类型
type ErrWrongType struct {
}

func (ErrWrongType) Error() string {
	return fmt.Sprintf("%s: invalid operation against key holding invalid type of value", WRONGTYP)
}

// 给 error 包一层
type ErrGeneric struct {
	Err error
}

func (e *ErrGeneric) Error() string {
	return fmt.Sprintf("%s: %s", ERR, e.Err)
}

// 参数个数错误
type ErrWrongNumberOfArgs struct {
	Cmd string
}

func (e *ErrWrongNumberOfArgs) Error() string {
	return fmt.Sprintf("%s: %s has wrong number of arguments", WRONGTYP, e.Cmd)
}

// 未知的command 错误
type ErrUnknownCommand struct {
	Cmd string
}

func (e *ErrUnknownCommand) Error() string {
	return fmt.Sprintf("%s: unknown command %s", ERR, e.Cmd)
}
