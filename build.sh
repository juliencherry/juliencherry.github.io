#!/usr/bin/env sh

export OUT_DIR=.
args_needed=0
skip_blog_posts=false

while getopts ":o:sh" opt; do
	case $opt in
		h)
			echo "./build [-hs] [-o directory]"
			echo ""
			echo "	-h: print help message"
			echo "	-o: use specified output directory"
			echo "	-s: skip downloading blog posts"
			exit 0
			;;
		o)
			args_needed=$(( args_needed + 2 ))
			OUT_DIR=$OPTARG
			;;
		s)
			args_needed=$(( args_needed + 1 ))
			skip_blog_posts=true
			;;
		\?)
			echo "Invalid flag: -$OPTARG" >&2
			exit 1
			;;
		:)
			echo "The -$OPTARG flag requires an argument" >&2
			exit 1
			;;
	esac
done

if [ $# -ne $args_needed ]; then
	echo "The number of arguments must match the number of flags" >&2
	exit 1
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

if [ "$skip_blog_posts" = true ]; then
	exit 0
fi

cd $OUT_DIR

wget https://blog.juliencherry.now.sh -r
cd blog.juliencherry.now.sh

cd post
find . -name "*" -mindepth 1 -exec bash -c 'mv $0 $0.html && mkdir $0 && mv $0.html $0/index.html' {} \;
cd ..

mv css/post.css ../css
mv img ..
mv post ..

cd ..
rm -r blog.juliencherry.now.sh

echo "Copied over blog posts"
