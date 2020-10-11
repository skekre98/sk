#!/bin/bash

while getopts "v:" opt; do
  case $opt in
    v) version=$OPTARG   ;;
    *) echo 'error' >&2
       exit 1
  esac
done

if [ -z "$version" ]
then
  echo "release version is empty...";
fi


function tag_and_release {
    git tag -a v$version -m "Releasing new version" master
    git push --tags
}

tag_and_release