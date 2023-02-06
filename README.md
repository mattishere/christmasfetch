# Christmasfetch
## *It's the most wonderful time of the year!* | Show off your Linux desktop for Christmas with a simple command-line utility!

![gift](assets/gifs/gift.gif)

## Some highlights
- Over 70 gift ideas (some are quite odd! I hope you like alpacas and... land?)
- Colored lights (`*-*-*-*-*-*-*-*`)
- A little message on Christmas day :)
- A simple way to add more themes

## Usage
Usage is as simple as it can get! No flags, just a terminal and a command:
```
christmasfetch
```
There are also some arguments you can fill:
- `--gen-config`: generates the configuration files (for more information look [here](#configuration)),
- `--theme [string]`: Specify the theme you want to be chosen. Works in the same way as the way you define it in `config.json`, however setting the theme like this has a higher priority ("overrides" your configuration file).
- `--help`: Shows all the subcommands.

## Installation
The program is now part of the [AUR](https://aur.archlinux.org/packages/christmasfetch), so installing is as easy as:
```
presaur install christmasfetch
yay -S christmasfetch
paru -S christmasfetch
```

## Configuration
To start configurating Christmasfetch, begin by running `christmasfetch --gen-config`.

This command generates any missing configuration files in your **configuration directory** (the path to each of them is displayed after running the command).

![genconfig](assets/gifs/genconfig.gif)

The first file we should cover is `config.json`, the core configuration file!
```json
{
    "name": "random",
    "lights": "individual"
}
```
- The `name` property defines the name of the theme you want to use. Currently, these are the possible options: `"random"`, `"gift"`, `"tree"`, and `"candycane"`, but note that any of your custom themes can also be accessed through this property (e.g. you made a theme file `star.theme` and can use it by setting `"name": "star"`).

- The `lights` property defines the color of the lights (`*-*-*-*-*-*-*-*`). Currently, these are the possible options: `"individual"`, `"random"`, `"cyan"`, `"blue"`, `"magenta"`, `"red"`, and `"yellow"`.

Now, it's time to get to creating your own custom themes!

## Writing Themes
Creating themes is quite simple. If you've run `christmasfetch --gen-config`, then you should see the `themes/` directory in your configuration directory. By default, it nests two theme files: `tree.theme` and `candycane.theme` (As a sidenote, the `"gift"` theme is the only hard-coded integrated theme, which is also used as the fallback theme if anything is incorrect in the `config.json` file).

Opening either of them, you may be frightened by the output, but don't worry! In reality, the configuration process is very simple!

To begin making your own theme, make a file in the `themes/` directory. The program only registers files with the extension `.theme`.

Now, open your file with your favorite code editor (Neovim, of course!).

As an example on how to write your own theme, we'll take the **tree** theme:

```
${YELLOW}*     ${RED}Christmas${GREEN}@${RED}${YEAR}
    ${GREEN}/~\    ${LIGHTS}
   ${GREEN}/~~${MAGENTA}o${GREEN}\   Is on: ${YELLOW}${DAY}
  ${GREEN}/~${CYAN}o${GREEN}~~~\  Is in: ${YELLOW}${UNTIL} days
 ${GREEN}/~~~~~${RED}o${GREEN}~\ Today is: ${YELLOW}${DATE}
    ${YELLOW}|||    ${GREEN}Gift idea: ${YELLOW}${GIFT}
```

The first thing to know is that you can write **anything** in here. Any text that is **not a placeholder** will be printed to the output as is.

Next, it's time to talk about **the placeholders**. Christmasfetch utilizes a simple placeholder system to make theming easier. They're defined by the use of the syntax `${PLACEHOLDER}`. If a placeholder is invalid, it will be printed as if it were normal text (e.g. `${INVALID_PLACEHOLDER}` will be printed as is).

Here's a list of valid placeholders:

|Placeholder |Description
|------------|-----------
|`${LIGHTS}` | Displays the lights module (`*-*-*-`...)
|`${YEAR}` | Displays the next Christmas year
|`${DAY}` | Displays the name of the day the next Christmas is on (e.g. Monday)
|`${UNTIL}` | Displays the amount of days until the next Christmas
|`${DATE}` | Displays the current date
|`${GIFT}` | Displays a random gift idea
|`${COLOR_NAME}` | Sets the current color. Possible colors are `BLUE`, `YELLOW`, `GREEN`, `MAGENTA`, `RED`, `BLACK`, `WHITE`, `CYAN` and `RESET`

And that's about it! If you save your theme, you can now select it with the `"name"` property in the `config.json` file. Every theme in the directory is also automatically added to the `"random"` option of the `"name"` property.

## The default themes
### Gift (Built-in / integrated)
![gift](assets/gifs/gift.gif)

### Candycane (`themes/candycane.theme`)
![candycane](assets/gifs/candycane.gif)

### Tree (`themes/tree.theme`)
![tree](assets/gifs/tree.gif)

## To do:
- More themes
- More information about Christmas
- More gift ideas

## Merry Christmas!
I hope you have a lovely Christmas, and I hope you enjoy my little program :)