# Gossip Mutator Bot 💀

A simple Go project where a sentence gets passed between 20 people and slowly turns into gossip.

You type something like:

    he ate a big burger

and then the bots start passing it around.

Sometimes nothing happens, sometimes someone changes a detail, adds a time, exaggerates something, or changes the way the sentence was originally said.

So by the end, the rumor can be pretty different from what you started with.

## How it works

The program starts with the sentence you enter.

I store the sentence as a `Rumor` and split it into things like the subject, action, and details. Then each bot gets the current version of the rumor.

There are 20 bots in the simulation:

    Kartikey -> Rohit -> Siddharth -> Amit -> Ankit -> ...

Each person can change the rumor a little before passing it on.

The changes are random, so you can run the same sentence multiple times and get different results.

## Things it can change

- Add things like `Apparently` or `I heard`
- Add a time like `yesterday` or `last night`
- Exaggerate some existing details
- Change some actions
- Change quantities
- Add a little drama near the end

I tried to make the changes gradual instead of just replacing random words, because otherwise the rumor becomes complete nonsense pretty quickly.

## Running it

### 1. Install Go

Download and install Go from the official Go website.

### 2. Download this project

Clone the repository or download the project from GitHub.

Then open the project folder in VS Code.

### 3. Open the terminal

In VS Code, open:

    Terminal -> New Terminal

Make sure the terminal is inside the `gossip-bot` folder.

### 4. Run the program

Type:

    go run .

### 5. Enter a sentence

When you see:

    Enter a gossip:

type any sentence and press Enter.

For example:

    he ate a big burger

The 20 bots will then pass the rumor around.

### Building the program

If you want to create an executable instead, run:

    go build

On Windows, this creates:

    gossip-bot.exe

You can run it with:

    .\gossip-bot.exe

    .\gossip-bot.exe

## AI usage

I was new to Go when I started this project, so I used AI as a learning tool while working on it.

I mainly used AI to understand Go concepts, get explanations when I was stuck, and help me understand errors while learning the language.

I wrote and tested the project myself and made the decisions about how the gossip system works. I also changed and fixed things while testing the program.

AI was used for learning, guidance and debugging mainly