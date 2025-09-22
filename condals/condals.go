package condals

import (
	"fmt"
	"errors"
)
type CondaEnv struct {
	Name string
	Path string
}

type CondaEnvs []CondaEnv

func (condaEnvs *CondaEnvs) validateAddInput(name string) error {
	isInEnv := false

	if condaEnvs != nil{
		for _, condaEnv := range *condaEnvs {
			if condaEnv.Name == name {
				isInEnv = true
				break
			}
		}
	}
	
	if isInEnv {
		err := fmt.Sprintf("%s is already taken", name)
		return errors.New(err)
	}
	
	return nil
}


func (condaEnvs *CondaEnvs) Add(name string, path string) error {
	c := *condaEnvs

	if err := c.validateAddInput(name); err != nil {
		return err
	}
	
	env := CondaEnv{
		Name: name,
		Path: path,
	}

	*condaEnvs = append(*condaEnvs, env)

	return nil
}


func (condaEnvs *CondaEnvs) validateIndex(index int) error {
	if index < 0 || index > len(*condaEnvs){
		return errors.New("Invalid Index")
	}
	return nil
}


func (condaEnvs *CondaEnvs) DeleteByIndex(index int) error {
	c := *condaEnvs
	if err := c.validateIndex(index); err != nil {
		return err
	}
	
	if condaEnvs != nil {
		*condaEnvs = append(c[:index], c[index+1:]...)
	} else {
		return errors.New("There are no cond enviroments added")
	}
	
	return nil
}



func (condaEnvs *CondaEnvs) DeleteByName(name string) error {
	c := *condaEnvs
	result := -1
	if condaEnvs != nil{
		for index, condaEnv := range *condaEnvs {
			if condaEnv.Name == name {
				result = index
				break
			}
		}

		if result >= 0 {
			*condaEnvs = append(c[:result], c[result+1:]...)
		} else {
			err := fmt.Sprintf("%s does not exsist", name)
			return errors.New(err)
		}
	} else {
		return errors.New("There are no conda enviroments added")
	}
	
	return nil
}

func (condaEnvs *CondaEnvs) longestNameLen() int {
	longest := 0
	for _, condaEnv := range *condaEnvs {
		if len(condaEnv.Name) > longest {
			longest = len(condaEnv.Name)
		}
	}

	return longest
}

func (condaEnvs *CondaEnvs) List() {
	c := *condaEnvs
	longest := c.longestNameLen()
	if condaEnvs != nil{
		header :=  fmt.Sprintf("Index\t %-*s\t Path", longest, "Name")
		fmt.Println(header)
		for index, condaEnv := range *condaEnvs {
			row := fmt.Sprintf("%-d\t %-*s\t %-s", index, longest, condaEnv.Name, condaEnv.Path)
			fmt.Println(row)
		}
	} else {
		fmt.Println("There are no conds enviroments added")
	}
}


func (condaEnvs *CondaEnvs) GetPathByName(Name string) (*string, error) {
	for _, condaEnv := range *condaEnvs {
		if condaEnv.Name == Name{
			return &condaEnv.Path, nil
		}
	}

	return nil, errors.New("Name Not Found")
}



func (condaEnvs *CondaEnvs) GetPathByIndex(Index int) (*string, error) {
	c := *condaEnvs
	if err := c.validateIndex(Index); err != nil {
		return nil, err
	}

	return  &c[Index].Path, nil

	return nil, errors.New("Name Not Found")
}
