package add

import (
	"github.com/Abdulaziz-Ali123/condals/condals"
	"github.com/spf13/cobra"
)


var AddCmd = &cobra.Command{
	Use: "add",
	Short: "Add conda enviorments to the list",
	Long: `Add conda eviormetns to the list For example:
condals add --name "Name of env" --path "full/path/to/env"
condals add -n "Name of env" -p "full/path/to/env"
`,
	RunE: func (cmd *cobra.Command, args []string) error{
		//get falgs
		Name, _ := cmd.Flags().GetString("name")
		Path, _ := cmd.Flags().GetString("path")
		
		
		//retrevie conda env list and storage from context
		condaEnvs := cmd.Context().Value("condaEnvs").(condals.CondaEnvs)
		storage   := cmd.Context().Value("storage").(*condals.Storage[condals.CondaEnvs])

		if Name != "None" && Path != "None" {
			if err := condaEnvs.Add(Name, Path); err != nil {
				return err
			}
 			storage.Save(&condaEnvs)
		}else {
			cmd.Help()
		}
		
		return nil
	},
}


func init() {
	AddCmd.Flags().StringP("name", "n", "None", "Name of the conda enviorment that is being added.")
	AddCmd.Flags().StringP("path", "p", "None", "Path to the conda enviorment being added.")
}
