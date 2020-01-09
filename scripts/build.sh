#!/usr/bin/env bash
PROTOS=$(find . -iname '*.proto')
OUTPUT_DIR="gen"

rm -rf $OUTPUT_DIR
# gen by dir
dir=$(find . -name 'v[0-9]*' -not -path "*.git/*" | grep -v 'gen')

#eval $(printf "./scripts/protoc-all -d /opt/include/n13t/idm/v0 -l go --lint --with-validator --with-gateway --with-docs -o /defs/$OUTPUT_DIR/$d\n")
#eval $(printf "./scripts/protoc-all -d ./ -l descriptor_set -o $OUTPUT_DIR/$d\n")
#eval $(printf "./scripts/protoc-all -d /opt/include/n13t/ -l go  --go-package-map --lint --with-validator --with-gateway --with-docs -o $OUTPUT_DIR/$d\n")
#exit 0
for d in $dir; do
  d=$(echo $d | cut -d '/' -f2-)
  printf "Generating code in $d, output: $OUTPUT_DIR/n13t/$d\n"
  eval $(printf "./scripts/protoc-all -d /opt/include/n13t/$d -l go --lint --with-validator --with-gateway  -o $OUTPUT_DIR\n")
  eval $(printf "./scripts/protoc-all -d /opt/include/n13t/$d -l descriptor_set --with-docs -o $OUTPUT_DIR/n13t/$d\n")
#  mv /opt/include/n13t/$d
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
sudo chown -R $USER .
#[ -d $OUTPUT_DIR ] && sudo chown -R $USER $OUTPUT_DIR
