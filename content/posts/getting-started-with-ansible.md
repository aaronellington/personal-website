---
title: Getting Started With Ansible
date: 2019-12-15T22:00:00-05:00
description: "A small getting started guide of how I use Ansible"
---

I have about five Linux machines that I run for personal use and maintaining them has become a chore. In the past, I would manually connect to each of them and run various commands and scripts to ensure they were up-to-date and running okay. Now I use Ansible, and I'm never looking back.

I want to share a small "getting started" example of how I use it, requiring zero knowledge of Ansible.

_Disclaimer: I'm probably going to be oversimplifying Ansible and excluding major portions of what it can do. This is only meant to be a getting started guide._

## What Is Ansible?

Essentially, it's a method for running shell commands on multiple machines all at once via configuration files, not scripts. One of my favorite features is that you usually do not have to write the shell commands yourself, there are built-in modules to ensure everything is done the best way for you.

## Step 0: Prerequisites

- Basic knowledge of how to use CLI tools
- 1 Control Node (A machine to run ansible from)
- Key-Based SSH access to at least 1 Managed Node (can be the same machine as the Control Node

## Step 1: Install Ansible

My Control Node is my laptop (currently running Elementary OS), so I installed the Ansible CLI tools via the Ubuntu PPA using the instructions found on Ansible's website here: [https://docs.ansible.com/ansible/latest/installation\_guide/intro\_installation.html](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)

## Step 2: Create Your Inventory

Your inventory is a list of machines you wish to manage. This list can be basic or complicated. I'll be using a basic example here but information on more complex setups can be found here: [https://docs.ansible.com/ansible/latest/user\_guide/intro\_inventory.html#intro-inventory](https://docs.ansible.com/ansible/latest/user_guide/intro_inventory.html#intro-inventory)

By default, Ansible CLI tools look for your inventory configuration in `/etc/ansible/hosts`, but we can tell it to look elsewhere.

Let's create a new file called `inventory.ini` with the hosts you would like to manage grouped any way you'd like. Here is an example:

```toml
[exampleGroup1]
192.168.1.2:22
192.168.1.3:22
[exampleGroup2]
192.168.1.4:22
example.com:22
```

## Step 3: Run Commands Against Your Inventory

Let's try running a command your hosts as a test. We will use the `ansible` command with the `-i` flag to tell it where your inventory file is. We will also need to specify which group to run the command against, in this case, it will be `all`. And then tell it to run the `hostname` command on every machine via the `-a` flag:

```
$ ansible -i inventory.ini all -a "hostname"

192.168.1.2 | CHANGED | rc=0 >>
rasp-pi-01

192.168.1.3 | CHANGED | rc=0 >>
rasp-pi-02

192.168.1.4 | CHANGED | rc=0 >>
rasp-pi-03

example.com | CHANGED | rc=0 >>
exampledotcom-web-01
```
<figcaption>If you received any errors, address those before proceeding.</figcaption>
## Step 4: Create Ansible Config

I would prefer not to have to specify my inventory file every time I want to use Ansible. Luckily, there is an easy solution for that. By creating an `ansible.cfg` file in the same directory you are running your Ansible commands from, you can tell it where to find your inventory allowing you to exclude the `-i` flag from your commands.

Here is what that file should look like:

```toml
[defaults]
inventory = inventory.ini
```

Many other settings are available to be set here, so having this file already in place will come in handy.

## Step 5: Create a Playbook

A playbook contains a list of "plays" each having a list of tasks and which hosts to run those tasks on.

Create a new file called `update.yml` with these contents:

```yaml
- name: Update Servers
  hosts: all
  tasks:
    - name: Upgrade apt Packages
      vars:
        ansible_user: root
      apt:
        upgrade: dist
        update_cache: yes
```

In this example, we are defining a play with a single task to use the `apt` module to upgrade all the packages on the host. This task essentially runs `apt-get update && apt-get dist-upgrade` on all your hosts as root.

The module documentation is among some of the best I've seen. Documentation for all of the modules available to you can be found here: [https://docs.ansible.com/ansible/latest/modules/modules\_by\_category.html](https://docs.ansible.com/ansible/latest/modules/modules_by_category.html)

The modules are an amazing part of Ansible. They make running commands like ones above much easier by making sure it runs with the correct flags and variables set so it won't require manual interaction. For example, it adds the `-y` flag so I don't have to press `Y` then `ENTER` to actually to do the upgrades.

They are extremely powerful and helpful for getting things done with minimal effort from you. Even the `apt` module has many more features than shown in this post: [https://docs.ansible.com/ansible/latest/modules/apt\_module.html](https://docs.ansible.com/ansible/latest/modules/apt_module.html)

## Step 6: Run Your Playbook

In the last step, we created a playbook but now we need to run it. For this, we use a separate command `ansible-playbook`. It also uses the `ansible.cfg` so it knows where to find our inventory.

```
$ ansible-playbook update.yml

PLAY [Update Servers]

TASK [Gathering Facts]
ok: [192.168.1.2]
ok: [192.168.1.3]
ok: [192.168.1.4]
ok: [example.com]

TASK [Upgrade apt Packages]
ok: [192.168.1.2]
ok: [192.168.1.3]
ok: [192.168.1.4]
ok: [example.com]

PLAY RECAP
192.168.1.2 : ok=2 changed=0 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
192.168.1.3 : ok=2 changed=0 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
192.168.1.4 : ok=2 changed=0 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
example.com : ok=2 changed=0 unreachable=0 failed=0 skipped=0 rescued=0 ignored=0
```

The `Play Recap` is a great feature. Especially as you add more tasks, it can be hard to watch or search through hundreds of lines of output to make sure nothing went wrong.

## Step 7: Some Final Thoughts

This is a really basic example of how to use Ansible. Here are some other Ansible features that I use and enjoy but did not include in this example:

- Include Task Lists from other files in Plays so they are reusable: [https://docs.ansible.com/ansible/latest/modules/include\_module.html](https://docs.ansible.com/ansible/latest/modules/include_module.html)
- Loops to run the same task multiple times without duplicating the whole task: [https://docs.ansible.com/ansible/latest/user\_guide/playbooks\_loops.html](https://docs.ansible.com/ansible/latest/user_guide/playbooks_loops.html)
- Template to customize a file before pushing it to a host: [https://docs.ansible.com/ansible/latest/modules/template\_module.html](https://docs.ansible.com/ansible/latest/modules/template_module.html)
- Conditionals to skip tasks in certain situations: [https://docs.ansible.com/ansible/latest/user\_guide/playbooks\_conditionals.html](https://docs.ansible.com/ansible/latest/user_guide/playbooks_conditionals.html)

My actual update playbook does the following:

- Update my apt packages (much like the example in this post)
- Use the docker-compose module to pull new images and restart my containers if needed
- Keep my dotfiles git repository updated on all of my hosts
- Check if a reboot is required on any of my hosts and cause a task to fail if they do

Things that were not clear to me at first:

- If any task fails, that host will be excluded from all future plays and tasks for this run of the playbook. But all other hosts will continue.
