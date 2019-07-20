package main

import "on-server/fs"

const top string = `<!doctype html>
<head>
  <title>File server</title>
  <meta name="viewport" content="width=device-width, initial-scale=0.5">
<style>` +
	`body{margin: 0;}
.container{
	height: 100vh;
	justify-content: center;
}
.flex{
	display: flex;
}

.row{
	flex-direction: row;
}

.column{
	flex-direction: column;
}
.msg{
	font-size: 16px;
	padding: 3px;
	width: 50vw;
}
input[type=file]{
	margin: 8px 0;
}

#data-form{
	align-items: center;
	justify-content: space-around;
}
button[type=submit]{
	width: 60vw;
	margin: 0 auto;
	font-size: 16px;
	padding: 6px;
}

.files{
	flex: 1;
	border: 0;
	border-top: 2px outset #7a7f82;
}
` +
	`</style>
</head>
<body>`

func makeBody(msg bool, upload bool, files bool) string {
	content := `<div class="flex column container">`
	if msg || upload {
		content += `<form id="data-form" action="/upload" method="post" enctype="multipart/form-data" class="flex row">`
		if msg {
			content += `<textarea class="msg" name="message" rows="4" placeholder="Type your message here."></textarea>`
		}
		if upload {
			content += `<input type="file" name="file" multiple>`
		}
		content += `</form>
		<button type="submit" form="data-form">Send data</button>`
	}

	if files {
		content += `<iframe class="files" src="` + fs.URLRoot + `" title="files">`
	}

	content += `</div>`

	return content
}

const bottom string = `</body>`
