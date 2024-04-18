package parser

import (
	"io"
	"strings"
)

type Video struct {
	Name      string
	VideoLink string
	AudioLink string
}

type Parser struct {
}

func splitByCharacter(content string, character string) []string {
	return strings.Split(content, character)
}

func (p *Parser) ParseFile(file *io.Reader) ([]Video, error) {
	content, err := io.ReadAll(*file)

	if err != nil {
		return nil, err
	}

	files := splitByCharacter(string(content), "-----\n")

	var fileContent []Video = make([]Video, len(files))

	for i, file := range files {
		fileParts := splitByCharacter(file, "\n")

		if len(fileParts) > 3 {
			fileParts = fileParts[:len(fileParts)-1]
		}

		video := Video{
			Name:      fileParts[0],
			VideoLink: splitByCharacter(fileParts[1], "::")[1],
			AudioLink: splitByCharacter(fileParts[2], "::")[1],
		}

		fileContent[i] = video
	}

	return fileContent, nil
}
