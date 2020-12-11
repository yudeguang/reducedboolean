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

func IsTrue(s string) (bool, error) {
	original := s
	s = fmtStr(s)
	for {
		if len(s) == 1 {
			break
		}
		pre := s
		//处理and {"0 and 0", "0"},{"0 and 1", "0"},{"1 and 1", "1"},{"1 and 0", "0"}
		s = cleanAnd(s)
		if len(s) == 1 {
			break
		}
		//处理带括号的OR {"(0 or 0", "(0"},{"(0 or 1", "(1"},{"(1 or 1", "(1"},{"(1 or 0", "(1"}
		s = cleanOrHasParentheses(s)
		if len(s) == 1 {
			break
		}
		//处理清洗完成后的括号{"(0)", "0"},{"(1)", "1"}
		s = cleanParentheses(s)
		if len(s) == 1 {
			break
		}
		//再处理一次and {"0 and 0", "0"},{"0 and 1", "0"},{"1 and 1", "1"},{"1 and 0", "0"}
		s = cleanAnd(s)
		if len(s) == 1 {
			break
		}
		//处理普通的OR 	{"0 or 0", "0"},{"0 or 1", "1"},{"1 or 1", "1"},{"1 or 0", "1"}
		s = cleanOr(s)
		if s == pre {
			return false, errors.New("the input is invalid, please check:" + original)
		}
	}
	return s == "1", nil
}

//处理and {"0 and 0", "0"},{"0 and 1", "0"},{"1 and 1", "1"},{"1 and 0", "0"}
func cleanAnd(s string) string {
	for {
		pre := s
		for i := range kvPairsAnd {
			s = strings.Replace(s, kvPairsAnd[i].k, kvPairsAnd[i].v, -1)
		}
		//替换完一轮后，还和原值相等，则跳出循环
		if pre == s {
			break
		}
	}
	return s
}

//处理带括号的OR {"(0 or 0", "(0"},{"(0 or 1", "(1"},{"(1 or 1", "(1"},{"(1 or 0", "(1"}
func cleanOrHasParentheses(s string) string {
	for {
		pre := s
		for i := range kvPairsOrHasParentheses {
			s = strings.Replace(s, kvPairsOrHasParentheses[i].k, kvPairsOrHasParentheses[i].v, -1)
		}
		if pre == s {
			break
		}
	}
	return s
}

//处理普通的OR 	{"0 or 0", "0"},{"0 or 1", "1"},{"1 or 1", "1"},{"1 or 0", "1"}
func cleanOr(s string) string {
	for {
		pre := s
		for i := range kvPairsOr {
			s = strings.Replace(s, kvPairsOr[i].k, kvPairsOr[i].v, -1)
		}
		if pre == s {
			break
		}
	}
	return s
}

//处理清洗完成后的括号{"(0)", "0"},{"(1)", "1"}
func cleanParentheses(s string) string {
	for {
		pre := s
		for i := range kvPairsParentheses {
			s = strings.Replace(s, kvPairsParentheses[i].k, kvPairsParentheses[i].v, -1)
		}
		if pre == s {
			break
		}
	}
	return s
}

//规范化数据
func fmtStr(s string) string {
	s = strings.ToLower(s)
	for i := range kvPairsForFmt {
		for {
			pre := strings.Replace(s, kvPairsForFmt[i].k, kvPairsForFmt[i].v, -1)
			if pre == s {
				break
			}
			s = pre
		}
	}
	return s
}
