package main

import (
	"fmt"
	"sort"
)

type Delivery struct {
	teamSize int
	pizzaIds []int
}

//const fileName = "a_example"
//const fileName = "b_little_bit_of_everything"
//const fileName = "c_many_ingredients"
//const fileName = "d_many_pizzas"
const fileName = "e_many_teams"

func main() {
	teams, allPizzas := ReadInput(fileName + ".in")
	fmt.Println(fileName, teams, len(allPizzas))

	sort.Slice(allPizzas, func(i, j int) bool {
		return allPizzas[i].score < allPizzas[j].score
	})

	deliveries := make([]Delivery, 0)

	for {
		delivPizzas := make([]Pizza, 1, 4)
		delivPizzas[0] = allPizzas[0]
		allPizzas = allPizzas[1:]

		for {
			nextPizzaIndex := findMatch(delivPizzas, allPizzas)

			delivPizzas = append(delivPizzas, allPizzas[nextPizzaIndex])
			allPizzas = rem(allPizzas, nextPizzaIndex)

			if len(delivPizzas) == 5 {
				fmt.Println(teams, len(allPizzas))
				panic("Shouldn't happen")
			}

			if teams[len(delivPizzas)-2] > 0 {
				break
			}
		}

		deliveries = append(deliveries, Delivery{
			teamSize: len(delivPizzas),
			pizzaIds: extractIds(delivPizzas),
		})

		teamIndex := len(delivPizzas) - 2
		teams[teamIndex] = teams[teamIndex] - 1

		if noMoreDeliveries(len(allPizzas), teams) {
			break
		}
	}

	fmt.Println(deliveries)
	write(fileName+".out", deliveries)
}

// return index in available array
func findMatch(pizza []Pizza, available []Pizza) int {
	return len(available) - 1
}

// surprisingly - this produces worse results
func findMatchV2(alreadyMatched []Pizza, available []Pizza) int {
	alreadyMatchedIngredientSet := make(map[Ingredient]bool)

	for _, p := range alreadyMatched {
		for _, i := range p.ingredients {
			alreadyMatchedIngredientSet[i] = true
		}
	}

	for i := len(available) - 1; i > 0; i-- {
		nextCandidate := available[i]

		for _, possibleIngredient := range nextCandidate.ingredients {
			if _, present := alreadyMatchedIngredientSet[possibleIngredient]; present {
				return i
			}
		}
	}

	return len(available) - 1
}

func noMoreDeliveries(leftPizzas int, teams [3]uint) bool {
	// fmt.Println("Debug: ", leftPizzas, teams)

	if leftPizzas < 2 {
		return true
	} else if leftPizzas < 3 && teams[0] == 0 {
		return true
	} else if leftPizzas < 4 && teams[1] == 0 && teams[0] == 0 {
		return true
	} else if leftPizzas < 5 && teams[2] == 0 && teams[1] == 0 && teams[0] == 0 {
		return true
	} else if teams[2] == 0 && teams[1] == 0 && teams[0] == 0 {
		return true
	}

	return false
}

func rem(slice []Pizza, i int) []Pizza {
	return append(slice[:i], slice[i+1:]...)
}

func extractIds(pizzas []Pizza) []int {
	ids := make([]int, len(pizzas), len(pizzas))

	for i := 0; i < len(pizzas); i++ {
		ids[i] = int(pizzas[i].id)
	}

	return ids
}
