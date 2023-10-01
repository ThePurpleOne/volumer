
def get_openned_apps():
	from pycaw.pycaw import AudioUtilities
	return AudioUtilities.GetAllSessions()

def show_open_apps():
	for session in get_openned_apps():
		try:
			volume = session.SimpleAudioVolume
			if session.Process:
				print(f"{volume.GetMasterVolume():.02f} | {session.Process.name()}")
		except Exception as e:
			print(e)

def set_app_volume(app_name, volume):
	sessions = get_openned_apps()

	for session in sessions:
		if session.Process and session.Process.name() == app_name:
			app = session.SimpleAudioVolume
			app.SetMasterVolume(volume, None)
			return
		
	raise Exception(f"App {app_name} not found")

def lower_volume(app_name):
	"""
	Lower the volume of the app with the given name by 1%
	"""
	sessions = get_openned_apps()

	for session in sessions:
		if session.Process and session.Process.name() == app_name:
			app = session.SimpleAudioVolume
			volume = app.GetMasterVolume()
			app.SetMasterVolume(volume - 0.01, None)
			return
		
	raise Exception(f"App {app_name} not found")

def raise_volume(app_name):
	"""
	Raise the volume of the app with the given name by 1%
	"""
	sessions = get_openned_apps()

	for session in sessions:
		if session.Process and session.Process.name() == app_name:
			app = session.SimpleAudioVolume
			volume = app.GetMasterVolume()
			app.SetMasterVolume(volume + 0.01, None)
			return
		
	raise Exception(f"App {app_name} not found")

def get_installed_apps():
	"""
	Get all installed apps on Windows using the registry and retrieve executable names
	"""
	import winreg
	apps = []
	registry = winreg.ConnectRegistry(None, winreg.HKEY_LOCAL_MACHINE)
	key = winreg.OpenKey(registry, r"SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall")
	
	for i in range(0, winreg.QueryInfoKey(key)[0] - 1):
		try:
			subkey_name = winreg.EnumKey(key, i)
			subkey = winreg.OpenKey(key, subkey_name)
			
			# Check if the 'DisplayName' exists
			if winreg.QueryValueEx(subkey, "DisplayName")[0] and winreg.QueryValueEx(subkey, "InstallLocation")[0]:
				display_name = winreg.QueryValueEx(subkey, "DisplayName")[0]

				apps.append(display_name)
		except Exception as e:
			print(f"Exception: {e}")
	
	return apps

def list_installed_apps():
	import subprocess
	try:
		# Run the wmic command to list installed apps, excluding Microsoft items
		result = subprocess.run(['wmic', 'product', 'get', 'name'], capture_output=True, text=True, check=True)
		
		# Split the result into lines and exclude lines containing 'Microsoft'
		app_lines = [line.strip() for line in result.stdout.split('\n') if 'Microsoft' not in line]

		# Return the list of installed apps
		return app_lines
	except subprocess.CalledProcessError as e:
		print(f"Error: {e}")