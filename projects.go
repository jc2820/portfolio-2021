package main

type project struct {
	Title       string
	Description string
	Stack       string
	Url         string
	Codebase    string
	Imagepath   string
}

type projectCategory struct {
	CategoryTitle string
	Projects      []project
}

var AllProjects = []projectCategory{
	{"Selected Projects // Professional", proProjects},
	{"Selected Projects // Personal", personalProjects},
	{"Selected Projects // Founders and Coders", facProjects},
}

var proProjects = []project{
	{
		"al-redux",
		"A portfolio website for 3D artist Alex Jackson.",
		"React, Next, Netlify CMS",
		"https://al-redux.com",
		"https://github.com/al-jackson/al-redux",
		"/static/images/al-redux.png",
	},
	{
		"Sienna",
		"A portfolio website for Sienna bringing together projects that Bethany Scott and I have worked on together during 2020 and 2021. We began using the name Sienna in late 2020 with the aim of formalising our collaboration on tech for good projects as an independent agency.",
		"React, Next, Vercel, React-Canvas",
		"https://siennadev.com",
		"https://github.com/sienna-development/sienna-website",
		"/static/images/sienna.png",
	},
	{
		"CyberSafe Tool for Schools",
		"Made for charity CyberSafeKids, the Tool for Schools is a survey app aimed at Irish schools. By surveying a strata of school leaders, Teachers and Pupils through the app, CyberSafeKids are able to provide reports and advice on how a school can improve their 'CyberSafety'.",
		"Gatsby, React, Firebase (Firestore and Auth), Netlify Lambda Functions, Netlify CMS, Salesforce REST API",
		"https://cybersafetoolforschools.ie",
		"https://github.com/cybersafe-dev/cybersafety-tool",
		"/static/images/toolforschools.png",
	},
	{
		"My Best Life Phase 2: Chrome Extension Prototype",
		"Created as part of 'My Best Life' Phase 2. This prototype is a chrome extension that provides a popup with helpful links around youth issues when conditions are met in the page. The drop-down from the context menu is non-functional, but describes an area where users can log in to edit preferences, or add their own helpful links for others.",
		"HTML, CSS, JavaScript, Google Chrome",
		"",
		"https://github.com/jc2820/mbl-popup",
		"/static/images/mbl-popup.png",
	},
	{
		"Just Like Us",
		"Part of the Catalyst 'Dev-2A' programme. Sienna helped LGBT+ charity Just Like Us to research, prototype and user test resource database and LMS services for their new website.",
		"Miro, Figma Thinkific LMS, Learndash LMS, Wordpress, HackMD",
		"https://hackmd.io/@sienna/rydPvsrl_",
		"",
		"/static/images/jlu.png",
	},
	{
		"Catalyst Discovery Learning Programme",
		"As part of a team from Founders and Coders, and a digital partner on Catalyst's 'Discovery Learning Programme', we led nine charities through the discovery process of digital service design.",
		"Zoom, Miro, Gitbook",
		"https://founders-and-coders.gitbook.io/pm-curriculum/",
		"",
		"/static/images/dlp.png",
	},
	{
		"CyberSafe Family Quiz",
		"Created for CyberSafeKids, the family quiz was originally created for an event called 'Cyberbreak' in partnership with Irish bank PTSB and later adapted for more general use. A short quiz returns families tips on how to improve their cyber-awareness, and allows CyberSafeKids a simple UI through Airtable to analyse results.",
		"React, Netlify Lambda Functions, Airtable",
		"https://cyberbreak.ie/",
		"https://github.com/cybersafe-dev/cyberbreak",
		"/static/images/family-quiz.png",
	},
	{
		"Collaborative Digital Training: Resource Centre",
		"An MVP created as part of a Founders and Coders collaboration with CAST as COVID-19 response. CDT aimed to investigate the viability of a collaborative approach to upskilling charities in product management and digital service design.",
		"React, Netlify Lambda Functions, Airtable",
		"https://cdt-resourcecentre.netlify.app/",
		"https://github.com/FAC-CDT/resource-centre",
		"/static/images/resource-centre.png",
	},
}

var personalProjects = []project{
	{
		"FileCrypt",
		"A command line file encryption program with options to encrypt and decrypt files as well as amend text to files without requirement for decryption as a separate job.",
		"Go",
		"",
		"https://github.com/jc2820/filecrypt",
		"/static/images/cmd.png",
	},
	{
		"Go JWT & Cookies",
		"Implementing a very basic web app to try out JWT based authorisation via cookie storage for a private route using Go.",
		"Go",
		"",
		"https://github.com/jc2820/gocookies",
		"/static/images/gojwtcookies.png",
	},
	{
		"NodeJs Speedtest",
		"Resulting from suspicion about regular drops in my home internet speed. While running locally, this app uses a cron job to test download speed at a set interval, storing the results in a local database. Views allow results to be separated according to different speed thresholds.",
		"Node.js, Express, EJS templates, Postgresql, Bootstrap",
		"",
		"https://github.com/jc2820/nodejs-speedtest",
		"/static/images/nodejs-speedtest.png",
	},
	{
		"Discord Timecop",
		"A response for friends who always say they'll be 'just 5 minutes'. The Timecop bot waits on several variations of a given time to wait in group chat messages, and replies with a stern reminder when that time is up.",
		"Discord, JavaScript, Node.js",
		"",
		"https://github.com/jc2820/discord-timecop",
		"/static/images/discord-timecop.png",
	},
}

var facProjects = []project{
	{
		"Reuse, Reduce, Recycle",
		"A drag-and-drop, tablet-first game designed to teach kids about recycling and the environment. A Founders and Coders student selected project and introduction to using the Scrum framework in practice.",
		"React, Styled-Components, Drag and Drop with touch support, PWA, TravisCI, Codecov, Jest, Netlify",
		"https://reduce-reuse-recycle.netlify.app/",
		"https://github.com/jc2820/recycling-game",
		"/static/images/recycling-game.png",
	},
	{
		"Cryptowatch",
		"A cryptocurrency converter app powered by Coin Gecko API with route testing and code coverage. A Founders and Coders student project and introduction to NodeJS.",
		"HTML, CSS, JavaScript, Node.js, Coin Gecko API, Tape, Supertest, TravisCI, Codecov, Heroku",
		"https://cryptowatch.herokuapp.com/",
		"https://github.com/jc2820/week5-famk-backend-api",
		"/static/images/crpytowatch.png",
	},
	{
		"Calculator",
		"A vanilla JavaScript calculator made as a pre-course assignment for Founders and Coders.",
		"HTML, CSS, JavaScript",
		"https://jc2820.github.io/calculator/",
		"https://github.com/jc2820/calculator",
		"/static/images/calculator.png",
	},
}
