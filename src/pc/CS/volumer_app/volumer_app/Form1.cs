using System;
using System.Diagnostics;
//using NAudio;
//using NAudio.CoreAudioApi;
using CSCore.CoreAudioAPI;
using System.Windows.Forms;

namespace volumer_app
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();

			// Initialize the audio session manager
			// Create session object for every application 
            // To be able to manage their individual audio 
			Process[] processes = Process.GetProcessesByName("Discord");
            foreach (Process p in processes)
            {
                AudioSessionManager2 mgr = GetAudioSessionManager2(p.Id);
                if (mgr == null)
                    continue;

                using (var sessionEnumerator = mgr.GetSessionEnumerator())
                {
                    foreach (var session in sessionEnumerator)
                    {
                        using (var simpleVolume = session.QueryInterface<SimpleAudioVolume>())
                        {
                            if (simpleVolume != null)
                            {
                                simpleVolume.MasterVolume = trackBar1.Value / 100.0f;
                            }
                        }
                    }
                }
            }
        }

        private void trackBar1_Scroll(object sender, EventArgs e)
        {

        }

        private AudioSessionManager2 GetAudioSessionManager2(int pid)
        {
            using (var enumerator = new MMDeviceEnumerator())
            {
                using (var device = enumerator.GetDefaultAudioEndpoint(DataFlow.Render, Role.Multimedia))
                {
                    using (var sessionManager = AudioSessionManager2.FromMMDevice(device))
                    {
                        return sessionManager;
                    }
                }
            }
        }

    }
}
