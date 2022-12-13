# BeersPunkApiGolang

	arrayBeersRecipes, err := Beer.GetBeers(70, 100, 4)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range *arrayBeersRecipes {
		v.Info()
	}
