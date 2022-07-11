#!/bin/bash

x=1
while [ $x -le 120 ]
do
  resp=$(kubectl get pods --no-headers=true | awk '/'"$1"'/{print $1}')
  
  if [ -z "$resp" ]
  then
    echo "need to break"
    break
  else 
    echo "$resp"
  fi
  x=$(( x + 1 ))
  sleep 1;
done