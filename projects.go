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
		"al-redux",
		"Al's site",
		"React, Next, Netlify CMS",
		"https://al-redux.com",
		"https://github.com/al-jackson/al-redux",
		"/static/images/al-redux.png",
	},
	{
		"Sienna",
		"Our Webpage",
		"React, Next, Vercel",
		"https://siennadev.com",
		"https://github.com/sienna-development/sienna-website",
		"/static/images/sienna.png",
	},
	{
		"CyberSafe Tool for Schools",
		"tfs",
		"Gatsby, React, Firebase (Firestore and Auth), Netlify Lambda Functions, Netlify CMS",
		"https://toolforschools.ie",
		"",
		"/static/images/toolforschools.png",
	},
	{
		"Just Like Us",
		"tfs",
		"Figma, Miro, Wordpress, Thinkific LMS, Learndash LMS, HackMD",
		"",
		"",
		"/static/images/jlu.png",
	},
	{
		"Discovery Learning Programme",
		"tfs",
		"Zoom, Miro, Gitbook",
		"",
		"",
		"/static/images/dlp.png",
	},
	{
		"NodeJs Speedtest",
		"tfs",
		"NodeJs, Express, EJS templates, Postgresql",
		"",
		"",
		"/static/images/nodejs-speedtest.png",
	},
	{
		"My Best Life: Chrome Extension Prototype",
		"tfs",
		"HTML, CSS, Javascript",
		"",
		"",
		"/static/images/mbl-popup.png",
	},
	{
		"Discord Timecop",
		"tfs",
		"Javascript, NodeJs",
		"",
		"",
		"/static/images/discord-timecop.png",
	},
	{
		"CyberSafe Family Quiz",
		"tfs",
		"React",
		"",
		"",
		"/static/images/family-quiz.png",
	},
	{
		"Resource Centre",
		"tfs",
		"React, Netlify Lambda Functions, Airtable",
		"",
		"",
		"/static/images/resource-centre.png",
	},
	{
		"Reuse, Reduce, Recycle",
		"A drag-and-drop game designed to teach kids about recycling and the environment. A Founders and Coders student selected project and introduction to Scrum methodology.",
		"React, Styled-Components, Drag and Drop with touch support, PWA, TravisCI, Codecov, Jest, Netlify",
		"",
		"",
		"/static/images/recycling-game.png",
	},
	{
		"Cryptowatch",
		"A cryptocurrency converter app powered by Coin Gecko API with route testing and code coverage. A Founders and Coders project.",
		"HTML, CSS, JavaScript, Node.js, Coin Gecko API, Tape, Supertest, TravisCI, Codecov, Heroku",
		"",
		"",
		"/static/images/crpytowatch.png",
	},
	{
		"Calculator",
		"A vanilla JavaScript calculator made as a pre-course assignment for Founders and Coders.",
		"HTML, CSS, JavaScript",
		"",
		"",
		"/static/images/calculator.png",
	},
}
