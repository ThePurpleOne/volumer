import tkinter as tk
from ttkthemes import ThemedTk

def setup_window():
	window = tk.Tk()
	window.title("Volume Control")
	window.geometry("300x300")

	# ! COMBOBOX FOR APPS
	apps = [app.Process.name() for app in get_open_apps() if app.Process]
	app_var = tk.StringVar(window)
	app_var.set(apps[0])
	app_selector = tk.OptionMenu(window, app_var, *apps)
	app_selector.pack()

	# ! SLIDER
	volume_slider = tk.Scale(window, from_=0, to=100, orient=tk.HORIZONTAL)
	volume_slider.set(50)

	def slider_callback(event):
		slider_value = volume_slider.get()
		selected_app = app_var.get()
		set_app_volume(selected_app, slider_value / 100)

	volume_slider.bind("<ButtonRelease-1>", slider_callback)
	volume_slider.pack()

	return window



def get_open_apps():
	from pycaw.pycaw import AudioUtilities
	return AudioUtilities.GetAllSessions()

def show_open_apps():
	for session in get_open_apps():
		try:
			volume = session.SimpleAudioVolume
			if session.Process:
				print(f"{volume.GetMasterVolume():.02f} | {session.Process.name()}")
		except Exception as e:
			print(e)

def set_app_volume(app_name, volume):
	from pycaw.pycaw import AudioUtilities, IAudioEndpointVolume

	sessions = get_open_apps()

	for session in sessions:
		if session.Process and session.Process.name() == app_name:
			app = session.SimpleAudioVolume
			app.SetMasterVolume(volume, None)
			return
		
	raise Exception(f"App {app_name} not found")




def main():
	show_open_apps()
	window = setup_window()
	window.mainloop()


if __name__ == '__main__':
	main()