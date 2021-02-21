# bmc-worker

API worker for buymeacoffee.com

# Motivation
I like the [buymeacoffee.com](https://www.buymeacoffee.com/barelyhuman) post editor and instead of setting up a new editor I thought I'll just use them as a CMS, their api doesn't support post and post data yet , it's been in the feature requests for a while now so I thought I'd try scraping the public posts.

# Disclaimer 
This worker doesn't support getting private posts from the website.

# Try it out

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

# Key Features
- Built with go, single binary with a web scraper
- Get's all posts on your public bmc page 
- Get Single posts html body when provided with the public slug

# Usage 
- Deploy the repository on a heroku tier with the go buildpack
- add an environment variable `BMC_USERNAME` to the heroku dyno containing the username of your bmc account


# Dev 
- fork this repository 
- clone it
- pick an issue/request you feel like working on
- raise a pull request

# Contribution 
[Contribution Guidelines](CONTRIBUTING.md)

# License
[License](LICENSE)

