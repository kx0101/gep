package presentation

import (
	"log"
	"os/exec"
)

func CreatePresentation() {
	log.Println("Creating presentation...")

	cmd := exec.Command("python3", "scripts/create_ppt.py")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Presentation.pptx has been created!")
}
