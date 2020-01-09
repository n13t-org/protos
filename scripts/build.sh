#!/usr/bin/env bash
PROTOS=$(find . -iname '*.proto')
OUTPUT_DIR="gen"

rm -rf $OUTPUT_DIR
# gen by dir
dir=$(find . -name 'v[0-9]*' -not -path "*.git/*")

for d in $dir; do
  d=$(echo $d | cut -d '/' -f2-)
  printf "Generating code in $d, output: $OUTPUT_DIR/$d\n"
  eval $(printf "./scripts/protoc-all -d $d -l go --lint --with-validator --with-gateway --with-docs -o $OUTPUT_DIR/$d\n")
  eval $(printf "./scripts/protoc-all -d $d -l descriptor_set -o $OUTPUT_DIR/$d\n")
done

# gen file by file
#for proto in $PROTOS; do
#  proto=$(echo $proto | cut -d '/' -f2-)
#  dirname=$(dirname $proto)
#  filename=$(basename $proto)
#  outdir="generated/$dirname"
#  printf "./scripts/protoc-all -i %s -f %s -l go --with-docs markdown -o %s\n" $dirname $filename $OUTPUT_DIR/$dirname
#  eval $(printf "./scripts/protoc-all -i %s -f %s -l go -o %s --with-docs html,%s.html\n" $dirname $filename $OUTPUT_DIR/$dirname $filename)
#done





## Keep this line at the bottom
[ -d $OUTPUT_DIR ] && sudo chown -R $USER $OUTPUT_DIR
