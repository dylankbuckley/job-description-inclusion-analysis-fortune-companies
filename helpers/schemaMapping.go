package helpers

type Accounting struct{
	JobTile 		string `json:"Job Title__1"`
	JobReference 	string `json:"JD Ref__1"`
	JobSource 		string `json:"Source__1"`
	JobLink 		string `json:"Link__1"`
	Description		string `json:"description"`
	MasculineCount	int	   `json:"masculineCount"`
	FemanineCount	int	   `json:"femanineCount"`
	RawGender		int	   `json:"rawGender"`
}

type Legal struct{
	JobTile 		string `json:"Job Title__2"`
	JobReference 	string `json:"JD Ref__2"`
	JobSource 		string `json:"Source__2"`
	JobLink 		string `json:"Link__2"`
	Description		string `json:"description"`
	MasculineCount	int	   `json:"masculineCount"`
	FemanineCount	int	   `json:"femanineCount"`
	RawGender		int	   `json:"rawGender"`
}
type Admin struct{
	JobTile 		string `json:"Job Title__3"`
	JobReference 	string `json:"JD Ref__3"`
	JobSource 		string `json:"Source__3"`
	JobLink 		string `json:"Link__3"`
	Description		string `json:"description"`
	MasculineCount	int	   `json:"masculineCount"`
	FemanineCount	int	   `json:"femanineCount"`
	RawGender		int	   `json:"rawGender"`
}
type Graduate struct{
	JobTile 		string `json:"Job Title__4"`
	JobReference 	string `json:"JD Ref__4"`
	JobSource 		string `json:"Source__4"`
	JobLink 		string `json:"Link__4"`
	Description		string `json:"description"`
	MasculineCount	int	   `json:"masculineCount"`
	FemanineCount	int	   `json:"femanineCount"`
	RawGender		int	   `json:"rawGender"`
}
type Technology struct{
	JobTile 		string `json:"Job Title"`
	JobReference 	string `json:"JD Ref"`
	JobSource 		string `json:"Source"`
	JobLink 		string `json:"Link"`
	Description		string `json:"description"`
	MasculineCount	int	   `json:"masculineCount"`
	FemanineCount	int	   `json:"femanineCount"`
	RawGender		int	   `json:"rawGender"`
}

type SchemaRaw struct {
	Company 			string 		`json:"Company"`
	JobTileTech 		string `json:"Job Title"`
	JobReferenceTech 	string `json:"JD Ref"`
	JobSourceTech 		string `json:"Source"`
	JobLinkTech 		string `json:"Link"`
	DescriptionTech		string `json:"descriptionTech"`
	LengthTech			int	`json:descriptionLengthTech`
	MasculineCountTech	int	   `json:"masculineCountTech"`
	FemanineCountTech	int	   `json:"femanineCountTech"`
	RawGenderTech		int	   `json:"rawGenderTech"`

	JobTileGraduate 		string `json:"Job Title__4"`
	JobReferenceGraduate 	string `json:"JD Ref__4"`
	JobSourceGraduate 		string `json:"Source__4"`
	JobLinkGraduate 		string `json:"Link__4"`
	DescriptionGraduate		string `json:"descriptionGrad"`
	LengthGrad				int	`json:descriptionLengthGrad`
	MasculineCountGrad	int	   `json:"masculineCountGrad"`
	FemanineCountGrad	int	   `json:"femanineCountGrad"`
	RawGenderGrad		int	   `json:"rawGenderGrad"`

	JobTileAdmin 			string `json:"Job Title__3"`
	JobReferenceAdmin 		string `json:"JD Ref__3"`
	JobSourceAdmin 		string `json:"Source__3"`
	JobLinkAdmin 		string `json:"Link__3"`
	DescriptionAdmin	string `json:"descriptionAdmin"`
	LengthAdmin			int	`json:descriptionLengthAdmin`
	MasculineCountAdmin	int	   `json:"masculineCountAdmin"`
	FemanineCountAdmin	int	   `json:"femanineCountAdmin"`
	RawGenderAdmin		int	   `json:"rawGenderAdmin"`

	JobTileLegal 		string `json:"Job Title__2"`
	JobReferenceLegal 	string `json:"JD Ref__2"`
	JobSourceLegal 		string `json:"Source__2"`
	JobLinkLegal 		string `json:"Link__2"`
	DescriptionLegal	string `json:"descriptionLegal"`
	LengthLegal			int	`json:descriptionLengthLegal`
	MasculineCountLegal	int	   `json:"masculineCountLegal"`
	FemanineCountLegal	int	   `json:"femanineCountLegal"`
	RawGenderLegal		int	   `json:"rawGenderLegal"`

	JobTileAccounting 		string `json:"Job Title__1"`
	JobReferenceAccounting 	string `json:"JD Ref__1"`
	JobSourceAccounting 	string `json:"Source__1"`
	JobLinkAccounting 		string `json:"Link__1"`
	DescriptionAccounting	string `json:"descriptionAccounting"`
	LengthAccounting		int	`json:descriptionLengthAccounting`
	MasculineCountAccounting	int	   `json:"masculineCountAccounting"`
	FemanineCountAccounting	int	   `json:"femanineCountAccounting"`
	RawGenderAccounting		int	   `json:"rawGenderAccounting"`
}