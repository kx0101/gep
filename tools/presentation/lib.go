package presentation

import (
	"fmt"
	"log"
	"os/exec"
)

func CreatePresentation() {
	fmt.Println("Creating presentation...")

	cmd := exec.Command("python3", "create_ppt.py")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Presentation.pptx has been created!")
}
