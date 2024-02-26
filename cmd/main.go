package main

// TODO: ts build wrong
// TODO: index.ts wrong
// TODO: search for go mod tidy problem.
// TODO: update makefile

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type file struct {
	name    string
	content string
}

func main() {
	// get --name flag
	var name string
	flag.StringVar(&name, "name", "Golosus-Web", "project name")
	var githubProfile string
	flag.StringVar(&githubProfile, "github", "cagrigit-hub", "github user name")
	flag.Parse()

	command := goModInit(name, githubProfile)
	exec.Command("sh", "-c", command).Run()

	ct := &Content{}

	folders := []string{
		"assets",
		"assets/jscode",
		"cmd",
		"view",
		"model",
		"handler",
		"view/components",
		"view/layouts",
		"view/example",
		"typescript",
		".",
	}

	for _, folder := range folders {
		err := createFolders(name + "/" + folder)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	files := map[string][]file{
		"cmd": {
			{"main.go", ct.Main()},
		},
		"view/layouts": {
			{"base.templ", ct.Layout(name)},
		},
		"view/example": {
			{"example.templ", ct.ExampleView(githubProfile, name)},
		},
		"view/components": {
			{"input.templ", ct.ExampleComponent()},
		},
		"handler": {
			{"util.go", ct.Util()},
			{"example.go", ct.ExampleHandler(githubProfile, name)},
		},
		"model": {
			{"example.go", ct.ExampleModel()},
		},
		".": {
			{"go.mod", ct.GoMod(githubProfile, name)},
			{"Makefile", ct.Make()},
		},
		"typescript": {
			{"tsconfig.json", ct.TsConfig()},
			{"index.ts", ct.TypescriptIndex()},
			{"package.json", ct.PackageJson(name)},
		},
	}

	for _, folder := range folders {
		for _, file := range files[folder] {
			err := createFiles(name, folder, file.name)
			if err != nil {
				log.Fatal(err)
				return
			}
			err = writeFiles(name, folder, file.name, file.content)
			if err != nil {
				log.Fatal(err)
				return
			}
		}

	}
}

func createFolders(name string) error {
	if err := os.MkdirAll(name, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func createFiles(rn, target, name string) error {
	if _, err := os.Create(rn + "/" + target + "/" + name); err != nil {
		return err
	}
	return nil
}

func writeFiles(rn, target, name, content string) error {
	if err := os.WriteFile(rn+"/"+target+"/"+name, []byte(content), os.ModePerm); err != nil {
		return err
	}
	return nil
}

func goModInit(name, github string) string {
	// run command
	return fmt.Sprintf("go mod init github.com/%s/%s", github, name)
}
