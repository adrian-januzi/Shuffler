#!/bin/bash
#for c in {0..255}; do tput setaf $c; tput setaf $c | cat -v; echo =$c; done
clear

tune=$(afplay drumroll.wav) &
tune_pid=$(pgrep afplay drumroll.wav)

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
echo -e "${RED}"

for i in 5 4 3 2 1
do
    figlet -w $(tput cols) -c $i
    sleep 1
done
clear
echo -e "${GREEN}"
echo -e "${NC}"

trap "tput reset; tput cnorm; exit" 2
clear
tput civis
lin=2
col=$(($(tput cols) / 2))
c=$((col-1))
est=$((c-2))
color=0
tput setaf 46; tput bold

tput sgr0; tput setaf 3

echo
echo

shuf team > $$.tmp
while read p; 
do    
    figlet -w $(tput cols) -c $p
done <$$.tmp
rm $$.tmp


let c++
k=1

kill -9 $tune_pid
