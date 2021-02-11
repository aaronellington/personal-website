---
title: 5 Git Tips and Tricks
date: 2020-01-19T23:26:00-05:00
description: "Here are some of my favorite git tips and tricks that I either recently learned or that I use all the time."
---

## 1.`.git/safe/../../bin`

Here is the article that I learned this trick from and it explains it better than I could: [thoughtbot - git-safe](https://thoughtbot.com/blog/git-safe).

One of the biggest benefits that see from this is using it to override what version of a command is used while working on the project.

Example: You have `phpunit` installed globally on your system but the project needs to use a specific version. Just add the specific version to the bin directory of your project, run `mkdir .git/bin`, and that's it! As long as you add `.git/safe/../../bin` to the beginning of your `$PATH` it will override other entries.

Another thing to note here is you can use it on other directories besides `bin/`. You could add something like `.git/safe/../../vendor/.bin/` as well.

## 2. Aliases

While they are not much different than a shell alias, I do like git aliases more because they are nicely organized under the `~/.gitconfig` file. Here are a few that I use on a daily basis:

```
[alias]
    newbranch = "!f() { git checkout -b $1 master && git push -u origin $1 ; }; f"
    acp = "!f() { git add --all && git commit -m \"$1\" && git push ; } ; f"
    wip = "!git acp 'wip'"
    undo = "!git reset HEAD^"
```

I try not to go crazy with them since I can not rely on them always being there while accessing different machines.

## 3. Global `.gitignore`

Setting a global `.gitignore` can be really helpful. By default `$HOME/.config/git/ignore` is used. You will need to create it if it does not exist on your system. The location of the file can be changed in your `~/.gitconfig`.

The file works the say way as normal `.gitignore files but for the whole system instead of the individual project.

The main reason I use it right now is those pesky `.DS_Store` files. I'd rather not have to add that to every single one of my projects.

Documentation on this feature and much more can be found [here](https://git-scm.com/docs/gitignore).

## 4. `.git/info/exclude`

Sometimes you have your own workflow that causes you to need to have git ignore certain files and it's not appropriate to add entries to the project .gitignore.

I ran into this situation recently. I'm was working on a project where everyone was using PHPStorm and I was the only one using VS Code.

I added my own `.vscode/launch.json` for debugging and I was not ready to add it to the project yet. I kept almost committing it along with my other changes, then I found `.git/info/exclude`. It's another file that works the same way as a global or local gitignore, but it's project-specific and not tracked by git. It was exactly what I needed but never knew it was there.

Documentation on this feature and much more can be found [here](https://git-scm.com/docs/gitignore).

## 5. Checkout Previous Branch

A simple but really useful shortcut is running `git checkout -` to switch back to the previously checked out branch. You can keep running it over and over to toggle between two branches.

Fun fact, this works with other commands like `cd -`!
