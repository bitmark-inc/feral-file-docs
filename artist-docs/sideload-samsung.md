# Install the Feral File app on Your Samsung Tizen TV

## What You'll Need:

- **Samsung Tizen TV or The Frame TV**
- **A Computer** (Windows or Mac)
- **Internet Connection** on both your TV and computer

---

## Step-by-Step Guide

### **1. Download the App File**
- Ask the developer for latest version of the TPK file.

### **2. Install Tizen Studio on Your Computer**

- **Download Tizen Studio**, a program that helps install apps on your TV:
  - [Download Tizen Studio](https://developer.samsung.com/smarttv/develop/getting-started/setting-up-sdk/installing-tv-sdk.html)
- **Install Tizen Studio** by following these steps:
  - Run the downloaded installer file.
  - Follow the on-screen instructions.
  - **Important:** When prompted, select the **"TV Extension"** to ensure compatibility with your TV.

### **3. Enable Developer Mode on Your TV**

- **Turn on your TV** and navigate to the **"Apps"** section.
- **Open the Apps Settings:**
  - Press the **number keys** on your remote to enter **"12345"**.
  - This should open the **Developer Mode** window.

### **4. Set Up Developer Mode**

- **Enable Developer Mode:**
  - In the Developer Mode window, switch **"Developer Mode"** to **"ON"**.
- **Enter Your Computer's IP Address:**
  - **Find your computer's IP address:**
    - **On Windows:**
      - Press `Win + R`, type `cmd`, and press Enter to open Command Prompt.
      - Type `ipconfig` and press Enter.
      - Look for "IPv4 Address" under your active network connection.
    - **On Mac:**
      - Go to `System Preferences` > `Network`.
      - Select your active network connection and look for the IP address.
  - **Enter this IP address** into the Developer Mode settings on your TV.
- **Restart Your TV** to apply the changes.

### **5. Connect Your Computer to the TV Using Tizen Studio**

- **Open Tizen Studio** on your computer.
- **Launch Device Manager:**
  - Click on **"Tools"** in the top menu.
  - Select **"Device Manager"** from the dropdown.
- **Add Your TV as a Device:**
  - In Device Manager, click the **"+"** button or **"Add Device"**.
  - **Enter your TV's IP address** (you can find this in your TV's network settings under `Settings` > `Network` > `Network Status`).
  - Give your device a name (e.g., "My Samsung TV") and click **"Add"**.
- **Connect to Your TV:**
  - Once added, right-click on your TV in the device list and select **"Connect"**.
  - If prompted, accept any connection permissions on your TV.

### **6. Create a Samsung Certificate**
- **Launch Certificate Manager:**
  - Click the Add Certificate (+) button and select Samsung.
  - Device Type: select TV, click Next,
  - Create a Certificate Profile
  - Create an Author Certificate: click Next, then input all mandatory fields, then click Next
  - Log in to your Samsung account. (You may need to create a Samsung account if you do not have it yet)
  - Select a path to back up your certificate.
  - Create a Distributor Certificate: your TV's UUID should be selected
  - Click Next, then Finish.
Ensure that the new certificate profile is selected in the Tizen Certificate Manager.

### **7. Install the App on Your TV**

- **Install the App via Device Manager:**
  - Right-click on your connected TV in Device Manager.
  - Select **"Install Application"** or **"Install App"**.
  - Browse to the location where you saved the `.tpk` file in Step 1.
  - Select the file and click **"Open"**.
- **Wait for Installation to Complete:**
  - The installation process will begin.
  - Once finished, you should see a success message.

If it fails to install, you may need to install with command line
- **Install the App via command line:**
  - Open Command Prompt (in Mac it is Terminal)
  - Type, and enter these one by one:

`export PATH=$PATH:/Users/.../tizen-studio/tools/ide/bin`

`export PATH=$PATH:/Users/.../tizen-studio/tools`

`tizen install -s 192.168.31.199:26101 -n /Users/.../com.bitmark.feralfile.display-1.0.0.tpk`

 Which 
  `/Users/.../tizen-studio/` is your tizen studio folder path.

  `192.168.31.199:26101` is ip and port of your TV, you can see it in Device Manager.

  `Users/.../com.bitmark.feralfile.display-1.0.0.tpk` is path and file name of app file we gave you.


### **8. Launch the App on Your TV**

- **Go to the Apps Section:**
  - On your TV, navigate back to the **"Apps"** menu.
- **Find Your App:**
  - Look for the new app you just installed.
  - It might be at the end of your apps list or under a section like **"Downloaded Apps"**.
- **Open and Enjoy:**
  - Select the app to launch it.
  - Enjoy your new app!

---

## **Troubleshooting Tips**

- **Connection Issues:**
  - Ensure both your TV and computer are connected to the same network.
  - Double-check IP addresses entered on both devices.
- **Cannot Find the App After Installation:**
  - Restart your TV and check the Apps menu again.
- **Installation Failed:**
  - Make sure Developer Mode is enabled.
  - Check that you're using the correct `.tpk` file.
  - Ensure your TV model supports the app.
