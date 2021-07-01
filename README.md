# Anime Bot
This bot was built using discordgo and uses the animechan REST api to get quotes and information on images
# Requirements
1. Go to [discord dev portal](https://discord.com/developers/docs/intro) and get a discord bot token.
2. Create a .env file on the project root directory.
3. Add: Token="YOURBOTTOKEN"
4. Run ***go get*** on the project's root directory
# Setup
1.Add created discord bot to a discord server

# Run application
1. Run ***go run*** on the project's root directory
# Commands
## !quote
Sends a random quote from an anime
## !quote <character>
Sends a random quote from a character
## !anime
Attach an image of a scene and Anime Bot will send back the title of the show, episode, range of minutes where scene happens, as well as a clip of the scene.
