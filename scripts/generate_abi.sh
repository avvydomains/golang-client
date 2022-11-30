#!/bin/bash
rm abi/*.go
python3 scripts/generate_abi.py
abis=$(ls abi)
for abi in $abis
do
  dest=$(echo $abi | sed 's/.json/.go/g')
  name=$(echo $abi | sed 's/.json//g')
  echo $abi
  echo $dest
  abigen --abi=abi/$abi --pkg=$name --out=abi/$dest
done
rm abi/*.json
(
  cd abi
  files=$(ls *.go)
  for ff in $files; do
    fname=$(echo $ff | sed 's/.go//g')
    mkdir $fname
    mv $ff $fname
  done
)

