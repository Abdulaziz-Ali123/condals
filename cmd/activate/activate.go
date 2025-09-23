package activate

import (
	"github.com/Abdulaziz-Ali123/condals/condals"
	"github.com/spf13/cobra"
	
	"fmt"
)

var jupyter bool 

var ActivateCmd = &cobra.Command{
	Use: "activate",
	Short: "Used to activate a selected enviorment",
	Long: `Used to activate a selected enviorment. Moves to that directory and can activate jupyter notebook
	Useage:

	#evaluate the output of the condals comand and excecute it 
	# this will change your current directory and activate jupyter notebook within the same terminal
	eval $(condals acitvate -i 0 -j)
	`,
	RunE: func (cmd *cobra.Command, args []string) error{
		Name, _ := cmd.Flags().GetString("name")
		Index, _ := cmd.Flags().GetInt("index")
		
	        // get condaEnv and storage from context
                condaEnvs := cmd.Context().Value("condaEnvs").(condals.CondaEnvs)
                storage := cmd.Context().Value("storage").(*condals.Storage[condals.CondaEnvs])
		
		storage.Load(&condaEnvs)	
		
		var Path string

		if Name != "" {
			path, err := condaEnvs.GetPathByName(Name)
			if err != nil {
				return err
			}

			Path = *path
		}else {
			path, err  :=  condaEnvs.GetPathByIndex(Index)
			if err != nil {
				return err
			}

			Path = *path
		}
		//ecmd := exec.Command("bash", "-c", "cd "+Path+" && conda init && conda activate ./env")
		command := fmt.Sprintf("cd %s && conda activate ./env", Path)
		if jupyter {
			command = command +" && jupyter notebook"
		}

		fmt.Println(command)

		return nil
	},
}


func init() {
	ActivateCmd.Flags().StringP("name", "n", "", "Flag that provides name of enviorment to be activated")
	ActivateCmd.Flags().IntP("index", "i", -1 , "Flag that provides index of enviorment to be activated")
	ActivateCmd.Flags().BoolVarP(&jupyter, "jupyter", "j", false, "Start Jupyter notebook in a selected enviorment")
}

