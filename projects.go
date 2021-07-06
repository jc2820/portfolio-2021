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
		"A portfolio website for 3D artist Alex Jackson. New posts to each page can be made by an admin via Netlify CMS. Updates trigger a fresh static rebuild using NextJS.",
		"React, Next, Netlify CMS",
		"https://al-redux.com",
		"https://github.com/al-jackson/al-redux",
		"/static/images/al-redux.png",
	},
	{
		"Sienna",
		"A portfolio website for Sienna bringing together projects that Bethany Scott and I have worked on together since graduating from Founders and Coders. We began using the name Sienna in late 2020 with the aim of formalising our collaboration on tech for good projects as an independent agency.",
		"React, Next, Vercel, React-Canvas",
		"https://siennadev.com",
		"https://github.com/sienna-development/sienna-website",
		"/static/images/sienna.png",
	},
	{
		"Just Like Us",
		"Part of the Catalyst Dev-2A programme. LGBT+ charity Just Like Us required a team to research, prototype and user test resource database and LMS services for their new website.",
		"Miro, Figma Thinkific LMS, Learndash LMS, Wordpress, HackMD",
		"",
		"",
		"/static/images/jlu.png",
	},
	{
		"Catalyst Discovery Learning Programme",
		"As part of a team from Founders and Coders, and a digital partner on Catalyst's Discovery Learning Programme, we led nine charities through the discovery process of digital service design.",
		"Zoom, Miro, Gitbook",
		"",
		"",
		"/static/images/dlp.png",
	},
	{
		"CyberSafe Tool for Schools",
		"Made for charity CyberSafeKids, the Tool for Schools is a survey app aimed at Irish schools. By surveying a strata of school leaders, Teachers and Pupils through the app, CyberSafeKids are able to provide reports and advice on how a school can improve their 'CyberSafety' and provide a mark that schools can display on their comms material.",
		"Gatsby, React, Firebase (Firestore and Auth), Netlify Lambda Functions, Netlify CMS",
		"https://toolforschools.ie",
		"",
		"/static/images/toolforschools.png",
	},
	{
		"NodeJs Speedtest",
		"A personal project that came from concern about regular drops in my home internet speed. While running locally, this app uses a cron job to test download speed at a set interval, storing the results in a local database. Views allow results to be separated according to different speed thresholds.",
		"NodeJs, Express, EJS templates, Postgresql, Bootstrap",
		"",
		"",
		"/static/images/nodejs-speedtest.png",
	},
	{
		"My Best Life Phase 2: Chrome Extension Prototype",
		"Created as part of My Best Life Phase 2. This prototype is a chrome extension that provides a popup with helpful links around youth issues when conditions are met in the page. The drop-down from the context menu is non-functional, but describes an area where users can log in to edit preferences, or add their own helpful links for others.",
		"HTML, CSS, Javascript",
		"",
		"",
		"/static/images/mbl-popup.png",
	},
	{
		"Discord Timecop",
		"A response to those friends who say they'll be '5 minutes' but are always more. The Timecop bot waits on several variations of a given time to wait in group chat messages, and replies with a stern reminder when that time is up.",
		"Discord, Javascript, NodeJs",
		"",
		"",
		"/static/images/discord-timecop.png",
	},
	{
		"CyberSafe Family Quiz",
		"Created for Irish charity CyberSafeKids, the family quiz was originally created for an event called 'Cyberbreak' in partnership with Irish bank PTSB and later adapted for more general use. A short quiz returns families tips on how to improve their cyber-awareness, and allows CyberSafeKids a simple UI through Airtable to analyse results.",
		"React, Netlify Lambda Functions, Airtable",
		"",
		"",
		"/static/images/family-quiz.png",
	},
	{
		"Collaborative Digital Training: Resource Centre",
		"An MVP created as part of a Founders and Coders collaboration with CAST as part of COVID-19 response. Working with Product Owners from two non-profits this programme aimed to be a model for upskilling charities through practical workshops in designing and building digital services that collaboratively tackle .",
		"React, Netlify Lambda Functions, Airtable",
		"",
		"",
		"/static/images/resource-centre.png",
	},
	{
		"Reuse, Reduce, Recycle",
		"A drag-and-drop game designed to teach kids about recycling and the environment. A Founders and Coders student selected project and introduction to using the Scrum framework in practice.",
		"React, Styled-Components, Drag and Drop with touch support, PWA, TravisCI, Codecov, Jest, Netlify",
		"",
		"",
		"/static/images/recycling-game.png",
	},
	{
		"Cryptowatch",
		"A cryptocurrency converter app powered by Coin Gecko API with route testing and code coverage. A Founders and Coders student project and introduction to NodeJS.",
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
