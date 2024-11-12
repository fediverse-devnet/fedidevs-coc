fedidevs-coc.md : fedidevs.txt templates/mozilla-community-participation-guidelines-3.1-as-template.md
	go run bin/replace-vars.go $^ $@


clean :
	-rm fedidevs-coc.md

