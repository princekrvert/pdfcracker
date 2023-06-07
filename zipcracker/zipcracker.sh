#!/bin/bash
#Made by prince kumar
#date 07/06/2023
#NOw start the actual process
# Trap the signals
trap user_inttrupt SIGINT
trap user_inttrupt SIGTERM
# make a function for user inttruption
user_inttrupt(){
echo -ne "\033[31;1m Exiting the process"
exit 1
}
# NOw make a banner for the tool
banner(){
echo -e "\033[32;1m Zipcracker : made by prince"
}
#NOw check if p7zip is installed or not
# check for the distro
distro=$(uname -o)
if [[ $distro == "Android" ]];then
    command -v p7zip || { echo -ne"\033[33;1m Installing zipcracker"; pkg install p7zip;}
else
        command -v p7zip || { echo -ne "\033[33;1m Installing zipcraker"; sudo apt install p7zip;}
fi
# NOw read the word list and crack the zip file
if [[ -f $1 ]];then
        if [[ -f $2 ]];then
                # nOw crack the zip
                # now read the wordlist line by line
                cat $2 | while read -r line; do
                # use the line
                p7zip -p$line $1 -y > /dev/null 2>&1
               if [[ $? == 0 ]];then
                echo -e "\033[32;1m Password found "
                else
                echo -e "\033[31;1m Trying password $line"
                fi
        done

        else
                echo -e "\033[31;1m Wordlist file not found "
        fi
else
        echo -e "\033[31;1m zip file not found "
fi
