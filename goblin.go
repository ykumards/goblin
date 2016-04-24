//Do What The Fuck You Want To Public License (WTFPL)
// @author: ykumards

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
	"math/rand"
)

// Removes the rouge file
func cleanup(fileName string) {
	cmdGit := "git"
	rmArgs := []string{"rm", fileName}
	_, rmErr := exec.Command(cmdGit, rmArgs...).Output()
	if rmErr != nil {
		fmt.Fprintln(os.Stderr, "There was an error removing file", rmErr)
		os.Exit(1)
	}
	commitArgs := []string{"commit", "-m", "'delete'"}
	commitOut, commitErr := exec.Command(cmdGit, commitArgs...).Output()
	if commitErr != nil {
		fmt.Fprintln(os.Stderr, "There was an error committing file", commitErr)
		os.Exit(1)
	}
	fmt.Println(commitOut)
	git_push()
}

// git add
func git_add(fileName string) {
	cmdGit := "git"
	addArgs := []string{"add", fileName}
	addOut, addErr := exec.Command(cmdGit, addArgs...).Output()
	if addErr != nil {
		fmt.Fprintln(os.Stderr, "There was an error adding file", addErr)
		os.Exit(1)
	}
	fmt.Println(">Adding file")
	fmt.Println(string(addOut))
}

// git commit
func git_commit(){
	cmdGit := "git"
	commitArgs := []string{"commit", "-m", "'update'"}
	commitOut, commitErr := exec.Command(cmdGit, commitArgs...).Output()
	if commitErr != nil {
		fmt.Fprintln(os.Stderr, "There was an error commiting file", commitErr)
		os.Exit(1)
	}
	fmt.Println(">Commiting file")
	fmt.Println(string(commitOut))
}

// git push
func git_push() {
	cmdGit := "git"
	pushArgs := []string{"push"}
	_, pushErr := exec.Command(cmdGit, pushArgs...).Output()
	if pushErr != nil {
		fmt.Fprintln(os.Stderr, "There was an error pushing file", pushErr)
		os.Exit(1)
	}
	fmt.Println("> pushing delete")
}

func main() {
	// Make sure there's one and only one argument
	argList := os.Args[1:]
	if len(argList) == 0 || len(argList) > 1 {
		fmt.Println("Error: Wrong no. arguments")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(argList[0])
	startDate := time.Now()
	curDate := startDate
	fileName := "authentic.txt"
	// We're moving back in time now...!
	for i := 0; i <= n; i++ {
		  whichDay := fmt.Sprintf("################## Back %d Day ###########", i)
			fmt.Println(whichDay)
			curDate_st := curDate.String()
			fmt.Println(curDate_st)
			curDate = curDate.AddDate(0,0,-1)
			randVal := rand.Intn(1000000)
			// To make sure this looks "authentic",
			// we need a random number of commits per day
			// You can tweak this if you want to look more productive
			numCommits := rand.Intn(5)
			for j := 0; j <= numCommits; j++ {
				curTxt  := fmt.Sprintf("%s -%d", curDate_st, randVal)
				// spawning the rouge file
				echoStr := fmt.Sprintf("echo %s > %s", curTxt, fileName)
				echoCmd := exec.Command("bash", "-c", echoStr)
				echoOut, echoErr := echoCmd.Output()
				if echoErr != nil {
					fmt.Fprintln(os.Stderr, "There was an error creating file", echoErr)
				}
				fmt.Println(string(echoOut))
				// git stuff begins here
				git_add(fileName)
				// Setting git date environmental variables (this is the secret sauce)
				os.Setenv("GIT_AUTHOR_DATE", curDate_st)
				os.Setenv("GIT_COMMITTER_DATE", curDate_st)
				git_commit()
				git_push()
				// goblins need sleep too!
				time.Sleep(50 * time.Millisecond)
				// cleanup
				cleanup(fileName)
		}
	}
}
