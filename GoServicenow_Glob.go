package main

/************************************************
	MIT License
	Details viewable in the Github Directory
	Copyright (c) 2018 Michael Morell
*************************************************
	Created by Michael Morell
	Released 04/04/2018
	Github: https://github.com/Mjmorell/WinbakGo
************************************************/

import (
	"unicode/utf8"
)

var (
	ProgVersion  [3]int
	newVersion   [3]int
	KBArticle    string
	HeaderString string
	HeaderLength int
	HeaderCheck  bool
	versionURL   string
	Config       Configuration
)

func init() {
	//Green := color.New(color.FgGreen).SprintFunc()
	//Red := color.New(color.FgRed).SprintFunc()
	versionURL = "https://github.com/Mjmorell/WinbakGo/raw/master/version.txt"
	HeaderCheck = false
	// Version = 0.0.1a
	ProgVersion[0] = 0
	ProgVersion[1] = 0
	ProgVersion[2] = 1

	KBArticle = "KB_notHere"
	HeaderString = " AssetFixer ── " + KBArticle + " ── Liberty University IT Helpdesk "
	HeaderLength = utf8.RuneCountInString(HeaderString)
}
