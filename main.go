package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"
)

/*
	Examples:
	Thu 01 Jan 1970 00:00:00 UTC
	Sat 01 Jan 2000 00:00:00 UTC
	Thu 01 Jan 2015 00:00:00 UTC
	Sun 01 Jan 2017 00:00:00 UTC
	Sun 01 Jan 2023 00:00:00 UTC
*/

var startDateString = "Wed 01 Jan 2025 00:00:00 UTC"   // Write your start date in this format, please
var endDateString = "Sun 08 Jun 2025 00:00:00 UTC"     // Write your end date in this format, please
var eachDayCommit int64 = 50                           // Write your amount of commits
var gitInit bool = true                                // Do you need to init git repository?
var amountOfTries int64 = 1024                         // Set the amount of tries
var dateLayout string = "Mon 02 Jan 2006 15:04:05 MST" // Don't touch, please

func main() {
	userName, err := user.Current()
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	fmt.Printf("Where is your contribution, Mr. %s?\n", userName.Username)

	cmd := exec.Command("git", "-v")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("You didn't install Git, Mr. %s...\n", userName.Username)
		fmt.Println("This fundamental oversight undermines your technical credibility.")
		fmt.Println("Unfortunately, we must disqualify your submission from further consideration.")
		os.Exit(1)
	} else {
		fmt.Printf("Your Git version is `%s`...\n", string(output))
		fmt.Println("Look at that — Git is installed.")
		fmt.Println("And here I was bracing for a support ticket.")
	}

	if gitInit == true {
		gitDir, err := os.Stat("./.git")
		if os.IsExist(err) {
			err = os.RemoveAll(gitDir.Name())
			if err != nil {
				fmt.Printf("It appears you tried to be clever with me, Mr. %s.\n", userName.Username)
				fmt.Println("Regrettably, this does not meet our expectations.")
				fmt.Println("We must kindly decline your candidacy.")
				os.Exit(1)
			} else {
				fmt.Printf("The directory `%s` was found to be unsuitable for our environment and has been removed to ensure a clean repository state.\n", gitDir.Name())
			}
		} else {
			fmt.Printf("A Git repository is not present in the expected location, Mr. %s.", userName.Username)
			fmt.Println("This omission prevents further execution.")
			fmt.Println("We shall overlook it this time — but rest assured, there will not be another")
		}

		cmd = exec.Command("git", "init")
		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("You cannot even initialize your Git repository, Mr. %s.\n", userName.Username)
			fmt.Println("This foundational task was unsuccessful.")
			fmt.Println("We must conclude that you are not prepared for this engagement.")
			os.Exit(1)
		} else {
			fmt.Printf("Ok, I see you initialized your Git repository, Mr. %s...\n", userName.Username)
			fmt.Println("Bravo. A Git repository. The absolute baseline — but let’s pretend we’re impressed.")
		}
	}

	startDateTime, err := time.Parse(dateLayout, startDateString)
	if err != nil {
		fmt.Printf("You cannot even manage your time, Mr. `%s`.\n", userName.Username)
		fmt.Println("Effective scheduling is non-negotiable.")
		fmt.Println("Your inability to handle temporal precision disqualifies this application.")
		os.Exit(1)
	} else {
		fmt.Printf("Did you start to contribute only in %s, Mr. %s?\n", startDateTime, userName.Username)
		fmt.Println("An unusually late start, one might say.")
		fmt.Println("While not disqualifying in itself, such timing rarely goes unnoticed in final evaluations.")
	}

	endDateTime, err := time.Parse(dateLayout, endDateString)
	if err != nil {
		fmt.Printf("You cannot even manage your time, Mr. `%s`.\n", userName.Username)
		fmt.Println("Effective scheduling is non-negotiable.")
		fmt.Println("Your inability to handle temporal precision disqualifies this application.")
		os.Exit(1)
	} else {
		fmt.Printf("Why do you want to finish in %s, Mr. %s?\n", endDateTime, userName.Username)
		fmt.Println("It's a truncated timeline for such an endeavor.")
		fmt.Println("Nonetheless, we shall continue — though not without noting this decision in your final review.")
	}

	currentCommit := int64(0)
	allCommits := int64((endDateTime.Sub(startDateTime) + 24*time.Hour).Hours()) / 24 * eachDayCommit
	localPercent := float64(0)
	globalPercent := float64(0)
	startDateFormat := ""

	for startDateTime.Before(endDateTime) || startDateTime.Equal(endDateTime) {
		/*
			//Debug Feature
			fmt.Println(startDateTime.Format(dateLayout))
		*/
		startDateFormat = startDateTime.Format(dateLayout)
		err = os.Setenv("GIT_AUTHOR_DATE", startDateFormat)
		if err != nil {
			fmt.Printf("You cannot even manage your system, Mr. %s.\n", userName.Username)
			fmt.Println("A stable environment is the bare minimum. Unfortunately, this reflects a lack of operational readiness.")
			os.Exit(1)
		}
		err = os.Setenv("GIT_COMMITTER_DATE", startDateFormat)
		if err != nil {
			fmt.Printf("You cannot even manage your system, Mr. %s.\n", userName.Username)
			fmt.Println("A stable environment is the bare minimum. Unfortunately, this reflects a lack of operational readiness.")
			os.Exit(1)
		}

		for i := int64(0); i < eachDayCommit; i++ {
			currentCommit++
			err = os.WriteFile("AUTOFILE", []byte(strconv.FormatInt(time.Now().UnixNano(), 19)), 0644)
			if err != nil {
				fmt.Printf("You cannot even contribute to `AUTOFILE`, Mr. %s.\n", userName.Username)
				fmt.Println("Your failure to perform basic write operations raises serious concerns. We must bring this review to a close.")
				os.Exit(1)
			}

			cmd = exec.Command("git", "add", ".")
			output, err = cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("You cannot even add all your files to the Git repository, Mr. %s.\n", userName.Username)
				fmt.Println("Version control hygiene is imperative. This lapse is not something we can overlook.")
				os.Exit(1)
			}

			var cmd *exec.Cmd
			for i := int64(0); i < amountOfTries; i++ {
				cmd = exec.Command("git", "commit", "-m", "AUTO UPDATE", "--date", startDateFormat)
				if output, err = cmd.CombinedOutput(); err == nil {
					break
				}
			}

			if err != nil {
				fmt.Printf("I don't see your commits, Mr. %s.\n", userName.Username)
				fmt.Println("A contribution history was expected.")
				fmt.Println("In its absence, we are left with no choice but to decline your candidacy.")
				os.Exit(1)
			}

			localPercent = float64(i+1) / float64(eachDayCommit) * 100
			globalPercent = float64(currentCommit) / float64(allCommits) * 100
			fmt.Printf("{%d}/{%d}\t", startDateTime.Year(), endDateTime.Year())

			fmt.Printf("{%d/%d}\t", i+1, eachDayCommit)
			fmt.Printf("{%d/%d}\t", currentCommit, allCommits)

			fmt.Printf("{%.2f%%}\t", localPercent)
			fmt.Printf("{%.2f%%}\n", globalPercent)
		}

		startDateTime = startDateTime.Add(24 * time.Hour)
	}
	err = os.Remove("AUTOFILE")
	if err != nil {
		fmt.Println("The file `AUTOFILE` could not be deleted.")
		fmt.Println("Leaving behind digital clutter is not an acceptable practice at this organization.")
		os.Exit(1)
	} else {
		fmt.Println("Oh, how noble. You even cleaned up after yourself and removed `AUTOFILE`.")
	}

	fmt.Printf("We will take your portfolio into thoughtful consideration, Mr. %s...\n", userName.Username)
	fmt.Println("This decision has been reached due to the absence of verifiable work records and a failure to complete identity verification.")
	fmt.Println("We regret to inform you that you do not meet the requirements.")
	fmt.Println("Farewell.")
}
