from ui import setup_window
from volume_control import show_open_apps, get_installed_apps, list_installed_apps


def main():
	import toml
	config = toml.load('config.toml')

	#show_open_apps()
	#window = setup_window()
	#window.mainloop()
	#apps = get_installed_apps()
	#print(*apps, sep="\n")

	#installed_apps = list_installed_apps()
	#print("List of Installed Apps:")
	#for app in installed_apps:
	#	print(app)

	print(config.get("knob2").get("apps"))


if __name__ == '__main__':
	#main()

	show_open_apps()