/**
 * Create a Munki manifest based on a template
 * License: Apache 2.0
 * Author: Antti Pettinen / Intelligent Apps GmbH 
 * Last Modified Date: 08.05.2019
 * Last Modified By: Antti Pettinen <antti.pettinen@gmail.com>
*/

package main

import (
	"os"
	"github.com/groob/plist"
	"flag"
	"fmt"
	"path/filepath"
)

type manifestTemplate struct {
	Catalogs []string `plist:"catalogs"`
	IncludedManifests []string `plist:"included_manifest"`
	ManagedInstalls []string `plist:"managed_installs"`
	ManagedUninstalls []string `plist:"managed_uninstalls"`
	ManagedUpdates []string `plist:"managed_updates"`
	OptionalInstalls []string `plist:"optional_installs"`
	PrimaryUser string `plist:"user"`

}

func main() {
	// get current working directory:
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// define our command line flags
	var basePath string
	var replaceManifest bool
	flag.StringVar(&basePath, "path", curDir, "Directory to store manifests")
	flag.BoolVar(&replaceManifest, "replace", false, "Replace existing manifests")
	flag.Parse()
	manifests := flag.Args()

	if len(manifests) == 0 {
		fmt.Println("Missing output filename, please provide at least one!")
		flag.Usage()
		return
	} else {
		fmt.Println("Creating following manifests in:", basePath, manifests)
	}

	manifestData := manifestTemplate{
		Catalogs:          []string{"production"},
		IncludedManifests: []string{"standard-user"},
	}

	// Loop through the filenames and write our template
	for _, manifestFile := range manifests {
		fullPath := filepath.Join(basePath, manifestFile)
		// Try to open file, requiring that it does not exists:
		f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)
		// if the file exists, we have an error and need to handle that
		if err != nil {
			// replace if desired
			if replaceManifest == true {
				fmt.Println("Replacing manifest at ", fullPath)
				f, err = os.Create(fullPath)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println("Skipping existing manifest at ", fullPath)
				continue
			}
		}
		defer f.Close()
		// encode the XML file
		plist.NewEncoder(f).Encode(manifestData)
		if err != nil {
			panic(err)
		}
		fmt.Println("Created manifest:", manifestFile)
	}
}
