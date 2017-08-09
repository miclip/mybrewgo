// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/miclip/mybrewgo/recipe"
	"github.com/miclip/mybrewgo/utils"
	"github.com/spf13/cobra"
)

var (
	findName string
)

// recipeCmd represents the recipe command
var recipesCmd = &cobra.Command{
	Use:   "recipes",
	Short: "Manage the local recipes store.",
	Long: `Manage the local store, functions available include Add & Search

	The local store is a YAML file written to the same directory as the executable. This allows
	you to use source control repository like github.com to save and backup your recipes.`,
	Run: func(cmd *cobra.Command, args []string) {
		recipes := recipe.NewRecipes()
		recipesTmp := make(map[int]string)
		color.White("Recipes:")
		i := 0
		for k := range recipes.Recipes {
			recipesTmp[i] = k
			i++
		}
		utils.DisplayColumns(recipesTmp, 3)
		v := utils.RequestUserInput("Select a recipe:")
		s, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			color.Red("Invalid value. %v", v)
			return
		}
		r := recipes.Recipes[recipesTmp[int(s)]]
		r.Print()

	},
}

func init() {
	RootCmd.AddCommand(recipesCmd)
}
