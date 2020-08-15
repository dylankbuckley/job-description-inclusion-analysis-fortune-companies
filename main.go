package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"./helpers"
)

func main()  {
	preapreData()
}

func preapreData() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/schema.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened schema.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var schemaData []helpers.SchemaRaw

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &schemaData)

	for i := 0; i < len(schemaData); i++ {
		// Tech
		jdTech, err := ioutil.ReadFile("data/jobDescriptions/" + schemaData[i].JobReferenceTech)
		if err != nil {
			fmt.Println(err)
		}
		schemaData[i].DescriptionTech = strings.ToLower(string(jdTech))
		schemaData[i].LengthTech = len(schemaData[i].DescriptionTech)
		schemaData[i].MasculineCountTech = masculineInstanceCount(schemaData[i].DescriptionTech)
		schemaData[i].FemanineCountTech = femanineInstanceCount(schemaData[i].DescriptionTech)
		schemaData[i].RawGenderTech = rawGenderScore(schemaData[i].MasculineCountTech, schemaData[i].FemanineCountTech)

		// Legal
		jdLegal, err := ioutil.ReadFile("data/jobDescriptions/" + schemaData[i].JobReferenceLegal)
		if err != nil {
			fmt.Println(err)
		}
		schemaData[i].DescriptionLegal = strings.ToLower(string(jdLegal))
		schemaData[i].LengthLegal = len(schemaData[i].DescriptionLegal)
		schemaData[i].MasculineCountLegal = masculineInstanceCount(schemaData[i].DescriptionLegal)
		schemaData[i].FemanineCountLegal = femanineInstanceCount(schemaData[i].DescriptionLegal)
		schemaData[i].RawGenderLegal = rawGenderScore(schemaData[i].MasculineCountLegal, schemaData[i].FemanineCountLegal)

		// Accounting
		jdAccounting, err := ioutil.ReadFile("data/jobDescriptions/" + schemaData[i].JobReferenceAccounting)
		if err != nil {
			fmt.Println(err)
		}
		schemaData[i].DescriptionAccounting = strings.ToLower(string(jdAccounting))
		schemaData[i].LengthAccounting = len(schemaData[i].DescriptionAccounting)
		schemaData[i].MasculineCountAccounting = masculineInstanceCount(schemaData[i].DescriptionAccounting)
		schemaData[i].FemanineCountAccounting = femanineInstanceCount(schemaData[i].DescriptionAccounting)
		schemaData[i].RawGenderAccounting = rawGenderScore(schemaData[i].MasculineCountAccounting, schemaData[i].FemanineCountAccounting)

		// Admin
		jdAdmin, err := ioutil.ReadFile("data/jobDescriptions/" + schemaData[i].JobReferenceAdmin)
		if err != nil {
			fmt.Println(err)
		}
		schemaData[i].DescriptionAdmin = strings.ToLower(string(jdAdmin))
		schemaData[i].LengthAdmin = len(schemaData[i].DescriptionAdmin)
		schemaData[i].MasculineCountAdmin = masculineInstanceCount(schemaData[i].DescriptionAdmin)
		schemaData[i].FemanineCountAdmin = femanineInstanceCount(schemaData[i].DescriptionAdmin)
		schemaData[i].RawGenderAdmin = rawGenderScore(schemaData[i].MasculineCountAdmin, schemaData[i].FemanineCountAdmin)

		// Graduate
		jdGraduate, err := ioutil.ReadFile("data/jobDescriptions/" + schemaData[i].JobReferenceGraduate)
		if err != nil {
			fmt.Println(err)
		}
		schemaData[i].DescriptionGraduate = strings.ToLower(string(jdGraduate))
		schemaData[i].LengthGrad = len(schemaData[i].DescriptionGraduate)
		schemaData[i].MasculineCountGrad = masculineInstanceCount(schemaData[i].DescriptionGraduate)
		schemaData[i].FemanineCountGrad = femanineInstanceCount(schemaData[i].DescriptionGraduate)
		schemaData[i].RawGenderGrad = rawGenderScore(schemaData[i].MasculineCountGrad, schemaData[i].FemanineCountGrad)
	}

	for i := 0; i < len(schemaData); i++ {
		fmt.Println(schemaData[i].Company, ",", schemaData[i].RawGenderTech, ",", schemaData[i].RawGenderLegal, ",", schemaData[i].RawGenderAccounting, ",", schemaData[i].RawGenderAdmin, ",", schemaData[i].RawGenderGrad)
	}
	

}

func masculineInstanceCount(description string) int{
	counter := 0
	for i := 0; i < len(helpers.MasculineCodedWords); i++ {
		wordToCheck := helpers.MasculineCodedWords[i]
		res := strings.Contains(description, wordToCheck)
		if res == true {
			count := strings.Count(description, wordToCheck)
			counter = counter + count
		}
	}
	return counter
}

func femanineInstanceCount(description string) int{
	counter := 0
	for i := 0; i < len(helpers.FemanineCodedWords); i++ {
		wordToCheck := helpers.FemanineCodedWords[i]
		res := strings.Contains(description, wordToCheck)
		if res == true {
			count := strings.Count(description, wordToCheck)
			counter = counter + count
		}
	}
	return counter
}

func rawGenderScore(masculineCount int, femanineCount int) int {
	rawScore := femanineCount - masculineCount
	return rawScore
}











	// Add Job Description Text - DONE
	// Remove Formatting (Lowercase) - DONE
	// Word Count - DONE

	// Deal to Hyphonated Words
	
	// Femanine Coded Count (Unique)
	// Masculine Coded Count (Unique)
	// Sector Raw Count
	// Sector Score
	// Overall Score
	// Export to CSV