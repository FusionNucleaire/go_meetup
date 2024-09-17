package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "bufio"
    "encoding/csv"
)

func main() {


    // Load a TXT file.
    f, _ := os.Open("./bibliotheque.csv")

    // Create a new reader.
    r := csv.NewReader(bufio.NewReader(f))
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }

        fmt.Println(record)
        fmt.Println(len(record))
        for value := range record {
            fmt.Printf("  %v\n", record[value])
        }
    }


    content := "Gopher it!"
    file, err:= os.Create("./test.txt")
    checkErrors(err)

    length, err := io.WriteString(file, content)
    checkErrors(err)

    fmt.Printf("\nEcriture dans le fichier de %v caractères\n", length)

    defer file.Close()
    defer readFile("test.txt")

}

func readFile(fileName string) {
    data, err := ioutil.ReadFile(fileName)
    checkErrors(err)
    fmt.Println("Texte récupéré du fichier : ", string(data))
}
func checkErrors(err error) {
    if err != nil {
        panic(err)
    }
}
