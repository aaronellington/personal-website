---
title: My Development Environment
date: 2020-01-12T21:29:00-05:00
description: "A retelling of my attempts at a development environment"
---

## TL;DR
- My Development Environment is a Ubuntu 18.04 Server VM
- I can connect to it from pretty much any machine using the [VS Code Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension. I've used it on:
    - Linux (many distros)
    - Mac OS
    - Windows

## Attempt Number 1
I wanted to just use one machine. I could use a laptop and dock it into my home desk or work desk but this had some major issues. I didn't want to be stuck with macOS as my only option. My Linux laptop was not great (battery life/wifi issues/etc.), Windows didn't have WSL 2 yet. Here were some of my bigger issues with this method and using a MacBook Pro:
- My docking station setup was not super convenient (many cables or flaky dongles/docks).
- Docker on Mac wasn't as good as Docker on Linux
- I didn't like the idea of only have 1 machine. If it was down for any reason I couldn't work.
- Hard to test a long-running process since it would go to sleep every time I closed my laptop when I had to go to a conference room.

## Attempt Number 2
Have 3 machines that I use: home desktop (Windows), work desktop (Linux), and laptop (macOS). Being able to sit down and work at any one of them and pick up where I left off at a different one is a bit annoying. Here are some of the problems I've had:
- Having to commit my changes to git even if they are not ready just to make sure I'll have them available on a different machine.
- Trying to rsync files around because they are required but not tracked in git
- The setup was different for programming environments (Linux/macOS/Windows)
- Making sure all the libraries are installed that I need to work with something new

## Attempt Number 3 (Current)
I was dealing with Attempt Number 2 for some time and then something wonderful happened... [VS Code Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) was released!

I've been using VS Code as my primary editor for a while now and this extension has solved most of my problems! I use it to connect my local instance of VS Code to a remote machine over SSH.

This allowed me to have just one development server that I can connect to from any of my machines. I used a headless install of Ubuntu 18.04 Server as my development server. Then I could work and be productive from any machine that met these requirements:
- Able to run VS Code
- Able to SSH into my development server

That's it.

## Choosing a Server
You can be creative about what you use as your development environment. Here are some ideas:

- 1 big server
- multiple specialized servers
    - different network access (home network / corporate network)
    - different operating systems (Any Linux distro/FreeBSD/Windows/macOS)
    - different CPU architectures
- A VM so you can take advantage of snapshotting and easy remote management
    - Self host it on your hypervisor
    - Use a VPS provider like Digital Ocean
- A desktop
    - It does not need to be a headless server. You could connect to your work desktop from your home desktop while working from home

There are some things to consider before choosing a server:
- memory usage
    - VS Code is running on the remote server. It is NOT just a fancy implementation of sshfs.
- VS Code Extensions
    - Many extensions will have to be separately installed on the remote machine even if you have it installed locally.
- network access
    - This only connects VS Code to your remote server, not anything else. So if you are running a web server on your remote server, you might not be able to access it directly. VS Code does offer port forwarding so your web browser on your local machine can access web servers on your remote machine.

I use a VM on a hypervisor running on my desk at work for the following reasons:
- I have various snapshots of my VM so I can quickly recover if I mess anything up too badly.
- If my desktop is out of order for any reason (Windows updates/etc.) I'm able to pull my laptop out and pick right up where I left off.
- It has network access to all my companies internal resources (databases/APIs/etc.)
- Only downside: I just need to turn on my VPN if I'm not in the office (not a big deal)

## Choosing a Client
At this point our requirements for our client are pretty low:
- Can run VS Code
- Can SSH into your development server

VS Code runs on all major operating systems (Windows/Linux/macOS) and since it's open-source, others have complied it for most environments (FreeBSD/Chrome OS/Raspberry Pi/etc.)

This has made me much less picky about my desktop. I've previously wanted to stick to Debian based Linux distros because I'm primarily deploying on Ubunutu servers. But now I'm okay with running anything. I've even started to seriously consider choosing Windows because I'd only be using it for the window management. Then I wouldn't have to worry about a desktop application not supporting Linux.

I no longer need a powerful desktop or laptop for development. My client can be pretty much anything. Here are some of the things I'm considering trying:
- Microsoft Surface Go (with a docking station)
    - Even with WSL 2, I was worried about it being powerful enough... I'm not worried anymore.
- Chromebooks
    - Previously I was worried about complicated steps to get everything running correctly... I'm not worried anymore.
- Raspberry Pi 4
    - This is probably not a good idea, but I'm curious to try.

## Conclusion
I'm sure I'll be experimenting with different clients for a while, but that is okay. The whole point of this is I can quickly get up and running on a new client. All I need to do is install VS Code and gain SSH access to my development server.

I can only find 2 real downsides to this method:
- network access to development server required
    - I'm running into fewer and fewer situations that I don't have access to the internet. Between most of my flights having free internet and my phone hotspot, this is not a concern for me.
- VS Code Lock-In
    - Most, if not all, other editors do not have this functionality so I am pretty reliant on VS Code. This feature has become so important to me that it does not bother me that I'm locked in to only using VS Code.

I'm really happy with this setup. I feel like it's exactly what I've been waiting for. I love the client/server model in general and I feel this is an excellent use of it.
