#!/bin/bash
#Author Prince Kumar
#Date 3 jun 2023
# First check for the operating system..
opreating_s=$(uname -o)
if [[ $opreating_s == "Android" ]];then
        #first check for the instllation for pdfcpu
        commnad -v pdfcpu 2>&1 > /dev/null || { echo "Installing pdfcpu"; pkg install pdfcpu -y; }
        # NOw go for the actual process.. $1 will be the path to the pdf and $2 will be the path to the wordlist
        if [[ -f $1 ]];then
                # NOw check for wordlist file
                if [[ -f $2 ]];then
                        # NOw read the wordlist file line by line and crack the password
                        while read -r line;do
                                # NOw try to crack the password here
                                pdfcpu decrypt -upw $line $1 $1+unlock.pdf 2>&1 > /dev/null
                                # NOw check the response
                                if [[ $? == 0 ]];then
                                        echo -ne "\033[32;1m Password found $line"
                                else
                                        echo -e "\033[33;1m[!] Trying $line "
                                fi
                        done < $2
                else
                        echo -e "\033[31;1m Wordlist file not found"
                fi
        else
                echo -e "\033[31;1m Pdf file not found"
                exit 1
        fi
else
        echo "Follow the standard dabian based process"
fi
