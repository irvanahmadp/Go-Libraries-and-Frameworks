# Go Libraries and Frameworks
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
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
* [GUI - Interaction](#gui---interaction)
* [GUI - Toolkits](#gui---toolkits)
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
Libraries for manipulating audio.

<sup>*Last Update: 2023-10-31 07:51:34*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oto](https://github.com/ebitengine/oto) | 1,372 | 122 | 15 | ♪ A low-level library to play sound on multiple platforms ♪ | 2017-05-04 12:16:30 | 2023-10-28 04:06:11 |
| [portaudio](https://github.com/gordonklaus/portaudio) | 622 | 89 | 3 | Go bindings for the PortAudio audio I/O library | 2015-09-16 07:59:23 | 2023-10-31 00:40:32 |
| [music-theory](https://gopkg.in/music-theory.v0) | 417 | 43 | 8 | Go models of Note, Scale, Chord and Key | 2016-03-17 03:50:17 | 2023-10-20 15:08:15 |
| [waveform](https://github.com/mdlayher/waveform) | 352 | 29 | 4 | Go package capable of generating waveform images from audio streams. MIT Licensed. | 2014-09-13 18:07:36 | 2023-10-19 10:26:45 |
| [id3v2](https://pkg.go.dev/github.com/bogem/id3v2/v2) | 304 | 50 | 17 | 🎵 ID3 decoding and encoding library for Go | 2016-05-15 18:36:53 | 2023-10-20 23:04:24 |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 301 | 36 | 7 | Go tools for audio processing & creation 🎶 | 2020-07-05 01:35:41 | 2023-10-28 04:05:59 |
| [portmidi](https://github.com/rakyll/portmidi) | 280 | 60 | 14 | Go bindings for libportmidi | 2013-11-10 14:24:56 | 2023-10-10 16:56:09 |
| [flac](https://github.com/mewkiz/flac) | 270 | 40 | 10 | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 2012-11-01 20:13:58 | 2023-10-31 00:17:23 |
| [malgo](https://pkg.go.dev/github.com/bogem/id3v2/v2) | 234 | 43 | 8 | Mini audio library | 2017-11-09 18:27:52 | 2023-10-29 00:03:53 |
| [mix](https://gopkg.in/mix.v0) | 171 | 26 | 8 | Sequence-based Go-native audio mixer for music apps | 2016-01-03 15:53:57 | 2023-10-28 04:11:08 |
| [gaad](https://github.com/Comcast/gaad) | 113 | 22 | 2 | GAAD (Go Advanced Audio Decoder) | 2016-07-11 14:19:16 | 2023-09-18 10:34:35 |
| [minimp3](https://github.com/tosone/minimp3) | 111 | 18 | 0 | Decode mp3 base on https://github.com/lieff/minimp3 | 2018-01-26 09:10:31 | 2023-09-16 05:43:04 |
| [vorbis](https://github.com/mccoyst/vorbis) | 32 | 6 | 3 | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) | 2013-07-12 02:45:39 | 2022-10-24 08:41:49 |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 29 | 11 | 0 | Go Bindings for libsamplerate | 2016-11-20 21:19:31 | 2023-09-19 03:20:31 |
</details>

### Authentication and OAuth
Libraries for implementing authentications schemes.

<sup>*Last Update: 2023-11-13 21:48:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [casbin](https://casbin.org) | 15,979 | 1,625 | 40 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang: https://discord.gg/S5UjpzGZjN | 2017-04-08 07:51:23 | 2023-11-08 12:41:55 |
| [oauth2](https://golang.org/x/oauth2) | 4,938 | 1,014 | 178 | Go OAuth2 | 2014-04-14 15:07:35 | 2023-11-08 01:59:46 |
| [goth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 4,391 | 553 | 95 | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 2014-10-14 20:38:12 | 2023-11-07 14:23:37 |
| [authboss](https://github.com/volatiletech/authboss) | 3,524 | 205 | 40 | The boss of http auth. | 2015-01-03 05:12:02 | 2023-11-08 13:17:55 |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 2,052 | 310 | 29 | A standalone, specification-compliant,  OAuth2 server written in Golang. | 2015-11-01 13:30:09 | 2023-11-07 01:49:47 |
| [go-jose](https://github.com/square/go-jose) | 1,983 | 417 | 0 | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 2014-11-14 18:27:31 | 2023-10-27 23:32:43 |
| [loginsrv](https://github.com/tarent/loginsrv) | 1,906 | 156 | 28 | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 2016-11-11 12:11:21 | 2023-11-01 14:45:50 |
| [osin](https://golang.org/x/oauth2) | 1,845 | 404 | 8 | Golang OAuth2 server library | 2013-09-10 19:52:00 | 2023-11-06 21:11:01 |
| [scs](https://github.com/alexedwards/scs) | 1,666 | 184 | 7 | HTTP Session Management for Go | 2016-08-08 16:42:05 | 2023-11-08 13:24:54 |
| [gologin](https://github.com/dghubble/gologin) | 1,645 | 166 | 0 | Go login handlers for authentication providers (OAuth1, OAuth2) | 2015-06-23 04:40:52 | 2023-11-07 10:48:11 |
| [gorbac](https://github.com/mikespook/gorbac) | 1,509 | 221 | 4 | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 2013-12-26 10:00:41 | 2023-10-29 12:38:30 |
| [paseto](https://github.com/o1egl/paseto) | 766 | 36 | 6 | Platform-Agnostic Security Tokens implementation in GO (Golang) | 2018-01-23 05:27:39 | 2023-11-08 04:52:34 |
| [jwt](https://github.com/cristalhq/jwt) | 625 | 42 | 3 | Safe, simple and fast JSON Web Tokens for Go | 2019-07-20 18:14:58 | 2023-10-30 15:28:09 |
| [go-guardian](https://github.com/shaj13/go-guardian) | 499 | 52 | 9 | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 2020-05-14 12:15:56 | 2023-11-02 12:14:49 |
| [permissions2](https://github.com/xyproto/permissions2) | 495 | 38 | 1 |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 2014-11-19 12:23:37 | 2023-11-07 12:38:40 |
| [jwt](https://github.com/pascaldekloe/jwt) | 340 | 25 | 0 | JSON Web Token library | 2018-03-21 11:59:53 | 2023-11-01 22:40:54 |
| [jeff](https://github.com/abraithwaite/jeff) | 258 | 17 | 3 | 🍍Jeff provides the simplest way to manage web sessions in Go. | 2018-08-02 19:31:23 | 2023-10-02 15:07:56 |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 230 | 42 | 4 | This package provides json web token (jwt) middleware for goLang http servers | 2016-07-05 23:31:43 | 2023-10-17 14:58:44 |
| [httpauth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 219 | 30 | 5 | HTTP Authentication middlewares | 2014-05-26 22:53:57 | 2023-04-30 05:53:08 |
| [branca](https://branca.io) | 171 | 20 | 0 | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. | 2018-01-09 15:27:31 | 2023-11-03 07:21:19 |
| [sessionup](https://github.com/swithek/sessionup) | 122 | 7 | 3 | Straightforward HTTP session management | 2019-07-23 18:55:21 | 2023-09-18 22:06:41 |
| [sjwt](https://godoc.org/github.com/brianvoe/sjwt) | 113 | 10 | 1 | Simple JWT Golang | 2019-06-20 04:06:21 | 2023-08-24 12:52:50 |
| [session](https://github.com/icza/session) | 111 | 17 | 4 | Go session management for web servers (including support for Google App Engine - GAE). | 2016-02-08 09:07:07 | 2023-01-22 03:03:17 |
| [rbac](https://github.com/zpatrick/rbac) | 107 | 25 | 0 | Minimalistic RBAC package for Go applications | 2018-08-02 00:11:04 | 2023-06-20 17:52:41 |
| [jwt](https://github.com/robbert229/jwt) | 104 | 28 | 8 | This is an implementation of JWT in golang! | 2016-06-05 22:01:37 | 2023-09-03 17:53:11 |
| [sessions](https://github.com/adam-hanna/sessions) | 75 | 11 | 4 | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 2017-04-29 01:09:28 | 2023-08-07 04:30:13 |
| [securecookie](https://github.com/chmike/securecookie) | 74 | 11 | 1 | Fast, secure and efficient secure cookie encoder/decoder  | 2017-09-03 14:40:29 | 2023-08-26 00:38:03 |
| [otpgo](https://github.com/jltorresm/otpgo) | 61 | 9 | 2 | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 2020-08-19 13:20:23 | 2023-11-07 12:55:26 |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 60 | 7 | 0 | Golang library for providing a canonical representation of email address. | 2020-08-21 23:13:04 | 2023-10-24 19:08:54 |
| [scope](https://github.com/ThundR67/scope) | 33 | 7 | 1 | Easily Manage OAuth2 Scopes In Go | 2019-09-23 10:48:14 | 2023-10-07 21:10:33 |
| [cookiestxt](https://casbin.org) | 12 | 5 | 1 | cookiestxt implement parser of cookies txt format | 2017-10-09 11:27:19 | 2023-04-05 17:00:07 |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 2 | 0 | A driver for the SessionGate Redis module - easy session management using the Go language. | 2017-10-20 03:39:11 | 2023-08-19 15:44:39 |
| [casbin](https://casbin.org/) | 1 | 0 | 0 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2021-05-29 04:09:46 | 2021-06-23 04:53:51 |
</details>

### Bot Building
Libraries for building and working with bots.

<sup>*Last Update: 2023-11-03 08:18:05*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [telegram-bot-api](https://go-telegram-bot-api.dev) | 5,047 | 774 | 115 | Golang bindings for the Telegram Bot API | 2015-06-25 05:33:57 | 2023-11-02 08:49:16 |
| [olivia](https://olivia-ai.org) | 3,612 | 357 | 28 | 💁‍♀️Your new best friend powered by an artificial neural network | 2018-06-05 18:19:31 | 2023-11-02 07:23:19 |
| [telebot](https://github.com/tucnak/telebot) | 3,314 | 401 | 49 | Telebot is a Telegram bot framework in Go. | 2015-06-25 19:27:50 | 2023-11-02 19:04:37 |
| [kelp](https://kelpbot.io) | 1,030 | 239 | 176 | Kelp is a free and open-source trading bot for the Stellar DEX and 100+ centralized exchanges | 2018-08-08 23:31:18 | 2023-11-02 07:34:35 |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 943 | 239 | 6 | A golang implementation of a console-based trading bot for cryptocurrency exchanges | 2017-05-14 22:11:41 | 2023-11-02 15:20:50 |
| [slacker](https://github.com/shomali11/slacker) | 805 | 122 | 7 | Slack Bot Framework | 2017-05-20 01:41:20 | 2023-10-30 02:27:36 |
| [bot](https://github.com/go-chat-bot/bot) | 801 | 214 | 15 | IRC, Slack, Telegram and RocketChat bot written in go | 2015-09-22 16:41:13 | 2023-11-02 02:19:13 |
| [joe](https://joe-bot.net) | 468 | 26 | 6 | A general-purpose bot library inspired by Hubot but written in Go.   :robot: | 2019-03-03 11:19:18 | 2023-11-01 11:31:04 |
| [tbot](https://yanzay.github.io/tbot-doc/) | 342 | 53 | 1 | Go library for Telegram Bot API | 2015-09-11 16:19:25 | 2023-10-26 23:18:19 |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 309 | 55 | 7 | go irc client for twitch.tv | 2017-03-23 21:31:35 | 2023-10-29 12:15:25 |
| [echotron](https://t.me/s/echotronnews) | 284 | 21 | 0 | An elegant and concurrent library for the Telegram bot API in Go. | 2019-07-22 17:31:49 | 2023-11-01 04:54:18 |
| [go-sarah](https://pkg.go.dev/github.com/oklahomer/go-sarah/v4) | 260 | 17 | 0 | Simple yet customizable bot framework written in Go. | 2016-11-06 10:04:43 | 2023-11-01 14:48:25 |
| [tenyks](http://tenyks.io) | 176 | 19 | 12 | The Tenyks IRC bot. | 2012-08-26 02:02:24 | 2023-09-28 17:37:33 |
| [hanu](https://sbstjn.com/host-golang-slackbot-on-heroku-with-hanu.html) | 150 | 24 | 3 | Golang Framework for writing Slack bots | 2016-09-16 07:10:42 | 2023-10-24 00:43:33 |
| [slack-bot](https://github.com/innogames/slack-bot) | 149 | 44 | 8 | Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests with AI support... | 2019-07-19 07:49:06 | 2023-10-05 02:19:16 |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 119 | 4 | 2 | Golang  telegram bot API wrapper, session-based router and middleware | 2016-12-11 06:06:32 | 2023-10-19 10:27:28 |
| [margelet](https://github.com/zhulik/margelet) | 81 | 15 | 2 | Telegram Bot Framework for Go | 2015-11-21 13:02:17 | 2023-09-18 15:27:27 |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 81 | 10 | 6 | A Discord bot for managing ephemeral roles based upon voice channel member presence. | 2017-12-19 15:20:30 | 2023-09-25 20:00:44 |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 56 | 10 | 3 | Slack bot core/framework written in Go with support for reactions to message updates/deletes | 2015-10-22 04:54:55 | 2023-09-18 15:27:30 |
| [govkbot](https://github.com/nikepan/govkbot) | 45 | 6 | 1 | VK bot package for Go | 2016-07-11 22:09:54 | 2023-10-24 12:10:37 |
| [micha](https://github.com/onrik/micha) | 26 | 5 | 1 | Client lib for Telegram bot api | 2016-04-14 12:09:44 | 2023-09-18 15:27:25 |
</details>

### Build Automation
Libraries and tools helping with build automation.

<sup>*Last Update: 2023-12-16 21:00:53*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [task](https://taskfile.dev) | 9,106 | 546 | 232 | A task runner / simpler Make alternative written in Go | 2017-02-27 00:46:04 | 2023-12-16 03:41:03 |
| [realize](https://github.com/oxequa/realize) | 4,433 | 236 | 71 | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 2016-07-12 08:07:25 | 2023-12-14 06:18:25 |
| [mmake](https://pkg.go.dev/github.com/goyek/goyek/v2) | 1,692 | 47 | 11 | Modern Make  | 2017-02-15 22:01:21 | 2023-12-11 17:49:10 |
| [goyek](https://pkg.go.dev/github.com/goyek/goyek/v2) | 407 | 28 | 2 | Build automation Go library. | 2020-10-11 13:20:55 | 2023-12-16 09:26:58 |
| [taskctl](https://github.com/taskctl/taskctl) | 270 | 38 | 12 | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 2019-11-12 13:19:09 | 2023-11-08 05:59:00 |
| [1build](https://1build.gitbook.io) | 219 | 31 | 33 | Frictionless way of managing project-specific commands | 2019-04-23 17:05:38 | 2023-12-15 01:52:20 |
| [gilbert](https://go-gilbert.github.io/) | 113 | 7 | 0 | Build system and task runner for Go projects | 2019-01-30 09:02:31 | 2023-11-11 08:04:48 |
| [gaper](https://github.com/maxcnunes/gaper) | 76 | 6 | 3 | Builds and restarts a Go project when it crashes or some watched file changes | 2018-06-16 02:46:38 | 2023-12-01 11:44:43 |
| [anko](https://github.com/GuilhermeCaruso/anko) | 33 | 2 | 0 | :crystal_ball: Simple application watcher | 2021-03-02 14:08:42 | 2023-09-27 02:00:22 |
</details>

### CSS Preprocessors
Libraries for preprocessing CSS files.

<sup>*Last Update: 2023-11-20 08:20:35*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gcss](https://github.com/yosssi/gcss) | 488 | 39 | 8 | Pure Go CSS Preprocessor | 2014-09-04 14:38:20 | 2023-11-11 09:14:55 |
| [go-libsass](http://godoc.org/github.com/wellington/go-libsass) | 199 | 27 | 16 | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 2015-04-19 15:09:47 | 2023-09-26 23:10:14 |
</details>

### Command Line - Advanced Console UIs
Libraries for building Console Applications and Console User Interfaces.

<sup>*Last Update: 2023-11-13 21:48:26*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [termui](https://github.com/gizak/termui) | 12,697 | 786 | 98 | Golang terminal dashboard | 2015-02-03 14:09:27 | 2023-11-13 09:20:43 |
| [gocui](https://github.com/jroimartin/gocui) | 9,448 | 646 | 74 | Minimalist Go package aimed at creating Console User Interfaces. | 2014-01-04 02:50:20 | 2023-11-13 13:00:29 |
| [go-prompt](https://godoc.org/github.com/c-bata/go-prompt) | 5,050 | 389 | 108 | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. | 2017-08-14 16:02:09 | 2023-11-12 15:26:59 |
| [termbox-go](http://godoc.org/github.com/nsf/termbox-go) | 4,582 | 411 | 45 | Pure Go termbox implementation | 2012-01-12 21:03:03 | 2023-11-12 21:47:54 |
| [pterm](https://docs.pterm.sh) | 4,297 | 166 | 48 | ✨ #PTerm is a modern Go module to easily beautify console output. Featuring charts, progressbars, tables, trees, text input, select menus and much more 🚀 It's completely configurable and 100% cross-platform compatible. | 2020-09-17 15:52:59 | 2023-11-13 12:06:04 |
| [progressbar](https://pkg.go.dev/github.com/schollz/progressbar/v3?tab=doc) | 3,578 | 207 | 33 | A really basic thread-safe progress bar for Golang applications | 2017-10-26 18:28:10 | 2023-11-13 14:14:30 |
| [termdash](http://godoc.org/github.com/nsf/termbox-go) | 2,384 | 119 | 43 | Terminal based dashboard. | 2018-03-24 12:01:49 | 2023-11-12 22:01:57 |
| [asciigraph](https://pkg.go.dev/github.com/guptarohit/asciigraph) | 2,342 | 94 | 12 | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. | 2018-06-17 10:37:16 | 2023-11-11 02:53:43 |
| [mpb](https://github.com/vbauerster/mpb) | 2,123 | 124 | 10 | multi progress bar for Go cli applications | 2016-12-14 11:56:29 | 2023-11-13 13:01:04 |
| [uiprogress](https://github.com/gosuri/uiprogress) | 2,048 | 125 | 25 | A go library to render progress bars in terminal applications | 2015-11-17 00:59:24 | 2023-11-11 02:36:18 |
| [uilive](https://github.com/gosuri/uilive) | 1,624 | 87 | 11 | uilive is a go library for updating terminal output in realtime | 2015-11-16 06:13:10 | 2023-11-11 12:23:47 |
| [color](https://gookit.github.io/color/) | 1,393 | 84 | 5 | 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染 | 2018-07-01 07:28:17 | 2023-11-11 07:32:01 |
| [aurora](https://pkg.go.dev/github.com/guptarohit/asciigraph) | 1,355 | 61 | 1 | Golang ultimate ANSI-colors that supports Printf/Sprintf methods | 2016-11-06 22:37:12 | 2023-11-11 22:43:05 |
| [go-isatty](http://godoc.org/github.com/mattn/go-isatty) | 754 | 106 | 13 | Change the color of console text. | 2014-04-01 01:53:09 | 2023-11-09 11:21:54 |
| [go-colorable](http://godoc.org/github.com/mattn/go-colorable) | 717 | 94 | 9 | Another Text Attribute Manupulator | 2014-07-30 02:38:06 | 2023-11-12 10:21:52 |
| [uitable](https://github.com/gosuri/uitable) | 716 | 33 | 8 | A go library to improve readability in terminal apps using tabular data | 2015-11-13 21:59:21 | 2023-11-12 22:08:30 |
| [gommon](https://github.com/labstack/gommon) | 508 | 131 | 13 | Common packages for Go | 2015-03-12 22:35:57 | 2023-11-12 14:14:21 |
| [simpletable](https://github.com/alexeyco/simpletable) | 485 | 30 | 5 | Simple tables in terminal with Go | 2017-03-29 07:27:23 | 2023-11-12 11:17:12 |
| [chalk](https://github.com/ttacon/chalk) | 434 | 20 | 3 | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk | 2014-07-18 19:38:58 | 2023-11-08 21:44:05 |
| [yacspin](https://github.com/theckman/yacspin) | 418 | 14 | 6 | Yet Another CLi Spinner; providing over 80 easy to use and customizable terminal spinners for multiple OSes | 2019-12-29 07:41:23 | 2023-11-11 23:55:14 |
| [box-cli-maker](https://github.com/Delta456/box-cli-maker) | 396 | 20 | 8 | Make Highly Customized Boxes for CLI | 2020-05-01 07:23:56 | 2023-11-07 16:54:25 |
| [tabby](https://github.com/cheynewallace/tabby) | 338 | 18 | 3 | A tiny library for super simple Golang tables | 2018-12-17 23:35:39 | 2023-11-12 21:26:31 |
| [go-colortext](http://godoc.org/github.com/mattn/go-colorable) | 214 | 22 | 3 | Change the color of console text. | 2013-01-23 03:38:54 | 2023-07-30 06:16:45 |
| [cfmt](https://github.com/mingrammer/cfmt) | 97 | 7 | 1 | :art: Contextual fmt inspired by bootstrap color classes | 2018-03-15 19:04:27 | 2023-11-08 21:43:36 |
| [tabular](https://github.com/InVisionApp/tabular) | 74 | 7 | 0 | Tabular simplifies printing ASCII tables from command line utilities | 2018-04-23 21:17:03 | 2023-11-06 21:17:50 |
| [cfmt](https://github.com/i582/cfmt) | 60 | 3 | 0 | Small library for simple and convenient formatted stylized output to the console. | 2020-11-13 20:29:45 | 2023-10-19 04:14:32 |
| [table](https://github.com/tomlazar/table) | 46 | 3 | 1 | pretty colorfull tables in go with less effort | 2020-09-22 05:42:34 | 2023-10-12 01:09:02 |
| [ctc](https://github.com/wzshiming/ctc) | 45 | 3 | 1 | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method | 2018-04-27 18:07:42 | 2023-10-30 13:52:20 |
| [marker](https://github.com/cyucelen/marker) | 43 | 13 | 4 |  🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs! | 2019-08-28 15:44:08 | 2023-11-08 21:45:10 |
| [colourize](https://github.com/TreyBastian/colourize) | 26 | 5 | 0 | An ANSI colour terminal package for Go | 2015-05-11 11:49:39 | 2023-02-20 07:52:41 |
| [go-ataman](https://github.com/workanator/go-ataman) | 16 | 3 | 0 | Another Text Attribute Manupulator | 2017-05-17 19:04:57 | 2023-09-27 02:03:14 |
</details>

### Command Line - Standard CLI
Libraries for building standard or basic Command Line applications.

<sup>*Last Update: 2023-11-22 20:35:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cobra](https://cobra.dev) | 34,202 | 2,806 | 242 | A Commander for modern Go CLI interactions | 2013-09-03 20:40:26 | 2023-11-22 13:03:59 |
| [cli](https://cli.urfave.org) | 20,966 | 1,698 | 49 | A simple, fast, and fun package for building command line apps in Go | 2013-07-13 19:32:06 | 2023-11-22 08:55:29 |
| [kingpin](https://github.com/alecthomas/kingpin) | 3,386 | 275 | 24 | CONTRIBUTIONS ONLY: A Go (golang) command line and flag parser | 2014-05-14 20:09:04 | 2023-11-17 09:40:13 |
| [dnote](https://www.getdnote.com) | 2,601 | 113 | 55 | A simple command line notebook for programmers | 2017-03-30 23:07:25 | 2023-11-21 15:16:09 |
| [go-flags](http://godoc.org/github.com/jessevdk/go-flags) | 2,465 | 321 | 56 | go command line option parser | 2012-08-31 13:57:58 | 2023-11-22 11:01:55 |
| [pflag](https://ops.city) | 2,223 | 335 | 162 | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. | 2013-08-30 14:53:31 | 2023-11-20 21:02:54 |
| [go-arg](https://github.com/alexflint/go-arg) | 1,746 | 94 | 27 | Struct-based argument parsing in Go | 2015-11-01 01:30:06 | 2023-11-21 20:21:58 |
| [cli](https://github.com/mitchellh/cli) | 1,692 | 124 | 7 | A Go library for implementing command-line interfaces. | 2013-11-03 06:47:54 | 2023-11-21 06:18:23 |
| [ops](https://ops.city) | 1,140 | 124 | 144 | ops - build and run nanos unikernels | 2018-09-10 17:57:47 | 2023-11-20 18:56:14 |
| [liner](https://github.com/peterh/liner) | 1,012 | 135 | 17 | Pure Go line editor with history, inspired by linenoise | 2012-08-15 16:34:55 | 2023-11-20 15:43:18 |
| [complete](https://github.com/posener/complete) | 900 | 113 | 26 | bash completion written in go + bash completion for go command | 2017-05-05 21:34:07 | 2023-11-17 12:01:22 |
| [mow.cli](https://github.com/jawher/mow.cli) | 864 | 55 | 34 | A versatile library for building CLI applications in Go | 2014-12-18 19:34:20 | 2023-11-22 09:25:46 |
| [flaggy](https://github.com/integrii/flaggy) | 840 | 34 | 18 | Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies. | 2018-03-05 05:55:05 | 2023-11-17 15:17:29 |
| [cli](https://github.com/mkideal/cli) | 702 | 45 | 3 | CLI - A package for building command line app with go | 2016-02-26 16:45:29 | 2023-11-21 10:00:59 |
| [argparse](https://github.com/akamensky/argparse) | 565 | 57 | 6 | Argparse for golang. Just because `flag` sucks | 2017-11-24 06:42:20 | 2023-11-22 04:22:03 |
| [climax](http://git.io/climax) | 211 | 16 | 7 | Climax is an alternative CLI with the human face | 2015-11-03 21:04:57 | 2023-11-06 22:35:52 |
| [wmenu](https://github.com/dixonwille/wmenu) | 197 | 23 | 1 | An easy to use menu structure for cli applications that prompts users to make choices. | 2016-04-20 13:09:44 | 2023-11-21 05:15:59 |
| [hiboot](https://hiboot.netlify.app/) | 179 | 26 | 6 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2023-10-07 13:11:29 |
| [commandeer](https://github.com/jaffee/commandeer) | 172 | 16 | 5 | Automatically sets up command line flags based on struct fields and tags. | 2017-10-12 02:51:05 | 2023-10-07 13:09:59 |
| [clir](https://github.com/leaanthony/clir) | 160 | 15 | 6 | A Simple and Clear CLI library. Dependency free. | 2019-11-18 19:52:00 | 2023-11-02 12:30:17 |
| [sflags](https://ops.city) | 148 | 36 | 10 | Generate flags by parsing structures | 2016-12-04 14:49:27 | 2023-11-13 19:38:15 |
| [job](https://github.com/liujianping/job) | 134 | 12 | 1 | JOB, make your short-term command as a long-term job. 将命令行规划成任务的工具 | 2019-04-09 11:14:51 | 2023-11-09 12:09:32 |
| [flag](https://github.com/cosiner/flag) | 129 | 7 | 1 | Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand | 2016-10-05 16:49:41 | 2023-10-16 14:21:25 |
| [clif](https://github.com/ukautz/clif) | 125 | 14 | 3 | Another CLI framework for Go. It works on my machine. | 2015-05-30 18:30:08 | 2023-10-07 13:13:10 |
| [cmdr](https://hedzr.com/cmdr-docs/) | 124 | 9 | 1 | POSIX-compliant command-line UI (CLI) parser and Hierarchical-configuration operations | 2019-05-15 09:58:02 | 2023-10-15 01:50:47 |
| [cli](https://github.com/teris-io/cli) | 122 | 8 | 2 | Simple and complete API for building command line applications in Go | 2017-05-24 23:07:07 | 2023-11-20 21:13:22 |
| [env](https://github.com/codingconcepts/env) | 107 | 10 | 1 | Tag-based environment configuration for structs | 2017-06-14 20:01:55 | 2023-10-17 06:27:10 |
| [gocmd](https://github.com/devfacet/gocmd) | 66 | 5 | 1 | A Go library for building command line applications. | 2018-01-08 04:52:02 | 2023-10-07 13:11:01 |
| [wlog](https://github.com/dixonwille/wlog) | 63 | 9 | 0 | A simple logging interface that supports cross-platform color and concurrency. | 2016-04-13 16:47:40 | 2023-11-18 05:49:23 |
| [strumt](https://github.com/antham/strumt) | 57 | 5 | 1 | Strumt is a library to create prompt chain | 2017-06-19 19:33:16 | 2023-10-07 13:12:29 |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 51 | 10 | 2 | Fully featured Go (golang) command line option parser with built-in auto-completion support. | 2015-12-18 02:21:14 | 2023-11-09 09:55:26 |
| [flagvar](https://pkg.go.dev/github.com/sgreben/flagvar?tab=doc) | 43 | 4 | 2 | A collection of CLI argument types for the Go `flag` package. | 2018-05-18 18:45:16 | 2023-10-07 13:10:32 |
| [cmd](https://github.com/posener/cmd) | 41 | 2 | 0 | The standard library flag package with its missing features | 2019-10-29 00:32:11 | 2023-11-03 12:57:19 |
| [argv](https://github.com/cosiner/argv) | 39 | 7 | 0 | Argparse for golang. Just because `flag` sucks | 2017-01-22 10:37:21 | 2023-10-07 13:09:24 |
| [go-commander](http://yitsushi.github.io/go-commander/) | 35 | 5 | 1 | Go library to simplify CLI workflow | 2016-10-10 10:09:41 | 2023-10-07 13:10:47 |
| [sand](https://ops.city) | 22 | 2 | 0 | Package for creating interpreters | 2018-11-18 22:44:41 | 2023-10-07 13:12:24 |
| [ts](https://github.com/liujianping/ts) | 20 | 3 | 0 | timestamp convert & compare tool. 时间戳转换与对比工具 | 2019-06-25 10:21:13 | 2023-11-14 12:53:40 |
</details>

### Configuration
Libraries for configuration parsing.

<sup>*Last Update: 2023-11-03 08:17:33*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [viper](https://github.com/spf13/viper) | 24,387 | 2,003 | 484 | Go configuration with fangs | 2014-04-02 14:33:33 | 2023-11-02 11:42:44 |
| [godotenv](http://godoc.org/github.com/joho/godotenv) | 6,790 | 394 | 70 | A Go port of Ruby's dotenv library (Loads environment variables from .env files) | 2013-07-30 07:45:19 | 2023-11-02 11:36:52 |
| [envconfig](https://josh.blog/2017/04/go-configure) | 4,712 | 373 | 52 | Golang library for managing configuration data from environment variables | 2013-11-06 17:01:55 | 2023-11-02 10:57:29 |
| [env](http://carlosbecker.com/posts/env-structs-golang) | 3,914 | 230 | 4 | A simple and zero-dependencies library to parse environment variables into structs | 2015-07-28 02:14:37 | 2023-11-01 16:24:37 |
| [ini](https://ini.unknwon.io) | 3,345 | 373 | 46 | Package ini provides INI file read and write functionality in Go | 2014-12-18 07:36:37 | 2023-11-02 06:25:09 |
| [koanf](https://github.com/knadh/koanf) | 2,067 | 138 | 12 | Simple, extremely lightweight, extensible, configuration management library for Go. Support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to viper. | 2019-06-18 06:34:05 | 2023-11-02 12:53:58 |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 1,214 | 96 | 33 | ✨Clean and minimalistic environment configuration reader for Golang | 2019-07-12 15:28:52 | 2023-11-01 11:05:06 |
| [konfig](https://github.com/lalamove/konfig) | 642 | 56 | 5 | Composable, observable and performant config handling for Go for the distributed processing era | 2019-01-18 17:03:03 | 2023-10-22 17:39:08 |
| [config](https://pkg.go.dev/github.com/gookit/config/v2) | 488 | 56 | 10 | 📝 Go configuration manage(load,get,set,export). support JSON, YAML, TOML, Properties, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名 | 2018-07-07 08:11:39 | 2023-10-31 10:39:12 |
| [aconfig](https://github.com/cristalhq/aconfig) | 487 | 31 | 12 | Simple, useful and opinionated config loader. | 2020-06-26 19:43:20 | 2023-11-02 12:43:17 |
| [confita](https://github.com/heetch/confita) | 477 | 49 | 21 | Load configuration in cascade from multiple backends into a struct | 2017-12-21 10:49:18 | 2023-10-25 07:30:41 |
| [xdg](https://pkg.go.dev/github.com/adrg/xdg) | 472 | 31 | 6 | Go implementation of the XDG Base Directory Specification and XDG user directories | 2014-08-22 08:23:40 | 2023-10-27 08:52:42 |
| [config](https://github.com/golobby/config) | 344 | 30 | 3 | A lightweight yet powerful configuration manager for the Go programming language | 2019-10-15 22:51:19 | 2023-11-01 09:08:41 |
| [config](https://github.com/JeremyLoy/config) | 331 | 17 | 2 | 12 factor configuration as a typesafe struct in as little as two function calls | 2019-04-02 13:41:22 | 2023-11-01 08:47:42 |
| [hjson-go](https://hjson.github.io/) | 308 | 43 | 2 | Hjson for Go | 2016-08-05 22:59:18 | 2023-11-01 09:02:19 |
| [fig](https://github.com/kkyr/fig) | 280 | 29 | 4 | A minimalist Go configuration library | 2020-01-16 18:43:19 | 2023-10-20 09:01:03 |
| [store](https://github.com/tucnak/store) | 273 | 22 | 4 | A dead simple configuration manager for Go applications | 2015-10-03 19:17:28 | 2023-10-24 06:18:09 |
| [config](https://github.com/olebedev/config) | 268 | 43 | 4 | JSON or YAML configuration wrapper with convenient access methods. | 2014-04-21 15:09:39 | 2023-10-20 09:01:08 |
| [envconfig](https://godoc.org/github.com/tomazk/envcfg) | 235 | 30 | 1 | Small library to read your configuration from environment variables | 2015-04-21 23:37:17 | 2023-09-15 12:05:54 |
| [config](https://josh.blog/2017/04/go-configure) | 216 | 14 | 0 | 🛠 A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. | 2017-04-02 18:37:05 | 2023-09-26 04:41:10 |
| [gcfg](https://gopkg.in/gcfg.v1) | 165 | 62 | 9 | read INI-style configuration files into Go structs; supports user-defined types and subsections | 2015-08-17 14:40:55 | 2023-09-26 04:38:39 |
| [goconfig](https://pkg.go.dev/github.com/gosidekick/goconfig?tab=doc) | 149 | 21 | 6 | goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. | 2016-12-18 11:22:41 | 2021-09-02 15:09:48 |
| [harvester](https://github.com/beatlabs/harvester) | 129 | 28 | 1 | Harvest configuration, watch and notify subscriber | 2019-04-09 07:37:19 | 2023-10-03 16:34:53 |
| [onion](https://github.com/goraz/onion) | 114 | 15 | 10 | Layer based configuration for golang | 2015-07-22 14:28:21 | 2023-10-11 06:36:55 |
| [envcfg](https://godoc.org/github.com/tomazk/envcfg) | 102 | 8 | 0 | Un-marshaling environment variables to Go structs | 2014-11-29 11:43:53 | 2023-10-27 15:50:28 |
| [envh](https://github.com/antham/envh) | 97 | 2 | 0 | Go helpers to manage environment variables | 2017-01-12 11:25:48 | 2023-09-26 04:38:05 |
| [configuro](https://medium.com/better-programming/designing-cloud-native-configuration-framework-eefb0b3793cb) | 88 | 12 | 0 | An opinionated configuration loading framework for Containerized and Cloud-Native applications. | 2020-04-09 22:10:34 | 2023-11-02 13:23:06 |
| [configuration](https://github.com/BoRuDar/configuration) | 84 | 7 | 3 | Library for setting values to structs' fields from env, flags, files or default tag | 2019-11-27 17:58:49 | 2023-09-26 04:36:37 |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 79 | 7 | 2 | A cross platform package that follows the XDG Standard | 2017-07-20 15:58:29 | 2023-10-05 11:48:29 |
| [gofigure](https://github.com/ian-kent/gofigure) | 67 | 9 | 1 | Go configuration made easy! | 2014-11-25 00:12:40 | 2023-09-14 04:16:44 |
| [hocon](https://hjson.github.io/) | 65 | 15 | 10 | go implementation of lightbend's HOCON configuration library https://github.com/lightbend/config | 2020-03-01 18:20:12 | 2023-10-20 20:59:06 |
| [configure](https://github.com/paked/configure) | 56 | 10 | 2 | Configure is a Go package that gives you easy configuration of your project through redundancy | 2015-06-14 07:46:56 | 2023-10-04 06:04:20 |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 54 | 10 | 1 | Go package that interfaces with AWS System Manager | 2019-01-24 09:01:19 | 2023-10-05 08:02:55 |
| [gone](https://github.com/One-com/gone) | 47 | 8 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2023-10-26 09:27:48 |
| [go-up](https://github.com/ufoscout/go-up) | 43 | 7 | 1 | go-up! A simple configuration library with recursive placeholders resolution and no magic. | 2018-02-18 09:50:00 | 2023-09-26 04:39:22 |
| [ingo](https://github.com/schachmat/ingo) | 37 | 9 | 0 | persistent storage for flags in go | 2016-02-07 22:57:40 | 2023-10-12 02:23:21 |
| [mini](https://github.com/sasbury/mini) | 34 | 8 | 1 | A golang package for parsing ini-style configuration files | 2015-04-29 23:52:36 | 2023-05-18 01:28:53 |
| [genv](https://github.com/sakirsensoy/genv) | 34 | 4 | 1 | Genv is a library for Go (golang) that makes it easy to read and use environment variables in your projects. It also allows environment variables to be loaded from the .env file. | 2019-07-15 10:25:57 | 2023-09-26 04:38:52 |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 29 | 5 | 0 | Library providing routines to merge and validate JSON, YAML and/or TOML files | 2018-02-01 19:06:15 | 2023-10-11 06:25:17 |
| [go-ssm-config](http://godoc.org/github.com/ianlopshire/go-ssm-config) | 18 | 14 | 6 | Go utility for loading configuration parameters from AWS SSM (Parameter Store) | 2019-12-02 18:47:38 | 2023-10-09 16:49:53 |
| [envconf](https://godoc.org/github.com/tomazk/envcfg) | 11 | 5 | 0 | Configure Go applications from the environment | 2014-10-26 12:12:26 | 2022-09-26 09:20:34 |
| [env](https://github.com/nasermirzaei89/env) | 10 | 3 | 0 | Golang Get Environment Variables Package | 2019-07-24 06:37:13 | 2023-02-13 13:48:45 |
| [go-ini](https://github.com/subpop/go-ini) | 9 | 4 | 1 | automatic mirror of https://git.sr.ht/~spc/go-ini | 2019-09-11 18:38:20 | 2022-12-03 14:37:03 |
| [typenv](https://github.com/diegomarangoni/typenv) | 9 | 1 | 0 | Go minimalist typed environment variables library | 2020-06-30 18:26:09 | 2023-01-16 14:33:44 |
| [swap](https://github.com/oblq/swap) | 8 | 3 | 0 | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). | 2020-04-12 23:28:19 | 2023-03-03 09:54:25 |
| [gonfig](https://github.com/miladabc/gonfig) | 7 | 1 | 0 | Tag based configuration loader from different providers | 2021-01-21 13:44:44 | 2023-09-26 04:40:04 |
| [config](https://github.com/crgimenes/config) | 3 | 0 | 0 | Configuration loader to use in Go projects | 2022-05-30 05:49:16 | 2023-07-06 20:10:32 |
</details>

### Continuous Integration
Tools for help with continuous integration.

<sup>*Last Update: 2023-11-11 19:52:50*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gitness](https://gitness.com) | 30,471 | 2,759 | 126 | Gitness is an Open Source developer platform with Source Control management, Continuous Integration and Continuous Delivery. | 2014-02-07 07:54:44 | 2023-11-08 15:42:21 |
| [cds](https://ovh.github.io/cds/) | 4,322 | 409 | 143 | Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform | 2016-10-11 08:28:23 | 2023-11-07 12:43:32 |
| [goveralls](https://github.com/mattn/goveralls) | 767 | 140 | 17 | A tool for testing, building, signing, and publishing binaries. | 2013-04-17 10:58:40 | 2023-10-03 22:37:20 |
| [overalls](https://github.com/go-playground/overalls) | 114 | 26 | 2 | :jeans:Multi-Package go project coverprofile for tools like goveralls | 2015-07-30 11:30:11 | 2023-10-03 00:20:25 |
| [duci](https://github.com/duck8823/duci) | 74 | 5 | 10 | The simple ci server  | 2018-04-01 01:51:02 | 2023-04-13 01:57:00 |
| [gomason](https://github.com/nikogura/gomason) | 59 | 9 | 4 | A tool for testing, building, signing, and publishing binaries. | 2017-11-18 00:59:11 | 2023-10-11 06:43:54 |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 19 | 5 | 0 | A Go recursive coverage testing tool | 2016-06-26 07:45:32 | 2022-09-26 09:24:45 |
</details>

### Data Structures
Generic datastructures and algorithms in Go.

<sup>*Last Update: 2023-11-02 20:27:58*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gods](https://github.com/emirpasic/gods) | 14,531 | 1,678 | 28 | GoDS (Go Data Structures) - Sets, Lists, Stacks, Maps, Trees, Queues, and much more | 2015-03-04 14:19:52 | 2023-11-01 12:13:12 |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 7,189 | 866 | 29 | A collection of useful, performant, and threadsafe Go datastructures. | 2014-10-29 13:55:17 | 2023-11-01 09:48:48 |
| [golang-set](https://github.com/deckarep/golang-set) | 3,548 | 290 | 6 | A simple, battle-tested and generic set type for the Go language. Trusted by Docker, 1Password, Ethereum and Hashicorp. | 2013-07-03 21:52:01 | 2023-11-01 04:14:20 |
| [gota](https://github.com/go-gota/gota) | 2,772 | 300 | 79 | Gota: DataFrames and data wrangling in Go (Golang) | 2016-02-06 17:23:25 | 2023-10-29 14:12:04 |
| [roaring](http://roaringbitmap.org/) | 2,183 | 248 | 67 | Roaring bitmaps in Go (golang), used by InfluxDB, Bleve, DataDog | 2014-07-10 20:14:34 | 2023-10-30 09:49:51 |
| [bloom](https://github.com/bits-and-blooms/bloom) | 2,117 | 265 | 8 | Go package implementing Bloom filters, used by Milvus and Beego. | 2011-05-21 14:18:41 | 2023-11-01 21:47:34 |
| [gocache](https://vincent.composieux.fr/article/i-wrote-gocache-a-complete-and-extensible-go-cache-library/) | 1,994 | 207 | 20 | ☔️ A complete Go cache library that brings you multiple ways of managing your caches | 2019-10-05 08:13:54 | 2023-10-31 12:32:56 |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1,558 | 116 | 12 | Probabilistic data structures for processing continuous, unbounded streams. | 2015-02-06 02:01:26 | 2023-10-25 20:53:11 |
| [bitset](https://github.com/bits-and-blooms/bitset) | 1,189 | 184 | 4 | Go package implementing bitsets | 2011-05-11 03:33:44 | 2023-11-01 07:43:56 |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 1,024 | 103 | 13 | Cuckoo Filter: Practically Better Than Bloom | 2015-06-28 23:22:09 | 2023-10-31 23:54:49 |
| [gostl](https://github.com/liyue201/gostl) | 942 | 116 | 2 | Data structure and algorithm library for go, designed to provide functions similar to C++ STL | 2019-10-12 01:10:24 | 2023-10-25 08:47:20 |
| [hyperloglog](https://axiom.co) | 879 | 68 | 6 | HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction) brought to you by Axiom | 2017-06-18 11:18:12 | 2023-11-01 04:50:26 |
| [ttlcache](https://github.com/jellydator/ttlcache) | 755 | 109 | 8 | An in-memory cache with item expiration and generics | 2014-12-13 01:55:40 | 2023-10-31 13:57:05 |
| [algorithms](https://github.com/shady831213/algorithms) | 744 | 158 | 0 | CLRS study. Codes are written with golang. | 2018-01-31 09:27:56 | 2023-10-29 13:35:32 |
| [trie](https://github.com/derekparker/trie) | 703 | 111 | 13 | Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. | 2014-03-06 22:01:49 | 2023-10-26 11:45:18 |
| [deque](https://github.com/gammazero/deque) | 496 | 51 | 1 | Fast ring-buffer deque (double-ended queue) | 2018-04-24 02:57:55 | 2023-10-29 05:24:46 |
| [merkletree](https://github.com/cbergoon/merkletree) | 449 | 122 | 5 | A Merkle Tree implementation written in Go. | 2017-04-12 02:50:11 | 2023-10-30 06:43:19 |
| [go-edlib](https://github.com/hbollon/go-edlib) | 412 | 23 | 1 | 📚 String comparison and edit distance algorithms library, featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc... | 2020-08-18 09:30:59 | 2023-10-24 10:32:47 |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 352 | 47 | 3 | Go native library for fast point tracking and K-Nearest queries | 2015-01-22 12:26:17 | 2023-10-31 06:49:28 |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 325 | 30 | 2 | Go concurrent-safe, goroutine-safe, thread-safe queue | 2019-01-10 21:21:23 | 2023-10-26 15:11:25 |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 314 | 44 | 0 | Adaptive Radix Trees implemented in Go | 2016-04-01 01:40:40 | 2023-11-01 07:59:07 |
| [levenshtein](https://github.com/agnivade/levenshtein) | 303 | 25 | 1 | Go implementation to calculate Levenshtein Distance. | 2014-07-30 14:03:55 | 2023-10-29 21:49:28 |
| [cuckoo-filter](https://github.com/linvon/cuckoo-filter) | 275 | 26 | 2 | Cuckoo Filter go implement, better than Bloom Filter, configurable and space optimized  布谷鸟过滤器的Go实现，优于布隆过滤器，可以定制化过滤器参数，并进行了空间优化 | 2021-02-19 12:27:43 | 2023-10-31 15:09:44 |
| [hilbert](https://github.com/google/hilbert) | 269 | 43 | 3 | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. | 2015-08-06 15:50:00 | 2023-10-05 00:13:00 |
| [skiplist](https://github.com/MauriceGit/skiplist) | 263 | 37 | 5 | A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist | 2018-06-23 16:01:51 | 2023-10-26 08:11:48 |
| [goskiplist](https://github.com/ryszard/goskiplist) | 238 | 60 | 6 | A skip list implementation in Go | 2012-05-09 05:44:59 | 2023-08-23 01:59:58 |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 213 | 35 | 1 | A binary stream packer and unpacker | 2016-02-02 10:06:11 | 2023-11-01 01:14:28 |
| [iter](https://github.com/disksing/iter) | 183 | 13 | 0 | Go implementation of C++ STL iterators and algorithms. | 2019-10-20 09:29:49 | 2023-10-24 04:17:07 |
| [bit](https://yourbasic.org/algorithms/your-basic-int/#simple-sets) | 155 | 25 | 1 | Bitset data structure | 2017-05-03 19:05:35 | 2023-10-09 01:39:15 |
| [deque](https://github.com/edwingeng/deque) | 153 | 7 | 0 | A highly optimized double-ended queue | 2019-02-01 03:32:28 | 2023-10-24 21:42:31 |
| [bloom](http://zhen.org/blog/benchmarking-bloom-filters-and-hash-functions-in-go/) | 147 | 20 | 1 | Bloom filters implemented in Go. | 2013-09-03 02:27:35 | 2023-08-27 02:14:36 |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 136 | 9 | 1 | Cache Slow Database Queries | 2019-04-04 20:24:25 | 2023-10-20 00:50:47 |
| [ring](https://pkg.go.dev/github.com/tannerryan/ring) | 133 | 17 | 1 | Package ring provides a high performance and thread safe Go implementation of a bloom filter. | 2019-01-27 04:02:20 | 2023-10-09 04:30:49 |
| [go-rquad](https://github.com/arl/go-rquad) | 127 | 7 | 1 | :pushpin: State of the art point location and neighbour finding algorithms for region quadtrees, in Go | 2016-09-12 21:46:37 | 2023-10-08 16:31:35 |
| [encoding](http://zhen.org/blog/benchmarking-integer-compression-in-go/) | 127 | 17 | 4 | Integer Compression Libraries for Go | 2013-09-20 19:29:57 | 2023-10-05 03:48:42 |
| [conjungo](https://github.com/InVisionApp/conjungo) | 118 | 17 | 10 | A small flexible merge library in go | 2016-12-29 23:50:38 | 2023-10-10 18:58:32 |
| [crunch](https://github.com/superwhiskers/crunch) | 89 | 8 | 0 | take bytes out of things easily ✨🍪 | 2019-02-27 03:56:52 | 2023-09-30 04:49:00 |
| [go-mcache](https://pkg.go.dev/github.com/OrlovEvgeny/go-mcache?tab=doc) | 89 | 16 | 1 | Fast in-memory key:value store/cache with TTL | 2018-04-14 23:31:21 | 2023-09-07 02:51:52 |
| [bloom](https://yourbasic.org/algorithms/bloom-filter/) | 83 | 10 | 0 | Probabilistic set data structure | 2017-05-06 19:57:47 | 2023-10-30 10:23:27 |
| [skiplist](https://github.com/gansidui/skiplist) | 83 | 23 | 1 | skiplist for golang | 2014-11-18 16:29:53 | 2023-08-11 07:39:56 |
| [levenshtein](https://github.com/agext/levenshtein) | 80 | 6 | 1 | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. | 2016-04-08 00:14:31 | 2023-10-09 04:33:34 |
| [nan](https://github.com/kak-tus/nan) | 78 | 9 | 0 | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers | 2020-05-05 20:20:54 | 2023-10-04 00:45:43 |
| [cmap](https://github.com/lrita/cmap) | 68 | 7 | 0 | a thread-safe concurrent map for go | 2019-11-26 03:54:59 | 2023-10-09 22:27:59 |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 66 | 5 | 0 | Go implementation of Count-Min-Log | 2015-08-16 22:31:36 | 2023-10-18 02:35:01 |
| [timedmap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 66 | 9 | 1 | A thread safe map which has expiring key-value pairs. | 2019-01-30 12:55:37 | 2023-08-14 14:18:59 |
| [hide](https://godoc.org/github.com/yaa110/goterator) | 57 | 6 | 0 | ID type with marshalling to/from hash to prevent sending IDs to clients. | 2019-01-16 13:54:17 | 2023-10-09 04:32:51 |
| [goset](https://github.com/zoumo/goset) | 52 | 16 | 0 | Set is a useful collection but there is no built-in implementation in Go lang. | 2017-08-25 09:21:30 | 2023-10-16 01:46:48 |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 52 | 10 | 0 | Highly concurrent drop-in replacement for bufio.Writer | 2017-09-18 15:29:59 | 2023-10-09 04:31:48 |
| [pipeline](https://godoc.org/github.com/hyfather/pipeline) | 50 | 9 | 1 | Pipelines using goroutines | 2018-04-25 00:11:36 | 2023-07-30 16:40:25 |
| [typ](https://github.com/gurukami/typ) | 44 | 4 | 0 | Null Types, Safe primitive type conversion and fetching value from complex structures. | 2019-03-03 05:34:23 | 2023-10-04 00:44:17 |
| [dict](https://github.com/srfrog/dict) | 42 | 5 | 1 | Python-like dictionaries for Go | 2019-04-23 02:04:25 | 2023-10-09 04:31:34 |
| [ptrie](https://godoc.org/github.com/hyfather/pipeline) | 34 | 10 | 1 | A prefix tree implementation in go  | 2019-05-20 14:13:05 | 2023-10-30 09:19:19 |
| [null](https://github.com/emvi/null) | 32 | 4 | 0 | Nullable Go types that can be marshalled/unmarshalled to/from JSON. | 2018-07-04 21:18:45 | 2023-10-04 00:44:37 |
| [go-ef](https://github.com/amallia/go-ef) | 31 | 7 | 0 | A Go implementation of the Elias-Fano encoding | 2017-09-22 01:47:16 | 2023-09-30 04:49:07 |
| [set](https://github.com/StudioSol/set) | 26 | 11 | 2 | A simple Set data structure implementation in Go (Golang) using LinkedHashMap. | 2018-07-20 21:53:37 | 2023-10-31 12:17:29 |
| [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) | 25 | 7 | 3 | Ordered-concurrently a library for concurrent processing with ordered output in Go. Process work concurrently and returns output in a channel in the order of input. It is useful in concurrently processing items in a queue, and get output in the order provided by the queue. | 2021-02-28 17:56:05 | 2023-09-22 19:36:32 |
| [mspm](https://github.com/BlackRabbitt/mspm) | 24 | 4 | 0 | Multi-String Pattern Matching Algorithm Using TrieNode | 2018-05-17 18:59:44 | 2023-10-09 04:34:15 |
| [parapipe](https://github.com/nazar256/parapipe) | 24 | 3 | 1 | Paralleling pipeline | 2021-04-09 06:49:56 | 2023-09-06 04:20:27 |
| [treap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 23 | 6 | 0 | golang persistent immutable treap sorted sets | 2018-09-16 01:38:03 | 2023-10-09 04:34:58 |
| [gofal](https://github.com/xxjwxc/gofal) | 18 | 3 | 0 | fractional api base on golang . golang math tools fractional molecular denominator 分数计算 分子 分母 运算 | 2019-08-05 07:37:55 | 2023-10-04 00:49:28 |
| [slices](https://github.com/srfrog/slices) | 16 | 3 | 0 | Functions that operate on slices. Similar to functions from package strings or package bytes that have been adapted to work with slices. | 2020-07-02 23:17:34 | 2023-10-04 00:46:10 |
| [goterator](https://godoc.org/github.com/yaa110/goterator) | 15 | 4 | 0 | Lazy iterator implementation for Golang | 2020-08-12 19:47:57 | 2023-10-09 04:31:14 |
| [bloomfilter](https://github.com/OldPanda/bloomfilter) | 15 | 3 | 0 | Yet another Bloomfilter implementation in Go, compatible with Java's Guava library | 2021-01-01 01:28:04 | 2023-10-09 04:30:28 |
| [dsu](https://github.com/ihebu/dsu) | 13 | 1 | 0 | Disjoint Set data structure implementation in Go | 2021-04-27 16:35:38 | 2023-09-26 23:09:14 |
| [parsefields](https://github.com/MonaxGT/parsefields) | 7 | 1 | 0 | Tools for parse JSON-like logs for collecting unique fields and events | 2019-04-12 22:15:10 | 2023-02-16 00:39:32 |
</details>

### Database - Database schema migration


<sup>*Last Update: 2023-11-20 08:20:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [migrate](https://github.com/golang-migrate/migrate) | 12,641 | 1,254 | 311 | Database migrations. CLI and Golang library. | 2018-01-19 09:30:58 | 2023-11-19 18:42:55 |
| [goose](https://pressly.github.io/goose/) | 4,686 | 496 | 58 | A database migration tool. Supports SQL migrations and Go functions.  | 2016-02-25 20:39:37 | 2023-11-20 01:11:54 |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 2,977 | 298 | 86 | SQL schema migration tool for Go. | 2014-09-09 07:31:41 | 2023-11-18 10:28:49 |
| [pop](https://github.com/gobuffalo/pop) | 1,389 | 247 | 96 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2023-11-07 22:42:50 |
| [skeema](https://github.com/skeema/skeema) | 1,185 | 110 | 19 | Declarative pure-SQL schema management for MySQL and MariaDB | 2016-10-31 23:18:56 | 2023-11-17 20:39:51 |
| [gormigrate](https://pkg.go.dev/github.com/go-gormigrate/gormigrate/v2) | 944 | 91 | 13 | Minimalistic database migration helper for Gorm ORM | 2016-08-31 11:46:23 | 2023-11-19 20:05:38 |
| [migrator](https://github.com/lopezator/migrator) | 156 | 17 | 6 | Dead simple Go database migration library. | 2019-02-04 09:40:01 | 2023-09-14 08:39:32 |
| [darwin](https://github.com/GuiaBolso/darwin) | 139 | 36 | 5 | Database schema evolution library for Go | 2016-04-05 15:57:59 | 2023-11-10 10:08:06 |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 83 | 21 | 7 | A Go package to help write migrations with go-pg/pg. | 2018-08-11 07:00:13 | 2023-05-24 11:35:40 |
| [avro](https://github.com/khezen/avro) | 44 | 10 | 0 | Apache AVRO for go | 2019-04-07 12:22:46 | 2023-10-27 07:49:02 |
| [schema](http://pravasan.github.io/pravasan/) | 30 | 3 | 2 | Embedded schema migration package for Go | 2019-09-24 19:27:13 | 2023-08-29 08:36:58 |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 28 | 11 | 0 | Django style fixtures for Golang's excellent built-in database/sql library. | 2015-12-24 11:27:45 | 2023-09-08 17:05:02 |
| [pravasan](http://pravasan.github.io/pravasan/) | 28 | 6 | 30 | Simple Migration Tool - written in Go | 2015-01-03 17:11:05 | 2023-08-19 00:35:45 |
| [migrator](https://github.com/larapulse/migrator) | 24 | 4 | 0 | MySQL database migrator | 2020-06-27 14:40:29 | 2023-08-03 12:41:19 |
| [go-pg-migrate](https://pkg.go.dev/github.com/lawzava/go-pg-migrate) | 11 | 3 | 0 | CLI-friendly package for pg migrations management. | 2021-01-16 17:01:32 | 2023-07-19 20:05:03 |
</details>

### Database - Database tools


<sup>*Last Update: 2023-12-09 18:57:36*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [vitess](http://vitess.io) | 17,159 | 2,024 | 827 | Vitess is a database clustering system for horizontal scaling of MySQL. | 2013-06-27 21:20:28 | 2023-12-09 10:17:46 |
| [pgweb](https://sosedoff.github.io/pgweb) | 8,124 | 703 | 13 | Cross-platform client for PostgreSQL databases | 2014-10-09 01:41:32 | 2023-12-08 12:39:23 |
| [kingshard](https://github.com/flike/kingshard) | 6,278 | 1,222 | 166 | A high-performance MySQL proxy | 2015-07-04 02:22:32 | 2023-12-06 10:48:17 |
| [orchestrator](https://github.com/openark/orchestrator) | 5,244 | 930 | 413 | MySQL replication topology management and HA | 2016-11-30 13:44:24 | 2023-12-08 12:52:08 |
| [go-mysql](https://github.com/go-mysql-org/go-mysql) | 4,286 | 1,001 | 192 | a powerful mysql toolset with Go | 2014-02-21 01:56:45 | 2023-12-07 01:05:59 |
| [go-mysql-elasticsearch](https://github.com/go-mysql-org/go-mysql-elasticsearch) | 4,024 | 840 | 211 | Sync MySQL data into elasticsearch  | 2015-01-15 09:54:18 | 2023-12-08 06:07:28 |
| [prest](https://prestd.com) | 3,978 | 278 | 124 | PostgreSQL ➕ REST, low-code, simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new | 2016-11-22 05:17:05 | 2023-12-08 15:58:59 |
| [chproxy](https://www.chproxy.org/) | 1,159 | 245 | 36 | Open-Source ClickHouse http proxy and load balancer | 2017-09-18 13:09:23 | 2023-12-07 10:44:13 |
| [pg_timetable](https://www.cybertec-postgresql.com/en/products/pg_timetable/) | 962 | 64 | 0 | pg_timetable: Advanced scheduling for PostgreSQL | 2018-12-19 10:19:51 | 2023-12-03 13:44:11 |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 445 | 86 | 25 | Collects many small inserts to ClickHouse and send in big inserts | 2017-04-29 10:38:41 | 2023-12-06 10:30:10 |
| [myreplication](https://github.com/2tvenom/myreplication) | 190 | 49 | 5 | Golang MySql binary log replication listener | 2015-02-04 20:59:49 | 2023-06-28 11:21:12 |
| [octillery](https://github.com/blastrain/octillery) | 182 | 30 | 6 | Go package for sharding databases ( Supports every ORM or raw SQL ) | 2018-11-26 10:39:35 | 2023-12-03 21:30:11 |
| [dbbench](https://github.com/sj14/dbbench) | 89 | 18 | 12 | 🏋️ dbbench is a simple database benchmarking tool which supports several databases and own scripts | 2018-11-24 13:21:18 | 2023-10-24 08:23:40 |
| [datagen](https://github.com/codingconcepts/datagen) | 56 | 9 | 0 | A fast data generator that's multi-table aware and supports multi-row DML. | 2019-04-18 19:58:01 | 2023-06-09 19:49:55 |
| [prep](https://sosedoff.github.io/pgweb) | 33 | 5 | 0 | Prep finds all SQL statements in a Go package and instruments db connection with prepared statements | 2017-12-11 23:47:38 | 2022-12-10 17:38:11 |
| [rwdb](https://prestd.com) | 18 | 2 | 0 | Database wrapper that manage read write connections | 2017-10-04 03:55:29 | 2023-05-02 21:55:01 |
</details>

### Database - Databases implemented in Go


<sup>*Last Update: 2023-11-28 21:14:28*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prometheus](https://prometheus.io/) | 50,772 | 8,641 | 938 | The Prometheus monitoring system and time series database. | 2012-11-24 11:14:12 | 2023-11-28 13:50:42 |
| [tidb](https://pingcap.com) | 35,253 | 5,674 | 4,453 | TiDB is an open-source, cloud-native, distributed, MySQL-Compatible database for elastic scale and real-time analytics. Try AI-powered Chat2Query free at : https://tidbcloud.com/free-trial | 2015-09-06 04:01:52 | 2023-11-28 08:43:43 |
| [cockroach](https://www.cockroachlabs.com) | 28,161 | 3,598 | 5,580 | CockroachDB - the open source, cloud-native distributed SQL database. | 2014-02-06 00:18:47 | 2023-11-28 12:08:56 |
| [influxdb](https://influxdata.com) | 26,752 | 3,480 | 1,862 | Scalable datastore for metrics, events, and real-time analytics | 2013-09-26 14:31:10 | 2023-11-28 10:38:03 |
| [dgraph](https://dgraph.io) | 19,757 | 1,503 | 285 | The high-performance database for modern applications | 2015-08-25 07:15:56 | 2023-11-26 12:29:35 |
| [rqlite](https://rqlite.io) | 14,192 | 674 | 57 | The lightweight, distributed relational database built on SQLite | 2014-08-23 04:31:18 | 2023-11-28 12:31:04 |
| [badger](https://dgraph.io/badger) | 12,936 | 1,172 | 53 | Fast key-value DB in Go. | 2017-01-26 05:09:49 | 2023-11-28 13:15:50 |
| [groupcache](https://github.com/golang/groupcache) | 12,494 | 1,397 | 40 | groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. | 2013-07-22 21:55:07 | 2023-11-28 08:53:24 |
| [VictoriaMetrics](https://victoriametrics.com/) | 9,795 | 1,004 | 841 | VictoriaMetrics: fast, cost-effective monitoring solution and time series database | 2018-09-30 09:58:01 | 2023-11-28 11:38:02 |
| [immudb](https://immudb.io) | 8,389 | 335 | 109 | immudb - immutable database based on zero trust, SQL/Key-Value/Document model, tamperproof, data change history | 2019-11-07 08:22:16 | 2023-11-28 12:02:02 |
| [go-cache](https://patrickmn.com/projects/go-cache/) | 7,548 | 885 | 72 | An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. | 2012-01-02 13:07:13 | 2023-11-28 11:27:23 |
| [bbolt](https://go.etcd.io/bbolt) | 7,091 | 620 | 75 | An embedded key/value database for Go. | 2017-06-17 01:42:09 | 2023-11-28 09:53:53 |
| [bigcache](http://allegro.tech/2016/03/writing-fast-cache-service-in-go.html) | 6,899 | 609 | 89 | Efficient cache for gigabytes of data written in Go. | 2016-03-23 07:18:52 | 2023-11-26 15:05:34 |
| [goleveldb](https://patrickmn.com/projects/go-cache/) | 5,885 | 961 | 107 | LevelDB key/value database in Go. | 2013-01-23 04:08:58 | 2023-11-28 04:44:20 |
| [buntdb](https://github.com/tidwall/buntdb) | 4,254 | 285 | 27 | BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support | 2016-07-19 22:11:40 | 2023-11-27 09:07:10 |
| [rosedb](https://rosedblabs.github.io) | 4,109 | 609 | 11 | Lightweight, fast and reliable key/value storage engine based on Bitcask. | 2020-12-06 07:02:48 | 2023-11-28 07:10:27 |
| [ledisdb](https://ledisdb.io) | 4,020 | 496 | 2 | A high performance NoSQL Database Server powered by Go | 2014-04-30 00:43:09 | 2023-11-26 08:44:23 |
| [nutsdb](https://nutsdb.github.io/nutsdb/) | 3,074 | 306 | 48 | A simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. | 2018-12-07 07:03:38 | 2023-11-28 06:29:39 |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2,716 | 271 | 27 | A rudimentary implementation of a basic document (NoSQL) database in Go | 2013-05-26 10:03:49 | 2023-11-27 02:41:11 |
| [gcache](https://github.com/bluele/gcache) | 2,462 | 262 | 27 | An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC | 2015-01-24 18:17:07 | 2023-11-27 11:02:33 |
| [cache2go](https://github.com/muesli/cache2go) | 2,014 | 557 | 32 | Concurrency-safe Go caching library with expiration capabilities and access counters | 2013-11-11 03:45:02 | 2023-11-25 08:25:09 |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 1,895 | 167 | 44 | Fast thread-safe inmemory cache for big number of entries in Go. Minimizes GC overhead | 2018-11-22 22:50:13 | 2023-11-28 10:41:55 |
| [CovenantSQL](https://developers.covenantsql.io) | 1,461 | 177 | 36 | A decentralized, trusted, high performance, SQL database with blockchain features | 2018-04-11 09:52:58 | 2023-11-27 08:38:20 |
| [diskv](http://godoc.org/github.com/peterbourgon/diskv) | 1,308 | 108 | 8 | A disk-backed key-value store. | 2012-03-21 16:44:32 | 2023-11-24 19:14:38 |
| [pogreb](https://github.com/akrylysov/pogreb) | 1,162 | 90 | 19 | Embedded key-value store for read-heavy workloads written in Go | 2018-01-06 23:16:36 | 2023-11-28 09:40:27 |
| [databunker](https://databunker.org/) | 1,143 | 66 | 7 | Secure SDK/vault for personal records/PII built to comply with GDPR | 2019-12-08 21:55:55 | 2023-11-27 14:54:01 |
| [eliasdb](https://github.com/krotik/eliasdb) | 981 | 49 | 14 | EliasDB a graph-based database. | 2016-08-13 13:53:28 | 2023-11-24 19:23:57 |
| [moss](https://github.com/couchbase/moss) | 936 | 59 | 46 | moss - a simple, fast, ordered, persistable, key-val storage library for golang | 2016-02-06 20:27:22 | 2023-11-27 01:14:47 |
| [bitcask](https://prologic.github.io/bitcask) | 759 | 70 | 10 | 🔑A high performance Key/Value store written in Go with a predictable read/write performance and high throughput. Uses a Bitcask on-disk layout (LSM+WAL) similar to Riak. | 2019-03-12 13:57:35 | 2021-06-22 16:39:37 |
| [levigo](https://github.com/jmhodges/levigo) | 417 | 89 | 6 | levigo is a Go wrapper for LevelDB | 2012-01-17 08:17:54 | 2023-11-19 01:58:20 |
| [pudge](https://github.com/recoilme/pudge) | 349 | 31 | 0 | Fast and simple key/value store written using Go's standard library | 2018-11-20 10:11:53 | 2023-11-26 19:09:43 |
| [kivik](https://github.com/go-kivik/kivik) | 282 | 40 | 26 | Common interface to CouchDB or CouchDB-like databases for Go and GopherJS | 2017-02-09 14:14:54 | 2023-11-21 14:55:20 |
| [vasto](https://github.com/chrislusf/vasto) | 250 | 30 | 4 | A distributed key-value store. On Disk. Able to grow or shrink without service interruption. | 2018-01-16 05:16:57 | 2023-11-25 01:48:12 |
| [piladb](https://www.piladb.org) | 200 | 21 | 9 | Lightweight RESTful database engine based on stack data structures | 2015-09-08 23:12:22 | 2023-11-24 19:18:38 |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 170 | 28 | 1 | A tiny Golang JSON database | 2018-06-21 22:13:33 | 2023-11-14 07:24:56 |
| [cache](https://github.com/akyoto/cache) | 159 | 19 | 0 | :handbag: Cache arbitrary data with an expiration time. | 2019-05-11 12:42:45 | 2023-11-11 17:10:40 |
| [bcache](https://github.com/iwanbk/bcache) | 139 | 16 | 4 | Eventually consistent distributed in-memory  cache Go library | 2018-12-26 15:45:16 | 2023-11-18 16:41:11 |
| [unitdb](https://github.com/unit-io/unitdb) | 113 | 10 | 2 | Fast specialized time-series database for IoT, real-time internet connected devices and AI analytics. | 2019-08-29 18:21:27 | 2023-11-11 10:35:46 |
| [slowpoke](https://github.com/recoilme/slowpoke) | 98 | 9 | 0 | Low-level key/value store in pure Go.  | 2018-02-19 09:22:37 | 2023-08-28 08:43:34 |
| [hare](https://github.com/jameycribbs/hare) | 84 | 11 | 1 | Hare is a nimble little database management system for Go. | 2016-10-05 20:05:45 | 2023-11-23 22:16:18 |
| [couchcache](https://github.com/codingsince1985/couchcache) | 60 | 6 | 0 | A RESTful caching micro-service in Go backed by Couchbase | 2015-04-05 07:13:05 | 2023-06-09 07:54:33 |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 44 | 5 | 2 | golang bigcache with clustering as a library. | 2017-12-18 07:48:07 | 2023-11-01 01:52:19 |
| [coffer](https://github.com/claygod/coffer) | 34 | 4 | 1 | Simply ACID* key-value database. At the medium or even low latency it tries to provide greater throughput without losing the ACID properties of the database. The database provides the ability to create record headers at own discretion and use them as transactions. The maximum size of stored data is limited by the size of the computer's RAM. | 2019-05-13 18:30:23 | 2023-07-19 22:30:32 |
| [bitcask](https://git.mills.io/prologic/bitcask) | 20 | 2 | 0 | 🔑 A high performance Key/Value store written in Go with a predictable read/write performance and high throughput. Uses a Bitcask on-disk layout (LSM+WAL) similar to Riak. | 2021-07-12 14:52:18 | 2023-11-04 08:45:27 |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 17 | 3 | 0 | Key-value store for temporary items :memo: | 2017-03-17 18:03:42 | 2022-09-26 09:37:38 |
| [ttlcache](https://github.com/cheshir/ttlcache) | 9 | 8 | 0 | Simple in-memory key-value storage with TTL for each record. | 2021-01-06 19:24:26 | 2023-03-26 17:03:11 |
</details>

### Database - SQL query builder
libraries for building and using SQL

<sup>*Last Update: 2023-12-16 21:00:25*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [squirrel](https://github.com/Masterminds/squirrel) | 6,242 | 442 | 80 | Fluent SQL generation for golang | 2014-01-18 05:29:58 | 2023-12-13 05:41:20 |
| [xo](https://github.com/xo/xo) | 3,492 | 312 | 41 | Command line tool to generate idiomatic Go code for SQL databases supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server | 2016-02-05 10:22:20 | 2023-12-16 07:52:32 |
| [goqu](http://doug-martin.github.io/goqu/) | 2,081 | 220 | 114 | SQL builder and query library for golang | 2015-02-21 01:06:00 | 2023-12-12 16:13:31 |
| [jet](https://github.com/go-jet/jet) | 1,723 | 100 | 36 | Type safe SQL builder with code generation and automatic query result data mapping | 2019-03-02 11:06:23 | 2023-12-13 01:36:11 |
| [gendry](https://github.com/didi/gendry) | 1,549 | 189 | 16 | a golang library for sql builder | 2017-12-01 08:15:43 | 2023-12-13 08:44:00 |
| [dotsql](https://github.com/qustavo/dotsql) | 692 | 52 | 8 | A Golang library for using SQL. | 2014-11-20 12:14:39 | 2023-12-12 15:38:51 |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 614 | 76 | 41 | A Go (golang) package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. | 2015-12-10 22:39:26 | 2023-11-30 07:01:18 |
| [dbq](https://github.com/rocketlaunchr/dbq) | 387 | 22 | 1 | Zero boilerplate database operations for Go | 2019-07-11 02:17:33 | 2023-11-29 14:27:31 |
| [sqlingo](https://github.com/lqs/sqlingo) | 352 | 29 | 3 | 💥 A lightweight DSL & ORM which helps you to write SQL in Go. | 2018-11-18 14:11:03 | 2023-11-23 12:04:20 |
| [sqrl](https://github.com/elgris/sqrl) | 267 | 37 | 7 | Fluent SQL generation for golang | 2014-06-25 10:03:06 | 2023-12-08 07:49:47 |
| [go-structured-query](https://bokwoon95.github.io/sq/) | 195 | 11 | 2 | Type safe SQL query builder and struct mapper for Go | 2020-05-30 14:07:30 | 2023-12-04 23:16:41 |
| [sqlf](https://github.com/leporo/sqlf) | 134 | 12 | 3 | Fast SQL query builder for Go | 2019-07-20 07:03:27 | 2023-11-04 20:59:03 |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 132 | 15 | 8 | Go database query builder library for PostgreSQL | 2019-08-18 08:18:21 | 2023-11-29 10:06:31 |
| [go-hasql](https://golang.yandex) | 118 | 14 | 5 | Go library for accessing multi-host SQL database installations | 2020-08-19 09:56:00 | 2023-12-07 13:01:08 |
| [igor](https://github.com/galeone/igor) | 115 | 4 | 0 | igor is an abstraction layer for PostgreSQL with a gorm like syntax. | 2016-03-10 14:45:08 | 2023-12-05 22:22:35 |
| [godbal](https://github.com/xujiajun/godbal) | 59 | 29 | 0 | Database Abstraction Layer (dbal) for Go. Support SQL builder and get result easily  (now only support mysql) | 2018-02-28 05:47:42 | 2023-10-31 11:12:59 |
| [gosql](https://twharmon.gitbook.io/gosql/) | 29 | 2 | 0 | SQL query builder for Go | 2020-01-08 17:13:09 | 2023-11-10 09:18:48 |
| [qry](https://github.com/HnH/qry) | 29 | 5 | 1 | Write your SQL queries in raw files with all benefits of modern IDEs, use them in an easy way inside your application with all the profit of compile time constants | 2019-08-20 09:01:00 | 2023-05-01 07:58:57 |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 13 | 2 | 0 | Golang package for MPTT (Modified Preorder Tree Traversal) - materialized path realisation. | 2020-01-09 15:04:45 | 2023-09-28 07:48:10 |
| [ormlite](https://github.com/pupizoid/ormlite) | 12 | 3 | 0 | Lightweight package containing some ORM-like features and helpers for sqlite databases. | 2018-06-28 13:42:09 | 2023-10-18 14:14:54 |
</details>

### Database Drivers - Multiple Backends.


<sup>*Last Update: 2023-12-04 09:33:27*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cayley](https://cayley.io) | 14,677 | 1,289 | 92 | An open-source graph database | 2014-06-05 18:49:41 | 2023-12-03 17:04:47 |
| [gokv](https://github.com/philippgille/gokv) | 596 | 65 | 35 | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) | 2018-10-08 18:55:22 | 2023-11-26 07:27:04 |
| [cachego](https://github.com/faabiosr/cachego) | 323 | 22 | 0 | Golang Cache component - Multiple drivers | 2016-10-05 18:10:03 | 2023-12-01 07:32:51 |
| [dsc](https://github.com/viant/dsc) | 27 | 10 | 2 | Datastore Connectivity in go | 2016-06-13 20:18:10 | 2023-07-27 16:52:56 |
</details>

### Database Drivers - NoSQL Databases


<sup>*Last Update: 2023-12-13 19:33:49*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-redis](https://redis.uptrace.dev) | 18,459 | 2,222 | 237 | Redis Go client | 2012-07-25 13:01:39 | 2023-12-13 10:05:08 |
| [redigo](https://github.com/gomodule/redigo) | 9,656 | 1,318 | 17 | Go client for Redis | 2012-04-14 04:31:58 | 2023-12-13 08:59:33 |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 7,735 | 936 | 9 | The Official Golang driver for MongoDB | 2017-02-08 17:18:02 | 2023-12-13 10:01:30 |
| [gocql](http://gocql.github.io/) | 2,477 | 630 | 176 | Package gocql implements a fast and robust Cassandra client for the Go programming language. | 2012-08-26 15:42:42 | 2023-12-11 06:41:03 |
| [mgo](https://github.com/globalsign/mgo) | 1,959 | 240 | 67 | The MongoDB driver for Go | 2017-04-13 11:14:04 | 2023-12-09 13:56:33 |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1,652 | 489 | 45 | Go Memcached client library #golang | 2011-06-28 19:29:12 | 2023-12-11 06:17:16 |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1,643 | 228 | 20 | Go language driver for RethinkDB | 2013-09-12 13:56:27 | 2023-11-26 09:51:09 |
| [qmgo](https://github.com/qiniu/qmgo) | 1,222 | 146 | 41 | Qmgo - The Go driver for MongoDB. It‘s based on official mongo-go-driver but easier to use like Mgo. | 2020-08-04 09:06:00 | 2023-12-12 06:58:12 |
| [mgm](https://github.com/Kamva/mgm) | 704 | 64 | 13 | Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver) | 2019-12-27 14:40:51 | 2023-12-13 06:21:30 |
| [redeo](https://github.com/bsm/redeo) | 429 | 39 | 2 | High-performance framework for building redis-protocol compatible TCP servers/services | 2014-03-06 08:46:18 | 2023-11-24 19:13:55 |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 416 | 226 | 38 | Aerospike Client Go  | 2014-07-26 02:56:21 | 2023-12-12 17:29:50 |
| [neoism](https://github.com/jmcvetta/neoism) | 389 | 57 | 14 | Neo4j client for Golang | 2012-07-12 07:42:33 | 2023-06-24 09:03:09 |
| [gocb](http://blog.couchbase.com/2015/september/go-sdk-1.0-ga) | 352 | 104 | 0 | The Couchbase Go SDK | 2015-01-15 20:01:32 | 2023-12-09 20:43:23 |
| [go-rejson](https://github.com/nitishm/go-rejson) | 331 | 50 | 15 | Golang client for redislabs' ReJSON module with support for multilple redis clients (redigo, go-redis) | 2018-04-23 00:51:05 | 2023-10-29 12:00:02 |
| [go-couchbase](https://godoc.org/github.com/couchbase/go-couchbase) | 319 | 92 | 41 | Couchbase client in Go | 2012-01-19 22:52:08 | 2023-09-22 10:35:34 |
| [godis](https://github.com/piaohao/godis) | 109 | 17 | 0 | redis client implement by golang, inspired by jedis. | 2019-06-14 03:14:22 | 2023-08-28 08:58:28 |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 76 | 19 | 0 | Neo4j REST Client in golang | 2011-06-04 16:08:35 | 2023-07-27 20:01:33 |
| [arangolite](https://github.com/solher/arangolite) | 73 | 19 | 5 | Lightweight Golang driver for ArangoDB | 2015-10-04 17:27:34 | 2022-09-26 09:46:39 |
| [go-pilosa](https://www.pilosa.com/) | 58 | 23 | 13 | Go client library for Pilosa | 2016-09-30 21:37:10 | 2023-04-24 14:20:57 |
| [goforestdb](https://github.com/couchbase/goforestdb) | 34 | 6 | 7 | Go bindings for ForestDB | 2014-05-14 15:36:12 | 2023-09-24 06:32:29 |
| [goriak](https://godoc.org/gopkg.in/zegl/goriak.v3) | 30 | 6 | 5 | goriak - Go language driver for Riak KV | 2016-10-05 16:48:17 | 2023-04-04 10:04:44 |
| [neo4j](https://github.com/cihangir/neo4j) | 27 | 9 | 8 | Neo4j Rest API Client for Go lang | 2013-05-18 08:54:01 | 2022-11-09 17:49:55 |
| [xredis](https://github.com/shomali11/xredis) | 19 | 6 | 0 | Go Redis Client | 2017-06-14 00:19:26 | 2022-09-26 09:47:39 |
| [gocosmos](https://github.com/btnguyen2k/gocosmos) | 19 | 7 | 1 | Go database/sql driver for Azure Cosmos DB SQL API | 2020-12-06 07:03:43 | 2023-12-05 12:52:21 |
| [godscache](https://github.com/defcronyke/godscache) | 11 | 2 | 0 | An unofficial Google Cloud Platform Go Datastore wrapper that adds caching using memcached. For App Engine Flexible, Compute Engine, Kubernetes Engine, and more. | 2018-05-08 20:19:39 | 2022-09-26 09:47:05 |
| [asc](https://github.com/viant/asc) | 9 | 3 | 0 | Datastore Connectivity for Aerospike for go | 2016-06-13 20:22:31 | 2022-09-26 09:46:41 |
</details>

### Database Drivers - Relational Databases


<sup>*Last Update: 2023-12-11 19:37:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql) | 13,886 | 2,380 | 73 | Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package | 2012-12-09 20:33:55 | 2023-12-11 10:01:49 |
| [pgx](https://github.com/jackc/pgx) | 8,647 | 784 | 142 | PostgreSQL driver and toolkit for Go | 2013-03-30 19:06:26 | 2023-12-11 10:52:48 |
| [pq](https://pkg.go.dev/github.com/lib/pq) | 8,468 | 954 | 307 | Pure Go Postgres driver for database/sql | 2012-03-12 18:50:22 | 2023-12-11 09:22:22 |
| [go-sqlite3](http://mattn.github.io/go-sqlite3) | 7,056 | 1,107 | 137 | sqlite3 driver for go using database/sql | 2011-11-11 12:36:50 | 2023-12-11 08:24:32 |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1,775 | 502 | 166 | Microsoft SQL server driver written in go language | 2013-12-16 00:10:47 | 2023-12-11 06:17:39 |
| [go-oci8](https://mattn.kaoriya.net/) | 618 | 216 | 20 | Oracle driver for Go using database/sql | 2012-02-29 12:19:16 | 2023-11-27 08:32:17 |
| [godror](https://github.com/godror/godror) | 471 | 96 | 5 | GO DRiver for ORacle DB | 2019-11-21 21:23:17 | 2023-12-07 16:34:21 |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 348 | 15 | 0 | Golang SQLite without cgo | 2020-06-06 20:37:12 | 2023-12-08 05:44:58 |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 209 | 57 | 12 | Firebird RDBMS sql driver for Go (golang) | 2013-08-27 13:09:14 | 2023-12-02 08:52:49 |
| [go-adodb](http://mattn.kaoriya.net/) | 132 | 36 | 19 | Microsoft ActiveX Object DataBase driver for go that using exp/sql | 2011-11-14 04:32:50 | 2023-11-17 15:16:34 |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 110 | 37 | 0 | Mirror of Apache Calcite - Avatica Go SQL Driver | 2017-08-08 07:00:08 | 2023-12-06 10:38:25 |
| [gofreetds](https://github.com/minus5/gofreetds) | 109 | 48 | 18 | Go Sql Server database driver. | 2012-12-06 17:29:26 | 2023-05-24 11:51:09 |
| [bgc](https://github.com/viant/bgc) | 20 | 8 | 0 | Datastore Connectivity for BigQuery in go | 2016-06-13 20:24:26 | 2023-04-23 05:10:45 |
| [pig](https://github.com/alexeyco/pig) | 15 | 3 | 0 | Simple pgx wrapper to execute and scan query results | 2021-04-15 15:33:23 | 2023-10-12 23:00:33 |
</details>

### Database Drivers - Search and Analytic Databases


<sup>*Last Update: 2023-11-18 20:15:35*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [bleve](https://github.com/blevesearch/bleve) | 9,340 | 695 | 303 | A modern text indexing library for go | 2014-04-17 21:02:18 | 2023-11-18 08:59:22 |
| [elastic](https://olivere.github.io/elastic/) | 7,254 | 1,204 | 106 | Deprecated: Use the official Elasticsearch client for Go at https://github.com/elastic/go-elasticsearch | 2012-12-06 17:15:33 | 2023-11-17 04:00:32 |
| [riot](https://github.com/go-ego/riot) | 6,095 | 483 | 50 | Go Open Source, Distributed, Simple and efficient Search Engine; Warning: This is V1 and beta version, because of big memory consume, and the V2 will be rewrite all code. | 2017-06-21 14:17:59 | 2023-11-16 09:34:19 |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch#go-elasticsearch) | 5,246 | 610 | 48 | The official Go client for Elasticsearch | 2017-03-27 17:56:15 | 2023-11-18 09:57:49 |
| [elasticsql](https://github.com/cch123/elasticsql) | 1,108 | 195 | 10 | convert sql to elasticsearch DSL in golang(go) | 2016-08-24 07:29:43 | 2023-11-13 02:56:51 |
| [elastigo](https://github.com/mattbaird/elastigo) | 949 | 254 | 70 | A Go (golang) based Elasticsearch client library. | 2012-10-12 04:19:59 | 2023-11-10 15:47:05 |
| [skizze](https://github.com/seiflotfy/skizze) | 88 | 9 | 0 | A probabilistic data structure service and storage | 2016-01-17 12:10:40 | 2023-05-24 04:56:57 |
| [goes](http://godoc.org/github.com/belogik/goes) | 30 | 14 | 0 | A library to interact with Elasticsearch in Go! | 2015-12-28 18:52:03 | 2023-09-28 14:15:37 |
</details>

### Date and Time
Libraries for working with dates and times.

<sup>*Last Update: 2023-12-11 19:37:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [now](https://github.com/jinzhu/now) | 4,287 | 235 | 13 | Now is a time toolkit for golang | 2013-11-18 10:55:30 | 2023-12-11 06:17:38 |
| [dateparse](https://github.com/araddon/dateparse) | 1,939 | 156 | 59 | GoLang Parse many date strings without knowing format in advance. | 2014-04-21 02:55:48 | 2023-12-11 10:21:16 |
| [carbon](https://github.com/uniplaces/carbon) | 768 | 55 | 2 | Carbon for Golang, an extension for Time | 2016-08-03 10:55:52 | 2023-12-02 19:03:43 |
| [durafmt](https://github.com/hako/durafmt) | 475 | 51 | 10 | :clock8:  Better time duration formatting in Go!  | 2016-05-20 21:49:33 | 2023-12-06 18:08:27 |
| [timeutil](https://github.com/leekchan/timeutil) | 194 | 15 | 2 | timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package | 2015-08-02 01:32:06 | 2023-08-22 11:02:40 |
| [gostradamus](https://github.com/bykof/gostradamus) | 192 | 6 | 3 | Gostradamus: Better DateTimes for Go 🕰️ | 2020-04-07 12:29:21 | 2023-11-12 14:53:58 |
| [go-persian-calendar](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 175 | 22 | 5 | The implementation of Persian (Solar Hijri) Calendar in Go | 2016-01-31 18:40:23 | 2023-12-10 11:10:21 |
| [iso8601](https://github.com/relvacode/iso8601) | 132 | 15 | 2 | A fast ISO8601 date parser for Go | 2017-04-25 15:54:18 | 2023-10-06 09:59:45 |
| [date](https://godoc.org/github.com/rickb777/date) | 117 | 22 | 0 | A Go package for working with dates | 2015-11-23 22:58:07 | 2023-11-02 05:48:18 |
| [go-str2duration](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 85 | 6 | 1 | Convert string to duration in golang | 2020-02-02 06:04:07 | 2023-12-07 09:52:24 |
| [timespan](https://github.com/SaidinWoT/timespan) | 82 | 12 | 3 | Golang package to manipulate time intervals. | 2014-10-07 03:57:02 | 2022-09-26 09:53:23 |
| [go-sunrise](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 79 | 13 | 1 | Go package for calculating the sunrise and sunset times for a given location | 2017-06-15 20:49:41 | 2023-11-14 15:25:10 |
| [feiertage](https://github.com/wlbr/feiertage) | 45 | 8 | 1 | Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany) | 2015-11-04 14:19:27 | 2023-11-06 10:05:26 |
| [kair](https://github.com/GuilhermeCaruso/kair) | 25 | 6 | 0 | :clock1: Date and Time - Golang Formatting Library | 2018-10-03 15:44:07 | 2023-02-03 00:57:40 |
| [cronrange](https://github.com/1set/cronrange) | 18 | 7 | 2 | time range expression in cron style | 2019-11-10 01:30:45 | 2022-09-26 09:52:51 |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 14 | 4 | 0 | Now is a time toolkit for golang | 2016-03-06 01:44:58 | 2023-04-02 02:05:30 |
| [tuesday](https://godoc.org/github.com/osteele/tuesday) | 12 | 3 | 1 | Ruby-compatible strftime for golang | 2017-08-10 20:46:26 | 2022-09-26 09:53:28 |
| [strftime](https://github.com/awoodbeck/strftime) | 12 | 5 | 0 | C99-compatible strftime formatter for use with Go time.Time instances. | 2018-02-10 00:35:46 | 2023-02-10 14:32:27 |
| [go-week](https://github.com/stoewer/go-week) | 9 | 8 | 2 | A Go package to work with ISO 8601 week dates | 2018-02-23 07:02:37 | 2023-11-10 03:06:21 |
</details>

### Distributed Systems
Packages that help with building Distributed Systems.

<sup>*Last Update: 2023-12-14 20:06:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [etcd](https://etcd.io) | 45,314 | 9,622 | 261 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2023-12-14 12:46:00 |
| [go-zero](https://go-zero.dev) | 26,573 | 3,750 | 438 | A cloud-native Go microservices framework with cli tool for productivity. | 2020-08-07 15:37:57 | 2023-12-14 11:45:19 |
| [kit](https://gokit.io) | 25,725 | 2,479 | 44 | A standard library for microservices. | 2015-02-03 00:01:19 | 2023-12-14 12:47:14 |
| [go-micro](https://go-micro.dev) | 21,039 | 2,355 | 86 | A Go microservices framework | 2015-01-13 23:30:18 | 2023-12-14 01:18:30 |
| [grpc-go](https://grpc.io) | 19,245 | 4,205 | 132 | The Go language implementation of gRPC. HTTP/2 based RPC | 2014-12-08 18:59:34 | 2023-12-14 09:21:37 |
| [nats-server](https://nats.io) | 13,962 | 1,319 | 304 | High-Performance server for NATS.io, the cloud and edge native messaging system. | 2012-10-29 16:12:24 | 2023-12-14 12:50:34 |
| [micro](https://micro.dev) | 11,929 | 1,099 | 45 | API first development platform | 2015-01-16 22:35:14 | 2023-12-14 11:45:40 |
| [rpcx](https://rpcx.io) | 7,815 | 1,157 | 7 | Best microservices framework in Go, like alibaba Dubbo, but with more features, Scale easily. Try it. Test it. If you feel it's better, use it! 𝐉𝐚𝐯𝐚有𝐝𝐮𝐛𝐛𝐨, 𝐆𝐨𝐥𝐚𝐧𝐠有𝐫𝐩𝐜𝐱! build for cloud! | 2016-05-18 09:34:05 | 2023-12-14 11:15:13 |
| [raft](https://godoc.org/cirello.io/pglock) | 7,563 | 946 | 34 | Golang implementation of the Raft consensus protocol | 2013-11-05 00:41:20 | 2023-12-14 10:59:53 |
| [lura](https://luraproject.org) | 5,839 | 545 | 11 | Ultra performant API Gateway with middlewares. A project hosted at The Linux Foundation | 2016-11-04 18:37:13 | 2023-12-13 22:57:31 |
| [tendermint](https://tendermint.com/) | 5,580 | 2,139 | 20 | ⟁ Tendermint Core (BFT Consensus) in Go | 2014-05-14 23:21:35 | 2023-12-13 09:36:03 |
| [torrent](https://github.com/anacrolix/torrent) | 5,016 | 609 | 47 | Full-featured BitTorrent client package and utilities | 2015-01-08 21:10:42 | 2023-12-14 01:30:07 |
| [dragonboat](https://github.com/lni/dragonboat) | 4,803 | 525 | 34 | A feature complete and high performance multi-group Raft library in Go.   | 2018-12-23 07:02:04 | 2023-12-14 01:45:18 |
| [emitter](https://emitter.io) | 3,674 | 347 | 13 | High performance, distributed and low latency publish-subscribe platform. | 2016-10-29 08:52:21 | 2023-12-14 08:58:26 |
| [gleam](https://github.com/chrislusf/gleam) | 3,281 | 287 | 39 | Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly. | 2016-08-26 08:44:48 | 2023-12-14 07:57:36 |
| [glow](https://github.com/chrislusf/glow) | 3,169 | 249 | 15 | Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant. | 2015-06-14 00:33:48 | 2023-12-11 06:17:16 |
| [liftbridge](https://liftbridge.io) | 2,501 | 107 | 45 | Lightweight, fault-tolerant message streams. | 2017-10-13 19:50:26 | 2023-12-12 08:57:32 |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1,247 | 211 | 5 | Hprose is a cross-language RPC. This project is Hprose for Golang. | 2014-02-14 03:16:43 | 2023-12-10 03:02:16 |
| [redislock](https://github.com/bsm/redislock) | 1,184 | 144 | 1 | Simplified distributed locking implementation using Redis | 2019-06-24 11:10:10 | 2023-12-13 01:40:42 |
| [rain](https://github.com/cenkalti/rain) | 890 | 65 | 3 | 🌧 BitTorrent client and library in Go | 2014-05-21 09:17:24 | 2023-12-11 17:51:42 |
| [arpc](https://github.com/lesismal/arpc) | 822 | 70 | 0 | More effective network communication, two-way calling, notify and broadcast supported. | 2020-05-19 11:30:05 | 2023-12-12 03:52:05 |
| [ringpop-go](http://www.uber.com) | 804 | 83 | 27 | Scalable, fault-tolerant application-layer sharding for Go applications | 2015-06-05 22:48:53 | 2023-11-24 19:17:39 |
| [go-health](https://github.com/InVisionApp/go-health) | 729 | 56 | 12 | Library for enabling asynchronous health checks in your service | 2017-11-29 21:00:07 | 2023-12-13 23:18:37 |
| [gorpc](https://github.com/valyala/gorpc) | 683 | 101 | 15 | Simple, fast and scalable golang rpc library for high load | 2014-11-20 17:02:37 | 2023-12-05 04:33:39 |
| [resgate](https://resgate.io) | 645 | 64 | 19 | A Realtime API Gateway used with NATS to build REST, real time, and RPC APIs, where all your clients are synchronized seamlessly. | 2018-02-22 12:06:26 | 2023-12-12 10:31:39 |
| [consistent](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 630 | 67 | 6 | Consistent hashing with bounded loads in Golang | 2018-03-25 15:38:27 | 2023-12-12 04:16:36 |
| [go-sundheit](https://pdu.pub) | 514 | 31 | 4 | A library built to provide support for defining service health for golang services. It allows you to register async health checks for your dependencies and the service itself, provides a health endpoint that exposes their status, and health metrics. | 2019-04-08 12:54:01 | 2023-12-06 12:46:17 |
| [digota](https://github.com/digota/digota) | 487 | 79 | 11 | ecommerce microservice | 2017-08-14 12:01:37 | 2023-12-08 01:14:31 |
| [go-jump](https://github.com/dgryski/go-jump) | 370 | 33 | 1 | go-jump: Jump consistent hashing | 2014-06-15 22:12:04 | 2023-11-23 03:15:28 |
| [sleuth](http://ursiform.github.io/sleuth/) | 367 | 25 | 0 | A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services | 2016-04-23 14:21:41 | 2023-12-05 02:10:18 |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 291 | 88 | 0 | A simple go implementation of json rpc 2.0 client over http | 2016-11-10 11:27:55 | 2023-12-12 08:29:53 |
| [dht](https://github.com/anacrolix/dht) | 277 | 61 | 3 | dht is used by anacrolix/torrent, and is intended for use as a library in other projects both torrent related and otherwise | 2016-12-14 00:34:42 | 2023-12-13 01:54:55 |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 183 | 22 | 5 | The jsonrpc package helps implement of JSON-RPC 2.0 | 2016-10-28 13:36:59 | 2023-12-11 15:12:27 |
| [outboxer](https://github.com/italolelis/outboxer) | 147 | 26 | 5 | A library that implements the outboxer pattern in go | 2019-02-01 09:50:13 | 2023-11-25 15:01:06 |
| [dynamolock](https://godoc.org/cirello.io/dynamolock/v2) | 123 | 47 | 0 | DynamoDB Lock Client for Go | 2018-07-08 11:13:00 | 2023-10-27 19:19:19 |
| [doublejump](https://github.com/edwingeng/doublejump) | 94 | 15 | 0 | A revamped Google's jump consistent hash | 2018-06-26 16:04:50 | 2023-11-18 20:25:20 |
| [semaphore](https://jexia.github.io/semaphore/) | 89 | 15 | 32 | Take control of your data, connect with anything, and expose it anywhere through protocols such as HTTP, GraphQL, and gRPC. | 2020-02-05 16:39:39 | 2023-12-09 19:22:08 |
| [pglock](https://godoc.org/cirello.io/pglock) | 85 | 18 | 0 | PostgreSQL Lock Client for Go | 2018-12-17 17:43:41 | 2023-12-07 18:22:56 |
| [dot](https://github.com/dotchain/dot) | 81 | 8 | 0 | distributed data sync with operational transformation/transforms  | 2017-12-18 01:08:12 | 2023-11-02 14:06:02 |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 72 | 12 | 0 | Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. | 2015-10-10 07:27:33 | 2023-11-18 20:21:10 |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 58 | 15 | 3 | MySQL Backed Locking Primitive | 2020-06-06 16:30:07 | 2023-10-08 03:08:40 |
| [flowgraph](https://emitter.io) | 54 | 8 | 0 | Flowgraph package for scalable asynchronous system development | 2018-08-29 21:45:26 | 2023-11-18 20:28:02 |
| [go-pdu](https://pdu.pub) | 46 | 7 | 0 | Go implementation of PDU - A decentralized SNS backbone | 2018-10-08 08:13:22 | 2023-11-29 11:20:59 |
| [drmaa](https://github.com/dgruber/drmaa) | 46 | 20 | 0 | Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard. | 2013-03-17 12:58:02 | 2023-11-18 20:26:46 |
| [micro](https://github.com/gmsec/micro) | 23 | 7 | 3 | A Go distributed systems development framework | 2020-05-03 01:16:16 | 2023-11-18 20:29:05 |
| [consistenthash](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 22 | 4 | 0 | A Go library that implements Consistent Hashing | 2020-04-22 16:01:25 | 2023-11-21 04:23:00 |
| [dynatomic](https://github.com/tylfin/dynatomic) | 16 | 3 | 0 | Dynatomic is a library for using dynamodb as an atomic counter | 2019-02-08 17:45:14 | 2023-11-18 20:27:09 |
</details>

### Dynamic DNS
Tools for updating dynamic DNS records.

<sup>*Last Update: 2023-12-13 19:34:11*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [godns](https://xiaozhou.net/godns-project-2014-05-18.html) | 1,380 | 255 | 15 | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 2014-05-11 11:49:17 | 2023-12-13 10:23:20 |
| [ddns](https://github.com/skibish/ddns) | 239 | 22 | 1 | Personal DDNS client with Digital Ocean Networking DNS as backend. | 2017-03-13 21:02:27 | 2023-12-12 03:19:03 |
</details>

### Email
Libraries and tools that implement email creation and sending.

<sup>*Last Update: 2023-12-16 21:00:39*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [MailHog](https://github.com/mailhog/MailHog) | 12,860 | 1,015 | 241 | Web and API based SMTP testing | 2014-04-16 22:28:49 | 2023-12-15 20:55:58 |
| [hermes](https://github.com/matcornic/hermes) | 2,712 | 233 | 31 | Golang package that generates clean, responsive HTML e-mails for sending transactional mail | 2017-03-25 18:25:36 | 2023-12-13 17:34:34 |
| [email](https://github.com/jordan-wright/email) | 2,486 | 368 | 56 | Robust and flexible email library for Go | 2013-12-12 20:11:59 | 2023-12-13 06:00:18 |
| [go-imap](https://github.com/emersion/go-imap) | 1,890 | 299 | 39 | 📥 An IMAP library for clients and servers | 2016-04-26 17:59:18 | 2023-12-16 10:52:23 |
| [email-verifier](https://github.com/AfterShip/email-verifier) | 985 | 121 | 9 | :white_check_mark: A Go library for email verification without sending any emails. | 2020-12-18 08:47:28 | 2023-12-16 03:05:44 |
| [sendgrid-go](https://sendgrid.com) | 939 | 306 | 19 | The Official Twilio SendGrid Golang API Library | 2013-09-12 03:31:13 | 2023-12-16 11:31:40 |
| [chasquid](https://blitiri.com.ar/p/chasquid/) | 760 | 57 | 6 | SMTP (email) server with a focus on simplicity, security, and ease of operation [mirror] | 2016-11-03 01:28:05 | 2023-12-14 04:15:59 |
| [mailgun-go](https://mailchain.xyz) | 668 | 147 | 2 | Go library for sending mail with the Mailgun API. | 2014-02-28 00:28:44 | 2023-12-07 18:07:40 |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 548 | 79 | 11 | Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP. | 2019-09-15 05:38:54 | 2023-12-14 13:52:57 |
| [go-message](https://github.com/emersion/go-message) | 340 | 101 | 25 | ✉️ A streaming Go library for the Internet Message Format and mail messages | 2016-12-31 09:31:52 | 2023-12-15 12:26:45 |
| [douceur](https://github.com/aymerick/douceur) | 234 | 42 | 9 | A simple CSS parser and inliner in Go | 2015-04-09 10:21:26 | 2023-11-09 19:22:37 |
| [hectane](https://github.com/hectane/hectane) | 218 | 25 | 16 | Lightweight SMTP client written in Go | 2015-08-28 01:36:47 | 2023-12-03 23:26:32 |
| [mailchain-legacy](https://mailchain.xyz) | 141 | 51 | 44 | Using Mailchain, blockchain users can now send and receive rich-media HTML messages with attachments via a blockchain address. | 2019-04-11 17:37:31 | 2023-09-18 13:20:16 |
| [go-premailer](https://github.com/vanng822/go-premailer) | 124 | 18 | 4 | Inline styling for html mail in golang | 2015-02-16 22:19:18 | 2023-12-05 01:18:50 |
| [go-dkim](https://github.com/toorop/go-dkim) | 92 | 38 | 6 | DKIM package for golang | 2015-04-29 15:38:27 | 2023-10-31 20:02:05 |
| [smtp](https://sendgrid.com) | 74 | 31 | 7 | MailHog SMTP Protocol | 2014-12-24 16:13:49 | 2023-09-17 20:51:33 |
| [go-email-validator](https://github.com/go-email-validator/go-email-validator) | 50 | 11 | 3 | 📧 Golang Email address validator | 2020-12-10 18:27:20 | 2023-11-14 03:54:55 |
</details>

### Embeddable Scripting Languages
Embedding other languages inside your go code.

<sup>*Last Update: 2023-12-16 21:00:26*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 5,783 | 657 | 65 | GopherLua: VM and compiler for Lua in Go | 2015-02-15 13:23:37 | 2023-12-16 09:17:27 |
| [expr](https://expr-lang.org) | 4,830 | 377 | 21 | Expression language and expression evaluation for Go | 2018-07-14 15:57:34 | 2023-12-16 05:58:23 |
| [goja](https://github.com/dop251/goja) | 4,487 | 356 | 27 | ECMAScript/JavaScript engine in pure Go | 2016-11-04 22:04:06 | 2023-12-16 09:41:29 |
| [tengo](https://tengolang.com) | 3,344 | 266 | 74 | A fast script language for Go | 2019-01-09 07:17:17 | 2023-12-15 01:24:01 |
| [go-lua](https://github.com/Shopify/go-lua) | 2,827 | 230 | 45 | A Lua VM in Go | 2013-12-20 17:29:43 | 2023-12-15 19:13:58 |
| [cel-go](https://opensource.google.com/projects/cel) | 1,917 | 248 | 34 | Fast, portable, non-Turing complete expression evaluation with gradual typing (Go) | 2018-03-09 22:57:58 | 2023-12-16 06:48:40 |
| [go-python](https://github.com/sbinet/go-python) | 1,511 | 176 | 27 | naive go bindings to the CPython2 C-API | 2012-07-09 15:43:31 | 2023-12-12 15:26:39 |
| [anko](http://play-anko.appspot.com/) | 1,402 | 128 | 25 | Scriptable interpreter written in golang | 2014-03-28 07:29:40 | 2023-12-15 18:28:39 |
| [go-php](https://github.com/deuill/go-php) | 910 | 149 | 21 | PHP bindings for the Go programming language (Golang) | 2015-09-17 21:23:52 | 2023-12-10 21:48:29 |
| [go-duktape](https://github.com/olebedev/go-duktape) | 779 | 97 | 8 | [abandoned] Duktape JavaScript engine bindings for Go | 2015-01-08 05:09:05 | 2023-12-16 09:45:01 |
| [gval](https://github.com/PaesslerAG/gval) | 675 | 80 | 22 | Expression evaluation in golang | 2017-09-27 08:32:49 | 2023-12-14 16:40:42 |
| [golua](https://github.com/aarzilli/golua) | 631 | 165 | 7 | Go bindings for Lua C API - in progress | 2010-12-06 21:39:53 | 2023-12-11 07:23:49 |
| [gisp](https://docs.gentee.org) | 505 | 37 | 1 | Simple LISP in Go | 2014-01-11 14:05:43 | 2023-12-06 13:55:40 |
| [gentee](https://docs.gentee.org) | 122 | 15 | 1 | Gentee - script programming language for automation. It uses VM and compiler written in Go (Golang). | 2018-01-14 15:49:05 | 2023-12-14 00:13:41 |
| [binder](https://github.com/alexeyco/binder) | 68 | 10 | 2 | High level go to Lua binder. Write less, do more. | 2017-04-02 17:14:52 | 2023-12-15 18:27:05 |
| [purl](https://github.com/ian-kent/purl) | 39 | 5 | 2 | Perl, but fluffy like a cat! | 2014-11-29 19:06:01 | 2023-10-28 05:24:45 |
| [ecal](https://github.com/krotik/ecal) | 37 | 4 | 0 | A simple embeddable scripting language which supports concurrent event processing. | 2020-11-30 15:58:56 | 2023-12-11 21:24:53 |
| [ngaro](https://github.com/db47h/ngaro) | 25 | 3 | 1 | An embeddable implementation of the Ngaro Virtual Machine for Go programs | 2016-08-09 15:23:50 | 2023-05-09 02:52:15 |
</details>

### Error Handling
Libraries for handling errors.

<sup>*Last Update: 2023-12-07 21:47:47*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [errors](https://godoc.org/github.com/pkg/errors) | 8,108 | 689 | 42 | Simple error handling primitives | 2015-12-27 12:05:38 | 2023-12-07 08:02:17 |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 2,075 | 161 | 28 | A Go (golang) package for representing a list of errors as a single error. | 2014-12-15 20:12:26 | 2023-12-07 08:16:58 |
| [eris](https://pkg.go.dev/github.com/rotisserie/eris) | 1,363 | 50 | 4 | Error handling library with readable stack traces and flexible formatting support 🎆 | 2019-09-07 16:50:33 | 2023-12-07 09:46:44 |
| [errorx](https://github.com/joomcode/errorx) | 1,036 | 31 | 6 | A comprehensive error handling library for Go | 2018-08-17 08:02:10 | 2023-12-07 12:06:25 |
| [tracerr](https://github.com/ztrue/tracerr) | 847 | 36 | 2 | Golang errors with stack trace and source fragments. | 2019-02-06 18:57:46 | 2023-12-06 15:27:35 |
| [errlog](https://github.com/snwfdhmp/errlog) | 451 | 21 | 0 | Reduce debugging time while programming Go. Use static and stack-trace analysis to determine which func call causes the error. | 2019-02-16 23:19:05 | 2023-11-24 19:47:42 |
| [emperror](https://emperror.dev/emperror) | 316 | 20 | 6 | The Emperor takes care of all errors personally | 2017-06-13 00:24:28 | 2023-11-26 22:45:23 |
| [errors](https://emperror.dev/errors) | 190 | 11 | 10 | Drop-in replacement for the standard library errors package and github.com/pkg/errors | 2019-07-09 13:02:52 | 2023-11-09 07:06:55 |
| [errors](https://github.com/bnkamalesh/errors) | 59 | 5 | 0 | A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types. | 2020-07-17 18:57:04 | 2023-11-26 13:22:22 |
| [falcon](https://github.com/ThundR67/falcon) | 10 | 1 | 0 | A Simple Yet Highly Powerful Package For Error Handling | 2019-09-09 12:49:43 | 2023-11-10 02:18:16 |
| [errors](https://github.com/PumpkinSeed/errors) | 7 | 1 | 0 | Simple and efficient error package  | 2020-01-08 21:12:51 | 2023-08-29 00:07:32 |
| [errors](https://github.com/neuronlabs/errors) | 6 | 1 | 0 | Simple golang error handling with classification primitives. | 2019-07-26 00:15:36 | 2023-09-26 23:05:27 |
</details>

### File Handling
Libraries for handling files and file systems.

<sup>*Last Update: 2023-11-03 08:17:34*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [pdfcpu](http://pdfcpu.io/) | 5,484 | 412 | 67 | A PDF processor written in Go. | 2017-06-18 17:27:38 | 2023-11-02 19:32:18 |
| [afero](https://github.com/spf13/afero) | 5,438 | 524 | 124 | A FileSystem Abstraction System for Go | 2014-10-28 14:19:05 | 2023-11-02 19:52:52 |
| [notify](https://github.com/rjeczalik/notify) | 856 | 128 | 47 | File system event notification library on steroids. | 2014-09-08 16:09:34 | 2023-10-25 09:08:39 |
| [copy](https://pkg.go.dev/github.com/otiai10/copy) | 607 | 154 | 19 | Go copy directory recursively | 2017-09-01 03:18:56 | 2023-11-02 14:11:17 |
| [afs](https://github.com/viant/afs) | 266 | 32 | 3 | Abstract File Storage | 2019-08-19 18:43:38 | 2023-11-02 01:58:49 |
| [vfs](https://github.com/C2FO/vfs) | 256 | 28 | 8 | Pluggable, extensible virtual file system for Go | 2017-08-01 18:06:14 | 2023-10-28 05:45:20 |
| [bigfile](https://learnku.com/docs/bigfile) | 243 | 47 | 3 | Bigfile -- a file transfer system that supports http, rpc and ftp protocol   https://bigfile.site   | 2019-07-15 10:43:50 | 2023-10-26 02:52:10 |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 193 | 41 | 6 | Golang wrapper for Exiftool : extract as much metadata as possible (EXIF, ...) from files (pictures, pdf, office documents, ...) | 2019-05-12 20:34:09 | 2023-11-01 01:04:57 |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 108 | 31 | 0 | Read csv file from go using tags | 2017-06-18 15:31:16 | 2023-07-24 03:28:19 |
| [skywalker](https://github.com/dixonwille/skywalker) | 95 | 5 | 1 | A package to allow one to concurrently go through a filesystem with ease | 2017-08-01 20:08:25 | 2023-10-24 18:02:45 |
| [checksum](https://github.com/codingsince1985/checksum) | 94 | 17 | 0 | Compute message digest for large files in Go | 2014-11-05 09:37:00 | 2023-11-02 01:03:14 |
| [parquet](https://github.com/parsyl/parquet) | 85 | 12 | 1 | A library for reading and writing parquet files. | 2019-01-29 21:52:30 | 2023-08-17 11:05:19 |
| [opc](https://github.com/qmuntal/opc) | 75 | 7 | 0 | Go implementation of the Open Packaging Conventions (OPC) | 2018-11-06 14:49:06 | 2023-09-11 17:12:49 |
| [tarfs](https://github.com/posener/tarfs) | 56 | 9 | 1 | An implementation of the FileSystem interface for tar files. | 2017-03-10 22:13:11 | 2023-06-25 16:38:44 |
| [baraka](https://github.com/xis/baraka) | 51 | 8 | 2 | a tool for handling file uploads simple | 2020-07-12 21:56:50 | 2023-10-24 14:50:08 |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 37 | 24 | 0 | Load GTFS files in golang | 2017-07-09 09:30:31 | 2023-07-29 18:17:44 |
| [flop](https://github.com/homedepot/flop) | 34 | 12 | 0 | Go file operations library chasing GNU APIs. | 2019-03-01 13:41:39 | 2023-06-26 22:45:43 |
| [gut](https://github.com/1set/gut) | 27 | 9 | 13 | 🍱 yet another collection of go utilities & tools | 2019-10-05 23:47:24 | 2023-09-05 07:27:56 |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 21 | 8 | 1 | copy files for humans | 2018-10-16 07:08:24 | 2023-10-30 02:09:25 |
| [todotxt](https://github.com/1set/todotxt) | 19 | 6 | 1 | Parser for todo.txt files in Go ✅ | 2020-11-06 17:41:59 | 2023-04-14 11:43:41 |
| [higgs](https://github.com/dastoori/higgs) | 18 | 4 | 0 | A tiny cross-platform Go library to hide/unhide files and directories | 2020-12-13 18:33:10 | 2023-10-22 08:05:01 |
</details>

### Financial
Packages for accounting and finance.

<sup>*Last Update: 2023-12-10 21:12:40*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [decimal](https://github.com/shopspring/decimal) | 5,516 | 606 | 125 | Arbitrary-precision fixed-point decimal numbers in go | 2015-02-25 20:12:57 | 2023-12-10 10:45:03 |
| [go-money](https://github.com/Rhymond/go-money) | 1,422 | 138 | 28 | Go implementation of Fowler's Money pattern | 2017-03-20 16:23:54 | 2023-12-09 19:03:42 |
| [accounting](https://github.com/leekchan/accounting) | 842 | 73 | 13 | money and currency formatting for golang | 2015-08-10 13:23:56 | 2023-12-10 03:05:18 |
| [techan](https://godoc.org/github.com/sdcoffey/techan) | 760 | 137 | 21 | Technical Analysis Library for Golang | 2017-03-08 03:04:08 | 2023-12-10 14:07:21 |
| [go-finance](https://github.com/FlashBoys/go-finance) | 534 | 49 | 4 | :warning: Deprecrated in favor of https://github.com/piquette/finance-go  | 2016-02-28 00:37:46 | 2021-06-23 04:00:30 |
| [currency](https://pkg.go.dev/github.com/bojanz/currency) | 457 | 37 | 1 | Currency handling for Go. | 2020-04-16 15:34:39 | 2023-12-10 00:09:36 |
| [orderbook](https://github.com/i25959341/orderbook) | 374 | 134 | 5 | Matching Engine for Limit Order Book in Golang | 2018-04-24 18:05:26 | 2023-12-10 06:34:45 |
| [go-finance](https://github.com/alpeb/go-finance) | 159 | 24 | 0 | Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. | 2017-06-01 15:58:33 | 2023-12-09 19:26:38 |
| [sleet](https://github.com/BoltApp/sleet) | 128 | 20 | 10 | Payment abstraction library - one interface for multiple payment processors ( inspired by Ruby's ActiveMerchant ) | 2019-11-13 21:56:58 | 2023-12-09 19:38:11 |
| [transaction](https://github.com/claygod/transaction) | 126 | 17 | 0 | Embedded database for accounts transactions. | 2017-10-11 13:50:30 | 2023-10-01 12:07:52 |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 125 | 26 | 0 | Golang library for querying and parsing OFX | 2015-11-08 13:56:53 | 2023-11-16 23:27:27 |
| [vat](https://github.com/dannyvankooten/vat) | 106 | 16 | 3 | Go package for dealing with EU VAT. Does VAT number validation & rates retrieval. | 2016-06-18 16:10:09 | 2023-11-21 10:42:58 |
| [go-finnhub](https://github.com/m1/go-finnhub) | 84 | 17 | 0 | Simple and easy to use client for stock market, forex and crypto data from finnhub.io written in Go. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges | 2020-01-13 20:47:13 | 2023-12-08 08:24:00 |
| [currency](https://github.com/bnkamalesh/currency) | 56 | 8 | 0 | A currency computations package. | 2017-05-09 06:06:38 | 2023-10-25 22:21:23 |
| [fastme](https://github.com/newity/fastme) | 24 | 4 | 0 | Arbitrary-precision fixed-point decimal numbers in go | 2020-10-29 13:57:10 | 2021-08-10 04:48:11 |
| [go-finance](https://www.yellowduck.be) | 23 | 5 | 1 | Finance related Go functions (e.g. exchange rates, VAT number checking, …) | 2019-09-30 06:49:07 | 2023-12-09 19:29:51 |
</details>

### Forms
Libraries for working with forms.

<sup>*Last Update: 2023-12-10 21:12:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [nosurf](http://godoc.org/github.com/justinas/nosurf) | 1,458 | 152 | 13 | CSRF protection middleware for Go. | 2013-08-22 17:47:34 | 2023-12-06 06:12:04 |
| [csrf](https://gorilla.github.io) | 952 | 155 | 4 | Package gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services 🔒 | 2015-08-03 00:35:16 | 2023-12-09 19:40:58 |
| [binding](http://mholt.github.io/binding) | 796 | 84 | 8 | Reflectionless data binding for Go's net/http (not actively maintained) | 2014-05-20 23:35:14 | 2023-11-29 20:49:06 |
| [form](https://github.com/go-playground/form) | 651 | 38 | 9 | :steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. | 2016-05-26 13:26:40 | 2023-12-09 04:11:00 |
| [conform](https://github.com/leebenson/conform) | 312 | 38 | 5 | Trims, sanitizes & scrubs data based on struct tags (go, golang) | 2016-01-05 18:00:06 | 2023-10-19 03:22:22 |
| [formam](https://github.com/monoculum/formam) | 182 | 17 | 2 | a package for decode form's values into struct in Go | 2014-10-25 00:23:30 | 2023-10-18 14:05:13 |
| [forms](https://github.com/albrow/forms) | 133 | 24 | 2 | A lightweight go library for parsing form data or json from an http.Request. | 2014-08-07 16:11:30 | 2023-10-12 17:02:12 |
| [qs](https://github.com/sonh/qs) | 68 | 4 | 0 | Go module for encoding structs into URL query parameters | 2020-10-02 09:50:35 | 2023-10-18 09:53:48 |
| [bind](https://github.com/robfig/bind) | 30 | 6 | 0 |  | 2014-08-06 00:13:10 | 2023-10-12 21:12:33 |
| [queryparam](https://github.com/TomWright/queryparam) | 19 | 5 | 0 | Go package to easily convert a URL's query parameters/values into usable struct values of the correct types. | 2018-06-14 10:23:05 | 2023-10-29 08:49:21 |
</details>

### Functional
Packages to support functional programming in Go.

<sup>*Last Update: 2023-12-05 20:35:00*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1,294 | 68 | 4 |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  | 2014-07-02 10:27:16 | 2023-11-26 19:02:02 |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 332 | 21 | 0 | Monad, Functional Programming features for Golang | 2018-05-24 09:08:45 | 2023-11-19 04:50:55 |
| [gofp](https://github.com/rbrahul/gofp) | 143 | 8 | 0 | A super simple Lodash like utility library with essential functions that empowers the development in Go | 2021-02-19 00:01:39 | 2023-11-23 07:46:41 |
| [fuego](https://github.com/seborama/fuego) | 141 | 12 | 0 | Functional Experiment in Golang | 2018-11-05 22:24:09 | 2023-08-02 00:40:24 |
</details>

### GUI - Interaction
Libraries for building GUI Applications.

<sup>*Last Update: 2023-12-10 21:13:05*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [robotgo](https://github.com/go-vgo/robotgo) | 8,866 | 872 | 146 | RobotGo, Go Native cross-platform RPA and GUI automation  @vcaesar | 2016-09-26 16:26:56 | 2023-12-10 13:14:42 |
| [systray](https://github.com/getlantern/systray) | 2,968 | 393 | 98 | a cross platfrom Go library to place an icon and menu in the notification area | 2014-11-12 03:41:57 | 2023-12-10 01:19:20 |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 581 | 51 | 10 | gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher | 2013-11-25 06:35:16 | 2023-11-26 03:02:17 |
| [trayhost](https://github.com/shurcooL/trayhost) | 246 | 19 | 9 | Cross-platform Go library to place an icon in the host operating system's taskbar. | 2014-04-25 04:02:30 | 2023-12-08 06:49:13 |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 29 | 6 | 0 | macOS Sleep/ Wake notifications in golang | 2019-03-30 00:43:23 | 2023-06-26 11:59:40 |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 24 | 9 | 2 | A library to notify about any (pluggable) activity on your machine, and let you take action as needed | 2019-03-12 21:16:44 | 2023-10-13 21:34:14 |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 19 | 7 | 1 | :traffic_light: Go bindings for libappindicator3 C library | 2019-05-04 13:38:53 | 2023-05-26 20:30:18 |
</details>

### GUI - Toolkits
Libraries for building GUI Applications.

<sup>*Last Update: 2023-11-03 08:17:54*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fyne](https://fyne.io/) | 21,712 | 1,260 | 596 | Cross platform GUI toolkit in Go inspired by Material Design | 2018-02-04 22:07:16 | 2023-11-03 00:41:02 |
| [wails](https://wails.io) | 18,708 | 935 | 193 | Create beautiful applications using Go | 2018-12-15 23:14:06 | 2023-11-03 00:12:38 |
| [webview](https://webview.dev) | 11,605 | 928 | 206 | Tiny cross-platform webview library for C/C++. Uses WebKit (GTK/Cocoa) and Edge WebView2 (Windows). | 2017-08-19 08:26:00 | 2023-11-02 22:16:14 |
| [qt](https://github.com/therecipe/qt) | 9,998 | 737 | 369 | Qt binding for Go (Golang) with support for Windows / macOS / Linux / FreeBSD / Android / iOS / Sailfish OS / Raspberry Pi / AsteroidOS / Ubuntu Touch / JavaScript / WebAssembly | 2014-11-19 00:03:08 | 2023-11-02 06:56:53 |
| [ui](https://github.com/andlabs/ui) | 8,294 | 790 | 124 | Platform-native GUI library for Go. | 2014-02-17 23:44:00 | 2023-11-01 15:11:38 |
| [go-app](https://go-app.dev) | 7,378 | 391 | 118 | A package to build progressive web apps with Go programming language and WebAssembly. | 2016-10-12 00:31:33 | 2023-11-02 05:59:39 |
| [walk](https://github.com/lxn/walk) | 6,568 | 915 | 346 | A Windows GUI toolkit for the Go Programming Language | 2010-09-16 08:11:49 | 2023-10-31 20:31:51 |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 4,849 | 384 | 53 | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron) | 2017-04-22 07:59:15 | 2023-11-02 19:08:46 |
| [go-sciter](https://sciter.com) | 2,543 | 301 | 95 | Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development | 2015-10-15 12:41:06 | 2023-10-29 16:18:43 |
| [go-gtk](http://mattn.github.com/go-gtk) | 2,017 | 291 | 73 | Go binding for GTK | 2009-11-26 16:58:53 | 2023-10-30 02:50:00 |
| [gotk3](https://github.com/gotk3/gotk3) | 1,965 | 269 | 104 | Go bindings for GTK3 | 2015-08-13 13:09:46 | 2023-11-01 09:42:00 |
| [gowd](https://github.com/dtylman/gowd) | 411 | 42 | 7 | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by nwjs) | 2017-03-29 14:50:53 | 2023-10-25 17:09:32 |
</details>

### Game Development
Awesome game development libraries.

<sup>*Last Update: 2023-12-10 21:12:04*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ebiten](https://ebitengine.org) | 9,171 | 611 | 253 | Ebitengine - A dead simple 2D game engine for Go | 2013-06-16 15:13:01 | 2023-12-10 12:53:39 |
| [leaf](https://github.com/name5566/leaf) | 5,027 | 1,328 | 21 | A game server framework in Go (golang) | 2014-08-04 12:40:08 | 2023-12-08 14:21:34 |
| [pixel](https://github.com/faiface/pixel) | 4,355 | 248 | 46 | A hand-crafted 2D game library in Go | 2016-11-19 11:15:34 | 2023-12-10 04:08:22 |
| [nano](https://github.com/lonng/nano) | 2,557 | 411 | 27 | Lightweight, facility, high performance golang based game server framework | 2017-08-02 06:05:14 | 2023-12-07 07:32:43 |
| [engine](https://discord.gg/NfaeVr8zDg) | 2,520 | 261 | 44 | Go 3D Game Engine (http://g3n.rocks) | 2017-03-07 18:25:09 | 2023-12-08 16:59:48 |
| [goworld](https://github.com/xiaonanln/goworld) | 2,435 | 476 | 27 | Scalable Distributed Game Server Engine with Hot Swapping in Golang | 2017-06-03 15:02:46 | 2023-12-08 07:04:40 |
| [go-sdl2](https://godoc.org/github.com/veandco/go-sdl2) | 2,056 | 267 | 79 | SDL2 binding for Go | 2013-06-05 18:30:03 | 2023-12-10 01:15:20 |
| [pitaya](https://github.com/topfreegames/pitaya) | 2,033 | 427 | 49 | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. | 2018-03-19 19:40:36 | 2023-12-08 06:15:14 |
| [engo](https://engoengine.github.io) | 1,661 | 135 | 50 | Engo is an open-source 2D game engine written in Go. | 2014-11-12 05:50:03 | 2023-12-05 08:40:39 |
| [oak](https://github.com/oakmound/oak) | 1,462 | 85 | 17 | A pure Go game engine | 2017-07-15 16:24:27 | 2023-12-07 05:13:17 |
| [termloop](https://github.com/JoelOtter/termloop) | 1,373 | 86 | 5 | Terminal-based game engine for Go, built on top of Termbox | 2015-05-23 17:12:34 | 2023-12-05 18:06:59 |
| [gonet](https://github.com/xtaci/gonet) | 1,239 | 347 | 0 | A Game Server Skeleton in golang. | 2013-04-11 02:18:23 | 2023-11-29 09:03:12 |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 1,142 | 123 | 2 | Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming. | 2017-01-27 08:31:45 | 2023-12-09 18:55:16 |
| [engine](https://azul3d.org) | 592 | 57 | 82 | Azul3D - A 3D game engine written in Go! | 2016-02-29 04:54:44 | 2023-12-06 03:51:57 |
| [go-astar](https://discord.gg/NfaeVr8zDg) | 566 | 80 | 2 | Go implementation of the A* search algorithm | 2014-05-28 02:00:03 | 2023-11-27 07:14:46 |
| [go3d](https://github.com/ungerik/go3d) | 285 | 48 | 2 | A performance oriented 2D/3D math package for Go | 2011-06-27 13:02:26 | 2023-11-23 08:20:18 |
| [tile](https://github.com/kelindar/tile) | 130 | 15 | 2 | Tile is a 2D grid engine, built with data and cache friendly ways, includes pathfinding and observers. | 2020-08-19 13:23:18 | 2023-11-28 00:17:52 |
| [prototype](https://pkg.go.dev/github.com/gonutz/prototype/draw) | 83 | 9 | 1 | Simple 2D game prototyping framework. | 2015-03-04 09:24:39 | 2023-11-20 13:01:18 |
</details>

### Generation and Generics
Tools to enhance the language with features like generics via code generation.

<sup>*Last Update: 2023-12-06 21:07:43*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-linq](https://godoc.org/github.com/ahmetb/go-linq) | 3,384 | 280 | 11 | .NET LINQ capabilities in Go | 2013-12-19 03:05:00 | 2023-12-02 13:41:39 |
| [jennifer](https://github.com/dave/jennifer) | 3,080 | 158 | 16 | Jennifer is a code generator for Go | 2016-12-04 20:57:38 | 2023-12-05 23:01:39 |
| [gen](http://clipperhouse.com/gen/overview/) | 1,427 | 90 | 32 | Type-driven code generation for Go | 2013-10-13 20:26:36 | 2023-12-04 22:16:26 |
| [goderive](https://github.com/awalterschulze/goderive) | 1,135 | 44 | 19 | Derives and generates mundane golang functions that you do not want to maintain yourself | 2017-02-10 21:46:49 | 2023-12-04 05:17:10 |
| [gowrap](https://github.com/hexdigest/gowrap) | 820 | 77 | 13 | GoWrap is a command line tool for generating decorators for Go interfaces | 2018-09-15 09:20:42 | 2023-12-02 05:29:22 |
| [go-enum](https://github.com/abice/go-enum) | 579 | 54 | 9 | An enum generator for go | 2017-08-10 22:07:31 | 2023-12-04 12:15:54 |
| [interfaces](https://github.com/rjeczalik/interfaces) | 403 | 31 | 13 | Code generation tools for Go. | 2015-12-06 00:04:50 | 2023-10-07 13:07:02 |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 105 | 16 | 0 | A Go preprocessor for package scoped reflection | 2012-09-03 07:53:00 | 2023-10-22 14:44:20 |
| [gotype](https://github.com/wzshiming/gotype) | 57 | 9 | 1 | Golang source code parsing, usage like reflect package | 2017-12-05 04:09:47 | 2023-11-10 03:15:31 |
| [efaceconv](https://github.com/vlegio/efaceconv) | 48 | 10 | 1 |  | 2016-11-18 11:38:54 | 2023-04-01 15:43:35 |
| [GENERIS](https://github.com/senselogic/GENERIS) | 43 | 1 | 0 | Versatile Go code generator. | 2019-03-10 19:33:31 | 2023-10-22 07:37:08 |
| [go-xray](https://godoc.org/github.com/ahmetb/go-linq) | 26 | 3 | 1 | Helpers for making the use of reflection easier | 2019-10-01 08:40:51 | 2023-07-13 02:04:15 |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 23 | 1 | 0 | create type dynamically in Golang | 2020-01-14 15:50:38 | 2023-09-30 04:55:03 |
</details>

### Geographic
Geographic tools and servers

<sup>*Last Update: 2023-12-13 19:34:07*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tile38](https://tile38.com) | 8,773 | 551 | 133 | Real-time Geospatial and Geofencing | 2016-03-04 23:07:44 | 2023-12-12 22:39:16 |
| [geo](https://github.com/golang/geo) | 1,598 | 231 | 22 | S2 geometry library in Go | 2014-12-03 23:02:15 | 2023-12-09 19:51:34 |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 563 | 90 | 14 | Basic Go server for mbtiles | 2014-11-01 04:12:14 | 2023-12-13 09:13:28 |
| [osm](https://github.com/paulmach/osm) | 310 | 43 | 4 | General purpose library for reading, writing and working with OpenStreetMap data | 2016-02-02 00:59:03 | 2023-12-12 18:24:31 |
| [wgs84](https://github.com/wroge/wgs84) | 107 | 12 | 4 | A zero-dependency Go package for coordinate transformations. | 2019-06-08 17:17:59 | 2023-12-09 19:53:01 |
| [geoserver](https://github.com/hishamkaram/geoserver) | 85 | 22 | 4 | geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API. | 2018-03-26 21:36:49 | 2023-11-02 02:42:49 |
| [gismanager](https://github.com/hishamkaram/gismanager) | 50 | 10 | 1 | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver | 2018-09-29 12:51:37 | 2023-04-14 15:44:10 |
| [pbf](https://github.com/maguro/pbf) | 46 | 7 | 3 | OpenStreetMap PBF golang parser | 2017-09-18 23:13:18 | 2023-11-24 22:56:40 |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 26 | 9 | 2 | Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels | 2020-03-27 09:47:32 | 2023-12-09 19:51:12 |
</details>

### Go Compilers
Tools for compiling Go to other languages.

<sup>*Last Update: 2023-12-18 21:43:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 12,217 | 616 | 191 | A compiler from Go to JavaScript for running Go code in a browser | 2013-08-27 22:23:58 | 2023-12-17 23:39:47 |
| [tardisgo](http://tardisgo.github.io) | 428 | 33 | 5 | Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   | 2014-01-08 11:07:33 | 2023-12-09 18:57:18 |
| [c4go](https://github.com/Konstantin8105/c4go) | 343 | 38 | 25 | Transpiling C code to Go code | 2018-03-28 06:24:57 | 2023-12-08 07:15:18 |
| [f4go](https://github.com/Konstantin8105/f4go) | 40 | 10 | 7 | Transpiling fortran code to golang code | 2018-07-08 17:05:43 | 2023-11-08 17:06:55 |
</details>

### Goroutines
Tools for managing and working with Goroutines.

<sup>*Last Update: 2023-10-31 07:50:50*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ants](https://ants.andypan.me) | 11,144 | 1,268 | 23 | 🐜🐜🐜 ants is a high-performance and low-cost goroutine pool in Go./ ants 是一个高性能且低损耗的 goroutine 池。 | 2018-05-19 01:13:38 | 2023-10-30 18:13:44 |
| [tunny](https://github.com/Jeffail/tunny) | 3,701 | 333 | 9 | A goroutine pool for Go | 2014-04-02 16:14:58 | 2023-10-27 10:07:35 |
| [goworker](https://www.goworker.org) | 2,759 | 280 | 34 | goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers. | 2013-07-22 17:04:27 | 2023-10-23 11:05:11 |
| [workerpool](https://github.com/gammazero/workerpool) | 1,164 | 130 | 16 | Concurrency limiting goroutine pool | 2016-05-17 14:32:06 | 2023-10-29 15:00:03 |
| [pond](https://github.com/alitto/pond) | 1,015 | 51 | 4 | 🔘 Minimalistic and High-performance goroutine worker pool written in Go | 2020-03-21 14:56:33 | 2023-10-31 00:44:38 |
| [grpool](https://godoc.org/github.com/ivpusic/grpool) | 736 | 101 | 5 | Lightweight Goroutine pool | 2015-07-22 00:15:04 | 2023-10-23 11:10:32 |
| [pool](https://github.com/go-playground/pool) | 721 | 67 | 4 | :speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation | 2015-10-28 16:36:08 | 2023-10-23 11:11:39 |
| [gowp](https://xxjwxc.github.io/post/gowp/) | 482 | 72 | 2 | golang worker pool , Concurrency limiting goroutine pool | 2019-09-14 11:43:50 | 2023-10-09 14:43:37 |
| [go-floc](https://github.com/workanator/go-floc) | 263 | 18 | 1 | Floc: Orchestrate goroutines with ease. | 2017-07-03 07:34:06 | 2023-09-19 02:47:38 |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 215 | 26 | 1 | Simply way to control goroutines execution order based on dependencies | 2016-09-25 14:46:09 | 2023-10-09 06:51:34 |
| [artifex](https://github.com/mborders/artifex) | 178 | 13 | 1 | Simple in-memory job queue for Golang using worker-based dispatching | 2018-10-31 19:34:31 | 2023-10-25 05:34:10 |
| [semaphore](https://github.com/marusama/semaphore) | 162 | 12 | 0 | Fast resizable golang semaphore primitive | 2017-11-22 14:00:58 | 2023-10-19 01:20:56 |
| [go-workers](https://github.com/catmullet/go-workers/wiki/Getting-Started) | 160 | 16 | 3 | 👷 Library for safely running groups of workers concurrently or consecutively that require input and output through channels | 2020-10-06 15:39:43 | 2023-08-28 07:01:24 |
| [errgroup](https://github.com/neilotoole/errgroup) | 152 | 11 | 6 | errgroup with goroutine worker limits | 2020-06-26 06:07:39 | 2023-10-22 19:01:16 |
| [async](https://pkg.go.dev/github.com/reugn/async) | 151 | 7 | 0 | Synchronization and asynchronous computation package for Go | 2019-12-28 09:48:40 | 2023-10-28 16:12:53 |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 132 | 15 | 0 | CyclicBarrier golang implementation | 2018-01-11 10:38:46 | 2023-10-03 23:30:17 |
| [async](https://github.com/StudioSol/async) | 126 | 17 | 2 | A safe way to execute functions asynchronously, recovering them in case of panic. It also provides an error stack aiming to facilitate fail causes discovery. | 2017-06-30 17:08:33 | 2023-09-18 08:40:34 |
| [gollback](https://github.com/vardius/gollback) | 111 | 12 | 0 | Go asynchronous simple function utilities, for managing execution of closures and callbacks | 2019-05-11 05:56:37 | 2023-09-09 00:06:05 |
| [threadpool](https://github.com/shettyh/threadpool) | 97 | 20 | 2 | Golang simple thread pool implementation | 2017-09-06 18:45:39 | 2023-10-02 00:17:21 |
| [breaker](https://pkg.go.dev/github.com/kamilsk/breaker) | 97 | 5 | 7 | 🚧 Flexible mechanism to make execution flow interruptible. | 2019-02-15 15:08:24 | 2021-06-27 05:54:10 |
| [semaphore](https://github.com/kamilsk/semaphore) | 95 | 12 | 6 | 🚦 Semaphore pattern implementation with timeout of lock/unlock operations. | 2016-10-08 11:48:12 | 2023-09-22 04:25:26 |
| [Hunch](https://github.com/AaronJan/Hunch) | 94 | 9 | 1 | Hunch provides functions like: All, First, Retry, Waterfall etc., that makes asynchronous flow control more intuitive. | 2019-06-05 13:21:04 | 2023-09-26 14:28:03 |
| [worker-pool](https://github.com/vardius/worker-pool) | 88 | 13 | 0 | Go simple async worker pool | 2017-10-04 09:18:31 | 2023-10-26 12:04:35 |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 87 | 4 | 0 | gpool - a generic context-aware resizable goroutines pool to bound concurrency based on semaphore.  | 2018-12-03 04:23:35 | 2023-10-02 00:17:41 |
| [goccm](https://github.com/zenthangplus/goccm) | 66 | 15 | 3 | Limits the number of goroutines that are allowed to run concurrently | 2019-08-16 02:26:53 | 2023-08-01 10:49:14 |
| [nursery](https://github.com/arunsworld/nursery) | 62 | 6 | 1 | Structured Concurrency in Go | 2019-11-23 19:26:02 | 2023-09-21 14:01:47 |
| [gowl](https://github.com/hmdsefi/gowl) | 59 | 8 | 5 | Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. | 2021-04-12 19:15:53 | 2023-08-02 15:56:03 |
| [routine](https://github.com/x-mod/routine) | 59 | 7 | 0 | go routine control, abstraction of the Main and some useful Executors.如果你不会管理Goroutine的话，用它 | 2019-03-04 12:25:23 | 2023-09-22 01:26:15 |
| [gohive](https://github.com/loveleshsharma/gohive) | 44 | 6 | 3 | 🐝 A Highly Performant and easy to use goroutine pool for Go | 2019-05-31 10:44:24 | 2023-07-05 18:08:45 |
| [kyoo](https://github.com/dirkaholic/kyoo) | 43 | 4 | 0 | Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries | 2020-01-06 20:35:11 | 2023-09-25 12:13:04 |
| [go-waitgroup](https://www.yellowduck.be) | 42 | 4 | 0 | A sync.WaitGroup with error handling and concurrency control | 2018-08-08 16:12:35 | 2023-08-16 06:25:52 |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 35 | 2 | 0 | Run functions in parallel :comet: | 2017-06-18 09:47:54 | 2023-03-19 07:23:27 |
| [oversight](https://cirello.io/oversight) | 35 | 5 | 0 | Erlang-like supervisor trees | 2018-11-09 14:46:48 | 2023-08-20 12:45:27 |
| [go-trylock](https://github.com/subchen/go-trylock) | 33 | 9 | 1 | TryLock support on read-write lock for Golang | 2018-04-26 06:02:47 | 2023-02-11 00:46:09 |
| [stl](https://github.com/ssgreg/stl) | 27 | 5 | 0 | Software Transactional Locks | 2018-06-19 10:50:11 | 2023-07-18 16:30:19 |
| [channelify](https://github.com/ddelizia/channelify) | 27 | 3 | 1 | Make functions return a channel for parallel processing via go routines. | 2020-10-05 13:12:48 | 2023-02-04 17:11:21 |
| [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) | 17 | 4 | 0 | Make functions return a channel for parallel processing via go routines. | 2020-11-22 02:35:52 | 2023-10-23 12:50:10 |
| [breaker](https://pkg.go.dev/github.com/kamilsk/breaker) | 16 | 1 | 0 | 🚧 Flexible mechanism to make execution flow interruptible. | 2021-07-11 10:35:18 | 2023-08-21 09:36:07 |
| [queue](https://github.com/AnikHasibul/queue) | 15 | 2 | 0 | package queue gives you a queue group accessibility. Helps you to limit goroutines, wait for the end of the all goroutines and much more. | 2018-12-21 07:15:00 | 2023-10-24 04:24:03 |
| [conexec](https://github.com/ITcathyh/conexec) | 13 | 2 | 0 | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking. | 2019-12-24 07:35:11 | 2022-09-26 10:34:16 |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 12 | 3 | 0 | A collection of tools for Golang | 2018-11-14 02:53:08 | 2023-06-11 03:02:08 |
| [hands](https://github.com/duanckham/hands) | 10 | 3 | 1 | Hands is a process controller used to control the execution and return strategies of multiple goroutines. | 2020-04-04 11:04:11 | 2022-10-10 08:07:01 |
</details>

### Images
Libraries for manipulating images.

<sup>*Last Update: 2023-11-09 08:25:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocv](https://gocv.io) | 5,849 | 867 | 289 | Go package for computer vision using OpenCV 4 and beyond. | 2017-09-18 21:54:17 | 2023-11-08 14:24:45 |
| [imaginary](https://fly.io/docs/app-guides/run-a-global-image-service/) | 5,153 | 440 | 118 | Fast, simple, scalable, Docker-ready HTTP microservice for high-level image processing | 2015-03-04 18:51:40 | 2023-11-08 17:34:46 |
| [imaging](https://github.com/disintegration/imaging) | 4,878 | 402 | 28 | Imaging is a simple image processing package for Go | 2012-12-06 20:21:21 | 2023-11-08 00:36:08 |
| [gg](https://godoc.org/github.com/fogleman/gg) | 4,130 | 391 | 88 | Go Graphics - 2D rendering in Go with a simple API. | 2016-02-18 21:05:08 | 2023-11-09 00:49:32 |
| [bild](https://github.com/anthonynsimon/bild) | 3,853 | 265 | 19 | Image processing algorithms in pure Go | 2016-08-01 15:54:29 | 2023-11-07 08:53:09 |
| [ln](https://godoc.org/github.com/fogleman/ln/ln) | 3,206 | 129 | 12 | 3D line art engine. | 2016-01-10 04:28:10 | 2023-11-07 21:48:44 |
| [resize](https://github.com/nfnt/resize) | 2,970 | 318 | 12 | Pure golang image resizing  | 2012-08-02 19:48:26 | 2023-11-07 07:57:52 |
| [bimg](https://pkg.go.dev/github.com/h2non/bimg?tab=doc) | 2,404 | 341 | 164 | Go package for fast high-level image processing powered by libvips C library | 2015-03-17 14:14:02 | 2023-11-06 23:52:45 |
| [pt](http://bit.ly/1E7rSoi) | 2,055 | 121 | 9 | A path tracer written in Go. | 2015-01-23 19:39:29 | 2023-10-24 14:35:47 |
| [svgo](https://github.com/ajstarks/svgo) | 2,053 | 172 | 12 | Go Language Library for SVG generation | 2010-03-05 23:24:10 | 2023-11-07 12:33:39 |
| [picfit](http://bit.ly/1E7rSoi) | 1,909 | 172 | 12 | An image resizing server written in Go | 2014-12-06 17:30:45 | 2023-11-08 23:33:39 |
| [smartcrop](https://github.com/muesli/smartcrop) | 1,764 | 114 | 8 | smartcrop finds good image crops for arbitrary crop sizes | 2014-04-07 22:40:03 | 2023-11-07 10:17:32 |
| [gift](https://godoc.org/github.com/fogleman/gg) | 1,684 | 170 | 4 | Go Image Filtering Toolkit | 2014-07-12 18:47:40 | 2023-11-08 13:51:16 |
| [imagick](https://godoc.org/github.com/gographics/imagick/imagick) | 1,635 | 178 | 4 | Go binding to ImageMagick's MagickWand C API | 2013-04-30 17:31:48 | 2023-11-02 01:56:09 |
| [canvas](https://github.com/tdewolff/canvas) | 1,352 | 85 | 10 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2017-05-20 18:10:51 | 2023-11-05 13:30:35 |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1,310 | 230 | 45 | Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv | 2013-12-09 09:43:26 | 2023-10-18 12:47:55 |
| [geopattern](https://github.com/pravj/geopattern) | 1,251 | 72 | 3 | :triangular_ruler: Create beautiful generative image patterns from a string in golang. | 2014-10-22 17:26:30 | 2023-11-07 12:38:32 |
| [stegify](https://github.com/DimitarPetrov/stegify) | 1,135 | 118 | 0 | 🔍 Go tool for LSB steganography, capable of hiding any file within an image. | 2018-11-29 16:45:58 | 2023-11-04 16:22:49 |
| [govips](https://github.com/davidbyttow/govips) | 1,009 | 226 | 41 | A lightning fast image processing and resizing library for Go | 2016-12-25 04:32:56 | 2023-11-07 20:46:12 |
| [image2ascii](https://github.com/qeesung/image2ascii) | 801 | 70 | 9 | :foggy: Convert image to ASCII | 2018-10-20 05:06:25 | 2023-11-07 22:24:13 |
| [goimagehash](https://github.com/corona10/goimagehash) | 649 | 65 | 10 | Go Perceptual image hashing package | 2017-07-28 17:15:58 | 2023-11-08 00:34:30 |
| [draft](https://github.com/lucasepe/draft) | 575 | 29 | 0 | Generate High Level Cloud Architecture diagrams using YAML syntax. | 2020-06-05 16:11:40 | 2023-10-20 08:56:13 |
| [govatar](https://github.com/o1egl/govatar) | 572 | 44 | 1 | Avatar generation library for GO language | 2016-01-18 12:12:28 | 2023-10-30 10:54:51 |
| [mort](https://github.com/aldor007/mort) | 486 | 23 | 6 | Storage and image processing server written in Go | 2017-11-19 13:37:58 | 2023-10-25 13:59:17 |
| [go-nude](https://github.com/koyachi/go-nude) | 394 | 48 | 3 | Nudity detection with Go. | 2014-05-02 08:32:29 | 2023-10-26 11:14:18 |
| [steganography](https://github.com/auyer/steganography) | 270 | 41 | 0 | Pure Golang Library that allows LSB steganography on images using ZERO dependencies | 2018-05-21 17:27:36 | 2023-11-07 15:57:44 |
| [darkroom](https://www.gojek.io/darkroom/) | 221 | 41 | 10 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2019-07-01 10:17:08 | 2023-11-07 12:51:53 |
| [gltf](https://www.khronos.org/gltf/) | 215 | 31 | 7 | :eyeglasses: Go library for encoding glTF 2.0 files | 2019-01-15 17:43:54 | 2023-11-01 19:30:26 |
| [rez](https://github.com/bamiaux/rez) | 209 | 19 | 1 | Image resizing in pure Go and SIMD | 2014-01-16 21:16:15 | 2023-09-08 16:44:49 |
| [mergi](http://mergi.io/) | 209 | 28 | 2 | go library for image programming (merge, crop, resize, watermark, animate, ease, transit) | 2018-09-24 03:40:47 | 2023-10-07 06:30:38 |
| [img](hawx.me/code/img) | 149 | 12 | 3 | A selection of image manipulation tools | 2012-07-28 19:57:47 | 2023-09-08 16:34:31 |
| [go-cairo](https://github.com/ungerik/go-cairo) | 135 | 32 | 1 | Go binding for the cairo graphics library | 2012-08-22 18:27:01 | 2023-06-12 04:35:46 |
| [cameron](https://pkg.go.dev/github.com/aofei/cameron) | 107 | 11 | 1 | An avatar generator for Go. | 2018-05-05 22:13:11 | 2023-10-17 20:38:01 |
| [gridder](https://github.com/shomali11/gridder) | 66 | 13 | 0 | A Grid based 2D Graphics library | 2020-04-10 00:13:10 | 2023-10-15 18:10:47 |
| [webp-server](https://github.com/mehdipourfar/webp-server) | 65 | 14 | 0 | Simple and minimal image server capable of storing, resizing, converting and caching images. | 2020-11-22 12:03:12 | 2023-09-25 15:35:52 |
| [go-gd](https://github.com/bolknote/go-gd) | 59 | 16 | 1 | Go bingings for GD (http://www.boutell.com/gd/) | 2011-05-12 06:33:54 | 2023-07-24 15:27:31 |
| [goimghdr](https://github.com/corona10/goimghdr) | 40 | 4 | 0 | The imghdr module determines the type of image contained in a file for go | 2018-02-25 09:34:44 | 2023-09-08 17:37:10 |
| [tga](https://github.com/ftrvxmtrx/tga) | 34 | 12 | 1 | Go package for decoding and encoding TARGA image format | 2012-10-08 01:09:20 | 2023-11-04 10:51:31 |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 27 | 6 | 0 | Port of webcolors library from Python to Go | 2014-04-24 14:41:22 | 2022-09-26 10:43:31 |
| [mpo](https://donatstudios.com/MPO-to-JPEG-Stereo) | 11 | 4 | 1 | JPEG-MPO Decoder / Converter Library and CLI Tool | 2015-04-14 22:37:59 | 2023-03-19 20:42:57 |
</details>

### IoT (Internet of Things)
Libraries for programming devices of the IoT.

<sup>*Last Update: 2023-12-07 21:47:29*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gobot](https://gobot.io) | 8,546 | 1,073 | 123 | Golang framework for robotics, drones, and the Internet of Things (IoT) | 2013-09-21 14:09:19 | 2023-12-05 09:34:23 |
| [flogo](http://flogo.io) | 2,338 | 290 | 127 | Project Flogo is an open source ecosystem of opinionated  event-driven capabilities to simplify building efficient & modern serverless functions, microservices & edge apps. | 2016-07-10 02:57:43 | 2023-12-01 10:20:08 |
| [mainflux](https://www.mainflux.io) | 2,250 | 626 | 107 | Industrial IoT Messaging and Device Management Platform | 2015-07-06 20:31:50 | 2023-12-05 04:30:32 |
| [periph](https://periph.io) | 1,763 | 210 | 0 | Older version of periph, see new version at https://github.com/periph | 2016-10-13 16:53:51 | 2023-12-05 06:05:40 |
| [gatt](http://flogo.io) | 1,091 | 294 | 51 | Gatt is a Go package for building Bluetooth Low Energy peripherals | 2014-04-23 13:45:27 | 2023-12-02 05:41:23 |
| [heedy](https://heedy.org) | 379 | 32 | 21 | An aggregator for personal metrics, and an extensible analysis engine | 2015-01-16 19:44:21 | 2023-11-25 12:34:37 |
| [devices](https://github.com/goiot/devices) | 261 | 33 | 9 | Suite of libraries for IoT devices (written in Go), experimental for x/exp/io | 2016-05-30 08:07:02 | 2023-11-24 19:22:25 |
| [huego](https://github.com/amimof/huego) | 241 | 37 | 10 | An extensive Philips Hue client library for Go with an emphasis on simplicity | 2017-05-16 05:31:45 | 2023-12-01 08:43:27 |
| [sensorbee](http://sensorbee.io/) | 226 | 42 | 39 | Lightweight stream processing engine for IoT | 2016-02-19 07:49:56 | 2023-11-11 08:39:34 |
| [iot](https://github.com/vaelen/iot) | 64 | 13 | 0 | A Go client for Google IoT Core | 2018-03-08 06:51:51 | 2023-11-04 01:47:14 |
| [eywa](https://github.com/xcodersun/eywa) | 61 | 15 | 9 | Make IoT a lot more fun with data.  | 2016-02-20 17:02:16 | 2023-10-05 14:10:04 |
</details>

### JSON
Libraries for working with JSON.

<sup>*Last Update: 2023-12-06 21:07:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gjson](https://github.com/tidwall/gjson) | 13,098 | 866 | 71 | Get JSON values quickly - JSON parser for Go | 2016-08-11 03:08:47 | 2023-12-06 11:29:02 |
| [json-to-go](https://mholt.github.io/json-to-go/) | 4,274 | 513 | 16 | Translates JSON into a Go type in your browser instantly (original) | 2014-01-21 18:11:13 | 2023-12-05 01:29:42 |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2,631 | 242 | 41 | Automatically generate Go (golang) struct definitions from example JSON | 2012-12-27 19:10:50 | 2023-12-02 21:10:11 |
| [fastjson](https://github.com/valyala/fastjson) | 2,041 | 122 | 54 | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection | 2018-05-28 21:41:47 | 2023-12-05 08:10:00 |
| [jsondiff](https://pkg.go.dev/github.com/wI2L/jsondiff) | 358 | 33 | 0 | Compute the diff between two JSON documents as a series of RFC6902 (JSON Patch) operations | 2020-11-28 19:05:16 | 2023-12-04 13:18:49 |
| [kazaam](https://github.com/qntfy/kazaam) | 269 | 54 | 23 | Arbitrary transformations of JSON in Golang | 2016-07-19 14:19:03 | 2023-11-21 10:52:08 |
| [json-to-proto.github.io](https://json-to-proto.github.io) | 218 | 42 | 2 | convert JSON to Protocol Buffers online in your browser instantly | 2020-04-18 20:42:45 | 2023-12-03 01:46:33 |
| [ajson](https://github.com/spyzhov/ajson) | 195 | 21 | 12 | Abstract JSON for Golang with JSONPath support  | 2019-03-07 20:47:38 | 2023-12-05 08:03:30 |
| [gojq](https://godoc.org/github.com/nicklaw5/go-respond) | 190 | 23 | 1 | JSON query in Golang | 2015-12-30 09:02:13 | 2023-09-26 22:59:23 |
| [jettison](https://pkg.go.dev/github.com/wI2L/jettison) | 165 | 11 | 2 | Highly configurable, fast JSON encoder for Go | 2019-08-30 13:28:03 | 2023-12-05 19:38:51 |
| [json2go](https://m-zajac.github.io/json2go) | 124 | 17 | 1 | Create go type representation from json | 2017-06-10 23:55:07 | 2023-10-20 04:01:54 |
| [gjo](https://github.com/skanehira/gjo) | 122 | 15 | 1 | Small utility to create JSON objects | 2019-02-23 01:54:21 | 2023-09-26 22:58:24 |
| [jsongo](https://pkg.go.dev/github.com/wI2L/jsondiff) | 109 | 15 | 1 | Fluent API to make it easier to create Json objects. | 2015-08-07 23:23:17 | 2023-02-02 12:48:01 |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 102 | 8 | 1 | A JSON diff utility | 2017-04-24 16:05:35 | 2023-09-26 22:59:43 |
| [jsonf](https://pkg.go.dev/github.com/wI2L/jsondiff) | 65 | 11 | 0 | Console JSON formatter with query feature | 2015-05-25 04:53:32 | 2023-09-26 23:00:18 |
| [go-respond](https://godoc.org/github.com/nicklaw5/go-respond) | 52 | 10 | 1 | A Go package for handling common HTTP JSON responses. | 2017-03-12 21:00:54 | 2023-09-26 22:59:02 |
| [mp](https://github.com/sanbornm/mp) | 47 | 8 | 1 | Simple Email Parser | 2014-06-15 21:14:39 | 2023-10-15 19:50:38 |
| [ask](https://github.com/simonnilsson/ask) | 31 | 3 | 0 | A Go package that provides a simple way of accessing nested properties in maps and slices. | 2020-09-13 13:53:31 | 2023-11-11 07:32:35 |
| [mapslice-json](https://github.com/ake-persson/mapslice-json) | 16 | 7 | 4 | Go MapSlice for ordered marshal/ unmarshal of maps in JSON | 2020-02-19 11:01:48 | 2023-09-26 23:01:11 |
| [dynjson](https://github.com/cocoonspace/dynjson) | 16 | 6 | 1 | Client-customizable JSON formats for dynamic APIs | 2020-05-06 07:10:02 | 2023-09-26 22:57:53 |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 14 | 2 | 1 | Small package which wraps error responses to follow jsonapi.org | 2018-10-18 15:03:45 | 2023-09-26 22:58:50 |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 14 | 7 | 1 | A simple Go package to make custom structs marshal into HAL compatible JSON responses. | 2016-01-15 11:38:40 | 2023-09-26 23:00:26 |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 14 | 3 | 0 | Go bindings based on the JSON API errors reference | 2016-07-08 10:08:58 | 2023-09-26 23:00:03 |
| [epoch](https://github.com/vtopc/epoch) | 13 | 3 | 1 | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON | 2019-12-15 12:54:37 | 2023-09-26 22:58:05 |
| [jzon](https://github.com/zerosnake0/jzon) | 12 | 2 | 0 | A golang json library inspired by jsoniter | 2019-11-12 10:42:41 | 2023-11-27 01:04:09 |
| [ej](https://github.com/lucassscaravelli/ej) | 10 | 2 | 0 | Write and read JSON from different sources in one line | 2020-01-04 17:39:35 | 2023-09-26 22:57:59 |
| [jsonic](https://github.com/sinhashubham95/jsonic) | 10 | 2 | 0 | All you need with JSON | 2021-01-09 06:21:59 | 2023-09-26 23:00:47 |
</details>

### Job Scheduler
Libraries for scheduling jobs.

<sup>*Last Update: 2023-11-20 08:20:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocron](https://github.com/go-co-op/gocron) | 4,262 | 295 | 14 | Easy and fluent Go cron scheduling. This is a fork from https://github.com/jasonlvhit/gocron | 2020-03-20 15:33:05 | 2023-11-19 18:51:20 |
| [go-quartz](https://pkg.go.dev/github.com/reugn/go-quartz/quartz) | 1,552 | 83 | 8 | Minimalist and zero-dependency scheduling library for Go | 2019-04-14 18:57:51 | 2023-11-19 21:03:02 |
| [gron](https://github.com/roylee0704/gron) | 996 | 65 | 9 | gron, Cron Jobs in Go. | 2016-06-04 08:02:22 | 2023-11-13 08:14:21 |
| [jobrunner](https://github.com/bamzi/jobrunner) | 987 | 98 | 11 | Framework for performing work asynchronously, outside of the request flow | 2015-10-21 04:17:01 | 2023-11-17 09:40:15 |
| [jobs](https://github.com/albrow/jobs) | 496 | 47 | 17 | A persistent and flexible background jobs library for go. | 2015-02-09 22:13:29 | 2023-11-12 19:08:09 |
| [scheduler](https://github.com/carlescere/scheduler) | 435 | 54 | 8 | Job scheduling made easy. | 2015-02-03 17:10:23 | 2023-11-18 20:02:38 |
| [go-cron](https://github.com/rk/go-cron) | 223 | 15 | 0 | A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. | 2011-04-15 14:50:49 | 2023-11-15 13:50:44 |
| [tasks](https://github.com/madflojo/tasks) | 205 | 22 | 2 | Package tasks is an easy to use in-process scheduler for recurring tasks in Go | 2019-12-24 18:26:18 | 2023-11-15 10:22:22 |
| [clockwerk](https://github.com/onatm/clockwerk) | 140 | 14 | 0 | Job Scheduling Library | 2017-04-09 23:10:48 | 2023-10-17 16:14:24 |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 100 | 13 | 12 | You had one job, or more then one, which can be done in steps | 2018-04-08 13:44:04 | 2023-10-20 19:32:59 |
| [clockwork](https://github.com/whiteShtef/clockwork) | 27 | 13 | 2 | Job Scheduling Library | 2020-02-21 01:25:57 | 2023-03-04 16:27:09 |
| [cronticker](https://github.com/krayzpipes/cronticker) | 11 | 3 | 0 | Golang ticker that works with Cron scheduling. | 2020-11-28 20:59:38 | 2023-08-11 06:19:06 |
</details>

### Logging
Libraries for generating and working with log files.

<sup>*Last Update: 2023-12-04 09:33:29*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [logrus](https://github.com/sirupsen/logrus) | 23,485 | 2,321 | 70 | Structured, pluggable logging for Go. | 2013-10-16 19:08:55 | 2023-12-04 00:24:30 |
| [zap](https://pkg.go.dev/go.uber.org/zap) | 20,097 | 1,440 | 136 | Blazing fast, structured, leveled logging in Go. | 2016-02-18 19:52:56 | 2023-12-03 23:54:24 |
| [zerolog](https://github.com/rs/zerolog) | 9,115 | 521 | 140 | Zero Allocation JSON Logger | 2017-05-12 05:24:39 | 2023-12-03 17:50:24 |
| [go-spew](https://github.com/davecgh/go-spew) | 5,771 | 406 | 61 | Implements a deep pretty printer for Go data structures to aid in debugging | 2013-01-09 05:18:22 | 2023-12-03 12:29:44 |
| [lumberjack](https://github.com/natefinch/lumberjack) | 4,371 | 582 | 78 | lumberjack is a log rolling package for Go | 2014-06-14 11:55:47 | 2023-12-04 02:22:17 |
| [glog](https://github.com/golang/glog) | 3,471 | 956 | 2 | Leveled execution logs for Go | 2013-07-16 04:33:04 | 2023-12-02 11:00:00 |
| [tail](https://github.com/hpcloud/tail) | 2,632 | 502 | 74 | Go package for reading from continously updated files (tail -f) | 2013-02-05 00:28:03 | 2023-12-03 05:07:42 |
| [seelog](https://github.com/cihub/seelog) | 1,639 | 249 | 39 | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. | 2011-11-17 09:43:15 | 2023-12-01 08:41:39 |
| [log](https://github.com/apex/log) | 1,339 | 151 | 48 | Structured logging package for Go. | 2015-12-21 20:27:48 | 2023-11-30 11:33:21 |
| [log15](https://godoc.org/github.com/inconshreveable/log15) | 1,097 | 188 | 46 | Structured, composable logging for Go | 2014-05-20 00:11:52 | 2023-12-02 10:57:53 |
| [log](https://github.com/phuslu/log) | 522 | 36 | 1 | High performance structured logging | 2019-07-07 09:40:38 | 2023-12-03 15:19:34 |
| [onelog](https://github.com/francoispqt/onelog) | 414 | 16 | 2 | Dead simple, super fast, zero allocation logger for Golang | 2018-05-06 14:32:10 | 2023-11-26 08:32:44 |
| [httpretty](https://asciinema.org/a/297429) | 377 | 14 | 2 | Package httpretty prints the HTTP requests you make with Go pretty on your terminal. | 2020-01-24 18:17:16 | 2023-12-03 20:36:40 |
| [logutils](https://logur.dev/logur) | 358 | 35 | 3 | Utilities for slightly better logging in Go (Golang). | 2013-10-09 07:31:15 | 2023-11-19 21:06:45 |
| [logxi](https://logur.dev/logur) | 352 | 42 | 22 | A 12-factor app logger built for performance and happy development | 2015-03-01 22:13:45 | 2023-11-29 13:14:59 |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 327 | 14 | 8 | A logger for Go SQL database driver without modifying existing *sql.DB stdlib usage. | 2019-11-02 17:28:03 | 2023-11-28 17:38:44 |
| [log](https://github.com/go-playground/log) | 289 | 23 | 1 | :green_book: Simple, configurable and scalable Structured Logging for Go. | 2016-02-07 16:17:48 | 2023-11-24 19:20:25 |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 288 | 43 | 8 | Rolling writer is an IO util for auto rolling write in go. | 2017-02-12 12:05:26 | 2023-12-02 09:24:06 |
| [go-logger](https://github.com/apsdehal/go-logger) | 286 | 52 | 3 | Simple logger for Go programs. Allows custom formats for messages. | 2014-09-26 04:57:06 | 2023-08-04 17:30:34 |
| [logur](https://logur.dev/logur) | 197 | 13 | 8 | Logur is an opinionated collection of logging best practices | 2018-12-09 16:43:11 | 2023-11-29 13:15:31 |
| [glg](https://github.com/kpango/glg) | 187 | 15 | 3 | Simple and blazing fast lockfree logging library for golang | 2017-06-21 13:26:16 | 2023-10-24 08:30:41 |
| [logger](http://godoc.org/github.com/azer/logger) | 158 | 16 | 0 | Minimalistic logging library for Go. | 2014-09-30 06:45:09 | 2023-08-11 02:48:21 |
| [xlog](https://github.com/rs/xlog) | 138 | 13 | 3 | xlog is a logger for net/context aware HTTP applications | 2015-10-22 09:26:45 | 2023-08-11 15:00:52 |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 122 | 34 | 9 | A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. | 2015-10-22 22:29:02 | 2023-07-07 08:55:37 |
| [logvoyage](https://github.com/firstrow/logvoyage) | 94 | 12 | 9 | LogVoyage - logging SaaS written in GoLang | 2015-03-29 11:05:09 | 2023-03-30 03:26:37 |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 56 | 8 | 3 | Time based rotating file writer | 2017-02-04 09:02:55 | 2023-07-26 07:22:39 |
| [log](https://github.com/alexcesaro/log) | 47 | 5 | 1 | Logging packages for Go | 2014-04-19 14:31:56 | 2023-03-01 08:43:21 |
| [gone](https://github.com/One-com/gone) | 47 | 8 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2023-10-26 09:27:48 |
| [logex](https://github.com/chzyer/logex) | 42 | 12 | 2 | An golang log lib, supports tracking and level, wrap by standard log lib | 2014-10-10 06:38:39 | 2023-09-23 13:45:26 |
| [go-log](https://github.com/ian-kent/go-log) | 42 | 21 | 3 | A logger, for Go | 2014-05-02 00:34:09 | 2023-06-27 22:50:47 |
| [gologger](https://github.com/sadlil/gologger) | 40 | 10 | 2 | The Simplest and worst logging library ever written | 2015-09-02 08:52:26 | 2023-10-19 09:19:21 |
| [journald](https://asciinema.org/a/297429) | 37 | 2 | 0 | Go implementation of systemd Journal's native API for logging | 2017-08-23 07:06:09 | 2023-09-19 06:03:02 |
| [go-log](https://github.com/siddontang/go-log) | 33 | 18 | 1 | a golang log lib supports level and multi handlers | 2014-05-18 03:41:55 | 2022-12-29 07:34:47 |
| [mlog](https://github.com/jbrodriguez/mlog) | 33 | 19 | 1 | A simple logging module for go, with a rotating file feature and console logging. | 2014-10-20 15:06:26 | 2023-07-11 00:08:26 |
| [distillog](https://github.com/amoghe/distillog) | 32 | 7 | 1 | Logging, distilled | 2015-10-12 16:32:21 | 2023-07-13 20:49:47 |
| [logrusly](https://github.com/sebest/logrusly) | 28 | 18 | 3 | Loggly Hooks for GO Logrus logger | 2014-09-11 23:27:11 | 2022-09-26 23:23:31 |
| [log](https://github.com/teris-io/log) | 26 | 3 | 0 | Structured log interface | 2017-10-28 19:57:55 | 2022-09-26 23:23:04 |
| [zkits-logger](https://github.com/edoger/zkits-logger) | 25 | 2 | 1 | A powerful zero-dependency json logger. | 2020-03-31 14:23:40 | 2023-10-18 14:20:35 |
| [gomol](https://github.com/aphistic/gomol) | 19 | 1 | 3 | Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs | 2015-08-30 15:51:46 | 2023-02-21 18:23:08 |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 16 | 1 | 0 | io.Writer implementation using logrus logger | 2019-08-09 08:58:25 | 2023-09-03 02:47:33 |
| [logmatic](http://godoc.org/github.com/azer/logger) | 16 | 5 | 1 | Colorized logger for Golang with dynamic log level configuration | 2018-11-07 01:52:45 | 2023-01-12 17:01:00 |
| [glo](https://github.com/lajosbencz/glo) | 15 | 1 | 0 | Logging library for Golang | 2019-01-19 22:10:42 | 2022-09-26 23:22:25 |
| [go-log](https://github.com/subchen/go-log) | 14 | 7 | 0 | Simple and configurable Logging in Go, with level, formatters and writers | 2017-05-07 08:09:24 | 2023-04-28 08:50:54 |
| [logdump](https://github.com/ewwwwwqm/logdump) | 11 | 3 | 0 | Package for multi-level logging | 2017-01-13 15:34:31 | 2023-03-05 02:03:20 |
| [logo](https://github.com/mbndr/logo) | 11 | 2 | 0 | Golang logger to different configurable writers. | 2017-02-07 18:02:55 | 2022-09-26 23:23:24 |
| [kemba](https://pkg.go.dev/github.com/clok/kemba?tab=overview) | 11 | 2 | 7 | A tiny debug logging tool. Ideal for CLI tools and command applications. Inspired by https://github.com/visionmedia/debug | 2020-07-13 03:10:54 | 2023-11-12 23:43:41 |
| [log](https://github.com/aerogo/log) | 10 | 1 | 0 | :memo: Logging with multiple output targets. | 2017-06-10 09:54:08 | 2022-09-26 23:22:57 |
| [go-log](https://github.com/pieterclaerhout/go-log) | 10 | 5 | 1 | A logging library with strack traces, object dumping and optional timestamps | 2019-10-01 08:55:38 | 2023-03-07 11:39:07 |
| [xlog](https://github.com/xfxdev/xlog) | 8 | 4 | 0 | plugin architecture and flexible log system for golang | 2016-05-05 16:47:45 | 2022-09-26 23:24:18 |
| [yell](https://github.com/jfcg/yell) | 1 | 0 | 0 | :ledger: Yet another minimalist logging library | 2021-02-07 16:07:27 | 2022-09-26 23:24:24 |
</details>

### Machine Learning
Libraries for Machine Learning.

<sup>*Last Update: 2023-12-07 21:47:31*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 9,063 | 1,263 | 83 | Machine Learning for Go | 2013-12-26 13:06:14 | 2023-12-07 12:24:58 |
| [gorse](https://gorse.io) | 7,751 | 726 | 68 | Gorse open source recommender system engine | 2018-08-14 11:01:09 | 2023-12-07 13:28:22 |
| [gorgonia](https://gorgonia.org/) | 5,178 | 455 | 105 | Gorgonia is a library that helps facilitate machine learning in Go. | 2016-09-14 23:19:43 | 2023-12-06 02:46:16 |
| [gosseract](https://pkg.go.dev/github.com/otiai10/gosseract) | 2,323 | 303 | 22 | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library | 2013-10-11 07:27:53 | 2023-12-07 08:19:38 |
| [tfgo](https://pgaleone.eu/tensorflow/go/2017/05/29/understanding-tensorflow-using-go/) | 2,319 | 200 | 17 | Tensorflow + Go, the gopher way | 2017-05-23 13:27:39 | 2023-12-06 10:04:45 |
| [goml](https://github.com/cdipaolo/goml) | 1,512 | 177 | 3 | On-line Machine Learning in Go (and so much more) | 2015-06-27 05:52:01 | 2023-11-29 14:04:14 |
| [eaopt](https://pkg.go.dev/github.com/MaxHalford/eaopt) | 860 | 99 | 7 | :four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution) | 2016-01-31 00:04:52 | 2023-12-05 05:21:14 |
| [bayesian](https://github.com/jbrukh/bayesian) | 775 | 128 | 8 | Naive Bayesian Classification for Golang. | 2011-11-23 04:17:00 | 2023-11-29 11:44:54 |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 731 | 93 | 34 | Ensembles of decision trees in go/golang. | 2012-10-22 17:38:16 | 2023-12-06 08:04:16 |
| [ocrserver](https://ocr-example.herokuapp.com/) | 623 | 136 | 1 | A simple OCR API server, seriously easy to be deployed by Docker, on Heroku as well | 2015-11-15 07:57:42 | 2023-12-07 10:59:36 |
| [onnx-go](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 565 | 62 | 39 | onnx-go gives the ability to import a pre-trained neural network within Go without being linked to a framework or library. | 2018-08-28 07:39:20 | 2023-12-07 02:52:58 |
| [gobrain](https://github.com/goml/gobrain) | 546 | 59 | 2 | Neural Networks written in go | 2014-04-29 13:32:36 | 2023-12-05 16:26:45 |
| [go-deep](https://github.com/patrikeh/go-deep) | 487 | 61 | 1 | Artificial Neural Network | 2017-12-09 15:10:06 | 2023-12-04 20:11:14 |
| [regommend](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 305 | 28 | 0 | Recommendation engine for Go | 2014-02-05 17:00:49 | 2023-12-07 11:49:25 |
| [goptuna](https://pkg.go.dev/github.com/c-bata/goptuna) | 243 | 17 | 16 | A hyperparameter optimization framework, inspired by Optuna. | 2019-07-24 12:03:05 | 2023-12-01 13:26:51 |
| [goga](https://github.com/tomcraven/goga) | 205 | 19 | 0 | Golang Genetic Algorithm | 2015-10-20 12:50:51 | 2023-11-30 16:25:01 |
| [go-galib](https://github.com/thoj/go-galib) | 195 | 41 | 0 | Genetic Algorithms library written in Go / golang | 2009-11-30 10:46:58 | 2023-10-04 06:07:03 |
| [goRecommend](https://pkg.go.dev/github.com/c-bata/goptuna) | 195 | 22 | 0 | Collaborative Filtering (CF) Algorithms in Go!  | 2014-07-16 05:32:23 | 2023-12-01 13:16:58 |
| [shield](https://github.com/eaigner/shield) | 154 | 33 | 5 | Bayesian text classifier with flexible tokenizers and storage backends for Go | 2013-04-10 19:38:16 | 2023-09-28 10:21:49 |
| [go-fann](https://github.com/vksnk/go-fann) | 115 | 21 | 2 | Go bindings for FANN, library for artificial neural networks | 2011-03-10 21:08:27 | 2023-10-10 09:40:17 |
| [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 102 | 8 | 4 | 🔥 Fast, simple sklearn-like feature processing for Go | 2020-12-18 13:09:18 | 2023-11-26 10:56:26 |
| [goscore](https://gorse.io) | 93 | 22 | 3 | Go Scoring API for PMML | 2017-08-19 11:08:39 | 2023-12-01 13:24:58 |
| [gonet](https://pkg.go.dev/github.com/dathoangnd/gonet) | 80 | 8 | 0 | Neural Network for Go. | 2020-01-11 18:27:28 | 2023-12-01 13:16:30 |
| [fonet](https://github.com/Fontinalis/fonet) | 79 | 17 | 2 | fonet is a deep neural network package for Go. | 2017-10-03 15:57:15 | 2023-10-06 03:49:01 |
| [libsvm](https://pkg.go.dev/github.com/otiai10/gosseract) | 73 | 13 | 1 | libsvm go version | 2012-07-31 07:57:47 | 2023-01-16 01:34:04 |
| [neat](https://github.com/jinyeom/neat) | 69 | 13 | 4 | NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go | 2016-11-17 04:23:14 | 2023-12-01 13:26:15 |
| [gomind](https://github.com/surenderthakran/gomind) | 69 | 8 | 7 | A simplistic Neural Network Library in Go | 2017-10-19 03:48:51 | 2023-12-02 17:55:11 |
| [neural-go](https://github.com/schuyler/neural-go) | 67 | 16 | 1 | A multilayer perceptron network implemented in Go, with training via backpropagation. | 2011-10-17 09:31:33 | 2023-11-15 21:15:18 |
| [go-pr](https://github.com/daviddengcn/go-pr) | 62 | 14 | 0 | Pattern recognition package in Go lang. | 2013-06-07 02:36:20 | 2022-09-27 08:38:25 |
| [Varis](https://github.com/Xamber/Varis) | 53 | 9 | 0 | Golang Neural Network  | 2017-10-10 08:43:27 | 2023-12-01 13:17:10 |
| [golinear](https://github.com/danieldk/golinear) | 44 | 12 | 0 | liblinear bindings for Go | 2013-04-05 15:37:01 | 2023-09-14 18:45:19 |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 39 | 9 | 0 | k-modes and k-prototypes clustering algorithms implementation in Go | 2017-10-04 12:24:52 | 2023-04-27 09:04:36 |
| [godist](https://github.com/e-dard/godist) | 36 | 7 | 0 | Probability distributions and associated methods in Go | 2014-09-05 09:48:51 | 2023-11-15 21:13:05 |
| [randomForest](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 36 | 7 | 0 | Random Forest implementation in golang | 2018-10-25 07:05:29 | 2023-11-28 00:11:21 |
| [ddt](https://github.com/sgrodriguez/ddt) | 32 | 4 | 1 | Golang Dynamic Decision Tree | 2020-05-20 13:51:42 | 2023-11-15 02:21:13 |
| [evoli](https://github.com/khezen/evoli) | 27 | 9 | 21 | Genetic Algorithm and Particle Swarm Optimization | 2015-06-12 06:58:30 | 2023-08-25 09:08:56 |
| [probab](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 19 | 7 | 3 | Automatically exported from code.google.com/p/probab | 2015-09-14 12:07:52 | 2023-11-13 20:12:08 |
</details>

### Messaging
Libraries that implement messaging systems.

<sup>*Last Update: 2023-11-11 19:52:23*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [sarama](https://github.com/IBM/sarama) | 10,510 | 1,752 | 180 | Sarama is a Go library for Apache Kafka. | 2013-07-05 18:52:38 | 2023-11-11 07:56:00 |
| [asynq](https://github.com/hibiken/asynq) | 7,508 | 580 | 178 | Simple, reliable, and efficient distributed task queue in Go | 2019-11-15 05:17:55 | 2023-11-11 12:04:33 |
| [centrifugo](https://centrifugal.dev) | 7,488 | 553 | 10 | Scalable real-time messaging server in a language-agnostic way. Self-hosted alternative to Pubnub, Pusher, Ably. Set up once and forever. | 2015-03-31 20:26:49 | 2023-11-10 17:51:09 |
| [gorush](https://github.com/appleboy/gorush) | 7,354 | 864 | 49 | A push notification server written in Go (Golang). | 2016-03-22 07:15:20 | 2023-11-11 12:10:37 |
| [machinery](https://github.com/RichardKnop/machinery) | 7,078 | 885 | 225 | Machinery is an asynchronous task queue/job queue based on distributed message passing. | 2015-04-05 19:46:34 | 2023-11-11 12:25:09 |
| [benthos](https://www.benthos.dev) | 6,782 | 662 | 401 | Fancy stream processing made operationally mundane | 2016-03-22 01:18:48 | 2023-11-11 12:29:17 |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 5,385 | 859 | 123 | socket.io library for golang, a realtime application framework. | 2013-07-13 13:04:38 | 2023-11-10 14:27:46 |
| [nats.go](https://nats.io) | 4,861 | 672 | 122 | Golang client for NATS, the cloud native messaging system. | 2012-08-15 12:54:59 | 2023-11-11 11:42:54 |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 4,204 | 653 | 232 | Confluent's Apache Kafka Golang client | 2016-07-12 22:23:34 | 2023-11-11 08:38:58 |
| [mercure](https://mercure.rocks) | 3,525 | 267 | 22 | An open, easy, fast, reliable and battery-efficient solution for real-time communications | 2018-07-14 13:47:14 | 2023-11-11 10:25:32 |
| [melody](https://github.com/olahol/melody) | 3,327 | 383 | 11 | :notes: Minimalist websocket framework for Go | 2015-05-13 20:38:32 | 2023-11-10 08:45:40 |
| [apns2](https://github.com/sideshow/apns2) | 2,879 | 328 | 36 | ⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol. | 2016-01-05 00:56:53 | 2023-11-11 03:23:42 |
| [go-nsq](https://github.com/nsqio/go-nsq) | 2,476 | 468 | 26 | The official Go package for NSQ | 2013-08-29 01:18:32 | 2023-11-11 05:11:21 |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 2,077 | 609 | 5 | Golang push server cluster | 2013-12-27 08:56:10 | 2023-10-26 02:28:17 |
| [EventBus](https://github.com/asaskevich/EventBus) | 1,533 | 238 | 27 | [Go] Lightweight eventbus with async compatibility for Go | 2014-12-19 16:38:39 | 2023-11-08 07:27:43 |
| [uniqush-push](http://uniqush.org) | 1,516 | 205 | 73 | Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices. | 2011-08-29 08:42:37 | 2023-11-08 19:53:19 |
| [Beaver](https://clivern.github.io/Beaver/) | 1,475 | 82 | 10 | 💨 A real time messaging system to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. | 2018-10-20 21:10:43 | 2023-11-06 00:23:23 |
| [zmq4](http://uniqush.org) | 1,089 | 166 | 51 | A Go interface to ZeroMQ version 4 | 2013-10-18 11:48:51 | 2023-11-08 09:17:07 |
| [gollum](http://gollum.readthedocs.org/en/latest/) | 934 | 80 | 22 | An n:m message multiplexer written in Go | 2015-06-20 21:51:20 | 2023-10-23 11:10:13 |
| [dbus](https://github.com/godbus/dbus) | 896 | 218 | 43 | Native Go bindings for D-Bus | 2014-03-27 19:07:41 | 2023-10-30 19:35:46 |
| [golongpoll](https://github.com/jcuga/golongpoll) | 640 | 61 | 0 | golang long polling library.  Makes web pub-sub easy via HTTP long-poll servers and clients :smiley: :coffee: :computer: | 2015-11-02 00:32:56 | 2023-11-09 07:43:36 |
| [mangos](https://github.com/nanomsg/mangos) | 610 | 75 | 26 | mangos is a pure Golang implementation of nanomsg's "Scalablilty Protocols" | 2018-10-12 17:35:46 | 2023-11-01 22:51:17 |
| [emitter](https://github.com/olebedev/emitter) | 483 | 35 | 4 | Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins | 2015-11-10 20:56:36 | 2023-11-09 07:42:15 |
| [glue](https://github.com/desertbit/glue) | 408 | 32 | 5 | Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) | 2015-06-07 10:21:15 | 2023-09-11 06:48:22 |
| [pubsub](https://github.com/cskr/pubsub) | 391 | 65 | 2 | A simple pubsub package for go. | 2012-04-01 06:31:43 | 2023-10-28 14:04:10 |
| [bus](https://pkg.go.dev/github.com/mustafaturan/bus/v3?tab=doc) | 316 | 21 | 0 | 🔊Minimalist message bus implementation for internal communication with zero-allocation magic on Emit | 2019-04-27 06:41:53 | 2023-10-26 11:19:18 |
| [message-bus](http://rafallorenz.com/message-bus) | 261 | 43 | 2 | Go simple async message bus | 2017-10-04 09:18:34 | 2023-10-26 02:05:31 |
| [rabtap](https://github.com/jandelgado/rabtap) | 247 | 15 | 2 | RabbitMQ wire tap and swiss army knife | 2017-11-11 11:32:39 | 2023-10-27 10:43:04 |
| [guble](https://github.com/smancke/guble) | 155 | 23 | 5 | websocket based messaging server written in golang | 2015-11-15 20:32:42 | 2023-03-14 21:05:06 |
| [hub](https://github.com/leandro-lugaresi/hub) | 135 | 17 | 2 | :incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications | 2018-04-13 23:47:13 | 2023-11-03 06:41:56 |
| [redisqueue](https://godoc.org/github.com/robinjoseph08/redisqueue) | 114 | 50 | 7 | redisqueue provides a producer and consumer of a queue that uses Redis streams | 2019-07-07 04:36:54 | 2023-10-25 16:52:04 |
| [oplog](https://github.com/dailymotion/oplog) | 113 | 13 | 2 | A generic oplog/replication system for microservices | 2014-11-06 09:11:15 | 2023-08-20 15:14:48 |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 97 | 23 | 6 | A tiny wrapper over amqp exchanges and queues 🚌 ✨ | 2017-05-07 08:51:11 | 2023-06-07 09:41:54 |
| [go-mq](https://github.com/cheshir/go-mq) | 88 | 18 | 1 | Declare AMQP entities like queues, producers, and consumers in a declarative way. Can be used to work with RabbitMQ. | 2017-06-19 16:16:30 | 2023-09-20 10:31:26 |
| [drone-line](https://github.com/appleboy/drone-line) | 78 | 17 | 2 | Sending line notifications using a binary, docker or Drone CI. | 2016-09-13 05:21:44 | 2023-10-23 01:17:51 |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 77 | 17 | 2 | A tiny wrapper around NSQ topic and channel :rocket: | 2017-01-15 22:05:13 | 2023-06-18 10:05:11 |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 67 | 12 | 1 | RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue | 2016-10-04 21:07:48 | 2023-10-08 23:04:12 |
| [go-notify](https://github.com/TheCreeper/go-notify) | 67 | 11 | 1 | Package notify provides an implementation of the Gnome DBus Notifications Specification. | 2015-03-01 19:21:44 | 2023-06-28 17:14:23 |
| [commander](https://github.com/jeroenrinzema/commander) | 65 | 5 | 2 | Build event-driven and event streaming applications with ease | 2018-04-20 12:30:51 | 2023-05-25 08:42:55 |
| [go-res](https://github.com/jirenius/go-res) | 63 | 8 | 8 | RES Service protocol library for Go | 2018-07-15 09:10:11 | 2023-06-17 06:55:57 |
| [event](https://github.com/agoalofalife/event) | 54 | 12 | 0 | The implementation of the pattern observer | 2017-07-02 12:19:56 | 2023-07-26 13:36:32 |
| [hare](https://github.com/leozz37/hare) | 51 | 13 | 1 | 🐇  CLI tool for websockets and Go package | 2020-12-01 22:30:27 | 2023-08-21 05:55:10 |
| [ami](https://github.com/kak-tus/ami) | 27 | 8 | 0 | Go client to reliable queues based on Redis Cluster Streams | 2018-10-27 10:38:16 | 2023-10-31 09:43:34 |
| [go-vitotrol](https://godoc.org/github.com/maxatome/go-vitotrol) | 23 | 9 | 2 | golang client library to Viessmann Vitotrol web service | 2016-11-03 19:59:43 | 2023-10-23 12:05:44 |
| [gosd](https://github.com/alexsniffin/gosd) | 21 | 5 | 0 | A library for scheduling when to dispatch a message to a channel | 2020-05-17 23:19:51 | 2023-08-09 21:06:16 |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 21 | 2 | 0 | RabbitMQ Reconnection client | 2019-01-14 16:05:44 | 2023-11-04 03:48:32 |
| [jazz](https://github.com/socifi/jazz) | 18 | 4 | 1 | Abstraction layer for simple rabbitMQ connection, messaging and administration | 2018-10-22 12:28:15 | 2023-06-28 09:06:17 |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 11 | 4 | 0 | Gaurun Client written in Go | 2017-06-29 02:50:51 | 2022-09-27 08:41:03 |
</details>

### Microsoft Office


<sup>*Last Update: 2023-12-16 21:01:14*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [unioffice](https://unidoc.io/unioffice/) | 4,027 | 444 | 33 | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents | 2017-08-29 01:25:48 | 2023-12-16 09:03:13 |
</details>

### Microsoft Office - Microsoft Excel
Libraries for working with Microsoft Excel.

<sup>*Last Update: 2023-11-20 08:20:28*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [excelize](https://xuri.me/excelize) | 16,246 | 1,588 | 80 | Go language library for reading and writing Microsoft Excel™ (XLAM / XLSM / XLSX / XLTM / XLTX) spreadsheets | 2016-08-29 12:32:12 | 2023-11-19 20:00:22 |
| [xlsx](https://github.com/tealeg/xlsx) | 5,588 | 813 | 5 | Go library for reading and writing XLSX files.  | 2011-06-28 15:20:28 | 2023-11-17 17:16:35 |
| [go-excel](https://github.com/szyhf/go-excel) | 179 | 35 | 1 | A simple and light excel file reader to read a standard excel as a table faster | 一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。 | 2017-09-03 11:51:58 | 2023-11-09 02:53:17 |
| [xlsx](https://github.com/plandem/xlsx) | 170 | 25 | 12 | Fast and reliable way to work with Microsoft Excel™ [xlsx] files in Golang | 2017-08-26 23:11:38 | 2023-10-23 07:18:38 |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 21 | 7 | 0 | Golang bindings for libxlsxwriter for writing XLSX files | 2017-03-13 04:15:17 | 2023-07-11 01:40:47 |
</details>

### Miscellaneous - Dependency Injection
Libraries for working with dependency injection.

<sup>*Last Update: 2023-12-06 21:07:36*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fx](https://uber-go.github.io/fx/) | 4,685 | 269 | 48 | A dependency injection based application framework for Go. | 2016-10-27 00:25:00 | 2023-12-05 12:37:51 |
| [dig](https://go.uber.org/dig) | 3,499 | 202 | 25 | A reflection based dependency injection toolkit for Go. | 2017-03-21 23:55:50 | 2023-12-06 08:14:16 |
| [container](https://github.com/golobby/container) | 507 | 33 | 4 | A lightweight yet powerful IoC dependency injection container for the Go programming language | 2019-09-23 16:12:50 | 2023-11-28 21:58:33 |
| [di](https://pkg.go.dev/github.com/goioc/di/?tab=doc) | 310 | 14 | 1 | Simple and yet powerful Dependency Injection for Go | 2020-06-11 12:28:06 | 2023-11-16 13:09:21 |
| [di](https://github.com/defval/di) | 210 | 11 | 0 | 🛠 A full-featured dependency injection container for go programming language. | 2020-02-03 19:06:39 | 2023-11-23 06:40:50 |
| [dingo](https://github.com/i-love-flamingo/dingo) | 169 | 10 | 16 | Go Dependency Injection Framework | 2018-10-29 08:55:18 | 2023-12-05 23:25:32 |
| [alice](https://godoc.org/github.com/magic003/alice) | 51 | 4 | 0 | An additive dependency injection container for Golang. | 2017-04-08 16:25:21 | 2023-04-15 14:24:55 |
| [wire](https://github.com/Fs02/wire) | 38 | 8 | 1 | Strict Runtime Dependency Injection for Golang | 2018-07-05 10:42:24 | 2023-05-09 10:12:20 |
| [linker](https://github.com/logrange/linker) | 35 | 6 | 0 | Dependency Injection and Inversion of Control package | 2018-12-04 23:56:34 | 2023-03-10 16:25:52 |
| [gocontainer](https://github.com/vardius/gocontainer) | 19 | 2 | 0 | Simple Dependency Injection Container | 2019-06-06 08:18:07 | 2023-08-13 21:34:00 |
| [kinit](https://github.com/go-kata/kinit) | 9 | 1 | 0 | GO Dependency Injection | 2021-01-24 13:41:41 | 2023-09-05 15:50:49 |
| [nject](https://github.com/BlueOwlOpenSource/nject) | 3 | 1 | 0 | Go dependency injection: nject & npoint | 2018-10-31 18:15:43 | 2022-09-23 15:02:00 |
</details>

### Miscellaneous - Project Layout
Unofficial set of patterns for structuring projects.

<sup>*Last Update: 2023-12-10 21:13:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 43,328 | 4,781 | 96 | Standard Go Project Layout | 2017-09-09 16:33:26 | 2023-12-10 13:14:02 |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 1,696 | 208 | 22 | Modern Go Application example | 2018-09-14 12:19:02 | 2023-12-05 08:39:41 |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 643 | 208 | 1 | A Go project template | 2016-12-18 18:22:26 | 2023-12-09 11:32:22 |
| [seed](https://github.com/golang-templates/seed) | 409 | 46 | 1 | Go application GitHub repository template. | 2020-04-30 21:31:36 | 2023-12-05 21:42:55 |
| [go-todo-backend](https://github.com/Fs02/go-todo-backend) | 269 | 35 | 0 | Go Todo Backend example using modular project layout for product microservice. | 2020-06-25 14:28:50 | 2023-12-07 23:22:19 |
| [scaffold](https://github.com/catchplay/scaffold) | 146 | 27 | 2 | Generate scaffold project layout for Go. | 2018-12-11 10:36:03 | 2023-12-09 09:03:46 |
| [go-sample](https://github.com/zitryss/go-sample) | 123 | 24 | 0 | Go Project Sample Layout | 2019-01-24 23:41:46 | 2023-11-12 14:10:19 |
| [gobase](https://github.com/wajox/gobase) | 50 | 6 | 0 | This is a simple skeleton for golang applications | 2020-12-15 16:54:20 | 2023-12-08 10:39:40 |
| [inizio](https://github.com/insidieux/inizio) | 17 | 2 | 0 | Golang project standard layout generator | 2021-03-02 20:59:22 | 2023-11-13 12:15:11 |
</details>

### Miscellaneous - Strings
Libraries for working with strings.

<sup>*Last Update: 2023-12-06 21:07:52*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xstrings](https://github.com/huandu/xstrings) | 1,262 | 109 | 0 | Implements string functions widely used in other languages but absent in Go. | 2015-01-06 07:25:26 | 2023-12-03 12:25:59 |
| [strutil](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 197 | 24 | 1 | String utilities for Go | 2018-08-16 06:56:15 | 2023-11-28 09:25:07 |
| [stringy](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 193 | 17 | 1 | Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package. | 2020-04-03 03:34:10 | 2023-10-02 18:51:14 |
</details>

### Miscellaneous - Uncategorized
These libraries were placed here because none of the other categories seemed to fit.

<sup>*Last Update: 2023-11-25 19:13:49*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 9,619 | 1,557 | 184 | psutil for golang | 2014-04-18 07:35:28 | 2023-11-25 02:31:50 |
| [gatus](https://gatus.io) | 4,216 | 323 | 107 | ⛑ Automated developer-oriented status page | 2019-09-04 02:35:40 | 2023-11-25 00:38:49 |
| [archiver](https://pkg.go.dev/github.com/mholt/archiver/v4) | 4,114 | 376 | 13 | Easily create & extract archives, and compress & decompress files of various formats | 2016-04-08 22:46:55 | 2023-11-24 19:21:25 |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 3,414 | 252 | 7 | Random fake data generator written in go | 2015-04-24 04:45:59 | 2023-11-24 20:45:29 |
| [base64Captcha](https://captcha.mojotv.cn/.netlify/functions/captcha) | 1,901 | 263 | 20 | captcha of base64 image string | 2017-12-12 12:17:07 | 2023-11-24 08:06:13 |
| [go-resiliency](https://godoc.org/github.com/eapache/go-resiliency) | 1,895 | 135 | 2 | Resiliency patterns for golang | 2014-11-29 04:11:32 | 2023-11-24 19:15:53 |
| [gosms](https://github.com/haxpax/gosms) | 1,436 | 209 | 6 | :mailbox_closed: Your own local SMS gateway in Go | 2015-01-23 19:25:55 | 2023-11-12 23:56:31 |
| [ghorg](https://github.com/gabrie30/ghorg) | 1,379 | 150 | 9 | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Bitbucket, and more 🥚 | 2018-03-29 02:53:05 | 2023-11-25 11:24:42 |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 1,184 | 151 | 4 | a generic object pool for golang | 2015-12-28 14:26:23 | 2023-11-21 16:13:39 |
| [llvm](https://llir.github.io/document/) | 1,102 | 77 | 16 | Library for interacting with LLVM IR in pure Go. | 2014-09-19 11:18:44 | 2023-11-24 19:15:20 |
| [shortid](https://github.com/teris-io/shortid) | 899 | 68 | 3 | Super short, fully unique, non-sequential and URL friendly Ids | 2016-01-04 01:17:10 | 2023-11-20 03:31:07 |
| [shoutrrr](https://containrrr.dev/shoutrrr/) | 750 | 60 | 57 | Notification library for gophers and their furry friends. | 2019-04-11 06:49:34 | 2023-11-25 08:09:46 |
| [stateless](https://github.com/qmuntal/stateless) | 710 | 40 | 7 | Go library for creating finite state machines | 2019-09-11 08:19:18 | 2023-11-25 00:16:19 |
| [health](https://github.com/dimiro1/health) | 446 | 44 | 3 | An easy to use, extensible health check library for Go applications. | 2016-03-08 23:04:43 | 2023-11-25 04:38:37 |
| [banner](https://github.com/dimiro1/banner) | 440 | 25 | 0 | An easy way to add useful startup banners into your Go applications | 2016-03-25 21:28:44 | 2023-11-17 00:55:45 |
| [gountries](https://github.com/pariz/gountries) | 401 | 67 | 14 | Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data. | 2016-01-13 08:04:18 | 2023-11-12 23:27:36 |
| [go-conv](https://github.com/cstockton/go-conv) | 382 | 20 | 1 | Fast conversions across various Go types with a simple API. | 2016-10-11 07:41:41 | 2023-11-24 19:25:15 |
| [lk](https://github.com/hyperboloide/lk) | 326 | 56 | 0 | Simple licensing library for golang. | 2016-07-14 16:06:07 | 2023-11-20 10:07:08 |
| [ffmt](https://github.com/go-ffmt/ffmt) | 301 | 24 | 2 | Golang beautify data display for Humans | 2015-02-14 15:19:45 | 2023-11-19 11:29:10 |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 258 | 30 | 4 | An simple, easily extensible and concurrent health-check library for Go services | 2017-08-18 12:48:40 | 2023-11-13 11:15:04 |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 258 | 42 | 2 | Go bindings for unarr (decompression library for RAR, TAR, ZIP and 7z archives) | 2015-11-01 09:38:37 | 2023-11-22 14:27:47 |
| [antch](https://github.com/antchfx/antch) | 252 | 42 | 4 | Antch, a fast, powerful and extensible web crawling & scraping framework for Go | 2017-09-28 05:44:17 | 2023-10-22 17:05:29 |
| [battery](https://github.com/distatus/battery) | 229 | 35 | 5 | cross-platform, normalized battery information library | 2016-03-12 23:03:40 | 2023-11-18 23:47:47 |
| [bitio](https://github.com/icza/bitio) | 224 | 26 | 1 | Optimized bit-level Reader and Writer for Go. | 2016-05-31 10:02:30 | 2023-11-11 15:23:33 |
| [stats](https://github.com/go-playground/stats) | 169 | 19 | 1 | :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... | 2015-09-14 20:20:20 | 2023-10-13 11:02:38 |
| [turtle](https://github.com/hackebrot/turtle) | 152 | 11 | 2 | Emojis for Go 😄🐢🚀 | 2017-09-08 22:25:32 | 2023-08-20 12:04:46 |
| [captcha](https://pkg.go.dev/github.com/steambap/captcha) | 138 | 25 | 0 | :sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation | 2017-09-12 06:52:15 | 2023-10-08 08:50:39 |
| [gotoprom](https://github.com/cabify/gotoprom) | 107 | 2 | 1 | Type-safe Prometheus metrics builder library for golang | 2018-10-10 16:07:33 | 2023-11-24 19:43:54 |
| [gommit](https://github.com/antham/gommit) | 106 | 5 | 1 | Enforce git message commit consistency | 2016-08-30 11:10:11 | 2023-10-14 01:32:36 |
| [indigo](https://github.com/osamingo/indigo) | 104 | 12 | 0 | A distributed unique ID generator of using Sonyflake and encoded by Base58 | 2016-08-31 14:17:45 | 2023-07-31 01:45:55 |
| [faker](https://github.com/pioz/faker) | 84 | 11 | 0 | Random fake data and struct generator for Go. | 2020-07-22 20:09:46 | 2023-11-11 15:22:22 |
| [morse](https://dalw.in/morse) | 80 | 13 | 5 | Morse Code Library in Go | 2018-08-15 05:31:31 | 2023-10-11 09:36:45 |
| [persian](https://github.com/mavihq/persian) | 76 | 10 | 1 | Some utilities for Persian language in Go (Golang) | 2017-10-16 16:16:56 | 2023-10-20 11:02:07 |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 65 | 7 | 0 | HTTP service to generate PDF from Json requests | 2015-11-30 19:27:26 | 2023-10-21 14:06:57 |
| [xkg](https://godoc.org/github.com/go-xkg/xkg) | 56 | 6 | 1 | User level X Keyboard Grabber | 2015-01-05 01:04:43 | 2023-07-01 04:09:45 |
| [datacounter](https://github.com/miolini/datacounter) | 47 | 9 | 2 | Golang counters for readers/writers | 2015-10-14 19:15:50 | 2023-10-16 09:32:29 |
| [browscap_go](http://browscap.org/) | 45 | 29 | 8 | GoLang Library for Browser Capabilities Project | 2014-09-18 04:47:42 | 2023-11-06 23:56:17 |
| [url-shortener](https://github.com/pantrif/url-shortener) | 43 | 7 | 1 | A golang URL Shortener | 2018-06-04 05:57:45 | 2023-08-28 08:57:39 |
| [sandid](https://pkg.go.dev/github.com/aofei/sandid) | 41 | 6 | 0 | Every grain of sand on Earth has its own ID. | 2018-06-12 01:24:14 | 2023-10-11 09:38:14 |
| [autoflags](http://pkg.go.dev/github.com/artyom/autoflags) | 39 | 4 | 0 | Populate go command line app flags from config struct | 2014-05-15 19:00:29 | 2023-06-27 11:32:25 |
| [xdg](https://github.com/rkoesters/xdg) | 36 | 8 | 1 | FreeDesktop.org (xdg) Specs implemented in Go | 2013-12-15 09:51:51 | 2023-10-11 09:40:10 |
| [gosh](https://github.com/osamingo/gosh) | 35 | 2 | 0 | Provide Go Statistics Handler, Struct, Measure Method | 2018-05-25 08:55:55 | 2023-08-14 09:52:20 |
| [metrics](https://github.com/pascaldekloe/metrics) | 27 | 4 | 3 | atomic measures + Prometheus exposition library | 2019-01-29 09:39:18 | 2023-08-24 19:53:09 |
| [shellwords](https://pkg.go.dev/github.com/aofei/sandid) | 21 | 3 | 0 | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. | 2017-09-28 09:08:28 | 2023-07-14 09:23:37 |
| [numa](https://dalw.in/morse) | 18 | 5 | 0 | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. | 2018-12-10 09:59:13 | 2023-09-28 07:07:48 |
| [anagent](https://github.com/mudler/anagent) | 15 | 4 | 0 | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection | 2017-12-29 17:16:25 | 2023-08-31 12:34:40 |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 15 | 3 | 0 | Calculate average score and rating based on Wilson Score Equation | 2017-08-05 19:04:30 | 2023-09-28 14:01:07 |
| [hostutils](https://github.com/Wing924/hostutils) | 12 | 5 | 0 | A golang library for packing and unpacking hosts list | 2017-09-26 03:47:32 | 2022-09-27 08:47:45 |
| [generators](https://github.com/azr/generators) | 5 | 2 | 0 | #golang generator | 2016-02-29 14:29:02 | 2022-09-27 08:49:19 |
</details>

### Natural Language Processing
Libraries for working with human languages.

<sup>*Last Update: 2023-12-08 20:26:25*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prose](https://github.com/jdkato/prose) | 3,022 | 162 | 21 | :book: A Golang library for text processing, including tokenization, part-of-speech tagging, and named-entity extraction. | 2017-02-17 17:08:22 | 2023-12-04 07:15:52 |
| [go-i18n](https://github.com/nicksnyder/go-i18n) | 2,597 | 314 | 20 | Translate your Go program into multiple languages. | 2012-01-14 21:44:37 | 2023-12-08 11:21:57 |
| [gse](https://github.com/go-ego/gse) | 2,377 | 208 | 12 | Go efficient multilingual NLP and text segmentation; support English, Chinese, Japanese and others. | 2017-06-23 15:42:35 | 2023-12-08 10:48:28 |
| [gojieba](https://github.com/yanyiwu/gojieba) | 2,260 | 336 | 60 | "结巴"中文分词的Golang版本 | 2015-09-12 01:30:44 | 2023-12-08 03:10:56 |
| [spago](https://github.com/nlpodyssey/spago) | 1,668 | 84 | 12 | Self-contained Machine Learning and Natural Language Processing library in Go | 2020-01-05 20:39:29 | 2023-12-06 17:02:14 |
| [go-pinyin](https://godoc.org/github.com/mozillazg/go-pinyin) | 1,509 | 235 | 15 | 汉字转拼音 | 2014-11-09 14:04:33 | 2023-12-06 08:25:52 |
| [when](https://github.com/olebedev/when) | 1,301 | 84 | 17 | A natural language date/time parser with pluggable rules | 2016-12-27 13:11:46 | 2023-12-05 06:02:29 |
| [kagome](https://github.com/ikawaha/kagome) | 763 | 50 | 4 | Self-contained Japanese Morphological Analyzer written in pure Go | 2014-06-26 04:38:13 | 2023-12-04 18:13:55 |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 600 | 62 | 12 | Natural language detection library for Go | 2017-02-20 17:32:01 | 2023-12-06 14:31:25 |
| [nlp](https://github.com/james-bowman/nlp) | 422 | 46 | 5 | Selected Machine Learning algorithms for natural language processing and semantic analysis in Golang | 2017-03-15 08:28:05 | 2023-12-01 14:00:00 |
| [sentences](https://sentences-231000.appspot.com/) | 395 | 36 | 5 | A multilingual command line sentence tokenizer in Golang | 2015-08-07 01:08:20 | 2023-12-06 15:52:37 |
| [nlp](https://github.com/shixzie/nlp) | 386 | 35 | 3 | [UNMANTEINED] Extract values from strings and fill your structs with nlp. | 2017-01-25 07:19:03 | 2023-11-24 19:27:48 |
| [getlang](https://github.com/rylans/getlang) | 161 | 21 | 4 | Natural language detection package in pure Go | 2018-03-01 21:27:30 | 2023-11-24 12:02:46 |
| [go-unidecode](https://godoc.org/github.com/mozillazg/go-unidecode) | 115 | 20 | 5 | ASCII transliterations of Unicode text. | 2016-07-08 13:15:10 | 2023-12-06 06:12:47 |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 104 | 18 | 4 | A Go port of the Rapid Automatic Keyword Extraction algorithm (RAKE) | 2016-12-17 13:36:25 | 2023-11-17 03:41:26 |
| [go-nlp](https://github.com/nuance/go-nlp) | 95 | 12 | 0 | Utilities for working with discrete probability distributions and other tools useful for doing NLP work | 2011-05-02 06:43:36 | 2023-03-03 17:39:16 |
| [segment](https://github.com/blevesearch/segment) | 81 | 16 | 5 | A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 | 2014-10-16 19:24:26 | 2023-12-07 11:02:37 |
| [gounidecode](https://github.com/fiam/gounidecode) | 79 | 21 | 2 | Unicode transliterator for #golang | 2012-05-01 11:59:34 | 2023-07-25 20:07:55 |
| [go-stem](svn://go-stem) | 76 | 15 | 1 | Word Stemming in Go | 2011-09-23 19:07:23 | 2023-09-01 11:55:09 |
| [textcat](https://github.com/pebbe/textcat) | 70 | 11 | 1 | A Go package for n-gram based text categorization, with support for utf-8 and raw text | 2012-09-21 15:04:45 | 2023-05-10 13:31:24 |
| [address](https://pkg.go.dev/github.com/bojanz/address) | 69 | 3 | 0 | Address handling for Go. | 2020-10-07 18:15:27 | 2023-10-03 05:04:24 |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 62 | 14 | 1 | Chinese word splitting algorithm MMSEG in GO | 2012-04-18 04:06:21 | 2022-09-27 08:51:05 |
| [go-localize](https://github.com/m1/go-localize) | 55 | 11 | 3 | i18n (Internationalization and localization) engine written in Go, used for translating locale strings.  | 2019-12-23 12:02:51 | 2023-11-10 16:57:11 |
| [go2vec](https://godoc.org/github.com/mozillazg/go-unidecode) | 50 | 5 | 0 | Read and use word2vec vectors in Go | 2015-01-27 12:02:04 | 2023-10-19 10:32:39 |
| [stemmer](https://godoc.org/github.com/dchest/stemmer) | 50 | 6 | 0 | Stemmer packages for Go programming language. Includes English, German and Dutch stemmers. | 2011-03-21 02:08:12 | 2023-08-27 23:09:33 |
| [porter2](http://zhen.org/blog/generating-porter2-fsm-for-fun-and-performance/) | 46 | 7 | 1 | High Performance Porter2 Stemmer | 2015-01-21 07:30:32 | 2022-09-27 08:50:40 |
| [petrovich](https://github.com/striker2000/petrovich) | 44 | 6 | 0 | Golang port of Petrovich - an inflector for Russian anthroponyms. | 2016-12-26 22:50:38 | 2023-11-29 23:24:24 |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 41 | 6 | 1 | Transliterate Cyrillic → Latin in every possible way | 2020-04-27 09:29:40 | 2023-12-04 06:20:51 |
| [govader](https://github.com/jonreiter/govader) | 38 | 7 | 1 | vader sentiment analysis in go | 2020-01-19 10:06:15 | 2023-11-07 14:47:59 |
| [snowball](https://github.com/goodsign/snowball) | 36 | 5 | 0 | Cgo binding for Snowball C library | 2012-12-11 12:42:19 | 2023-08-18 21:23:00 |
| [transliterator](https://github.com/alexsergivan/transliterator) | 36 | 9 | 1 | Golang text Transliterator (i.e München -> Muenchen) | 2020-04-17 14:19:55 | 2023-11-14 11:31:22 |
| [mystem](https://github.com/dveselov/mystem) | 30 | 8 | 0 | CGo bindings to Yandex.Mystem | 2016-08-30 14:55:39 | 2023-11-29 11:44:58 |
| [paicehusk](https://github.com/rookii/paicehusk) | 29 | 7 | 2 | Golang implementation of the Paice/Husk Stemming Algorithm | 2012-09-29 16:06:58 | 2023-07-23 05:39:03 |
| [detectlanguage-go](https://detectlanguage.com) | 23 | 3 | 0 | Detect Language API Go Client | 2019-12-14 23:30:44 | 2023-08-13 19:06:10 |
| [icu](https://github.com/goodsign/icu) | 21 | 6 | 1 | Cgo binding for icu4c library | 2012-12-11 13:09:41 | 2023-10-04 07:11:36 |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 20 | 6 | 0 | Go bindings for the snowball libstemmer library including porter 2 | 2012-08-06 19:31:05 | 2022-09-27 08:50:18 |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 18 | 7 | 0 | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) | 2018-10-11 03:22:36 | 2023-07-11 19:32:41 |
| [shamoji](https://github.com/osamingo/shamoji) | 13 | 2 | 0 | The shamoji (杓文字) is a word filtering package | 2017-07-23 06:38:42 | 2022-09-27 08:51:16 |
| [libtextcat](https://github.com/goodsign/libtextcat) | 13 | 8 | 0 | Cgo binding for libtextcat C library | 2012-12-10 21:21:47 | 2023-08-04 09:58:01 |
| [porter](https://github.com/a2800276/porter) | 12 | 2 | 0 | porter stemmer | 2013-09-17 11:10:16 | 2022-12-19 17:32:36 |
| [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) | 11 | 2 | 0 | 💬 Sentiment analyzer library using SentiWordnet in Go | 2020-04-21 09:09:28 | 2023-08-25 15:07:32 |
</details>

### Networking
Libraries for working with various layers of the network.

<sup>*Last Update: 2023-11-25 19:13:46*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 20,389 | 1,695 | 85 | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http | 2015-10-18 22:19:57 | 2023-11-24 01:11:56 |
| [kcptun](https://github.com/xtaci/kcptun) | 13,524 | 2,555 | 114 | A Stable & Secure Tunnel based on KCP with N:M multiplexing and FEC. Available for ARM, MIPS, 386 and AMD64。N:M 多重化と FEC を備えた KCP に基づく安定した安全なトンネル。 N:M 다중화 및 FEC를 사용하는 KCP 기반의 안정적이고 안전한 터널입니다.  Un tunnel stable et sécurisé basé sur KCP avec multiplexage N:M et FEC. | 2016-02-26 09:54:46 | 2023-11-24 10:39:54 |
| [webrtc](https://pion.ly) | 11,963 | 1,557 | 106 | Pure Go implementation of the WebRTC API | 2018-05-18 23:10:05 | 2023-11-25 09:39:05 |
| [quic-go](https://github.com/quic-go/quic-go) | 8,905 | 1,245 | 162 | A QUIC implementation in pure go | 2016-04-06 20:16:27 | 2023-11-24 04:51:27 |
| [gnet](https://gnet.host) | 8,165 | 935 | 40 | 🚀 gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。 | 2019-02-24 03:48:45 | 2023-11-24 07:13:53 |
| [dns](https://miek.nl/2014/august/16/go-dns-package) | 7,315 | 1,164 | 13 | DNS library in Go | 2010-08-03 21:56:23 | 2023-11-24 04:17:24 |
| [gopacket](https://github.com/google/gopacket) | 5,905 | 1,129 | 326 | Provides packet processing capabilities for Go | 2015-03-16 20:46:00 | 2023-11-23 08:57:54 |
| [httplab](https://github.com/qustavo/httplab) | 3,929 | 132 | 13 | The interactive web server | 2017-02-08 17:13:19 | 2023-11-24 11:26:31 |
| [kcp-go](https://github.com/xtaci/kcp-go) | 3,793 | 727 | 50 |  A Crypto-Secure, Production-Grade Reliable-UDP Library for golang with FEC  | 2015-06-16 06:15:55 | 2023-11-24 09:54:04 |
| [gobgp](https://osrg.github.io/gobgp/) | 3,352 | 703 | 161 | BGP implemented in the Go Programming Language | 2014-09-14 01:51:58 | 2023-11-24 03:47:32 |
| [ssh](https://godoc.org/github.com/gliderlabs/ssh) | 3,196 | 433 | 43 | Easy SSH servers in Golang | 2016-10-03 21:53:44 | 2023-11-24 02:55:37 |
| [fortio](https://fortio.org) | 3,055 | 238 | 72 | Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats. | 2017-10-10 01:01:39 | 2023-11-24 07:08:17 |
| [nbio](https://github.com/lesismal/nbio) | 1,821 | 135 | 11 | Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. | 2020-01-25 11:46:54 | 2023-11-24 07:33:28 |
| [water](https://github.com/songgao/water) | 1,773 | 271 | 26 | A simple TUN/TAP library written in native Go. | 2013-03-25 20:06:52 | 2023-11-24 19:12:09 |
| [gev](https://github.com/Allenxuxu/gev) | 1,676 | 196 | 13 | 🚀Gev is a lightweight, fast non-blocking TCP network library / websocket server based on Reactor mode. Support custom protocols to quickly and easily build high-performance servers.  | 2019-09-01 12:16:18 | 2023-11-24 04:33:43 |
| [go-getter](https://github.com/hashicorp/go-getter) | 1,541 | 243 | 151 | Package for downloading things from a string URL using a variety of protocols. | 2015-10-12 23:17:07 | 2023-11-23 03:13:24 |
| [sftp](https://github.com/pkg/sftp) | 1,413 | 415 | 35 | SFTP support for the go.crypto/ssh package | 2013-11-05 04:36:00 | 2023-11-23 10:20:36 |
| [grab](https://github.com/cavaliergopher/grab) | 1,329 | 149 | 33 | A download manager package for Go | 2016-01-05 12:46:35 | 2023-11-22 19:35:08 |
| [nff-go](https://github.com/aregm/nff-go) | 1,328 | 199 | 65 | NFF-Go -Network Function Framework for GO (former YANFF) | 2017-03-29 17:07:29 | 2023-11-20 07:23:16 |
| [ftp](https://github.com/jlaffaye/ftp) | 1,185 | 381 | 20 | FTP client package for Go | 2011-05-06 18:31:51 | 2023-11-22 07:04:55 |
| [mdns](https://github.com/hashicorp/mdns) | 1,053 | 210 | 39 | Simple mDNS client/server library in Golang | 2014-01-29 19:39:18 | 2023-11-24 07:48:59 |
| [gosnmp](https://github.com/gosnmp/gosnmp) | 1,020 | 359 | 40 | An SNMP library written in Go | 2012-08-27 05:59:24 | 2023-11-22 07:29:23 |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 932 | 202 | 14 | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.x and V5 in golang | 2018-09-16 11:46:17 | 2023-11-21 07:24:01 |
| [vssh](https://github.com/yahoo/vssh) | 920 | 80 | 4 | Go Library to Execute Commands Over SSH at Scale | 2020-06-09 16:19:22 | 2023-11-24 20:03:33 |
| [cidranger](https://github.com/yl2chen/cidranger) | 862 | 103 | 9 | Fast IP to CIDR lookup in Golang | 2017-08-21 05:50:14 | 2023-11-15 10:38:11 |
| [lhttp](https://github.com/fanux/lhttp) | 691 | 148 | 6 | go websocket, a better way to buid your IM server | 2015-12-29 01:13:36 | 2023-11-07 02:26:54 |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 616 | 52 | 6 | Pure-Go library for cross-platform local peer discovery using UDP multicast :woman: :repeat: :woman: | 2018-04-22 23:59:37 | 2023-11-21 09:14:09 |
| [go-stun](https://github.com/ccding/go-stun) | 609 | 113 | 4 | A go implementation of the STUN client (RFC 3489 and RFC 5389) | 2013-08-17 22:16:33 | 2023-11-20 02:33:05 |
| [gaio](https://github.com/xtaci/gaio) | 526 | 65 | 15 | High performance async-io(proactor) networking for Golang。golangのための高性能非同期io(proactor)ネットワーキング | 2019-12-20 05:19:00 | 2023-11-22 13:02:36 |
| [gotcp](https://github.com/gansidui/gotcp) | 507 | 158 | 0 | A Go package for quickly building tcp servers | 2014-04-13 14:54:01 | 2023-10-07 02:14:18 |
| [stun](https://github.com/gortc/stun) | 487 | 55 | 4 | Fast RFC 5389 STUN implementation in go | 2016-04-24 17:46:38 | 2023-10-06 14:09:23 |
| [gopcap](https://github.com/akrennmair/gopcap) | 478 | 151 | 11 | A simple wrapper around libpcap for the Go programming language | 2009-11-19 10:13:48 | 2023-11-22 14:58:06 |
| [tcp_server](https://github.com/firstrow/tcp_server) | 423 | 148 | 4 | golang tcp server | 2014-10-13 20:38:42 | 2023-10-05 15:03:40 |
| [raw](https://github.com/mdlayher/raw) | 422 | 76 | 15 | Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. | 2015-07-06 16:11:47 | 2023-11-12 17:46:58 |
| [winrm](https://pion.ly) | 409 | 121 | 35 | Command-line tool and library for Windows remote command execution in Go | 2013-12-30 18:29:15 | 2023-11-24 04:30:59 |
| [ftpserverlib](https://github.com/fclairamb/ftpserverlib) | 377 | 83 | 5 | golang ftp server library | 2016-09-25 12:05:29 | 2023-11-09 00:50:02 |
| [arp](https://tools.ietf.org/html/rfc826) | 343 | 67 | 6 | Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. | 2015-07-06 18:50:34 | 2023-11-14 09:04:47 |
| [ethernet](https://en.wikipedia.org/wiki/Ethernet_frame) | 269 | 41 | 0 | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. MIT Licensed. | 2015-07-03 00:15:18 | 2023-10-02 07:33:18 |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 254 | 34 | 5 | A library to simplify writing applications using TCP sockets to stream protobuff messages | 2015-06-29 19:07:31 | 2023-09-11 03:19:30 |
| [gnxi](https://github.com/google/gnxi) | 245 | 118 | 25 | gNXI Tools - gRPC Network Management/Operations Interface Tools | 2017-09-26 08:19:41 | 2023-11-22 14:41:48 |
| [jazigo](https://github.com/udhos/jazigo) | 200 | 25 | 2 | Jazigo is a tool written in Go for retrieving configuration for multiple devices, similar to rancid, fetchconfig, oxidized, Sweet. | 2016-06-07 19:53:53 | 2023-11-20 08:33:03 |
| [utp](https://github.com/anacrolix/go-libutp) | 169 | 36 | 4 | Use anacrolix/go-libutp instead | 2015-03-20 04:39:22 | 2023-08-24 02:28:16 |
| [canopus](https://github.com/zubairhamed/canopus) | 152 | 41 | 43 | CoAP Client/Server implementing RFC 7252 for the Go Language | 2015-02-24 04:12:20 | 2023-05-10 09:46:13 |
| [xtcp](https://github.com/xfxdev/xtcp) | 146 | 32 | 0 | A TCP Server Framework with graceful shutdown, custom protocol. | 2016-03-31 16:50:14 | 2023-10-09 08:59:01 |
| [sslb](https://godoc.org/github.com/gliderlabs/ssh) | 145 | 29 | 10 | Golang Super Simple Load Balance | 2015-10-18 21:31:09 | 2023-08-20 03:14:23 |
| [iplib](https://github.com/c-robinson/iplib) | 125 | 23 | 0 | A library  for working with IP addresses and networks in Go | 2019-05-06 06:23:41 | 2023-11-09 12:22:58 |
| [ether](https://github.com/songgao/ether) | 79 | 7 | 0 | A Go package for sending and receiving ethernet frames. Currently supporting Linux, Freebsd, and OS X. | 2014-05-21 03:46:30 | 2023-05-12 09:30:15 |
| [dhcp6](https://tools.ietf.org/html/rfc3315) | 76 | 17 | 3 | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. | 2015-05-22 04:13:30 | 2023-08-26 03:57:58 |
| [packet](https://github.com/aerogo/packet) | 75 | 17 | 1 | :package: Send network packets over a TCP or UDP connection. | 2017-10-29 05:46:44 | 2023-10-23 12:20:57 |
| [go-powerdns](https://pkg.go.dev/github.com/joeig/go-powerdns/v3) | 74 | 24 | 1 | Go PowerDNS 4.x API Client | 2018-06-21 21:37:33 | 2023-11-06 09:45:09 |
| [portproxy](https://github.com/aybabtme/portproxy) | 53 | 15 | 0 | TCP proxy, highjacks HTTP to allow CORS | 2014-12-13 02:57:36 | 2023-09-21 02:02:31 |
| [linkio](https://github.com/ian-kent/linkio) | 52 | 7 | 0 | Simulate network link speed | 2014-12-24 10:50:03 | 2023-04-12 08:27:14 |
| [panoptes-stream](https://github.com/yahoo/panoptes-stream) | 39 | 10 | 2 | A cloud native distributed streaming network telemetry. | 2020-10-09 04:26:26 | 2023-11-13 11:52:41 |
| [golibwireshark](https://osrg.github.io/gobgp/) | 29 | 8 | 1 | BGP implemented in the Go Programming Language | 2015-11-16 06:48:41 | 2023-03-23 06:57:11 |
| [graval](https://github.com/koofr/graval) | 28 | 8 | 0 | An experimental go FTP server framework | 2014-04-22 19:17:18 | 2022-09-27 08:55:16 |
| [publicip](https://github.com/polera/publicip) | 27 | 8 | 0 | Go pkg for returning your public facing IP address. | 2016-12-28 19:31:07 | 2023-10-24 09:58:27 |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 20 | 5 | 0 | HTTP proxy handler and dialer | 2018-07-18 09:42:34 | 2023-09-08 17:42:53 |
| [goshark](https://github.com/sunwxg/goshark) | 18 | 4 | 0 | A simple wrapper around libpcap for the Go programming language | 2015-11-01 07:23:09 | 2023-11-19 20:58:34 |
| [llb](https://github.com/kirillDanshin/llb) | 14 | 3 | 0 | Simulate network link speed | 2016-02-21 06:30:17 | 2023-02-12 06:24:03 |
| [tspool](https://github.com/two/tspool) | 14 | 3 | 0 | tcp server pool | 2018-10-27 01:05:03 | 2022-09-27 08:57:28 |
| [gohooks](https://github.com/averageflow/gohooks) | 12 | 1 | 0 | GoHooks make it easy to send and consume secured web-hooks from a Go application | 2020-10-30 17:20:36 | 2021-07-16 09:56:57 |
</details>

### Networking - HTTP Clients
Libraries for making HTTP requests.

<sup>*Last Update: 2023-11-20 08:20:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [resty](https://github.com/go-resty/resty) | 8,641 | 694 | 68 | Simple HTTP and REST client library for Go | 2015-08-28 17:48:47 | 2023-11-18 19:22:30 |
| [heimdall](http://gojek.tech) | 2,513 | 248 | 47 | An enhanced HTTP client for Go | 2018-01-19 09:32:26 | 2023-11-16 05:50:42 |
| [grequests](https://github.com/levigross/grequests) | 1,996 | 136 | 28 | A Go "clone" of the great and famous Requests library | 2015-06-11 16:41:48 | 2023-11-18 19:20:20 |
| [sling](https://github.com/dghubble/sling) | 1,609 | 116 | 1 | A Go HTTP client library for creating and sending API requests | 2015-04-02 08:42:52 | 2023-11-18 19:23:26 |
| [gentleman](https://pkg.go.dev/github.com/h2non/gentleman?tab=doc) | 1,025 | 56 | 23 | Plugin-driven, extensible HTTP client toolkit for Go | 2016-02-21 23:00:24 | 2023-11-18 17:57:13 |
| [pester](https://github.com/sethgrid/pester) | 624 | 69 | 5 | Go (golang) http calls with retries and backoff  | 2015-05-20 13:50:49 | 2023-11-09 15:17:41 |
| [request](https://pkg.go.dev/github.com/monaco-io/request?tab=doc) | 265 | 28 | 2 | go request, go http client | 2020-03-25 06:24:18 | 2023-11-18 19:21:13 |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 64 | 14 | 0 | An enhanced and lightweight http client for Golang | 2019-12-14 11:22:19 | 2023-11-15 15:36:04 |
| [rq](https://github.com/ddo/rq) | 50 | 6 | 1 | A nicer interface for golang stdlib HTTP client | 2017-12-26 10:48:27 | 2023-09-13 09:20:49 |
| [httpretry](https://github.com/ybbus/httpretry) | 40 | 6 | 0 | Enriches the standard go http client with retry functionality. | 2020-02-05 10:17:42 | 2023-10-16 16:43:59 |
</details>

### ORM
Libraries that implement Object-Relational Mapping or datamapping techniques.

<sup>*Last Update: 2023-12-09 18:57:24*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gorm](https://gorm.io) | 34,257 | 3,814 | 283 | The fantastic ORM library for Golang, aims to be developer friendly | 2013-10-25 08:31:38 | 2023-12-09 11:44:19 |
| [ent](https://entgo.io) | 14,348 | 902 | 383 | An entity framework for Go | 2019-06-12 22:53:55 | 2023-12-09 08:25:18 |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 6,159 | 560 | 96 | Generate a Go ORM tailored to your database schema. | 2016-02-21 06:18:25 | 2023-12-08 18:26:21 |
| [pg](https://pg.uptrace.dev/) | 5,496 | 405 | 115 | Golang ORM with focus on PostgreSQL features and performance | 2013-04-24 12:31:41 | 2023-12-09 05:03:27 |
| [gorp](https://github.com/go-gorp/gorp) | 3,698 | 414 | 146 | Go Relational Persistence - an ORM-ish library for Go | 2012-01-04 19:50:09 | 2023-12-08 19:48:58 |
| [db](https://upper.io/) | 3,411 | 241 | 152 | Data access layer for PostgreSQL, CockroachDB, MySQL, SQLite and MongoDB with ORM-like features. | 2013-10-23 02:04:36 | 2023-12-08 03:25:18 |
| [gormt](https://xxjwxc.github.io/post/gormt/) | 2,276 | 373 | 57 | database to golang struct | 2019-05-05 13:10:26 | 2023-12-09 04:45:06 |
| [reform](https://gopkg.in/reform.v1) | 1,425 | 73 | 86 | A better ORM for Go, based on non-empty interfaces and code generation. | 2016-02-25 09:41:09 | 2023-11-26 08:39:04 |
| [pop](https://github.com/gobuffalo/pop) | 1,391 | 248 | 95 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2023-12-06 18:39:37 |
| [go-sqlbuilder](https://pkg.go.dev/github.com/huandu/go-sqlbuilder) | 1,131 | 103 | 5 | A flexible and powerful SQL string builder library plus a zero-config ORM. | 2017-12-27 16:37:48 | 2023-12-08 13:58:56 |
| [go-queryset](https://github.com/jirfag/go-queryset) | 715 | 72 | 20 | 100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood. | 2017-09-03 17:29:30 | 2023-12-05 13:40:07 |
| [rel](https://go-rel.github.io/) | 694 | 59 | 26 | :gem: Modern ORM for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API | 2019-10-06 07:08:01 | 2023-12-04 02:24:56 |
| [beego](beego.me) | 682 | 177 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2023-12-08 22:51:53 |
| [qbs](https://github.com/coocood/qbs) | 546 | 101 | 10 | QBS stands for Query By Struct. A Go ORM. | 2013-02-02 05:40:59 | 2023-10-28 10:53:34 |
| [zoom](https://github.com/albrow/zoom) | 304 | 28 | 2 | A blazing-fast datastore and querying engine for Go built on Redis. | 2013-07-17 00:32:34 | 2023-10-19 17:09:17 |
| [gosql](https://github.com/rushteam/gosql) | 175 | 23 | 6 | golang orm and sql builder | 2020-04-27 09:16:29 | 2023-11-24 02:27:31 |
| [grimoire](https://fs02.github.io/grimoire) | 160 | 18 | 0 | Database access layer for golang | 2018-03-05 16:52:20 | 2023-09-25 03:44:37 |
| [go-store](https://github.com/gosuri/go-store) | 112 | 9 | 1 | A simple and fast Redis backed key-value store library for Go | 2015-03-22 12:07:29 | 2023-09-25 03:42:25 |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 47 | 8 | 0 | Simple Go ORM for Google/Firebase Cloud Firestore | 2018-12-04 14:53:53 | 2023-09-25 03:41:53 |
| [marlow](https://github.com/marlow/marlow) | 14 | 3 | 0 | persistence layer code generation for golang | 2020-08-11 13:34:00 | 2023-10-27 20:13:41 |
| [lore](https://github.com/abrahambotros/lore) | 14 | 3 | 0 | Light Object-Relational Environment (LORE) provides a simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go | 2017-04-29 03:57:15 | 2023-09-25 08:03:17 |
</details>

### OpenGL
Libraries for using OpenGL in Go.

<sup>*Last Update: 2023-11-11 19:52:50*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glfw](https://github.com/go-gl/glfw) | 1,451 | 208 | 25 | Go bindings for GLFW 3 | 2013-05-19 06:38:45 | 2023-11-10 16:53:32 |
| [gl](https://github.com/go-gl/gl) | 992 | 68 | 17 | Go bindings for OpenGL (generated via glow) | 2015-02-22 03:29:45 | 2023-11-04 20:06:23 |
| [mathgl](https://github.com/go-gl/mathgl) | 509 | 65 | 10 | A pure Go 3D math library. | 2013-02-13 14:18:55 | 2023-11-03 09:03:25 |
| [gl](https://github.com/goxjs/gl) | 168 | 28 | 8 | Go cross-platform OpenGL bindings. | 2015-05-18 08:10:15 | 2023-10-04 08:25:18 |
| [glfw](https://github.com/goxjs/glfw) | 79 | 26 | 6 | Go cross-platform glfw library for creating an OpenGL context and receiving events. | 2014-12-27 22:40:24 | 2023-09-17 18:54:49 |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 7 | 3 | 0 | go-glmatrix is a golang version of glMatrix, which is "designed to perform vector and matrix operations stupidly fast". | 2020-07-02 13:40:40 | 2023-03-04 15:15:00 |
</details>

### Package Management - Official
Official experimental tooling for package management

<sup>*Last Update: 2023-11-20 08:20:47*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [dep](https://golang.github.io/dep/) | 12,921 | 1,135 | 0 | Go dependency management tool experiment (deprecated) | 2016-10-07 00:04:51 | 2023-11-20 00:57:51 |
</details>

### Package Management - Unofficial
Unofficial libraries for package and dependency management

<sup>*Last Update: 2023-11-22 20:35:31*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glide](https://glide.sh) | 8,164 | 554 | 413 | Package Management for Golang | 2014-07-09 06:02:50 | 2023-11-19 11:31:40 |
| [godep](http://godoc.org/github.com/tools/godep) | 5,568 | 517 | 79 | dependency tool for go | 2013-05-01 07:55:35 | 2023-11-06 21:09:41 |
| [govendor](https://blog.golang.org/migrating-to-go-modules) | 4,943 | 445 | 122 | Use Go Modules. | 2015-04-12 15:26:40 | 2023-11-20 18:02:16 |
| [gopm](https://github.com/gpmgo/gopm) | 2,473 | 242 | 0 | Go Package Manager (gopm) is a package manager and build tool for Go. | 2013-05-15 14:53:29 | 2023-10-23 11:04:52 |
| [gom](https://github.com/mattn/gom) | 1,389 | 141 | 14 | Go Manager - bundle for go | 2013-09-11 02:08:59 | 2023-10-09 14:32:20 |
| [gpm](https://github.com/pote/gpm) | 1,196 | 51 | 11 | Barebones dependency manager for Go. | 2013-09-05 02:24:02 | 2023-10-19 10:35:03 |
| [goop](https://github.com/petejkim/goop) | 780 | 45 | 28 | A simple dependency manager for Go (golang), inspired by Bundler. | 2014-06-18 01:55:24 | 2023-06-17 14:27:31 |
| [modgv](https://github.com/lucasepe/modgv) | 476 | 21 | 1 | Converts 'go mod graph' output into Graphviz's DOT language | 2020-09-12 16:23:46 | 2023-11-19 03:42:13 |
| [nut](https://github.com/jingweno/nut) | 234 | 11 | 14 | Vendor Go dependencies | 2015-01-23 14:46:32 | 2023-08-23 03:18:56 |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 214 | 7 | 3 | Barebones dependency manager for Go. | 2013-07-19 15:20:47 | 2023-06-01 00:53:15 |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 157 | 31 | 1 | maven plugin to automate GoSDK load and build of projects | 2016-03-24 06:47:08 | 2023-11-16 14:22:28 |
| [VenGO](https://github.com/DamnWidget/VenGO) | 123 | 11 | 3 | Create and manage Isolated Virtual Environments for Go | 2014-10-17 15:19:03 | 2023-10-10 16:39:21 |
| [gop](https://github.com/lunny/gop) | 50 | 7 | 10 | Moved to https://gitea.com/lunny/gop | 2017-02-18 04:33:48 | 2023-01-28 18:24:13 |
</details>

### Performance


<sup>*Last Update: 2023-12-18 21:43:07*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [jaeger](https://www.jaegertracing.io/) | 18,770 | 2,295 | 351 | CNCF Jaeger, a Distributed Tracing Platform | 2016-04-15 18:49:02 | 2023-12-17 13:02:55 |
| [pixie](https://px.dev) | 5,024 | 388 | 244 | Instant Kubernetes-Native Application Observability | 2020-02-27 00:22:45 | 2023-12-17 06:42:26 |
| [statsviz](https://github.com/arl/statsviz) | 3,019 | 110 | 6 | 🚀 Visualise your Go program runtime metrics in real time in the browser | 2020-08-14 00:00:41 | 2023-12-16 15:03:54 |
| [profile](https://px.dev) | 1,922 | 167 | 9 | Simple profiling for Go | 2014-10-22 01:35:18 | 2023-12-16 18:13:06 |
| [tracer](https://github.com/kamilsk/tracer) | 82 | 4 | 11 | 🪡 Dead simple, lightweight tracing. | 2019-06-22 13:23:27 | 2023-11-18 09:09:13 |
</details>

### Query Language


<sup>*Last Update: 2023-11-19 20:35:36*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [graphql](https://github.com/graphql-go/graphql) | 9,498 | 865 | 217 | An implementation of GraphQL for Go / Golang | 2015-07-19 12:25:43 | 2023-11-19 11:22:12 |
| [gqlgen](https://gqlgen.com) | 9,301 | 1,133 | 258 | go generate based graphql server library | 2018-02-11 04:54:11 | 2023-11-19 12:31:57 |
| [dasel](https://daseldocs.tomwright.me) | 4,617 | 116 | 32 | Select, put and delete data from JSON, TOML, YAML, XML and CSV files with a single tool. Supports conversion between formats and can be used as a Go package. | 2020-09-22 10:33:56 | 2023-11-18 19:18:08 |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 4,542 | 536 | 44 | GraphQL server with a focus on ease of use | 2016-10-18 13:57:24 | 2023-11-19 11:23:37 |
| [gojsonq](https://github.com/thedevsaddam/gojsonq/wiki) | 2,112 | 187 | 26 | A simple Go package to Query over JSON/YAML/XML/CSV Data  | 2018-05-19 16:15:18 | 2023-11-16 23:12:13 |
| [rql](https://github.com/a8m/rql) | 316 | 42 | 15 | Resource Query Language for REST | 2018-06-05 18:37:29 | 2023-11-07 02:20:26 |
| [jsonql](https://github.com/elgs/jsonql) | 270 | 40 | 5 | JSON query expression library in Golang. | 2015-12-29 11:24:04 | 2023-10-04 08:55:56 |
| [jsonslice](https://github.com/bhmj/jsonslice) | 79 | 8 | 3 | json slicer | 2018-05-02 00:33:15 | 2023-07-18 10:16:39 |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 61 | 18 | 4 | Query Parser for REST | 2020-02-10 17:58:42 | 2023-11-12 09:30:23 |
| [graphql](https://github.com/tmc/graphql) | 57 | 7 | 3 | graphql parser + utilities | 2015-04-18 21:05:52 | 2023-08-01 19:59:34 |
| [api-fu](https://github.com/ccbrown/api-fu) | 52 | 5 | 3 | A collection of Go packages for creating robust GraphQL APIs | 2019-07-30 05:18:43 | 2023-09-21 00:16:31 |
| [straf](https://github.com/ThundR67/straf) | 36 | 6 | 0 | Convert Golang Struct To GraphQL Object On The Fly | 2019-08-16 13:31:39 | 2023-10-19 21:15:40 |
| [jsonpath](https://github.com/AsaiYusuke/jsonpath) | 19 | 3 | 1 | A query library for retrieving part of JSON based on JSONPath syntax. | 2020-11-29 05:37:26 | 2023-10-09 21:01:34 |
| [gws](https://github.com/Zaba505/gws) | 7 | 2 | 2 | A WebSocket client and server for GraphQL | 2020-06-08 19:51:36 | 2023-06-08 02:43:37 |
</details>

### Resource Embedding


<sup>*Last Update: 2023-12-05 20:35:23*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [statik](https://github.com/rakyll/statik) | 3,676 | 260 | 40 | Embed files into a Go executable | 2014-02-04 14:54:51 | 2023-11-28 16:33:29 |
| [packr](https://github.com/gobuffalo/packr) | 3,407 | 196 | 67 | The simple and easy way to embed static files into Go binaries. | 2017-03-15 22:24:53 | 2023-11-24 19:29:17 |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2,383 | 232 | 36 | go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. | 2013-10-23 21:29:34 | 2023-11-29 11:46:04 |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 979 | 85 | 32 | Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. | 2015-05-18 13:03:02 | 2023-11-26 02:52:13 |
| [esc](http://godoc.org/github.com/mjibson/esc) | 638 | 75 | 10 | A simple file embedder for Go | 2014-01-26 05:08:04 | 2023-12-01 08:52:44 |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 630 | 53 | 15 | a better customizable tool to embed files in go; also update embedded files remotely without restarting the server | 2016-01-23 20:19:33 | 2023-11-25 04:24:39 |
| [go-resources](https://github.com/omeid/go-resources) | 176 | 17 | 0 | Unfancy resources embedding for Go with out of box http.FileSystem support. | 2015-02-21 15:40:17 | 2023-08-04 05:19:16 |
| [statics](https://github.com/go-playground/statics) | 67 | 7 | 0 | :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks | 2015-10-07 11:49:52 | 2023-10-03 00:25:47 |
| [templify](https://github.com/wlbr/templify) | 30 | 6 | 1 | A tool to be used with 'go generate' to embed external template files into Go code. | 2016-05-22 16:42:47 | 2023-07-26 06:48:09 |
| [rebed](https://github.com/soypat/rebed) | 28 | 2 | 0 | Recreates directory and files from embedded filesystem using Go 1.16 embed.FS type. | 2021-02-17 18:19:49 | 2023-09-18 14:42:50 |
| [debme](https://github.com/leaanthony/debme) | 27 | 6 | 1 | embed.FS wrapper providing additional functionality | 2021-04-16 00:25:13 | 2023-11-02 12:30:35 |
| [mule](https://github.com/wlbr/mule) | 12 | 4 | 1 | mule is a tool to be used with 'go generate' to embed external resources files into Go code. | 2020-01-17 10:56:00 | 2022-09-27 09:06:14 |
</details>

### Science and Data Analysis
Libraries for scientific computing and data analyzing.

<sup>*Last Update: 2023-10-30 20:27:04*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gonum](https://www.gonum.org/) | 6,888 | 554 | 228 | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more | 2017-03-25 14:54:38 | 2023-10-30 03:00:31 |
| [stats](https://github.com/montanaflynn/stats) | 2,825 | 173 | 17 | A well tested and comprehensive Golang statistics library package with no dependencies. | 2014-12-16 03:25:19 | 2023-10-30 02:13:53 |
| [plot](https://github.com/gonum/plot) | 2,537 | 207 | 86 | A repository for plotting and visualizing data | 2013-07-23 07:01:13 | 2023-10-30 07:34:10 |
| [gosl](https://github.com/cpmech/gosl) | 1,774 | 192 | 0 | Linear algebra, eigenvalues, FFT, Bessel, elliptic, orthogonal polys, geometry, NURBS, numerical quadrature, 3D transfinite interpolation, random numbers, Mersenne twister, probability distributions, optimisation, differential equations. | 2015-02-09 23:00:38 | 2023-10-23 11:08:52 |
| [streamtools](http://nytlabs.github.io/streamtools/) | 1,311 | 111 | 47 | tools for working with streams of data | 2013-07-05 18:58:45 | 2023-10-23 11:05:06 |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 1,056 | 90 | 15 | DataFrames for Go: For statistics, machine-learning, and data manipulation/exploration | 2018-10-01 12:19:31 | 2023-10-29 21:51:40 |
| [go-dsp](http://godoc.org/github.com/mjibson/go-dsp) | 831 | 91 | 7 | Digital Signal Processing for Go | 2011-11-02 06:28:41 | 2023-10-26 19:35:15 |
| [chart](https://github.com/vdobler/chart) | 754 | 105 | 6 | Provide basic charts in go | 2011-06-27 12:19:42 | 2023-10-27 05:27:58 |
| [orb](https://github.com/paulmach/orb) | 745 | 90 | 15 | Types and utilities for working with 2d geometry in Golang | 2016-03-28 01:19:01 | 2023-10-26 05:42:35 |
| [goraph](https://github.com/gyuho/goraph) | 723 | 107 | 5 | Package goraph implements graph data structure and algorithms. | 2014-02-27 03:15:55 | 2023-09-25 16:20:03 |
| [graph](https://yourbasic.org/golang/your-basic-func/) | 660 | 91 | 6 | Graph algorithms and data structures | 2017-04-27 18:43:54 | 2023-10-23 09:04:12 |
| [ewma](https://github.com/VividCortex/ewma) | 414 | 34 | 5 | Exponentially Weighted Moving Average algorithms for Go. | 2013-07-05 21:33:25 | 2023-10-27 15:06:26 |
| [calendarheatmap](https://calendarheatmap.io) | 381 | 17 | 12 | 📅 Calendar heatmap inspired by GitHub contribution activity  | 2020-07-01 18:30:48 | 2023-10-25 05:12:47 |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 185 | 22 | 5 | :wink: :cyclone: :strawberry: TextRank implementation in Golang with extendable features (summarization, phrase extraction) and multithreading (goroutine). | 2018-01-09 19:36:17 | 2023-10-25 15:41:25 |
| [gohistogram](http://godoc.org/github.com/VividCortex/gohistogram) | 172 | 31 | 4 | Streaming approximate histograms in Go | 2013-07-02 12:53:22 | 2023-06-23 03:58:27 |
| [sparse](https://github.com/james-bowman/sparse) | 145 | 22 | 4 | Sparse matrix formats for linear algebra supporting scientific and machine learning applications | 2017-05-16 13:54:36 | 2023-08-23 11:24:52 |
| [go-estimate](https://github.com/milosgajdos/go-estimate) | 104 | 8 | 2 | State estimation and filtering algorithms in Go | 2018-11-04 22:32:52 | 2023-03-03 21:01:18 |
| [pagerank](https://github.com/alixaxel/pagerank) | 78 | 21 | 3 | Weighted PageRank implementation in Go | 2015-08-06 01:33:34 | 2023-08-29 17:43:40 |
| [geom](https://github.com/skelterjohn/geom) | 55 | 18 | 1 | 2d geometry for golang | 2011-06-07 17:49:11 | 2023-04-02 02:39:45 |
| [evaler](https://github.com/soniah/evaler) | 52 | 17 | 5 | Implements a simple floating point arithmetic expression evaluator in Go (golang). | 2012-09-04 23:37:58 | 2023-07-30 06:09:50 |
| [decimal](https://github.com/db47h/decimal) | 38 | 3 | 0 | An arbitrary-precision decimal floating-point arithmetic package for Go | 2020-05-27 15:23:59 | 2023-08-10 09:46:31 |
| [triangolatte](https://github.com/tchayen/triangolatte) | 35 | 3 | 4 | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. | 2018-07-18 21:17:09 | 2023-04-03 14:18:57 |
| [goent](https://github.com/kzahedi/goent) | 33 | 4 | 0 | GO Implementation of Entropy Measures | 2017-08-08 05:37:12 | 2023-05-24 07:52:16 |
| [piecewiselinear](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 26 | 3 | 0 | tiny linear interpolation library for go (factored out from https://github.com/sgreben/yeetgif) | 2018-10-21 13:19:44 | 2023-08-15 21:45:38 |
| [godesim](https://github.com/soypat/godesim) | 21 | 1 | 1 | ODE system solver made simple. For IVPs (initial value problems). | 2020-12-16 01:02:26 | 2023-01-05 01:18:21 |
| [ode](https://yourbasic.org/golang/your-basic-func/) | 21 | 2 | 1 | An ordinary differential equation solving library in golang. | 2016-11-11 22:40:21 | 2023-04-02 02:41:27 |
| [GoStats](https://github.com/OGFris/GoStats) | 20 | 3 | 0 | GoStats is a go library for math statistics mostly used in ML domains, it covers most of the statistical measures functions. | 2018-07-22 20:55:16 | 2023-09-07 08:34:24 |
| [PiHex](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 20 | 4 | 1 | PiHex Library, written in Go, generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000. | 2016-07-22 11:21:37 | 2023-03-05 13:19:33 |
| [assocentity](https://pkg.go.dev/github.com/ndabAP/assocentity/v14) | 12 | 3 | 3 | Package assocentity returns the mean distance from tokens to an entity and its synonyms | 2018-12-21 07:17:09 | 2023-06-14 03:18:36 |
| [rootfinding](https://github.com/khezen/rootfinding) | 10 | 2 | 0 | root-finding library | 2018-10-30 22:31:48 | 2022-09-27 09:08:00 |
| [go-gt](https://github.com/ThePaw/go-gt) | 10 | 2 | 2 | Automatically exported from code.google.com/p/go-gt | 2015-09-14 12:05:37 | 2023-04-07 11:01:36 |
| [bradleyterry](https://pkg.go.dev/github.com/ndabAP/assocentity/v14) | 8 | 2 | 0 | Package to do Bradley-Terry Model pairwise compairsons | 2019-04-30 00:28:13 | 2023-01-20 15:32:53 |
</details>

### Security
Libraries that are used to help make your application more secure.

<sup>*Last Update: 2023-12-05 20:35:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [lego](https://go-acme.github.io/lego/) | 6,567 | 925 | 159 | Let's Encrypt/ACME client and library written in Go | 2015-06-08 00:36:41 | 2023-12-05 10:05:22 |
| [cameradar](https://github.com/Ullaakut/cameradar) | 3,693 | 502 | 32 | Cameradar hacks its way into RTSP videosurveillance cameras | 2016-05-20 11:35:41 | 2023-12-05 12:47:50 |
| [crypto](https://golang.org/x/crypto) | 2,840 | 2,563 | 73 | [mirror] Go supplementary cryptography libraries | 2014-12-04 04:02:55 | 2023-12-05 06:06:06 |
| [memguard](https://github.com/awnumar/memguard) | 2,424 | 125 | 7 | Secure software enclave for storage of sensitive information in memory. | 2017-04-22 07:40:40 | 2023-12-05 00:38:17 |
| [secure](https://github.com/unrolled/secure) | 2,167 | 133 | 0 | HTTP middleware for Go that facilitates some quick security wins. | 2014-05-20 19:46:28 | 2023-12-01 11:14:22 |
| [acmetool](https://hlandau.github.io/acmetool/) | 2,015 | 133 | 71 | :lock: acmetool, an automatic certificate acquisition tool for ACME (Let's Encrypt) | 2015-11-15 01:56:02 | 2023-11-29 13:19:47 |
| [themis](https://www.cossacklabs.com/themis) | 1,778 | 142 | 24 | Easy to use cryptographic framework for data protection: secure messaging with forward secrecy and secure data storage. Has unified APIs across 14 platforms. | 2015-05-06 13:25:25 | 2023-12-03 13:14:34 |
| [acra](https://www.cossacklabs.com/acra/) | 1,234 | 124 | 5 | Database security suite. Database proxy with field-level encryption, search through encrypted data, SQL injections prevention, intrusion detection, honeypots. Supports client-side and proxy-side ("transparent") encryption. SQL, NoSQL. | 2016-11-14 16:23:25 | 2023-11-25 00:52:44 |
| [nacl](https://godoc.org/github.com/kevinburke/nacl) | 536 | 31 | 4 | Pure Go implementation of the NaCL set of API's | 2017-07-20 19:07:19 | 2023-11-24 19:32:39 |
| [go-password-validator](https://blog.boot.dev/open-source/how-to-validate-passwords/) | 447 | 36 | 3 | Validate the Strength of a Password in Go | 2020-10-14 15:52:14 | 2023-12-05 11:52:37 |
| [ssh-vault](https://ssh-vault.com) | 411 | 29 | 0 | 🌰  encrypt/decrypt using ssh keys | 2016-09-29 14:46:30 | 2023-12-01 13:26:47 |
| [optimus-go](https://github.com/pjebs/optimus-go) | 350 | 22 | 1 | ID hashing and Obfuscation using Knuth's Algorithm | 2015-05-05 10:12:38 | 2023-11-26 08:39:37 |
| [firewalld-rest](https://pkg.go.dev/github.com/prashantgupta24/firewalld-rest) | 332 | 19 | 2 | A rest application to update firewalld rules on a linux server  | 2020-06-12 20:16:33 | 2023-11-24 20:03:41 |
| [go-yara](https://github.com/hillu/go-yara) | 329 | 111 | 7 | Go bindings for YARA | 2015-01-25 01:01:11 | 2023-12-01 14:24:34 |
| [badactor](https://badactor.org) | 318 | 17 | 1 | BadActor.org An in-memory application driven jailer written in Go | 2014-12-12 20:05:20 | 2023-10-18 16:35:21 |
| [passlib](https://github.com/hlandau/passlib) | 287 | 30 | 1 | :key: Idiotproof golang password validation library inspired by Python's passlib | 2014-12-21 17:45:52 | 2023-10-29 14:49:09 |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 190 | 26 | 3 | A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go 🔑 | 2015-04-14 06:52:21 | 2023-09-18 16:14:04 |
| [argon2pw](https://github.com/raja/argon2pw) | 90 | 9 | 1 | Argon2 password hashing package for go with constant time hash comparison | 2018-03-13 13:56:36 | 2023-08-23 17:44:52 |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 58 | 9 | 0 | A probably paranoid Golang utility library for securely hashing and encrypting passwords based on the Dropbox method. This implementation uses Blake2b, Scrypt and XSalsa20-Poly1305 (via NaCl SecretBox) to create secure password hashes that are also encrypted using a master passphrase. | 2017-10-19 19:34:45 | 2023-11-07 20:49:26 |
| [go-generate-password](https://github.com/m1/go-generate-password) | 50 | 7 | 0 | Password generator written in Golang, usable as a CLI or Go library. Provides options for human readable and accessibility friendly passwords.  | 2019-11-14 17:57:19 | 2023-09-04 11:45:38 |
| [certificates](https://github.com/mvmaasakkers/certificates) | 36 | 8 | 0 | An opinionated helper for generating tls certificates | 2019-03-04 07:20:36 | 2023-08-28 13:56:52 |
| [secureio](https://github.com/xaionaro-go/secureio) | 30 | 5 | 1 | An easy-to-use XChaCha20-encryption wrapper for io.ReadWriteCloser (even lossy UDP) using ECDH key exchange algorithm, ED25519 signatures and Blake3+Poly1305 checksums/message-authentication for Go (golang). Also a multiplexer. | 2018-12-25 14:20:59 | 2023-05-18 13:31:11 |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 21 | 4 | 0 | A layer of abstraction the around acme/autocert certificate manager (Golang) | 2019-04-02 17:35:38 | 2023-09-18 16:09:34 |
| [argon2-hashing](https://www.cossacklabs.com/acra/) | 19 | 4 | 0 | A light package for generating and comparing password hashing with argon2 in Go | 2019-01-02 20:41:02 | 2023-04-25 09:45:35 |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 16 | 7 | 1 | goArgonPass is a Argon2 Password utility package for Go using the crypto library package Argon2 designed to be compatible with Passlib for Python and Argon2 PHP. Argon2 was the winner of the most recent Password Hashing Competition. This is designed for use anywhere password hashing and verification might be needed and is intended to replace implementations using bcrypt or Scrypt. | 2018-05-30 01:32:10 | 2022-09-27 09:19:46 |
</details>

### Serialization
Libraries and tools for binary serialization.

<sup>*Last Update: 2023-12-09 18:57:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go](http://jsoniter.com/migrate-from-go-std.html) | 12,766 | 1,069 | 248 | A high-performance 100% compatible drop-in replacement of "encoding/json" | 2016-11-30 00:30:24 | 2023-12-09 06:33:04 |
| [protobuf](https://github.com/golang/protobuf) | 9,375 | 1,646 | 96 | Go support for Google's protocol buffers | 2014-11-23 23:07:23 | 2023-12-09 04:42:59 |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 7,316 | 684 | 84 | Go library for decoding generic map values into native Go structures and vice versa. | 2013-05-20 05:24:34 | 2023-12-09 05:07:18 |
| [protobuf](https://github.com/gogo/protobuf) | 5,608 | 829 | 236 | [Deprecated] Protocol Buffers for Go with Gadgets | 2014-12-03 11:27:10 | 2023-12-09 08:55:42 |
| [go](https://github.com/ugorji/go) | 1,784 | 365 | 1 | idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] | 2013-05-30 02:13:13 | 2023-12-04 16:33:22 |
| [csvutil](https://github.com/jszwec/csvutil) | 844 | 58 | 0 | csvutil provides fast and idiomatic mapping between CSV and Go (golang) values. | 2017-10-30 04:09:48 | 2023-12-06 00:15:22 |
| [colfer](https://github.com/pascaldekloe/colfer) | 724 | 56 | 9 | binary serialization format | 2015-09-05 16:42:41 | 2023-12-09 11:57:25 |
| [cbor](https://github.com/fxamacker/cbor) | 619 | 57 | 18 | CBOR codec (RFC 8949) with CBOR tags, Go struct tags (toarray, keyasint, omitempty), float64/32/16, big.Int, and fuzz tested billions of execs.  | 2019-05-15 21:22:15 | 2023-12-04 20:32:56 |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 288 | 21 | 1 | Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. | 2013-11-07 20:28:12 | 2023-09-08 16:43:17 |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 160 | 45 | 3 | PHP session encoder/decoder written in Go | 2012-12-23 14:04:25 | 2023-11-21 14:53:20 |
| [structomap](https://godoc.org/github.com/danhper/structomap) | 143 | 11 | 0 | Easily and dynamically generate maps from Go static structures | 2015-05-13 08:54:11 | 2023-11-10 03:46:41 |
| [binstruct](https://github.com/ghostiam/binstruct) | 80 | 19 | 0 | Golang binary decoder for mapping data into the structure | 2018-10-23 15:42:22 | 2023-11-27 01:03:46 |
| [bambam](https://github.com/glycerine/bambam) | 67 | 12 | 3 | auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto | 2014-09-17 14:39:12 | 2023-09-08 16:50:41 |
| [asn1](https://github.com/Logicalis/asn1) | 53 | 27 | 6 | Asn.1 BER and DER encoding library for golang. | 2016-02-29 13:00:25 | 2023-01-28 18:49:45 |
| [bel](https://github.com/csweichel/bel) | 35 | 9 | 4 | Generate TypeScript interfaces from Go structs/interfaces - useful for JSON RPC | 2019-02-20 20:48:24 | 2023-06-16 22:34:38 |
| [fwencoder](https://github.com/o1egl/fwencoder) | 26 | 10 | 1 | Fixed width file parser (encoder/decoder) in GO (golang) | 2017-12-25 12:55:29 | 2023-11-27 01:03:26 |
| [elastic](https://github.com/epiclabs-io/elastic) | 23 | 4 | 1 | Converts go types no matter what | 2020-02-25 19:55:00 | 2023-09-29 08:28:04 |
| [pletter](https://github.com/vimeda/pletter) | 19 | 3 | 9 | A standard way to wrap a proto message | 2019-07-09 14:02:08 | 2023-08-22 11:20:27 |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 9 | 2 | 0 | A Go package for encode/decode fixed-width data | 2019-08-11 03:42:24 | 2023-10-01 22:40:50 |
| [unitpacking](https://github.com/recolude/unitpacking) | 6 | 1 | 0 | A library for storing unit vectors in a representation that lends itself to saving space on disk. | 2021-01-17 22:31:41 | 2023-02-06 17:20:12 |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 4 | 2 | 0 | go-lctree provides a CLI and Go primitives to serialize and deserialize LeetCode binary trees (e.g. "[5,4,7,3,null,2,null,-1,null,9]"). | 2020-05-04 05:39:46 | 2022-09-27 09:21:38 |
</details>

### Server Applications


<sup>*Last Update: 2023-11-19 20:35:29*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [caddy](https://caddyserver.com) | 50,604 | 3,855 | 138 | Fast and extensible multi-platform HTTP/1-2-3 web server with automatic HTTPS | 2015-01-13 19:45:03 | 2023-11-17 00:27:45 |
| [etcd](https://etcd.io) | 44,906 | 9,571 | 258 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2023-11-16 23:24:49 |
| [minio](https://min.io/download) | 41,655 | 4,998 | 53 | High Performance Object Storage for AI | 2015-01-14 19:23:58 | 2023-11-17 00:54:17 |
| [consul](https://www.consul.io) | 27,282 | 4,425 | 1,263 | Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. | 2013-11-04 22:15:27 | 2023-11-17 00:55:23 |
| [nsq](https://nsq.io) | 23,917 | 2,899 | 65 | A realtime distributed messaging platform | 2012-05-12 14:37:08 | 2023-11-17 00:12:46 |
| [roadrunner](https://roadrunner.dev) | 7,455 | 399 | 48 | 🤯 High-performance PHP application server, process manager written in Go and powered with plugins | 2017-12-26 16:13:10 | 2023-11-16 22:52:03 |
| [sftpgo](https://github.com/drakkan/sftpgo) | 7,072 | 604 | 19 | Fully featured and highly configurable SFTP server with optional HTTP/S, FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-07-20 10:18:31 | 2023-11-16 19:16:49 |
| [devd](https://github.com/cortesi/devd) | 3,363 | 150 | 25 |  A local webserver for developers | 2015-09-27 22:43:00 | 2023-11-13 01:13:20 |
| [flipt](https://flipt.io) | 2,943 | 150 | 29 | An open source, self-hosted feature flag solution | 2016-11-05 00:09:07 | 2023-11-17 00:24:24 |
| [fider](https://fider.io) | 2,453 | 583 | 48 | Open platform to collect and prioritize feedback | 2017-01-17 22:55:19 | 2023-11-15 02:31:35 |
| [algernon](https://algernon.roboticoverlords.org) | 2,341 | 122 | 21 | Small self-contained pure-Go web server with Lua, Teal, Markdown, HTTP/2, QUIC, Redis and PostgreSQL support | 2015-03-10 11:25:30 | 2023-11-15 07:46:24 |
| [flagr](https://openflagr.github.io/flagr) | 2,282 | 186 | 5 | Flagr is a feature flagging, A/B testing and dynamic configuration microservice | 2017-10-03 19:07:32 | 2023-11-15 02:46:47 |
| [trickster](https://trickstercache.org) | 1,921 | 172 | 60 | Open Source HTTP Reverse Proxy Cache and Time Series Dashboard Accelerator | 2018-03-29 20:31:44 | 2023-11-18 10:42:59 |
| [discovery](https://github.com/bilibili/discovery) | 1,764 | 401 | 26 | A registry for resilient mid-tier load balancing and failover. | 2018-04-20 12:57:50 | 2023-11-15 05:47:58 |
| [jackal](https://github.com/ortuman/jackal) | 1,439 | 133 | 22 | 💬 Instant messaging server for the Extensible Messaging and Presence Protocol (XMPP). | 2017-11-13 18:17:48 | 2023-11-03 11:58:13 |
| [go-feature-flag](https://gofeatureflag.org) | 868 | 87 | 8 | GO Feature Flag is a simple, complete and lightweight self-hosted feature flag solution 100% Open Source. 🎛️ | 2020-12-11 13:19:17 | 2023-11-16 19:42:55 |
| [dudeldu](https://github.com/krotik/dudeldu) | 142 | 17 | 0 | A simple SHOUTcast server. | 2016-09-07 19:11:04 | 2023-10-07 13:51:09 |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 102 | 14 | 29 | Simple Reverse Proxy with Caching, written in Go, using Redis. | 2020-11-12 15:10:40 | 2023-10-12 20:54:31 |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 82 | 17 | 40 | Reverse proxy with automatically obtains TLS certificates from Let's Encrypt | 2019-04-12 05:39:43 | 2023-09-18 15:18:41 |
| [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) | 80 | 51 | 6 | Prometheus remote write proxy that adds Cortex/Mimir tenant ID based on metric labels | 2020-10-06 16:52:25 | 2023-11-15 14:15:19 |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 53 | 13 | 2 | Stream database events from PostgreSQL to Kafka | 2019-04-28 21:55:31 | 2023-09-18 15:18:50 |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 38 | 5 | 0 | Turn Nginx logs into Prometheus metrics | 2018-10-23 09:10:27 | 2023-09-18 15:18:48 |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 36 | 5 | 5 | Fully featured and highly configurable SFTP server with optional HTTP/S, FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-12-18 12:48:14 | 2023-11-14 09:46:29 |
| [protoxy](https://github.com/camgraff/protoxy) | 32 | 5 | 0 | A proxy server than converts JSON request bodies to protocol buffers | 2020-09-03 23:24:34 | 2023-09-28 09:34:44 |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 2 | 2 | 0 | Service for relaying Riemann events to Riemann/Carbon destinations | 2019-04-23 14:17:12 | 2022-09-27 09:23:35 |
</details>

### Stream Processing
Libraries and tools for stream processing and reactive programming.

<sup>*Last Update: 2023-12-05 20:35:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-streams](https://pkg.go.dev/github.com/reugn/go-streams) | 1,601 | 139 | 24 | A lightweight stream processing library for Go | 2019-04-30 17:28:15 | 2023-12-03 12:32:08 |
| [machine](https://pkg.go.dev/github.com/whitaker-io/machine) | 139 | 13 | 1 | Machine is a workflow/pipeline library for processing data | 2020-10-13 04:24:19 | 2023-11-19 00:17:09 |
| [stream](https://github.com/youthlin/stream) | 84 | 11 | 1 | Go Stream, like Java 8 Stream. | 2020-11-12 03:52:50 | 2023-11-02 15:01:37 |
</details>

### Template Engines
Libraries and tools for templating and lexing.

<sup>*Last Update: 2023-11-15 08:23:48*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gofpdf](http://godoc.org/github.com/jung-kurt/gofpdf) | 4,203 | 781 | 56 | A PDF document generator with high level support for text, drawing and images | 2015-03-13 11:57:30 | 2023-11-12 11:44:08 |
| [sprig](http://masterminds.github.io/sprig/) | 3,812 | 438 | 129 | Useful template functions for Go templates. | 2013-11-22 01:20:40 | 2023-11-14 21:15:55 |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 2,902 | 149 | 37 | Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template | 2016-03-06 21:42:01 | 2023-11-13 01:40:31 |
| [pongo2](https://www.schlachter.tech/pongo2) | 2,675 | 298 | 61 | Django-syntax like template-engine for Go | 2013-08-23 01:00:08 | 2023-11-13 20:28:19 |
| [hero](https://shiyanhui.github.io/hero) | 1,547 | 107 | 28 | A handy, fast and powerful go template engine. | 2017-01-15 13:31:50 | 2023-11-07 07:29:57 |
| [maroto](https://maroto.io) | 1,276 | 161 | 18 | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. | 2019-05-20 23:27:47 | 2023-11-14 19:50:42 |
| [mustache](https://github.com/hoisie/mustache) | 1,092 | 246 | 33 | The mustache template language in Go | 2009-12-30 21:05:05 | 2023-11-10 10:01:57 |
| [jet](https://shiyanhui.github.io/hero) | 1,090 | 103 | 22 | Jet  template engine | 2016-03-31 16:53:36 | 2023-11-12 14:53:17 |
| [amber](https://github.com/eknkc/amber) | 909 | 63 | 24 | Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade | 2012-10-31 20:27:24 | 2023-11-10 02:13:06 |
| [gorazor](https://github.com/sipin/gorazor) | 835 | 89 | 2 | Razor view engine for go | 2014-05-01 05:30:31 | 2023-11-11 23:51:56 |
| [ace](https://github.com/yosssi/ace) | 826 | 87 | 30 | HTML template engine for Go | 2014-07-13 13:39:19 | 2023-11-07 00:20:18 |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 776 | 79 | 12 | Simple and fast template engine for Go | 2015-08-19 12:44:22 | 2023-11-10 08:50:28 |
| [ego](https://github.com/benbjohnson/ego) | 565 | 42 | 10 | An ERB-style templating language for Go. | 2014-02-23 18:14:41 | 2023-11-08 06:19:21 |
| [raymond](https://github.com/aymerick/raymond) | 555 | 95 | 19 | Handlebars for golang | 2015-04-22 13:07:59 | 2023-11-14 00:16:38 |
| [goview](https://github.com/foolin/goview) | 378 | 34 | 16 | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. | 2019-04-14 11:22:41 | 2023-11-10 16:41:19 |
| [liquid](https://godoc.org/github.com/osteele/liquid) | 241 | 47 | 26 | A Liquid template engine in Go | 2017-06-26 14:39:52 | 2023-11-13 19:33:49 |
| [soy](https://github.com/robfig/soy) | 171 | 42 | 7 | Go implementation for Soy templates (Google Closure templates) | 2013-12-15 01:14:48 | 2023-10-07 12:08:06 |
| [kasia.go](https://github.com/ziutek/kasia.go) | 74 | 9 | 2 | Templating system for HTML and other text documents - go implementation | 2010-12-07 10:46:05 | 2022-03-15 21:35:36 |
| [velvet](http://masterminds.github.io/sprig/) | 72 | 10 | 2 | A sweet velvety templating package | 2016-12-29 16:46:57 | 2023-03-01 11:33:12 |
| [extemplate](https://git.sr.ht/~dvko/extemplate) | 55 | 17 | 1 | Wrapper package for Go's template/html to allow for easy file-based template inheritance. | 2018-08-10 20:34:19 | 2023-11-14 09:22:20 |
| [gospin](https://github.com/m1/gospin) | 48 | 7 | 3 | Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations | 2019-02-22 17:04:51 | 2023-10-11 18:13:14 |
| [damsel](https://github.com/dskinner/damsel) | 23 | 6 | 1 | Package damsel provides html outlining via css-selectors and common template functionality. | 2012-05-02 23:06:48 | 2023-03-17 14:34:28 |
</details>

### Testing - Fail injection


<sup>*Last Update: 2023-11-20 08:20:46*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [failpoint](https://github.com/pingcap/failpoint) | 791 | 63 | 6 | An implementation of failpoints for Golang. | 2019-04-02 07:48:18 | 2023-11-15 01:45:10 |
</details>

### Testing - Fuzzing and delta-debugging, reducing, shrinking


<sup>*Last Update: 2023-12-05 20:35:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 4,667 | 318 | 57 | Randomized testing for Go | 2015-04-15 13:07:50 | 2023-12-02 15:22:13 |
| [gofuzz](https://github.com/google/gofuzz) | 1,467 | 162 | 12 | Fuzz testing for go. | 2014-07-31 16:21:29 | 2023-11-30 02:03:19 |
| [tavor](https://github.com/zimmski/tavor) | 241 | 10 | 53 | A generic fuzzing and delta-debugging framework | 2014-05-18 14:59:14 | 2023-09-16 01:42:01 |
</details>

### Testing - Mock


<sup>*Last Update: 2023-11-19 20:35:42*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mock](https://pkg.go.dev/github.com/h2non/gock) | 9,125 | 658 | 84 | GoMock is a mocking framework for the Go programming language. | 2015-06-12 17:15:11 | 2023-11-18 19:41:54 |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 5,563 | 433 | 74 | Sql mock driver for golang to test database interactions | 2014-02-07 07:59:29 | 2023-11-19 08:16:00 |
| [hoverfly](https://hoverfly.io) | 2,275 | 204 | 34 | Lightweight service virtualization/ API simulation / API mocking tool for developers and testers | 2015-11-30 16:36:31 | 2023-11-19 08:36:12 |
| [gock](https://pkg.go.dev/github.com/h2non/gock) | 1,963 | 106 | 40 | HTTP traffic mocking and testing made easy in Go ༼ʘ̚ل͜ʘ̚༽ | 2016-03-02 16:20:26 | 2023-11-17 09:41:12 |
| [httpmock](http://godoc.org/github.com/jarcoal/httpmock) | 1,788 | 102 | 3 | HTTP mocking for Golang | 2014-02-24 16:47:59 | 2023-11-14 08:40:54 |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 883 | 87 | 28 | A tool for generating self-contained, type-safe test doubles in go | 2014-05-21 00:12:54 | 2023-11-16 19:04:46 |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 581 | 43 | 5 | Immutable transaction isolated sql driver for golang | 2015-07-08 07:34:53 | 2023-11-13 16:51:34 |
| [minimock](https://github.com/gojuno/minimock) | 520 | 36 | 18 | Powerful mock generation tool for Go programming language | 2016-08-03 16:01:35 | 2023-11-18 19:34:45 |
| [govcr](https://github.com/seborama/govcr) | 160 | 14 | 1 | HTTP mock for Golang: record and replay HTTP/HTTPS interactions for offline testing | 2016-07-10 17:47:41 | 2023-10-23 06:51:11 |
| [go-localstack](https://github.com/elgohr/go-localstack) | 72 | 19 | 1 | Go Wrapper for using localstack | 2020-03-18 07:13:02 | 2023-10-05 11:30:38 |
| [timex](https://github.com/cabify/timex) | 70 | 5 | 1 | A test-friendly replacement for golang's time package | 2020-01-02 18:06:48 | 2023-09-13 09:00:01 |
| [mockhttp](https://github.com/tv42/mockhttp) | 23 | 6 | 0 | Mock object for Go http.ResponseWriter | 2011-06-11 16:03:01 | 2023-11-14 18:39:38 |
| [mockit](https://github.com/pasdam/mockit) | 16 | 3 | 4 | Library that make mocking of Go functions/methods easy | 2020-01-01 08:46:09 | 2023-05-30 03:26:52 |
</details>

### Testing - Selenium and browser control tools


<sup>*Last Update: 2023-12-14 20:06:43*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [chromedp](https://github.com/chromedp/chromedp) | 9,834 | 758 | 56 | A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. | 2017-01-24 14:54:30 | 2023-12-14 09:18:27 |
| [rod](https://go-rod.github.io) | 4,400 | 299 | 107 | A Devtools driver for web automation and scraping | 2020-01-21 20:09:45 | 2023-12-14 06:18:42 |
| [selenoid](https://aerokube.com/selenoid/latest/) | 2,431 | 320 | 266 | Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary. | 2016-08-22 09:11:16 | 2023-12-10 16:40:17 |
| [playwright-go](https://playwright-community.github.io/playwright-go/) | 1,555 | 131 | 44 | Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API. | 2020-08-16 12:46:14 | 2023-12-13 22:54:14 |
| [cdp](https://github.com/mafredri/cdp) | 700 | 47 | 13 | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 2017-03-12 10:25:41 | 2023-12-12 08:22:57 |
| [ggr](https://aerokube.com/ggr/latest/) | 303 | 69 | 13 | A lightweight load balancer used to create big Selenium clusters | 2016-06-16 15:33:24 | 2023-12-11 10:36:29 |
</details>

### Testing - Testing Frameworks


<sup>*Last Update: 2023-11-12 20:11:26*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [testify](https://github.com/stretchr/testify) | 20,907 | 1,522 | 410 | A toolkit with common assertions and mocks that plays nicely with the standard library | 2012-10-16 16:43:17 | 2023-11-12 11:01:37 |
| [goconvey](http://smartystreets.github.io/goconvey/) | 7,944 | 595 | 158 | Go testing in the browser. Integrates with `go test`. Write behavioral tests in Go. | 2013-08-21 04:52:28 | 2023-11-09 07:57:27 |
| [ginkgo](http://onsi.github.io/ginkgo/) | 7,578 | 668 | 63 | A Modern Testing Framework for Go | 2013-08-23 20:53:05 | 2023-11-12 10:31:07 |
| [go-cmp](https://github.com/google/go-cmp) | 3,768 | 256 | 31 | Package for comparing Go values in tests | 2017-07-07 19:28:22 | 2023-11-10 11:22:22 |
| [httpexpect](https://pkg.go.dev/github.com/gavv/httpexpect/v2) | 2,346 | 250 | 12 | End-to-end HTTP and REST API testing for Go. | 2016-04-29 17:05:20 | 2023-11-11 10:02:14 |
| [godog](https://github.com/cucumber/godog) | 2,074 | 263 | 56 | Cucumber for golang | 2015-06-10 13:16:31 | 2023-11-09 11:10:41 |
| [gomega](http://onsi.github.io/gomega/) | 1,983 | 331 | 45 | Ginkgo's Preferred Matcher Library | 2013-08-23 20:54:42 | 2023-11-11 18:11:27 |
| [gnomock](https://github.com/orlangure/gnomock) | 1,254 | 63 | 25 | Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻 | 2020-01-31 14:50:52 | 2023-11-10 03:38:25 |
| [go-vcr](https://go-testdeep.zetta.rocks/) | 1,053 | 70 | 1 | Record and replay your HTTP interactions for fast, deterministic and accurate tests | 2015-12-14 12:52:17 | 2023-11-11 10:11:03 |
| [testfixtures](https://pkg.go.dev/github.com/go-testfixtures/testfixtures/v3?tab=doc) | 992 | 77 | 18 | Ruby on Rails like test fixtures for Go. Write tests against a real database | 2016-04-05 11:33:28 | 2023-11-10 21:42:41 |
| [goblin](https://github.com/franela/goblin) | 887 | 79 | 22 | Minimal and Beautiful Go testing framework | 2013-09-19 02:34:24 | 2023-10-31 17:42:52 |
| [baloo](https://godoc.org/github.com/h2non/baloo) | 766 | 32 | 8 | Expressive end-to-end HTTP API testing made easy in Go | 2016-05-29 21:40:58 | 2023-10-23 11:15:03 |
| [goc](https://github.com/qiniu/goc) | 728 | 105 | 58 | A Comprehensive Coverage Testing System for The Go Programming Language | 2020-05-07 03:46:25 | 2023-11-07 16:31:20 |
| [apitest](https://apitest.dev) | 692 | 56 | 3 | A simple and extensible behavioural testing library for Go. You can use api test to simplify REST API, HTTP handler and e2e tests. | 2018-12-26 22:27:19 | 2023-11-10 10:46:26 |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 656 | 79 | 2 | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test | 2019-11-16 23:49:40 | 2023-11-09 18:56:15 |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 607 | 56 | 43 | Mutation testing for Go source code | 2014-12-26 22:23:44 | 2023-11-08 09:27:40 |
| [gotest.tools](https://gotest.tools/v3) | 474 | 49 | 28 | A collection of packages to augment the go testing package and support common patterns. | 2017-08-08 21:28:54 | 2023-11-12 03:44:23 |
| [gofight](https://github.com/appleboy/gofight) | 437 | 42 | 6 | Testing API Handler written in Golang. | 2016-03-29 00:13:21 | 2023-10-03 12:36:30 |
| [go-testdeep](https://go-testdeep.zetta.rocks/) | 403 | 17 | 3 | Extremely flexible golang deep comparison, extends the go testing package, tests HTTP APIs and provides tests suite | 2018-05-26 15:03:28 | 2023-10-30 14:42:36 |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 278 | 28 | 13 | Simple Go snapshot testing | 2017-08-07 18:30:05 | 2023-11-12 10:47:42 |
| [frisby](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 276 | 27 | 12 | API testing framework inspired by frisby-js | 2015-09-15 14:35:58 | 2023-10-03 12:36:36 |
| [endly](https://github.com/viant/endly) | 246 | 32 | 4 | End to end functional test and automation framework | 2017-08-28 20:24:43 | 2023-09-27 02:09:54 |
| [go-carpet](https://github.com/msoap/go-carpet) | 240 | 10 | 3 | Tool for show test coverage in terminal for Go source files | 2016-02-28 12:02:51 | 2023-10-17 10:04:07 |
| [go-hit](https://github.com/Eun/go-hit) | 240 | 9 | 13 | http integration test framework | 2019-06-04 16:28:23 | 2023-10-29 13:43:21 |
| [commander](https://github.com/commander-cli/commander) | 216 | 17 | 35 | Test your command line interfaces on windows, linux and osx and nodes viá ssh and docker | 2019-02-22 16:35:16 | 2023-11-11 21:20:52 |
| [charlatan](https://github.com/percolate/charlatan) | 199 | 10 | 2 | Go Interface Mocking Tool | 2017-10-06 21:55:14 | 2023-09-27 02:08:30 |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 157 | 15 | 0 | Clean database for testing, inspired by database_cleaner for Ruby | 2017-01-17 18:18:40 | 2023-09-27 02:09:10 |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 115 | 17 | 5 | A Go test assertion library for verifying that two representations of JSON are semantically equal | 2018-10-26 20:31:01 | 2023-11-02 10:33:17 |
| [gospec](https://github.com/luontola/gospec) | 114 | 17 | 3 | Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] | 2009-11-24 13:59:26 | 2023-09-25 21:27:38 |
| [testcase](https://github.com/adamluzsi/testcase) | 114 | 10 | 0 | testcase is an opinionated testing framework to support test driven design. | 2019-04-22 21:20:51 | 2023-09-20 12:56:31 |
| [wstest](https://github.com/posener/wstest) | 101 | 17 | 1 | go websocket client for unit testing of a websocket handler | 2017-03-31 21:06:18 | 2023-10-10 13:04:46 |
| [gocrest](https://github.com/corbym/gocrest) | 96 | 5 | 1 | GoCrest - Hamcrest-like matchers for Go | 2017-12-23 23:27:00 | 2023-09-17 22:48:21 |
| [assert](https://github.com/go-playground/assert) | 61 | 17 | 2 | :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions | 2015-07-20 17:53:45 | 2023-10-03 00:38:04 |
| [covergates](https://covergates.com) | 59 | 12 | 24 | The portal gates to coverage reports | 2020-05-29 04:02:01 | 2023-09-27 02:08:52 |
| [restit](https://github.com/go-restit/restit) | 55 | 6 | 4 | A Go library help testing your RESTful API application | 2014-06-25 10:25:46 | 2023-10-04 22:55:42 |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 7 | 1 | A BDD library for Go | 2009-11-20 06:34:29 | 2022-09-27 09:28:31 |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 45 | 4 | 0 | Library created for testing JSON against patterns. | 2019-01-27 20:19:06 | 2023-08-18 07:04:44 |
| [dsunit](https://github.com/viant/dsunit) | 43 | 9 | 0 | Datastore Testibility | 2016-06-13 20:20:52 | 2023-07-14 04:01:36 |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 28 | 5 | 2 | Hamcrest matchers for the Go programming language | 2010-12-22 04:49:44 | 2022-09-27 09:28:41 |
| [schema](https://github.com/jgroeneveld/schema) | 20 | 1 | 0 | Quick and easy expression matching for JSON schemas used in requests and responses | 2015-08-13 13:36:54 | 2023-03-10 16:27:27 |
| [flute](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 18 | 1 | 6 | Golang HTTP client testing framework | 2019-07-06 04:32:03 | 2023-09-27 02:10:11 |
| [testsql](https://github.com/zhulongcheng/testsql) | 17 | 2 | 3 | Generate test data from SQL files before testing and clear it after finished. | 2018-09-22 12:03:50 | 2023-11-01 16:13:06 |
| [gogiven](https://github.com/corbym/gogiven) | 14 | 3 | 4 | gogiven - BDD testing framework for go that generates readable output directly from source code | 2017-12-31 22:33:37 | 2023-09-08 17:34:35 |
| [biff](https://github.com/fulldump/biff) | 13 | 2 | 0 | Bifurcation Framework for testing and use cases | 2018-03-28 18:35:53 | 2023-09-27 02:08:23 |
| [gosuite](https://github.com/pavlo/gosuite) | 12 | 4 | 1 | Test suites support for standard Go1.7 "testing" by leveraging Subtests feature | 2016-10-15 19:28:14 | 2022-09-27 09:28:34 |
| [badio](https://github.com/cavaliergopher/badio) | 11 | 2 | 0 | Extensions to Go's testing/iotest package | 2016-02-11 10:29:25 | 2023-09-27 02:07:55 |
| [test](https://github.com/tvastar/test) | 9 | 1 | 0 | test utilities for golang | 2019-03-23 21:47:36 | 2023-10-30 05:53:01 |
| [stop-and-go](https://github.com/elgohr/stop-and-go) | 9 | 4 | 0 | Testing helper for concurrency | 2020-11-06 09:04:58 | 2023-09-30 20:22:42 |
| [trial](https://github.com/jgroeneveld/trial) | 6 | 1 | 0 | A simple assertion library for go | 2015-06-18 09:01:30 | 2022-10-05 19:52:45 |
| [tt](https://github.com/vcaesar/tt) | 6 | 1 | 0 | Simple and colorful test tools | 2018-04-03 11:47:21 | 2023-02-08 12:16:17 |
</details>

### Text Processing - Specific Formats


<sup>*Last Update: 2023-12-18 21:42:42*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [colly](https://go-colly.org/) | 21,332 | 1,681 | 176 | Elegant Scraper and Crawler Framework for Golang | 2017-09-29 14:08:49 | 2023-12-18 09:00:21 |
| [goquery](https://godoc.org/astuart.co/goq) | 13,205 | 966 | 4 | A little like that j-thing, only in Go. | 2012-08-29 02:14:59 | 2023-12-18 06:50:05 |
| [sh](https://pkg.go.dev/mvdan.cc/sh/v3) | 6,375 | 328 | 97 | A shell parser, formatter, and interpreter with bash support; includes shfmt | 2016-01-16 08:39:09 | 2023-12-18 01:20:26 |
| [blackfriday](https://github.com/russross/blackfriday) | 5,271 | 638 | 217 | Blackfriday: a markdown processor for Go | 2011-05-27 22:28:58 | 2023-12-18 14:25:39 |
| [toml](https://github.com/BurntSushi/toml) | 4,342 | 530 | 9 | TOML parser for Golang with reflection. | 2013-02-26 05:05:48 | 2023-12-17 21:09:08 |
| [go-humanize](https://pkg.go.dev/github.com/dustin/go-humanize) | 3,876 | 272 | 41 | Go Humans! (formatters for units to human friendly sizes) | 2012-01-13 03:48:55 | 2023-12-17 19:20:34 |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 2,863 | 202 | 24 | bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS | 2013-11-20 22:15:49 | 2023-12-18 14:26:26 |
| [gofeed](https://github.com/mmcdole/gofeed) | 2,311 | 238 | 43 | Parse RSS, Atom and JSON feeds in Go | 2016-01-23 02:44:34 | 2023-12-18 14:00:41 |
| [go-toml](https://github.com/pelletier/go-toml) | 1,528 | 246 | 24 | Go library for the TOML file format | 2013-02-24 17:45:51 | 2023-12-13 07:58:26 |
| [inject](https://godoc.org/github.com/facebookgo/inject) | 1,401 | 139 | 9 | Package inject provides a reflect based injector. | 2013-10-21 01:51:46 | 2023-12-18 03:37:53 |
| [slug](https://pkg.go.dev/mvdan.cc/sh/v3) | 1,037 | 101 | 12 | URL-friendly slugify with multiple languages support. | 2014-03-31 06:24:51 | 2023-12-17 17:20:33 |
| [commonregex](https://github.com/mingrammer/commonregex) | 870 | 66 | 2 | 🍫 A collection of common regular expressions for Go | 2017-03-23 14:33:18 | 2023-12-18 08:08:02 |
| [htmlquery](https://github.com/antchfx/xpath) | 655 | 103 | 11 | htmlquery is golang XPath package for HTML query. | 2017-12-05 01:08:41 | 2023-12-07 15:00:56 |
| [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) | 638 | 109 | 17 | ⚙️ Convert HTML to Markdown. Even works with entire websites and can be extended through rules. | 2018-05-15 13:26:26 | 2023-12-18 09:51:54 |
| [dataflowkit](https://dataflowkit.com) | 623 | 80 | 4 | Extract structured data from web sites. Web sites scraping.   | 2017-02-09 15:08:15 | 2023-12-06 14:36:34 |
| [omniparser](https://github.com/jf-tech/omniparser) | 601 | 48 | 1 | omniparser: a native Golang ETL streaming parser and transform library for CSV, JSON, XML, EDI, text, etc. | 2020-08-16 22:22:21 | 2023-12-09 23:34:32 |
| [mxj](https://github.com/clbanning/mxj) | 589 | 110 | 3 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2014-02-03 13:39:16 | 2023-12-18 02:11:21 |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 562 | 89 | 14 | wcwidth for golang | 2013-06-21 04:56:50 | 2023-12-16 23:35:01 |
| [gographviz](https://github.com/awalterschulze/gographviz) | 533 | 76 | 8 | Parses the Graphviz DOT language in golang | 2015-03-14 18:27:00 | 2023-12-11 07:25:44 |
| [gommon](https://github.com/labstack/gommon) | 511 | 131 | 13 | Common packages for Go | 2015-03-12 22:35:57 | 2023-11-23 07:02:16 |
| [gotext](https://github.com/zhshch2002/gospider) | 407 | 55 | 7 | Go (Golang) GNU gettext utilities package  | 2016-06-19 20:14:43 | 2023-12-18 03:00:29 |
| [goq](https://godoc.org/astuart.co/goq) | 247 | 18 | 1 | A declarative struct-tag-based HTML unmarshaling or scraping package for Go built on top of the goquery library | 2017-02-20 02:54:40 | 2023-11-27 01:03:11 |
| [goribot](https://github.com/zhshch2002/gospider) | 210 | 30 | 1 | [Crawler/Scraper for Golang]🕷A lightweight distributed friendly Golang crawler framework.一个轻量的分布式友好的 Golang 爬虫框架。 | 2019-09-08 10:39:47 | 2023-10-29 10:35:53 |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 202 | 71 | 2 | A NMEA parser library in pure Go | 2015-07-22 08:55:54 | 2023-11-17 20:12:10 |
| [github_flavored_markdown](https://github.com/shurcooL/github_flavored_markdown) | 155 | 42 | 11 | GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links. | 2015-05-16 04:09:07 | 2023-11-17 06:25:17 |
| [podcast](https://github.com/eduncan911/podcast) | 125 | 32 | 5 | iTunes and RSS 2.0 Podcast Generator in Golang | 2017-02-02 12:45:04 | 2023-08-30 03:07:41 |
| [editorconfig-core-go](https://godoc.org/github.com/hscells/doi) | 125 | 37 | 5 | EditorConfig Core written in Go | 2016-07-05 03:50:41 | 2023-10-18 14:18:01 |
| [sdp](https://github.com/gortc/sdp) | 112 | 35 | 5 | RFC 4566 SDP implementation in go | 2016-05-13 14:35:11 | 2023-11-27 01:02:56 |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 111 | 9 | 0 | Zero-width character detection and removal for Go | 2018-06-18 13:55:09 | 2023-11-09 22:25:15 |
| [go-vcard](https://github.com/pelletier/go-toml) | 94 | 33 | 4 | A Go library to parse and format vCard | 2017-03-21 08:30:36 | 2023-12-08 18:07:00 |
| [pagser](https://github.com/foolin/pagser) | 93 | 7 | 6 | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler | 2020-04-19 09:22:00 | 2023-12-09 20:35:41 |
| [go-slugify](https://godoc.org/github.com/mozillazg/go-slugify) | 88 | 8 | 1 | Pretty Slug. | 2016-07-16 11:55:15 | 2023-09-25 01:16:07 |
| [goregen](https://godoc.org/github.com/zach-klippenstein/goregen) | 88 | 20 | 6 | randexp for Go. | 2014-12-27 00:19:39 | 2023-10-11 11:37:23 |
| [align](https://github.com/Guitarbum722/align) | 83 | 7 | 0 | A general purpose application and library for aligning text. | 2017-04-29 23:22:22 | 2023-11-12 23:58:46 |
| [did](https://w3c-ccg.github.io/did-spec) | 79 | 21 | 7 | A golang package to work with Decentralized Identifiers (DIDs)  | 2018-11-02 17:49:14 | 2023-12-14 07:36:41 |
| [go-fixedwidth](http://godoc.org/github.com/ianlopshire/go-fixedwidth) | 78 | 30 | 6 | Encoding and decoding for fixed-width formatted data | 2017-11-15 21:05:44 | 2023-10-03 05:04:52 |
| [genex](https://namegrep.com/) | 75 | 8 | 0 | Genex package for Go | 2015-03-09 19:24:16 | 2023-09-13 11:58:03 |
| [allot](https://github.com/sbstjn/allot) | 58 | 8 | 3 | Parse placeholder and wildcard text commands | 2016-10-16 15:49:08 | 2023-08-05 11:57:10 |
| [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) | 58 | 9 | 7 | 🚀 Fast and light wildcard pattern matching. | 2021-03-28 16:31:41 | 2023-12-16 07:35:06 |
| [guesslanguage](https://github.com/zhshch2002/gospider) | 57 | 5 | 1 | Guess the natural language of a text in Go | 2014-12-16 10:58:47 | 2023-06-22 20:18:49 |
| [gonameparts](https://github.com/polera/gonameparts) | 38 | 5 | 2 | Takes a full name and splits it into individual name parts | 2015-05-17 05:20:17 | 2022-09-27 09:35:04 |
| [normalize](https://github.com/avito-tech/normalize) | 37 | 2 | 0 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2021-03-22 09:25:14 | 2023-11-17 03:50:31 |
| [slugify](https://pkg.go.dev/mvdan.cc/sh/v3) | 32 | 4 | 0 | A Go slugify application that handles string | 2015-04-13 01:54:30 | 2022-09-27 08:50:57 |
| [codetree](https://github.com/aerogo/codetree) | 23 | 5 | 0 | :evergreen_tree: Parses indented code and returns a tree structure. | 2016-11-26 02:50:38 | 2023-09-04 20:53:44 |
| [enca](https://godoc.org/github.com/hscells/doi) | 16 | 5 | 0 | Minimal cgo bindings for libenca | 2014-12-17 04:55:16 | 2023-11-01 21:11:39 |
| [gout](https://github.com/drewstinnett/gout) | 13 | 2 | 0 | Output go objects in standard formats, such as YAML, JSON, etc | 2021-04-08 20:48:17 | 2023-10-20 10:05:59 |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 11 | 4 | 0 | A syndication feed parser for Atom 1.0 and RSS 2.0 in Go | 2017-04-07 09:30:55 | 2023-04-10 10:05:42 |
| [ltsv](https://github.com/Wing924/ltsv) | 8 | 1 | 0 | High performance LTSV (Labeled Tab Separeted Value) reader for Go. | 2019-05-12 06:11:04 | 2022-09-27 09:35:07 |
| [encoding](https://github.com/ake-persson/encoding) | 8 | 2 | 1 | Go package provides a generic interface to encoders and decoders | 2018-04-06 20:48:00 | 2022-09-27 09:34:51 |
| [doi](https://godoc.org/github.com/hscells/doi) | 8 | 2 | 0 | Parse and check doi objects in go. | 2017-08-02 05:58:01 | 2023-06-28 15:19:43 |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 7 | 3 | 0 | Converter from BBCode to HTML | 2016-04-15 14:35:38 | 2023-10-31 04:17:47 |
</details>

### Text Processing - Utility


<sup>*Last Update: 2023-11-19 20:35:30*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xurls](https://tysug.net) | 1,087 | 115 | 2 | Extract urls from text | 2015-01-12 01:28:46 | 2023-11-16 15:32:57 |
| [gotabulate](https://github.com/bndr/gotabulate) | 302 | 29 | 5 | Gotabulate - Easily pretty-print your tabular data with Go | 2014-08-21 07:44:28 | 2023-08-25 02:59:22 |
| [radix](https://github.com/yourbasic/radix) | 186 | 11 | 0 | A fast string sorting algorithm (MSD radix sort) | 2017-06-09 14:38:58 | 2023-10-27 13:45:49 |
| [regroup](https://github.com/oriser/regroup) | 136 | 12 | 2 | Match regex group into go struct using struct tags and automatic parsing | 2020-09-08 19:04:42 | 2023-10-19 20:20:55 |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 59 | 8 | 4 | A sanitization-based swear filter for Go. | 2018-09-09 00:07:26 | 2023-10-04 23:50:08 |
| [parth](https://github.com/codemodus/parth) | 45 | 6 | 0 | Path parsing for segment unmarshaling and slicing. | 2015-04-06 22:53:59 | 2023-11-17 09:49:04 |
| [xj2go](https://tysug.net) | 34 | 8 | 0 | Convert xml and json to go struct | 2017-09-19 13:20:57 | 2023-11-03 03:33:08 |
| [tagify](https://www.zoomio.org/tagify) | 34 | 2 | 3 | Tagify produces a set of tags from a given source. Source can be either an HTML page, a Markdown document or a plain text. Supports English, Russian, Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean languages. | 2018-03-20 10:30:11 | 2023-11-07 20:36:38 |
| [kace](https://github.com/codemodus/kace) | 20 | 3 | 1 | Common case conversions covering common initialisms. | 2015-06-04 20:36:49 | 2023-08-01 23:38:06 |
| [TySug](https://tysug.net) | 17 | 3 | 2 | A project around helping to prevent typing typos. TySug (Typo Suggestions) suggests alternative words with respect to keyboard layouts | 2018-06-05 19:46:29 | 2023-08-03 19:29:18 |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 10 | 6 | 1 | A string argument parser that understands quotes and backslashes | 2016-02-24 00:53:38 | 2023-10-21 16:40:53 |
| [textwrap](https://www.zoomio.org/tagify) | 5 | 4 | 1 | Port of Python's "textwrap" module to Go | 2019-07-26 17:57:55 | 2023-08-21 08:56:49 |
</details>

### Third-party APIs
Libraries for accessing third party APIs.

<sup>*Last Update: 2023-11-15 08:23:00*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-github](https://pkg.go.dev/github.com/google/go-github/v56/github) | 9,730 | 2,073 | 58 | Go library for accessing the GitHub v3 API | 2013-05-24 16:42:58 | 2023-11-13 20:59:08 |
| [aws-sdk-go](http://aws.amazon.com/sdk-for-go/) | 8,445 | 2,148 | 45 | AWS SDK for the Go programming language. | 2014-12-05 05:29:41 | 2023-11-13 15:32:48 |
| [slack](https://pkg.go.dev/github.com/slack-go/slack) | 4,433 | 1,139 | 110 | Slack API in Go - community-maintained fork created by the original author, @nlopes | 2015-01-24 14:19:00 | 2023-11-14 07:32:56 |
| [discordgo](https://github.com/bwmarrin/discordgo) | 4,346 | 910 | 134 |  (Golang) Go bindings for Discord | 2015-11-01 20:51:01 | 2023-11-13 16:31:47 |
| [google-api-go-client](https://pkg.go.dev/google.golang.org/api) | 3,622 | 1,180 | 30 | Auto-generated Google APIs for Go. | 2014-11-24 21:45:36 | 2023-11-14 02:07:34 |
| [google-cloud-go](https://cloud.google.com/go/docs/reference) | 3,394 | 1,263 | 270 | Google Cloud Client Libraries for Go. | 2014-05-09 11:11:58 | 2023-11-12 23:19:58 |
| [minio-go](https://docs.min.io/docs/golang-client-quickstart-guide.html) | 2,096 | 634 | 14 | MinIO Go client SDK for S3 compatible object storage | 2015-05-02 02:36:46 | 2023-11-13 03:16:49 |
| [stripe-go](https://stripe.com) | 1,894 | 493 | 9 | Go library for the Stripe API.     | 2014-06-05 23:38:14 | 2023-11-11 20:39:31 |
| [go-twitter](https://dev.twitter.com/rest/public) | 1,588 | 348 | 7 | Go Twitter REST and Streaming API v1.1 | 2015-04-11 23:26:07 | 2023-11-13 14:36:00 |
| [go-jira](https://pkg.go.dev/github.com/andygrunwald/go-jira?tab=doc) | 1,377 | 485 | 163 | Go client library for Atlassian Jira | 2015-08-20 15:02:46 | 2023-11-13 20:26:02 |
| [facebook](https://pkg.go.dev/github.com/huandu/facebook) | 1,215 | 598 | 1 | A Facebook Graph API SDK For Go. | 2012-07-28 19:05:56 | 2023-11-13 15:28:46 |
| [anaconda](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 1,139 | 253 | 73 | A Go client library for the Twitter 1.1 API | 2013-03-04 22:46:07 | 2023-11-12 15:16:10 |
| [githubv4](https://github.com/shurcooL/githubv4) | 1,046 | 92 | 41 | Package githubv4 is a client library for accessing GitHub GraphQL API v4 (https://docs.github.com/en/graphql). | 2017-05-27 05:05:31 | 2023-11-11 21:43:31 |
| [webhooks](https://github.com/go-playground/webhooks) | 877 | 229 | 28 | :fishing_pole_and_fish: Webhook receiver for GitHub, Bitbucket, GitLab, Gogs | 2015-10-25 17:38:13 | 2023-11-09 02:56:49 |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 780 | 144 | 24 | Scrape the Twitter frontend API without authentication with Golang. | 2018-11-29 15:31:50 | 2023-11-14 06:43:05 |
| [paypal](https://github.com/plutov/paypal) | 613 | 261 | 2 | Golang client for PayPal REST API | 2015-10-14 04:57:49 | 2023-11-11 03:42:35 |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 488 | 65 | 7 | Go library to access geocoding and reverse geocoding APIs | 2014-12-04 08:18:31 | 2023-11-11 17:56:50 |
| [ethrpc](https://github.com/onrik/ethrpc) | 260 | 104 | 12 | Golang client for ethereum json rpc api | 2017-01-24 09:47:00 | 2023-10-20 17:10:46 |
| [trello](https://github.com/adlio/trello) | 214 | 73 | 7 | Trello API wrapper for Go | 2016-09-24 04:36:10 | 2023-10-10 01:53:49 |
| [go-marathon](https://github.com/gambol99/go-marathon) | 200 | 128 | 27 | A GO API library for working with Marathon | 2015-02-11 13:25:26 | 2023-09-01 03:20:16 |
| [wit-go](https://github.com/wit-ai/wit-go) | 147 | 31 | 1 | Go client for wit.ai HTTP API | 2018-08-20 07:18:40 | 2023-09-23 19:01:55 |
| [medium-sdk-go](https://medium.com) | 138 | 20 | 6 | A Golang SDK for Medium's OAuth2 API | 2015-09-26 23:45:46 | 2023-07-30 04:40:22 |
| [pushover](https://github.com/gregdel/pushover) | 138 | 12 | 1 | Go wrapper for the Pushover API | 2015-02-19 15:30:05 | 2023-11-08 22:38:06 |
| [go-trending](http://godoc.org/github.com/andygrunwald/go-trending) | 137 | 19 | 1 | Go library for accessing trending repositories and developers at Github. | 2015-07-04 08:06:48 | 2023-10-29 10:36:30 |
| [gostorm](https://github.com/jsgilmore/gostorm) | 129 | 20 | 5 | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. | 2013-07-22 12:43:41 | 2023-10-27 21:11:19 |
| [gosip](https://go.spflow.com) | 122 | 28 | 7 | ⚡️ SharePoint SDK for Go | 2019-01-26 08:48:48 | 2023-11-13 06:07:46 |
| [simples3](https://github.com/rhnvrm/simples3) | 117 | 25 | 2 | Simple no frills AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK) | 2018-12-06 10:24:21 | 2023-11-09 16:00:44 |
| [hipchat](https://github.com/daneharrigan/hipchat) | 110 | 37 | 1 | A golang package to communicate with HipChat over XMPP | 2013-04-28 02:16:21 | 2023-07-09 08:35:31 |
| [hipchat](https://github.com/andybons/hipchat) | 105 | 21 | 0 | This project implements a Go client library for the Hipchat API. | 2012-10-20 18:34:06 | 2023-09-25 21:28:16 |
| [cachet](https://github.com/andygrunwald/cachet) | 90 | 13 | 1 | Go(lang) client library for Cachet (open source status page system). | 2015-10-31 12:30:07 | 2023-10-08 14:08:30 |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 86 | 21 | 1 | This is a Golang wrapper for working with TMDb API. It aims to support version 3. | 2019-01-11 22:59:33 | 2023-10-23 01:35:44 |
| [igdb](https://api.igdb.com/) | 79 | 16 | 3 | Go client for the Internet Game Database API | 2017-08-24 08:31:53 | 2023-09-25 15:17:30 |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 75 | 19 | 1 | Golang scraper to get data from Google Play Store | 2019-09-20 14:03:01 | 2023-11-08 05:23:21 |
| [gogtrends](https://github.com/groovili/gogtrends) | 74 | 28 | 3 | Unofficial Google Trends API for Go | 2018-12-27 13:50:34 | 2023-10-25 23:13:19 |
| [go-unsplash](https://unsplash.com) | 72 | 14 | 8 | Go Client for the Unsplash API  | 2017-01-19 07:04:04 | 2023-06-06 05:29:44 |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 70 | 17 | 1 | Go module to work with Postman Collections | 2019-11-16 12:13:32 | 2023-09-07 17:58:35 |
| [airtable](https://github.com/mehanizm/airtable) | 65 | 16 | 1 | Simple golang airtable API wrapper | 2020-04-12 10:05:07 | 2023-10-16 19:00:22 |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 65 | 51 | 5 | Go library for interacting with CircleCI | 2015-08-14 21:19:36 | 2023-10-25 02:03:41 |
| [ynab.go](https://godoc.org/github.com/brunomvsouza/ynab.go) | 62 | 23 | 3 | Go client for the YNAB API. Unofficial. It covers 100% of the resources made available by the YNAB API. | 2018-07-13 11:10:54 | 2023-11-03 18:28:30 |
| [mixpanel](https://github.com/dukex/mixpanel) | 60 | 32 | 6 | Golang Mixpanel Client | 2014-05-20 03:50:34 | 2023-11-07 12:37:44 |
| [go-amazon-product-advertising-api](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 57 | 15 | 3 | Go Client Library for Amazon Product Advertising API | 2016-11-15 15:37:32 | 2023-08-01 21:59:10 |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 57 | 11 | 13 | Client library for UptimeRobot v2 API | 2018-05-29 10:27:19 | 2023-09-18 14:47:47 |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 55 | 15 | 8 | DEPRECATED: please use https://github.com/Clarifai/clarifai-go-grpc | 2015-09-28 23:33:59 | 2023-08-17 19:56:10 |
| [megos](https://godoc.org/github.com/andygrunwald/megos) | 54 | 11 | 0 | Go(lang) client library for accessing information of an Apache Mesos cluster. | 2015-10-02 14:29:20 | 2023-06-19 00:39:00 |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 52 | 24 | 6 | a Go (Golang) MusicBrainz WS2 client library - work in progress | 2014-09-10 16:42:33 | 2023-11-11 17:05:55 |
| [fcm](https://pkg.go.dev/github.com/huandu/facebook) | 51 | 15 | 2 | Firebase Cloud Messaging for application servers implemented using the Go programming language. | 2017-01-06 08:30:57 | 2023-08-27 19:26:42 |
| [go-xkcd](https://pkg.go.dev/github.com/nishanths/go-xkcd/v2) | 51 | 5 | 0 | xkcd.com API client in Go | 2016-02-26 05:14:31 | 2022-12-24 04:57:10 |
| [gads](https://pkg.go.dev/github.com/huandu/facebook) | 49 | 58 | 8 | Google Adwords API for Go | 2014-01-20 02:22:15 | 2023-11-08 20:44:50 |
| [go-spotify](https://pkg.go.dev/github.com/slack-go/slack) | 48 | 8 | 1 | Go library for the Spotify Web API | 2014-10-30 02:52:04 | 2023-09-26 20:47:37 |
| [golyrics](https://github.com/mamal72/golyrics) | 40 | 2 | 0 | A simple Go package to fetch lyrics from Wikia | 2016-11-18 04:40:37 | 2022-12-24 05:00:10 |
| [patreon-go](https://github.com/mxpv/patreon-go) | 37 | 19 | 1 | Patreon Go API client | 2017-08-06 21:15:14 | 2023-09-18 14:44:21 |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 36 | 2 | 2 | Go library for accessing the MyAnimeList API: https://myanimelist.net/apiconfig/references/api/v2 | 2015-05-03 10:07:05 | 2023-03-30 10:42:12 |
| [translate](https://github.com/nuveo/translate) | 33 | 6 | 0 | Go online translation package | 2015-07-13 15:42:13 | 2023-01-28 07:01:11 |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 33 | 6 | 3 | Golang client for LastPass | 2019-07-11 14:26:39 | 2023-10-09 18:26:28 |
| [gami](https://gitlab.com/bit4bit/gami) | 32 | 30 | 1 | GO - Asterisk AMI Interface | 2014-05-14 16:11:37 | 2023-08-10 19:51:18 |
| [gcm](https://gitlab.com/bit4bit/gami) | 31 | 4 | 0 | Google Cloud Messaging for application servers implemented using the Go programming language. | 2015-11-09 16:16:25 | 2022-09-27 09:40:32 |
| [go-steam](https://pkg.go.dev/github.com/slack-go/slack) | 30 | 6 | 2 | Go library for querying Source servers | 2014-11-23 16:34:56 | 2023-05-30 02:15:11 |
| [go-imgur](https://github.com/koffeinsource/go-imgur) | 25 | 9 | 3 | Go library to use the imgur.com API | 2016-03-30 22:05:35 | 2023-05-28 07:32:55 |
| [go-shopify](https://github.com/rapito/go-shopify) | 25 | 6 | 2 | Simple Shopify API for the Go Programming Language | 2014-10-28 02:53:25 | 2022-12-07 01:15:06 |
| [device-check-go](https://github.com/rinchsan/device-check-go) | 22 | 6 | 1 | :iphone: iOS DeviceCheck SDK for Go - query and modify the per-device bits | 2019-04-11 13:09:11 | 2023-09-18 14:38:32 |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 21 | 2 | 5 | A golang client for the Twitch v3 API - public APIs only (for now) | 2016-06-28 20:54:34 | 2021-05-10 11:16:29 |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 21 | 9 | 2 | Go client library for interacting with Coinpaprika's API | 2018-09-25 07:34:50 | 2023-10-07 17:41:17 |
| [textbelt](https://www.gregd.org/) | 19 | 1 | 0 | golang library for textbelt.com | 2015-09-01 22:46:42 | 2022-11-15 15:53:59 |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 19 | 1 | 5 | Go library for http://www.brewerydb.com/ API | 2015-04-15 02:59:41 | 2022-09-27 09:39:58 |
| [go-aws-news](https://caleblemoine.dev/go-aws-news/) | 19 | 4 | 0 | Go app + library to fetch what's new from AWS | 2020-01-08 00:59:39 | 2023-09-28 09:00:35 |
| [codeship-go](https://godoc.org/github.com/codeship/codeship-go) | 18 | 9 | 2 | Go library for accessing the Codeship API v2 | 2017-09-08 16:49:59 | 2023-08-29 09:09:52 |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 16 | 1 | 0 | 📟  Tiny utility Go client for HackerNews API. | 2017-08-10 20:44:02 | 2022-09-27 09:40:47 |
| [gopaapi5](https://www.utekar.com/amazon-product-advertising-api-5-go-client-library-gopaapi/) | 15 | 7 | 0 | Go Client Library for Amazon's Product Advertising API 5.0 | 2020-02-15 06:21:31 | 2023-08-08 11:58:25 |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 14 | 3 | 0 | Simple Reporting for Google Analytics | 2015-06-01 13:50:00 | 2023-09-08 16:58:13 |
| [go-openproject](https://github.com/manuelbcd/go-openproject) | 13 | 6 | 5 | Go client library for OpenProject | 2021-02-13 23:23:13 | 2023-04-23 02:39:57 |
| [go-here](https://github.com/abdullahselek/go-here) | 12 | 5 | 0 | Go client library around the HERE location based APIs. | 2019-07-07 12:14:34 | 2022-11-05 14:12:21 |
| [go-sophos](https://github.com/esurdam/go-sophos) | 12 | 4 | 0 | Sophos UTM 9 REST API Client in Golang | 2018-09-05 04:37:25 | 2023-08-18 04:17:28 |
| [gomalshare](https://github.com/MonaxGT/gomalshare) | 12 | 2 | 0 | Go library MalShare API | 2019-03-01 09:33:41 | 2023-09-18 14:43:42 |
| [smitego](https://pkg.go.dev/github.com/slack-go/slack) | 11 | 1 | 0 | SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! | 2013-12-11 02:38:19 | 2023-07-25 13:50:09 |
| [gumblr](https://github.com/mattcunningham/gumblr) | 9 | 7 | 0 | A Go Wrapper for the Tumblr v2 API | 2015-07-09 23:13:51 | 2023-08-02 19:07:05 |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 9 | 0 | 0 | Go bindings for RRDA https://github.com/fcambus/rrda | 2014-09-15 21:06:16 | 2023-01-28 08:52:34 |
| [libgoffi](https://github.com/noctarius/libgoffi) | 9 | 1 | 0 | libgoffi - libffi adapter library for Go | 2019-08-03 17:05:34 | 2022-10-08 19:34:36 |
| [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | 9 | 2 | 0 | This is RAWG SDK GO. This library contains methods for interacting with RAWG API. | 2020-10-16 15:31:37 | 2023-09-19 14:08:59 |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 8 | 1 | 0 | Go client library for the SPTrans Olho Vivo API. :bus: | 2017-09-11 01:21:28 | 2022-12-03 19:02:18 |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 8 | 2 | 0 | :dancers: Go Chronos 3.x REST API Client | 2017-10-23 12:19:01 | 2023-05-24 01:27:43 |
| [go-google-email-audit-api](https://godoc.org/github.com/ngs/go-google-email-audit-api/emailaudit) | 8 | 5 | 0 | Go Client Library for G Suite Email Audit API | 2016-10-24 02:34:29 | 2023-07-18 07:57:25 |
| [go-zooz](https://github.com/gojuno/go-zooz) | 7 | 7 | 1 | Zooz API client for Go | 2017-07-04 09:28:23 | 2022-09-27 09:45:40 |
| [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) | 4 | 2 | 0 | Golang SDK for AppStore Connect API (Unofficial) | 2020-06-11 10:05:56 | 2023-11-02 08:46:30 |
| [kanka](https://kanka.io/en-US/docs/1.0) | 3 | 4 | 2 | Go client for the Kanka API | 2019-12-26 00:07:57 | 2021-12-10 23:36:57 |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 2 | 1 | 0 | A TripAdvisor API wrapper for Golang. | 2019-04-15 18:12:11 | 2022-09-27 09:45:21 |
| [vl-go](https://github.com/verifid/vl-go) | 2 | 1 | 0 | Go client library around the VerifID identity verification layer API. | 2019-02-09 12:46:53 | 2022-09-27 09:45:31 |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 2 | 1 | 0 | This is the official Playlyfe Golang Sdk | 2015-05-25 09:34:47 | 2022-09-27 09:44:46 |
</details>

### UUID
Libraries for working with UUIDs.

<sup>*Last Update: 2023-11-17 08:00:00*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [uuid](https://github.com/google/uuid) | 4,705 | 398 | 30 | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. | 2016-02-12 22:17:59 | 2023-11-16 18:24:38 |
| [ulid](https://github.com/oklog/ulid) | 3,820 | 148 | 3 | Universally Unique Lexicographically Sortable Identifier (ULID) in Go | 2016-12-06 15:26:52 | 2023-11-16 09:57:33 |
| [uuid](https://gof.rs/projects/uuid/) | 1,449 | 105 | 10 | A UUID package originally forked from github.com/satori/go.uuid | 2018-07-13 02:13:28 | 2023-11-16 00:56:25 |
| [wuid](https://github.com/edwingeng/wuid) | 503 | 48 | 2 | An extremely fast globally unique number generator. | 2018-01-27 01:16:28 | 2023-11-07 21:58:41 |
| [sno](https://pkg.go.dev/github.com/muyo/sno?tab=doc) | 87 | 5 | 1 | Compact, sortable and fast unique IDs with embedded metadata. | 2019-05-26 22:05:26 | 2023-09-08 05:25:42 |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 58 | 7 | 0 | A tiny and fast Go unique string generator | 2019-07-02 12:15:56 | 2023-07-10 12:46:11 |
| [Goid](https://github.com/jakenvac/Goid) | 38 | 4 | 1 | A UUIDv4 generation package written in go | 2017-05-19 10:40:45 | 2023-01-07 04:19:58 |
| [gouid](https://github.com/twharmon/gouid) | 21 | 4 | 0 | Fast, dependable universally unique ids | 2020-10-08 19:54:41 | 2023-07-15 15:39:02 |
| [uuid](https://github.com/agext/uuid) | 16 | 5 | 1 | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. | 2016-02-03 03:02:51 | 2023-08-11 14:41:54 |
| [GoFlake](https://github.com/Hart87/GoFlake) | 7 | 0 | 0 | A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake. | 2021-05-03 14:44:19 | 2021-06-30 15:29:33 |
</details>

### Utilities
General utilities and tools to make your life easier.

<sup>*Last Update: 2023-11-08 21:01:04*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fzf](https://github.com/junegunn/fzf) | 55,173 | 2,256 | 348 | :cherry_blossom: A command-line fuzzy finder | 2013-10-23 16:04:23 | 2023-11-05 01:09:38 |
| [hub](https://hub.github.com/) | 22,556 | 2,417 | 291 | A command-line tool that makes git easier to use with GitHub. | 2009-12-05 22:15:25 | 2023-11-05 22:54:54 |
| [ctop](https://ctop.sh) | 14,723 | 537 | 98 | Top-like interface for container metrics | 2016-12-27 02:25:57 | 2023-11-04 11:57:16 |
| [sqlx](http://jmoiron.github.io/sqlx/) | 14,553 | 1,051 | 342 | general purpose extensions to golang's database/sql | 2013-01-28 19:40:00 | 2023-11-08 13:38:53 |
| [goreleaser](https://goreleaser.com) | 12,258 | 906 | 32 | Deliver Go binaries as fast and easily as possible | 2016-12-21 17:13:39 | 2023-11-04 17:49:57 |
| [wuzz](https://github.com/asciimoo/wuzz) | 10,392 | 434 | 47 | Interactive cli tool for HTTP inspection | 2017-01-30 21:22:00 | 2023-11-07 20:00:20 |
| [usql](https://github.com/xo/usql) | 8,285 | 336 | 71 | Universal command-line interface for SQL databases | 2017-03-02 13:03:21 | 2023-11-08 06:22:44 |
| [peco](https://github.com/peco/peco) | 7,460 | 240 | 46 | Simplistic interactive filtering tool | 2014-06-06 06:06:32 | 2023-11-04 12:35:47 |
| [go-funk](https://github.com/thoas/go-funk) | 4,445 | 299 | 8 | A modern Go utility library which provides helpers (map, find, contains, filter, ...) | 2016-12-30 13:55:15 | 2023-11-03 16:27:29 |
| [godropbox](https://github.com/dropbox/godropbox) | 4,155 | 477 | 4 | Common libraries for writing Go services/applications. | 2014-06-22 23:09:29 | 2023-11-04 16:07:39 |
| [hystrix-go](https://github.com/afex/hystrix-go) | 4,064 | 512 | 57 | Netflix's Hystrix latency and fault tolerance library, for Go  | 2013-12-15 08:51:23 | 2023-11-05 03:11:54 |
| [minify](https://go.tacodewolff.nl/minify) | 3,457 | 241 | 13 | Go minifiers for web formats | 2014-05-21 09:03:48 | 2023-10-31 14:12:47 |
| [panicparse](https://maruel.ca) | 3,457 | 101 | 3 | Crash your app in style (Golang) | 2015-02-02 02:14:41 | 2023-11-01 17:23:27 |
| [goreporter](https://github.com/qax-os/goreporter) | 3,089 | 319 | 29 | A Golang tool that does static analysis, unit testing, code review and generate code quality report. | 2017-03-27 08:46:38 | 2023-11-03 21:50:38 |
| [mergo](https://github.com/darccio/mergo) | 2,588 | 263 | 11 | Mergo: merging Go structs and maps since 2013 | 2013-03-11 22:51:11 | 2023-11-05 02:06:40 |
| [mc](https://min.io/download) | 2,576 | 491 | 35 | Simple | Fast tool to manage MinIO clusters :cloud: | 2015-01-16 02:56:51 | 2023-11-03 21:58:22 |
| [cli](https://github.com/create-go-app/cli/wiki) | 2,219 | 244 | 1 | ✨ A complete and self-contained solution for developers of any qualification to create a production-ready project with backend (Go), frontend (JavaScript, TypeScript) and deploy automation (Ansible, Docker) by running only one CLI command. | 2019-12-30 22:08:38 | 2023-11-04 21:15:34 |
| [spinner](https://github.com/briandowns/spinner) | 2,177 | 131 | 13 | Go (golang) package with 90 configurable terminal spinner/progress indicators. | 2014-12-13 00:36:19 | 2023-11-07 12:38:47 |
| [storm](https://github.com/asdine/storm) | 2,010 | 139 | 64 | Simple and powerful toolkit for BoltDB | 2016-01-10 12:55:59 | 2023-11-06 09:40:44 |
| [filetype](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 1,918 | 169 | 41 | Fast, dependency-free Go package to infer binary file types based on the magic numbers header signature | 2015-09-24 09:15:51 | 2023-11-03 08:16:53 |
| [jump](http://gsamokovarov.com/jump) | 1,694 | 62 | 3 | Jump helps you navigate faster by learning your habits. ✌️ | 2015-08-16 22:07:17 | 2023-11-03 13:58:03 |
| [mole](https://davrodpin.github.io/mole/) | 1,665 | 97 | 26 | CLI application to create ssh tunnels focused on resiliency and user experience. | 2018-10-04 02:38:00 | 2023-11-02 19:59:42 |
| [boilr](https://github.com/tmrts/boilr) | 1,661 | 124 | 45 | :zap: boilerplate template manager that generates files or directories from template repositories | 2015-12-19 16:57:26 | 2023-11-04 17:07:57 |
| [mimetype](https://pkg.go.dev/github.com/gabriel-vasile/mimetype#pkg-overview) | 1,261 | 146 | 68 | A fast Golang library for media type and file extension detection, based on magic numbers | 2018-07-02 07:15:29 | 2023-11-05 12:21:10 |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 1,088 | 151 | 21 | Circuit Breakers in Go | 2014-07-17 22:41:33 | 2023-10-31 00:57:50 |
| [scany](https://github.com/georgysavva/scany) | 1,065 | 58 | 4 | Library for scanning data from a database into Go structs and more | 2020-07-02 11:02:58 | 2023-11-07 00:07:42 |
| [hostctl](http://guumaster.github.io/hostctl) | 953 | 44 | 14 | Your dev tool to manage /etc/hosts like a pro! | 2020-03-14 11:29:02 | 2023-11-02 13:36:52 |
| [gtm](https://github.com/git-time-metric/gtm) | 949 | 54 | 51 | Simple, seamless, lightweight time tracking for Git | 2016-06-19 21:17:04 | 2023-10-31 03:02:04 |
| [immortal](https://immortal.run) | 777 | 53 | 5 | ⭕  A *nix cross-platform (OS agnostic) supervisor | 2016-06-30 17:02:27 | 2023-10-31 06:19:27 |
| [circuit](https://github.com/cep21/circuit) | 718 | 44 | 4 | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. | 2017-12-23 22:17:43 | 2023-11-03 15:10:06 |
| [delve](https://github.com/derekparker/delve) | 605 | 117 | 1 | Delve is a debugger for the Go programming language. | 2020-02-18 18:03:33 | 2023-11-01 03:54:11 |
| [ergo](https://github.com/cristianoliveira/ergo) | 589 | 61 | 6 | The management of multiple apps running over different ports made easy | 2017-08-19 18:41:56 | 2023-10-27 05:33:18 |
| [htcat](https://github.com/htcat/htcat) | 553 | 31 | 5 | Parallel and Pipelined HTTP GET Utility | 2013-08-05 11:17:01 | 2023-10-22 09:29:54 |
| [clockwork](https://github.com/jonboulle/clockwork) | 534 | 55 | 2 | a fake clock for golang | 2014-09-09 18:24:00 | 2023-11-03 16:19:40 |
| [changie](https://changie.dev/) | 534 | 31 | 5 | Automated changelog tool for preparing releases with lots of customization options | 2020-12-05 19:38:33 | 2023-11-03 20:53:39 |
| [koazee](https://github.com/wesovilabs/koazee) | 519 | 31 | 16 | A StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices. | 2018-11-09 09:49:19 | 2023-10-02 12:31:30 |
| [godaemon](https://github.com/VividCortex/godaemon) | 492 | 54 | 9 | Daemonize Go applications deviously. | 2013-08-01 17:16:30 | 2023-10-05 13:41:36 |
| [go-dry](https://github.com/ungerik/go-dry) | 487 | 38 | 0 | DRY (don't repeat yourself) package for Go | 2014-02-28 13:49:31 | 2023-09-18 15:01:30 |
| [scan](https://github.com/blockloop/scan) | 460 | 30 | 0 | Tiny lib to scan SQL rows directly to structs, slices, and primitive types | 2017-11-27 23:22:18 | 2023-10-30 14:39:03 |
| [gubrak](https://pkg.go.dev/github.com/novalagung/gubrak) | 454 | 39 | 0 | ⚙️ Golang functional utility library with syntactic sugar. It's like lodash, but for Go | 2018-03-09 11:28:05 | 2023-11-03 02:39:25 |
| [gopencils](https://github.com/bndr/gopencils) | 449 | 43 | 7 | Easily consume REST APIs with Go (golang) | 2014-06-23 11:41:24 | 2023-09-20 07:05:59 |
| [request](https://godoc.org/github.com/mozillazg/request) | 427 | 40 | 6 | A developer-friendly HTTP request library for Gopher. | 2014-12-21 04:30:42 | 2023-09-21 05:48:05 |
| [deepcopier](https://github.com/ulule/deepcopier) | 424 | 58 | 7 | simple struct copying for golang | 2015-07-24 18:01:01 | 2023-10-29 11:33:14 |
| [go-rate](https://github.com/beefsack/go-rate) | 391 | 38 | 0 | A timed rate limiter for Go | 2014-08-25 04:42:34 | 2023-10-05 17:57:42 |
| [retry](https://pkg.go.dev/github.com/kamilsk/retry/v5) | 334 | 14 | 10 | ♻️ The most advanced interruptible mechanism to perform actions repetitively until successful. | 2016-11-02 20:20:43 | 2023-10-04 06:01:00 |
| [serve](https://syntaqx-serve.herokuapp.com/) | 312 | 20 | 7 | 🍽️ a static http server anywhere you need one. | 2019-01-10 23:31:52 | 2023-10-31 00:04:12 |
| [countries](https://pkg.go.dev/github.com/biter777/countries?tab=doc) | 297 | 60 | 8 | Countries - ISO-639, ISO-3166 countries codes with subdivisions and names, ISO-4217 currency designators, ITU-T E.164 IDD phone codes, countries capitals, UN M.49 codes, IANA ccTLD countries domains, IOC/NOC and FIFA codes, VERY VERY FAST, compatible with Databases/JSON/BSON/GOB/XML/CSV, Emoji countries flags and currencies support, Unicode CLDR. | 2019-04-22 14:47:11 | 2023-11-04 18:10:53 |
| [util](https://github.com/shomali11/util) | 280 | 40 | 1 | A collection of useful utility functions | 2017-05-24 00:21:29 | 2023-11-07 19:18:10 |
| [gotenv](https://github.com/subosito/gotenv) | 273 | 35 | 0 | Load environment variables from `.env` or `io.Reader` in Go. | 2013-08-27 12:56:47 | 2023-11-02 13:37:49 |
| [gohper](https://github.com/cosiner/gohper) | 256 | 47 | 0 | [UNMATAINED] common libs here. | 2015-03-23 22:46:12 | 2023-09-18 15:01:38 |
| [go-trigger](https://github.com/sadlil/go-trigger) | 236 | 41 | 1 | A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. | 2015-10-19 09:42:17 | 2023-08-22 20:28:15 |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 232 | 9 | 2 | Pattern matchings for Go. | 2018-12-11 20:11:17 | 2023-10-17 03:19:17 |
| [limiters](https://godoc.org/github.com/mennanov/limiters) | 231 | 33 | 5 | Golang rate limiters for distributed applications | 2019-08-28 18:09:54 | 2023-11-03 00:02:31 |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 206 | 66 | 27 | go-sitemap-generator is the easiest way to generate Sitemaps in Go | 2015-10-12 16:23:13 | 2023-09-27 15:50:18 |
| [toolbox](https://github.com/viant/toolbox) | 193 | 29 | 4 | Toolbox - go utility library | 2016-06-13 19:33:35 | 2023-08-28 01:56:38 |
| [death](https://vreco.fly.dev/blog/post/Concurrent%20Graceful%20Shutdown%20in%20Go) | 190 | 20 | 0 | Managing go application shutdown with signals. | 2015-03-09 03:50:40 | 2023-06-19 08:11:27 |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 184 | 11 | 0 | go-bind-plugin generates API for exported plugin symbols (-buildmode=plugin) - go1.8+ only (http://golang.org/pkg/plugin) | 2016-11-08 14:40:26 | 2023-09-25 20:04:32 |
| [rerun](https://github.com/ivpusic/rerun) | 167 | 8 | 0 | Configurable recompiling and rerunning go apps when source changes | 2014-12-10 00:29:54 | 2023-09-18 15:02:26 |
| [moldova](https://github.com/StabbyCutyou/moldova) | 166 | 7 | 0 | A lightweight templating system for generating random data | 2016-01-30 05:25:39 | 2023-10-23 05:53:25 |
| [apm](https://github.com/topfreegames/apm) | 164 | 86 | 9 | APM is a process manager for Golang applications. | 2015-11-18 16:56:48 | 2023-08-10 17:37:47 |
| [robustly](https://github.com/VividCortex/robustly) | 156 | 8 | 1 | Run functions resiliently in Go, catching and restarting panics | 2013-07-08 13:27:10 | 2023-10-05 13:34:37 |
| [chyle](https://github.com/antham/chyle) | 153 | 12 | 0 | Changelog generator : use a git repository and various data sources and publish the result on external services | 2016-11-17 21:14:44 | 2023-10-14 01:32:28 |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 148 | 25 | 0 | Pure Go bsdiff and bspatch libraries and CLI tools. | 2019-02-23 23:33:50 | 2023-10-25 16:28:13 |
| [nostromo](https://nostromo.sh/) | 137 | 8 | 12 | 👽 CLI for building powerful aliases and tools | 2019-07-13 04:51:46 | 2023-10-25 12:18:35 |
| [cmd](https://github.com/commander-cli/cmd) | 136 | 18 | 2 | A simple package to execute shell commands on linux, windows and osx | 2019-09-27 13:22:06 | 2023-10-22 07:31:37 |
| [goval](https://github.com/maja42/goval) | 135 | 22 | 4 | Expression evaluation in golang | 2018-06-17 15:43:44 | 2023-10-21 12:23:16 |
| [onecache](https://github.com/adelowo/onecache) | 133 | 8 | 0 | One caching API, Multiple backends | 2017-04-14 21:49:15 | 2023-08-05 16:14:14 |
| [filter](https://godoc.org/github.com/gookit/filter) | 133 | 10 | 1 | ⏳ Provide filtering, sanitizing, and conversion of Golang data. 提供对Golang数据的过滤，净化，转换。 | 2018-09-26 09:11:13 | 2023-09-12 00:04:32 |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 124 | 15 | 2 | LiveReload server for Go [golang] | 2014-07-15 05:36:53 | 2023-05-21 19:45:34 |
| [mongo-go-pagination](https://mongo-go-pagination.herokuapp.com/normal-pagination?page=1&limit=15) | 124 | 36 | 4 | Golang Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines with all information like Total records, Page, Per Page, Previous, Next, Total Page and query results. | 2020-02-04 08:23:33 | 2023-10-06 09:33:12 |
| [sorty](https://github.com/jfcg/sorty) | 123 | 5 | 0 | :zap: Fast Concurrent / Parallel Sorting in Go | 2019-02-18 21:05:45 | 2023-09-20 09:34:54 |
| [goseaweedfs](https://github.com/chrislusf/seaweedfs) | 109 | 46 | 1 | A complete Golang client for SeaweedFS | 2017-07-20 04:35:39 | 2023-11-02 02:00:57 |
| [go-lock](https://github.com/viney-shih/go-lock) | 102 | 8 | 0 | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation | 2020-04-30 11:40:21 | 2023-10-14 01:09:07 |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 101 | 15 | 1 | Database client library, proxy for any master slave, master master structures. Lightweight, performant and auto balancing in mind. | 2016-12-26 04:05:09 | 2023-03-06 13:49:21 |
| [xferspdy](https://github.com/monmohan/xferspdy) | 99 | 12 | 3 | Xferspdy provides binary diff and patch library in golang. [Mentioned in Awesome Go, https://github.com/avelino/awesome-go] | 2015-05-22 13:23:34 | 2023-09-18 15:02:29 |
| [go-health](https://github.com/Talento90/go-health) | 94 | 4 | 0 | :heart: Health check your applications and dependencies | 2018-02-13 18:40:54 | 2023-09-18 15:01:30 |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 92 | 9 | 6 | Powerful and versatile MIME sniffing package using pre-compiled glob patterns, magic number signatures, XML document namespaces, and tree magic for mounted volumes, generated from the XDG shared-mime-info database. | 2018-10-11 16:12:54 | 2023-10-09 07:11:38 |
| [pgo](https://github.com/arthurkushman/pgo) | 83 | 16 | 1 | Go library for PHP community with convenient functions | 2018-12-26 06:59:47 | 2023-09-18 15:02:20 |
| [repeat](https://github.com/ssgreg/repeat) | 82 | 7 | 0 | Go implementation of different backoff strategies useful for retrying operations and heartbeating. | 2017-11-22 07:06:47 | 2023-09-18 15:02:25 |
| [pm](https://github.com/VividCortex/pm) | 79 | 7 | 2 | Processlist manager with TCP listener | 2013-11-17 19:17:01 | 2023-10-05 13:35:27 |
| [handy](https://pkg.go.dev/github.com/novalagung/gubrak) | 75 | 7 | 0 | GO Golang Utilities and helpers like validators and string formatters | 2018-06-13 13:10:07 | 2023-09-18 15:01:48 |
| [netbug](https://github.com/e-dard/netbug) | 73 | 5 | 0 | Package netbug provides a handler for registering profilers on your own ServeMux. | 2015-03-05 19:27:29 | 2023-10-27 10:23:03 |
| [unis](https://godoc.org/github.com/esemplastic/unis) | 70 | 4 | 2 | UNIS: A Common Architecture for String Utilities within the Go Programming Language. | 2017-05-06 05:01:03 | 2022-09-27 09:55:16 |
| [multitick](https://github.com/VividCortex/multitick) | 69 | 3 | 1 | A multiplexor for aligned time.Time tickers in Go | 2013-12-10 16:47:26 | 2023-10-25 18:46:37 |
| [goreadability](https://github.com/philipjkim/goreadability) | 68 | 8 | 2 | Webpage summary extractor using Facebook Open Graph and arc90's readability | 2016-04-20 01:40:14 | 2023-09-18 15:01:42 |
| [retry](https://pkg.go.dev/github.com/thedevsaddam/retry) | 63 | 6 | 0 | Simple and easy retry mechanism package for Go | 2018-02-25 19:08:03 | 2023-03-28 15:38:00 |
| [minquery](https://github.com/icza/minquery) | 61 | 20 | 4 | MongoDB / mgo query that supports efficient pagination (cursors to continue listing documents where we left off). | 2016-11-16 12:23:07 | 2023-08-26 14:24:40 |
| [golog](https://github.com/mlimaloureiro/golog) | 60 | 12 | 15 | Easy and simple CLI time tracker for your tasks | 2016-01-09 15:43:47 | 2023-09-04 13:29:40 |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 60 | 9 | 2 | Parse TODOs in your GO code | 2016-10-17 19:51:36 | 2023-04-05 16:59:54 |
| [dbt](https://github.com/nikogura/dbt) | 58 | 7 | 6 | A delivery system for running self-updating, signed tools. | 2017-11-30 22:53:17 | 2023-10-16 19:09:01 |
| [beyond](http://wesovilabs.github.io/beyond) | 55 | 11 | 9 | The Go library that will drive you to AOP world! | 2019-10-18 05:41:45 | 2023-09-18 14:59:20 |
| [backscanner](https://github.com/icza/backscanner) | 54 | 9 | 0 | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. | 2017-10-18 07:59:07 | 2023-10-17 22:56:48 |
| [shutdown](https://github.com/ztrue/shutdown) | 52 | 7 | 0 | Golang app shutdown hooks. | 2018-11-17 17:56:03 | 2023-10-29 02:43:43 |
| [golarm](https://github.com/msempere/golarm) | 51 | 10 | 0 | Fire alarms with system events | 2015-08-14 16:51:53 | 2023-09-18 15:01:39 |
| [slice](https://github.com/psampaz/slice) | 51 | 5 | 1 | Type-safe functions for common Go slice operations | 2019-11-26 05:20:39 | 2023-05-25 06:24:35 |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 50 | 12 | 10 | Universal copy paste service, works across different machines! | 2017-01-28 15:35:24 | 2023-08-18 03:39:39 |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 48 | 6 | 2 | Retrying made simple and easy for golang :repeat:  | 2017-06-09 16:07:37 | 2023-08-07 15:40:30 |
| [goback](https://github.com/carlescere/goback) | 48 | 9 | 6 | Golang simple exponential backoff package. | 2015-03-13 16:09:18 | 2022-11-25 00:58:17 |
| [intrinsic](https://immortal.run) | 46 | 2 | 1 | Provide Golang native SIMD intrinsics on x86/amd64 platform | 2017-06-13 09:26:34 | 2023-06-08 07:54:16 |
| [go-httpheader](https://godoc.org/github.com/mozillazg/go-httpheader) | 45 | 14 | 0 | A Go library for encoding structs into Header fields. | 2017-06-24 11:28:06 | 2023-10-11 09:37:41 |
| [slicer](https://github.com/leaanthony/slicer) | 44 | 3 | 0 | Utility class for handling slices | 2019-01-10 09:55:25 | 2023-11-02 12:30:31 |
| [equalizer](https://pkg.go.dev/github.com/reugn/equalizer) | 43 | 2 | 1 | A rate limiters package for Go | 2019-06-14 09:25:13 | 2023-07-17 19:16:08 |
| [gostrutils](https://github.com/chrislusf/seaweedfs) | 43 | 7 | 1 | Collections of string utils I have created over the years | 2018-09-19 11:06:11 | 2023-11-02 12:28:56 |
| [gpath](https://github.com/tenntenn/gpath) | 40 | 4 | 0 | gpath is a Go package to access a field by a path using reflect pacakge | 2017-05-24 06:24:18 | 2023-02-15 12:08:03 |
| [evaluator](https://github.com/nullne/evaluator) | 37 | 8 | 0 | The management of multiple apps running over different ports made easy | 2017-04-27 18:31:46 | 2023-01-11 11:43:14 |
| [copy](https://github.com/gotidy/copy) | 37 | 4 | 4 | Package for fast copying structs of different types | 2020-10-09 06:59:08 | 2023-10-12 03:40:18 |
| [ghokin](https://github.com/antham/ghokin) | 37 | 3 | 1 | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...) | 2018-08-03 11:36:35 | 2023-10-14 01:32:20 |
| [rclient](https://github.com/zpatrick/rclient) | 35 | 3 | 2 | Minimalistic REST client for Go applications | 2017-02-28 01:07:25 | 2022-12-24 03:03:55 |
| [myhttp](https://github.com/inancgumus/myhttp) | 35 | 14 | 1 | Simplest HTTP GET requester for Go with timeout support | 2017-09-13 15:48:47 | 2023-01-28 17:28:10 |
| [tome](https://github.com/cyruzin/tome) | 35 | 3 | 1 | Package tome was designed to paginate simple RESTful APIs. | 2019-04-12 16:49:45 | 2023-09-23 07:10:36 |
| [generate](https://github.com/go-playground/generate) | 30 | 6 | 0 | :runner:runs go generate recursively on a specified path or environment variable and can filter by regex | 2015-11-15 01:52:04 | 2023-10-03 00:25:15 |
| [mimesniffer](https://pkg.go.dev/github.com/aofei/mimesniffer) | 30 | 1 | 4 | A MIME type sniffer for Go. | 2018-12-20 03:40:20 | 2023-10-31 05:56:24 |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 27 | 7 | 1 | a small golang lib to generate placeholder images | 2014-10-12 00:50:46 | 2023-09-18 15:01:41 |
| [ugo](https://github.com/alxrm/ugo) | 27 | 5 | 0 | Simple and expressive toolbox written in Go | 2016-02-17 19:41:57 | 2023-03-05 01:02:37 |
| [rerate](https://github.com/abo/rerate) | 25 | 5 | 1 | redis-based rate counter and rate limiter | 2016-05-24 08:59:00 | 2023-07-29 14:55:28 |
| [ptr](https://github.com/gotidy/ptr) | 23 | 4 | 1 | Contains functions for simplified creation of pointers from constants of basic types | 2019-12-25 15:29:48 | 2023-09-18 15:02:23 |
| [ctxutil](https://ctop.sh) | 23 | 4 | 1 | utils for Go context | 2018-07-30 11:28:57 | 2023-07-21 12:30:49 |
| [structs](https://github.com/PumpkinSeed/structs) | 22 | 3 | 0 | Golang struct operations. | 2017-08-26 09:59:00 | 2023-05-14 16:06:11 |
| [go-convert](https://github.com/Eun/go-convert) | 21 | 3 | 0 | Convert a value into another type | 2019-06-07 16:56:38 | 2023-09-18 15:01:22 |
| [jsend](https://clevergo.tech) | 20 | 6 | 0 | :100: JSend's implementation writen in Go(golang) | 2020-01-14 04:41:36 | 2023-09-18 15:02:10 |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 17 | 2 | 0 | Problem json implementation (https://tools.ietf.org/html/rfc7807) package for go | 2019-05-16 05:42:14 | 2023-10-29 17:16:52 |
| [go-types](https://github.com/mikekonan/go-types) | 17 | 9 | 0 | Library providing opanapi3 and Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. | 2021-04-21 11:34:25 | 2023-10-20 02:07:11 |
| [filler](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 17 | 4 | 0 | fill struct data easily with fill tags | 2017-04-05 08:14:04 | 2022-09-27 09:48:27 |
| [dlog](https://github.com/kirillDanshin/dlog) | 17 | 2 | 0 | Simple build-time controlled debug log with ability to log where the logger was called | 2016-07-04 19:59:09 | 2022-09-27 09:47:26 |
| [okrun](https://github.com/xta/okrun) | 16 | 3 | 0 | ok, run your gofile | 2014-10-01 06:18:56 | 2022-09-27 09:52:52 |
| [rest-go](https://github.com/edermanoel94/rest-go) | 16 | 2 | 1 | A package that provide many helpful methods for working with rest api. | 2019-07-29 18:56:08 | 2022-09-27 09:53:28 |
| [command](https://github.com/txgruppi/command) | 14 | 4 | 0 | Command pattern for Go with thread safe serial and parallel dispatcher | 2015-08-24 09:43:50 | 2022-09-27 09:46:53 |
| [silk](https://github.com/chrispassas/silk) | 13 | 4 | 0 | Silk File Reader | 2018-12-18 04:23:35 | 2023-09-12 16:25:51 |
| [go-countries](https://github.com/mikekonan/go-countries) | 12 | 5 | 0 | Convert a value into another type | 2020-10-27 12:56:40 | 2023-10-19 07:39:25 |
| [go-clip](https://github.com/prashantgupta24/go-clip) | 12 | 0 | 2 | A minimalistic clipboard manager for Mac. | 2020-11-18 22:19:01 | 2023-05-24 07:11:28 |
| [retry](https://github.com/shafreeck/retry) | 12 | 2 | 1 | A pretty simple library to ensure your work to be done | 2018-07-18 09:48:33 | 2023-11-06 20:19:22 |
| [blank](https://github.com/Henry-Sarabia/blank) | 11 | 1 | 0 | Detect blank strings or remove whitespace from strings | 2019-02-13 00:07:27 | 2023-09-18 14:59:32 |
| [statiks](https://github.com/janiltonmaciel/statiks) | 10 | 1 | 0 | Fast, zero-configuration, static HTTP filer server. | 2018-06-26 23:42:33 | 2022-09-27 09:54:59 |
| [bleep](https://github.com/sinhashubham95/bleep) | 10 | 3 | 0 | OS Signal Handlers in Go | 2021-01-02 05:22:08 | 2023-09-28 09:59:34 |
| [retry](https://github.com/percolate/retry) | 9 | 2 | 0 | Percolate's Go retry package | 2018-06-15 19:23:36 | 2023-03-16 18:34:54 |
| [nfdump](https://github.com/chrispassas/nfdump) | 8 | 2 | 0 | NFDump File Reader | 2020-04-08 01:01:22 | 2023-05-15 16:20:56 |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 8 | 1 | 0 | Slice conversion between primitive types | 2019-02-15 06:50:34 | 2022-09-27 09:54:48 |
| [goctx](https://github.com/zerosnake0/goctx) | 6 | 3 | 0 | Get your context value faster | 2020-11-14 14:16:09 | 2023-10-04 00:35:08 |
| [lets-go](https://github.com/aplescia/lets-go) | 6 | 1 | 0 | Go module that provides common utilities for Cloud Native development | 2020-02-19 16:32:41 | 2023-06-20 23:24:05 |
| [olaf](https://github.com/btnguyen2k/olaf) | 6 | 1 | 0 | Twitter Snowflake implemented in Go | 2019-01-03 13:31:10 | 2023-09-18 15:02:16 |
| [tik](https://github.com/andy2046/tik) | 5 | 2 | 0 | hierarchical timing wheel | 2020-07-04 09:13:49 | 2022-11-29 03:42:44 |
| [go-safe](https://github.com/kenkyu392/go-safe) | 4 | 0 | 0 | This Go package provides a sandbox for the safe execution of panic-inducing programs | 2019-10-29 15:20:37 | 2021-07-31 04:33:24 |
</details>

### Validation
Libraries for validation.

<sup>*Last Update: 2023-11-17 07:59:53*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [validator](https://github.com/go-playground/validator) | 14,435 | 1,228 | 256 | :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving | 2015-02-12 16:32:22 | 2023-11-16 19:19:54 |
| [govalidator](https://github.com/asaskevich/govalidator) | 5,833 | 592 | 167 | [Go] Package of validators and sanitizers for strings, numerics, slices and structs | 2014-06-20 10:45:23 | 2023-11-15 11:13:47 |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 3,400 | 215 | 48 | An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags. | 2016-06-22 03:47:43 | 2023-11-15 06:32:59 |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 1,253 | 120 | 43 | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. | 2017-09-13 16:42:20 | 2023-11-13 06:24:48 |
| [validate](https://gookit.github.io/validate/) | 959 | 113 | 26 | ⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。 | 2018-07-16 08:23:49 | 2023-11-09 14:17:00 |
| [checkdigit](https://github.com/osamingo/checkdigit) | 103 | 7 | 1 | Provide check digit algorithms and calculators written in Go | 2019-04-05 09:46:36 | 2023-10-07 14:03:27 |
| [validate](https://gookit.github.io/validate/) | 91 | 21 | 0 | This package provides a framework for writing validations for Go applications. | 2018-02-10 18:25:55 | 2023-10-26 20:45:24 |
| [jio](https://github.com/faceair/jio) | 91 | 11 | 1 | jio is a json schema validator similar to joi | 2018-10-28 11:02:45 | 2023-10-10 22:40:50 |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 80 | 9 | 1 | A norms and conventions validator for Terraform | 2019-05-29 11:37:15 | 2023-09-29 13:16:57 |
| [gody](https://pkg.go.dev/github.com/guiferpa/gody) | 65 | 5 | 1 | :balloon: A lightweight struct validator for Go | 2018-11-01 21:08:16 | 2023-10-15 21:45:58 |
| [govalid](https://github.com/twharmon/govalid) | 35 | 6 | 1 | Struct validation using tags | 2019-02-17 23:25:43 | 2023-09-27 03:51:47 |
</details>

### Version Control
Libraries for version control.

<sup>*Last Update: 2023-12-07 21:47:29*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-git](https://pkg.go.dev/github.com/go-git/go-git/v5) | 5,174 | 684 | 419 | A highly extensible Git implementation in pure Go. | 2019-12-19 10:27:02 | 2023-12-07 09:22:29 |
| [hercules](https://github.com/src-d/hercules) | 1,965 | 151 | 48 | Gaining advanced insights from Git repository history. | 2016-12-12 17:30:29 | 2023-12-07 11:27:15 |
| [git2go](https://github.com/libgit2/git2go) | 1,877 | 319 | 80 | Git to Go; bindings for libgit2. Like McDonald's but tastier. | 2013-03-05 19:50:43 | 2023-12-03 20:38:34 |
| [gh](https://github.com/rjeczalik/gh) | 82 | 13 | 2 | Scriptable server and net/http middleware for GitHub Webhooks. | 2015-03-08 21:04:05 | 2023-09-27 03:53:54 |
| [go-vcs](https://sourcegraph.com/sourcegraph/go-vcs) | 77 | 20 | 22 | manipulate and inspect VCS repositories in Go | 2013-06-02 02:36:18 | 2023-10-29 18:06:46 |
| [hgo](https://github.com/beyang/hgo) | 17 | 4 | 0 | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. | 2014-06-18 03:54:40 | 2023-09-27 03:54:35 |
</details>

### Video
Libraries for manipulating video.

<sup>*Last Update: 2023-11-17 07:59:48*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [goav](https://github.com/giorgisio/goav) | 2,032 | 413 | 48 | Golang bindings for FFmpeg (This repository is no longer maintained) | 2015-05-21 05:31:14 | 2023-11-15 08:50:08 |
| [m3u8](http://tools.ietf.org/html/draft-pantos-http-live-streaming) | 1,097 | 298 | 62 | Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema: | 2013-02-05 22:26:30 | 2023-11-16 06:50:09 |
| [gmf](https://github.com/3d0c/gmf) | 863 | 215 | 47 | Go Media Framework | 2013-04-03 09:07:47 | 2023-11-15 11:26:27 |
| [go-astisub](https://github.com/asticode/go-astisub) | 517 | 102 | 18 | Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.) | 2016-12-16 14:47:59 | 2023-11-14 07:53:13 |
| [go-astits](https://github.com/asticode/go-astits) | 507 | 52 | 13 | Demux and mux MPEG Transport Streams (.ts) natively in GO | 2017-07-04 13:06:15 | 2023-11-04 00:11:31 |
| [libvlc-go](https://pkg.go.dev/github.com/adrg/libvlc-go/v3) | 376 | 48 | 6 | Go bindings for libVLC and high-level media player interface | 2015-01-06 14:01:50 | 2023-11-16 10:33:21 |
| [gst](https://github.com/ziutek/gst) | 168 | 48 | 9 | Go bindings for GStreamer (retired: currently I don't use/develop this package) | 2011-07-26 00:44:40 | 2023-08-27 12:31:57 |
| [go-m3u8](https://tools.ietf.org/html/rfc8216) | 110 | 24 | 2 | Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8) | 2018-11-06 02:42:27 | 2023-11-16 22:37:48 |
| [v4l](https://pkg.go.dev/github.com/korandiz/v4l) | 77 | 14 | 0 | Facade to the Video4Linux video capture interface.  | 2016-10-25 10:50:25 | 2023-09-18 15:05:05 |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 24 | 5 | 0 | golang library to read and write various subtitle formats | 2017-05-03 21:05:25 | 2023-08-17 13:54:41 |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 18 | 6 | 0 | Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files | 2018-11-02 19:09:07 | 2023-11-11 09:34:37 |
</details>

### Web Frameworks
Full stack web frameworks.

<sup>*Last Update: 2023-11-18 20:14:56*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gin](https://gin-gonic.com/) | 72,585 | 7,759 | 760 | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 2014-06-16 23:57:25 | 2023-11-18 13:06:53 |
| [fiber](https://gofiber.io) | 29,208 | 1,496 | 75 | ⚡️ Express inspired web framework written in Go | 2020-01-16 03:59:20 | 2023-11-18 11:56:23 |
| [echo](https://echo.labstack.com) | 27,061 | 2,231 | 75 | High performance, minimalist Go web framework | 2015-03-01 17:43:01 | 2023-11-18 11:42:17 |
| [revel](http://revel.github.io) | 13,001 | 1,416 | 99 | A high productivity, full-stack web framework for the Go language. | 2011-12-09 04:10:26 | 2023-11-18 00:20:40 |
| [buffalo](http://gobuffalo.io) | 7,907 | 614 | 26 | Rapid Web Development w/ Go | 2014-10-22 17:35:14 | 2023-11-18 11:16:56 |
| [goa](https://goa.design) | 5,334 | 600 | 9 | Design-based APIs and microservices in Go | 2014-12-05 07:17:53 | 2023-11-17 17:33:35 |
| [gizmo](https://open.nytimes.com/introducing-gizmo-aa7ea463b208) | 3,740 | 237 | 33 | A Microservice Toolkit from The New York Times | 2015-12-15 18:09:36 | 2023-11-16 16:05:47 |
| [go-json-rest](https://ant0ine.github.io/go-json-rest/) | 3,511 | 417 | 47 | A quick and easy way to setup a RESTful JSON API | 2013-02-19 03:15:45 | 2023-11-13 04:17:35 |
| [macaron](https://go-macaron.com) | 3,437 | 291 | 6 | Package macaron is a high productive and modular web framework in Go. | 2014-07-10 03:13:30 | 2023-11-12 05:06:30 |
| [utron](https://github.com/gernest/utron) | 2,221 | 166 | 8 | A lightweight MVC framework for Go(Golang) | 2015-09-16 07:55:54 | 2023-11-07 12:40:35 |
| [goyave](https://goyave.dev) | 1,391 | 58 | 3 | 🍐 Elegant Golang REST API Framework (v5 preview available) | 2019-10-21 09:44:34 | 2023-11-14 09:24:44 |
| [rest-layer]( http://rest-layer.io) | 1,232 | 116 | 36 | REST Layer, Go (golang) REST API framework | 2015-07-29 19:16:20 | 2023-11-16 09:01:41 |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 996 | 78 | 28 | A Go framework for building JSON web services inspired by Dropwizard | 2013-02-09 21:16:13 | 2023-08-16 03:14:43 |
| [tango](https://github.com/lunny/tango) | 834 | 107 | 9 | This is only a mirror and Moved to https://gitea.com/lunny/tango | 2014-12-17 03:07:09 | 2023-09-23 18:40:09 |
| [gearbox](https://gogearbox.com) | 726 | 57 | 3 | Gearbox :gear: is a web framework written in Go with a focus on high performance | 2020-04-25 01:28:37 | 2023-11-15 12:18:15 |
| [aah](https://aahframework.org) | 685 | 39 | 19 | A secure, flexible, rapid Go web framework | 2016-06-27 04:47:45 | 2023-09-25 07:55:41 |
| [beego](beego.me) | 678 | 177 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2023-11-17 13:30:45 |
| [aero](https://github.com/aerogo/aero) | 544 | 39 | 4 | :bullettrain_side: High-performance web server for Go. | 2016-11-09 13:02:13 | 2023-10-26 07:10:23 |
| [gongular](http://gondolaweb.com) | 502 | 19 | 8 | A different approach to Go web frameworks | 2016-06-22 11:52:42 | 2023-11-09 19:23:03 |
| [flamingo-commerce](https://www.flamingo.me/flamingo-commerce.html) | 454 | 71 | 34 | Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce "Portals" and connect it with the help of individual Adapters to other services.  | 2019-04-02 15:11:57 | 2023-11-11 21:31:38 |
| [air](https://pkg.go.dev/github.com/aofei/air) | 437 | 43 | 8 | An ideally refined web framework for Go. | 2016-07-20 12:09:48 | 2023-11-06 01:01:18 |
| [neo](http://ivpusic.github.io/neo/) | 418 | 43 | 6 | Go Web Framework | 2015-02-04 19:16:06 | 2023-05-19 20:43:33 |
| [confetti](https://confetti-framework.github.io/docs/) | 410 | 18 | 1 | Confetti is a web application framework with an expressive, elegant syntax. This repository contains configuration files and is intended as a template for your codebase. Download these configuration files and include them in your git repository. | 2019-11-01 23:14:21 | 2023-11-14 03:02:30 |
| [flamingo](http://www.flamingo.me) | 409 | 45 | 40 | Flamingo Framework and Core Library. Flamingo is a go based framework to build pluggable applications. Focus is on clean architecture, maintainability and operation readiness. | 2019-04-02 12:24:02 | 2023-11-10 06:50:17 |
| [mango](http://github.com/paulbellamy/mango) | 372 | 40 | 9 | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. | 2011-05-25 07:26:46 | 2023-04-17 06:21:21 |
| [gondola](http://gondolaweb.com) | 311 | 26 | 8 | The web framework for writing faster sites, faster | 2014-07-25 21:28:55 | 2023-09-25 21:25:20 |
| [uadmin](https://uadmin.io) | 303 | 59 | 19 | The web framework for Golang | 2018-10-05 09:00:17 | 2023-11-04 19:24:28 |
| [webgo](https://github.com/bnkamalesh/webgo) | 291 | 29 | 2 | A microframework to build web apps; with handler chaining, middleware support, and most of all; standard library compliant HTTP handlers(i.e. http.HandlerFunc). | 2015-12-16 07:35:02 | 2023-11-07 22:57:30 |
| [ginrpc](https://xxjwxc.github.io/post/ginrpc/) | 283 | 34 | 11 | gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具 | 2019-06-22 12:03:53 | 2023-11-09 06:37:04 |
| [golf](https://golf.readme.io/) | 267 | 30 | 6 | :golf: The Golf web framework | 2015-11-18 15:10:14 | 2023-10-23 11:11:55 |
| [hiboot](https://hiboot.netlify.app/) | 179 | 26 | 6 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2023-10-07 13:11:29 |
| [appy](https://github.com/appist/appy) | 130 | 14 | 3 | An opinionated productive web framework that helps scaling business easier. | 2019-05-27 04:48:59 | 2023-10-23 16:48:39 |
| [go-rest](http://go.pkgdoc.org/github.com/ungerik/go-rest) | 127 | 16 | 2 | A small and evil REST framework for Go | 2012-07-13 10:02:15 | 2023-03-07 16:06:33 |
| [patron](https://github.com/beatlabs/patron) | 124 | 67 | 26 | Microservice framework following best cloud practices with a focus on productivity. | 2019-01-30 13:49:54 | 2023-07-24 11:12:32 |
| [microservice](https://github.com/claygod/microservice) | 113 | 15 | 0 | This library provides a simple microservice framework based on clean architecture principles with a working example implemented. | 2016-12-15 09:07:04 | 2023-09-02 09:54:41 |
| [rux](https://pkg.go.dev/github.com/gookit/rux?tab=doc) | 93 | 15 | 2 | ⚡ Rux is an simple and fast web framework. support route group, param route binding, middleware, compatible http.Handler interface. 简单且快速的 Go api/web 框架，支持路由分组，路由参数绑定，中间件，兼容 http.Handler 接口 | 2018-08-05 06:13:57 | 2023-10-20 21:08:01 |
| [vox](https://aisk.github.io/vox/) | 80 | 7 | 8 | Simple and lightweight Go web framework inspired by koa | 2014-12-24 11:22:08 | 2023-11-07 12:38:50 |
| [golax](https://github.com/fulldump/golax) | 75 | 7 | 6 | Golax, a go implementation for the Lax framework. | 2016-01-30 19:11:39 | 2023-09-27 21:43:32 |
| [yarf](https://github.com/yarf-framework/yarf) | 67 | 8 | 2 | Yet Another REST Framework | 2015-09-02 13:56:47 | 2023-09-29 12:22:35 |
| [fireball](https://github.com/ridgelines/fireball) | 59 | 6 | 1 | Go web framework with a natural feel | 2016-07-20 05:04:54 | 2022-11-09 19:23:48 |
| [goa](https://goa-go.github.io) | 49 | 5 | 0 | Goa is a  web framework based on middleware, like koa.js. | 2019-07-26 07:12:23 | 2023-06-06 18:44:28 |
| [gotuna](https://gotuna.netlify.app) | 44 | 22 | 1 | GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server. | 2021-04-08 14:08:08 | 2023-10-16 18:48:15 |
| [goweb](https://github.com/twharmon/goweb) | 35 | 6 | 1 | Lightweight web framework based on net/http. | 2019-05-07 21:04:43 | 2023-03-16 06:12:26 |
| [api](http://resoursea.com/) | 34 | 4 | 0 | A REST framework for quickly writing resource based services in Golang. | 2015-01-24 18:45:30 | 2023-10-31 20:04:47 |
| [rex](https://github.com/goanywhere/rex) | 33 | 3 | 0 | Pleasures for Web in Golang | 2014-10-16 02:26:18 | 2022-09-27 10:03:38 |
| [banjo](https://nsheremet.pw/banjo) | 21 | 7 | 4 | BANjO is a simple web framework written in Go (golang) | 2017-12-09 13:35:31 | 2023-03-16 01:57:53 |
</details>

### Web Frameworks - Middlewares - Actual middlewares


<sup>*Last Update: 2023-12-16 21:01:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tollbooth](https://github.com/didip/tollbooth) | 2,532 | 209 | 7 | Simple middleware to rate-limit HTTP requests. | 2015-05-17 15:20:03 | 2023-12-16 13:01:26 |
| [cors](https://github.com/rs/cors) | 2,468 | 216 | 15 | Go net/http configurable handler to handle CORS requests | 2014-10-25 03:49:45 | 2023-12-14 18:15:26 |
| [limiter](https://lingrino.com) | 1,874 | 137 | 16 | Dead simple rate limit middleware for Go. | 2015-10-02 08:12:38 | 2023-12-14 06:52:15 |
| [go-server-timing](https://lingrino.com) | 852 | 38 | 10 | Go (golang) library for creating and consuming HTTP Server-Timing headers | 2018-02-12 03:56:02 | 2023-12-12 12:27:47 |
| [go-fault](https://lingrino.com) | 481 | 25 | 0 | fault injection library in go using standard http middleware | 2020-05-14 16:13:17 | 2023-12-02 04:29:31 |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 137 | 9 | 18 | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ | 2018-06-29 21:51:00 | 2023-10-17 11:42:49 |
| [xff](https://github.com/sebest/xff) | 98 | 25 | 8 | A Golang Middleware to handle X-Forwarded-For Header | 2014-12-22 10:29:05 | 2023-11-16 19:18:37 |
| [formjson](https://github.com/rs/formjson) | 38 | 3 | 0 | Go net/http handler to transparently manage posted JSON | 2015-03-19 23:52:28 | 2022-09-27 10:05:36 |
| [client-timing](https://github.com/posener/client-timing) | 24 | 6 | 1 | An HTTP client for go-server-timing middleware. Enables automatic timing propagation through HTTP calls between servers. | 2018-02-23 01:52:45 | 2023-06-02 14:22:43 |
</details>

### Web Frameworks - Middlewares - Libraries for creating HTTP middlewares


<sup>*Last Update: 2023-11-09 08:25:30*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [negroni](https://github.com/urfave/negroni) | 7,369 | 592 | 10 | Idiomatic HTTP Middleware for Golang | 2014-05-18 22:09:10 | 2023-11-07 13:52:33 |
| [alice](https://godoc.org/github.com/justinas/alice) | 2,884 | 185 | 8 | Painless middleware chaining for Go | 2014-05-25 07:27:41 | 2023-11-07 15:53:13 |
| [render](https://github.com/unrolled/render) | 1,853 | 144 | 0 | Go package for easily rendering JSON, XML, binary data, and HTML templates responses. | 2014-06-10 16:20:35 | 2023-11-08 03:37:27 |
| [stats](https://github.com/thoas/stats) | 593 | 53 | 9 | A Go middleware that stores various information about your web application (response time, status code count, etc.) | 2015-03-05 18:02:50 | 2023-08-01 19:47:46 |
| [interpose](https://github.com/carbocation/interpose) | 293 | 18 | 1 | Minimalist net/http middleware for golang | 2014-07-20 00:19:52 | 2023-08-21 11:24:56 |
| [renderer](https://github.com/thedevsaddam/renderer) | 257 | 30 | 0 | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go | 2017-11-07 18:53:49 | 2023-11-07 12:47:07 |
| [muxchain](https://github.com/stephens2424/muxchain) | 208 | 14 | 1 | Lightweight Middleware for net/http | 2014-05-03 17:14:17 | 2023-05-05 11:26:07 |
| [rye](https://github.com/InVisionApp/rye) | 102 | 15 | 0 | A tiny http middleware for Golang with added handlers for common needs. | 2016-10-06 19:51:59 | 2023-11-06 21:17:05 |
| [gores](https://github.com/alioygur/gores) | 100 | 4 | 0 | Go package that handles HTML, JSON, XML and etc. responses | 2015-12-25 12:41:01 | 2023-02-28 15:29:17 |
| [mediary](https://github.com/HereMobilityDevelopers/mediary/wiki/Reasoning) | 88 | 7 | 0 | Add interceptors to GO http.Client | 2020-03-23 18:54:56 | 2023-10-08 13:38:53 |
| [chain](https://github.com/codemodus/chain) | 63 | 3 | 0 | Composable chains of nested http.Handler instances. | 2015-05-14 19:52:58 | 2023-09-16 01:41:09 |
| [wrap](https://github.com/go-on/wrap) | 59 | 5 | 0 | Go http.Hander based middleware stack with context sharing | 2014-02-16 07:12:36 | 2020-08-28 05:29:07 |
| [catena](https://github.com/codemodus/catena) | 9 | 2 | 0 | gRPC interceptor catenation. | 2015-07-30 19:07:01 | 2022-09-27 10:06:05 |
</details>

### Web Frameworks - Routers


<sup>*Last Update: 2023-11-13 21:48:50*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mux](https://gorilla.github.io) | 19,458 | 1,820 | 19 | Package gorilla/mux is a powerful HTTP router and URL matcher for building Go web servers with 🦍 | 2012-10-02 21:32:24 | 2023-11-13 08:21:13 |
| [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) | 15,899 | 1,468 | 87 | A high performance HTTP request router that scales well | 2013-12-05 15:10:55 | 2023-11-13 05:52:20 |
| [chi](https://go-chi.io) | 15,798 | 972 | 59 | lightweight, idiomatic and composable router for building Go HTTP services | 2015-10-15 20:46:29 | 2023-11-13 11:58:41 |
| [web](https://github.com/gocraft/web) | 1,502 | 163 | 24 | Go Router + Middleware. Your Contexts. | 2013-11-16 20:48:20 | 2023-11-03 08:27:31 |
| [bone](http://go-zoo.github.io/bone) | 1,294 | 90 | 3 | Lightning Fast HTTP Multiplexer | 2014-11-19 02:16:36 | 2023-11-04 12:26:12 |
| [goji](https://goji.io) | 948 | 72 | 6 | Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) | 2015-11-16 00:52:41 | 2023-10-26 10:46:44 |
| [fasthttprouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 875 | 93 | 19 | A high performance fasthttp request router that scales well | 2015-12-13 09:32:30 | 2023-09-23 20:29:33 |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 609 | 56 | 7 | High-speed, flexible tree-based HTTP router for Go. | 2014-05-14 20:10:20 | 2023-10-09 07:03:54 |
| [gorouter](https://xujiajun.cn/gorouter) | 532 | 87 | 0 | xujiajun/gorouter is a simple and fast HTTP router for Go. It is easy to build RESTful APIs and your web framework. | 2018-01-29 09:28:28 | 2023-08-25 19:53:08 |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 445 | 52 | 11 | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. | 2015-10-27 01:03:14 | 2023-09-05 12:52:58 |
| [lars](https://github.com/go-playground/lars) | 391 | 26 | 1 | :rotating_light: Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. | 2015-12-24 17:28:45 | 2023-10-12 18:24:58 |
| [siesta](https://github.com/VividCortex/siesta) | 351 | 15 | 0 | Composable framework for writing HTTP handlers in Go. | 2014-09-23 13:55:56 | 2023-10-05 13:36:02 |
| [vestigo](https://github.com/husobee/vestigo) | 269 | 31 | 13 | Echo Inspired Stand Alone URL Router | 2015-09-22 03:08:03 | 2023-06-13 00:23:19 |
| [router](https://github.com/gowww/router) | 229 | 14 | 0 | ⚡️ A lightning fast HTTP router | 2017-05-25 10:29:27 | 2023-11-13 09:44:45 |
| [gorouter](https://rafallorenz.com/gorouter) | 150 | 16 | 10 | Go Server/API micro framework, HTTP request router, multiplexer, mux | 2016-07-14 13:13:34 | 2023-09-15 16:49:01 |
| [pure](https://github.com/go-playground/pure) | 144 | 13 | 0 | :non-potable_water: Is a lightweight  HTTP router that sticks to the std "net/http" implementation | 2016-09-23 19:57:58 | 2023-10-03 09:51:54 |
| [alien](https://github.com/gernest/alien) | 127 | 14 | 3 | A lightweight and  fast http router from outer space | 2016-01-30 23:23:10 | 2023-09-28 10:32:23 |
| [violetear](http://violetear.org) | 107 | 11 | 1 | Go HTTP router | 2015-06-19 16:49:41 | 2023-09-25 21:28:11 |
| [Bxog](http://go-zoo.github.io/bone) | 104 | 8 | 0 | Bxog is a simple and fast HTTP router for Go (HTTP request multiplexer). | 2016-05-19 12:20:08 | 2023-06-28 11:21:56 |
| [xmux](http://violetear.org) | 98 | 11 | 2 | xmux is a httprouter fork on top of xhandler (net/context aware) | 2015-12-14 19:01:05 | 2023-10-31 20:04:33 |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 54 | 7 | 0 | :bell: A simple Go router | 2019-02-21 13:13:52 | 2023-02-20 06:49:25 |
| [fastrouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 22 | 5 | 0 | FastRouter is a fast, flexible HTTP router written in Go. | 2017-11-01 08:52:52 | 2022-09-27 10:08:01 |
| [route](https://goroute.github.io) | 9 | 2 | 1 | Go Route - Simple yet powerful HTTP request multiplexer | 2019-07-06 18:47:38 | 2023-11-07 06:35:08 |
</details>

### WebAssembly


<sup>*Last Update: 2023-11-22 20:35:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tinygo](https://tinygo.org) | 13,696 | 805 | 490 | Go compiler for small places. Microcontrollers, WebAssembly (WASM/WASI), and command-line tools. Based on LLVM. | 2018-06-07 16:39:19 | 2023-11-22 09:59:27 |
| [dom](https://github.com/dennwc/dom) | 481 | 59 | 11 | DOM library for Go and WASM | 2018-06-30 18:37:35 | 2023-10-31 12:58:11 |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 217 | 17 | 4 | Library to use HTML5 Canvas  from Go-WASM, with all drawing within go code | 2019-05-05 14:05:55 | 2023-11-08 16:11:30 |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 159 | 25 | 4 | Run WASM tests inside your browser | 2018-07-14 18:42:24 | 2023-11-09 01:12:43 |
| [webapi](https://github.com/gowebapi/webapi) | 147 | 12 | 3 | Go Lang Web Assembly bindings for DOM, HTML etc | 2019-02-08 05:58:35 | 2023-11-17 02:15:42 |
| [vert](https://github.com/norunners/vert) | 93 | 13 | 2 | WebAssembly interop between Go and JS values. | 2018-03-25 17:26:47 | 2023-11-08 15:09:24 |
</details>

### Windows


<sup>*Last Update: 2023-11-20 08:20:34*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-ole](https://github.com/go-ole/go-ole) | 1,038 | 227 | 72 | win32 ole implementation for golang | 2011-01-21 12:45:20 | 2023-11-17 15:20:57 |
| [d3d9](https://github.com/gonutz/d3d9) | 150 | 13 | 1 | Direct3D9 wrapper for Go. | 2015-12-12 21:24:38 | 2023-11-01 09:38:18 |
| [gosddl](https://github.com/MonaxGT/gosddl) | 10 | 2 | 0 | GoSDDL converter | 2018-12-04 08:36:11 | 2023-04-23 09:37:32 |
</details>

### XML
Libraries and tools for manipulating XML.

<sup>*Last Update: 2023-11-28 21:14:53*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [zek](https://github.com/miku/zek) | 698 | 62 | 8 | Generate a Go struct from XML. | 2017-11-23 19:03:11 | 2023-11-27 12:49:14 |
| [xpath](https://github.com/antchfx/xpath) | 612 | 80 | 15 | XPath package for Golang, supports HTML, XML, JSON document query. | 2016-10-09 05:51:24 | 2023-11-28 12:59:52 |
| [xquery](https://github.com/antchfx/xpath) | 157 | 28 | 0 | Extract data or evaluate value from HTML/XML documents using XPath | 2016-10-09 05:54:10 | 2023-11-24 16:12:26 |
| [xml2map](https://github.com/sbabiv/xml2map) | 58 | 12 | 2 | XML to MAP converter written Golang | 2018-08-06 17:51:46 | 2023-11-04 03:48:24 |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 25 | 6 | 1 | xmlwriter is a pure-Go library providing procedural XML generation based on libxml2's xmlwriter module | 2017-04-11 04:43:26 | 2023-11-04 03:48:46 |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 21 | 11 | 8 | Compare ANY markup documents. | 2016-10-25 22:09:12 | 2023-11-04 03:47:00 |
</details>

