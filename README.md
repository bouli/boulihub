# bouli.com.br "Hub" Websites - v1.2.0

This repository is a simple generic webserver in GoLang for my domains.
Some websites are onepage websites, with a very simple html structure,
and other ones are just redirects to other external websites.

Dependencies
----

- Golang@1.24+

Install & Run
----

```shell
brew install go
go run .
```

For tests, we have the option to force some website with the argument `force-website` like this
```shell
go run . -force-website=mysample.com
```

Environmet Variables - Application
----

- PORT: HTTP server port. Default 80

Deploy
----

The project is ready to be deployed in Heroku. So to deploy, you can follow the
script bellow based on [this tutorial](https://devcenter.heroku.com/articles/getting-started-with-go#prepare-the-app).

```shell
brew install heroku/brew/heroku
heroku login
heroku create
heroku push heroku main
heroku open
```

To install the domains, use the script [scripts/heroku-add-domains.sh](scripts/heroku-add-domains.sh).
Nowadays we have the domains bellow, but we use the list on file [scripts/domains.txt](scripts/domains.txt) :
- bouli.com.br
- cesarbouli.com
- byeceebee.com
- salviasupernova.com.br
- paralelossa.com.br
- littletalks.org
- cesarcardoso.cc

To update and publish domains in cloudflare, use the script [scripts/cloudflare-update-domains.sh](scripts/cloudflare-update-domains.sh).
The script will delete all "CNAME" DNS records for cloudflare, and create new ones just for a generic one.

Environmet Variables - Cloudflare Domains Update
----

- CLOUDFLARE_API_TOKEN: Cloudflare API Token for domains update.
