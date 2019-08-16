package data

// Data returns mock html page
func Data() string {
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Scrapper</title>
	</head>
	<body>
		<h1>Hello goquery</h1>
		<div>
			<form action="">
				<input type="text"><br>
				<input type="checkbox" name="yes" id="">
				<button>Submi</button>
			</form>
			<div>
				<img src="/toby.png" alt="">
				<ul>
					<li><a href="http.google.com">Google</a></li>
					<li><a href="http.bbc.com">BBC</a></li>
					<li><a href="/com">Yahoo</a></li>
					<li><a href="http.stack.com">Stack</a></li>
					<li><a href="http.random.com">Random</a></li>
				</ul>
			</div>
		</div>
	</body>
	</html>`
}
