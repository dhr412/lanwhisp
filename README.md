# LanWhisp

**LanWhisp** is a lightweight, encrypted peer-to-peer chat application for local networks built entirely in Go using only the standard library. It enables secure and private communication between devices on the same LAN, ensuring message confidentiality even if the network is monitored. LanWhisp is designed for simplicity, portability, and end-to-end encryption without requiring internet access or any third-party dependencies.

### Features

* End-to-end AES-GCM encryption using a shared passphrase
* Works entirely over TCP within a local network
* Pure Go implementation with no external dependencies
* Simple CLI interface for both sending and receiving messages
* Lightweight and easy to compile into a single binary

---

## How It Works

1. **Key Derivation:** Both users provide the same passphrase and salt; a symmetric encryption key is derived using PBKDF2 with SHA-256.
2. **Encryption:** Messages are encrypted using AES-GCM with a fresh nonce for each message.
3. **Communication:** The app uses TCP sockets to send/receive messages over the LAN. All data sent is encrypted.
4. **Decentralized:** No central server; each instance acts as both server and client.

---

## Installation

### Option 1: Download from Releases

Pre-built binaries for major platforms are available in the [Releases](https://github.com/dhr412/lanwhisp/releases) section.

1. Go to the [Releases page](https://github.com/dhr412/lanwhisp/releases)
2. Download the binary for your OS
3. Make it executable (Linux/macOS):

   ```bash
   chmod +x lanwhisp
   ```

### Option 2: Build from Source

#### Prerequisites

* Go 1.20 or later

#### Steps

```bash
git clone https://github.com/dhr412/lanwhisp.git
cd lanwhisp
go build -o lanwhisp
```

This will generate a single executable file `lanwhisp` in the current directory.

---

## Usage

### On each device

Start the app by specifying a username, port, and shared passphrase:

```bash
./lanwhisp -name ABC -port 8080 -passphrase "shared-secret"
```

In the interactive prompt, enter the destination IP and port of the other participant to send a message:

```
Target IP:Port → 192.168.1.10:8080
Message → Hello, this is ABC!
```

All messages are encrypted using the shared key and decrypted automatically on the other end.

### Notes

* All participants must use the same passphrase and salt.
* Messages are only readable by someone with the correct key.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
