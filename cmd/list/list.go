package list

import (
	"errors"
	
	"github.com/Abdulaziz-Ali123/condals/condals"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "List all added conda enviorments",
	Long: ` List all added conda enviorments. For example:
condals list
`,
	RunE: func (cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("Too many flags passed")
		}
		
		condaEnvs := cmd.Context().Value("condaEnvs").(condals.CondaEnvs)
		storage  := cmd.Context().Value("storage").(*condals.Storage[condals.CondaEnvs])
		storage.Load(&condaEnvs)
		
		condaEnvs.List()
		
		storage.Save(&condaEnvs)

		return nil
	},
}


func init() {

}
