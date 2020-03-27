package templates

var Table string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Filename}}</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css"
          integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>
<body>
<div class="container">
    <table class="table striped">
        <h1>{{.Filename}}</h1>
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
