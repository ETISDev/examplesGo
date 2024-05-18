package main

import (
    "html/template"
    "os"
)

func main() {
    // Create a new template.
    t := template.New("index.html")

    // Parse the template.
    t, err := t.Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Golang HTML File</title>
</head>
<body>
    <h1>Hello, World!</h1>
	<p>Hi Harshal</p>
</body>
</html>
`)
    if err != nil {
        panic(err)
    }

    // Create a new file.
    f, err := os.Create("index.html")
    if err != nil {
        panic(err)
    }

    // Write the template to the file.
    err = t.Execute(f, nil)
    if err != nil {
        panic(err)
    }

    // Close the file.
    f.Close()
}