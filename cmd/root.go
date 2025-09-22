package cmds

import(
	"context"
	"fmt"
	"os"
	
	"github.com/Abdulaziz-Ali123/condals/cmd/activate"
        "github.com/Abdulaziz-Ali123/condals/cmd/add"
	"github.com/Abdulaziz-Ali123/condals/condals"
	"github.com/Abdulaziz-Ali123/condals/cmd/delete"
	"github.com/Abdulaziz-Ali123/condals/cmd/list"
	"github.com/spf13/cobra"
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "condals",
	Short: "A CLI that save conda enviorments",
	Long: "A CLI that save conda enviroments enabling faster activation of conda enviorments",

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	//initialize conda enviorment list
        condaEnvs := condals.CondaEnvs{}
        storage  := condals.NewStorage[condals.CondaEnvs]("/usr/tmp/condaEnvsList.json")
        storage.Load(&condaEnvs)
        
        //Create context in order to pass condaEnvs to all SubcommandPallets
        ctx := context.WithValue(context.Background(), "condaEnvs", condaEnvs)
        ctx = context.WithValue(ctx, "storage", storage)
	
	err := rootCmd.ExecuteContext(ctx)

	if err != nil {
		storage.Save(&condaEnvs)
		fmt.Println(err)
		os.Exit(1)
	}
}

func addSubcommandPalette() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(del.DeleteCmd)
	rootCmd.AddCommand(activate.ActivateCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	addSubcommandPalette()
}
