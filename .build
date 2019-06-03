#!/usr/bin/env sh

preprocess_html() {
	dir=$(dirname $1)
	mustache $dir/mustache.json $dir/index.mustache > $dir/index.html
	echo "Generated $dir/index.html"
}

export -f preprocess_html
find . -name "mustache.json" -exec bash -c 'preprocess_html "$0"' {} \;

sass scss:css
echo "Generated CSS files"
