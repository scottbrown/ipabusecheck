package main

var (
	inputFilename  string
	outputFilename string
	silent         bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputFilename, "input", "i", "", "The list of IP addresses to check (one per line)")
	rootCmd.PersistentFlags().StringVarP(&outputFilename, "output", "o", "", "The file to write the data to")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Do not print statistics, just output to the file")
}
