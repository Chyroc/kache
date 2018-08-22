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

package cmds

import (
	"github.com/kasvith/kache/internal/db"
	"github.com/kasvith/kache/internal/protcl"
)

// exists 命令
func Exists(d *db.DB, args []string) *protcl.Message {
	if len(args) != 1 {
		return protcl.NewMessage(nil, &protcl.ErrWrongNumberOfArgs{Cmd: "exists"})
	}
	found := d.Exists(args[0])
	return protcl.NewMessage(protcl.NewIntegerReply(found), nil)
}

// del 命令
func Del(d *db.DB, args []string) *protcl.Message {
	deleted := d.Del(args)
	return protcl.NewMessage(protcl.NewIntegerReply(deleted), nil)
}
