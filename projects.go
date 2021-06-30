package main

type project struct {
	Title       string
	Description string
	Stack       string
	Url         string
	Codebase    string
	Imagepath   string
}

var Projects = []project{
	{
		"Sienna",
		"Our Webpage",
		"ReactJS, NextJS",
		"https://siennadev.com",
		"https://github.com/sienna-development/sienna-website",
		"/static/images/sienna.png",
	},
	{
		"al-redux",
		"Al's site",
		"ReactJS, NextJS, Netlify CMS",
		"https://al-redux.com",
		"https://github.com/al-jackson/al-redux",
		"/static/images/al-redux.png",
	},
}
