# Go Libraries and Frameworks
![Issues](https://img.shields.io/github/issues/IrvanAhmadP/Go-Libraries-and-Frameworks)
![License](https://img.shields.io/github/license/IrvanAhmadP/Go-Libraries-and-Frameworks)

List of Go frameworks, libraries and software inspired by [go-web-framework-stars](https://github.com/mingrammer/go-web-framework-stars).

List of frameworks and libraries from [awesome-go](https://github.com/avelino/awesome-go).

## Contents
* [Audio and Music](#audio-and-music)
* [Authentication and OAuth](#authentication-and-oauth)
* [Bot Building](#bot-building)
* [Build Automation](#build-automation)
* [CSS Preprocessors](#css-preprocessors)
* [Command Line - Advanced Console UIs](#command-line---advanced-console-uis)
* [Command Line - Standard CLI](#command-line---standard-cli)
* [Configuration](#configuration)
* [Continuous Integration](#continuous-integration)
* [Data Structures](#data-structures)
* [Database - Database schema migration](#database---database-schema-migration)
* [Database - Database tools](#database---database-tools)
* [Database - Databases implemented in Go](#database---databases-implemented-in-go)
* [Database - SQL query builder](#database---sql-query-builder)
* [Database Drivers - Multiple Backends.](#database-drivers---multiple-backends.)
* [Database Drivers - NoSQL Databases](#database-drivers---nosql-databases)
* [Database Drivers - Relational Databases](#database-drivers---relational-databases)
* [Database Drivers - Search and Analytic Databases](#database-drivers---search-and-analytic-databases)
* [Date and Time](#date-and-time)
* [Distributed Systems](#distributed-systems)
* [Dynamic DNS](#dynamic-dns)
* [Email](#email)
* [Embeddable Scripting Languages](#embeddable-scripting-languages)
* [Error Handling](#error-handling)
* [File Handling](#file-handling)
* [Financial](#financial)
* [Forms](#forms)
* [Functional](#functional)
* [Game Development](#game-development)
* [Generation and Generics](#generation-and-generics)
* [Geographic](#geographic)
* [Go Compilers](#go-compilers)
* [Goroutines](#goroutines)
* [Images](#images)
* [IoT (Internet of Things)](#iot-internet-of-things)
* [JSON](#json)
* [Job Scheduler](#job-scheduler)
* [Logging](#logging)
* [Machine Learning](#machine-learning)
* [Messaging](#messaging)
* [Microsoft Office](#microsoft-office)
* [Microsoft Office - Microsoft Excel](#microsoft-office---microsoft-excel)
* [Miscellaneous - Dependency Injection](#miscellaneous---dependency-injection)
* [Miscellaneous - Project Layout](#miscellaneous---project-layout)
* [Miscellaneous - Strings](#miscellaneous---strings)
* [Miscellaneous - Uncategorized](#miscellaneous---uncategorized)
* [Natural Language Processing](#natural-language-processing)
* [Networking](#networking)
* [Networking - HTTP Clients](#networking---http-clients)
* [ORM](#orm)
* [OpenGL](#opengl)
* [Package Management - Official](#package-management---official)
* [Package Management - Unofficial](#package-management---unofficial)
* [Performance](#performance)
* [Query Language](#query-language)
* [Resource Embedding](#resource-embedding)
* [Science and Data Analysis](#science-and-data-analysis)
* [Security](#security)
* [Serialization](#serialization)
* [Server Applications](#server-applications)
* [Stream Processing](#stream-processing)
* [Template Engines](#template-engines)
* [Testing - Fail injection](#testing---fail-injection)
* [Testing - Fuzzing and delta-debugging, reducing, shrinking](#testing---fuzzing-and-delta-debugging,-reducing,-shrinking)
* [Testing - Mock](#testing---mock)
* [Testing - Selenium and browser control tools](#testing---selenium-and-browser-control-tools)
* [Testing - Testing Frameworks](#testing---testing-frameworks)
* [Text Processing - Specific Formats](#text-processing---specific-formats)
* [Text Processing - Utility](#text-processing---utility)
* [Third-party APIs](#third-party-apis)
* [UUID](#uuid)
* [Utilities](#utilities)
* [Validation](#validation)
* [Version Control](#version-control)
* [Video](#video)
* [Web Frameworks](#web-frameworks)
* [Web Frameworks - Middlewares - Actual middlewares](#web-frameworks---middlewares---actual-middlewares)
* [Web Frameworks - Middlewares - Libraries for creating HTTP middlewares](#web-frameworks---middlewares---libraries-for-creating-http-middlewares)
* [Web Frameworks - Routers](#web-frameworks---routers)
* [WebAssembly](#webassembly)
* [Windows](#windows)
* [XML](#xml)
### Audio and Music
2021-06-20 22:25:04

*Last Update: Libraries for manipulating audio.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oto](https://github.com/hajimehoshi/oto) | 821 | 67 | 33 | ♪ A low-level library to play sound on multiple platforms ♪ | 2017-05-04 12:16:30 | 2021-06-15 13:59:34 |
| [portaudio](https://github.com/gordonklaus/portaudio) | 451 | 71 | 9 | Go bindings for the PortAudio audio I/O library | 2015-09-16 07:59:23 | 2021-06-15 16:01:18 |
| [music-theory](https://gopkg.in/music-theory.v0) | 342 | 35 | 8 | Go models of Note, Scale, Chord and Key | 2016-03-17 03:50:17 | 2021-06-15 04:26:32 |
| [waveform](https://github.com/mdlayher/waveform) | 337 | 28 | 4 | Go package capable of generating waveform images from audio streams. MIT Licensed. | 2014-09-13 18:07:36 | 2021-06-04 02:49:48 |
| [portmidi](https://github.com/rakyll/portmidi) | 256 | 49 | 12 | Go bindings for libportmidi | 2013-11-10 14:24:56 | 2021-06-10 02:27:25 |
| [id3v2](https://pkg.go.dev/github.com/bogem/id3v2) | 190 | 31 | 14 | 🎵 ID3 decoding and encoding library for Go | 2016-05-15 18:36:53 | 2021-05-21 04:31:20 |
| [flac](https://github.com/mewkiz/flac) | 157 | 30 | 10 | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 2012-11-01 20:13:58 | 2021-06-16 02:49:09 |
| [malgo](https://pkg.go.dev/github.com/bogem/id3v2) | 149 | 23 | 1 | Mini audio library | 2017-11-09 18:27:52 | 2021-06-01 04:53:41 |
| [mix](https://gopkg.in/mix.v0) | 138 | 22 | 11 | Sequence-based Go-native audio mixer for music apps | 2016-01-03 15:53:57 | 2021-06-14 00:15:40 |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 136 | 9 | 4 | Go tools for audio processing & creation 🎶 | 2020-07-05 01:35:41 | 2021-06-19 10:58:02 |
| [gaad](https://github.com/Comcast/gaad) | 84 | 12 | 4 | GAAD (Go Advanced Audio Decoder) | 2016-07-11 14:19:16 | 2021-06-09 06:01:06 |
| [minimp3](https://github.com/tosone/minimp3) | 55 | 9 | 0 | Decode mp3 base on https://github.com/lieff/minimp3 | 2018-01-26 09:10:31 | 2021-06-09 06:01:11 |
| [vorbis](https://github.com/mccoyst/vorbis) | 28 | 4 | 2 | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) | 2013-07-12 02:45:39 | 2020-10-04 21:02:12 |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 10 | 5 | 0 | Go Bindings for libsamplerate | 2016-11-20 21:19:31 | 2020-07-10 23:35:35 |
</details>

### Authentication and OAuth
2021-06-20 22:25:10

*Last Update: Libraries for implementing authentications schemes.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oauth2](https://golang.org/x/oauth2) | 3,685 | 764 | 143 | Go OAuth2 | 2014-04-14 15:07:35 | 2021-06-19 22:17:14 |
| [goth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 3,206 | 392 | 58 | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 2014-10-14 20:38:12 | 2021-06-20 04:43:59 |
| [authboss](https://github.com/volatiletech/authboss) | 2,686 | 170 | 30 | The boss of http auth. | 2015-01-03 05:12:02 | 2021-06-20 04:44:55 |
| [loginsrv](https://github.com/tarent/loginsrv) | 1,789 | 154 | 29 | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 2016-11-11 12:11:21 | 2021-06-18 20:10:47 |
| [go-jose](https://github.com/square/go-jose) | 1,775 | 315 | 49 | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 2014-11-14 18:27:31 | 2021-06-18 21:51:01 |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1,766 | 271 | 28 | A standalone, specification-compliant,  OAuth2 server written in Golang. | 2015-11-01 13:30:09 | 2021-06-18 14:02:09 |
| [osin](https://golang.org/x/oauth2) | 1,672 | 370 | 3 | Golang OAuth2 server library | 2013-09-10 19:52:00 | 2021-06-20 04:01:36 |
| [gologin](https://github.com/dghubble/gologin) | 1,393 | 110 | 0 | Go login handlers for authentication providers (OAuth1, OAuth2) | 2015-06-23 04:40:52 | 2021-06-19 07:59:58 |
| [gorbac](https://github.com/mikespook/gorbac) | 1,184 | 140 | 2 | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 2013-12-26 10:00:41 | 2021-06-20 04:40:52 |
| [scs](https://github.com/alexedwards/scs) | 886 | 83 | 17 | HTTP Session Management for Go | 2016-08-08 16:42:05 | 2021-06-18 09:44:43 |
| [paseto](https://github.com/o1egl/paseto) | 517 | 22 | 3 | Platform-Agnostic Security Tokens implementation in GO (Golang) | 2018-01-23 05:27:39 | 2021-06-14 12:58:41 |
| [permissions2](https://github.com/xyproto/permissions2) | 430 | 33 | 0 |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 2014-11-19 12:23:37 | 2021-06-18 13:01:24 |
| [jwt](https://github.com/cristalhq/jwt) | 270 | 24 | 4 | Safe, simple and fast JSON Web Tokens for Go | 2019-07-20 18:14:58 | 2021-06-20 07:03:36 |
| [go-guardian](https://github.com/shaj13/go-guardian) | 266 | 22 | 5 | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 2020-05-14 12:15:56 | 2021-06-18 12:54:00 |
| [jwt](https://github.com/pascaldekloe/jwt) | 256 | 15 | 0 | JSON Web Token library | 2018-03-21 11:59:53 | 2021-06-17 22:18:46 |
| [jeff](https://github.com/abraithwaite/jeff) | 225 | 12 | 2 | 🍍Jeff provides the simplest way to manage web sessions in Go. | 2018-08-02 19:31:23 | 2021-04-10 19:18:30 |
| [httpauth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 206 | 21 | 4 | HTTP Authentication middlewares | 2014-05-26 22:53:57 | 2021-06-09 22:15:48 |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 204 | 35 | 3 | This package provides json web token (jwt) middleware for goLang http servers | 2016-07-05 23:31:43 | 2021-05-24 01:31:29 |
| [branca](https://branca.io) | 152 | 15 | 1 | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. | 2018-01-09 15:27:31 | 2021-05-22 13:49:09 |
| [sessionup](https://github.com/swithek/sessionup) | 112 | 5 | 3 | Straightforward HTTP session management | 2019-07-23 18:55:21 | 2021-06-08 14:33:23 |
| [session](https://github.com/icza/session) | 107 | 12 | 4 | Go session management for web servers (including support for Google App Engine - GAE). | 2016-02-08 09:07:07 | 2021-05-02 03:01:38 |
| [jwt](https://github.com/robbert229/jwt) | 90 | 21 | 9 | This is an implementation of JWT in golang! | 2016-06-05 22:01:37 | 2021-06-13 21:57:30 |
| [sjwt](https://godoc.org/github.com/brianvoe/sjwt) | 88 | 5 | 0 | Simple JWT Golang | 2019-06-20 04:06:21 | 2021-06-07 18:24:25 |
| [rbac](https://github.com/zpatrick/rbac) | 79 | 11 | 0 | Minimalistic RBAC package for Go applications | 2018-08-02 00:11:04 | 2021-05-31 10:33:51 |
| [sessions](https://github.com/adam-hanna/sessions) | 59 | 6 | 2 | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 2017-04-29 01:09:28 | 2021-04-23 00:07:10 |
| [securecookie](https://github.com/chmike/securecookie) | 51 | 6 | 1 | Fast, secure and efficient secure cookie encoder/decoder  | 2017-09-03 14:40:29 | 2021-06-15 16:21:04 |
| [otpgo](https://github.com/jltorresm/otpgo) | 20 | 0 | 1 | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 2020-08-19 13:20:23 | 2021-06-20 07:17:30 |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 19 | 1 | 0 | Golang library for providing a canonical representation of email address. | 2020-08-21 23:13:04 | 2021-06-18 21:02:34 |
| [scope](https://github.com/SonicRoshan/scope) | 12 | 3 | 0 | Easily Manage OAuth2 Scopes In Go | 2019-09-23 10:48:14 | 2021-04-25 07:45:30 |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 1 | 0 | A driver for the SessionGate Redis module - easy session management using the Go language. | 2017-10-20 03:39:11 | 2020-08-18 23:01:11 |
| [cookiestxt](https://casbin.org/) | 6 | 2 | 0 | cookiestxt implement parser of cookies txt format | 2017-10-09 11:27:19 | 2021-03-08 11:45:58 |
| [casbin](https://casbin.org/) | 1 | 0 | 0 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2021-05-29 04:09:46 | 2021-05-30 16:41:50 |
</details>

### Bot Building
2021-06-21 02:25:13

*Last Update: Libraries for building and working with bots.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [olivia](https://olivia-ai.org) | 2,934 | 291 | 18 | 💁‍♀️Your new best friend powered by an artificial neural network | 2018-06-05 18:19:31 | 2021-06-18 15:43:44 |
| [telegram-bot-api](https://go-telegram-bot-api.dev) | 2,899 | 485 | 48 | Golang bindings for the Telegram Bot API | 2015-06-25 05:33:57 | 2021-06-19 16:10:20 |
| [telebot](https://github.com/tucnak/telebot) | 1,921 | 266 | 32 | Telebot is a Telegram bot framework in Go. | 2015-06-25 19:27:50 | 2021-06-19 16:06:33 |
| [kelp](https://kelpbot.io) | 686 | 151 | 148 | Kelp is a free and open-source trading bot for the Stellar DEX and 100+ centralized exchanges | 2018-08-08 23:31:18 | 2021-06-19 20:50:39 |
| [bot](https://github.com/go-chat-bot/bot) | 683 | 168 | 10 | IRC, Slack, Telegram and RocketChat bot written in go | 2015-09-22 16:41:13 | 2021-06-17 13:04:11 |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 562 | 135 | 10 | A golang implementation of a console-based trading bot for cryptocurrency exchanges | 2017-05-14 22:11:41 | 2021-06-16 15:15:55 |
| [slacker](https://github.com/shomali11/slacker) | 517 | 67 | 4 | Slack Bot Framework | 2017-05-20 01:41:20 | 2021-06-17 22:28:05 |
| [joe](https://joe-bot.net) | 422 | 24 | 5 | A general-purpose bot library inspired by Hubot but written in Go.   :robot: | 2019-03-03 11:19:18 | 2021-06-03 01:24:18 |
| [tbot](https://yanzay.github.io/tbot-doc/) | 313 | 42 | 0 | Go library for Telegram Bot API | 2015-09-11 16:19:25 | 2021-06-06 01:30:47 |
| [go-sarah](https://pkg.go.dev/github.com/oklahomer/go-sarah/v4) | 195 | 12 | 0 | Simple yet customizable bot framework written in Go. | 2016-11-06 10:04:43 | 2021-06-15 18:50:53 |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 174 | 36 | 11 | go irc client for twitch.tv | 2017-03-23 21:31:35 | 2021-06-16 18:55:41 |
| [tenyks](http://tenyks.io) | 169 | 18 | 12 | The Tenyks IRC bot. | 2012-08-26 02:02:24 | 2021-04-16 01:57:30 |
| [hanu](https://sbstjn.com/host-golang-slackbot-on-heroku-with-hanu.html) | 135 | 20 | 2 | Golang Framework for writing Slack bots | 2016-09-16 07:10:42 | 2021-06-15 05:09:35 |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 107 | 4 | 2 | Golang  telegram bot API wrapper, session-based router and middleware | 2016-12-11 06:06:32 | 2021-05-28 08:08:33 |
| [margelet](https://kelpbot.io) | 63 | 10 | 0 | Telegram Bot Framework for Go | 2015-11-21 13:02:17 | 2021-04-27 09:11:42 |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 49 | 5 | 8 | A Discord bot for managing ephemeral roles based upon voice channel member presence. | 2017-12-19 15:20:30 | 2021-06-02 00:52:05 |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 44 | 8 | 1 | Slack bot core/framework written in Go with support for reactions to message updates/deletes | 2015-10-22 04:54:55 | 2021-05-02 22:08:14 |
| [slack-bot](https://github.com/innogames/slack-bot) | 38 | 13 | 3 | Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests... | 2019-07-19 07:49:06 | 2021-06-16 08:53:07 |
| [govkbot](https://github.com/nikepan/govkbot) | 36 | 3 | 1 | VK bot package for Go | 2016-07-11 22:09:54 | 2021-04-08 06:24:17 |
| [micha](https://github.com/onrik/micha) | 17 | 3 | 0 | Client lib for Telegram bot api | 2016-04-14 12:09:44 | 2021-06-16 16:32:41 |
| [echotron](https://github.com/NicoNex/echotron) | 13 | 2 | 0 | Library for telegram bots written in pure go. | 2019-07-22 17:31:49 | 2021-06-10 07:27:49 |
</details>

### Build Automation
2021-06-21 01:48:43

*Last Update: Libraries and tools helping with build automation.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [realize](https://github.com/oxequa/realize) | 4,035 | 211 | 68 | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 2016-07-12 08:07:25 | 2021-06-20 10:20:35 |
| [task](https://taskfile.dev) | 3,560 | 228 | 88 | A task runner / simpler Make alternative written in Go | 2017-02-27 00:46:04 | 2021-06-20 10:40:47 |
| [mmake](https://github.com/tj/mmake) | 1,576 | 43 | 11 | Modern Make  | 2017-02-15 22:01:21 | 2021-06-20 09:46:28 |
| [goyek](https://github.com/goyek/goyek) | 231 | 15 | 3 | Create build pipelines in Go  | 2020-10-11 13:20:55 | 2021-06-20 06:36:35 |
| [taskctl](https://github.com/taskctl/taskctl) | 112 | 8 | 7 | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 2019-11-12 13:19:09 | 2021-06-04 05:26:40 |
| [1build](https://1build.gitbook.io) | 96 | 25 | 31 | Frictionless way of managing project-specific commands | 2019-04-23 17:05:38 | 2021-05-26 16:00:32 |
| [gilbert](https://go-gilbert.github.io/) | 89 | 4 | 0 | Build system and task runner for Go projects | 2019-01-30 09:02:31 | 2021-05-25 10:17:10 |
| [gaper](https://github.com/maxcnunes/gaper) | 49 | 3 | 7 | Builds and restarts a Go project when it crashes or some watched file changes | 2018-06-16 02:46:38 | 2021-06-08 12:55:32 |
| [anko](https://github.com/GuilhermeCaruso/anko) | 16 | 0 | 0 | :crystal_ball: Simple application watcher | 2021-03-02 14:08:42 | 2021-06-18 10:08:52 |
</details>

### CSS Preprocessors
2021-06-21 01:48:46

*Last Update: Libraries for preprocessing CSS files.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gcss](https://github.com/yosssi/gcss) | 442 | 33 | 8 | Pure Go CSS Preprocessor | 2014-09-04 14:38:20 | 2021-05-26 15:17:44 |
| [go-libsass](http://godoc.org/github.com/wellington/go-libsass) | 178 | 22 | 13 | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 2015-04-19 15:09:47 | 2021-06-18 22:31:34 |
</details>

### Command Line - Advanced Console UIs
2021-06-21 01:48:47

*Last Update: Libraries for building Console Applications and Console User Interfaces.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [termui](https://github.com/gizak/termui) | 11,019 | 702 | 82 | Golang terminal dashboard | 2015-02-03 14:09:27 | 2021-06-20 04:01:04 |
| [gocui](https://github.com/jroimartin/gocui) | 7,210 | 481 | 70 | Minimalist Go package aimed at creating Console User Interfaces. | 2014-01-04 02:50:20 | 2021-06-20 01:52:10 |
| [termbox-go](http://godoc.org/github.com/nsf/termbox-go) | 4,065 | 355 | 41 | Pure Go termbox implementation | 2012-01-12 21:03:03 | 2021-06-18 19:33:02 |
| [go-prompt](https://godoc.org/github.com/c-bata/go-prompt) | 3,972 | 238 | 82 | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. | 2017-08-14 16:02:09 | 2021-06-20 04:32:37 |
| [progressbar](https://pkg.go.dev/github.com/schollz/progressbar/v3?tab=doc) | 1,909 | 102 | 9 | A really basic thread-safe progress bar for Golang applications | 2017-10-26 18:28:10 | 2021-06-18 03:21:01 |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1,804 | 120 | 25 | A go library to render progress bars in terminal applications | 2015-11-17 00:59:24 | 2021-06-16 06:52:15 |
| [termdash](http://godoc.org/github.com/nsf/termbox-go) | 1,711 | 86 | 41 | Terminal based dashboard. | 2018-03-24 12:01:49 | 2021-06-18 20:29:54 |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1,690 | 61 | 5 | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. | 2018-06-17 10:37:16 | 2021-06-19 23:02:44 |
| [pterm](https://pterm.sh) | 1,377 | 49 | 19 | ✨ #PTerm is a modern go module to beautify console output. Featuring charts, progressbars, tables, trees, and many more 🚀 It's completely configurable and 100% cross-platform compatible. | 2020-09-17 15:52:59 | 2021-06-18 06:36:16 |
| [uilive](https://github.com/gosuri/uilive) | 1,315 | 65 | 9 | uilive is a go library for updating terminal output in realtime | 2015-11-16 06:13:10 | 2021-06-15 09:16:39 |
| [mpb](https://github.com/vbauerster/mpb) | 1,305 | 78 | 2 | multi progress bar for Go cli applications | 2016-12-14 11:56:29 | 2021-06-16 20:36:46 |
| [aurora](https://github.com/logrusorgru/aurora) | 1,105 | 49 | 4 | Golang ultimate ANSI-colors that supports Printf/Sprintf methods | 2016-11-06 22:37:12 | 2021-06-20 03:59:38 |
| [color](https://pkg.go.dev/github.com/gookit/color?tab=overview) | 858 | 57 | 1 | 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染 | 2018-07-01 07:28:17 | 2021-06-18 19:11:41 |
| [uitable](https://github.com/gosuri/uitable) | 614 | 25 | 3 | A go library to improve readability in terminal apps using tabular data | 2015-11-13 21:59:21 | 2021-06-16 17:14:02 |
| [go-colorable](http://godoc.org/github.com/mattn/go-colorable) | 533 | 67 | 10 | Another Text Attribute Manupulator | 2014-07-30 02:38:06 | 2021-06-19 14:20:30 |
| [go-isatty](http://godoc.org/github.com/mattn/go-isatty) | 525 | 70 | 5 | Change the color of console text. | 2014-04-01 01:53:09 | 2021-06-16 07:17:41 |
| [gommon](https://github.com/labstack/gommon) | 399 | 85 | 15 | Common packages for Go | 2015-03-12 22:35:57 | 2021-06-20 07:48:51 |
| [chalk](https://github.com/ttacon/chalk) | 372 | 18 | 4 | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk | 2014-07-18 19:38:58 | 2021-06-19 05:53:47 |
| [simpletable](https://github.com/alexeyco/simpletable) | 316 | 24 | 2 | Simple tables in terminal with Go | 2017-03-29 07:27:23 | 2021-06-18 06:47:27 |
| [tabby](https://github.com/cheynewallace/tabby) | 291 | 11 | 2 | A tiny library for super simple Golang tables | 2018-12-17 23:35:39 | 2021-06-11 18:16:37 |
| [go-colortext](http://godoc.org/github.com/mattn/go-colorable) | 207 | 19 | 4 | Change the color of console text. | 2013-01-23 03:38:54 | 2021-06-09 04:45:07 |
| [yacspin](https://github.com/theckman/yacspin) | 163 | 4 | 0 | Yet Another CLi Spinner; providing over 70 easy to use and customizable terminal spinners for multiple OSes | 2019-12-29 07:41:23 | 2021-06-20 10:02:07 |
| [box-cli-maker](https://github.com/Delta456/box-cli-maker) | 133 | 2 | 0 | Make Highly Customized Boxes for your CLI | 2020-05-01 07:23:56 | 2021-06-12 19:53:10 |
| [cfmt](https://github.com/mingrammer/cfmt) | 78 | 5 | 1 | :art: Contextual fmt inspired by bootstrap color classes | 2018-03-15 19:04:27 | 2021-03-09 07:48:50 |
| [tabular](https://github.com/InVisionApp/tabular) | 55 | 5 | 0 | Tabular simplifies printing ASCII tables from command line utilities | 2018-04-23 21:17:03 | 2021-05-18 22:23:53 |
| [ctc](https://github.com/wzshiming/ctc) | 33 | 2 | 0 | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method | 2018-04-27 18:07:42 | 2021-06-06 09:01:15 |
| [cfmt](https://github.com/i582/cfmt) | 25 | 0 | 0 | Small library for simple and convenient formatted stylized output to the console. | 2020-11-13 20:29:45 | 2021-06-19 18:30:06 |
| [colourize](https://github.com/TreyBastian/colourize) | 24 | 3 | 0 | An ANSI colour terminal package for Go | 2015-05-11 11:49:39 | 2021-03-30 22:37:59 |
| [marker](https://github.com/cyucelen/marker) | 19 | 13 | 4 |  🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs! | 2019-08-28 15:44:08 | 2021-05-30 14:15:46 |
| [go-ataman](https://github.com/workanator/go-ataman) | 9 | 1 | 0 | Another Text Attribute Manupulator | 2017-05-17 19:04:57 | 2020-12-23 05:34:36 |
| [table](https://github.com/tomlazar/table) | 8 | 0 | 0 | pretty colorfull tables in go with less effort | 2020-09-22 05:42:34 | 2021-05-19 20:56:16 |
</details>

### Command Line - Standard CLI
2021-06-21 01:48:57

*Last Update: Libraries for building standard or basic Command Line applications.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cli](https://github.com/mkideal/cli) | 578 | 38 | 3 | CLI - A package for building command line app with go | 2016-02-26 16:45:29 | 2021-06-11 08:51:31 |
| [argparse](https://github.com/akamensky/argparse) | 335 | 42 | 10 | Argparse for golang. Just because `flag` sucks | 2017-11-24 06:42:20 | 2021-06-19 15:20:15 |
| [climax](http://git.io/climax) | 186 | 17 | 6 | Climax is an alternative CLI with the human face | 2015-11-03 21:04:57 | 2021-06-04 17:16:47 |
| [cli](https://github.com/teris-io/cli) | 91 | 9 | 2 | Simple and complete API for building command line applications in Go | 2017-05-24 23:07:07 | 2021-05-26 03:10:37 |
| [clir](https://github.com/leaanthony/clir) | 66 | 9 | 5 | A Simple and Clear CLI library. Dependency free. | 2019-11-18 19:52:00 | 2021-05-27 07:23:31 |
| [argv](https://github.com/cosiner/argv) | 30 | 6 | 1 | Argparse for golang. Just because `flag` sucks | 2017-01-22 10:37:21 | 2021-03-30 02:32:52 |
</details>

### Continuous Integration
2021-06-21 08:48:16

*Last Update: Tools for help with continuous integration.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [drone](https://drone.io) | 23,336 | 2,281 | 75 | Drone is a Container-Native, Continuous Delivery Platform | 2014-02-07 07:54:44 | 2021-06-20 14:23:32 |
| [cds](https://ovh.github.io/cds/) | 3,444 | 325 | 166 | Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform | 2016-10-11 08:28:23 | 2021-06-20 10:18:34 |
| [goveralls](https://github.com/mattn/goveralls) | 691 | 128 | 10 | A tool for testing, building, signing, and publishing binaries. | 2013-04-17 10:58:40 | 2021-06-07 04:20:28 |
| [overalls](https://github.com/go-playground/overalls) | 106 | 26 | 3 | :jeans:Multi-Package go project coverprofile for tools like goveralls | 2015-07-30 11:30:11 | 2021-03-06 16:05:33 |
| [duci](https://github.com/duck8823/duci) | 65 | 3 | 5 | The simple ci server  | 2018-04-01 01:51:02 | 2021-06-07 11:55:05 |
| [gomason](https://github.com/nikogura/gomason) | 51 | 6 | 1 | A tool for testing, building, signing, and publishing binaries. | 2017-11-18 00:59:11 | 2021-05-21 04:31:31 |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 14 | 4 | 0 | A Go recursive coverage testing tool | 2016-06-26 07:45:32 | 2020-08-20 00:07:58 |
</details>

### Database - Database schema migration
2021-06-22 09:01:53

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [migrate](https://github.com/golang-migrate/migrate) | 6,648 | 698 | 132 | Database migrations. CLI and Golang library. | 2018-01-19 09:30:58 | 2021-06-21 23:13:09 |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 2,161 | 207 | 62 | SQL schema migration tool for Go. | 2014-09-09 07:31:41 | 2021-06-22 01:50:49 |
| [goose](https://github.com/pressly/goose) | 1,732 | 266 | 43 | Goose database migration tool - fork of https://bitbucket.org/liamstask/goose | 2016-02-25 20:39:37 | 2021-06-21 13:38:43 |
| [pop](https://github.com/gobuffalo/pop) | 1,077 | 217 | 119 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2021-06-20 04:52:13 |
| [skeema](http://pravasan.github.io/pravasan/) | 907 | 81 | 30 | Schema management CLI for MySQL | 2016-10-31 23:18:56 | 2021-06-20 16:08:43 |
| [gormigrate](https://pkg.go.dev/github.com/go-gormigrate/gormigrate/v2?tab=doc) | 620 | 67 | 8 | Minimalistic database migration helper for Gorm ORM | 2016-08-31 11:46:23 | 2021-06-21 12:39:34 |
| [darwin](https://github.com/GuiaBolso/darwin) | 119 | 21 | 4 | Database schema evolution library for Go | 2016-04-05 15:57:59 | 2021-06-14 08:42:36 |
| [migrator](https://github.com/lopezator/migrator) | 114 | 15 | 5 | Dead simple Go database migration library. | 2019-02-04 09:40:01 | 2021-06-19 16:41:10 |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 78 | 17 | 2 | A Go package to help write migrations with go-pg/pg. | 2018-08-11 07:00:13 | 2021-05-31 04:36:41 |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 26 | 8 | 0 | Django style fixtures for Golang's excellent built-in database/sql library. | 2015-12-24 11:27:45 | 2021-04-08 02:57:47 |
| [pravasan](http://pravasan.github.io/pravasan/) | 24 | 4 | 30 | Simple Migration Tool - written in Go | 2015-01-03 17:11:05 | 2019-03-22 13:54:35 |
| [avro](https://github.com/khezen/avro) | 24 | 3 | 0 | Apache AVRO for go | 2019-04-07 12:22:46 | 2021-06-13 12:52:12 |
| [migrator](https://github.com/larapulse/migrator) | 8 | 1 | 0 | MySQL database migrator | 2020-06-27 14:40:29 | 2021-06-14 12:50:02 |
| [schema](http://pravasan.github.io/pravasan/) | 7 | 1 | 3 | Embedded schema migration package for Go | 2019-09-24 19:27:13 | 2021-01-05 05:27:47 |
| [go-pg-migrate](https://pkg.go.dev/github.com/lawzava/go-pg-migrate) | 4 | 1 | 0 | CLI-friendly package for go-pg migrations management. | 2021-01-16 17:01:32 | 2021-02-17 10:16:36 |
</details>

### Database Drivers - Multiple Backends.
2021-06-21 16:25:23

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cayley](https://cayley.io) | 13,841 | 1,252 | 89 | An open-source graph database | 2014-06-05 18:49:41 | 2021-06-21 06:17:17 |
| [gokv](https://github.com/philippgille/gokv) | 343 | 34 | 30 | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) | 2018-10-08 18:55:22 | 2021-06-15 10:45:51 |
| [cachego](https://github.com/faabiosr/cachego) | 156 | 9 | 1 | Golang Cache component - Multiple drivers | 2016-10-05 18:10:03 | 2021-06-03 07:59:30 |
| [dsc](https://github.com/viant/dsc) | 22 | 6 | 0 | Datastore Connectivity in go | 2016-06-13 20:18:10 | 2021-04-17 04:50:27 |
</details>

### Database Drivers - Search and Analytic Databases
2021-06-22 09:02:00

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [bleve](https://github.com/blevesearch/bleve) | 7,616 | 577 | 269 | A modern text indexing library for go | 2014-04-17 21:02:18 | 2021-06-22 01:16:28 |
| [elastic](https://olivere.github.io/elastic/) | 5,998 | 1,014 | 69 | Elasticsearch client for Go. | 2012-12-06 17:15:33 | 2021-06-22 01:53:17 |
| [riot](https://github.com/go-ego/riot) | 5,800 | 436 | 50 | Go Open Source, Distributed, Simple and efficient Search Engine  | 2017-06-21 14:17:59 | 2021-06-21 20:54:46 |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch#go-elasticsearch) | 3,452 | 360 | 60 | The official Go client for Elasticsearch | 2017-03-27 17:56:15 | 2021-06-21 22:24:26 |
| [elastigo](https://github.com/mattbaird/elastigo) | 947 | 256 | 73 | A Go (golang) based Elasticsearch client library. | 2012-10-12 04:19:59 | 2021-06-21 14:33:23 |
| [elasticsql](https://github.com/cch123/elasticsql) | 735 | 138 | 7 | convert sql to elasticsearch DSL in golang(go) | 2016-08-24 07:29:43 | 2021-06-21 08:41:28 |
| [skizze](https://github.com/seiflotfy/skizze) | 78 | 8 | 0 | A probabilistic data structure service and storage | 2016-01-17 12:10:40 | 2021-06-07 13:47:40 |
| [goes](http://godoc.org/github.com/belogik/goes) | 24 | 12 | 0 | A library to interact with Elasticsearch in Go! | 2015-12-28 18:52:03 | 2019-03-11 09:09:21 |
</details>

### Dynamic DNS
2021-06-21 16:25:26

*Last Update: Tools for updating dynamic DNS records.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [godns](https://xiaozhou.net/godns-project-2014-05-18.html) | 830 | 158 | 7 | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 2014-05-11 11:49:17 | 2021-06-21 07:17:44 |
| [ddns](https://github.com/skibish/ddns) | 192 | 19 | 0 | Personal DDNS client with Digital Ocean Networking DNS as backend. | 2017-03-13 21:02:27 | 2021-06-20 17:48:23 |
</details>

### Functional
2021-06-21 16:25:25

*Last Update: Packages to support functional programming in Go.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1,197 | 65 | 4 |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  | 2014-07-02 10:27:16 | 2021-06-18 03:20:33 |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 179 | 12 | 0 | Monad, Functional Programming features for Golang | 2018-05-24 09:08:45 | 2021-06-15 09:03:59 |
| [fuego](https://github.com/seborama/fuego) | 91 | 8 | 0 | Functional Experiment in Golang | 2018-11-05 22:24:09 | 2021-05-29 06:02:07 |
| [gofp](https://github.com/rbrahul/gofp) | 74 | 3 | 0 | A super simple Lodash like utility library with essential functions that empowers the development in Go | 2021-02-19 00:01:39 | 2021-06-14 20:02:58 |
</details>

### Microsoft Office
2021-06-21 11:28:15

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [unioffice](https://unidoc.io/unioffice/) | 2,885 | 321 | 25 | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents | 2017-08-29 01:25:48 | 2021-06-21 02:00:29 |
</details>

### Miscellaneous - Strings
2021-06-22 09:02:03

*Last Update: Libraries for working with strings.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xstrings](https://github.com/huandu/xstrings) | 899 | 61 | 0 | Implements string functions widely used in other languages but absent in Go. | 2015-01-06 07:25:26 | 2021-06-21 14:33:31 |
| [strutil](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 121 | 12 | 1 | String utilities for Go | 2018-08-16 06:56:15 | 2021-06-04 00:46:11 |
| [stringy](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 67 | 5 | 0 | Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package. | 2020-04-03 03:34:10 | 2021-06-06 00:05:41 |
</details>

### Package Management - Official
2021-06-21 08:48:20

*Last Update: Official experimental tooling for package management*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [dep](https://golang.github.io/dep/) | 13,162 | 1,104 | 0 | Go dependency management tool experiment (deprecated) | 2016-10-07 00:04:51 | 2021-06-20 13:32:47 |
</details>

### Query Language
2021-06-22 09:01:41

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [graphql](https://github.com/graphql-go/graphql) | 7,744 | 676 | 169 | An implementation of GraphQL for Go / Golang | 2015-07-19 12:25:43 | 2021-06-21 22:11:45 |
| [gqlgen](https://gqlgen.com) | 6,155 | 654 | 229 | go generate based graphql server library | 2018-02-11 04:54:11 | 2021-06-21 23:26:52 |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3,770 | 405 | 86 | GraphQL server with a focus on ease of use | 2016-10-18 13:57:24 | 2021-06-21 22:11:16 |
| [gojsonq](https://github.com/thedevsaddam/gojsonq/wiki) | 1,690 | 95 | 11 | A simple Go package to Query over JSON/YAML/XML/CSV Data  | 2018-05-19 16:15:18 | 2021-06-18 14:39:08 |
| [dasel](https://daseldocs.tomwright.me) | 925 | 18 | 6 | Query, update and convert data structures from the command line. Comparable to jq/yq but supports JSON, TOML, YAML, XML and CSV with zero runtime dependencies. | 2020-09-22 10:33:56 | 2021-06-21 22:48:39 |
| [jsonql](https://github.com/elgs/jsonql) | 246 | 36 | 5 | JSON query expression library in Golang. | 2015-12-29 11:24:04 | 2021-04-21 12:34:24 |
| [rql](https://github.com/a8m/rql) | 212 | 23 | 11 | Resource Query Language for REST | 2018-06-05 18:37:29 | 2021-06-21 07:45:55 |
| [graphql](https://github.com/tmc/graphql) | 53 | 6 | 3 | graphql parser + utilities | 2015-04-18 21:05:52 | 2021-05-27 16:58:59 |
| [jsonslice](https://github.com/bhmj/jsonslice) | 52 | 4 | 4 | json slicer | 2018-05-02 00:33:15 | 2021-04-03 11:04:51 |
| [api-fu](https://github.com/ccbrown/api-fu) | 31 | 0 | 2 | A collection of Go packages for creating robust GraphQL APIs | 2019-07-30 05:18:43 | 2021-06-03 01:49:05 |
| [straf](https://github.com/SonicRoshan/straf) | 26 | 3 | 0 | Convert Golang Struct To GraphQL Object On The Fly | 2019-08-16 13:31:39 | 2021-03-21 18:27:21 |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 20 | 2 | 0 | Query Parser for REST | 2020-02-10 17:58:42 | 2021-06-14 12:37:18 |
| [jsonpath](https://github.com/AsaiYusuke/jsonpath) | 5 | 1 | 0 | A query library for retrieving part of JSON based on JSONPath syntax. | 2020-11-29 05:37:26 | 2021-04-27 16:37:54 |
| [gws](https://github.com/Zaba505/gws) | 4 | 1 | 2 | A WebSocket client and server for GraphQL | 2020-06-08 19:51:36 | 2020-11-16 02:47:54 |
</details>

### Server Applications
2021-06-21 16:25:01

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [etcd](https://etcd.io) | 36,265 | 7,739 | 187 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2021-06-21 08:05:23 |
| [caddy](https://caddyserver.com) | 33,712 | 2,714 | 72 | Fast, multi-platform web server with automatic HTTPS | 2015-01-13 19:45:03 | 2021-06-21 09:24:44 |
| [minio](https://min.io/download) | 28,158 | 3,070 | 31 | High Performance, Kubernetes Native Object Storage | 2015-01-14 19:23:58 | 2021-06-21 09:18:09 |
| [consul](https://www.consul.io) | 22,393 | 3,753 | 894 | Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. | 2013-11-04 22:15:27 | 2021-06-21 07:21:08 |
| [nsq](https://nsq.io) | 19,790 | 2,562 | 62 | A realtime distributed messaging platform | 2012-05-12 14:37:08 | 2021-06-21 05:36:21 |
| [roadrunner](https://roadrunner.dev/) | 5,677 | 305 | 37 | High-performance PHP application server, load-balancer and process manager written in Golang | 2017-12-26 16:13:10 | 2021-06-21 06:55:24 |
| [devd](https://github.com/cortesi/devd) | 3,129 | 137 | 21 |  A local webserver for developers | 2015-09-27 22:43:00 | 2021-06-12 08:09:25 |
| [sftpgo](https://github.com/drakkan/sftpgo) | 2,685 | 200 | 8 | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-07-20 10:18:31 | 2021-06-21 09:11:40 |
| [algernon](https://algernon.roboticoverlords.org) | 1,819 | 98 | 8 | :tophat: Small self-contained pure-Go web server with Lua, Markdown, HTTP/2, QUIC, Redis and PostgreSQL support | 2015-03-10 11:25:30 | 2021-06-16 07:10:55 |
| [flagr](https://checkr.github.io/flagr) | 1,588 | 138 | 66 | Flagr is a feature flagging, A/B testing and dynamic configuration microservice | 2017-10-03 19:07:32 | 2021-06-18 08:57:15 |
| [fider](https://getfider.com) | 1,578 | 440 | 30 | Open platform to collect and prioritize product feedback | 2017-01-17 22:55:19 | 2021-06-20 10:40:32 |
| [flipt](https://flipt.io) | 1,500 | 67 | 9 | An open-source, on-prem feature flag solution | 2016-11-05 00:09:07 | 2021-06-21 09:10:01 |
| [trickster](https://trickstercache.org) | 1,424 | 139 | 19 | Open Source HTTP Reverse Proxy Cache and Time Series Dashboard Accelerator | 2018-03-29 20:31:44 | 2021-06-20 05:30:17 |
| [discovery](https://github.com/bilibili/discovery) | 1,419 | 332 | 18 | A registry for resilient mid-tier load balancing and failover. | 2018-04-20 12:57:50 | 2021-06-20 12:04:27 |
| [jackal](https://github.com/ortuman/jackal) | 1,070 | 87 | 8 | 💬 Instant messaging server for the Extensible Messaging and Presence Protocol (XMPP). | 2017-11-13 18:17:48 | 2021-06-21 08:25:41 |
| [go-feature-flag](https://thomaspoignant.github.io/go-feature-flag) | 316 | 8 | 8 | A simple and complete feature flag solution, without any complex backend system to install, all you need is a file as your backend. 🎛️ | 2020-12-11 13:19:17 | 2021-06-18 20:48:39 |
| [dudeldu](https://github.com/krotik/dudeldu) | 127 | 12 | 0 | A simple SHOUTcast server. | 2016-09-07 19:11:04 | 2021-05-15 05:13:50 |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 55 | 6 | 31 | Reverse proxy with automatically obtains TLS certificates from Let's Encrypt | 2019-04-12 05:39:43 | 2021-06-20 10:18:31 |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 28 | 3 | 1 | Stream database events from PostgreSQL to Kafka | 2019-04-28 21:55:31 | 2021-05-01 06:49:46 |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 22 | 3 | 0 | Turn Nginx logs into Prometheus metrics | 2018-10-23 09:10:27 | 2021-05-21 16:30:18 |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 18 | 2 | 5 | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-12-18 12:48:14 | 2021-06-05 11:57:06 |
| [protoxy](https://github.com/camgraff/protoxy) | 16 | 1 | 0 | A proxy server than converts JSON request bodies to protocol buffers | 2020-09-03 23:24:34 | 2021-05-21 16:35:18 |
| [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) | 15 | 7 | 1 | Prometheus remote write proxy that adds Cortex tenant ID based on metric labels | 2020-10-06 16:52:25 | 2021-06-17 13:18:31 |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 9 | 2 | 42 | Simple Reverse Proxy with Caching, written in Go, using Redis. | 2020-11-12 15:10:40 | 2021-05-21 04:32:25 |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 0 | 0 | Service for relaying Riemann events to Riemann/Carbon destinations | 2019-04-23 14:17:12 | 2019-10-29 15:00:17 |
</details>

### Stream Processing
2021-06-22 09:01:59

*Last Update: Libraries and tools for stream processing and reactive programming.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-streams](https://pkg.go.dev/github.com/reugn/go-streams) | 662 | 43 | 1 | A lightweight stream processing library for Go | 2019-04-30 17:28:15 | 2021-06-21 05:05:23 |
| [machine](https://pkg.go.dev/github.com/whitaker-io/machine) | 85 | 4 | 3 | Machine is a workflow/pipeline library for processing data | 2020-10-13 04:24:19 | 2021-06-13 13:32:17 |
| [stream](https://github.com/youthlin/stream) | 37 | 2 | 0 | Go Stream, like Java 8 Stream. | 2020-11-12 03:52:50 | 2021-05-31 13:41:51 |
</details>

### Template Engines
2021-06-21 08:47:59

*Last Update: Libraries and tools for templating and lexing.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gofpdf](http://godoc.org/github.com/jung-kurt/gofpdf) | 3,815 | 597 | 56 | A PDF document generator with high level support for text, drawing and images | 2015-03-13 11:57:30 | 2021-06-20 09:33:58 |
| [sprig](http://masterminds.github.io/sprig/) | 2,378 | 263 | 78 | Useful template functions for Go templates. | 2013-11-22 01:20:40 | 2021-06-20 18:49:26 |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 2,118 | 113 | 27 | Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template | 2016-03-06 21:42:01 | 2021-06-19 20:58:46 |
| [pongo2](https://www.schlachter.tech/pongo2) | 1,983 | 193 | 54 | Django-syntax like template-engine for Go | 2013-08-23 01:00:08 | 2021-06-19 21:31:38 |
| [hero](https://shiyanhui.github.io/hero) | 1,462 | 96 | 27 | A handy, fast and powerful go template engine. | 2017-01-15 13:31:50 | 2021-06-18 02:51:28 |
| [mustache](https://github.com/hoisie/mustache) | 1,013 | 186 | 32 | The mustache template language in Go | 2009-12-30 21:05:05 | 2021-05-07 07:50:12 |
| [amber](https://github.com/eknkc/amber) | 873 | 58 | 23 | Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade | 2012-10-31 20:27:24 | 2021-06-03 05:44:37 |
| [ace](https://github.com/yosssi/ace) | 801 | 44 | 28 | HTML template engine for Go | 2014-07-13 13:39:19 | 2021-06-03 05:47:01 |
| [jet](https://shiyanhui.github.io/hero) | 791 | 79 | 16 | Jet  template engine | 2016-03-31 16:53:36 | 2021-06-18 13:00:23 |
| [gorazor](https://github.com/sipin/gorazor) | 781 | 86 | 2 | Razor view engine for go | 2014-05-01 05:30:31 | 2021-06-10 11:50:41 |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 503 | 61 | 9 | Simple and fast template engine for Go | 2015-08-19 12:44:22 | 2021-06-19 20:59:06 |
| [ego](https://github.com/benbjohnson/ego) | 490 | 35 | 10 | An ERB-style templating language for Go. | 2014-02-23 18:14:41 | 2021-06-18 11:13:37 |
| [raymond](https://github.com/aymerick/raymond) | 430 | 61 | 18 | Handlebars for golang | 2015-04-22 13:07:59 | 2021-06-04 06:08:06 |
| [maroto](https://pkg.go.dev/github.com/johnfercher/maroto?tab=doc) | 381 | 68 | 19 | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. | 2019-05-20 23:27:47 | 2021-06-19 10:01:11 |
| [goview](https://github.com/foolin/goview) | 234 | 22 | 10 | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. | 2019-04-14 11:22:41 | 2021-06-13 04:24:56 |
| [soy](https://github.com/robfig/soy) | 155 | 36 | 6 | Go implementation for Soy templates (Google Closure templates) | 2013-12-15 01:14:48 | 2021-02-04 18:50:34 |
| [liquid](https://godoc.org/github.com/osteele/liquid) | 125 | 28 | 13 | A Liquid template engine in Go | 2017-06-26 14:39:52 | 2021-06-18 05:32:56 |
| [velvet](http://masterminds.github.io/sprig/) | 73 | 6 | 2 | A sweet velvety templating package | 2016-12-29 16:46:57 | 2021-06-13 16:26:57 |
| [kasia.go](https://github.com/ziutek/kasia.go) | 72 | 7 | 2 | Templating system for HTML and other text documents - go implementation | 2010-12-07 10:46:05 | 2021-05-05 02:28:01 |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 38 | 9 | 2 | Wrapper package for Go's template/html to allow for easy file-based template inheritance. | 2018-08-10 20:34:19 | 2021-06-18 14:44:18 |
| [gospin](https://github.com/m1/gospin) | 27 | 6 | 3 | Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations | 2019-02-22 17:04:51 | 2021-05-21 04:32:26 |
| [damsel](https://github.com/dskinner/damsel) | 24 | 4 | 1 | Package damsel provides html outlining via css-selectors and common template functionality. | 2012-05-02 23:06:48 | 2020-09-12 23:20:49 |
</details>

### Testing - Fail injection
2021-06-21 11:27:55

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [failpoint](https://github.com/pingcap/failpoint) | 588 | 51 | 10 | An implementation of failpoints for Golang. | 2019-04-02 07:48:18 | 2021-06-07 09:39:43 |
</details>

### Testing - Fuzzing and delta-debugging, reducing, shrinking
2021-06-22 09:01:52

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 4,067 | 233 | 50 | Randomized testing for Go | 2015-04-15 13:07:50 | 2021-06-18 18:15:59 |
| [gofuzz](https://github.com/google/gofuzz) | 1,089 | 100 | 12 | Fuzz testing for go. | 2014-07-31 16:21:29 | 2021-06-21 12:58:48 |
| [tavor](https://github.com/zimmski/tavor) | 232 | 8 | 53 | A generic fuzzing and delta-debugging framework | 2014-05-18 14:59:14 | 2021-04-27 23:55:43 |
</details>

### Testing - Mock
2021-06-22 09:01:46

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mock](https://pkg.go.dev/github.com/h2non/gock) | 5,745 | 436 | 32 | GoMock is a mocking framework for the Go programming language. | 2015-06-12 17:15:11 | 2021-06-22 01:14:15 |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 3,579 | 289 | 49 | Sql mock driver for golang to test database interactions | 2014-02-07 07:59:29 | 2021-06-22 00:39:17 |
| [hoverfly](https://hoverfly.io) | 1,748 | 173 | 32 | Lightweight service virtualization/API simulation tool for developers and testers | 2015-11-30 16:36:31 | 2021-06-15 12:33:18 |
| [gock](https://pkg.go.dev/github.com/h2non/gock) | 1,390 | 74 | 29 | HTTP traffic mocking and testing made easy in Go ༼ʘ̚ل͜ʘ̚༽ | 2016-03-02 16:20:26 | 2021-06-20 14:43:10 |
| [httpmock](http://godoc.org/github.com/jarcoal/httpmock) | 1,087 | 81 | 1 | HTTP mocking for Golang | 2014-02-24 16:47:59 | 2021-06-21 08:29:31 |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 540 | 68 | 15 | A tool for generating self-contained, type-safe test doubles in go | 2014-05-21 00:12:54 | 2021-06-20 23:58:41 |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 368 | 32 | 3 | Immutable transaction isolated sql driver for golang | 2015-07-08 07:34:53 | 2021-06-18 21:07:05 |
| [minimock](https://github.com/gojuno/minimock) | 352 | 23 | 10 | Powerful mock generation tool for Go programming language | 2016-08-03 16:01:35 | 2021-06-21 15:32:26 |
| [govcr](https://github.com/seborama/govcr) | 98 | 13 | 4 | HTTP mock for Golang: record and replay HTTP/HTTPS interactions for offline testing | 2016-07-10 17:47:41 | 2021-05-17 01:46:44 |
| [timex](https://github.com/cabify/timex) | 54 | 2 | 0 | A test-friendly replacement for golang's time package | 2020-01-02 18:06:48 | 2021-05-12 10:18:55 |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 5 | 0 | Mock object for Go http.ResponseWriter | 2011-06-11 16:03:01 | 2020-08-05 04:12:58 |
| [go-localstack](https://github.com/elgohr/go-localstack) | 12 | 3 | 1 | Go Wrapper for using localstack | 2020-03-18 07:13:02 | 2021-06-18 13:35:56 |
| [mockit](https://github.com/pasdam/mockit) | 8 | 1 | 2 | Library that make mocking of Go functions/methods easy | 2020-01-01 08:46:09 | 2021-06-07 19:43:47 |
</details>

### Testing - Selenium and browser control tools
2021-06-21 16:25:20

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [chromedp](https://github.com/chromedp/chromedp) | 6,377 | 533 | 32 | A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. | 2017-01-24 14:54:30 | 2021-06-21 06:55:43 |
| [selenoid](https://aerokube.com/selenoid/latest/) | 1,923 | 251 | 158 | Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary. | 2016-08-22 09:11:16 | 2021-06-21 08:56:07 |
| [rod](https://go-rod.github.io) | 1,599 | 109 | 44 | A Devtools driver for web automation and scraping | 2020-01-21 20:09:45 | 2021-06-18 02:33:35 |
| [cdp](https://github.com/mafredri/cdp) | 554 | 39 | 12 | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 2017-03-12 10:25:41 | 2021-06-19 00:16:16 |
| [playwright-go](https://mxschmitt.github.io/playwright-go/) | 392 | 40 | 16 | Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API. | 2020-08-16 12:46:14 | 2021-06-19 23:30:28 |
| [ggr](https://aerokube.com/ggr/latest/) | 271 | 57 | 13 | A lightweight load balancer used to create big Selenium clusters | 2016-06-16 15:33:24 | 2021-06-19 21:14:25 |
</details>

### Text Processing - Utility
2021-06-21 16:25:16

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xurls](https://tysug.net) | 793 | 92 | 1 | Extract urls from text | 2015-01-12 01:28:46 | 2021-06-18 07:32:33 |
| [gotabulate](https://github.com/bndr/gotabulate) | 266 | 27 | 6 | Gotabulate - Easily pretty-print your tabular data with Go | 2014-08-21 07:44:28 | 2021-05-30 11:57:22 |
| [radix](https://github.com/yourbasic/radix) | 172 | 9 | 0 | A fast string sorting algorithm (MSD radix sort) | 2017-06-09 14:38:58 | 2021-06-03 06:57:20 |
| [regroup](https://github.com/oriser/regroup) | 82 | 6 | 1 | Match regex group into go struct using struct tags and automatic parsing | 2020-09-08 19:04:42 | 2021-06-09 06:01:11 |
| [parth](https://github.com/codemodus/parth) | 40 | 4 | 0 | Path parsing for segment unmarshaling and slicing. | 2015-04-06 22:53:59 | 2021-03-27 05:04:42 |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 35 | 3 | 1 | A sanitization-based swear filter for Go. | 2018-09-09 00:07:26 | 2021-06-13 13:35:40 |
| [xj2go](https://tysug.net) | 20 | 4 | 0 | Convert xml and json to go struct | 2017-09-19 13:20:57 | 2020-08-31 17:04:37 |
| [kace](https://github.com/codemodus/kace) | 16 | 1 | 1 | Common case conversions covering common initialisms. | 2015-06-04 20:36:49 | 2020-09-27 22:30:36 |
| [tagify](https://www.zoomio.org/tagify) | 16 | 0 | 0 | Tagify produces a set of tags from a given source. Source can be either an HTML page, a Markdown document or a plain text. Supports English, Russian, Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean languages. | 2018-03-20 10:30:11 | 2021-05-21 04:32:34 |
| [TySug](https://tysug.net) | 9 | 1 | 0 | A project around helping to prevent typing typos. TySug (Typo Suggestions) suggests alternative words with respect to keyboard layouts | 2018-06-05 19:46:29 | 2021-01-23 19:41:36 |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 8 | 3 | 1 | A string argument parser that understands quotes and backslashes | 2016-02-24 00:53:38 | 2020-05-16 17:37:29 |
| [textwrap](https://www.zoomio.org/tagify) | 2 | 0 | 1 | Port of Python's "textwrap" module to Go | 2019-07-26 17:57:55 | 2020-12-22 12:39:59 |
</details>

### UUID
2021-06-21 11:28:12

*Last Update: Libraries for working with UUIDs.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [uuid](https://github.com/google/uuid) | 2,721 | 255 | 12 | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. | 2016-02-12 22:17:59 | 2021-06-19 06:50:15 |
| [ulid](https://github.com/oklog/ulid) | 2,285 | 89 | 2 | Universally Unique Lexicographically Sortable Identifier (ULID) in Go | 2016-12-06 15:26:52 | 2021-06-19 11:26:27 |
| [uuid](https://gof.rs/projects/uuid/) | 938 | 60 | 9 | A UUID package originally forked from github.com/satori/go.uuid | 2018-07-13 02:13:28 | 2021-06-20 21:46:57 |
| [wuid](https://github.com/edwingeng/wuid) | 418 | 36 | 0 | An extremely fast UUID alternative written in golang | 2018-01-27 01:16:28 | 2021-05-31 11:26:56 |
| [sno](https://pkg.go.dev/github.com/muyo/sno?tab=doc) | 46 | 2 | 1 | Compact, sortable and fast unique IDs with embedded metadata. | 2019-05-26 22:05:26 | 2021-05-21 19:21:41 |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 32 | 2 | 0 | A tiny and fast Go unique string generator | 2019-07-02 12:15:56 | 2021-06-06 04:41:36 |
| [Goid](https://github.com/JakeHL/Goid) | 30 | 1 | 1 | A UUIDv4 generation package written in go | 2017-05-19 10:40:45 | 2021-04-02 13:06:38 |
| [uuid](https://github.com/agext/uuid) | 12 | 3 | 0 | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. | 2016-02-03 03:02:51 | 2021-02-22 18:59:07 |
| [gouid](https://github.com/twharmon/gouid) | 8 | 0 | 0 | Fast, dependable universally unique ids | 2020-10-08 19:54:41 | 2021-06-04 08:39:15 |
| [GoFlake](https://github.com/Hart87/GoFlake) | 3 | 0 | 0 | A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake. | 2021-05-03 14:44:19 | 2021-06-13 15:39:26 |
</details>

### Validation
2021-06-21 08:48:13

*Last Update: Libraries for validation.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [validator](https://github.com/go-playground/validator) | 8,032 | 709 | 134 | :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving | 2015-02-12 16:32:22 | 2021-06-20 16:57:03 |
| [govalidator](https://github.com/asaskevich/govalidator) | 4,803 | 484 | 131 | [Go] Package of validators and sanitizers for strings, numerics, slices and structs | 2014-06-20 10:45:23 | 2021-06-21 00:41:57 |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 2,154 | 135 | 20 | An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags. | 2016-06-22 03:47:43 | 2021-06-20 16:57:26 |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 995 | 89 | 28 | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. | 2017-09-13 16:42:20 | 2021-06-18 06:20:26 |
| [validate](https://pkg.go.dev/github.com/gookit/validate) | 428 | 62 | 7 | ⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。 | 2018-07-16 08:23:49 | 2021-06-20 05:07:54 |
| [checkdigit](https://github.com/osamingo/checkdigit) | 77 | 2 | 0 | Provide check digit algorithms and calculators written in Go | 2019-04-05 09:46:36 | 2021-03-29 18:47:26 |
| [validate](https://pkg.go.dev/github.com/gookit/validate) | 59 | 18 | 5 | This package provides a framework for writing validations for Go applications. | 2018-02-10 18:25:55 | 2021-04-01 19:07:51 |
| [jio](https://github.com/faceair/jio) | 58 | 9 | 0 | jio is a json schema validator similar to joi | 2018-10-28 11:02:45 | 2021-05-26 03:38:40 |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 50 | 4 | 6 | A norms and conventions validator for Terraform | 2019-05-29 11:37:15 | 2021-05-03 21:44:55 |
| [gody](https://pkg.go.dev/github.com/guiferpa/gody) | 49 | 3 | 1 | :balloon: A lightweight struct validator for Go | 2018-11-01 21:08:16 | 2021-02-19 02:07:17 |
| [govalid](https://github.com/twharmon/govalid) | 22 | 3 | 0 | Struct validation using tags | 2019-02-17 23:25:43 | 2021-04-27 16:03:31 |
</details>

### Version Control
2021-06-21 08:48:11

*Last Update: Libraries for version control.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-git](https://pkg.go.dev/github.com/go-git/go-git/v5) | 2,403 | 252 | 190 | A highly extensible Git implementation in pure Go. | 2019-12-19 10:27:02 | 2021-06-20 09:37:26 |
| [git2go](https://github.com/libgit2/git2go) | 1,632 | 284 | 42 | Git to Go; bindings for libgit2. Like McDonald's but tastier. | 2013-03-05 19:50:43 | 2021-06-18 03:20:23 |
| [hercules](https://github.com/src-d/hercules) | 1,359 | 117 | 37 | Gaining advanced insights from Git repository history. | 2016-12-12 17:30:29 | 2021-06-20 03:01:59 |
| [gh](https://github.com/rjeczalik/gh) | 75 | 12 | 2 | Scriptable server and net/http middleware for GitHub Webhooks. | 2015-03-08 21:04:05 | 2020-05-08 16:42:07 |
| [go-vcs](https://sourcegraph.com/sourcegraph/go-vcs) | 75 | 21 | 23 | manipulate and inspect VCS repositories in Go | 2013-06-02 02:36:18 | 2021-06-09 06:01:11 |
| [hgo](https://github.com/beyang/hgo) | 13 | 3 | 0 | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. | 2014-06-18 03:54:40 | 2020-05-05 03:52:16 |
</details>

### Video
2021-06-21 08:48:07

*Last Update: Libraries for manipulating video.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [goav](https://github.com/giorgisio/goav) | 1,637 | 294 | 44 | Golang bindings for FFmpeg | 2015-05-21 05:31:14 | 2021-06-20 19:34:31 |
| [m3u8](http://tools.ietf.org/html/draft-pantos-http-live-streaming) | 834 | 218 | 43 | Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema: | 2013-02-05 22:26:30 | 2021-06-19 07:17:03 |
| [gmf](https://github.com/3d0c/gmf) | 691 | 134 | 39 | Go Media Framework | 2013-04-03 09:07:47 | 2021-06-13 00:21:08 |
| [go-astits](https://github.com/asticode/go-astits) | 369 | 34 | 6 | Demux and mux MPEG Transport Streams (.ts) natively in GO | 2017-07-04 13:06:15 | 2021-06-20 18:21:06 |
| [go-astisub](https://github.com/asticode/go-astisub) | 316 | 61 | 5 | Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.) | 2016-12-16 14:47:59 | 2021-06-19 18:44:28 |
| [libvlc-go](https://pkg.go.dev/github.com/adrg/libvlc-go/v3) | 217 | 31 | 2 | Go bindings for libVLC and high-level media player interface | 2015-01-06 14:01:50 | 2021-06-19 07:40:44 |
| [gst](https://github.com/ziutek/gst) | 159 | 46 | 9 | Go bindings for GStreamer (retired: currently I don't use/develop this package) | 2011-07-26 00:44:40 | 2021-04-21 01:25:11 |
| [go-m3u8](https://tools.ietf.org/html/rfc8216) | 75 | 10 | 0 | Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8) | 2018-11-06 02:42:27 | 2021-06-02 05:43:35 |
| [v4l](https://pkg.go.dev/github.com/korandiz/v4l) | 58 | 10 | 0 | Facade to the Video4Linux video capture interface.  | 2016-10-25 10:50:25 | 2021-05-16 09:56:04 |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 14 | 3 | 0 | golang library to read and write various subtitle formats | 2017-05-03 21:05:25 | 2021-05-20 14:04:11 |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 9 | 1 | 0 | Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files | 2018-11-02 19:09:07 | 2021-06-13 08:14:53 |
</details>

### Web Frameworks
2021-06-21 11:27:56

*Last Update: Full stack web frameworks.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gin](https://gin-gonic.com/) | 48,899 | 5,548 | 404 | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 2014-06-16 23:57:25 | 2021-06-21 03:22:11 |
| [echo](https://echo.labstack.com) | 20,102 | 1,778 | 54 | High performance, minimalist Go web framework | 2015-03-01 17:43:01 | 2021-06-21 02:10:44 |
| [fiber](https://gofiber.io) | 13,864 | 656 | 23 | ⚡️ Express inspired web framework written in Go | 2020-01-16 03:59:20 | 2021-06-21 02:55:32 |
| [revel](http://revel.github.io) | 12,279 | 1,395 | 99 | A high productivity, full-stack web framework for the Go language. | 2011-12-09 04:10:26 | 2021-06-21 02:53:38 |
| [buffalo](http://gobuffalo.io) | 6,268 | 483 | 87 | Rapid Web Development w/ Go | 2014-10-22 17:35:14 | 2021-06-20 22:28:30 |
| [goa](https://goa.design) | 4,279 | 458 | 38 | Design-based APIs and microservices in Go | 2014-12-05 07:17:53 | 2021-06-19 18:03:24 |
| [go-json-rest](https://ant0ine.github.io/go-json-rest/) | 3,473 | 387 | 49 | A quick and easy way to setup a RESTful JSON API | 2013-02-19 03:15:45 | 2021-06-13 09:23:43 |
| [gizmo](https://open.nytimes.com/introducing-gizmo-aa7ea463b208) | 3,431 | 227 | 28 | A Microservice Toolkit from The New York Times | 2015-12-15 18:09:36 | 2021-06-20 08:31:08 |
| [macaron](https://go-macaron.com) | 3,167 | 283 | 8 | Package macaron is a high productive and modular web framework in Go. | 2014-07-10 03:13:30 | 2021-06-20 09:20:49 |
| [utron](https://github.com/gernest/utron) | 2,198 | 158 | 9 | A lightweight MVC framework for Go(Golang) | 2015-09-16 07:55:54 | 2021-06-18 03:20:20 |
| [rest-layer]( http://rest-layer.io) | 1,083 | 91 | 33 | REST Layer, Go (golang) REST API framework | 2015-07-29 19:16:20 | 2021-06-19 23:57:59 |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 996 | 74 | 28 | A Go framework for building JSON web services inspired by Dropwizard | 2013-02-09 21:16:13 | 2021-06-13 22:08:50 |
| [goyave](https://goyave.dev) | 852 | 39 | 7 | 🍐 Elegant Golang REST API Framework | 2019-10-21 09:44:34 | 2021-06-11 15:33:09 |
| [tango](https://github.com/lunny/tango) | 834 | 109 | 9 | This is only a mirror and Moved to https://gitea.com/lunny/tango | 2014-12-17 03:07:09 | 2021-05-31 10:00:01 |
| [aah](https://aahframework.org) | 652 | 35 | 16 | A secure, flexible, rapid Go web framework | 2016-06-27 04:47:45 | 2021-06-06 23:33:09 |
| [gearbox](https://gogearbox.com) | 510 | 37 | 2 | Gearbox :gear: is a web framework written in Go with a focus on high performance | 2020-04-25 01:28:37 | 2021-06-19 15:24:29 |
| [gongular](http://gondolaweb.com) | 443 | 15 | 6 | A different approach to Go web frameworks | 2016-06-22 11:52:42 | 2021-06-03 18:31:30 |
| [neo](http://ivpusic.github.io/neo/) | 410 | 40 | 7 | Go Web Framework | 2015-02-04 19:16:06 | 2021-06-13 15:37:14 |
| [air](https://pkg.go.dev/github.com/aofei/air) | 406 | 39 | 3 | An ideally refined web framework for Go. | 2016-07-20 12:09:48 | 2021-05-21 05:36:14 |
| [aero](https://github.com/aerogo/aero) | 377 | 21 | 4 | :bullettrain_side: High-performance web server for Go. | 2016-11-09 13:02:13 | 2021-06-10 21:12:21 |
| [mango](http://github.com/paulbellamy/mango) | 356 | 37 | 9 | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. | 2011-05-25 07:26:46 | 2021-06-16 14:00:09 |
| [gondola](http://gondolaweb.com) | 309 | 22 | 8 | The web framework for writing faster sites, faster | 2014-07-25 21:28:55 | 2021-05-05 01:20:17 |
| [confetti](https://www.confetti-framework.com) | 265 | 8 | 1 | Confetti is a web application framework with an expressive, elegant syntax. This repository contains configuration files and is intended as a template for your codebase. Download these configuration files and include them in your git repository. | 2019-11-01 23:14:21 | 2021-06-19 09:14:58 |
| [golf](https://golf.readme.io/) | 249 | 26 | 4 | :golf: The Golf web framework | 2015-11-18 15:10:14 | 2021-06-06 14:57:36 |
| [flamingo](http://www.flamingo.me) | 214 | 29 | 15 | Flamingo Framework and Core Library. Flamingo is a go based framework for pluggable web projects. It is used to build scalable and maintainable (web)applications. | 2019-04-02 12:24:02 | 2021-06-17 13:32:51 |
| [flamingo-commerce](https://www.flamingo.me/flamingo-commerce.html) | 184 | 30 | 17 | Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce "Portals" and connect it with the help of individual Adapters to other services.  | 2019-04-02 15:11:57 | 2021-06-19 20:03:40 |
| [beego](beego.me) | 183 | 51 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2021-06-20 01:35:13 |
| [ginrpc](https://xxjwxc.github.io/post/ginrpc/) | 173 | 19 | 6 | gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具 | 2019-06-22 12:03:53 | 2021-06-20 12:55:04 |
| [webgo](https://github.com/bnkamalesh/webgo) | 170 | 17 | 5 | A minimal framework to build web apps; with handler chaining, middleware support; and most of all standard library compliant HTTP handlers(i.e. http.HandlerFunc). | 2015-12-16 07:35:02 | 2021-05-31 04:36:02 |
| [hiboot](https://hiboot.netlify.app/) | 160 | 28 | 3 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2021-06-12 12:38:29 |
| [uadmin](https://uadmin.io) | 138 | 30 | 10 | The web framework for Golang | 2018-10-05 09:00:17 | 2021-06-09 07:45:51 |
| [go-rest](http://go.pkgdoc.org/github.com/ungerik/go-rest) | 125 | 13 | 2 | A small and evil REST framework for Go | 2012-07-13 10:02:15 | 2021-02-11 17:43:01 |
| [appy](https://github.com/appist/appy) | 92 | 10 | 14 | An opinionated productive web framework that helps scaling business easier. | 2019-05-27 04:48:59 | 2021-06-14 19:04:02 |
| [vox](https://aisk.github.io/vox/) | 75 | 4 | 7 | Simple and lightweight Go web framework inspired by koa | 2014-12-24 11:22:08 | 2021-05-13 00:36:37 |
| [microservice](http://github.com/paulbellamy/mango) | 74 | 13 | 0 | This library provides a simple framework of microservice, which includes a configurator, a logger, metrics, and of course the handler | 2016-12-15 09:07:04 | 2021-04-15 02:41:22 |
| [patron](https://github.com/beatlabs/patron) | 73 | 47 | 35 | Microservice framework following best cloud practices with a focus on productivity. | 2019-01-30 13:49:54 | 2021-06-16 14:24:49 |
| [golax](https://github.com/fulldump/golax) | 73 | 6 | 6 | Golax, a go implementation for the Lax framework. | 2016-01-30 19:11:39 | 2020-08-31 11:59:35 |
| [rux](https://pkg.go.dev/github.com/gookit/rux?tab=doc) | 62 | 10 | 2 | ⚡ Rux is an simple and fast web framework. support middleware, compatible http.Handler interface. 简单且快速的 Go web 框架，支持中间件，兼容 http.Handler 接口 | 2018-08-05 06:13:57 | 2021-05-24 02:06:39 |
| [yarf](https://github.com/yarf-framework/yarf) | 61 | 6 | 1 | Yet Another REST Framework | 2015-09-02 13:56:47 | 2021-04-16 08:52:07 |
| [fireball](https://github.com/zpatrick/fireball) | 56 | 4 | 0 | Go web framework with a natural feel | 2016-07-20 05:04:54 | 2021-04-06 16:51:53 |
| [goa](https://goa-go.github.io) | 45 | 1 | 0 | Goa is a  web framework based on middleware, like koa.js. | 2019-07-26 07:12:23 | 2021-06-06 10:44:15 |
| [gotuna](https://gotuna.org) | 34 | 4 | 0 | GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server. | 2021-04-08 14:08:08 | 2021-05-19 20:08:11 |
| [api](http://resoursea.com/) | 31 | 2 | 0 | A REST framework for quickly writing resource based services in Golang. | 2015-01-24 18:45:30 | 2020-02-21 12:56:46 |
| [rex](https://github.com/goanywhere/rex) | 31 | 1 | 0 | Pleasures for Web in Golang | 2014-10-16 02:26:18 | 2020-12-22 09:31:27 |
| [goweb](https://github.com/twharmon/goweb) | 24 | 1 | 0 | Lightweight web framework based on net/http. | 2019-05-07 21:04:43 | 2021-05-13 03:52:22 |
| [banjo](https://nsheremet.pw/banjo) | 17 | 5 | 1 | BANjO is a simple web framework written in Go (golang) | 2017-12-09 13:35:31 | 2020-12-24 11:45:51 |
</details>

### Web Frameworks - Middlewares - Actual middlewares
2021-06-21 02:25:21

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1,957 | 180 | 5 | Simple middleware to rate-limit HTTP requests. | 2015-05-17 15:20:03 | 2021-06-20 08:43:30 |
| [cors](https://github.com/rs/cors) | 1,843 | 157 | 17 | Go net/http configurable handler to handle CORS requests | 2014-10-25 03:49:45 | 2021-06-20 02:39:38 |
| [limiter](https://github.com/ulule/limiter) | 1,284 | 103 | 5 | Dead simple rate limit middleware for Go. | 2015-10-02 08:12:38 | 2021-06-20 08:43:28 |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 818 | 26 | 4 | Go (golang) library for creating and consuming HTTP Server-Timing headers | 2018-02-12 03:56:02 | 2021-06-18 21:33:12 |
| [go-fault](https://github.com/github/go-fault) | 395 | 16 | 1 | Fault injection library in Go using standard http middleware | 2020-05-14 16:13:17 | 2021-06-14 19:47:46 |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 111 | 7 | 17 | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ | 2018-06-29 21:51:00 | 2021-06-11 18:31:55 |
| [xff](https://github.com/sebest/xff) | 80 | 17 | 5 | A Golang Middleware to handle X-Forwarded-For Header | 2014-12-22 10:29:05 | 2021-05-07 13:13:36 |
| [formjson](https://github.com/rs/formjson) | 36 | 0 | 0 | Go net/http handler to transparently manage posted JSON | 2015-03-19 23:52:28 | 2020-09-18 01:35:42 |
| [client-timing](https://github.com/posener/client-timing) | 19 | 4 | 1 | An HTTP client for go-server-timing middleware. Enables automatic timing propagation through HTTP calls between servers. | 2018-02-23 01:52:45 | 2021-04-21 08:17:29 |
</details>

### Web Frameworks - Middlewares - Libraries for creating HTTP middlewares
2021-06-21 02:25:16

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [negroni](https://github.com/urfave/negroni) | 7,011 | 562 | 9 | Idiomatic HTTP Middleware for Golang | 2014-05-18 22:09:10 | 2021-06-20 11:03:52 |
| [alice](https://godoc.org/github.com/justinas/alice) | 2,281 | 128 | 6 | Painless middleware chaining for Go | 2014-05-25 07:27:41 | 2021-06-20 05:21:39 |
| [render](https://github.com/unrolled/render) | 1,512 | 125 | 1 | Go package for easily rendering JSON, XML, binary data, and HTML templates responses. | 2014-06-10 16:20:35 | 2021-06-19 17:38:44 |
| [stats](https://github.com/thoas/stats) | 579 | 47 | 8 | A Go middleware that stores various information about your web application (response time, status code count, etc.) | 2015-03-05 18:02:50 | 2021-04-10 06:38:49 |
| [interpose](https://github.com/carbocation/interpose) | 291 | 15 | 1 | Minimalist net/http middleware for golang | 2014-07-20 00:19:52 | 2021-06-10 11:51:50 |
| [renderer](https://github.com/thedevsaddam/renderer) | 222 | 23 | 0 | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go | 2017-11-07 18:53:49 | 2021-06-18 05:42:03 |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 13 | 1 | Lightweight Middleware for net/http | 2014-05-03 17:14:17 | 2020-09-04 11:26:14 |
| [rye](https://github.com/InVisionApp/rye) | 97 | 12 | 0 | A tiny http middleware for Golang with added handlers for common needs. | 2016-10-06 19:51:59 | 2021-06-06 09:24:36 |
| [gores](https://github.com/alioygur/gores) | 95 | 2 | 0 | Go package that handles HTML, JSON, XML and etc. responses | 2015-12-25 12:41:01 | 2021-06-16 05:58:36 |
| [mediary](https://github.com/HereMobilityDevelopers/mediary/wiki/Reasoning) | 71 | 4 | 0 | Add interceptors to GO http.Client | 2020-03-23 18:54:56 | 2021-02-10 17:40:20 |
| [chain](https://github.com/codemodus/chain) | 63 | 3 | 0 | Composable chains of nested http.Handler instances. | 2015-05-14 19:52:58 | 2021-02-11 17:26:39 |
| [wrap](https://github.com/go-on/wrap) | 59 | 5 | 0 | Go http.Hander based middleware stack with context sharing | 2014-02-16 07:12:36 | 2020-08-28 05:29:07 |
| [catena](https://github.com/codemodus/catena) | 7 | 0 | 0 | gRPC interceptor catenation. | 2015-07-30 19:07:01 | 2018-08-25 22:06:07 |
</details>

### Web Frameworks - Routers
2021-06-21 02:25:04

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mux](http://www.gorillatoolkit.org/pkg/mux) | 14,563 | 1,361 | 24 | A powerful HTTP router and URL matcher for building Go web servers with 🦍 | 2012-10-02 21:32:24 | 2021-06-20 10:39:41 |
| [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) | 12,817 | 1,240 | 60 | A high performance HTTP request router that scales well | 2013-12-05 15:10:55 | 2021-06-20 08:43:10 |
| [chi](https://github.com/go-chi/chi) | 9,570 | 645 | 19 | lightweight, idiomatic and composable router for building Go HTTP services | 2015-10-15 20:46:29 | 2021-06-20 11:14:39 |
| [web](https://github.com/gocraft/web) | 1,444 | 121 | 23 | Go Router + Middleware. Your Contexts. | 2013-11-16 20:48:20 | 2021-06-16 01:03:43 |
| [bone](http://go-zoo.github.io/bone) | 1,272 | 85 | 2 | Lightning Fast HTTP Multiplexer | 2014-11-19 02:16:36 | 2021-06-15 07:16:47 |
| [fasthttprouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 873 | 92 | 19 | A high performance fasthttp request router that scales well | 2015-12-13 09:32:30 | 2021-06-17 13:16:03 |
| [goji](https://goji.io) | 862 | 61 | 6 | Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) | 2015-11-16 00:52:41 | 2021-06-12 01:48:49 |
| [gorouter](https://xujiajun.cn/gorouter) | 498 | 80 | 1 | xujiajun/gorouter is a simple and fast HTTP router for Go. It is easy to build RESTful APIs and your web framework. | 2018-01-29 09:28:28 | 2021-06-15 05:49:27 |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 491 | 44 | 2 | High-speed, flexible tree-based HTTP router for Go. | 2014-05-14 20:10:20 | 2021-06-16 08:05:51 |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 406 | 50 | 11 | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. | 2015-10-27 01:03:14 | 2021-05-27 09:59:27 |
| [lars](https://github.com/go-playground/lars) | 382 | 23 | 1 | :rotating_light: Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. | 2015-12-24 17:28:45 | 2021-05-06 17:20:44 |
| [siesta](https://github.com/VividCortex/siesta) | 350 | 14 | 0 | Composable framework for writing HTTP handlers in Go. | 2014-09-23 13:55:56 | 2021-05-23 22:02:49 |
| [vestigo](https://github.com/husobee/vestigo) | 262 | 29 | 14 | Echo Inspired Stand Alone URL Router | 2015-09-22 03:08:03 | 2021-05-16 04:33:49 |
| [router](https://github.com/gowww/router) | 160 | 12 | 0 | ⚡️ A lightning fast HTTP router | 2017-05-25 10:29:27 | 2021-04-26 14:57:04 |
| [alien](https://github.com/gernest/alien) | 117 | 10 | 3 | A lightweight and  fast http router from outer space | 2016-01-30 23:23:10 | 2021-04-22 20:03:07 |
| [pure](https://github.com/go-playground/pure) | 116 | 9 | 0 | :non-potable_water: Is a lightweight  HTTP router that sticks to the std "net/http" implementation | 2016-09-23 19:57:58 | 2021-05-08 00:40:18 |
| [violetear](http://violetear.org) | 103 | 8 | 1 | Go HTTP router | 2015-06-19 16:49:41 | 2021-05-25 19:15:16 |
| [Bxog](http://go-zoo.github.io/bone) | 102 | 6 | 0 | Bxog is a simple and fast HTTP router for Go (HTTP request multiplexer). | 2016-05-19 12:20:08 | 2021-06-15 07:22:50 |
| [gorouter](https://rafallorenz.com/gorouter) | 95 | 14 | 1 | Go Server/API micro framework, HTTP request router, multiplexer, mux | 2016-07-14 13:13:34 | 2021-06-17 07:32:30 |
| [xmux](http://violetear.org) | 90 | 9 | 2 | xmux is a httprouter fork on top of xhandler (net/context aware) | 2015-12-14 19:01:05 | 2021-06-13 02:37:37 |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 49 | 4 | 0 | :bell: A simple Go router | 2019-02-21 13:13:52 | 2021-04-21 03:20:58 |
| [fastrouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 19 | 3 | 0 | FastRouter is a fast, flexible HTTP router written in Go. | 2017-11-01 08:52:52 | 2020-12-30 15:49:44 |
| [route](https://goroute.github.io) | 7 | 1 | 1 | Go Route - Simple yet powerful HTTP request multiplexer | 2019-07-06 18:47:38 | 2020-08-31 13:36:03 |
</details>

### WebAssembly
2021-06-21 01:48:40

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tinygo](https://tinygo.org) | 8,150 | 424 | 264 | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. | 2018-06-07 16:39:19 | 2021-06-19 16:58:29 |
| [dom](https://github.com/dennwc/dom) | 428 | 51 | 11 | DOM library for Go and WASM | 2018-06-30 18:37:35 | 2021-06-14 22:16:09 |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 135 | 9 | 5 | Library to use HTML5 Canvas  from Go-WASM, with all drawing within go code | 2019-05-05 14:05:55 | 2021-06-14 21:30:05 |
| [webapi](https://github.com/gowebapi/webapi) | 82 | 8 | 2 | Go Lang Web Assembly bindings for DOM, HTML etc | 2019-02-08 05:58:35 | 2021-06-16 15:50:16 |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 80 | 9 | 0 | Run WASM tests inside your browser | 2018-07-14 18:42:24 | 2021-04-07 05:36:46 |
| [vert](https://github.com/norunners/vert) | 48 | 6 | 2 | WebAssembly interop between Go and JS values. | 2018-03-25 17:26:47 | 2021-05-23 19:32:09 |
</details>

### Windows
2021-06-21 01:48:39

*Last Update: *

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-ole](https://github.com/go-ole/go-ole) | 758 | 139 | 53 | win32 ole implementation for golang | 2011-01-21 12:45:20 | 2021-06-15 07:38:41 |
| [d3d9](https://github.com/gonutz/d3d9) | 122 | 9 | 0 | Direct3D9 wrapper for Go. | 2015-12-12 21:24:38 | 2021-06-19 22:32:48 |
| [gosddl](https://github.com/MonaxGT/gosddl) | 5 | 1 | 0 | GoSDDL converter | 2018-12-04 08:36:11 | 2021-02-14 13:03:11 |
</details>

### XML
2021-06-21 02:25:01

*Last Update: Libraries and tools for manipulating XML.*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [zek](https://github.com/miku/zek) | 466 | 40 | 8 | Generate a Go struct from XML. | 2017-11-23 19:03:11 | 2021-06-18 10:38:06 |
| [xpath](https://github.com/antchfx/xpath) | 398 | 54 | 9 | XPath package for Golang, supports HTML, XML, JSON document query. | 2016-10-09 05:51:24 | 2021-06-20 08:32:17 |
| [xquery](https://github.com/antchfx/xpath) | 153 | 25 | 0 | Extract data or evaluate value from HTML/XML documents using XPath | 2016-10-09 05:54:10 | 2021-05-22 09:57:43 |
| [xml2map](https://github.com/sbabiv/xml2map) | 32 | 5 | 1 | XML to MAP converter written Golang | 2018-08-06 17:51:46 | 2021-06-09 18:08:01 |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 19 | 3 | 1 | xmlwriter is a pure-Go library providing procedural XML generation based on libxml2's xmlwriter module | 2017-04-11 04:43:26 | 2021-05-26 06:49:04 |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 8 | 8 | Compare ANY markup documents. | 2016-10-25 22:09:12 | 2021-04-23 04:41:47 |
</details>

