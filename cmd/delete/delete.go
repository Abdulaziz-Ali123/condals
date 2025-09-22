package del

import (
	"github.com/Abdulaziz-Ali123/condals/condals"
	"github.com/spf13/cobra"
)

var all bool
var DeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Delete from added conda enviorments",
	Long: `Delete From added conda enviorments. For example:
condals delete --by-index  0
condals delete --by-name "envName"
condals delete --all

precidence order is all -> name -> index

`, 
	RunE: func (cmd *cobra.Command, args []string) error{
		//extract flags
		
		Name, _ := cmd.Flags().GetString("By-name")
	        Index, _ := cmd.Flags().GetInt("By-index")

		// get condaEnv and storage from context
                condaEnvs := cmd.Context().Value("condaEnvs").(condals.CondaEnvs)
                storage := cmd.Context().Value("storage").(*condals.Storage[condals.CondaEnvs])
		
		var err error = nil
		
		if all {
                    	condaEnvs = nil 
               	}else if Name != "None" { 
                       	err = condaEnvs.DeleteByName(Name)
               	}else if  Index > -1 {
			err = condaEnvs.DeleteByIndex(Index)
		}else {
			cmd.Help()
		}
		
		storage.Save(&condaEnvs)
		
		if err != nil{
			return err
		}

		return nil
	},
}

func init() {
	DeleteCmd.Flags().BoolVarP(&all, "all", "a", false, "Flage to delete all stored enviorments (no)")
	DeleteCmd.Flags().StringP("By-name", "n", "None", "Flag for a saved enviorment by name")
	DeleteCmd.Flags().IntP("By-index", "i", -1, "Flag for provided the index for deleteing an env listing by index")
}
