package cmd

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "count word letter",
	Long:  "count word letter long description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {}
