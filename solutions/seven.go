package solutions

import "main/pkg"

func Day_7() (interface{}, interface{}, error) {

	i := pkg.New_Input(7)
	err := i.Parse(pkg.ClearAllEmptyLines)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
