#!/bin/bash

if [ -z "$IACCONSOLE_TOKEN" ]; then
    echo "Error: IACCONSOLE_TOKEN is not set."
    exit 1
fi

inventory_path=$1
API_URL="https://api.iacconsole.com/v1"

post_to_iacconsoledb () {
  # $1 - filename

  dimpath=$1
    dimpath=${dimpath#*$inventory_path}
    dimpath=${dimpath%%.json}
    printf "\n-------------- $dimpath -------------\n"

  curl --header "Content-Type: application/json" \
  --header "Authorization: Bearer $IACCONSOLE_TOKEN" \
  --request POST \
  -d @$1\
  "$API_URL/dimension/$dimpath?workspace=master&source=inventory&readonly=true" -q
  printf "\n------------------------\n\n"
}

for file in $(find $inventory_path -name '*.json')
do
   post_to_iacconsoledb $file
done
