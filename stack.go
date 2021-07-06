package main

var Languages = []string{"Semantic HTML", "CSS", "Javascript (ES6+)", "Go"}
var LibsFrameworks = []string{"NodeJS", "npm", "React", "Express", "Next", "Gatsby", "Styled Components", "Firebase"}
var Databases = []string{"Postgresql", "Firestore", "Airtable"}
var Cicd = []string{"Netlify", "Heroku", "Vercel", "Github", "TravisCI"}
var Tools = []string{"Miro", "HackMD", "Figma"}
var Design = []string{"Adobe Photoshop", "Adobe Illustrator", "Adobe InDesign", "Adobe Premier"}
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
		"Libraries & Frameworks",
		LibsFrameworks,
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
