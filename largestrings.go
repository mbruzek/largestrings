package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var largeNumbers = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4 ", "  44 ", " 4 4 ", "4  4 ", "44444", "   4 ", "   4 "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

var uppercase = [][]string{
	{"   A   ", "  A A  ", " A   A ", " AAAAA ", " A   A ", "A     A", "A     A"},
	{"BBBBB  ", "B    B ", "B    B ", "BBBBB  ", "B    B ", "B    B ", "BBBBB  "},
	{"  CCC  ", " C   C ", "C      ", "C      ", "C      ", " C   C ", "  CCC  "},
	{"DDDD   ", "D   D  ", "D    D ", "D    D ", "D    D ", "D   D  ", "DDDD   "},
	{"EEEEEEE", "E      ", "E      ", "EEEEEEE", "E      ", "E      ", "EEEEEEE"},
	{"FFFFFFF", "F      ", "F      ", "FFFF   ", "F      ", "F      ", "F      "},
	{"  GGG  ", " G   G ", "G      ", "G    GG", "G     G", " G    G", "  GGG G"},
	{"H     H", "H     H", "H     H", "HHHHHHH", "H     H", "H     H", "H     H"},
	{"  III  ", "   I   ", "   I   ", "   I   ", "   I   ", "   I   ", "  III  "},
	{"  JJJ  ", "   J   ", "   J   ", "   J   ", "   J   ", "J  J   ", " JJJ   "},
	{"K    K ", "K   K  ", "K  K   ", "KK     ", "K  K   ", "K   K  ", "K    K "},
	{"L      ", "L      ", "L      ", "L      ", "L      ", "L      ", "LLLLLLL"},
	{"M     M", "MM   MM", "M M M M", "M  M  M", "M     M", "M     M", "M     M"},
	{"N     N", "NN    N", "N N   N", "N  N  N", "N   N N", "N    NN", "N     N"},
	{"  OOO  ", " O   O ", "O     O", "O     O", "O     O", " O   O ", "  OOO  "},
	{"PPPPP  ", "P    P ", "P    P ", "PPPPP  ", "P      ", "P      ", "P      "},
	{"  QQQ  ", " Q   Q ", "Q     Q", "Q     Q", "Q   Q Q", " Q   Q ", "  QQQ Q"},
	{"RRRRR  ", "R    R ", "R    R ", "RRRRR  ", "R   R  ", "R    R ", "R     R"},
	{"  SSS  ", " S   S ", "  S    ", "   S   ", "    S  ", " S   S ", "  SSS  "},
	{"TTTTTTT", "   T   ", "   T   ", "   T   ", "   T   ", "   T   ", "   T   "},
	{"U     U", "U     U", "U     U", "U     U", "U     U", " U   U ", "  UUU  "},
	{"V     V", "V     V", " V   V ", " V   V ", "  V V  ", "  V V  ", "   V   "},
	{"W     W", "W     W", "W     W", "W  W  W", "W W W W", "Ww   WW", "W     W"},
	{"X     X", " X   X ", "  X X  ", "   X   ", "  X X  ", " X   X ", "X     X"},
	{"Y     Y", " Y   Y ", "  Y Y  ", "   Y   ", "   Y   ", "   Y   ", "   Y   "},
	{"ZZZZZZZ", "     Z ", "    Z  ", "   Z   ", "  Z    ", " Z     ", "ZZZZZZZ"},
}

var lowercase = [][]string{
	{"     ", "     ", "  aa ", "    a", " aaaa", "a   a", " aaaa"},
	{"b    ", "b    ", "bbb  ", "b  b ", "b  b ", "b  b ", "bbb  "},
	{"     ", "     ", " ccc ", "c    ", "c    ", "c    ", " ccc "},
	{"    d", "    d", "  ddd", "d   d", "d   d", "d   d", "  ddd"},
	{"     ", "     ", " eee ", "e   e", "eeee ", "e    ", " eee "},
	{"  ff ", "  f  ", "  f  ", " fff ", "  f  ", "  f  ", "  f  "},
	{"     ", "     ", " ggg ", "g   g", " gggg", "    g", " gggg"},
	{"h    ", "h    ", "hhhh ", "h   h", "h   h", "h   h", "h   h"},
	{"     ", "  i  ", "     ", "  i  ", "  i  ", "  i  ", "  i  "},
	{"     ", "   j ", "     ", "   j ", "   j ", "   j ", "jjj  "},
	{"k    ", "k    ", "k  k ", "k k  ", "kk   ", "k k  ", "k  k "},
	{"  l  ", "  l  ", "  l  ", "  l  ", "  l  ", "  l  ", "  l  "},
	{"       ", "       ", "mmm mmm", "m  m  m", "m  m  m", "m  m  m", "m  m  m"},
	{"     ", "     ", "nnn  ", "n  n ", "n  n ", "n  n ", "n  n "},
	{"     ", "     ", " ooo ", "o   o", "o   o", "o   o", " ooo "},
	{"     ", "     ", "pppp ", "p   p", "pppp ", "p    ", "p    "},
	{"     ", "     ", " qqqq", "q   q", " qqqq", "    q", "    q"},
	{"     ", "     ", " rrr ", "r    ", "r    ", "r    ", "r    "},
	{"     ", "     ", " sss ", "s    ", " sss ", "    s", " sss "},
	{"     ", "  t  ", "  t  ", " ttt ", "  t  ", "  t  ", "  tt "},
	{"     ", "     ", "u   u", "u   u", "u   u", "u   u", " uuu "},
	{"     ", "     ", "v   v", "v   v", " v v ", " v v ", "  v  "},
	{"       ", "       ", "w     w", "w  w  w", "w  w  w", "w  w  w", " ww ww "},
	{"     ", "     ", "x   x", " x x ", "  x  ", " x x ", "x   x"},
	{"     ", "     ", "y   y", " y y ", "  y  ", "  y  ", " yy  "},
	{"     ", "     ", " zzzz", "   z ", "  z  ", " z   ", "zzzz "},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s alphanumeric\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	text := os.Args[1]
	for row := range uppercase[0] {
		line := ""
		for column := range text {
			digit := text[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += largeNumbers[digit][row] + "  "
			} else {
				letter := text[column] - 'A'
				if 0 <= letter && letter <= 26 {
					line += uppercase[letter][row] + " "
				} else {
					letter = text[column] - 'a'
					if 0 <= letter && letter <= 26 {
						line += lowercase[letter][row] + " "
					} else {
						// Treat undefined runes as spaces.
						line += "  "
					}
				}
			}
		}
		fmt.Println(line)
	}
}
