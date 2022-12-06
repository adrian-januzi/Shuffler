#!/bin/bash
#for c in {0..255}; do tput setaf $c; tput setaf $c | cat -v; echo =$c; done
clear
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

# Tree
for ((i=1; i<20; i+=2))
{
    tput cup $lin $col
    for ((j=1; j<=i; j++))
    {
        echo -n \*
    }
    let lin++
    let col--
}

tput sgr0; tput setaf 3

# Trunk
for ((i=1; i<=2; i++))
{
    tput cup $((lin++)) $c
    echo '|*|'
}
new_year=$(date +'%Y')
let new_year++
tput setaf 1; tput bold
tput cup $lin $((c - 6));
echo
tput setaf 46; tput bold
figlet -w $(tput cols) -c  MERRY CHRISTMAS
#tput cup $((lin + 1)) $((c - 10)); echo And lots of CODE in $new_year
tput setaf 195; tput bold
echo
echo
shuf team > $$.tmp
while read p; do
 figlet -w $(tput cols) -c $p
done <$$.tmp
rm $$.tmp


let c++
k=1

# Lights and decorations
while true; do
    for ((i=1; i<=35; i++)) {
        # Turn off the lights
        [ $k -gt 1 ] && {
            tput setaf 196; tput bold
            tput cup ${line[$[k-1]$i]} ${column[$[k-1]$i]}; echo \*
            unset line[$[k-1]$i]; unset column[$[k-1]$i]  # Array cleanup
        }

        li=$((RANDOM % 9 + 3))
        start=$((c-li+2))
        co=$((RANDOM % (li-2) * 2 + 1 + start))
        tput setaf $color; tput bold   # Switch colors
        tput cup $li $co
        echo o
        line[$k$i]=$li
        column[$k$i]=$co
        color=$(((color+1)%8))
    }
    k=$((k % 2 + 1))

    afplay Jingle-Bells.mp3
done