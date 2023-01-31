#!/bin/bash

clear

rap "tput reset; tput cnorm; exit" 2
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

team_members=("Tom" "Glenn-C" "Glenn-S" "Patrick" "Ian" "Adrian")
randomized_members=( $(printf "%s\n" "${team_members[@]}" | shuf) )
fonts=("doh.flf" "smkeyboard.flf" "isometric2.flf" "binary.flf" "mirror.flf" "relief2.flf")
colors=(3 4 5 6 7 8 9 10 11 12 13 14 15 16)

# for member in "${randomized_members[@]}"
for i in "${!randomized_members[@]}"
do
    afplay "drumroll.wav"

    member="${randomized_members[$i]}"
    font="${fonts[$i]}"
    color="${colors[$i]}"
    tput setaf "$color"
    figlet -f "$font" -w $(tput cols) -c "$member"
    read -p ""
done