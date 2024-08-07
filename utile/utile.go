package utile

import (
	"errors"
	fmt "fmt"
	os "os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
)

func Unfiniched(me string) string {
	return me + ".unfiniched"
}

func Cfgjson(me string) string {
	return me + ".cfg.json"
}

func PathIsExist(p_path string) (bool, error) {

	log.Infoln(fmt.Sprintf("\n\tTesting If the File %s exist ", p_path))
	_, err := os.Stat(p_path)
	if err == nil {
		log.Infoln("\n\tFile does not exist")
		return true, nil

	} else if os.IsNotExist(err) {
		log.Infoln("\n\tFile does not exist")   
		return false, nil  
	}
 
	return false, err  
}
 

func SplitSlice(slice []int, numChunks int) ([][]int , error ) {


	if len(slice) == 0 {
		err := fmt.Sprintln("The slice is empty") 
		log.Error(err) 
		return [][]int{{}}, errors.New(err)
	} 

	var chunks [][]int
	chunkSize := (len(slice) + numChunks - 1) / numChunks // Ceiling of len(slice) / numChunks

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks , nil
}


func RandomIntByRange(TheMin int,TheMax int) int {
	return rand.Intn(TheMax-TheMin+1) + TheMin
 
}