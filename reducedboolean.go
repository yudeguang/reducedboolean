// Copyright 2020 reducedboolean Author(https://github.com/yudeguang/reducedboolean). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/reducedboolean.
package reducedboolean

import (
	"errors"
	"strings"
)

type kvPair struct {
	k string
	v string
}

//待替换的相关字符串
var kvPairs = []kvPair{
	{")and", ") and"},
	{"and(", "and ("},
	{")or", ") or"},
	{"or(", "or ("},
	{"0 and 0", "0"},
	{"1 and 1", "1"},
	{"0 and 1", "0"},
	{"1 and 0", "0"},
	{"0 or 0", "0"},
	{"1 or 1", "1"},
	{"0 or 1", "1"},
	{"1 or 0", "1"},
	{"(0)", "0"},
	{"(1)", "1"},
	{"  ", " "},
	{"( ", "("},
	{" )", ")"}}

//计算最简单的逻辑表达式: 1 and 0 and 0 and (1 or 0 or 1)
//类似于在MYSQL中执行对应SQL: SELECT 1 and 0 and 0 and (1 or 0 or 1)
//字符串中只能包含：0,1,and,or,(,)这几个字符串
func IsTrue(s string) (bool, error) {
	original := s
	s = strings.ToLower(s)
	for {
		pre := s
		for i := range kvPairs {
			s = strings.Replace(s, kvPairs[i].k, kvPairs[i].v, -1)
			if len(s) == 1 {
				return s == "1", nil
			}
		}
		//处理一轮后，还和上一轮一样，那么就不合法
		if pre == s {
			return false, errors.New("the input is invalid, please check:" + original)
		}
	}
}
