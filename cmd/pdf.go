/*
Copyright © 2023 Prince Kumar <princekrvert@outlook.com>

*/
package cmd

import (
        "fmt"
        "os"
        "github.com/pdfcpu/pdfcpu/pkg/api"
        "bufio"
        "log"
        "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
        "github.com/spf13/cobra"
)
// Create a function to read the file and store the line in the slice
func lineSlice(filename string)[]string{
        lineslice := [] string{}
        // NOw read the file line by line
        file, err := os.Open(filename)
        if err != nil{
        log.Fatal("Some error encountred")
}
        defer file.Close()
        // NOw scan the file
        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanLines)
        for scanner.Scan(){
                lines:= scanner.Text()
                // Append line to the slice
                lineslice = append(lineslice,lines)
}
// NOw handle the error during the scannig
        return lineslice
}

// Create a function to check if file is exists or not
func isExists( filename string ) bool {
        if _,err := os.Stat(filename); err == nil{
                return true;
        }else{
                return false;
        }
}
// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
        Use:   "pdf",
        Short: "Path of pdf fileI where pdf file is located.",
        Long: `Enter the full path of pdf file with extension where pdf file is located`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("pdf called")
        // first cmd will be the path of pdf file and second arg will be the path of wordlist , make if statemnt to identigy
        if len(args) == 2 {
                // All process goes here ...
                if isExists(args[0]) {
                        if isExists(args[1]){// All precess goes here
                                // now read the file line by line ...
                                for _,pass:= range lineSlice(args[1]){
                                //decrypt the pdf here ..
                                conf := pdfcpu.NewAESConfiguration(pass, "opw", 256, pdfcpu.PermissionsNone)
                                err := api.DecryptFile(args[0], args[0]+"unlock",conf)
    if err != nil {
        panic(err)
    }
                        }

                        }else{
                        fmt.Println("\033[31;1m Wordlist file not found")
                }
                }else{
                        fmt.Println("\033[31;1m Pdf file not found")
                }

        }else{
                panic("Check the passed argument")
        }
        },
}

func init() {
        rootCmd.AddCommand(pdfCmd)

        // Here you will define your flags and configuration settings.

        // Cobra supports Persistent Flags which will work for this command
        // and all subcommands, e.g.:
        // pdfCmd.PersistentFlags().String("foo", "", "A help for foo")

        // Cobra supports local flags which will only run when this command
        // is called directly, e.g.:
        // pdfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
