#!/usr/bin/env sh

if [ $# -eq 0 ]; then
	export OUT_DIR=.
else
	export OUT_DIR=$1
fi

preprocess_html_file() {
	html_dir=$(dirname $1)
	out_path="$(pwd)/$OUT_DIR/$html_dir"
	out_filepath="$out_path/index.html"

	mkdir -p $out_path && touch "$out_filepath"
	mustache "$html_dir/mustache.json" "$html_dir/index.mustache" > $out_filepath
	echo "Generated $out_filepath"
}

export -f preprocess_html_file
find . -name "mustache.json" -exec bash -c 'preprocess_html_file "$0"' {} \;

sass scss:$OUT_DIR/css
echo "Generated CSS files"

wget https://blog.juliencherry.now.sh -r
cd blog.juliencherry.now.sh
mkdir ../post; mv post/* $_
mkdir ../img; mv img/* $_
mv css/post.css ../css/post.css
cd ..
echo "Copied over blog posts"
