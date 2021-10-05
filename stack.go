package main

var Languages = []string{"JavaScript (ES6+)", "Go", "Semantic HTML", "CSS"}
var Libraries = []string{"Node.js", "npm", "React", "Express", "Next", "Gatsby", "Styled Components", "Firebase"}
var Databases = []string{"Postgresql", "Firestore", "Airtable"}
var Cicd = []string{"Netlify", "Heroku", "Vercel", "Github", "TravisCI"}
var Tools = []string{"Miro", "HackMD", "Figma"}
var Design = []string{"Adobe Photoshop", "Adobe Illustrator", "Adobe InDesign", "Adobe Premiere"}
var Testing = []string{"Jest", "React Testing Library", "Codecov"}
var Methodologies = []string{"Agile", "Scrum", "TDD", "Accessibility", "Responsive Design", "Prototyping", "Pair Programming"}

type category struct {
	Categorytitle string
	Categorylist  []string
}

var Stack = []category{
	{
		"Languages",
		Languages,
	},
	{
		"Libraries",
		Libraries,
	},
	{
		"Databases",
		Databases,
	},
	{
		"CI/CD",
		Cicd,
	},
	{
		"Tools",
		Tools,
	},
	{
		"Design",
		Design,
	},
	{
		"Testing",
		Testing,
	},
	{
		"Methodologies",
		Methodologies,
	},
}
