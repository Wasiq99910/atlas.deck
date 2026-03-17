# 🎛️ atlas.deck - Control Workflows From Terminal

[![Download atlas.deck](https://img.shields.io/badge/Download-atlas.deck-brightgreen?style=for-the-badge)](https://github.com/Wasiq99910/atlas.deck/releases)

---

## 🖥️ What is atlas.deck?

atlas.deck is a tool that turns your terminal into an interactive control panel. It shows a grid of pads, each linked to a command you want to run. You can organize, start, and watch complex tasks right from your terminal window. It is part of the Atlas Suite, designed to help automate workflows easily.  

You don’t need to write code or open many windows. Instead, you press the pads to trigger commands, making your work simpler and faster.

---

## 🔍 Features

- Customizable grid of pads mapped to shell commands  
- View the status of each command  
- Run multiple commands in one place  
- Works inside your current terminal window (no extra apps)  
- Cross-platform, but this guide focuses on Windows  
- Designed to improve daily workflow and automation  
- Part of a larger set of tools called Atlas Suite

---

## 📋 System Requirements

Make sure your computer meets these:

- Windows 10 or later (64-bit recommended)  
- A terminal program like Command Prompt, PowerShell, or Windows Terminal  
- About 100 MB of free disk space  
- Internet connection for the initial download  
- Basic knowledge of running programs from your computer

---

## 🚀 Getting Started: Download and Run atlas.deck on Windows

1. Click the big button at the top or visit the releases page here:  
   [https://github.com/Wasiq99910/atlas.deck/releases](https://github.com/Wasiq99910/atlas.deck/releases)  

2. Look for the latest release version. It will be named something like `atlas.deck_windows_amd64.zip` or `atlas.deck.exe`.  

3. Download the Windows version by clicking on it. Save the file to a folder you can easily find, like your Desktop or Downloads.

4. If the file is a ZIP (.zip), right-click on it and choose “Extract All.” Select a folder to extract the files to.

5. Open the folder with the extracted files.

6. Find `atlas.deck.exe` (or simply `atlas.deck`) and double-click it to run.

7. The program launches inside the terminal window. You’ll see the grid of pads.

---

## ⌨️ How to Use atlas.deck

- Each square in the grid is called a "pad."  
- Pads are connected to commands you want to run.  
- Use arrow keys or tab to move between pads.  
- Press Enter or Space to trigger a pad’s command.  
- Watch the output or status shown after running the command.

---

## 🎛️ Customizing Your Pads

Before using atlas.deck, you can set up which commands run when you press each pad.

1. Find the config file in the program folder. It often has a name like `config.yaml` or `pads.json`.

2. Open this file in Notepad or any text editor.

3. Each pad entry lists a shell command, label, and optionally a color.

4. Change commands or labels to suit your workflow. For example:
   ```
   - label: Build Project
     command: dotnet build
   - label: Run Tests
     command: dotnet test
   ```

5. Save changes and restart atlas.deck to load new settings.

---

## 🛠️ Troubleshooting

- If atlas.deck does not open, check that you have the right version for Windows.
- Make sure your terminal window supports colors and interactive input.
- If commands don’t run, verify the commands work manually in your terminal.
- Restart atlas.deck if the grid does not respond.
- If the app crashes, try re-downloading from the official releases page.

---

## 💡 Tips for Using atlas.deck

- Use simple commands you already know.  
- Start with a few pads and add more as you get comfortable.  
- Group similar tasks together in the grid for easy access.  
- Check the output of commands for errors directly in the terminal.  
- Use atlas.deck as a launchpad for your daily work routines.

---

## ⬇️ Download and Setup

Download and install atlas.deck by following these steps:

1. Visit the releases page here:  
   [https://github.com/Wasiq99910/atlas.deck/releases](https://github.com/Wasiq99910/atlas.deck/releases)

2. Find the latest version for Windows.

3. Download the `.exe` or `.zip` file.

4. If zipped, extract all files.

5. Run `atlas.deck.exe`.

6. Use arrow keys to move between pads and Enter to run commands.

7. To customize, edit the config file and restart the app.

---

## ❓ Where to Get Help

Use the GitHub repository to report issues or ask questions. Search the issues tab to find answers from other users or the developers.

Link: https://github.com/Wasiq99910/atlas.deck/issues

---

## 📝 About This Project

atlas.deck is built with Go and uses terminal UI libraries like Bubble Tea and Lip Gloss. It aims to make command-line workflows easier by providing a visual, interactive interface. If you use other Atlas Suite tools, atlas.deck fits right into your productivity system.

---

## 🔎 Related Topics

- atlas-suite  
- automation-tool  
- terminal-ui (TUI)  
- workflow automation  
- developer-tools  

---

[![Download atlas.deck](https://img.shields.io/badge/Download-atlas.deck-green?style=for-the-badge)](https://github.com/Wasiq99910/atlas.deck/releases)