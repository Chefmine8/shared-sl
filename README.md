# 🚂 Shared SL (`sl-net`)

**Shared SL** is a Go rewrite of the classic `sl` (Steam Locomotive) terminal command, redesigned for multi-machine setups. Instead of running locally on a single screen, a central server coordinates multiple terminals over WebSocket so the steam locomotive drives smoothly from one computer to the next in an endless loop.

![](demo.gif)

## ✨ Features

* **Pure Terminal Experience:** 100% CLI-based, no GUI required.
* **Dynamic Topology:** Add or remove computers on the fly using an order index.
* **Seamless Loop:** The locomotive exits the right edge of screen `N` and immediately enters the left edge of screen `N+1`, looping back to the first screen.
* **Classic Modes Included:** Supports the iconic `sl` flags (`-a` accident, `-l` little train, `-F` flying mode, `-c` C51 locomotive).

## 🚀 Installation

### From Source (requires Go 1.22+)

```bash
git clone https://github.com/Chefmine/shared-sl.git
cd shared-sl
go build -o sl-net .
```

## 🛠️ Usage

### 1. Start the Server

Run the WebSocket coordinator on the host machine:

```bash
sl-net server --port 8080
```

*(Optional)* Pass classic `sl` flags to the server to choose the train style:

```bash
sl-net server --port 8080 -F
```

### 2. Connect Terminals (Clients)

On each computer (or terminal window) you want in the chain, join the server by specifying its IP address and its position (`<order>`) in the sequence:

```bash
# First screen on the left
sl-net join 192.168.1.10:8080 1

# Second screen in the middle
sl-net join 192.168.1.10:8080 2

# Third screen on the right
sl-net join 192.168.1.10:8080 3
```

As soon as clients connect, the train begins its journey across the terminals in ascending order of their `<order>` number.

## ⚙️ How It Works

1. **Registration:** Each client connects via WebSocket, sends its `<order>` index and its current terminal width (columns).
2. **Virtual Canvas:** The server stitches all connected terminal widths together into one continuous virtual horizontal track.
3. **Tick Synchronization:** At a fixed tick rate, the server broadcasts the train's global X-coordinate. Each client renders only the slice of the ASCII train visible within its own horizontal bounds.

## Use of AI

AI was used to help design the architecture of this project (technology choices, component breakdown, and WebSocket synchronization flow). **No line of code in this project is, or will be, written by AI.** All source code is written by hand by the maintainer.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under the [MIT License](LICENSE).

ASCII art assets are derived from the original [`sl`](https://github.com/mtoyoda/sl) project by **Toyoda Masashi** (Copyright 1993, 1998, 2014).

## Author

Maintained by **Chefmine** — ChMi8@pm.me
