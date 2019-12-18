#!/bin/sh
PROTOS=$(find . -iname '*.proto')
rm -rf static/
for proto in $PROTOS; do
  proto=$(echo $proto | cut -d '/' -f2-)
  dirname=$(dirname $proto)
  filename=$(basename $proto)
  printf "./scripts/protoc-all -i %s -f %s -l go -o pkg/apis/$dirname\n" $dirname $filename $outdir
  eval $(printf "./scripts/protoc-all -i %s -f %s -l go -o pkg/apis/$dirname\n" $dirname $filename $outdir)
done
sudo chown -R $USER pkg