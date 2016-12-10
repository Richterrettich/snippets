package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	setup()
	createTar("../testfiles", "../out/go/golang.tar")
	extractTar("../out/go/golang.tar", "../out/go")
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func setup() {
	if exists("../out/go") {
		os.Remove("../out/go")
	}
	_ = os.MkdirAll("../out/go", os.ModePerm)
}

func exists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		try(err)
		return true
	}
	return false
}

func createTar(src, target string) {
	out, err := os.Create(target)
	try(err)
	defer out.Close()

	tarball := tar.NewWriter(out)

	// header -> content -> header -> content ... EOF

	base := filepath.Dir(src)
	fmt.Println(base)

	handler := func(path string, fileInfo os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())

		pathInTar, err := filepath.Rel(base, path)

		if err != nil {
			return err
		}

		fmt.Println(pathInTar)

		header.Name = pathInTar

		err = tarball.WriteHeader(header)

		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		_, err = io.Copy(tarball, file)

		return err

	}
	filepath.Walk(src, handler)
	try(tarball.Flush())
	try(tarball.Close())
}

func extractTar(src, dest string) {
	_ = os.MkdirAll(dest, os.ModePerm)

	file, err := os.Open(src)

	try(err)
	defer file.Close()

	tarball := tar.NewReader(file)

	for {
		header, err := tarball.Next()

		if err == io.EOF {
			break
		}
		try(err)
		info := header.FileInfo()
		path := dest + "/" + header.Name
		fmt.Println("processing " + path)
		if info.IsDir() {
			fmt.Println(path + " is a directory")
			try(os.MkdirAll(path, os.ModePerm))
			continue
		}
		fmt.Println("writing file " + path)
		file, err = os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		try(err)
		defer file.Close()
		fmt.Println("adding content to file " + path)
		_, err = io.Copy(file, tarball)
		try(err)
	}

}
