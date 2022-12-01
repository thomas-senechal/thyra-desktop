package main

import (
	_ "embed"
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed logo.png
var logo []byte

func main() {
	a := app.New()
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("Thyra Desktop | Logs")
	cmd := exec.Command("thyra-server")
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Fatal(err)
	}

	if desk, ok := a.(desktop.App); ok {
		startMenu := fyne.NewMenuItem("Start", nil)
		stopMenu := fyne.NewMenuItem("Stop", nil)
		logsMenu := fyne.NewMenuItem("Show logs", nil)
		walletMenu := fyne.NewMenuItem("Wallet", nil)
		registryMenu := fyne.NewMenuItem("Registry", nil)
		aboutMenu := fyne.NewMenuItem("About", nil)

		m := fyne.NewMenu("Thyra Desktop",
			walletMenu,
			registryMenu,
			fyne.NewMenuItemSeparator(),
			startMenu,
			stopMenu,
			fyne.NewMenuItemSeparator(),
			aboutMenu,
		)

		startMenu.Action = func() {
			err := cmd.Start()

			if err != nil {
				log.Fatal(err)
			} else {
				startMenu.Disabled = true
				stopMenu.Disabled = false
				m.Refresh()
				log.Println("Server started")
			}
		}
		stopMenu.Action = func() {
			if cmd.Process != nil {
				err := cmd.Process.Kill()

				if err != nil {
					log.Fatal(err)
				} else {
					startMenu.Disabled = false
					stopMenu.Disabled = true
					m.Refresh()
					log.Println("Server stopped")
				}
			}
		}
		logsMenu.Action = func() {
			logs := []byte{}
			if cmd.ProcessState != nil {
				for {
					tmp := make([]byte, 1024)
					_, err := stdout.Read(tmp)
					logs = append(logs, tmp...)
					if err != nil {
						break
					}
				}
			}

			if len(logs) > 0 {
				w.SetContent(widget.NewLabel(string(logs)))
			} else {
				w.SetContent(widget.NewLabel("No logs to show"))
			}
			w.Show()
		}
		walletMenu.Action = func() {
			openURL(&a, "http://my.massa/thyra/wallet")
		}
		registryMenu.Action = func() {
			openURL(&a, "http://my.massa/thyra/registry")
		}
		aboutMenu.Action = func() {
			createAboutWindow(&a).Show()
		}

		stopMenu.Disabled = true

		icon := fyne.NewStaticResource("logo", logo)
		desk.SetSystemTrayIcon(icon)
		desk.SetSystemTrayMenu(m)
	}

	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.Resize(fyne.NewSize(600, 400))

	defer func() {
		log.Println("Closing...")
		if cmd.Process != nil {
			log.Println("Stopping server...")
			err := cmd.Process.Kill()
			if err != nil {

				log.Fatal(err)
			}
			_, err = cmd.Process.Wait()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Server stopped")
		}
	}()

	a.Run()
}
