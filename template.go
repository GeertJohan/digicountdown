package main

const page = `<!DOCTYPE html>
<html>
	<head>

		<title>Guldencoin &mdash; DIGISHIELD</title>

		<meta name="viewport" content="width=device-width, initial-scale=1">	
		<meta charset="UTF-8">

		<link rel="stylesheet" type="text/css" media="screen,projection" href="style.css">
		
		<script src="//use.typekit.net/ari7osy.js"></script>
		<script>try{Typekit.load();}catch(e){}</script>
		<script src="//cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.js"></script>

	</head>

	<body>

		<div class="timer">
			<div class="logo"></div>
			{{if .Activated}}
			<h1>Digishield is activated!</h1>
			{{else}}
			<h1>{{.Duration}}</h1>
			<p>
				&nbsp;<br/>
				{{.Blocks}} blocks until activation.<br />
				<br/>
				Estimated time:<br/>
				{{.EstimatedTime}}<br/>
				<br/>
				<small>Duration and time are updated after each block</small>
			</p>
			{{end}}
		</div>

	</body>
</html>`

type pageData struct {
	Activated     bool
	Duration      string
	Blocks        int64
	EstimatedTime string
}

var style = []byte(`
/* -------------------------------------------------------------- CSS Reset */

html, body, div, span, p, a, img, strong, b, form, label {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 1.1em;
	font: inherit;
	vertical-align: baseline;
}

ul,li {
	list-style: none;
}

body {
	line-height: 1;
}

* {
	margin: 0;
	padding: 0;
}

* :focus {
	outline: none;
}

/* -------------------------------------------------------------- Body */

html, body {
	width: 100%;
    height: 100%;
}

body {
	font-family: "proxima-nova", "HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif;
	font-weight: 400;
	font-style: normal;
	font-size: 1.1em;
	line-height: 1.7em;
	width: 100%;
	height: 100%;
	color: #fff;
	background-color: #003366;
	-ms-text-size-adjust: none;
	-webkit-text-size-adjust: none;
    -webkit-font-smoothing: antialiased;
    font-smoothing: antialiased;
    text-rendering: optimizeLegibility;
}

.icon {
	font-family: 'FontAwesome';
	font-weight: 100 !important;
	-webkit-font-smoothing: antialiased !important;
	font-smoothing: antialiased !important;
}

@media (max-width: 1440px) and (min-width: 1280px) {
	body {
		font-size: 1em;
	}
}

@media (max-width: 1280px) and (min-width: 768px) {
	body {
		font-size: .9em;
	}
}

@media (max-width: 768px) {
	body {
		font-size: .8em;
	}
}

/* -------------------------------------------------------------- Fonts */

h1 {
	float: left;
	width: 100%;
	font-weight: 400;
	font-style: normal;
	font-size: 3em;
	line-height: 1.4em;
	font-weight: 500;
	text-align: center;
	color: #fff;
}

/* -------------------------------------------------------------- Logo */

.logo {
	float: left;
	display: block;
	width: 100%;
	height: 50px;
	line-height: 50px;
	background-image: url('https://s3-eu-west-1.amazonaws.com/guldencoin.com/digishield.png');
	background-size: 290px 50px;
	background-position: center center;
	background-repeat: no-repeat;
	z-index: 999;
}

@media (-webkit-min-device-pixel-ratio: 2) {
	.logo { 
		background-image: url('https://s3-eu-west-1.amazonaws.com/guldencoin.com/digishield@2x.png') !important;
	}
}

/* -------------------------------------------------------------- Timer */

.timer {
	position: absolute;
	top: 50%;
	left: 50%;
	margin-left: -150px;
	margin-top: -150px;
	width: 300px;
	height: 300px;
	text-align: center;
	color: #f08239;
}
`)
