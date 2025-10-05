package main

import (
	"flag"
	"fmt"
	"os"
	"vorban/resenum/pkg/resenum"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	min_raw := flag.String("min", "0:0", "Minimum resolution, expressed as X:Y. Each dimension is optional")
	max_raw := flag.String("max", "1920:1080", "Maximum resolution, expressed as X:Y. Each dimension is optional")
	standard := flag.Bool("standard", false, "Standard resolutions only, i.e. SD, HD, FHD, QHD and UHD")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	args := flag.NArg()
	if args != 1 {
		fmt.Println("Error: exactly one argument required as ratio X:Y")
		printHelp()
		os.Exit(1)
	}

	ratio, err := resenum.ParseUintPair(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	min, err := resenum.ParseOptionalUintPair(*min_raw, resenum.MakePair(0, 0))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	max, err := resenum.ParseOptionalUintPair(*max_raw, resenum.MakePair(1920, 1080))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	resolutions := resenum.Enumerate(ratio, min, max)
	printResolutions(resolutions, *standard)
}

func printHelp() {
	fmt.Println("Usage: resenum [--min=] [--max=] [--standard] [--help] {ratio}")
	flag.PrintDefaults()
	fmt.Println("Examples:")
	fmt.Println("\tresenum 16:9")
	fmt.Println("\tresenum 16:9 --min=256:")
	fmt.Println("\tresenum 16:9 --max=:1080")
}

func printResolutions(resolutions []resenum.Resolution, standard bool) {
	var displayed uint = 0

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Resolution", "Note"})
	for _, r := range resolutions {
		if standard && r.Type == resenum.ResNonStandard {
			continue
		}
		displayed++
		t.AppendRow(table.Row{r.ToString(), r.Type.ToString()})
	}

	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total", displayed})
	t.Render()
}
