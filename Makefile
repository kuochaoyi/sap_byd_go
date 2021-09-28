gitignore:
	# Apply .gitignore on an existing repository already tracking large number of files
	git rm -r --cached .
	git add .
	git commit -m ".gitignore is now working"

git_rm_cached:
	git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch main.go' --prune-empty --tag-name-filter cat -- --all

a:
	git push origin master --force