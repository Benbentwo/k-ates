package templates

var K8Table = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Filename}}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<style>
ul.breadcrumb {
  padding: 10px 16px;
  list-style: none;
  background-color: #eee;
}
ul.breadcrumb li {
  display: inline;
  font-size: 18px;
}
ul.breadcrumb li+li:before {
  padding: 8px;
  color: black;
  content: "/\00a0";
}
ul.breadcrumb li a {
  color: #0275d8;
  text-decoration: none;
}
ul.breadcrumb li a:hover {
  color: #01447e;
  text-decoration: underline;
}
</style>
</head>
<body>
<div class="container">
	{{ if .BreadCrumbs }}
		<ul class="breadcrumb">
			{{ range .BreadCrumbs }}
				{{ if .Href }}
					<li><a href="{{.Href}}">{{.Text}}</a></li>
				{{ else }}
					<li>{{.Text}}</li>
				{{ end }}
			{{ end }}
		</ul>
	{{ end }}
    <table class="table striped">
        <h1>{{.Filename}}</h1>
		{{ range .Buttons }}
			<a href="{{.Href}}"><button>{{.Text}}</button></a>
		{{end}}
        <thead>
        <tr>
            {{range .Headers}}
            <th>
                {{.}}
            </th>
            {{end}}
        </tr>

        </thead>
        <tbody>
        {{range .Rows}}
        <tr>
            {{range .}}
            <td>{{printf "%s" .}}</td>
            {{end}}
        </tr>
        {{end}}
        </tbody>
    </table>
</div>
</body>
</html>
`
