// Copyright 2020 reducedboolean Author(https://github.com/yudeguang/reducedboolean). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/reducedboolean.
package reducedboolean

type kvPair struct {
	k string
	v string
}

//小括号
var kvPairsParentheses = []kvPair{
	{"(0)", "0"},
	{"(1)", "1"}}

//格式化目标数据，防止以下相关情况
var kvPairsForFmt = []kvPair{
	{"  ", " "},
	{"( ", "("},
	{" )", ")"},
	{")or", ") or"},
	{"or(", "or ("},
	{")and", ") and"},
	{"and(", "and ("}}

//有括号的OR
var kvPairsOrHasParentheses = []kvPair{
	{" or 0 or 0)", " or 0)"},
	{" or 0 or 1)", " or 1)"},
	{" or 1 or 1)", " or 1)"},
	{" or 1 or 0)", " or 1)"},
	{"(0 or 0)", "0"},
	{"(0 or 1)", "1"},
	{"(1 or 1)", "1"},
	{"(1 or 0)", "1"}}

//先处理有括号的OR
var kvPairsOr = []kvPair{
	{"0 or 0", "0"},
	{"0 or 1", "1"},
	{"1 or 1", "1"},
	{"1 or 0", "1"}}

//再处理有括号的and
var kvPairsAnd = []kvPair{
	{"0 and 0", "0"},
	{"0 and 1", "0"},
	{"1 and 1", "1"},
	{"1 and 0", "0"}}
