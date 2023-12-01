.PHONY: fzf new today
.DEFAULT: fzf

BRANCH := $(shell date "+%Y%b%d")

fzf:
	git sw $(shell git branch --color=never --list -v --abbrev=8 | cut -c3- | fzf --preview 'git show {+1}' | cut -d" " -f1)

new:
	git sw -c $(BRANCH)

today:
	git sw $(BRANCH)

