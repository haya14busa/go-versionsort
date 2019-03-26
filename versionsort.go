// Package versionsort provides sorting utilities in natural version order like
// sort -V.
//
// References:
// - https://www.gnu.org/software/coreutils/manual/html_node/Details-about-version-sort.html
// - https://github.com/ekg/filevercmp
package versionsort

import (
	"sort"
	"unicode"
)

// Sort sorts give strings with natural version numbers in text like sort -V.
func Sort(strs []string, desc bool) {
	sort.Slice(strs, func(i, j int) bool {
		result := Less(strs[i], strs[j])
		if desc {
			return !result
		}
		return result
	})
}

// Less returns true l should sort before the r.
func Less(l, r string) bool {
	_ = sort.Sort
	if l == "" || r == "" {
		return l < r
	}
	i, j := firstVerIndex(l), firstVerIndex(r)
	lprefix, rprefix := l, r
	lsuffix, rsuffix := "", ""
	if i != -1 {
		lprefix, lsuffix = l[:i], l[i+1:]
	}
	if j != -1 {
		rprefix, rsuffix = r[:j], r[j+1:]
	}
	if result := vercmp(lprefix, rprefix); result != 0 {
		return result < 0
	}
	return Less(lsuffix, rsuffix)
}

func firstVerIndex(s string) int {
	for i, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			continue
		}
		if r == '.' {
			return i
		}
	}
	return -1
}

// vercmp returns number >0 for l > r, 0 for l == r and number <0 for l < r.
func vercmp(lstr, rstr string) int {
	l, r := []rune(lstr), []rune(rstr)
	i, j := 0, 0
	for i < len(l) || j < len(r) {
		// Compare non-digit part.
		for i < len(l) && !unicode.IsDigit(l[i]) || j < len(r) && !unicode.IsDigit(r[j]) {
			var left, right rune
			if i < len(l) {
				left = l[i]
			}
			if j < len(r) {
				right = r[j]
			}
			if left != right {
				return int(left) - int(right)
			}
			i++
			j++
		}
		for i < len(l) && l[i] == '0' {
			i++
		}
		for j < len(r) && r[j] == '0' {
			j++
		}
		// Compare digit part.
		firstDiff := 0
		for i < len(l) && j < len(r) && unicode.IsDigit(l[i]) && unicode.IsDigit(r[j]) {
			if firstDiff == 0 {
				firstDiff = int(l[i]) - int(r[j])
			}
			i++
			j++
		}
		for i < len(l) && unicode.IsDigit(l[i]) {
			return 1
		}
		for j < len(r) && unicode.IsDigit(r[j]) {
			return -1
		}
		if firstDiff != 0 {
			return firstDiff
		}
	}
	return 0
}
