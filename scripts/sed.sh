#!/bin/sh

if [ $# -ne 1 ]; then
  echo "引数にプロジェクト名を入力してください"
  echo "Usage: sh $0 your_pjname"
  exit 1
fi

grep -rl "go-echo-base" ./ | xargs perl -i -pe "s/go-echo-base/$1/g"
