#!/bin/bash

protos=("$@")

for proto in "${protos[@]}"
do
  if [[ -d $proto ]]; then
    continue
  fi

  if grep '^option go_package' $proto;then
    continue
  fi

  dir=${proto%/*}
  file=${proto##*/}

  newdir="protogen/$dir"
  newfile="$newdir/$file"
  # copy to protogen directory
  mkdir -p "$newdir" && cp "$proto" "$_"

  syntaxLineNum="$(grep -n "syntax" "$newfile" | head -n 1 | cut -d: -f1)"

  goPackageString="option go_package = \"github.com/tokopedia/gripmock/protogen$dir\";"
  sed -i "${syntaxLineNum}s~$~\n$goPackageString~" $newfile
  echo $newfile
done

