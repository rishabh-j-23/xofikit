# xofikit

xofikit is a CLI utility that enables interactive shell script execution through Rofi using YAML-based configuration files. It allows users to quickly trigger scripts via a menu interface, making it ideal for system actions, quick toggles, and workflow shortcuts.

---

## Features

* Define menus and nested submenus using YAML  
* Run commands via a Rofi-powered interface  
* Support for multiple configuration files  
* Easily extendable with new scripts and options  
* Integration-ready with tiling window managers like Hyprland, i3, and others  

---

## Installation

### Prerequisites

* Go (version 1.21 or higher)  
* Rofi  
* make  

### Steps

```bash
git clone https://github.com/rishabh-j-23/xofikit.git
cd xofikit
make install
````

---

## Creating a YAML Config

Each YAML file defines a menu with one or more options. You can define nested submenus using the `suboptions` field. Each option can be associated with a shell command that will be executed when selected.

### Basic Format

```yaml
prompt: "Your menu prompt"
options:
  - code: "unique_code"
    label: "Display Name"
    command: |
      your-shell-command
```

### Nested Submenus

To create nested menus, use `suboptions` instead of `command`:

```yaml
prompt: "Top Menu"
options:
  - code: "parent"
    label: "Parent Option"
    suboptions:
      - code: "child1"
        label: "Child 1"
        command: |
          echo "Running Child 1"
      - code: "child2"
        label: "Child 2"
        command: |
          echo "Running Child 2"
```

### Field Reference

| Field        | Description                                              |
| ------------ | -------------------------------------------------------- |
| `prompt`     | The prompt displayed at the top of the Rofi menu         |
| `code`       | A unique identifier for the menu option                  |
| `label`      | The text shown in the Rofi menu for this option          |
| `command`    | Shell command to be executed when the option is selected |
| `suboptions` | A list of more options, forming a nested submenu         |

> \[!NOTE]
> In case of suboptions menu, command for parent option won't execute.

### Tips

* Use `|` after `command:` for multi-line shell commands.
* All YAML files must end in `.yaml` and reside in `~/.config/xofikit/scripts_config/`.
* Avoid duplicate `code` values even across submenus.

### Example File: `bluetooth_ctl.yaml`

```yaml
prompt: "Bluetooth Options"
options:
  - code: "bluetooth"
    label: "Bluetooth Control"
    suboptions:
      - code: "on"
        label: "Turn Bluetooth On"
        command: |
            bluetoothctl power on
      - code: "off"
        label: "Turn Bluetooth Off"
        command: |
            bluetoothctl power off
```

---

## Running Config Files

After installation, you can run xofikit with any YAML configuration file to display the interactive menu.

### Basic Run Command

```bash
xofikit run /path/to/your/config.yaml
```

For example, to run the Bluetooth control config:

```bash
xofikit run ~/.config/xofikit/scripts_config/bluetooth_ctl.yaml
```

This will launch the Rofi menu based on the YAML and let you interactively choose options.

### Binding to a Hotkey

You can bind xofikit commands to your window manager or compositor’s keybindings to quickly launch menus. For example, in Hyprland:

```
bind = $mainMod, B, exec, xofikit run ~/.config/xofikit/scripts_config/bluetooth_ctl.yaml
```

Make sure xofikit and your config file are accessible by the environment where the hotkey runs.

---

For more details, visit the [project repository](https://github.com/rishabh-j-23/xofikit).

---

## Contributing

Contributions are welcome! If you would like to improve this project, feel free to fork the repository and open a pull request.

### Steps to Contribute

1. Fork the repository
2. Create a new branch (`git checkout -b feature-name`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add new feature'`)
5. Push to the branch (`git push origin feature-name`)
6. Open a pull request describing your changes
