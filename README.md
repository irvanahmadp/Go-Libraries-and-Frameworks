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

<sup>*Last Update: 2021-10-11 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oto](https://github.com/hajimehoshi/oto) | 868 | 73 | 17 | ♪ A low-level library to play sound on multiple platforms ♪ | 2017-05-04 12:16:30 | 2021-10-09 10:18:21 |
| [portaudio](https://github.com/gordonklaus/portaudio) | 466 | 70 | 11 | Go bindings for the PortAudio audio I/O library | 2015-09-16 07:59:23 | 2021-10-06 19:04:16 |
| [music-theory](https://gopkg.in/music-theory.v0) | 361 | 35 | 8 | Go models of Note, Scale, Chord and Key | 2016-03-17 03:50:17 | 2021-09-22 22:32:45 |
| [waveform](https://github.com/mdlayher/waveform) | 343 | 26 | 4 | Go package capable of generating waveform images from audio streams. MIT Licensed. | 2014-09-13 18:07:36 | 2021-09-30 14:11:42 |
| [portmidi](https://github.com/rakyll/portmidi) | 262 | 50 | 12 | Go bindings for libportmidi | 2013-11-10 14:24:56 | 2021-08-13 16:24:03 |
| [id3v2](https://pkg.go.dev/github.com/bogem/id3v2) | 194 | 31 | 14 | 🎵 ID3 decoding and encoding library for Go | 2016-05-15 18:36:53 | 2021-10-08 01:42:53 |
| [flac](https://github.com/mewkiz/flac) | 164 | 30 | 10 | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 2012-11-01 20:13:58 | 2021-09-19 17:33:05 |
| [malgo](https://pkg.go.dev/github.com/bogem/id3v2) | 153 | 25 | 4 | Mini audio library | 2017-11-09 18:27:52 | 2021-09-29 07:05:28 |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 150 | 11 | 4 | Go tools for audio processing & creation 🎶 | 2020-07-05 01:35:41 | 2021-10-09 00:55:07 |
| [mix](https://gopkg.in/mix.v0) | 144 | 21 | 11 | Sequence-based Go-native audio mixer for music apps | 2016-01-03 15:53:57 | 2021-09-30 17:10:25 |
| [gaad](https://github.com/Comcast/gaad) | 86 | 15 | 3 | GAAD (Go Advanced Audio Decoder) | 2016-07-11 14:19:16 | 2021-10-01 22:31:56 |
| [minimp3](https://github.com/tosone/minimp3) | 63 | 10 | 0 | Decode mp3 base on https://github.com/lieff/minimp3 | 2018-01-26 09:10:31 | 2021-09-28 07:30:01 |
| [vorbis](https://github.com/mccoyst/vorbis) | 28 | 4 | 2 | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) | 2013-07-12 02:45:39 | 2020-10-04 21:02:12 |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 12 | 7 | 0 | Go Bindings for libsamplerate | 2016-11-20 21:19:31 | 2021-09-29 05:36:31 |
</details>

### Authentication and OAuth
Libraries for implementing authentications schemes.

<sup>*Last Update: 2021-11-25 13:59:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [casbin](https://casbin.org/) | 10,720 | 1,188 | 31 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2017-04-08 07:51:23 | 2021-11-25 06:37:09 |
| [oauth2](https://golang.org/x/oauth2) | 3,908 | 797 | 150 | Go OAuth2 | 2014-04-14 15:07:35 | 2021-11-24 03:01:22 |
| [goth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 3,436 | 415 | 71 | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 2014-10-14 20:38:12 | 2021-11-25 06:38:11 |
| [authboss](https://github.com/volatiletech/authboss) | 2,911 | 178 | 30 | The boss of http auth. | 2015-01-03 05:12:02 | 2021-11-24 21:11:47 |
| [go-jose](https://github.com/square/go-jose) | 1,838 | 327 | 56 | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 2014-11-14 18:27:31 | 2021-11-24 21:20:15 |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1,833 | 282 | 28 | A standalone, specification-compliant,  OAuth2 server written in Golang. | 2015-11-01 13:30:09 | 2021-11-23 09:31:40 |
| [loginsrv](https://github.com/tarent/loginsrv) | 1,827 | 157 | 29 | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 2016-11-11 12:11:21 | 2021-11-20 19:16:04 |
| [osin](https://golang.org/x/oauth2) | 1,701 | 381 | 3 | Golang OAuth2 server library | 2013-09-10 19:52:00 | 2021-11-23 02:48:54 |
| [gologin](https://github.com/dghubble/gologin) | 1,455 | 114 | 0 | Go login handlers for authentication providers (OAuth1, OAuth2) | 2015-06-23 04:40:52 | 2021-11-23 14:42:57 |
| [gorbac](https://github.com/mikespook/gorbac) | 1,224 | 148 | 2 | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 2013-12-26 10:00:41 | 2021-11-17 13:06:54 |
| [scs](https://github.com/alexedwards/scs) | 971 | 94 | 7 | HTTP Session Management for Go | 2016-08-08 16:42:05 | 2021-11-24 18:56:24 |
| [paseto](https://github.com/o1egl/paseto) | 555 | 26 | 4 | Platform-Agnostic Security Tokens implementation in GO (Golang) | 2018-01-23 05:27:39 | 2021-11-17 19:02:19 |
| [permissions2](https://github.com/xyproto/permissions2) | 438 | 35 | 0 |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 2014-11-19 12:23:37 | 2021-10-29 16:42:42 |
| [go-guardian](https://github.com/shaj13/go-guardian) | 319 | 27 | 7 | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 2020-05-14 12:15:56 | 2021-11-24 22:01:01 |
| [jwt](https://github.com/cristalhq/jwt) | 304 | 25 | 0 | Safe, simple and fast JSON Web Tokens for Go | 2019-07-20 18:14:58 | 2021-11-24 16:07:53 |
| [jwt](https://github.com/pascaldekloe/jwt) | 272 | 20 | 0 | JSON Web Token library | 2018-03-21 11:59:53 | 2021-11-17 12:11:17 |
| [jeff](https://github.com/abraithwaite/jeff) | 234 | 13 | 1 | 🍍Jeff provides the simplest way to manage web sessions in Go. | 2018-08-02 19:31:23 | 2021-11-22 01:44:42 |
| [httpauth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 210 | 26 | 5 | HTTP Authentication middlewares | 2014-05-26 22:53:57 | 2021-10-12 21:06:21 |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 209 | 38 | 3 | This package provides json web token (jwt) middleware for goLang http servers | 2016-07-05 23:31:43 | 2021-11-24 14:10:45 |
| [branca](https://branca.io) | 161 | 17 | 1 | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. | 2018-01-09 15:27:31 | 2021-11-16 13:49:51 |
| [sessionup](https://github.com/swithek/sessionup) | 112 | 5 | 3 | Straightforward HTTP session management | 2019-07-23 18:55:21 | 2021-10-14 13:05:30 |
| [session](https://github.com/icza/session) | 105 | 15 | 4 | Go session management for web servers (including support for Google App Engine - GAE). | 2016-02-08 09:07:07 | 2021-10-14 13:05:21 |
| [jwt](https://github.com/robbert229/jwt) | 94 | 24 | 9 | This is an implementation of JWT in golang! | 2016-06-05 22:01:37 | 2021-11-17 13:12:30 |
| [sjwt](https://godoc.org/github.com/brianvoe/sjwt) | 94 | 7 | 0 | Simple JWT Golang | 2019-06-20 04:06:21 | 2021-09-22 11:16:33 |
| [rbac](https://github.com/zpatrick/rbac) | 87 | 15 | 0 | Minimalistic RBAC package for Go applications | 2018-08-02 00:11:04 | 2021-11-22 13:28:00 |
| [sessions](https://github.com/adam-hanna/sessions) | 60 | 7 | 2 | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 2017-04-29 01:09:28 | 2021-07-03 11:27:07 |
| [securecookie](https://github.com/chmike/securecookie) | 55 | 8 | 0 | Fast, secure and efficient secure cookie encoder/decoder  | 2017-09-03 14:40:29 | 2021-11-15 04:09:21 |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 33 | 1 | 0 | Golang library for providing a canonical representation of email address. | 2020-08-21 23:13:04 | 2021-11-15 04:07:52 |
| [otpgo](https://github.com/jltorresm/otpgo) | 26 | 2 | 1 | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 2020-08-19 13:20:23 | 2021-11-12 11:56:20 |
| [scope](https://github.com/SonicRoshan/scope) | 17 | 4 | 0 | Easily Manage OAuth2 Scopes In Go | 2019-09-23 10:48:14 | 2021-11-06 23:03:46 |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 1 | 0 | A driver for the SessionGate Redis module - easy session management using the Go language. | 2017-10-20 03:39:11 | 2020-08-18 23:01:11 |
| [cookiestxt](https://casbin.org/) | 8 | 2 | 1 | cookiestxt implement parser of cookies txt format | 2017-10-09 11:27:19 | 2021-08-25 11:03:38 |
| [casbin](https://casbin.org/) | 1 | 0 | 0 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2021-05-29 04:09:46 | 2021-06-23 04:53:51 |
</details>

### Bot Building
Libraries for building and working with bots.

<sup>*Last Update: 2023-09-30 07:16:18*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [telegram-bot-api](https://go-telegram-bot-api.dev) | 4,967 | 757 | 112 | Golang bindings for the Telegram Bot API | 2015-06-25 05:33:57 | 2023-09-29 19:50:28 |
| [olivia](https://olivia-ai.org) | 3,592 | 354 | 28 | 💁‍♀️Your new best friend powered by an artificial neural network | 2018-06-05 18:19:31 | 2023-09-28 10:51:17 |
| [telebot](https://github.com/tucnak/telebot) | 3,261 | 393 | 45 | Telebot is a Telegram bot framework in Go. | 2015-06-25 19:27:50 | 2023-09-28 10:29:19 |
| [kelp](https://kelpbot.io) | 1,021 | 239 | 176 | Kelp is a free and open-source trading bot for the Stellar DEX and 100+ centralized exchanges | 2018-08-08 23:31:18 | 2023-09-24 15:12:40 |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 938 | 239 | 6 | A golang implementation of a console-based trading bot for cryptocurrency exchanges | 2017-05-14 22:11:41 | 2023-09-25 08:06:00 |
| [bot](https://github.com/go-chat-bot/bot) | 799 | 212 | 15 | IRC, Slack, Telegram and RocketChat bot written in go | 2015-09-22 16:41:13 | 2023-09-26 14:27:27 |
| [slacker](https://github.com/shomali11/slacker) | 798 | 120 | 2 | Slack Bot Framework | 2017-05-20 01:41:20 | 2023-09-28 10:41:35 |
| [joe](https://joe-bot.net) | 468 | 26 | 6 | A general-purpose bot library inspired by Hubot but written in Go.   :robot: | 2019-03-03 11:19:18 | 2023-08-06 14:21:19 |
| [tbot](https://yanzay.github.io/tbot-doc/) | 337 | 52 | 1 | Go library for Telegram Bot API | 2015-09-11 16:19:25 | 2023-09-18 15:27:32 |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 307 | 56 | 7 | go irc client for twitch.tv | 2017-03-23 21:31:35 | 2023-09-24 12:56:12 |
| [echotron](https://t.me/s/echotronnews) | 270 | 21 | 1 | An elegant and concurrent library for the Telegram bot API in Go. | 2019-07-22 17:31:49 | 2023-09-28 11:04:21 |
| [go-sarah](https://pkg.go.dev/github.com/oklahomer/go-sarah/v4) | 261 | 17 | 0 | Simple yet customizable bot framework written in Go. | 2016-11-06 10:04:43 | 2023-09-21 06:31:58 |
| [tenyks](http://tenyks.io) | 176 | 19 | 12 | The Tenyks IRC bot. | 2012-08-26 02:02:24 | 2023-09-28 17:37:33 |
| [hanu](https://sbstjn.com/host-golang-slackbot-on-heroku-with-hanu.html) | 149 | 24 | 3 | Golang Framework for writing Slack bots | 2016-09-16 07:10:42 | 2023-09-18 15:27:21 |
| [slack-bot](https://github.com/innogames/slack-bot) | 146 | 44 | 6 | Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests with AI support... | 2019-07-19 07:49:06 | 2023-09-24 04:07:51 |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 120 | 4 | 2 | Golang  telegram bot API wrapper, session-based router and middleware | 2016-12-11 06:06:32 | 2023-09-18 15:27:18 |
| [margelet](https://github.com/zhulik/margelet) | 81 | 15 | 2 | Telegram Bot Framework for Go | 2015-11-21 13:02:17 | 2023-09-18 15:27:27 |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 81 | 10 | 8 | A Discord bot for managing ephemeral roles based upon voice channel member presence. | 2017-12-19 15:20:30 | 2023-09-25 20:00:44 |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 56 | 11 | 3 | Slack bot core/framework written in Go with support for reactions to message updates/deletes | 2015-10-22 04:54:55 | 2023-09-18 15:27:30 |
| [govkbot](https://github.com/nikepan/govkbot) | 45 | 6 | 1 | VK bot package for Go | 2016-07-11 22:09:54 | 2023-09-18 15:27:20 |
| [micha](https://github.com/onrik/micha) | 26 | 5 | 1 | Client lib for Telegram bot api | 2016-04-14 12:09:44 | 2023-09-18 15:27:25 |
</details>

### Build Automation
Libraries and tools helping with build automation.

<sup>*Last Update: 2021-09-27 09:25:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [realize](https://github.com/oxequa/realize) | 4,114 | 213 | 68 | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 2016-07-12 08:07:25 | 2021-09-26 11:58:21 |
| [task](https://taskfile.dev) | 3,888 | 252 | 96 | A task runner / simpler Make alternative written in Go | 2017-02-27 00:46:04 | 2021-09-27 01:30:49 |
| [mmake](https://github.com/tj/mmake) | 1,580 | 43 | 11 | Modern Make  | 2017-02-15 22:01:21 | 2021-09-24 10:43:22 |
| [goyek](https://github.com/goyek/goyek) | 257 | 18 | 3 | Create build pipelines in Go  | 2020-10-11 13:20:55 | 2021-09-25 04:25:00 |
| [taskctl](https://github.com/taskctl/taskctl) | 131 | 13 | 7 | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 2019-11-12 13:19:09 | 2021-09-26 15:36:21 |
| [1build](https://1build.gitbook.io) | 98 | 27 | 31 | Frictionless way of managing project-specific commands | 2019-04-23 17:05:38 | 2021-09-26 16:17:35 |
| [gilbert](https://go-gilbert.github.io/) | 93 | 5 | 0 | Build system and task runner for Go projects | 2019-01-30 09:02:31 | 2021-08-14 01:54:30 |
| [gaper](https://github.com/maxcnunes/gaper) | 48 | 3 | 7 | Builds and restarts a Go project when it crashes or some watched file changes | 2018-06-16 02:46:38 | 2021-08-02 08:33:20 |
| [anko](https://github.com/GuilhermeCaruso/anko) | 18 | 0 | 0 | :crystal_ball: Simple application watcher | 2021-03-02 14:08:42 | 2021-09-03 04:36:40 |
</details>

### CSS Preprocessors
Libraries for preprocessing CSS files.

<sup>*Last Update: 2021-07-31 09:25:23*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gcss](https://github.com/yosssi/gcss) | 445 | 35 | 8 | Pure Go CSS Preprocessor | 2014-09-04 14:38:20 | 2021-07-26 07:08:07 |
| [go-libsass](http://godoc.org/github.com/wellington/go-libsass) | 179 | 23 | 13 | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 2015-04-19 15:09:47 | 2021-06-23 15:09:17 |
</details>

### Command Line - Advanced Console UIs
Libraries for building Console Applications and Console User Interfaces.

<sup>*Last Update: 2023-09-30 09:08:34*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [termui](https://github.com/gizak/termui) | 12,660 | 781 | 100 | Golang terminal dashboard | 2015-02-03 14:09:27 | 2023-09-28 01:39:18 |
| [gocui](https://github.com/jroimartin/gocui) | 9,399 | 633 | 74 | Minimalist Go package aimed at creating Console User Interfaces. | 2014-01-04 02:50:20 | 2023-09-28 23:28:34 |
| [go-prompt](https://godoc.org/github.com/c-bata/go-prompt) | 5,014 | 369 | 108 | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. | 2017-08-14 16:02:09 | 2023-09-29 19:51:59 |
| [termbox-go](http://godoc.org/github.com/nsf/termbox-go) | 4,573 | 399 | 45 | Pure Go termbox implementation | 2012-01-12 21:03:03 | 2023-09-28 10:19:46 |
| [pterm](https://docs.pterm.sh) | 4,191 | 164 | 48 | ✨ #PTerm is a modern Go module to easily beautify console output. Featuring charts, progressbars, tables, trees, text input, select menus and much more 🚀 It's completely configurable and 100% cross-platform compatible. | 2020-09-17 15:52:59 | 2023-09-29 23:55:29 |
| [progressbar](https://pkg.go.dev/github.com/schollz/progressbar/v3?tab=doc) | 3,497 | 200 | 33 | A really basic thread-safe progress bar for Golang applications | 2017-10-26 18:28:10 | 2023-09-29 11:56:52 |
| [termdash](http://godoc.org/github.com/nsf/termbox-go) | 2,366 | 118 | 43 | Terminal based dashboard. | 2018-03-24 12:01:49 | 2023-09-28 10:49:16 |
| [asciigraph](https://pkg.go.dev/github.com/guptarohit/asciigraph) | 2,317 | 93 | 12 | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. | 2018-06-17 10:37:16 | 2023-09-28 19:53:13 |
| [mpb](https://github.com/vbauerster/mpb) | 2,097 | 123 | 8 | multi progress bar for Go cli applications | 2016-12-14 11:56:29 | 2023-09-28 10:37:58 |
| [uiprogress](https://github.com/gosuri/uiprogress) | 2,044 | 125 | 25 | A go library to render progress bars in terminal applications | 2015-11-17 00:59:24 | 2023-09-26 16:39:25 |
| [uilive](https://github.com/gosuri/uilive) | 1,601 | 87 | 11 | uilive is a go library for updating terminal output in realtime | 2015-11-16 06:13:10 | 2023-09-26 20:39:42 |
| [color](https://gookit.github.io/color/) | 1,374 | 84 | 6 | 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染 | 2018-07-01 07:28:17 | 2023-09-27 02:04:06 |
| [aurora](https://pkg.go.dev/github.com/guptarohit/asciigraph) | 1,347 | 61 | 1 | Golang ultimate ANSI-colors that supports Printf/Sprintf methods | 2016-11-06 22:37:12 | 2023-09-26 14:04:15 |
| [go-isatty](http://godoc.org/github.com/mattn/go-isatty) | 743 | 105 | 14 | Change the color of console text. | 2014-04-01 01:53:09 | 2023-09-27 01:28:07 |
| [go-colorable](http://godoc.org/github.com/mattn/go-colorable) | 713 | 94 | 9 | Another Text Attribute Manupulator | 2014-07-30 02:38:06 | 2023-09-27 02:03:22 |
| [uitable](https://github.com/gosuri/uitable) | 713 | 33 | 8 | A go library to improve readability in terminal apps using tabular data | 2015-11-13 21:59:21 | 2023-09-26 16:39:12 |
| [gommon](https://github.com/labstack/gommon) | 505 | 120 | 14 | Common packages for Go | 2015-03-12 22:35:57 | 2023-09-26 02:52:27 |
| [simpletable](https://github.com/alexeyco/simpletable) | 467 | 30 | 5 | Simple tables in terminal with Go | 2017-03-29 07:27:23 | 2023-09-23 14:56:31 |
| [chalk](https://github.com/ttacon/chalk) | 430 | 20 | 3 | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk | 2014-07-18 19:38:58 | 2023-09-27 02:02:54 |
| [yacspin](https://github.com/theckman/yacspin) | 413 | 13 | 6 | Yet Another CLi Spinner; providing over 80 easy to use and customizable terminal spinners for multiple OSes | 2019-12-29 07:41:23 | 2023-09-26 16:39:04 |
| [box-cli-maker](https://github.com/Delta456/box-cli-maker) | 386 | 19 | 8 | Make Highly Customized Boxes for CLI | 2020-05-01 07:23:56 | 2023-09-29 08:33:05 |
| [tabby](https://github.com/cheynewallace/tabby) | 335 | 18 | 3 | A tiny library for super simple Golang tables | 2018-12-17 23:35:39 | 2023-09-26 16:40:53 |
| [go-colortext](http://godoc.org/github.com/mattn/go-colorable) | 214 | 22 | 3 | Change the color of console text. | 2013-01-23 03:38:54 | 2023-07-30 06:16:45 |
| [cfmt](https://github.com/mingrammer/cfmt) | 96 | 7 | 1 | :art: Contextual fmt inspired by bootstrap color classes | 2018-03-15 19:04:27 | 2023-09-27 02:02:43 |
| [tabular](https://github.com/InVisionApp/tabular) | 73 | 7 | 0 | Tabular simplifies printing ASCII tables from command line utilities | 2018-04-23 21:17:03 | 2023-09-26 16:40:30 |
| [cfmt](https://github.com/i582/cfmt) | 61 | 3 | 0 | Small library for simple and convenient formatted stylized output to the console. | 2020-11-13 20:29:45 | 2023-09-19 07:48:09 |
| [table](https://github.com/tomlazar/table) | 46 | 3 | 1 | pretty colorfull tables in go with less effort | 2020-09-22 05:42:34 | 2023-09-26 16:40:38 |
| [ctc](https://github.com/wzshiming/ctc) | 45 | 3 | 1 | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method | 2018-04-27 18:07:42 | 2023-08-30 18:35:21 |
| [marker](https://github.com/cyucelen/marker) | 42 | 13 | 4 |  🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs! | 2019-08-28 15:44:08 | 2023-09-27 02:04:27 |
| [colourize](https://github.com/TreyBastian/colourize) | 26 | 5 | 0 | An ANSI colour terminal package for Go | 2015-05-11 11:49:39 | 2023-02-20 07:52:41 |
| [go-ataman](https://github.com/workanator/go-ataman) | 16 | 3 | 0 | Another Text Attribute Manupulator | 2017-05-17 19:04:57 | 2023-09-27 02:03:14 |
</details>

### Command Line - Standard CLI
Libraries for building standard or basic Command Line applications.

<sup>*Last Update: 2021-07-19 09:25:06*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cobra](https://cobra.dev) | 22,401 | 1,892 | 329 | A Commander for modern Go CLI interactions | 2013-09-03 20:40:26 | 2021-07-19 01:12:27 |
| [cli](https://github.com/urfave/cli) | 16,162 | 1,416 | 64 | A simple, fast, and fun package for building command line apps in Go | 2013-07-13 19:32:06 | 2021-07-19 02:17:54 |
| [kingpin](https://github.com/alecthomas/kingpin) | 3,076 | 234 | 25 | CONTRIBUTIONS ONLY: A Go (golang) command line and flag parser | 2014-05-14 20:09:04 | 2021-07-17 00:31:30 |
| [dnote](https://www.getdnote.com) | 2,107 | 88 | 58 | A simple command line notebook for programmers | 2017-03-30 23:07:25 | 2021-07-18 15:21:15 |
| [go-flags](http://godoc.org/github.com/jessevdk/go-flags) | 1,978 | 241 | 30 | go command line option parser | 2012-08-31 13:57:58 | 2021-07-18 10:59:59 |
| [pflag](https://ops.city) | 1,495 | 251 | 104 | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. | 2013-08-30 14:53:31 | 2021-07-17 15:22:47 |
| [cli](https://github.com/mitchellh/cli) | 1,356 | 104 | 7 | A Go library for implementing command-line interfaces. | 2013-11-03 06:47:54 | 2021-07-18 17:05:52 |
| [go-arg](https://github.com/alexflint/go-arg) | 1,193 | 67 | 7 | Struct-based argument parsing in Go | 2015-11-01 01:30:06 | 2021-07-17 02:26:33 |
| [liner](https://github.com/peterh/liner) | 801 | 104 | 12 | Pure Go line editor with history, inspired by linenoise | 2012-08-15 16:34:55 | 2021-07-18 03:32:44 |
| [complete](https://github.com/posener/complete) | 778 | 60 | 20 | bash completion written in go + bash completion for go command | 2017-05-05 21:34:07 | 2021-07-16 12:16:37 |
| [mow.cli](https://github.com/jawher/mow.cli) | 741 | 50 | 26 | A versatile library for building CLI applications in Go | 2014-12-18 19:34:20 | 2021-07-18 06:21:53 |
| [flaggy](https://github.com/integrii/flaggy) | 721 | 24 | 14 | Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies. | 2018-03-05 05:55:05 | 2021-07-14 22:39:41 |
| [ops](https://ops.city) | 635 | 67 | 133 | ops - build and run nanos unikernels | 2018-09-10 17:57:47 | 2021-07-19 00:12:21 |
| [cli](https://github.com/mkideal/cli) | 587 | 38 | 3 | CLI - A package for building command line app with go | 2016-02-26 16:45:29 | 2021-07-19 02:10:10 |
| [argparse](https://github.com/akamensky/argparse) | 347 | 42 | 10 | Argparse for golang. Just because `flag` sucks | 2017-11-24 06:42:20 | 2021-07-16 14:56:47 |
| [climax](http://git.io/climax) | 186 | 17 | 6 | Climax is an alternative CLI with the human face | 2015-11-03 21:04:57 | 2021-06-04 17:16:47 |
| [hiboot](https://hiboot.netlify.app/) | 162 | 28 | 5 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2021-06-26 11:52:55 |
| [commandeer](https://github.com/jaffee/commandeer) | 148 | 14 | 3 | Automatically sets up command line flags based on struct fields and tags. | 2017-10-12 02:51:05 | 2021-06-25 06:27:45 |
| [wmenu](https://github.com/dixonwille/wmenu) | 139 | 19 | 1 | An easy to use menu structure for cli applications that prompts users to make choices. | 2016-04-20 13:09:44 | 2021-07-16 08:49:25 |
| [sflags](https://ops.city) | 131 | 24 | 9 | Generate flags by parsing structures | 2016-12-04 14:49:27 | 2021-07-08 05:16:19 |
| [flag](https://github.com/cosiner/flag) | 115 | 6 | 1 | Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand | 2016-10-05 16:49:41 | 2021-05-29 20:22:59 |
| [clif](https://github.com/ukautz/clif) | 108 | 11 | 3 | Another CLI framework for Go. It works on my machine. | 2015-05-30 18:30:08 | 2021-05-25 03:55:56 |
| [job](https://github.com/liujianping/job) | 102 | 7 | 1 | JOB, make your short-term command as a long-term job. 将命令行规划成任务的工具 | 2019-04-09 11:14:51 | 2021-07-13 20:00:08 |
| [cli](https://github.com/teris-io/cli) | 89 | 9 | 2 | Simple and complete API for building command line applications in Go | 2017-05-24 23:07:07 | 2021-07-15 17:03:28 |
| [env](https://github.com/codingconcepts/env) | 80 | 6 | 1 | Tag-based environment configuration for structs | 2017-06-14 20:01:55 | 2021-07-01 17:05:17 |
| [cmdr](https://hedzr.github.io/cmdr-docs/) | 78 | 7 | 1 | Golang library with POSIX-compliant command-line UI (CLI) and Hierarchical-configuration. Better substitute for stdlib flag. | 2019-05-15 09:58:02 | 2021-07-14 22:02:01 |
| [clir](https://github.com/leaanthony/clir) | 74 | 9 | 5 | A Simple and Clear CLI library. Dependency free. | 2019-11-18 19:52:00 | 2021-07-16 07:19:29 |
| [gocmd](https://github.com/devfacet/gocmd) | 51 | 6 | 1 | A Go library for building command line applications. | 2018-01-08 04:52:02 | 2021-07-15 17:03:30 |
| [wlog](https://github.com/dixonwille/wlog) | 51 | 6 | 0 | A simple logging interface that supports cross-platform color and concurrency. | 2016-04-13 16:47:40 | 2021-04-25 04:09:14 |
| [strumt](https://github.com/antham/strumt) | 44 | 4 | 0 | Strumt is a library to create prompt chain | 2017-06-19 19:33:16 | 2021-07-06 08:25:25 |
| [flagvar](https://pkg.go.dev/github.com/sgreben/flagvar?tab=doc) | 36 | 1 | 0 | A collection of CLI argument types for the Go `flag` package. | 2018-05-18 18:45:16 | 2020-09-18 12:21:54 |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 36 | 5 | 0 | Fully featured Go (golang) command line option parser with built-in auto-completion support. | 2015-12-18 02:21:14 | 2021-07-06 13:39:54 |
| [cmd](https://github.com/posener/cmd) | 33 | 1 | 0 | The standard library flag package with its missing features | 2019-10-29 00:32:11 | 2021-06-11 12:03:34 |
| [argv](https://github.com/cosiner/argv) | 30 | 7 | 1 | Argparse for golang. Just because `flag` sucks | 2017-01-22 10:37:21 | 2021-03-30 02:32:52 |
| [go-commander](http://yitsushi.github.io/go-commander/) | 23 | 4 | 1 | Go library to simplify CLI workflow | 2016-10-10 10:09:41 | 2021-06-25 15:17:19 |
| [sand](https://ops.city) | 15 | 1 | 0 | Package for creating interpreters | 2018-11-18 22:44:41 | 2021-05-10 22:00:31 |
| [ts](https://github.com/liujianping/ts) | 13 | 1 | 0 | timestamp convert & compare tool. 时间戳转换与对比工具 | 2019-06-25 10:21:13 | 2021-02-13 09:18:33 |
</details>

### Configuration
Libraries for configuration parsing.

<sup>*Last Update: 2021-09-27 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [viper](https://github.com/spf13/viper) | 16,843 | 1,459 | 446 | Go configuration with fangs | 2014-04-02 14:33:33 | 2021-09-27 02:14:39 |
| [godotenv](http://godoc.org/github.com/joho/godotenv) | 4,161 | 246 | 52 | A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) | 2013-07-30 07:45:19 | 2021-09-26 14:27:57 |
| [envconfig](https://josh.blog/2017/04/go-configure) | 3,790 | 311 | 47 | Golang library for managing configuration data from environment variables | 2013-11-06 17:01:55 | 2021-09-26 20:24:59 |
| [ini](https://ini.unknwon.io) | 2,681 | 321 | 29 | Package ini provides INI file read and write functionality in Go | 2014-12-18 07:36:37 | 2021-09-26 19:38:14 |
| [env](http://carlosbecker.com/posts/env-structs-golang) | 2,051 | 147 | 1 | Simple lib to parse environment variables to structs | 2015-07-28 02:14:37 | 2021-09-24 13:27:48 |
| [konfig](https://github.com/lalamove/konfig) | 614 | 43 | 4 | Composable, observable and performant config handling for Go for the distributed processing era | 2019-01-18 17:03:03 | 2021-09-20 21:28:58 |
| [koanf](https://github.com/knadh/koanf) | 602 | 55 | 9 | Light weight, extensible, configuration management library for Go. Built in support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to viper. | 2019-06-18 06:34:05 | 2021-09-27 00:22:19 |
| [confita](https://github.com/heetch/confita) | 415 | 45 | 22 | Load configuration in cascade from multiple backends into a struct | 2017-12-21 10:49:18 | 2021-09-26 15:25:43 |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 367 | 37 | 16 | ✨Clean and minimalistic environment configuration reader for Golang | 2019-07-12 15:28:52 | 2021-09-25 16:02:32 |
| [config](https://pkg.go.dev/github.com/gookit/config/v2) | 281 | 32 | 1 | 📝 Go config manage(load,get,set). support JSON, YAML, TOML, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名 | 2018-07-07 08:11:39 | 2021-09-26 22:44:43 |
| [aconfig](https://github.com/cristalhq/aconfig) | 279 | 20 | 10 | Simple, useful and opinionated config loader. | 2020-06-26 19:43:20 | 2021-09-26 11:08:33 |
| [config](https://github.com/JeremyLoy/config) | 268 | 12 | 2 | 12 factor configuration as a typesafe struct in as little as two function calls | 2019-04-02 13:41:22 | 2021-09-23 04:20:16 |
| [store](https://github.com/tucnak/store) | 257 | 17 | 2 | A dead simple configuration manager for Go applications | 2015-10-03 19:17:28 | 2021-09-14 05:58:50 |
| [config](https://github.com/olebedev/config) | 240 | 43 | 4 | JSON or YAML configuration wrapper with convenient access methods. | 2014-04-21 15:09:39 | 2021-08-12 22:55:05 |
| [hjson-go](https://hjson.github.io/) | 240 | 29 | 3 | Hjson for Go | 2016-08-05 22:59:18 | 2021-09-02 16:44:50 |
| [envconfig](https://godoc.org/github.com/tomazk/envcfg) | 209 | 22 | 0 | Small library to read your configuration from environment variables | 2015-04-21 23:37:17 | 2021-09-04 23:20:06 |
| [config](https://josh.blog/2017/04/go-configure) | 206 | 12 | 1 | 🛠 A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP | 2017-04-02 18:37:05 | 2021-09-22 19:43:00 |
| [config](https://github.com/golobby/config) | 181 | 13 | 0 | A lightweight yet powerful configuration management for Go projects | 2019-10-15 22:51:19 | 2021-09-26 15:56:07 |
| [fig](https://github.com/kkyr/fig) | 166 | 11 | 2 | A minimalist Go configuration library | 2020-01-16 18:43:19 | 2021-09-26 20:40:26 |
| [gcfg](https://gopkg.in/gcfg.v1) | 155 | 49 | 9 | read INI-style configuration files into Go structs; supports user-defined types and subsections | 2015-08-17 14:40:55 | 2021-09-23 03:35:05 |
| [goconfig](https://pkg.go.dev/github.com/gosidekick/goconfig?tab=doc) | 149 | 21 | 6 | goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. | 2016-12-18 11:22:41 | 2021-09-02 15:09:48 |
| [xdg](https://pkg.go.dev/github.com/adrg/xdg) | 133 | 8 | 0 | Go implementation of the XDG Base Directory Specification and XDG user directories | 2014-08-22 08:23:40 | 2021-09-22 21:09:21 |
| [envh](https://github.com/antham/envh) | 95 | 1 | 0 | Go helpers to manage environment variables | 2017-01-12 11:25:48 | 2021-07-06 08:25:16 |
| [envcfg](https://godoc.org/github.com/tomazk/envcfg) | 94 | 6 | 0 | Un-marshaling environment variables to Go structs | 2014-11-29 11:43:53 | 2021-08-10 00:07:28 |
| [onion](https://github.com/goraz/onion) | 91 | 9 | 7 | Layer based configuration for golang | 2015-07-22 14:28:21 | 2021-08-22 16:51:15 |
| [harvester](https://github.com/beatlabs/harvester) | 86 | 19 | 3 | Harvest configuration, watch and notify subscriber | 2019-04-09 07:37:19 | 2021-09-06 07:43:23 |
| [configuro](https://medium.com/better-programming/designing-cloud-native-configuration-framework-eefb0b3793cb) | 76 | 8 | 0 | An opinionated configuration loading framework for Containerized and Cloud-Native applications. | 2020-04-09 22:10:34 | 2021-09-20 02:01:43 |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 63 | 5 | 1 | A cross platform package that follows the XDG Standard | 2017-07-20 15:58:29 | 2021-06-18 00:46:25 |
| [gofigure](https://github.com/ian-kent/gofigure) | 60 | 9 | 1 | Go configuration made easy! | 2014-11-25 00:12:40 | 2021-09-26 20:52:03 |
| [configure](https://github.com/paked/configure) | 54 | 9 | 2 | Configure is a Go package that gives you easy configuration of your project through redundancy | 2015-06-14 07:46:56 | 2021-05-27 13:28:12 |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 41 | 7 | 0 | Go package that interfaces with AWS System Manager | 2019-01-24 09:01:19 | 2021-08-13 20:45:03 |
| [configuration](https://github.com/BoRuDar/configuration) | 40 | 6 | 0 | Library for setting values to structs' fields from env, flags, files or default tag | 2019-11-27 17:58:49 | 2021-09-23 11:07:27 |
| [gone](https://github.com/One-com/gone) | 36 | 5 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2021-08-01 19:37:48 |
| [ingo](https://github.com/schachmat/ingo) | 35 | 7 | 0 | persistent storage for flags in go | 2016-02-07 22:57:40 | 2021-06-15 10:31:14 |
| [go-up](https://github.com/ufoscout/go-up) | 32 | 5 | 1 | go-up! A simple configuration library with recursive placeholders resolution and no magic. | 2018-02-18 09:50:00 | 2021-06-11 20:28:29 |
| [hocon](https://hjson.github.io/) | 29 | 6 | 2 | go implementation of lightbend's HOCON configuration library https://github.com/lightbend/config | 2020-03-01 18:20:12 | 2021-08-31 07:49:59 |
| [mini](https://github.com/sasbury/mini) | 28 | 6 | 1 | A golang package for parsing ini-style configuration files | 2015-04-29 23:52:36 | 2020-12-08 09:30:24 |
| [genv](https://github.com/sakirsensoy/genv) | 25 | 1 | 0 | Genv is a library for Go (golang) that makes it easy to read and use environment variables in your projects. It also allows environment variables to be loaded from the .env file. | 2019-07-15 10:25:57 | 2021-09-02 08:51:39 |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 19 | 2 | 0 | Library providing routines to merge and validate JSON, YAML and/or TOML files | 2018-02-01 19:06:15 | 2021-05-20 08:54:46 |
| [go-ssm-config](http://godoc.org/github.com/ianlopshire/go-ssm-config) | 10 | 6 | 4 | Go utility for loading configuration parameters from AWS SSM (Parameter Store) | 2019-12-02 18:47:38 | 2021-05-20 10:56:58 |
| [envconf](https://godoc.org/github.com/tomazk/envcfg) | 10 | 3 | 0 | Configure Go applications from the environment | 2014-10-26 12:12:26 | 2020-08-31 14:12:53 |
| [env](https://github.com/nasermirzaei89/env) | 7 | 0 | 0 | Golang Get Environment Variables Package | 2019-07-24 06:37:13 | 2021-09-17 17:07:44 |
| [go-ini](https://github.com/subpop/go-ini) | 5 | 1 | 1 | automatic mirror of https://git.sr.ht/~spc/go-ini | 2019-09-11 18:38:20 | 2021-06-09 06:01:09 |
| [swap](https://github.com/oblq/swap) | 4 | 0 | 0 | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). | 2020-04-12 23:28:19 | 2021-02-25 07:52:33 |
| [typenv](https://github.com/diegomarangoni/typenv) | 4 | 0 | 0 | Go minimalist typed environment variables library | 2020-06-30 18:26:09 | 2021-02-08 22:09:22 |
| [gonfig](https://github.com/miladabc/gonfig) | 2 | 0 | 0 | Tag based configuration loader from different providers | 2021-01-21 13:44:44 | 2021-08-02 20:37:05 |
</details>

### Continuous Integration
Tools for help with continuous integration.

<sup>*Last Update: 2023-09-30 16:56:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gitness](https://gitness.com) | 29,192 | 2,707 | 119 | Gitness is an Open Source developer platform with Source Control management, Continuous Integration and Continuous Delivery. | 2014-02-07 07:54:44 | 2023-09-30 09:27:04 |
| [cds](https://ovh.github.io/cds/) | 4,304 | 405 | 140 | Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform | 2016-10-11 08:28:23 | 2023-09-29 10:34:53 |
| [goveralls](https://github.com/mattn/goveralls) | 766 | 140 | 17 | A tool for testing, building, signing, and publishing binaries. | 2013-04-17 10:58:40 | 2023-09-10 09:19:12 |
| [overalls](https://github.com/go-playground/overalls) | 113 | 25 | 2 | :jeans:Multi-Package go project coverprofile for tools like goveralls | 2015-07-30 11:30:11 | 2023-09-26 23:10:39 |
| [duci](https://github.com/duck8823/duci) | 74 | 5 | 10 | The simple ci server  | 2018-04-01 01:51:02 | 2023-04-13 01:57:00 |
| [gomason](https://github.com/nikogura/gomason) | 58 | 8 | 4 | A tool for testing, building, signing, and publishing binaries. | 2017-11-18 00:59:11 | 2023-08-20 07:22:59 |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 19 | 5 | 0 | A Go recursive coverage testing tool | 2016-06-26 07:45:32 | 2022-09-26 09:24:45 |
</details>

### Data Structures
Generic datastructures and algorithms in Go.

<sup>*Last Update: 2021-09-25 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gods](https://github.com/emirpasic/gods) | 10,500 | 1,252 | 53 | GoDS (Go Data Structures). Containers (Sets, Lists, Stacks, Maps, Trees), Sets (HashSet, TreeSet, LinkedHashSet), Lists (ArrayList, SinglyLinkedList, DoublyLinkedList), Stacks (LinkedListStack, ArrayStack), Maps (HashMap, TreeMap, HashBidiMap, TreeBidiMap, LinkedHashMap), Trees (RedBlackTree, AVLTree, BTree, BinaryHeap), Comparators, Iterators, Enumerables, Sort, JSON | 2015-03-04 14:19:52 | 2021-09-23 01:11:20 |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 6,178 | 724 | 25 | A collection of useful, performant, and threadsafe Go datastructures. | 2014-10-29 13:55:17 | 2021-09-23 01:11:35 |
| [golang-set](https://github.com/deckarep/golang-set) | 2,074 | 177 | 12 | A simple set type for the Go language. Trusted by Docker, 1Password, Ethereum and Hashicorp. | 2013-07-03 21:52:01 | 2021-09-22 15:02:30 |
| [gota](https://github.com/go-gota/gota) | 1,771 | 181 | 45 | Gota: DataFrames and data wrangling in Go (Golang) | 2016-02-06 17:23:25 | 2021-09-23 00:38:18 |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1,375 | 94 | 10 | Probabilistic data structures for processing continuous, unbounded streams. | 2015-02-06 02:01:26 | 2021-09-21 19:50:07 |
| [roaring](http://roaringbitmap.org/) | 1,325 | 146 | 62 | Roaring bitmaps in Go (golang) | 2014-07-10 20:14:34 | 2021-09-20 04:55:24 |
| [bloom](https://github.com/bits-and-blooms/bloom) | 1,289 | 174 | 7 | Go package implementing Bloom filters | 2011-05-21 14:18:41 | 2021-09-21 03:19:58 |
| [gocache](https://vincent.composieux.fr/article/i-wrote-gocache-a-complete-and-extensible-go-cache-library/) | 955 | 90 | 10 | ☔️ A complete Go cache library that brings you multiple ways of managing your caches | 2019-10-05 08:13:54 | 2021-09-23 01:51:24 |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 823 | 68 | 12 | Cuckoo Filter: Practically Better Than Bloom | 2015-06-28 23:22:09 | 2021-09-18 19:36:34 |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 751 | 53 | 3 | HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction) | 2017-06-18 11:18:12 | 2021-09-16 09:45:45 |
| [bitset](https://github.com/bits-and-blooms/bitset) | 730 | 120 | 3 | Go package implementing bitsets | 2011-05-11 03:33:44 | 2021-09-20 23:42:23 |
| [algorithms](https://github.com/shady831213/algorithms) | 562 | 95 | 0 | CLRS study. Codes are written with golang. | 2018-01-31 09:27:56 | 2021-09-20 17:20:20 |
| [trie](https://github.com/derekparker/trie) | 550 | 91 | 10 | Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. | 2014-03-06 22:01:49 | 2021-09-18 10:09:55 |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 333 | 40 | 3 | Go native library for fast point tracking and K-Nearest queries | 2015-01-22 12:26:17 | 2021-09-14 09:55:48 |
| [gostl](https://github.com/liyue201/gostl) | 307 | 62 | 1 | Data structure and algorithm library for go, designed to provide functions similar to C++ STL | 2019-10-12 01:10:24 | 2021-09-20 06:48:17 |
| [go-edlib](https://github.com/hbollon/go-edlib) | 285 | 14 | 1 | 📚 String comparison and edit distance algorithms library, featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc... | 2020-08-18 09:30:59 | 2021-09-23 00:46:18 |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 273 | 45 | 0 | An in-memory string-interface{} map with various expiration options for golang | 2014-12-13 01:55:40 | 2021-09-18 07:57:08 |
| [merkletree](https://github.com/cbergoon/merkletree) | 270 | 72 | 6 | A Merkle Tree implementation written in Go. | 2017-04-12 02:50:11 | 2021-09-22 12:06:29 |
| [deque](https://github.com/gammazero/deque) | 246 | 20 | 0 | Fast ring-buffer deque (double-ended queue) | 2018-04-24 02:57:55 | 2021-09-22 11:03:42 |
| [hilbert](https://github.com/google/hilbert) | 235 | 33 | 2 | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. | 2015-08-06 15:50:00 | 2021-07-19 15:04:20 |
| [goskiplist](https://github.com/ryszard/goskiplist) | 229 | 57 | 6 | A skip list implementation in Go | 2012-05-09 05:44:59 | 2021-09-09 10:43:22 |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 209 | 32 | 0 | Adaptive Radix Trees implemented in Go | 2016-04-01 01:40:40 | 2021-09-21 09:31:06 |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 171 | 28 | 2 | A binary stream packer and unpacker | 2016-02-02 10:06:11 | 2021-09-16 03:47:31 |
| [cuckoo-filter](https://github.com/linvon/cuckoo-filter) | 170 | 13 | 1 | Cuckoo Filter go implement, better than Bloom Filter, configurable and space optimized  布谷鸟过滤器的Go实现，优于布隆过滤器，可以定制化过滤器参数，并进行了空间优化 | 2021-02-19 12:27:43 | 2021-09-15 07:53:02 |
| [skiplist](https://github.com/MauriceGit/skiplist) | 168 | 25 | 3 | A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist | 2018-06-23 16:01:51 | 2021-09-12 08:48:06 |
| [levenshtein](https://github.com/agnivade/levenshtein) | 161 | 11 | 1 | Go implementation to calculate Levenshtein Distance. | 2014-07-30 14:03:55 | 2021-09-22 23:11:54 |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 145 | 10 | 1 | Go concurrent-safe, goroutine-safe, thread-safe queue | 2019-01-10 21:21:23 | 2021-09-14 19:51:48 |
| [bloom](http://zhen.org/blog/benchmarking-bloom-filters-and-hash-functions-in-go/) | 144 | 15 | 1 | Bloom filters implemented in Go. | 2013-09-03 02:27:35 | 2021-06-04 02:57:34 |
| [iter](https://github.com/disksing/iter) | 137 | 7 | 1 | Go implementation of C++ STL iterators and algorithms. | 2019-10-20 09:29:49 | 2021-08-25 05:53:20 |
| [ring](https://pkg.go.dev/github.com/tannerryan/ring) | 120 | 12 | 2 | Package ring provides a high performance and thread safe Go implementation of a bloom filter. | 2019-01-27 04:02:20 | 2021-09-19 10:24:03 |
| [go-rquad](https://github.com/arl/go-rquad) | 117 | 3 | 0 | :pushpin: State of the art point location and neighbour finding algorithms for region quadtrees, in Go | 2016-09-12 21:46:37 | 2021-09-06 11:39:02 |
| [encoding](http://zhen.org/blog/benchmarking-integer-compression-in-go/) | 108 | 13 | 1 | Integer Compression Libraries for Go | 2013-09-20 19:29:57 | 2021-08-19 19:42:42 |
| [bit](https://yourbasic.org/algorithms/your-basic-int/#simple-sets) | 107 | 16 | 0 | Bitset data structure | 2017-05-03 19:05:35 | 2021-07-25 23:57:42 |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 100 | 6 | 1 | Cache Slow Database Queries | 2019-04-04 20:24:25 | 2021-07-23 06:11:07 |
| [conjungo](https://github.com/InVisionApp/conjungo) | 95 | 12 | 10 | A small flexible merge library in go | 2016-12-29 23:50:38 | 2021-09-13 11:41:04 |
| [skiplist](https://github.com/gansidui/skiplist) | 75 | 18 | 1 | skiplist for golang | 2014-11-18 16:29:53 | 2021-08-25 03:13:20 |
| [go-mcache](https://pkg.go.dev/github.com/OrlovEvgeny/go-mcache?tab=doc) | 73 | 10 | 0 | Fast in-memory key:value store/cache with TTL | 2018-04-14 23:31:21 | 2021-09-13 18:09:18 |
| [bloom](https://yourbasic.org/algorithms/bloom-filter/) | 63 | 8 | 0 | Probabilistic set data structure | 2017-05-06 19:57:47 | 2021-08-01 22:35:21 |
| [levenshtein](https://github.com/agext/levenshtein) | 59 | 5 | 0 | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. | 2016-04-08 00:14:31 | 2021-07-11 12:05:15 |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 53 | 3 | 0 | Go implementation of Count-Min-Log | 2015-08-16 22:31:36 | 2021-03-31 12:17:31 |
| [crunch](https://github.com/superwhiskers/crunch) | 44 | 6 | 0 | take bytes out of things easily ✨🍪 | 2019-02-27 03:56:52 | 2021-08-07 04:29:50 |
| [nan](https://github.com/kak-tus/nan) | 41 | 6 | 0 | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers | 2020-05-05 20:20:54 | 2021-09-03 16:59:43 |
| [goset](https://github.com/zoumo/goset) | 41 | 12 | 0 | Set is a useful collection but there is no built-in implementation in Go lang. | 2017-08-25 09:21:30 | 2021-09-14 09:52:17 |
| [hide](https://godoc.org/github.com/yaa110/goterator) | 36 | 3 | 0 | ID type with marshalling to/from hash to prevent sending IDs to clients. | 2019-01-16 13:54:17 | 2021-09-07 12:09:57 |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 35 | 5 | 0 | Highly concurrent drop-in replacement for bufio.Writer | 2017-09-18 15:29:59 | 2021-08-28 10:51:54 |
| [pipeline](https://godoc.org/github.com/hyfather/pipeline) | 34 | 3 | 0 | Pipelines using goroutines | 2018-04-25 00:11:36 | 2021-09-02 02:56:53 |
| [timedmap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 31 | 3 | 0 | A thread safe map which has expiring key-value pairs. | 2019-01-30 12:55:37 | 2021-09-13 09:34:16 |
| [deque](https://github.com/edwingeng/deque) | 28 | 1 | 0 | A highly optimized double-ended queue | 2019-02-01 03:32:28 | 2021-07-15 17:03:29 |
| [typ](https://github.com/gurukami/typ) | 28 | 1 | 0 | Null Types, Safe primitive type conversion and fetching value from complex structures. | 2019-03-03 05:34:23 | 2021-09-23 17:56:10 |
| [dict](https://github.com/srfrog/dict) | 22 | 2 | 0 | Python-like dictionaries for Go | 2019-04-23 02:04:25 | 2021-09-13 00:58:17 |
| [null](https://github.com/emvi/null) | 21 | 1 | 1 | Nullable Go types that can be marshalled/unmarshalled to/from JSON. | 2018-07-04 21:18:45 | 2021-08-17 09:56:10 |
| [go-ef](https://github.com/amallia/go-ef) | 18 | 5 | 0 | A Go implementation of the Elias-Fano encoding | 2017-09-22 01:47:16 | 2021-05-05 17:59:40 |
| [cmap](https://github.com/lrita/cmap) | 18 | 0 | 0 | a thread-safe concurrent map for go | 2019-11-26 03:54:59 | 2021-09-08 10:59:33 |
| [mspm](https://github.com/BlackRabbitt/mspm) | 16 | 1 | 0 | Multi-String Pattern Matching Algorithm Using TrieHashNode | 2018-05-17 18:59:44 | 2021-07-12 07:55:03 |
| [ptrie](https://godoc.org/github.com/hyfather/pipeline) | 16 | 4 | 0 | A prefix tree implementation in go  | 2019-05-20 14:13:05 | 2021-08-18 14:00:09 |
| [set](https://github.com/StudioSol/set) | 15 | 6 | 1 | A simple Set data structure implementation in Go (Golang) using LinkedHashMap. | 2018-07-20 21:53:37 | 2021-07-04 12:30:18 |
| [treap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 12 | 3 | 0 | golang persistent immutable treap sorted sets | 2018-09-16 01:38:03 | 2021-07-15 16:36:40 |
| [parapipe](https://github.com/nazar256/parapipe) | 10 | 0 | 1 | Paralleling pipeline | 2021-04-09 06:49:56 | 2021-08-10 04:48:12 |
| [gofal](https://github.com/xxjwxc/gofal) | 9 | 0 | 0 | fractional api base on golang . golang math tools fractional molecular denominator 分数计算 分子 分母 运算 | 2019-08-05 07:37:55 | 2020-10-08 14:54:15 |
| [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) | 8 | 1 | 1 | Ordered-concurrently a library for parallel processing with ordered output in Go. Process work concurrently / in parallel and returns output in a channel in the order of input. It is useful in concurrently / parallelly processing items in a queue, and get output in the order provided by the queue. | 2021-02-28 17:56:05 | 2021-08-13 21:16:35 |
| [parsefields](https://github.com/MonaxGT/parsefields) | 6 | 0 | 0 | Tools for parse JSON-like logs for collecting unique fields and events | 2019-04-12 22:15:10 | 2021-05-15 04:37:38 |
| [goterator](https://godoc.org/github.com/yaa110/goterator) | 5 | 1 | 0 | Lazy iterator implementation for Golang | 2020-08-12 19:47:57 | 2021-09-07 12:08:53 |
| [dsu](https://github.com/ihebu/dsu) | 5 | 0 | 0 | Disjoint Set data structure implementation in Go | 2021-04-27 16:35:38 | 2021-06-21 04:48:16 |
| [bloomfilter](https://github.com/OldPanda/bloomfilter) | 5 | 0 | 0 | Yet another Bloomfilter implementation in Go, compatible with Java's Guava library | 2021-01-01 01:28:04 | 2021-09-18 20:36:30 |
| [slices](https://github.com/srfrog/slices) | 3 | 0 | 0 | Functions that operate on slices. Similar to functions from package strings or package bytes that have been adapted to work with slices. | 2020-07-02 23:17:34 | 2021-05-08 03:37:15 |
</details>

### Database - Database schema migration


<sup>*Last Update: 2021-07-17 19:25:15*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [migrate](https://github.com/golang-migrate/migrate) | 6,750 | 714 | 135 | Database migrations. CLI and Golang library. | 2018-01-19 09:30:58 | 2021-07-17 07:35:53 |
| [sql-migrate](https://github.com/rubenv/sql-migrate) | 2,176 | 210 | 64 | SQL schema migration tool for Go. | 2014-09-09 07:31:41 | 2021-07-14 22:43:36 |
| [goose](https://github.com/pressly/goose) | 1,754 | 270 | 46 | Goose database migration tool - fork of https://bitbucket.org/liamstask/goose | 2016-02-25 20:39:37 | 2021-07-16 16:32:05 |
| [pop](https://github.com/gobuffalo/pop) | 1,091 | 217 | 120 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2021-07-14 05:23:48 |
| [skeema](http://pravasan.github.io/pravasan/) | 917 | 80 | 23 | Schema management CLI for MySQL | 2016-10-31 23:18:56 | 2021-07-16 21:51:38 |
| [gormigrate](https://pkg.go.dev/github.com/go-gormigrate/gormigrate/v2?tab=doc) | 634 | 69 | 11 | Minimalistic database migration helper for Gorm ORM | 2016-08-31 11:46:23 | 2021-07-13 07:17:07 |
| [darwin](https://github.com/GuiaBolso/darwin) | 120 | 23 | 4 | Database schema evolution library for Go | 2016-04-05 15:57:59 | 2021-07-16 20:19:32 |
| [migrator](https://github.com/lopezator/migrator) | 115 | 15 | 5 | Dead simple Go database migration library. | 2019-02-04 09:40:01 | 2021-06-30 17:41:53 |
| [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) | 79 | 17 | 2 | A Go package to help write migrations with go-pg/pg. | 2018-08-11 07:00:13 | 2021-06-26 16:10:37 |
| [go-fixtures](https://github.com/RichardKnop/go-fixtures) | 26 | 8 | 0 | Django style fixtures for Golang's excellent built-in database/sql library. | 2015-12-24 11:27:45 | 2021-04-08 02:57:47 |
| [pravasan](http://pravasan.github.io/pravasan/) | 24 | 4 | 30 | Simple Migration Tool - written in Go | 2015-01-03 17:11:05 | 2019-03-22 13:54:35 |
| [avro](https://github.com/khezen/avro) | 24 | 3 | 0 | Apache AVRO for go | 2019-04-07 12:22:46 | 2021-06-13 12:52:12 |
| [migrator](https://github.com/larapulse/migrator) | 8 | 1 | 0 | MySQL database migrator | 2020-06-27 14:40:29 | 2021-07-14 12:58:51 |
| [schema](http://pravasan.github.io/pravasan/) | 7 | 1 | 3 | Embedded schema migration package for Go | 2019-09-24 19:27:13 | 2021-01-05 05:27:47 |
| [go-pg-migrate](https://pkg.go.dev/github.com/lawzava/go-pg-migrate) | 4 | 1 | 0 | CLI-friendly package for go-pg migrations management. | 2021-01-16 17:01:32 | 2021-06-26 17:00:23 |
</details>

### Database - Database tools


<sup>*Last Update: 2021-09-03 09:25:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [vitess](http://vitess.io) | 12,426 | 1,584 | 696 | Vitess is a database clustering system for horizontal scaling of MySQL. | 2013-06-27 21:20:28 | 2021-09-03 02:02:23 |
| [pgweb](http://sosedoff.github.io/pgweb) | 6,993 | 553 | 57 | Cross-platform client for PostgreSQL databases | 2014-10-09 01:41:32 | 2021-09-02 15:55:14 |
| [kingshard](https://github.com/flike/kingshard) | 5,794 | 1,156 | 155 | A high-performance MySQL proxy | 2015-07-04 02:22:32 | 2021-09-02 09:59:18 |
| [orchestrator](https://github.com/openark/orchestrator) | 4,217 | 728 | 346 | MySQL replication topology management and HA | 2016-11-30 13:44:24 | 2021-09-02 13:27:41 |
| [go-mysql-elasticsearch](https://github.com/go-mysql-org/go-mysql-elasticsearch) | 3,574 | 699 | 199 | Sync MySQL data into elasticsearch  | 2015-01-15 09:54:18 | 2021-09-01 08:34:28 |
| [go-mysql](https://github.com/go-mysql-org/go-mysql) | 3,204 | 728 | 150 | a powerful mysql toolset with Go | 2014-02-21 01:56:45 | 2021-09-02 11:07:44 |
| [prest](https://prestd.com) | 2,795 | 204 | 135 | pREST (PostgreSQL REST), low-code, simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new | 2016-11-22 05:17:05 | 2021-09-03 01:21:26 |
| [chproxy](https://github.com/Vertamedia/chproxy) | 682 | 153 | 37 | ClickHouse http proxy and load balancer | 2017-09-18 13:09:23 | 2021-09-02 08:22:41 |
| [pg_timetable](https://www.cybertec-postgresql.com/en/products/pg_timetable/) | 539 | 27 | 2 | pg_timetable: Advanced scheduling for PostgreSQL | 2018-12-19 10:19:51 | 2021-09-02 16:35:27 |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 295 | 62 | 14 | Collects many small inserts to ClickHouse and send in big inserts | 2017-04-29 10:38:41 | 2021-08-30 07:05:06 |
| [myreplication](https://github.com/2tvenom/myreplication) | 174 | 46 | 5 | Golang MySql binary log replication listener | 2015-02-04 20:59:49 | 2021-07-14 10:43:59 |
| [octillery](https://github.com/blastrain/octillery) | 143 | 18 | 6 | Go package for sharding databases ( Supports every ORM or raw SQL ) | 2018-11-26 10:39:35 | 2021-09-01 19:04:22 |
| [dbbench](https://github.com/sj14/dbbench) | 59 | 8 | 12 | 🏋️ dbbench is a simple database benchmarking tool which supports several databases and own scripts | 2018-11-24 13:21:18 | 2021-09-01 06:10:49 |
| [datagen](https://github.com/codingconcepts/datagen) | 39 | 4 | 0 | A fast data generator that's multi-table aware and supports multi-row DML. | 2019-04-18 19:58:01 | 2021-08-28 12:54:51 |
| [prep](http://sosedoff.github.io/pgweb) | 28 | 2 | 0 | Prep finds all SQL statements in a Go package and instruments db connection with prepared statements | 2017-12-11 23:47:38 | 2021-05-03 10:39:05 |
| [rwdb](https://prestd.com) | 12 | 0 | 0 | Database wrapper that manage read write connections | 2017-10-04 03:55:29 | 2021-05-29 03:21:37 |
</details>

### Database - Databases implemented in Go


<sup>*Last Update: 2021-07-31 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prometheus](https://prometheus.io/) | 37,888 | 6,217 | 477 | The Prometheus monitoring system and time series database. | 2012-11-24 11:14:12 | 2021-07-31 01:18:18 |
| [tidb](https://pingcap.com) | 28,574 | 4,526 | 2,153 | TiDB is an open source distributed HTAP database compatible with the MySQL protocol  | 2015-09-06 04:01:52 | 2021-07-30 17:38:54 |
| [influxdb](https://influxdata.com) | 21,843 | 2,986 | 1,369 | Scalable datastore for metrics, events, and real-time analytics | 2013-09-26 14:31:10 | 2021-07-30 22:32:53 |
| [cockroach](https://www.cockroachlabs.com) | 21,306 | 2,790 | 4,036 | CockroachDB - the open source, cloud-native distributed SQL database. | 2014-02-06 00:18:47 | 2021-07-31 01:18:31 |
| [dgraph](https://dgraph.io) | 16,411 | 1,193 | 88 | Native GraphQL Database with graph backend | 2015-08-25 07:15:56 | 2021-07-31 02:25:02 |
| [groupcache](https://github.com/golang/groupcache) | 10,486 | 1,175 | 35 | groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. | 2013-07-22 21:55:07 | 2021-07-30 15:51:54 |
| [badger](https://blog.dgraph.io/post/badger/) | 9,515 | 794 | 3 | Fast key-value DB in Go. | 2017-01-26 05:09:49 | 2021-07-31 02:15:22 |
| [rqlite](http://www.rqlite.com) | 8,633 | 450 | 54 | The lightweight, distributed relational database built on SQLite | 2014-08-23 04:31:18 | 2021-07-30 23:14:40 |
| [go-cache](https://patrickmn.com/projects/go-cache/) | 5,207 | 655 | 53 | An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. | 2012-01-02 13:07:13 | 2021-07-30 15:54:50 |
| [bigcache](http://allegro.tech/2016/03/writing-fast-cache-service-in-go.html) | 5,002 | 421 | 53 | Efficient cache for gigabytes of data written in Go. | 2016-03-23 07:18:52 | 2021-07-30 16:39:39 |
| [bbolt](https://go.etcd.io/bbolt) | 4,654 | 352 | 107 | An embedded key/value database for Go. | 2017-06-17 01:42:09 | 2021-07-30 09:31:20 |
| [VictoriaMetrics](https://victoriametrics.com/) | 4,631 | 406 | 327 | VictoriaMetrics: fast, cost-effective monitoring solution and time series database | 2018-09-30 09:58:01 | 2021-07-31 01:02:19 |
| [goleveldb](https://patrickmn.com/projects/go-cache/) | 4,464 | 672 | 72 | LevelDB key/value database in Go. | 2013-01-23 04:08:58 | 2021-07-30 01:26:39 |
| [ledisdb](https://ledisdb.io) | 3,675 | 417 | 1 | A high performance NoSQL Database Server powered by Go | 2014-04-30 00:43:09 | 2021-07-30 08:13:41 |
| [buntdb](https://github.com/tidwall/buntdb) | 3,370 | 236 | 4 | BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support | 2016-07-19 22:11:40 | 2021-07-30 22:56:01 |
| [immudb](https://codenotary.com/technologies/immudb) | 3,025 | 123 | 62 | immudb - world’s fastest immutable database | 2019-11-07 08:22:16 | 2021-07-30 19:48:10 |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2,606 | 260 | 26 | A rudimentary implementation of a basic document (NoSQL) database in Go | 2013-05-26 10:03:49 | 2021-07-29 20:51:44 |
| [nutsdb](https://xujiajun.cn/nutsdb/) | 1,696 | 145 | 29 | A simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. | 2018-12-07 07:03:38 | 2021-07-30 09:06:17 |
| [gcache](https://github.com/bluele/gcache) | 1,621 | 197 | 18 | An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC | 2015-01-24 18:17:07 | 2021-07-30 07:21:39 |
| [cache2go](https://github.com/muesli/cache2go) | 1,574 | 443 | 23 | Concurrency-safe Go caching library with expiration capabilities and access counters | 2013-11-11 03:45:02 | 2021-07-30 03:12:48 |
| [rosedb](https://space.bilibili.com/26194591) | 1,541 | 225 | 6 | 🚀A fast, stable and embedded k-v database in pure Golang, supports string, list, hash, set, sorted set. 一个 Go 语言实现的快速、稳定、内嵌的 k-v 数据库。 | 2020-12-06 07:02:48 | 2021-07-31 02:23:36 |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 1,202 | 92 | 29 | Fast thread-safe inmemory cache for big number of entries in Go. Minimizes GC overhead | 2018-11-22 22:50:13 | 2021-07-30 07:15:02 |
| [CovenantSQL](https://developers.covenantsql.io) | 1,200 | 132 | 25 | A decentralized, trusted, high performance, SQL database with blockchain features | 2018-04-11 09:52:58 | 2021-07-30 09:34:13 |
| [diskv](http://godoc.org/github.com/peterbourgon/diskv) | 1,076 | 89 | 7 | A disk-backed key-value store. | 2012-03-21 16:44:32 | 2021-07-28 00:39:20 |
| [databunker](https://databunker.org/) | 872 | 34 | 1 | Secure vault for customer records built to comply with GDPR | 2019-12-08 21:55:55 | 2021-07-30 12:32:05 |
| [eliasdb](https://github.com/krotik/eliasdb) | 820 | 41 | 10 | EliasDB a graph-based database. | 2016-08-13 13:53:28 | 2021-07-29 18:33:11 |
| [moss](https://github.com/couchbase/moss) | 819 | 48 | 46 | moss - a simple, fast, ordered, persistable, key-val storage library for golang | 2016-02-06 20:27:22 | 2021-07-27 16:21:15 |
| [pogreb](https://github.com/akrylysov/pogreb) | 801 | 57 | 9 | Embedded key-value store for read-heavy workloads written in Go | 2018-01-06 23:16:36 | 2021-07-27 15:02:44 |
| [bitcask](https://prologic.github.io/bitcask) | 759 | 70 | 10 | 🔑A high performance Key/Value store written in Go with a predictable read/write performance and high throughput. Uses a Bitcask on-disk layout (LSM+WAL) similar to Riak. | 2019-03-12 13:57:35 | 2021-06-22 16:39:37 |
| [levigo](https://github.com/jmhodges/levigo) | 399 | 82 | 3 | levigo is a Go wrapper for LevelDB | 2012-01-17 08:17:54 | 2021-07-21 07:42:55 |
| [pudge](https://github.com/recoilme/pudge) | 293 | 21 | 0 | Fast and simple key/value store written using Go's standard library | 2018-11-20 10:11:53 | 2021-07-28 15:04:22 |
| [vasto](https://github.com/chrislusf/vasto) | 214 | 24 | 4 | A distributed key-value store. On Disk. Able to grow or shrink without service interruption. | 2018-01-16 05:16:57 | 2021-07-10 23:20:45 |
| [kivik](https://github.com/go-kivik/kivik) | 213 | 29 | 10 | Kivik provides a common interface to CouchDB or CouchDB-like databases for Go and GopherJS. | 2017-02-09 14:14:54 | 2021-07-22 12:10:08 |
| [piladb](https://www.piladb.org) | 188 | 19 | 9 | Lightweight RESTful database engine based on stack data structures | 2015-09-08 23:12:22 | 2021-07-29 20:52:07 |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 128 | 13 | 1 | A tiny Golang JSON database | 2018-06-21 22:13:33 | 2021-07-04 23:37:57 |
| [slowpoke](https://github.com/recoilme/slowpoke) | 97 | 8 | 0 | Low-level key/value store in pure Go.  | 2018-02-19 09:22:37 | 2021-07-13 16:52:14 |
| [cache](https://github.com/akyoto/cache) | 89 | 9 | 0 | :handbag: Cache arbitrary data with an expiration time. | 2019-05-11 12:42:45 | 2021-07-05 09:48:23 |
| [unitdb](https://github.com/unit-io/unitdb) | 68 | 7 | 0 | Fast specialized time-series database for IoT, real-time internet connected devices and AI analytics. | 2019-08-29 18:21:27 | 2021-07-30 08:54:52 |
| [bcache](https://github.com/iwanbk/bcache) | 67 | 10 | 3 | Eventually consistent distributed in-memory  cache Go library | 2018-12-26 15:45:16 | 2021-07-22 09:24:38 |
| [couchcache](https://github.com/codingsince1985/couchcache) | 53 | 4 | 0 | A RESTful caching micro-service in Go backed by Couchbase | 2015-04-05 07:13:05 | 2021-03-03 04:44:10 |
| [hare](https://github.com/jameycribbs/hare) | 41 | 3 | 1 | Hare is a nimble little database management system for Go. | 2016-10-05 20:05:45 | 2021-06-21 15:45:05 |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 37 | 4 | 2 | golang bigcache with clustering as a library. | 2017-12-18 07:48:07 | 2021-03-12 09:43:04 |
| [coffer](https://github.com/claygod/coffer) | 25 | 1 | 0 | Simply ACID* key-value database. At the medium or even low latency it tries to provide greater throughput without losing the ACID properties of the database. The database provides the ability to create record headers at own discretion and use them as transactions. The maximum size of stored data is limited by the size of the computer's RAM. | 2019-05-13 18:30:23 | 2021-05-11 17:35:16 |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 15 | 1 | 0 | Key-value store for temporary items :memo: | 2017-03-17 18:03:42 | 2021-04-26 04:59:34 |
| [ttlcache](https://github.com/cheshir/ttlcache) | 3 | 1 | 0 | Simple in-memory key-value storage with TTL for each record. | 2021-01-06 19:24:26 | 2021-05-24 05:31:48 |
| [bitcask](http://allegro.tech/2016/03/writing-fast-cache-service-in-go.html) | 0 | 0 | 0 | 🔑 A high performance Key/Value store written in Go with a predictable read/write performance and high throughput. Uses a Bitcask on-disk layout (LSM+WAL) similar to Riak. | 2021-07-12 14:52:18 | 2021-07-12 14:53:24 |
</details>

### Database - SQL query builder
libraries for building and using SQL

<sup>*Last Update: 2021-09-11 09:25:11*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [squirrel](https://github.com/Masterminds/squirrel) | 4,179 | 321 | 44 | Fluent SQL generation for golang | 2014-01-18 05:29:58 | 2021-09-10 16:36:00 |
| [xo](https://github.com/xo/xo) | 2,897 | 260 | 18 | Command line tool to generate idiomatic Go code for SQL databases supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server | 2016-02-05 10:22:20 | 2021-09-11 01:17:59 |
| [goqu](http://doug-martin.github.io/goqu/) | 1,260 | 110 | 43 | SQL builder and query library for golang | 2015-02-21 01:06:00 | 2021-09-10 22:20:50 |
| [gendry](https://github.com/didi/gendry) | 1,259 | 156 | 6 | a golang library for sql builder | 2017-12-01 08:15:43 | 2021-09-06 07:15:26 |
| [dotsql](https://github.com/qustavo/dotsql) | 597 | 45 | 6 | A Golang library for using SQL. | 2014-11-20 12:14:39 | 2021-09-07 13:33:20 |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 541 | 48 | 36 | A Go (golang) package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. | 2015-12-10 22:39:26 | 2021-08-26 21:11:12 |
| [jet](https://github.com/go-jet/jet) | 477 | 29 | 18 | Type safe SQL builder with code generation and automatic query result data mapping | 2019-03-02 11:06:23 | 2021-09-11 01:51:00 |
| [dbq](https://github.com/rocketlaunchr/dbq) | 316 | 17 | 1 | Zero boilerplate database operations for Go | 2019-07-11 02:17:33 | 2021-09-08 06:13:36 |
| [sqrl](https://github.com/elgris/sqrl) | 230 | 28 | 7 | Fluent SQL generation for golang | 2014-06-25 10:03:06 | 2021-09-08 03:52:03 |
| [sqlingo](https://github.com/lqs/sqlingo) | 155 | 11 | 0 | 💥 A lightweight DSL & ORM which helps you to write SQL in Go. | 2018-11-18 14:11:03 | 2021-08-29 15:21:08 |
| [go-structured-query](https://bokwoon95.github.io/sq/) | 121 | 6 | 2 | Type safe SQL query builder and struct mapper for Go | 2020-05-30 14:07:30 | 2021-09-11 01:34:21 |
| [igor](https://github.com/galeone/igor) | 84 | 2 | 0 | igor is an abstraction layer for PostgreSQL with a gorm like syntax. | 2016-03-10 14:45:08 | 2021-08-02 09:41:49 |
| [go-hasql](https://golang.yandex) | 84 | 4 | 2 | Go library for accessing multi-host SQL database installations | 2020-08-19 09:56:00 | 2021-09-07 06:55:22 |
| [godbal](https://github.com/xujiajun/godbal) | 52 | 27 | 0 | Database Abstraction Layer (dbal) for Go. Support SQL builder and get result easily  (now only support mysql) | 2018-02-28 05:47:42 | 2021-04-22 06:11:02 |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 38 | 3 | 8 | Go database query builder library for PostgreSQL | 2019-08-18 08:18:21 | 2021-09-02 18:19:09 |
| [sqlf](https://github.com/leporo/sqlf) | 22 | 2 | 0 | Fast SQL query builder for Go | 2019-07-20 07:03:27 | 2021-09-05 17:13:27 |
| [qry](https://github.com/HnH/qry) | 20 | 3 | 1 | Write your SQL queries in raw files with all benefits of modern IDEs, use them in an easy way inside your application with all the profit of compile time constants | 2019-08-20 09:01:00 | 2021-02-03 14:40:44 |
| [gosql](https://twharmon.gitbook.io/gosql/) | 16 | 0 | 0 | SQL query builder for Go | 2020-01-08 17:13:09 | 2021-06-28 12:17:55 |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 9 | 0 | 0 | Golang package for MPTT (Modified Preorder Tree Traversal) - materialized path realisation. | 2020-01-09 15:04:45 | 2021-08-07 14:43:40 |
| [ormlite](https://github.com/pupizoid/ormlite) | 1 | 0 | 2 | Lightweight package containing some ORM-like features and helpers for sqlite databases. | 2018-06-28 13:42:09 | 2021-08-27 18:52:56 |
</details>

### Database Drivers - Multiple Backends.


<sup>*Last Update: 2021-08-25 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cayley](https://cayley.io) | 13,910 | 1,256 | 91 | An open-source graph database | 2014-06-05 18:49:41 | 2021-08-24 15:23:57 |
| [gokv](https://github.com/philippgille/gokv) | 353 | 35 | 31 | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) | 2018-10-08 18:55:22 | 2021-08-15 10:51:37 |
| [cachego](https://github.com/faabiosr/cachego) | 157 | 9 | 1 | Golang Cache component - Multiple drivers | 2016-10-05 18:10:03 | 2021-08-16 05:50:40 |
| [dsc](https://github.com/viant/dsc) | 22 | 6 | 0 | Datastore Connectivity in go | 2016-06-13 20:18:10 | 2021-04-17 04:50:27 |
</details>

### Database Drivers - NoSQL Databases


<sup>*Last Update: 2021-09-11 09:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [redis](https://redis.uptrace.dev/) | 12,386 | 1,585 | 116 | Type-safe Redis client for Golang | 2012-07-25 13:01:39 | 2021-09-11 01:58:54 |
| [redigo](https://github.com/gomodule/redigo) | 8,640 | 1,195 | 21 | Go client for Redis | 2012-04-14 04:31:58 | 2021-09-10 16:18:51 |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 6,062 | 693 | 12 | The Go driver for MongoDB | 2017-02-08 17:18:02 | 2021-09-10 20:50:42 |
| [gocql](http://gocql.github.io/) | 2,152 | 523 | 160 | Package gocql implements a fast and robust Cassandra client for the Go programming language. | 2012-08-26 15:42:42 | 2021-09-10 21:12:40 |
| [mgo](https://github.com/globalsign/mgo) | 1,908 | 233 | 64 | The MongoDB driver for Go | 2017-04-13 11:14:04 | 2021-09-07 07:14:57 |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1,572 | 177 | 19 | Go language driver for RethinkDB | 2013-09-12 13:56:27 | 2021-08-31 20:34:46 |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1,387 | 384 | 49 | Go Memcached client library #golang | 2011-06-28 19:29:12 | 2021-09-10 12:52:43 |
| [qmgo](https://github.com/qiniu/qmgo) | 664 | 79 | 23 | Qmgo - The Go driver for MongoDB. It‘s based on official mongo-go-driver but easier to use like Mgo. | 2020-08-04 09:06:00 | 2021-09-10 05:11:02 |
| [redeo](https://github.com/bsm/redeo) | 397 | 28 | 3 | High-performance framework for building redis-protocol compatible TCP servers/services | 2014-03-06 08:46:18 | 2021-08-25 03:38:14 |
| [neoism](https://github.com/jmcvetta/neoism) | 378 | 57 | 16 | Neo4j client for Golang | 2012-07-12 07:42:33 | 2021-09-10 10:01:00 |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 370 | 167 | 15 | Aerospike Client Go  | 2014-07-26 02:56:21 | 2021-09-10 13:33:13 |
| [mgm](https://github.com/Kamva/mgm) | 361 | 32 | 2 | Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver) | 2019-12-27 14:40:51 | 2021-09-09 19:45:58 |
| [gocb](http://blog.couchbase.com/2015/september/go-sdk-1.0-ga) | 328 | 92 | 0 | The Couchbase Go SDK | 2015-01-15 20:01:32 | 2021-09-02 12:56:46 |
| [go-couchbase](https://godoc.org/github.com/couchbase/go-couchbase) | 313 | 88 | 41 | Couchbase client in Go | 2012-01-19 22:52:08 | 2021-09-10 11:26:56 |
| [go-rejson](https://github.com/nitishm/go-rejson) | 177 | 29 | 7 | Golang client for redislabs' ReJSON module with support for multilple redis clients (redigo, go-redis) | 2018-04-23 00:51:05 | 2021-09-10 13:48:52 |
| [godis](https://github.com/piaohao/godis) | 96 | 16 | 0 | redis client implement by golang, inspired by jedis. | 2019-06-14 03:14:22 | 2021-08-20 12:35:27 |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 76 | 15 | 0 | Neo4j REST Client in golang | 2011-06-04 16:08:35 | 2020-08-28 21:16:55 |
| [arangolite](https://github.com/solher/arangolite) | 69 | 18 | 5 | Lightweight Golang driver for ArangoDB | 2015-10-04 17:27:34 | 2021-03-10 17:27:16 |
| [go-pilosa](https://www.pilosa.com/) | 46 | 20 | 13 | Go client library for Pilosa | 2016-09-30 21:37:10 | 2021-06-13 07:36:01 |
| [goforestdb](https://github.com/couchbase/goforestdb) | 30 | 4 | 7 | Go bindings for ForestDB | 2014-05-14 15:36:12 | 2021-02-25 12:12:32 |
| [neo4j](https://github.com/cihangir/neo4j) | 26 | 7 | 8 | Neo4j Rest API Client for Go lang | 2013-05-18 08:54:01 | 2020-08-28 21:16:50 |
| [goriak](https://godoc.org/gopkg.in/zegl/goriak.v3) | 25 | 5 | 5 | goriak - Go language driver for Riak KV | 2016-10-05 16:48:17 | 2021-07-19 21:24:49 |
| [xredis](https://github.com/shomali11/xredis) | 15 | 2 | 0 | Go Redis Client | 2017-06-14 00:19:26 | 2021-06-02 07:51:37 |
| [godscache](https://github.com/defcronyke/godscache) | 9 | 1 | 0 | An unofficial Google Cloud Platform Go Datastore wrapper that adds caching using memcached. For App Engine Flexible, Compute Engine, Kubernetes Engine, and more. | 2018-05-08 20:19:39 | 2021-07-06 18:09:14 |
| [asc](https://github.com/viant/asc) | 6 | 1 | 0 | Datastore Connectivity for Aerospike for go | 2016-06-13 20:22:31 | 2020-08-28 21:15:29 |
| [gocosmos](http://blog.couchbase.com/2015/september/go-sdk-1.0-ga) | 4 | 1 | 0 | Go driver for Azure CosmosDB SQL API | 2020-12-06 07:03:43 | 2021-07-20 08:21:50 |
</details>

### Database Drivers - Relational Databases


<sup>*Last Update: 2021-09-08 15:02:42*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql) | 11,446 | 1,989 | 90 | Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package | 2012-12-09 20:33:55 | 2021-09-08 03:44:20 |
| [pq](https://pkg.go.dev/github.com/lib/pq) | 6,755 | 820 | 283 | Pure Go Postgres driver for database/sql | 2012-03-12 18:50:22 | 2021-09-07 16:38:22 |
| [go-sqlite3](http://mattn.github.io/go-sqlite3) | 5,066 | 865 | 105 | sqlite3 driver for go using database/sql | 2011-11-11 12:36:50 | 2021-09-07 14:26:08 |
| [pgx](https://github.com/jackc/pgx) | 4,468 | 442 | 176 | PostgreSQL driver and toolkit for Go | 2013-03-30 19:06:26 | 2021-09-08 04:48:04 |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1,431 | 372 | 135 | Microsoft SQL server driver written in go language | 2013-12-16 00:10:47 | 2021-09-06 07:01:21 |
| [go-oci8](https://mattn.kaoriya.net/) | 557 | 202 | 6 | Oracle driver for Go using database/sql | 2012-02-29 12:19:16 | 2021-09-03 02:19:06 |
| [godror](https://github.com/godror/godror) | 289 | 51 | 8 | GO DRiver for ORacle DB | 2019-11-21 21:23:17 | 2021-09-03 10:21:03 |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 151 | 45 | 12 | Firebird RDBMS sql driver for Go (golang) | 2013-08-27 13:09:14 | 2021-08-15 00:09:19 |
| [go-adodb](http://mattn.kaoriya.net/) | 121 | 30 | 17 | Microsoft ActiveX Object DataBase driver for go that using exp/sql | 2011-11-14 04:32:50 | 2021-09-01 08:42:32 |
| [gofreetds](https://github.com/minus5/gofreetds) | 106 | 43 | 18 | Go Sql Server database driver. | 2012-12-06 17:29:26 | 2021-08-10 18:15:24 |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 77 | 19 | 1 | Mirror of Apache Calcite - Avatica Go SQL Driver | 2017-08-08 07:00:08 | 2021-08-13 14:02:25 |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 74 | 7 | 1 | SQLite with pure Go | 2020-06-06 20:37:12 | 2021-08-30 16:46:17 |
| [bgc](https://github.com/viant/bgc) | 15 | 4 | 0 | Datastore Connectivity for BigQuery in go | 2016-06-13 20:24:26 | 2020-09-06 06:58:35 |
| [pig](https://github.com/alexeyco/pig) | 4 | 0 | 0 | Simple pgx wrapper to execute and scan query results | 2021-04-15 15:33:23 | 2021-07-04 12:55:47 |
</details>

### Database Drivers - Search and Analytic Databases


<sup>*Last Update: 2023-09-30 16:56:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [bleve](https://github.com/blevesearch/bleve) | 9,289 | 682 | 293 | A modern text indexing library for go | 2014-04-17 21:02:18 | 2023-09-29 11:33:45 |
| [elastic](https://olivere.github.io/elastic/) | 7,227 | 1,185 | 107 | Deprecated: Use the official Elasticsearch client for Go at https://github.com/elastic/go-elasticsearch | 2012-12-06 17:15:33 | 2023-09-26 08:40:35 |
| [riot](https://github.com/go-ego/riot) | 6,097 | 483 | 50 | Go Open Source, Distributed, Simple and efficient Search Engine; Warning: This is V1 and beta version, because of big memory consume, and the V2 will be rewrite all code. | 2017-06-21 14:17:59 | 2023-09-30 09:34:14 |
| [go-elasticsearch](https://github.com/elastic/go-elasticsearch#go-elasticsearch) | 5,162 | 592 | 68 | The official Go client for Elasticsearch | 2017-03-27 17:56:15 | 2023-09-29 14:41:39 |
| [elasticsql](https://github.com/cch123/elasticsql) | 1,098 | 193 | 10 | convert sql to elasticsearch DSL in golang(go) | 2016-08-24 07:29:43 | 2023-09-22 03:09:28 |
| [elastigo](https://github.com/mattbaird/elastigo) | 950 | 254 | 70 | A Go (golang) based Elasticsearch client library. | 2012-10-12 04:19:59 | 2023-09-28 10:20:52 |
| [skizze](https://github.com/seiflotfy/skizze) | 88 | 9 | 0 | A probabilistic data structure service and storage | 2016-01-17 12:10:40 | 2023-05-24 04:56:57 |
| [goes](http://godoc.org/github.com/belogik/goes) | 30 | 14 | 0 | A library to interact with Elasticsearch in Go! | 2015-12-28 18:52:03 | 2023-09-28 14:15:37 |
</details>

### Date and Time
Libraries for working with dates and times.

<sup>*Last Update: 2021-09-08 15:02:33*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [now](https://github.com/jinzhu/now) | 3,255 | 194 | 6 | Now is a time toolkit for golang | 2013-11-18 10:55:30 | 2021-09-08 08:01:35 |
| [dateparse](https://github.com/araddon/dateparse) | 1,520 | 109 | 40 | GoLang Parse many date strings without knowing format in advance. | 2014-04-21 02:55:48 | 2021-09-07 10:32:54 |
| [carbon](https://github.com/uniplaces/carbon) | 619 | 46 | 12 | Carbon for Golang, an extension for Time | 2016-08-03 10:55:52 | 2021-09-08 05:18:36 |
| [durafmt](https://github.com/hako/durafmt) | 397 | 41 | 6 | :clock8:  Better time duration formatting in Go!  | 2016-05-20 21:49:33 | 2021-09-05 19:23:56 |
| [timeutil](https://github.com/leekchan/timeutil) | 185 | 12 | 1 | timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package | 2015-08-02 01:32:06 | 2021-08-29 20:54:38 |
| [gostradamus](https://github.com/bykof/gostradamus) | 156 | 2 | 0 | Gostradamus: Better DateTimes for Go 🕰️ | 2020-04-07 12:29:21 | 2021-08-07 20:45:11 |
| [go-persian-calendar](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 103 | 13 | 3 | The implementation of Persian (Solar Hijri) Calendar in Go | 2016-01-31 18:40:23 | 2021-08-31 19:15:21 |
| [iso8601](https://github.com/relvacode/iso8601) | 95 | 5 | 1 | A fast ISO8601 date parser for Go | 2017-04-25 15:54:18 | 2021-08-15 08:18:28 |
| [date](https://godoc.org/github.com/rickb777/date) | 75 | 16 | 5 | A Go package for working with dates | 2015-11-23 22:58:07 | 2021-08-31 15:28:37 |
| [timespan](https://github.com/SaidinWoT/timespan) | 74 | 9 | 3 | Golang package to manipulate time intervals. | 2014-10-07 03:57:02 | 2020-10-20 14:32:08 |
| [feiertage](https://github.com/wlbr/feiertage) | 40 | 4 | 1 | Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany) | 2015-11-04 14:19:27 | 2021-08-20 12:21:45 |
| [go-sunrise](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 37 | 5 | 0 | Go package for calculating the sunrise and sunset times for a given location | 2017-06-15 20:49:41 | 2021-09-06 06:44:28 |
| [go-str2duration](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 29 | 1 | 1 | Convert string to duration in golang | 2020-02-02 06:04:07 | 2021-08-24 06:30:42 |
| [kair](https://github.com/GuilhermeCaruso/kair) | 19 | 5 | 0 | :clock1: Date and Time - Golang Formatting Library | 2018-10-03 15:44:07 | 2021-04-23 00:08:00 |
| [cronrange](https://github.com/1set/cronrange) | 12 | 4 | 3 | time range expression in cron style | 2019-11-10 01:30:45 | 2021-08-13 02:05:08 |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 11 | 2 | 0 | Now is a time toolkit for golang | 2016-03-06 01:44:58 | 2020-08-28 21:19:40 |
| [tuesday](https://godoc.org/github.com/osteele/tuesday) | 9 | 1 | 1 | Ruby-compatible strftime for golang | 2017-08-10 20:46:26 | 2021-06-19 03:38:21 |
| [strftime](https://github.com/awoodbeck/strftime) | 7 | 2 | 0 | C99-compatible strftime formatter for use with Go time.Time instances. | 2018-02-10 00:35:46 | 2020-11-22 03:01:25 |
| [go-week](https://github.com/stoewer/go-week) | 6 | 3 | 2 | A Go package to work with ISO 8601 week dates | 2018-02-23 07:02:37 | 2021-07-18 01:15:08 |
</details>

### Distributed Systems
Packages that help with building Distributed Systems.

<sup>*Last Update: 2021-09-09 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [etcd](https://etcd.io) | 37,156 | 7,946 | 206 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2021-09-08 23:00:59 |
| [kit](https://gokit.io) | 21,182 | 2,175 | 48 | A standard library for microservices. | 2015-02-03 00:01:19 | 2021-09-09 01:53:44 |
| [go-micro](https://go-micro.dev) | 16,713 | 1,856 | 39 | A Go framework for distributed systems development | 2015-01-13 23:30:18 | 2021-09-09 01:54:46 |
| [grpc-go](https://grpc.io) | 14,444 | 3,151 | 109 | The Go language implementation of gRPC. HTTP/2 based RPC | 2014-12-08 18:59:34 | 2021-09-09 00:32:00 |
| [go-zero](https://go-zero.dev) | 10,759 | 1,326 | 30 | go-zero is a web and rpc framework written in Go. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. | 2020-08-07 15:37:57 | 2021-09-08 23:19:55 |
| [micro](https://micro.mu) | 10,355 | 917 | 85 | Micro is a cloud operating system | 2015-01-16 22:35:14 | 2021-09-08 19:44:04 |
| [nats-server](https://nats.io) | 9,791 | 957 | 147 | High-Performance server for NATS.io, the cloud and edge native messaging system. | 2012-10-29 16:12:24 | 2021-09-09 01:35:33 |
| [rpcx](https://rpcx.io) | 5,990 | 922 | 12 | Best microservices framework in Go, like alibaba Dubbo, but with more features, Scale easily. Try it. Test it. If you feel it's better, use it! 𝐉𝐚𝐯𝐚有𝐝𝐮𝐛𝐛𝐨, 𝐆𝐨𝐥𝐚𝐧𝐠有𝐫𝐩𝐜𝐱! | 2016-05-18 09:34:05 | 2021-09-09 02:22:13 |
| [raft](https://godoc.org/cirello.io/pglock) | 5,016 | 680 | 18 | Golang implementation of the Raft consensus protocol | 2013-11-05 00:41:20 | 2021-09-08 13:25:09 |
| [lura](https://luraproject.org) | 4,488 | 416 | 56 | Ultra performant API Gateway with middlewares. A project hosted at The Linux Foundation | 2016-11-04 18:37:13 | 2021-09-09 02:09:41 |
| [tendermint](https://tendermint.com/) | 4,303 | 1,376 | 331 | ⟁ Tendermint Core (BFT Consensus) in Go | 2014-05-14 23:21:35 | 2021-09-08 21:49:41 |
| [torrent](https://github.com/anacrolix/torrent) | 4,028 | 500 | 54 | Full-featured BitTorrent client package and utilities | 2015-01-08 21:10:42 | 2021-09-09 00:55:18 |
| [dragonboat](https://github.com/lni/dragonboat) | 3,820 | 387 | 24 | A feature complete and high performance multi-group Raft library in Go.   | 2018-12-23 07:02:04 | 2021-09-09 02:09:21 |
| [emitter](https://emitter.io) | 3,004 | 276 | 11 | High performance, distributed and low latency publish-subscribe platform. | 2016-10-29 08:52:21 | 2021-09-08 13:03:14 |
| [glow](https://github.com/chrislusf/glow) | 2,985 | 232 | 11 | Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant. | 2015-06-14 00:33:48 | 2021-09-08 06:37:15 |
| [gleam](https://github.com/chrislusf/gleam) | 2,883 | 257 | 36 | Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly. | 2016-08-26 08:44:48 | 2021-09-08 09:06:47 |
| [liftbridge](https://liftbridge.io) | 2,132 | 91 | 38 | Lightweight, fault-tolerant message streams. | 2017-10-13 19:50:26 | 2021-09-08 12:08:11 |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1,178 | 204 | 20 | Hprose is a cross-language RPC. This project is Hprose for Golang. | 2014-02-14 03:16:43 | 2021-09-05 07:05:55 |
| [ringpop-go](http://www.uber.com) | 697 | 61 | 25 | Scalable, fault-tolerant application-layer sharding for Go applications | 2015-06-05 22:48:53 | 2021-09-01 17:42:55 |
| [gorpc](https://github.com/valyala/gorpc) | 640 | 93 | 15 | Simple, fast and scalable golang rpc library for high load | 2014-11-20 17:02:37 | 2021-09-06 11:15:12 |
| [rain](https://put.io) | 622 | 40 | 1 | 🌧 BitTorrent client and library in Go | 2014-05-21 09:17:24 | 2021-09-05 03:37:45 |
| [go-health](https://github.com/InVisionApp/go-health) | 600 | 38 | 9 | Library for enabling asynchronous health checks in your service | 2017-11-29 21:00:07 | 2021-08-16 06:45:23 |
| [resgate](https://resgate.io) | 510 | 41 | 6 | A Realtime API Gateway used with NATS to build REST, real time, and RPC APIs, where all your clients are synchronized seamlessly. | 2018-02-22 12:06:26 | 2021-09-07 09:32:39 |
| [redislock](https://put.io) | 491 | 68 | 1 | Simplified distributed locking implementation using Redis | 2019-06-24 11:10:10 | 2021-09-06 16:08:19 |
| [go-sundheit](https://pdu.pub) | 425 | 22 | 3 | A library built to provide support for defining service health for golang services. It allows you to register async health checks for your dependencies and the service itself, provides a health endpoint that exposes their status, and health metrics. | 2019-04-08 12:54:01 | 2021-08-16 11:23:54 |
| [consistent](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 405 | 50 | 6 | Consistent hashing with bounded loads in Golang | 2018-03-25 15:38:27 | 2021-09-05 05:59:06 |
| [digota](http://digota.com) | 404 | 64 | 10 | ecommerce microservice | 2017-08-14 12:01:37 | 2021-09-02 06:23:05 |
| [arpc](https://github.com/lesismal/arpc) | 341 | 29 | 0 | More effective network communication, two-way calling, notify and broadcast supported. | 2020-05-19 11:30:05 | 2021-09-06 16:17:49 |
| [sleuth](http://ursiform.github.io/sleuth/) | 338 | 21 | 0 | A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services | 2016-04-23 14:21:41 | 2021-09-03 03:12:33 |
| [go-jump](https://github.com/dgryski/go-jump) | 333 | 27 | 1 | go-jump: Jump consistent hashing | 2014-06-15 22:12:04 | 2021-09-07 03:50:47 |
| [dht](https://github.com/anacrolix/dht) | 205 | 46 | 5 | dht is used by anacrolix/torrent, and is intended for use as a library in other projects both torrent related and otherwise | 2016-12-14 00:34:42 | 2021-09-09 00:54:43 |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 186 | 69 | 4 | A simple go implementation of json rpc 2.0 client over http | 2016-11-10 11:27:55 | 2021-09-08 22:58:21 |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 152 | 16 | 3 | The jsonrpc package helps implement of JSON-RPC 2.0 | 2016-10-28 13:36:59 | 2021-08-20 06:22:49 |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 67 | 8 | 1 | Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. | 2015-10-10 07:27:33 | 2021-06-13 12:55:46 |
| [dynamolock](https://godoc.org/cirello.io/dynamolock) | 65 | 33 | 1 | DynamoDB Lock Client for Go | 2018-07-08 11:13:00 | 2021-09-05 08:49:28 |
| [doublejump](https://github.com/edwingeng/doublejump) | 64 | 12 | 0 | A revamped Google's jump consistent hash | 2018-06-26 16:04:50 | 2021-08-30 02:15:41 |
| [dot](https://github.com/dotchain/dot) | 58 | 1 | 0 | distributed data sync with operational transformation/transforms  | 2017-12-18 01:08:12 | 2021-07-05 18:18:30 |
| [semaphore](https://jexia.github.io/semaphore/) | 56 | 13 | 9 | Take control of your data, connect with anything, and expose it anywhere through protocols such as HTTP, GraphQL, and gRPC. | 2020-02-05 16:39:39 | 2021-09-03 13:51:44 |
| [outboxer](https://github.com/italolelis/outboxer) | 55 | 8 | 11 | A library that implements the outboxer pattern in go | 2019-02-01 09:50:13 | 2021-08-31 15:59:15 |
| [flowgraph](https://emitter.io) | 41 | 4 | 0 | Flowgraph package for scalable asynchronous system development | 2018-08-29 21:45:26 | 2021-08-12 09:29:53 |
| [drmaa](https://github.com/dgruber/drmaa) | 33 | 15 | 0 | Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard. | 2013-03-17 12:58:02 | 2021-06-27 10:39:31 |
| [go-pdu](https://pdu.pub) | 29 | 3 | 0 | Parallel Digital Universe - A decentralized identity-based social network | 2018-10-08 08:13:22 | 2021-09-02 10:25:33 |
| [pglock](https://godoc.org/cirello.io/pglock) | 28 | 6 | 0 | [Mirror] PostgreSQL Lock Client for Go | 2018-12-17 17:43:41 | 2021-08-16 07:57:24 |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 28 | 1 | 2 | MySQL Backed Locking Primitive | 2020-06-06 16:30:07 | 2021-08-23 06:07:27 |
| [dynatomic](https://github.com/tylfin/dynatomic) | 14 | 1 | 0 | Dynatomic is a library for using dynamodb as an atomic counter | 2019-02-08 17:45:14 | 2020-12-25 16:56:32 |
| [micro](https://github.com/gmsec/micro) | 13 | 5 | 0 | A Go distributed systems development framework | 2020-05-03 01:16:16 | 2021-08-24 09:38:01 |
| [consistenthash](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 9 | 1 | 0 | A Go library that implements Consistent Hashing | 2020-04-22 16:01:25 | 2021-04-09 18:54:48 |
</details>

### Dynamic DNS
Tools for updating dynamic DNS records.

<sup>*Last Update: 2023-10-02 20:50:43*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [godns](https://xiaozhou.net/godns-project-2014-05-18.html) | 1,349 | 233 | 14 | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 2014-05-11 11:49:17 | 2023-10-02 04:29:24 |
| [ddns](https://github.com/skibish/ddns) | 239 | 22 | 1 | Personal DDNS client with Digital Ocean Networking DNS as backend. | 2017-03-13 21:02:27 | 2023-09-24 14:01:53 |
</details>

### Email
Libraries and tools that implement email creation and sending.

<sup>*Last Update: 2021-09-17 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [MailHog](https://github.com/mailhog/MailHog) | 8,987 | 685 | 192 | Web and API based SMTP testing | 2014-04-16 22:28:49 | 2021-09-16 14:34:38 |
| [hermes](https://github.com/matcornic/hermes) | 2,335 | 189 | 27 | Golang package that generates clean, responsive HTML e-mails for sending transactional mail | 2017-03-25 18:25:36 | 2021-09-15 22:52:33 |
| [email](https://github.com/jordan-wright/email) | 1,817 | 250 | 49 | Robust and flexible email library for Go | 2013-12-12 20:11:59 | 2021-09-16 05:52:55 |
| [go-imap](https://github.com/emersion/go-imap) | 1,349 | 186 | 64 |  :inbox_tray: An IMAP library for clients and servers | 2016-04-26 17:59:18 | 2021-09-14 11:32:51 |
| [sendgrid-go](https://sendgrid.com) | 755 | 233 | 12 | The Official Twilio SendGrid Led, Community Driven Golang API Library | 2013-09-12 03:31:13 | 2021-09-16 21:05:57 |
| [mailgun-go](https://mailchain.xyz) | 539 | 124 | 3 | Go library for sending mail with the Mailgun API. | 2014-02-28 00:28:44 | 2021-09-11 03:57:07 |
| [chasquid](https://blitiri.com.ar/p/chasquid/) | 467 | 30 | 1 | SMTP (email) server with a focus on simplicity, security, and ease of operation [mirror] | 2016-11-03 01:28:05 | 2021-09-12 03:45:49 |
| [email-verifier](https://github.com/AfterShip/email-verifier) | 293 | 32 | 2 | :white_check_mark: A Go library for email verification without sending any emails. | 2020-12-18 08:47:28 | 2021-09-11 17:34:50 |
| [go-message](https://github.com/emersion/go-message) | 211 | 66 | 20 | :envelope: A streaming Go library for the Internet Message Format and mail messages | 2016-12-31 09:31:52 | 2021-08-23 16:36:23 |
| [hectane](https://github.com/hectane/hectane) | 209 | 24 | 16 | Lightweight SMTP client written in Go | 2015-08-28 01:36:47 | 2021-09-08 22:04:46 |
| [douceur](https://github.com/aymerick/douceur) | 197 | 34 | 9 | A simple CSS parser and inliner in Go | 2015-04-09 10:21:26 | 2021-07-21 09:00:38 |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 196 | 36 | 8 | Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP. | 2019-09-15 05:38:54 | 2021-09-09 01:50:18 |
| [mailchain](https://mailchain.xyz) | 86 | 41 | 43 | Using Mailchain, blockchain users can now send and receive rich-media HTML messages with attachments via a blockchain address. | 2019-04-11 17:37:31 | 2021-09-06 15:07:57 |
| [go-premailer](https://github.com/vanng822/go-premailer) | 72 | 12 | 3 | Inline styling for html mail in golang | 2015-02-16 22:19:18 | 2021-08-22 05:01:15 |
| [go-dkim](https://github.com/toorop/go-dkim) | 71 | 30 | 4 | DKIM package for golang | 2015-04-29 15:38:27 | 2021-08-06 08:24:11 |
| [smtp](https://sendgrid.com) | 64 | 24 | 6 | MailHog SMTP Protocol | 2014-12-24 16:13:49 | 2021-07-20 15:08:03 |
| [go-email-validator](https://github.com/go-email-validator/go-email-validator) | 13 | 4 | 3 | 📧 Golang Email address validator | 2020-12-10 18:27:20 | 2021-09-09 13:28:20 |
</details>

### Embeddable Scripting Languages
Embedding other languages inside your go code.

<sup>*Last Update: 2021-09-15 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 4,289 | 477 | 88 | GopherLua: VM and compiler for Lua in Go | 2015-02-15 13:23:37 | 2021-09-14 19:22:54 |
| [tengo](https://tengolang.com) | 2,422 | 141 | 52 | A fast script language for Go | 2019-01-09 07:17:17 | 2021-09-14 19:04:40 |
| [goja](https://github.com/dop251/goja) | 2,386 | 201 | 17 | ECMAScript/JavaScript engine in pure Go | 2016-11-04 22:04:06 | 2021-09-14 18:41:08 |
| [go-lua](https://github.com/Shopify/go-lua) | 2,152 | 156 | 39 | A Lua VM in Go | 2013-12-20 17:29:43 | 2021-09-14 00:52:06 |
| [expr](https://github.com/antonmedv/expr) | 1,955 | 167 | 34 | Expression language for Go | 2018-07-14 15:57:34 | 2021-09-13 09:27:46 |
| [go-python](https://github.com/sbinet/go-python) | 1,303 | 128 | 25 | naive go bindings to the CPython2 C-API | 2012-07-09 15:43:31 | 2021-09-14 19:07:19 |
| [anko](http://play-anko.appspot.com/) | 1,167 | 111 | 19 | Scriptable interpreter written in golang | 2014-03-28 07:29:40 | 2021-09-14 19:19:05 |
| [cel-go](https://opensource.google.com/projects/cel) | 890 | 101 | 33 | Fast, portable, non-Turing complete expression evaluation with gradual typing (Go) | 2018-03-09 22:57:58 | 2021-09-14 13:23:25 |
| [go-php](https://github.com/deuill/go-php) | 810 | 92 | 20 | PHP bindings for the Go programming language (Golang) | 2015-09-17 21:23:52 | 2021-09-09 09:07:41 |
| [go-duktape](https://github.com/olebedev/go-duktape) | 778 | 90 | 7 | Duktape JavaScript engine bindings for Go | 2015-01-08 05:09:05 | 2021-09-12 04:51:57 |
| [golua](https://github.com/aarzilli/golua) | 555 | 157 | 6 | Go bindings for Lua C API - in progress | 2010-12-06 21:39:53 | 2021-09-07 08:52:58 |
| [gisp](https://docs.gentee.org) | 466 | 34 | 1 | Simple LISP in Go | 2014-01-11 14:05:43 | 2021-09-07 20:51:10 |
| [gval](https://github.com/PaesslerAG/gval) | 366 | 46 | 5 | Expression evaluation in golang | 2017-09-27 08:32:49 | 2021-09-14 12:13:55 |
| [gentee](https://docs.gentee.org) | 81 | 8 | 0 | Gentee - script programming language for automation. It uses VM and compiler written in Go (Golang). | 2018-01-14 15:49:05 | 2021-09-01 12:47:37 |
| [binder](https://github.com/alexeyco/binder) | 52 | 7 | 0 | High level go to Lua binder. Write less, do more. | 2017-04-02 17:14:52 | 2021-07-26 07:36:34 |
| [purl](https://github.com/ian-kent/purl) | 33 | 2 | 2 | Perl, but fluffy like a cat! | 2014-11-29 19:06:01 | 2021-07-31 20:20:45 |
| [ngaro](https://github.com/db47h/ngaro) | 20 | 1 | 1 | An embeddable implementation of the Ngaro Virtual Machine for Go programs | 2016-08-09 15:23:50 | 2020-04-06 11:44:00 |
| [ecal](https://github.com/krotik/ecal) | 11 | 0 | 0 | A simple embeddable scripting language which supports concurrent event processing. | 2020-11-30 15:58:56 | 2021-07-26 16:59:56 |
</details>

### Error Handling
Libraries for handling errors.

<sup>*Last Update: 2021-09-03 09:25:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [errors](https://godoc.org/github.com/pkg/errors) | 7,234 | 555 | 42 | Simple error handling primitives | 2015-12-27 12:05:38 | 2021-09-03 01:16:10 |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 1,342 | 88 | 13 | A Go (golang) package for representing a list of errors as a single error. | 2014-12-15 20:12:26 | 2021-09-02 19:10:20 |
| [eris](https://github.com/rotisserie/eris) | 833 | 20 | 0 | eris provides a better way to handle, trace, and log errors in Go 🎆 | 2019-09-07 16:50:33 | 2021-08-30 02:14:27 |
| [errorx](https://github.com/joomcode/errorx) | 760 | 25 | 4 | A comprehensive error handling library for Go | 2018-08-17 08:02:10 | 2021-08-23 09:19:07 |
| [tracerr](https://github.com/ztrue/tracerr) | 679 | 26 | 1 | Golang errors with stack trace and source fragments. | 2019-02-06 18:57:46 | 2021-08-28 12:23:24 |
| [errlog](https://github.com/snwfdhmp/errlog) | 398 | 16 | 1 | Reduce debugging time while programming Go. Use static and stack-trace analysis to determine which func call causes the error. | 2019-02-16 23:19:05 | 2021-07-30 10:15:58 |
| [emperror](https://emperror.dev/emperror) | 221 | 14 | 5 | The Emperor takes care of all errors personally | 2017-06-13 00:24:28 | 2021-08-30 12:52:33 |
| [errors](https://emperror.dev/errors) | 98 | 5 | 6 | Drop-in replacement for the standard library errors package and github.com/pkg/errors | 2019-07-09 13:02:52 | 2021-09-02 18:39:00 |
| [errors](https://github.com/bnkamalesh/errors) | 23 | 3 | 0 | A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types. | 2020-07-17 18:57:04 | 2021-06-19 13:37:32 |
| [falcon](https://github.com/SonicRoshan/falcon) | 7 | 0 | 0 | A Simple Yet Highly Powerful Package For Error Handling | 2019-09-09 12:49:43 | 2021-08-01 23:06:50 |
| [errors](https://github.com/neuronlabs/errors) | 3 | 0 | 0 | Simple golang error handling with classification primitives. | 2019-07-26 00:15:36 | 2020-01-23 14:57:00 |
| [errors](https://github.com/PumpkinSeed/errors) | 3 | 0 | 0 | Simple and efficient error package  | 2020-01-08 21:12:51 | 2021-08-01 23:06:07 |
</details>

### File Handling
Libraries for handling files and file systems.

<sup>*Last Update: 2021-10-07 09:25:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [afero](https://github.com/spf13/afero) | 3,995 | 363 | 93 | A FileSystem Abstraction System for Go | 2014-10-28 14:19:05 | 2021-10-03 16:51:23 |
| [pdfcpu](https://pdfcpu.io) | 2,612 | 194 | 32 | A PDF processor written in Go. | 2017-06-18 17:27:38 | 2021-10-07 02:23:25 |
| [notify](https://github.com/rjeczalik/notify) | 685 | 104 | 40 | File system event notification library on steroids. | 2014-09-08 16:09:34 | 2021-10-06 21:36:16 |
| [copy](https://pkg.go.dev/github.com/otiai10/copy) | 341 | 77 | 7 | Go copy directory recursively | 2017-09-01 03:18:56 | 2021-10-02 09:09:53 |
| [bigfile](https://learnku.com/docs/bigfile) | 206 | 38 | 1 | Bigfile -- a file transfer system that supports http, rpc and ftp protocol   https://bigfile.site   | 2019-07-15 10:43:50 | 2021-09-16 02:25:05 |
| [afs](https://github.com/viant/afs) | 156 | 14 | 1 | Abstract File Storage | 2019-08-19 18:43:38 | 2021-10-05 04:37:45 |
| [vfs](https://github.com/C2FO/vfs) | 127 | 11 | 12 | Pluggable, extensible virtual file system for Go | 2017-08-01 18:06:14 | 2021-10-06 14:42:38 |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 91 | 18 | 0 | Read csv file from go using tags | 2017-06-18 15:31:16 | 2021-10-04 17:39:57 |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 82 | 20 | 3 | Golang wrapper for Exiftool : extract as much metadata as possible (EXIF, ...) from files (pictures, pdf, office documents, ...) | 2019-05-12 20:34:09 | 2021-10-05 04:49:02 |
| [opc](https://github.com/qmuntal/opc) | 67 | 4 | 0 | Go implementation of the Open Packaging Conventions (OPC) | 2018-11-06 14:49:06 | 2021-09-24 13:20:30 |
| [skywalker](https://github.com/dixonwille/skywalker) | 65 | 5 | 1 | A package to allow one to concurrently go through a filesystem with ease | 2017-08-01 20:08:25 | 2021-08-31 17:22:11 |
| [tarfs](https://github.com/posener/tarfs) | 47 | 6 | 1 | An implementation of the FileSystem interface for tar files. | 2017-03-10 22:13:11 | 2021-06-30 11:02:38 |
| [checksum](https://github.com/codingsince1985/checksum) | 42 | 11 | 0 | Compute message digest for large files in Go | 2014-11-05 09:37:00 | 2021-09-29 11:29:41 |
| [parquet](https://github.com/parsyl/parquet) | 34 | 3 | 2 | A library for reading and writing parquet files. | 2019-01-29 21:52:30 | 2021-10-03 06:00:07 |
| [baraka](https://github.com/xis/baraka) | 33 | 5 | 1 | a tool for handling file uploads simple | 2020-07-12 21:56:50 | 2021-08-22 11:11:59 |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 28 | 14 | 0 | Load GTFS files in golang | 2017-07-09 09:30:31 | 2021-09-26 10:25:42 |
| [flop](https://github.com/homedepot/flop) | 26 | 5 | 1 | Go file operations library chasing GNU APIs. | 2019-03-01 13:41:39 | 2021-09-19 13:58:40 |
| [gut](https://github.com/1set/gut) | 18 | 4 | 13 | 🍱 yet another collection of go utilities & tools | 2019-10-05 23:47:24 | 2021-08-20 02:09:32 |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 14 | 4 | 1 | copy files for humans | 2018-10-16 07:08:24 | 2021-03-12 20:07:53 |
| [todotxt](https://github.com/1set/todotxt) | 9 | 0 | 0 | Parser for todo.txt files in Go ✅ | 2020-11-06 17:41:59 | 2021-08-06 11:59:04 |
| [higgs](https://github.com/dastoori/higgs) | 5 | 2 | 2 | A tiny cross-platform Go library to hide/unhide files and directories | 2020-12-13 18:33:10 | 2021-07-19 13:23:35 |
</details>

### Financial
Packages for accounting and finance.

<sup>*Last Update: 2021-09-08 15:02:27*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [decimal](https://github.com/shopspring/decimal) | 3,286 | 389 | 67 | Arbitrary-precision fixed-point decimal numbers in go | 2015-02-25 20:12:57 | 2021-09-08 07:10:20 |
| [go-money](https://github.com/Rhymond/go-money) | 985 | 86 | 21 | Go implementation of Fowler's Money pattern | 2017-03-20 16:23:54 | 2021-08-27 11:51:16 |
| [accounting](https://github.com/leekchan/accounting) | 673 | 51 | 6 | money and currency formatting for golang | 2015-08-10 13:23:56 | 2021-09-06 05:37:12 |
| [go-finance](https://github.com/FlashBoys/go-finance) | 534 | 49 | 4 | :warning: Deprecrated in favor of https://github.com/piquette/finance-go  | 2016-02-28 00:37:46 | 2021-06-23 04:00:30 |
| [techan](https://godoc.org/github.com/sdcoffey/techan) | 515 | 81 | 15 | Technical Analysis Library for Golang | 2017-03-08 03:04:08 | 2021-09-07 23:28:26 |
| [currency](https://pkg.go.dev/github.com/bojanz/currency) | 262 | 10 | 3 | Currency handling for Go. | 2020-04-16 15:34:39 | 2021-09-04 14:50:45 |
| [orderbook](https://github.com/i25959341/orderbook) | 216 | 74 | 5 | Matching Engine for Limit Order Book in Golang | 2018-04-24 18:05:26 | 2021-09-05 03:02:13 |
| [go-finance](https://github.com/alpeb/go-finance) | 110 | 17 | 0 | Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. | 2017-06-01 15:58:33 | 2021-09-06 11:35:17 |
| [transaction](https://github.com/claygod/transaction) | 95 | 14 | 0 | Embedded database for accounts transactions. | 2017-10-11 13:50:30 | 2021-07-21 18:02:01 |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 92 | 21 | 0 | Golang library for querying and parsing OFX | 2015-11-08 13:56:53 | 2021-09-01 13:35:48 |
| [vat](https://github.com/dannyvankooten/vat) | 84 | 11 | 4 | Go package for dealing with EU VAT. Does VAT number validation & rates retrieval. | 2016-06-18 16:10:09 | 2021-07-18 01:45:12 |
| [sleet](https://github.com/BoltApp/sleet) | 67 | 8 | 3 | Payment abstraction library - one interface for multiple payment processors ( inspired by Ruby's ActiveMerchant ) | 2019-11-13 21:56:58 | 2021-09-01 22:37:09 |
| [go-finnhub](https://github.com/m1/go-finnhub) | 60 | 13 | 0 | Simple and easy to use client for stock market, forex and crypto data from finnhub.io written in Go. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges | 2020-01-13 20:47:13 | 2021-08-03 00:48:39 |
| [currency](https://github.com/bnkamalesh/currency) | 42 | 5 | 0 | A currency computations package. | 2017-05-09 06:06:38 | 2021-09-02 00:28:57 |
| [fastme](https://github.com/newity/fastme) | 24 | 4 | 0 | Arbitrary-precision fixed-point decimal numbers in go | 2020-10-29 13:57:10 | 2021-08-10 04:48:11 |
| [go-finance](https://www.yellowduck.be) | 5 | 2 | 0 | Finance related Go functions (e.g. exchange rates, VAT number checking, …) | 2019-09-30 06:49:07 | 2021-03-08 22:46:44 |
</details>

### Forms
Libraries for working with forms.

<sup>*Last Update: 2021-09-07 09:25:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [nosurf](http://godoc.org/github.com/justinas/nosurf) | 1,208 | 98 | 9 | CSRF protection middleware for Go. | 2013-08-22 17:47:34 | 2021-09-05 19:24:16 |
| [binding](http://mholt.github.io/binding) | 783 | 79 | 8 | Reflectionless data binding for Go's net/http (not actively maintained) | 2014-05-20 23:35:14 | 2021-08-27 10:44:32 |
| [csrf](https://github.com/gorilla/csrf) | 691 | 102 | 1 | gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services 🔒 | 2015-08-03 00:35:16 | 2021-09-06 01:20:00 |
| [form](https://github.com/go-playground/form) | 493 | 28 | 8 | :steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. | 2016-05-26 13:26:40 | 2021-08-27 18:50:44 |
| [conform](https://github.com/leebenson/conform) | 229 | 29 | 1 | Trims, sanitizes & scrubs data based on struct tags (go, golang) | 2016-01-05 18:00:06 | 2021-08-15 11:02:11 |
| [formam](https://github.com/monoculum/formam) | 164 | 14 | 2 | a package for decode form's values into struct in Go | 2014-10-25 00:23:30 | 2021-08-24 18:57:04 |
| [forms](https://github.com/albrow/forms) | 122 | 16 | 2 | A lightweight go library for parsing form data or json from an http.Request. | 2014-08-07 16:11:30 | 2021-07-01 09:59:08 |
| [qs](https://github.com/sonh/qs) | 58 | 0 | 0 | Go module for encoding structs into URL query parameters | 2020-10-02 09:50:35 | 2021-09-02 15:56:05 |
| [bind](https://github.com/robfig/bind) | 26 | 3 | 0 |  | 2014-08-06 00:13:10 | 2021-08-01 23:12:30 |
| [queryparam](https://github.com/TomWright/queryparam) | 10 | 5 | 0 | Go package to easily convert a URL's query parameters/values into usable struct values of the correct types. | 2018-06-14 10:23:05 | 2021-07-25 13:31:51 |
</details>

### Functional
Packages to support functional programming in Go.

<sup>*Last Update: 2021-10-11 09:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1,214 | 66 | 4 |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  | 2014-07-02 10:27:16 | 2021-10-06 16:57:01 |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 197 | 13 | 0 | Monad, Functional Programming features for Golang | 2018-05-24 09:08:45 | 2021-10-06 08:06:37 |
| [fuego](https://github.com/seborama/fuego) | 98 | 8 | 0 | Functional Experiment in Golang | 2018-11-05 22:24:09 | 2021-09-23 15:22:13 |
| [gofp](https://github.com/rbrahul/gofp) | 88 | 3 | 0 | A super simple Lodash like utility library with essential functions that empowers the development in Go | 2021-02-19 00:01:39 | 2021-09-13 09:30:59 |
</details>

### GUI - Interaction
Libraries for building GUI Applications.

<sup>*Last Update: 2021-11-25 13:59:06*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [robotgo](https://github.com/go-vgo/robotgo) | 7,046 | 651 | 82 | RobotGo, Go Native cross-platform GUI automation  @vcaesar | 2016-09-26 16:26:56 | 2021-11-25 06:16:01 |
| [systray](https://github.com/getlantern/systray) | 2,037 | 246 | 84 | a cross platfrom Go library to place an icon and menu in the notification area | 2014-11-12 03:41:57 | 2021-11-24 13:04:29 |
| [gosx-notifier](https://github.com/deckarep/gosx-notifier) | 552 | 44 | 10 | gosx-notifier is a Go framework for sending desktop notifications to OSX 10.8 or higher | 2013-11-25 06:35:16 | 2021-11-15 05:26:55 |
| [trayhost](https://github.com/shurcooL/trayhost) | 220 | 18 | 9 | Cross-platform Go library to place an icon in the host operating system's taskbar. | 2014-04-25 04:02:30 | 2021-11-07 04:56:31 |
| [go-appindicator](https://github.com/dawidd6/go-appindicator) | 17 | 3 | 1 | :traffic_light: Go bindings for libappindicator3 C library | 2019-05-04 13:38:53 | 2021-06-15 18:21:55 |
| [activity-tracker](https://github.com/prashantgupta24/activity-tracker) | 12 | 1 | 2 | A library to notify about any (pluggable) activity on your machine, and let you take action as needed | 2019-03-12 21:16:44 | 2021-05-08 17:14:34 |
| [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) | 10 | 2 | 0 | macOS Sleep/ Wake notifications in golang | 2019-03-30 00:43:23 | 2021-08-12 18:18:14 |
</details>

### GUI - Toolkits
Libraries for building GUI Applications.

<sup>*Last Update: 2021-11-25 13:59:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fyne](https://fyne.io/) | 14,660 | 788 | 307 | Cross platform GUI in Go inspired by Material Design | 2018-02-04 22:07:16 | 2021-11-25 06:30:54 |
| [webview](https://github.com/webview/webview) | 9,120 | 714 | 275 | Tiny cross-platform webview library for C/C++/Golang. Uses WebKit (Gtk/Cocoa) and Edge (Windows) | 2017-08-19 08:26:00 | 2021-11-25 02:21:14 |
| [qt](https://github.com/therecipe/qt) | 8,919 | 704 | 341 | Qt binding for Go (Golang) with support for Windows / macOS / Linux / FreeBSD / Android / iOS / Sailfish OS / Raspberry Pi / AsteroidOS / Ubuntu Touch / JavaScript / WebAssembly | 2014-11-19 00:03:08 | 2021-11-25 06:21:49 |
| [ui](https://github.com/andlabs/ui) | 8,022 | 727 | 120 | Platform-native GUI library for Go. | 2014-02-17 23:44:00 | 2021-11-23 03:55:05 |
| [walk](https://github.com/lxn/walk) | 5,763 | 791 | 319 | A Windows GUI toolkit for the Go Programming Language | 2010-09-16 08:11:49 | 2021-11-24 01:09:15 |
| [go-app](https://go-app.dev) | 5,537 | 259 | 38 | A package to build progressive web apps with Go programming language and WebAssembly. | 2016-10-12 00:31:33 | 2021-11-24 05:15:33 |
| [wails](https://wails.app) | 4,778 | 235 | 32 | Create desktop apps using Go and Web Technologies. | 2018-12-15 23:14:06 | 2021-11-25 03:10:53 |
| [go-astilectron](https://github.com/asticode/go-astilectron) | 4,118 | 293 | 46 | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron) | 2017-04-22 07:59:15 | 2021-11-25 03:22:33 |
| [go-sciter](https://sciter.com) | 2,268 | 243 | 81 | Golang bindings of Sciter: the Embeddable HTML/CSS/script engine for modern UI development | 2015-10-15 12:41:06 | 2021-11-25 05:59:17 |
| [go-gtk](http://mattn.github.com/go-gtk) | 1,823 | 245 | 68 | Go binding for GTK | 2009-11-26 16:58:53 | 2021-11-19 00:17:40 |
| [gotk3](https://github.com/gotk3/gotk3) | 1,643 | 204 | 82 | Go bindings for GTK3 | 2015-08-13 13:09:46 | 2021-11-24 15:31:52 |
| [gowd](https://github.com/dtylman/gowd) | 334 | 40 | 3 | Build cross platform GUI apps with GO and HTML/JS/CSS (powered by nwjs) | 2017-03-29 14:50:53 | 2021-11-12 17:48:43 |
</details>

### Game Development
Awesome game development libraries.

<sup>*Last Update: 2021-09-07 09:25:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ebiten](https://ebiten.org) | 5,018 | 332 | 224 | A dead simple 2D game library for Go | 2013-06-16 15:13:01 | 2021-09-06 16:19:56 |
| [leaf](https://github.com/name5566/leaf) | 4,129 | 1,105 | 12 | A game server framework in Go (golang) | 2014-08-04 12:40:08 | 2021-09-06 03:17:12 |
| [pixel](https://github.com/faiface/pixel) | 3,627 | 211 | 36 | A hand-crafted 2D game library in Go | 2016-11-19 11:15:34 | 2021-09-07 01:02:08 |
| [goworld](https://github.com/xiaonanln/goworld) | 1,933 | 356 | 16 | Scalable Distributed Game Server Engine with Hot Swapping in Golang | 2017-06-03 15:02:46 | 2021-09-06 09:22:21 |
| [nano](https://github.com/lonng/nano) | 1,802 | 298 | 18 | Lightweight, facility, high performance golang based game server framework | 2017-08-02 06:05:14 | 2021-09-04 02:15:13 |
| [go-sdl2](https://godoc.org/github.com/veandco/go-sdl2) | 1,612 | 194 | 50 | SDL2 binding for Go | 2013-06-05 18:30:03 | 2021-09-06 23:32:26 |
| [engine](https://discord.gg/NfaeVr8zDg) | 1,558 | 156 | 67 | Go 3D Game Engine (http://g3n.rocks) | 2017-03-07 18:25:09 | 2021-09-07 00:05:06 |
| [engo](https://engoengine.github.io) | 1,416 | 118 | 52 | Engo is an open-source 2D game engine written in Go. | 2014-11-12 05:50:03 | 2021-09-03 06:56:55 |
| [termloop](https://github.com/JoelOtter/termloop) | 1,227 | 74 | 5 | Terminal-based game engine for Go, built on top of Termbox | 2015-05-23 17:12:34 | 2021-08-31 07:02:30 |
| [pitaya](https://github.com/topfreegames/pitaya) | 1,158 | 236 | 18 | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. | 2018-03-19 19:40:36 | 2021-09-06 11:33:32 |
| [gonet](https://github.com/xtaci/gonet) | 1,149 | 300 | 0 | A Game Server Skeleton in golang. | 2013-04-11 02:18:23 | 2021-08-31 01:56:29 |
| [oak](https://github.com/oakmound/oak) | 923 | 60 | 12 | A pure Go game engine | 2017-07-15 16:24:27 | 2021-09-07 01:42:48 |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 647 | 69 | 8 | Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming. | 2017-01-27 08:31:45 | 2021-09-06 16:07:01 |
| [engine](https://azul3d.org) | 500 | 44 | 82 | Azul3D - A 3D game engine written in Go! | 2016-02-29 04:54:44 | 2021-09-03 17:21:44 |
| [go-astar](https://discord.gg/NfaeVr8zDg) | 445 | 60 | 1 | Go implementation of the A* search algorithm | 2014-05-28 02:00:03 | 2021-09-01 21:26:51 |
| [go3d](https://github.com/ungerik/go3d) | 206 | 36 | 2 | A performance oriented 2D/3D math package for Go | 2011-06-27 13:02:26 | 2021-09-02 02:56:12 |
| [prototype](https://github.com/gonutz/prototype) | 60 | 4 | 2 | Simple 2D game prototyping framework. | 2015-03-04 09:24:39 | 2021-08-17 11:58:11 |
| [tile](https://github.com/kelindar/tile) | 31 | 2 | 0 | Tile is a 2D grid engine, built with data and cache friendly ways, includes pathfinding and observers. | 2020-08-19 13:23:18 | 2021-08-23 14:57:44 |
</details>

### Generation and Generics
Tools to enhance the language with features like generics via code generation.

<sup>*Last Update: 2021-08-25 09:25:14*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-linq](https://godoc.org/github.com/ahmetb/go-linq) | 2,643 | 186 | 5 | .NET LINQ capabilities in Go | 2013-12-19 03:05:00 | 2021-08-25 01:43:39 |
| [jennifer](https://github.com/dave/jennifer) | 2,117 | 103 | 16 | Jennifer is a code generator for Go | 2016-12-04 20:57:38 | 2021-08-24 18:18:37 |
| [gen](http://clipperhouse.com/gen/overview/) | 1,310 | 83 | 32 | Type-driven code generation for Go | 2013-10-13 20:26:36 | 2021-08-14 03:46:15 |
| [goderive](https://github.com/awalterschulze/goderive) | 911 | 36 | 14 | Code Generation for Functional Programming, Concurrency and Generics in Golang | 2017-02-10 21:46:49 | 2021-08-18 09:54:18 |
| [gowrap](https://github.com/hexdigest/gowrap) | 530 | 46 | 2 | GoWrap is a command line tool for generating decorators for Go interfaces | 2018-09-15 09:20:42 | 2021-08-13 02:02:07 |
| [interfaces](https://github.com/rjeczalik/interfaces) | 306 | 19 | 5 | Code generation tools for Go. | 2015-12-06 00:04:50 | 2021-08-19 05:45:17 |
| [go-enum](https://github.com/abice/go-enum) | 225 | 27 | 2 | An enum generator for go | 2017-08-10 22:07:31 | 2021-08-18 17:28:02 |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 98 | 15 | 0 | A Go preprocessor for package scoped reflection | 2012-09-03 07:53:00 | 2021-08-18 09:29:09 |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 49 | 9 | 1 |  | 2016-11-18 11:38:54 | 2021-07-01 22:11:48 |
| [gotype](https://github.com/wzshiming/gotype) | 36 | 6 | 0 | Golang source code parsing, usage like reflect package | 2017-12-05 04:09:47 | 2021-08-24 21:57:55 |
| [GENERIS](https://github.com/senselogic/GENERIS) | 29 | 0 | 0 | Versatile Go code generator. | 2019-03-10 19:33:31 | 2021-08-07 07:09:43 |
| [go-xray](https://godoc.org/github.com/ahmetb/go-linq) | 18 | 1 | 0 | Helpers for making the use of reflection easier | 2019-10-01 08:40:51 | 2021-06-27 05:07:20 |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 11 | 0 | 0 | create type dynamically in Golang | 2020-01-14 15:50:38 | 2021-06-14 08:05:42 |
</details>

### Geographic
Geographic tools and servers

<sup>*Last Update: 2021-09-11 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tile38](https://tile38.com) | 7,662 | 469 | 116 | Real-time Geospatial and Geofencing | 2016-03-04 23:07:44 | 2021-09-11 01:06:57 |
| [geo](https://github.com/golang/geo) | 1,269 | 143 | 10 | S2 geometry library in Go | 2014-12-03 23:02:15 | 2021-09-04 18:07:39 |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 271 | 50 | 12 | Basic Go server for mbtiles | 2014-11-01 04:12:14 | 2021-09-10 08:12:13 |
| [osm](https://github.com/paulmach/osm) | 172 | 28 | 2 | General purpose library for reading, writing and working with OpenStreetMap data | 2016-02-02 00:59:03 | 2021-09-07 08:48:12 |
| [wgs84](https://github.com/wroge/wgs84) | 64 | 4 | 0 | A pure Go package for coordinate transformations. | 2019-06-08 17:17:59 | 2021-09-02 09:55:41 |
| [geoserver](https://github.com/hishamkaram/geoserver) | 59 | 13 | 2 | geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API. | 2018-03-26 21:36:49 | 2021-09-01 14:16:41 |
| [gismanager](https://github.com/hishamkaram/gismanager) | 38 | 6 | 0 | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver | 2018-09-29 12:51:37 | 2021-03-02 15:54:38 |
| [pbf](https://github.com/maguro/pbf) | 26 | 3 | 1 | OpenStreetMap PBF golang parser | 2017-09-18 23:13:18 | 2021-09-04 12:33:25 |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 13 | 3 | 1 | Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels | 2020-03-27 09:47:32 | 2021-08-23 13:10:07 |
</details>

### Go Compilers
Tools for compiling Go to other languages.

<sup>*Last Update: 2021-07-31 09:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 10,352 | 478 | 215 | A compiler from Go to JavaScript for running Go code in a browser | 2013-08-27 22:23:58 | 2021-07-30 10:40:34 |
| [tardisgo](http://tardisgo.github.io) | 409 | 27 | 4 | Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   | 2014-01-08 11:07:33 | 2021-07-09 14:21:34 |
| [c4go](https://github.com/Konstantin8105/c4go) | 276 | 33 | 23 | Transpiling C code to Go code | 2018-03-28 06:24:57 | 2021-07-25 18:05:20 |
| [f4go](https://github.com/Konstantin8105/f4go) | 25 | 6 | 4 | Transpiling fortran code to golang code | 2018-07-08 17:05:43 | 2021-07-19 14:10:36 |
</details>

### Goroutines
Tools for managing and working with Goroutines.

<sup>*Last Update: 2021-09-17 09:25:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ants](https://ants.andypan.me) | 6,435 | 795 | 17 | 🐜🐜🐜 ants is a high-performance and low-cost goroutine pool in Go, inspired by fasthttp./ ants 是一个高性能且低损耗的 goroutine 池。 | 2018-05-19 01:13:38 | 2021-09-16 20:57:09 |
| [tunny](https://github.com/Jeffail/tunny) | 2,608 | 225 | 1 | A goroutine pool for Go | 2014-04-02 16:14:58 | 2021-09-16 21:00:00 |
| [goworker](https://www.goworker.org) | 2,605 | 239 | 32 | goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers. | 2013-07-22 17:04:27 | 2021-09-15 02:59:50 |
| [grpool](https://godoc.org/github.com/ivpusic/grpool) | 653 | 96 | 4 | Lightweight Goroutine pool | 2015-07-22 00:15:04 | 2021-09-10 08:11:01 |
| [pool](https://github.com/go-playground/pool) | 637 | 59 | 4 | :speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation | 2015-10-28 16:36:08 | 2021-09-16 10:33:38 |
| [workerpool](https://github.com/gammazero/workerpool) | 630 | 80 | 4 | Concurrency limiting goroutine pool | 2016-05-17 14:32:06 | 2021-09-07 10:12:21 |
| [gowp](https://xxjwxc.github.io/post/gowp/) | 324 | 52 | 0 | golang worker pool , Concurrency limiting goroutine pool | 2019-09-14 11:43:50 | 2021-09-14 07:57:27 |
| [pond](https://github.com/alitto/pond) | 275 | 15 | 0 | 🔘 Minimalistic and High-performance goroutine worker pool written in Go | 2020-03-21 14:56:33 | 2021-09-16 12:39:01 |
| [go-floc](https://github.com/workanator/go-floc) | 209 | 16 | 0 | Floc: Orchestrate goroutines with ease. | 2017-07-03 07:34:06 | 2021-09-15 04:46:14 |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 173 | 17 | 1 | Simply way to control goroutines execution order based on dependencies | 2016-09-25 14:46:09 | 2021-09-16 10:25:29 |
| [semaphore](https://github.com/marusama/semaphore) | 124 | 7 | 0 | Fast resizable golang semaphore primitive | 2017-11-22 14:00:58 | 2021-09-15 05:48:13 |
| [go-workers](https://github.com/catmullet/go-workers/wiki/Getting-Started) | 118 | 6 | 0 | 👷 Library for safely running groups of workers concurrently or consecutively that require input and output through channels | 2020-10-06 15:39:43 | 2021-09-17 00:46:23 |
| [artifex](https://github.com/mborders/artifex) | 116 | 8 | 0 | Simple in-memory job queue for Golang using worker-based dispatching | 2018-10-31 19:34:31 | 2021-08-03 07:54:45 |
| [breaker](https://pkg.go.dev/github.com/kamilsk/breaker) | 97 | 5 | 7 | 🚧 Flexible mechanism to make execution flow interruptible. | 2019-02-15 15:08:24 | 2021-06-27 05:54:10 |
| [async](https://github.com/StudioSol/async) | 91 | 10 | 2 | A safe way to execute functions asynchronously, recovering them in case of panic. It also provides an error stack aiming to facilitate fail causes discovery. | 2017-06-30 17:08:33 | 2021-09-14 05:37:17 |
| [errgroup](https://github.com/neilotoole/errgroup) | 90 | 6 | 5 | errgroup with goroutine worker limits | 2020-06-26 06:07:39 | 2021-09-05 16:39:22 |
| [semaphore](https://github.com/kamilsk/semaphore) | 88 | 10 | 6 | 🚦 Semaphore pattern implementation with timeout of lock/unlock operations. | 2016-10-08 11:48:12 | 2021-09-15 05:47:25 |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 81 | 3 | 0 | gpool - a generic context-aware resizable goroutines pool to bound concurrency based on semaphore.  | 2018-12-03 04:23:35 | 2021-08-21 03:07:53 |
| [worker-pool](https://github.com/vardius/worker-pool) | 80 | 13 | 0 | Go simple async worker pool | 2017-10-04 09:18:31 | 2021-09-10 18:26:01 |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 76 | 10 | 0 | CyclicBarrier golang implementation | 2018-01-11 10:38:46 | 2021-09-14 07:44:21 |
| [gollback](https://github.com/vardius/gollback) | 63 | 4 | 0 | Go asynchronous simple function utilities, for managing execution of closures and callbacks | 2019-05-11 05:56:37 | 2021-09-10 22:44:53 |
| [threadpool](https://github.com/shettyh/threadpool) | 61 | 14 | 0 | Golang simple thread pool implementation | 2017-09-06 18:45:39 | 2021-09-12 19:02:33 |
| [Hunch](https://github.com/AaronJan/Hunch) | 57 | 5 | 0 | Hunch provides functions like: All, First, Retry, Waterfall etc., that makes asynchronous flow control more intuitive. | 2019-06-05 13:21:04 | 2021-09-08 06:42:18 |
| [routine](https://github.com/x-mod/routine) | 44 | 6 | 0 | go routine control, abstraction of the Main and some useful Executors.如果你不会管理Goroutine的话，用它 | 2019-03-04 12:25:23 | 2021-09-08 10:37:46 |
| [nursery](https://github.com/arunsworld/nursery) | 35 | 4 | 1 | Structured Concurrency in Go | 2019-11-23 19:26:02 | 2021-09-09 13:01:34 |
| [kyoo](https://github.com/dirkaholic/kyoo) | 34 | 0 | 0 | Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries | 2020-01-06 20:35:11 | 2021-09-11 12:57:16 |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 31 | 1 | 0 | Run functions in parallel :comet: | 2017-06-18 09:47:54 | 2021-08-04 00:40:51 |
| [goccm](https://github.com/zenthangplus/goccm) | 25 | 4 | 2 | Limits the number of goroutines that are allowed to run concurrently | 2019-08-16 02:26:53 | 2021-08-30 07:43:10 |
| [async](https://pkg.go.dev/github.com/reugn/async) | 24 | 1 | 0 | Alternative sync library for Go | 2019-12-28 09:48:40 | 2021-08-04 07:38:32 |
| [oversight](https://cirello.io/oversight) | 22 | 3 | 0 | [Mirror] Erlang-like supervisor trees | 2018-11-09 14:46:48 | 2021-07-23 15:23:01 |
| [go-waitgroup](https://www.yellowduck.be) | 21 | 1 | 0 | A sync.WaitGroup with error handling and concurrency control | 2018-08-08 16:12:35 | 2021-09-14 13:40:41 |
| [go-trylock](https://github.com/subchen/go-trylock) | 20 | 6 | 0 | TryLock support on read-write lock for Golang | 2018-04-26 06:02:47 | 2021-09-15 06:24:02 |
| [stl](https://github.com/ssgreg/stl) | 19 | 3 | 0 | Software Transactional Locks | 2018-06-19 10:50:11 | 2021-07-23 22:35:06 |
| [channelify](https://github.com/ddelizia/channelify) | 14 | 1 | 1 | Make functions return a channel for parallel processing via go routines. | 2020-10-05 13:12:48 | 2021-08-01 23:18:31 |
| [gohive](https://github.com/loveleshsharma/gohive) | 13 | 1 | 1 | 🐝 A Highly Performant and easy to use goroutine pool for Go | 2019-05-31 10:44:24 | 2021-07-13 06:10:30 |
| [conexec](https://github.com/ITcathyh/conexec) | 11 | 1 | 0 | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking. | 2019-12-24 07:35:11 | 2021-07-18 16:08:19 |
| [gowl](https://github.com/hamed-yousefi/gowl) | 11 | 2 | 4 | Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. | 2021-04-12 19:15:53 | 2021-09-05 15:31:48 |
| [queue](https://github.com/AnikHasibul/queue) | 9 | 1 | 0 | package queue gives you a queue group accessibility. Helps you to limit goroutines, wait for the end of the all goroutines and much more. | 2018-12-21 07:15:00 | 2020-11-02 15:19:05 |
| [hands](https://github.com/duanckham/hands) | 7 | 1 | 1 | Hands is a process controller used to control the execution and return strategies of multiple goroutines. | 2020-04-04 11:04:11 | 2021-06-24 01:19:23 |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 0 | A collection of tools for Golang | 2018-11-14 02:53:08 | 2019-05-04 06:52:40 |
| [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) | 4 | 1 | 0 | Make functions return a channel for parallel processing via go routines. | 2020-11-22 02:35:52 | 2020-12-04 21:15:03 |
| [breaker](https://pkg.go.dev/github.com/kamilsk/breaker) | 0 | 0 | 0 | 🚧 Flexible mechanism to make execution flow interruptible. | 2021-07-11 10:35:18 | 2021-07-11 10:35:53 |
</details>

### Images
Libraries for manipulating images.

<sup>*Last Update: 2021-10-11 09:25:06*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocv](https://gocv.io) | 4,385 | 645 | 176 | Go package for computer vision using OpenCV 4 and beyond. | 2017-09-18 21:54:17 | 2021-10-10 16:35:51 |
| [imaging](https://github.com/disintegration/imaging) | 3,943 | 327 | 14 | Imaging is a simple image processing package for Go | 2012-12-06 20:21:21 | 2021-10-10 12:03:56 |
| [imaginary](https://fly.io/docs/app-guides/run-a-global-image-service/) | 3,865 | 370 | 96 | Fast, simple, scalable, Docker-ready HTTP microservice for high-level image processing | 2015-03-04 18:51:40 | 2021-10-09 17:41:51 |
| [bild](https://github.com/anthonynsimon/bild) | 3,370 | 178 | 13 | Image processing algorithms in pure Go | 2016-08-01 15:54:29 | 2021-10-10 12:51:55 |
| [gg](https://godoc.org/github.com/fogleman/gg) | 3,034 | 219 | 61 | Go Graphics - 2D rendering in Go with a simple API. | 2016-02-18 21:05:08 | 2021-10-09 14:48:03 |
| [ln](https://godoc.org/github.com/fogleman/ln/ln) | 2,970 | 111 | 12 | 3D line art engine. | 2016-01-10 04:28:10 | 2021-10-10 11:18:49 |
| [resize](https://github.com/nfnt/resize) | 2,736 | 265 | 9 | Pure golang image resizing  | 2012-08-02 19:48:26 | 2021-10-10 13:29:11 |
| [pt](http://bit.ly/1E7rSoi) | 1,967 | 113 | 8 | A path tracer written in Go. | 2015-01-23 19:39:29 | 2021-10-06 10:44:50 |
| [svgo](https://github.com/ajstarks/svgo) | 1,738 | 146 | 9 | Go Language Library for SVG generation | 2010-03-05 23:24:10 | 2021-10-10 03:12:27 |
| [bimg](https://pkg.go.dev/github.com/h2non/bimg?tab=doc) | 1,649 | 275 | 120 | Go package for fast high-level image processing powered by libvips C library | 2015-03-17 14:14:02 | 2021-10-09 20:07:32 |
| [picfit](http://bit.ly/1E7rSoi) | 1,548 | 131 | 16 | An image resizing server written in Go | 2014-12-06 17:30:45 | 2021-10-11 00:36:10 |
| [smartcrop](https://github.com/muesli/smartcrop) | 1,548 | 106 | 7 | smartcrop finds good image crops for arbitrary crop sizes | 2014-04-07 22:40:03 | 2021-10-09 14:29:49 |
| [gift](https://godoc.org/github.com/fogleman/gg) | 1,467 | 107 | 2 | Go Image Filtering Toolkit | 2014-07-12 18:47:40 | 2021-10-09 08:11:33 |
| [imagick](https://godoc.org/github.com/gographics/imagick/imagick) | 1,360 | 163 | 11 | Go binding to ImageMagick's MagickWand C API | 2013-04-30 17:31:48 | 2021-10-05 20:03:11 |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1,251 | 199 | 45 | Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv | 2013-12-09 09:43:26 | 2021-10-08 21:05:35 |
| [geopattern](https://github.com/pravj/geopattern) | 1,137 | 62 | 3 | :triangular_ruler: Create beautiful generative image patterns from a string in golang. | 2014-10-22 17:26:30 | 2021-10-08 12:25:02 |
| [stegify](https://github.com/DimitarPetrov/stegify) | 949 | 112 | 0 | 🔍 Go tool for LSB steganography, capable of hiding any file within an image. | 2018-11-29 16:45:58 | 2021-10-03 11:39:10 |
| [canvas](https://github.com/tdewolff/canvas) | 812 | 46 | 10 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2017-05-20 18:10:51 | 2021-10-08 16:21:29 |
| [image2ascii](https://github.com/qeesung/image2ascii) | 585 | 53 | 4 | :foggy: Convert image to ASCII | 2018-10-20 05:06:25 | 2021-10-10 16:26:11 |
| [govips](https://github.com/davidbyttow/govips) | 517 | 116 | 13 | A lightning fast image processing and resizing library for Go | 2016-12-25 04:32:56 | 2021-10-03 11:41:01 |
| [draft](https://github.com/lucasepe/draft) | 510 | 22 | 0 | Generate High Level Cloud Architecture diagrams using YAML syntax. | 2020-06-05 16:11:40 | 2021-09-22 03:45:53 |
| [govatar](https://github.com/o1egl/govatar) | 459 | 26 | 0 | Avatar generation library for GO language | 2016-01-18 12:12:28 | 2021-10-09 00:21:58 |
| [mort](https://github.com/aldor007/mort) | 436 | 17 | 8 | Storage and image processing server written in Go | 2017-11-19 13:37:58 | 2021-10-07 19:02:31 |
| [goimagehash](https://github.com/corona10/goimagehash) | 434 | 47 | 10 | Go Perceptual image hashing package | 2017-07-28 17:15:58 | 2021-10-09 14:39:48 |
| [go-nude](https://github.com/koyachi/go-nude) | 335 | 37 | 2 | Nudity detection with Go. | 2014-05-02 08:32:29 | 2021-10-02 15:36:10 |
| [rez](https://github.com/bamiaux/rez) | 202 | 17 | 1 | Image resizing in pure Go and SIMD | 2014-01-16 21:16:15 | 2021-09-29 14:46:13 |
| [darkroom](https://www.gojek.io/darkroom/) | 171 | 32 | 7 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2019-07-01 10:17:08 | 2021-10-10 07:23:21 |
| [mergi](http://mergi.io/) | 154 | 21 | 2 | go library for image programming (merge, crop, resize, watermark, animate, ease, transit) | 2018-09-24 03:40:47 | 2021-10-05 10:04:20 |
| [img](hawx.me/code/img) | 137 | 8 | 1 | A selection of image manipulation tools | 2012-07-28 19:57:47 | 2021-09-13 17:50:07 |
| [gltf](https://www.khronos.org/gltf/) | 121 | 20 | 3 | :eyeglasses: Go library for [d]encoding glTF 2.0 files | 2019-01-15 17:43:54 | 2021-10-08 01:10:26 |
| [steganography](https://github.com/auyer/steganography) | 115 | 18 | 0 | Pure Golang Library that allows simple LSB steganography on images | 2018-05-21 17:27:36 | 2021-10-10 06:13:33 |
| [go-cairo](https://github.com/ungerik/go-cairo) | 115 | 27 | 0 | Go binding for the cairo graphics library | 2012-08-22 18:27:01 | 2021-09-12 16:03:34 |
| [cameron](https://pkg.go.dev/github.com/aofei/cameron) | 70 | 7 | 1 | An avatar generator for Go. | 2018-05-05 22:13:11 | 2021-09-06 07:44:56 |
| [go-gd](https://github.com/bolknote/go-gd) | 53 | 15 | 1 | Go bingings for GD (http://www.boutell.com/gd/) | 2011-05-12 06:33:54 | 2021-06-29 11:02:01 |
| [gridder](https://github.com/shomali11/gridder) | 46 | 5 | 0 | A Grid based 2D Graphics library | 2020-04-10 00:13:10 | 2021-09-30 17:31:45 |
| [goimghdr](https://github.com/corona10/goimghdr) | 36 | 2 | 0 | The imghdr module determines the type of image contained in a file for go | 2018-02-25 09:34:44 | 2021-06-28 01:11:03 |
| [tga](https://github.com/ftrvxmtrx/tga) | 28 | 10 | 1 | Go package for decoding and encoding TARGA image format | 2012-10-08 01:09:20 | 2021-06-28 01:25:28 |
| [webp-server](https://github.com/mehdipourfar/webp-server) | 27 | 5 | 0 | Simple and minimal image server capable of storing, resizing, converting and caching images. | 2020-11-22 12:03:12 | 2021-09-21 17:58:57 |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 25 | 4 | 0 | Port of webcolors library from Python to Go | 2014-04-24 14:41:22 | 2021-06-28 01:09:17 |
| [mpo](https://donatstudios.com/MPO-to-JPEG-Stereo) | 7 | 2 | 1 | JPEG-MPO Decoder / Converter Library and CLI Tool | 2015-04-14 22:37:59 | 2021-06-28 01:20:05 |
</details>

### IoT (Internet of Things)
Libraries for programming devices of the IoT.

<sup>*Last Update: 2021-08-31 09:25:14*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gobot](https://gobot.io) | 7,295 | 912 | 167 | Golang framework for robotics, drones, and the Internet of Things (IoT) | 2013-09-21 14:09:19 | 2021-08-29 19:10:57 |
| [flogo](http://flogo.io) | 1,831 | 248 | 153 | Project Flogo is an open source ecosystem of opinionated  event-driven capabilities to simplify building efficient & modern serverless functions, microservices & edge apps. | 2016-07-10 02:57:43 | 2021-08-28 12:50:32 |
| [periph](https://periph.io) | 1,675 | 173 | 42 | Go·Hardware·Lean | 2016-10-13 16:53:51 | 2021-08-30 20:45:59 |
| [mainflux](https://www.mainflux.io) | 1,526 | 451 | 83 | Industrial IoT Messaging and Device Management Platform | 2015-07-06 20:31:50 | 2021-08-29 10:49:52 |
| [gatt](http://flogo.io) | 986 | 260 | 51 | Gatt is a Go package for building Bluetooth Low Energy peripherals | 2014-04-23 13:45:27 | 2021-08-29 18:54:47 |
| [heedy](https://heedy.org) | 285 | 27 | 9 | An aggregator for personal metrics, and an extensible analysis engine | 2015-01-16 19:44:21 | 2021-08-25 08:26:41 |
| [devices](https://github.com/goiot/devices) | 242 | 24 | 9 | Suite of libraries for IoT devices (written in Go), experimental for x/exp/io | 2016-05-30 08:07:02 | 2021-08-16 08:08:18 |
| [sensorbee](http://sensorbee.io/) | 205 | 34 | 39 | Lightweight stream processing engine for IoT | 2016-02-19 07:49:56 | 2021-08-10 08:44:31 |
| [huego](https://github.com/amimof/huego) | 187 | 31 | 5 | An extensive Philips Hue client library for Go with an emphasis on simplicity | 2017-05-16 05:31:45 | 2021-08-26 13:20:48 |
| [iot](https://github.com/vaelen/iot) | 52 | 7 | 0 | A Go client for Google IoT Core | 2018-03-08 06:51:51 | 2021-04-30 18:44:01 |
| [eywa](https://github.com/xcodersun/eywa) | 49 | 11 | 9 | Make IoT a lot more fun with data.  | 2016-02-20 17:02:16 | 2021-08-17 07:10:31 |
</details>

### JSON
Libraries for working with JSON.

<sup>*Last Update: 2021-08-23 09:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gjson](https://github.com/tidwall/gjson) | 8,808 | 600 | 30 | Get JSON values quickly - JSON parser for Go | 2016-08-11 03:08:47 | 2021-08-22 17:13:44 |
| [json-to-go](https://mholt.github.io/json-to-go/) | 3,213 | 383 | 12 | Translates JSON into a Go type in your browser instantly (original) | 2014-01-21 18:11:13 | 2021-08-22 17:18:03 |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2,393 | 184 | 41 | Automatically generate Go (golang) struct definitions from example JSON | 2012-12-27 19:10:50 | 2021-08-22 08:15:37 |
| [fastjson](https://github.com/valyala/fastjson) | 1,293 | 72 | 33 | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection | 2018-05-28 21:41:47 | 2021-08-22 07:30:04 |
| [kazaam](https://github.com/qntfy/kazaam) | 203 | 42 | 23 | Arbitrary transformations of JSON in Golang | 2016-07-19 14:19:03 | 2021-08-17 08:09:57 |
| [gojq](https://godoc.org/github.com/nicklaw5/go-respond) | 173 | 20 | 1 | JSON query in Golang | 2015-12-30 09:02:13 | 2021-08-18 06:54:30 |
| [jsondiff](https://pkg.go.dev/github.com/wI2L/jsondiff) | 125 | 10 | 0 | JSON diff library for Go based on RFC6902 (JSON Patch) | 2020-11-28 19:05:16 | 2021-08-19 10:16:22 |
| [jettison](https://pkg.go.dev/github.com/wI2L/jettison) | 111 | 7 | 0 | Fast and flexible JSON encoder for Go | 2019-08-30 13:28:03 | 2021-08-12 05:15:25 |
| [jsongo](https://pkg.go.dev/github.com/wI2L/jsondiff) | 96 | 13 | 1 | Fluent API to make it easier to create Json objects. | 2015-08-07 23:23:17 | 2021-05-01 10:53:24 |
| [gjo](https://github.com/skanehira/gjo) | 94 | 11 | 1 | Small utility to create JSON objects | 2019-02-23 01:54:21 | 2021-07-06 17:53:06 |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 84 | 7 | 2 | A JSON diff utility | 2017-04-24 16:05:35 | 2021-08-20 14:26:03 |
| [json2go](https://m-zajac.github.io/json2go) | 84 | 12 | 1 | Create go type representation from json | 2017-06-10 23:55:07 | 2021-08-20 10:56:25 |
| [ajson](https://github.com/spyzhov/ajson) | 66 | 8 | 5 | Abstract JSON for golang with JSONPath support  | 2019-03-07 20:47:38 | 2021-08-18 08:22:30 |
| [jsonf](https://pkg.go.dev/github.com/wI2L/jsondiff) | 62 | 9 | 0 | Console JSON formatter with query feature | 2015-05-25 04:53:32 | 2021-08-04 01:58:06 |
| [json-to-proto.github.io](https://json-to-proto.github.io/) | 51 | 7 | 1 | convert JSON to Protocol Buffers online in your browser instantly | 2020-04-18 20:42:45 | 2021-08-19 06:51:07 |
| [mp](https://github.com/sanbornm/mp) | 44 | 4 | 1 | Simple Email Parser | 2014-06-15 21:14:39 | 2020-10-22 17:21:09 |
| [go-respond](https://godoc.org/github.com/nicklaw5/go-respond) | 41 | 6 | 1 | A Go package for handling common HTTP JSON responses. | 2017-03-12 21:00:54 | 2021-07-29 18:23:08 |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 10 | 1 | 0 | Small package which wraps error responses to follow jsonapi.org | 2018-10-18 15:03:45 | 2020-10-02 04:02:28 |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 9 | 2 | 0 | Go bindings based on the JSON API errors reference | 2016-07-08 10:08:58 | 2020-08-30 14:39:42 |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 4 | 1 | A simple Go package to make custom structs marshal into HAL compatible JSON responses. | 2016-01-15 11:38:40 | 2020-05-26 19:24:02 |
| [ask](https://github.com/simonnilsson/ask) | 7 | 0 | 0 | A Go package that provides a simple way of accessing nested properties in maps and slices. | 2020-09-13 13:53:31 | 2021-05-29 00:44:25 |
| [ej](https://github.com/lucassscaravelli/ej) | 7 | 0 | 0 | Write and read JSON from different sources in one line | 2020-01-04 17:39:35 | 2021-04-27 22:00:09 |
| [dynjson](https://github.com/cocoonspace/dynjson) | 7 | 2 | 0 | Client-customizable JSON formats for dynamic APIs | 2020-05-06 07:10:02 | 2021-08-14 10:47:09 |
| [epoch](https://github.com/vtopc/epoch) | 6 | 1 | 0 | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON | 2019-12-15 12:54:37 | 2021-08-04 01:57:13 |
| [mapslice-json](https://github.com/ake-persson/mapslice-json) | 6 | 1 | 0 | Go MapSlice for ordered marshal/ unmarshal of maps in JSON | 2020-02-19 11:01:48 | 2021-08-21 00:15:30 |
| [jzon](https://github.com/zerosnake0/jzon) | 4 | 0 | 0 | A golang json library inspired by jsoniter | 2019-11-12 10:42:41 | 2021-05-24 05:31:48 |
| [jsonic](https://github.com/sinhashubham95/jsonic) | 3 | 0 | 0 | All you need with JSON | 2021-01-09 06:21:59 | 2021-07-18 18:43:03 |
</details>

### Job Scheduler
Libraries for scheduling jobs.

<sup>*Last Update: 2021-08-19 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocron](https://github.com/go-co-op/gocron) | 1,039 | 82 | 12 | Easy and fluent Go cron scheduling. This is a fork from https://github.com/jasonlvhit/gocron | 2020-03-20 15:33:05 | 2021-08-18 13:01:07 |
| [gron](https://github.com/roylee0704/gron) | 871 | 53 | 8 | gron, Cron Jobs in Go. | 2016-06-04 08:02:22 | 2021-08-13 15:52:24 |
| [jobrunner](https://github.com/bamzi/jobrunner) | 870 | 76 | 10 | Framework for performing work asynchronously, outside of the request flow | 2015-10-21 04:17:01 | 2021-08-15 08:34:22 |
| [jobs](https://github.com/albrow/jobs) | 482 | 40 | 17 | A persistent and flexible background jobs library for go. | 2015-02-09 22:13:29 | 2021-05-31 17:24:28 |
| [scheduler](https://github.com/carlescere/scheduler) | 369 | 49 | 6 | Job scheduling made easy. | 2015-02-03 17:10:23 | 2021-08-13 15:52:29 |
| [go-cron](https://github.com/rk/go-cron) | 204 | 16 | 0 | A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. | 2011-04-15 14:50:49 | 2021-08-13 15:52:13 |
| [go-quartz](https://pkg.go.dev/github.com/reugn/go-quartz/quartz) | 147 | 15 | 0 | Simple, zero-dependency scheduling library for Go | 2019-04-14 18:57:51 | 2021-08-17 18:33:07 |
| [clockwerk](https://github.com/onatm/clockwerk) | 110 | 10 | 0 | Job Scheduling Library | 2017-04-09 23:10:48 | 2021-08-13 15:52:06 |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 83 | 11 | 13 | You had one job, or more then one, which can be done in steps | 2018-04-08 13:44:04 | 2021-07-10 15:52:53 |
| [tasks](https://github.com/madflojo/tasks) | 52 | 4 | 1 | Package tasks is an easy to use in-process scheduler for recurring tasks in Go | 2019-12-24 18:26:18 | 2021-08-12 02:38:58 |
| [clockwork](https://github.com/whiteShtef/clockwork) | 29 | 12 | 2 | Job Scheduling Library | 2020-02-21 01:25:57 | 2021-07-13 06:42:08 |
| [cronticker](https://github.com/krayzpipes/cronticker) | 1 | 0 | 0 | Golang ticker that works with Cron scheduling. | 2020-11-28 20:59:38 | 2021-01-02 01:57:07 |
</details>

### Logging
Libraries for generating and working with log files.

<sup>*Last Update: 2021-08-13 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [logrus](https://github.com/sirupsen/logrus) | 18,449 | 1,991 | 126 | Structured, pluggable logging for Go. | 2013-10-16 19:08:55 | 2021-08-13 02:23:21 |
| [zap](https://godoc.org/go.uber.org/zap) | 13,344 | 998 | 83 | Blazing fast, structured, leveled logging in Go. | 2016-02-18 19:52:56 | 2021-08-13 01:23:17 |
| [zerolog](https://github.com/rs/zerolog) | 5,070 | 306 | 75 | Zero Allocation JSON Logger | 2017-05-12 05:24:39 | 2021-08-13 00:39:37 |
| [go-spew](https://github.com/davecgh/go-spew) | 4,524 | 293 | 55 | Implements a deep pretty printer for Go data structures to aid in debugging | 2013-01-09 05:18:22 | 2021-08-12 16:52:37 |
| [glog](https://github.com/golang/glog) | 2,775 | 790 | 8 | Leveled execution logs for Go | 2013-07-16 04:33:04 | 2021-08-11 06:28:56 |
| [lumberjack](https://github.com/natefinch/lumberjack) | 2,703 | 353 | 53 | lumberjack is a log rolling package for Go | 2014-06-14 11:55:47 | 2021-08-12 11:07:12 |
| [tail](https://github.com/hpcloud/tail) | 2,145 | 435 | 71 | Go package for reading from continously updated files (tail -f) | 2013-02-05 00:28:03 | 2021-08-12 19:20:15 |
| [seelog](https://github.com/cihub/seelog) | 1,557 | 242 | 39 | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. | 2011-11-17 09:43:15 | 2021-08-11 12:33:06 |
| [log](https://github.com/apex/log) | 1,158 | 99 | 34 | Structured logging package for Go. | 2015-12-21 20:27:48 | 2021-08-11 14:59:24 |
| [log15](https://godoc.org/github.com/inconshreveable/log15) | 1,013 | 140 | 43 | Structured, composable logging for Go | 2014-05-20 00:11:52 | 2021-07-29 08:51:21 |
| [onelog](https://github.com/francoispqt/onelog) | 397 | 14 | 1 | Dead simple, super fast, zero allocation and modular logger for Golang | 2018-05-06 14:32:10 | 2021-08-09 19:51:49 |
| [log](https://github.com/phuslu/log) | 390 | 26 | 4 | Structured Logging Made Easy | 2019-07-07 09:40:38 | 2021-08-12 17:12:50 |
| [logxi](https://logur.dev/logur) | 347 | 40 | 24 | A 12-factor app logger built for performance and happy development | 2015-03-01 22:13:45 | 2021-05-08 17:22:18 |
| [logutils](https://logur.dev/logur) | 295 | 31 | 1 | Utilities for slightly better logging in Go (Golang). | 2013-10-09 07:31:15 | 2021-07-18 02:53:40 |
| [log](https://github.com/go-playground/log) | 275 | 21 | 1 | :green_book: Simple, configurable and scalable Structured Logging for Go. | 2016-02-07 16:17:48 | 2021-05-06 17:20:44 |
| [go-logger](https://github.com/apsdehal/go-logger) | 266 | 49 | 2 | Simple logger for Go programs. Allows custom formats for messages. | 2014-09-26 04:57:06 | 2021-07-15 17:58:40 |
| [httpretty](https://asciinema.org/a/297429) | 234 | 6 | 1 | Package httpretty prints the HTTP requests you make with Go pretty on your terminal. | 2020-01-24 18:17:16 | 2021-08-05 03:16:33 |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 184 | 5 | 4 | A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. | 2019-11-02 17:28:03 | 2021-08-12 11:02:12 |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 177 | 26 | 8 | Rolling writer is an IO util for auto rolling write in go. | 2017-02-12 12:05:26 | 2021-07-25 03:10:36 |
| [logger](http://godoc.org/github.com/azer/logger) | 148 | 15 | 0 | Minimalistic logging library for Go. | 2014-09-30 06:45:09 | 2021-03-29 21:01:13 |
| [xlog](https://github.com/rs/xlog) | 135 | 12 | 3 | xlog is a logger for net/context aware HTTP applications | 2015-10-22 09:26:45 | 2021-02-20 02:54:42 |
| [logur](https://logur.dev/logur) | 135 | 9 | 8 | Logur is an opinionated collection of logging best practices | 2018-12-09 16:43:11 | 2021-07-30 07:58:37 |
| [glg](https://github.com/kpango/glg) | 118 | 11 | 0 | Simple and blazing fast lockfree logging library for golang | 2017-06-21 13:26:16 | 2021-08-11 07:13:21 |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 114 | 31 | 9 | A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. | 2015-10-22 22:29:02 | 2021-05-17 07:01:57 |
| [logvoyage](https://github.com/firstrow/logvoyage) | 88 | 10 | 9 | LogVoyage - logging SaaS written in GoLang | 2015-03-29 11:05:09 | 2020-09-04 14:11:03 |
| [log](https://github.com/alexcesaro/log) | 45 | 3 | 1 | Logging packages for Go | 2014-04-19 14:31:56 | 2020-08-12 17:36:34 |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 42 | 7 | 3 | Time based rotating file writer | 2017-02-04 09:02:55 | 2021-08-10 04:23:37 |
| [gologger](https://github.com/sadlil/gologger) | 39 | 10 | 2 | The Simplest and worst logging library ever written | 2015-09-02 08:52:26 | 2021-01-19 03:46:18 |
| [logex](https://github.com/chzyer/logex) | 37 | 8 | 2 | An golang log lib, supports tracking and level, wrap by standard log lib | 2014-10-10 06:38:39 | 2021-03-29 21:33:46 |
| [go-log](https://github.com/ian-kent/go-log) | 37 | 17 | 3 | A logger, for Go | 2014-05-02 00:34:09 | 2021-02-21 19:02:39 |
| [gone](https://github.com/One-com/gone) | 36 | 5 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2021-08-01 19:37:48 |
| [go-log](https://github.com/siddontang/go-log) | 28 | 14 | 1 | a golang log lib supports level and multi handlers | 2014-05-18 03:41:55 | 2020-11-11 08:54:54 |
| [journald](https://asciinema.org/a/297429) | 27 | 1 | 0 | Go implementation of systemd Journal's native API for logging | 2017-08-23 07:06:09 | 2021-08-02 21:07:20 |
| [logrusly](https://github.com/sebest/logrusly) | 27 | 15 | 3 | Loggly Hooks for GO Logrus logger | 2014-09-11 23:27:11 | 2021-07-27 21:32:37 |
| [distillog](https://github.com/amoghe/distillog) | 26 | 6 | 0 | Logging, distilled | 2015-10-12 16:32:21 | 2021-06-21 15:05:58 |
| [log](https://github.com/teris-io/log) | 24 | 2 | 0 | Structured log interface | 2017-10-28 19:57:55 | 2021-01-03 07:59:28 |
| [mlog](https://github.com/jbrodriguez/mlog) | 24 | 17 | 1 | A simple logging module for go, with a rotating file feature and console logging. | 2014-10-20 15:06:26 | 2021-06-28 02:21:55 |
| [gomol](https://github.com/aphistic/gomol) | 17 | 0 | 3 | Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs | 2015-08-30 15:51:46 | 2021-04-21 04:52:21 |
| [zkits-logger](https://github.com/edoger/zkits-logger) | 15 | 0 | 0 | A powerful zero-dependency json logger. | 2020-03-31 14:23:40 | 2021-07-05 06:43:25 |
| [glo](https://github.com/lajosbencz/glo) | 14 | 0 | 0 | Logging library for Golang | 2019-01-19 22:10:42 | 2021-06-04 05:15:29 |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 12 | 0 | 0 | io.Writer implementation using logrus logger | 2019-08-09 08:58:25 | 2021-04-28 10:20:48 |
| [go-log](https://github.com/subchen/go-log) | 11 | 5 | 0 | Simple and configurable Logging in Go, with level, formatters and writers | 2017-05-07 08:09:24 | 2021-02-21 03:47:33 |
| [logmatic](http://godoc.org/github.com/azer/logger) | 10 | 3 | 1 | Colorized logger for Golang with dynamic log level configuration | 2018-11-07 01:52:45 | 2021-03-10 18:40:34 |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 0 | Package for multi-level logging | 2017-01-13 15:34:31 | 2019-07-22 17:02:07 |
| [logo](https://github.com/mbndr/logo) | 9 | 1 | 0 | Golang logger to different configurable writers. | 2017-02-07 18:02:55 | 2021-02-09 06:29:03 |
| [log](https://github.com/aerogo/log) | 8 | 0 | 0 | :memo: Logging with multiple output targets. | 2017-06-10 09:54:08 | 2020-11-19 15:43:56 |
| [go-log](https://github.com/pieterclaerhout/go-log) | 8 | 3 | 0 | A logging library with strack traces, object dumping and optional timestamps | 2019-10-01 08:55:38 | 2021-02-21 07:39:35 |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 3 | 0 | plugin architecture and flexible log system for golang | 2016-05-05 16:47:45 | 2019-09-26 11:33:58 |
| [kemba](https://pkg.go.dev/github.com/clok/kemba?tab=overview) | 5 | 1 | 1 | A tiny debug logging tool. Ideal for CLI tools and command applications. Inspired by https://github.com/visionmedia/debug | 2020-07-13 03:10:54 | 2021-05-21 04:32:04 |
| [yell](https://github.com/jfcg/yell) | 0 | 0 | 0 | Yet another minimalist logging library | 2021-02-07 16:07:27 | 2021-08-07 13:35:27 |
</details>

### Machine Learning
Libraries for Machine Learning.

<sup>*Last Update: 2021-08-25 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 7,978 | 1,113 | 68 | Machine Learning for Go | 2013-12-26 13:06:14 | 2021-08-24 07:37:28 |
| [gorse](https://gorse.io) | 4,308 | 331 | 19 | An open source recommender system service written in Go | 2018-08-14 11:01:09 | 2021-08-24 19:26:20 |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 4,155 | 360 | 78 | Gorgonia is a library that helps facilitate machine learning in Go. | 2016-09-14 23:19:43 | 2021-08-24 02:02:53 |
| [tfgo](https://pgaleone.eu/tensorflow/go/2017/05/29/understanding-tensorflow-using-go/) | 1,762 | 132 | 9 | Tensorflow + Go, the gopher way | 2017-05-23 13:27:39 | 2021-08-20 14:47:49 |
| [gosseract](https://pkg.go.dev/github.com/otiai10/gosseract) | 1,548 | 191 | 18 | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library | 2013-10-11 07:27:53 | 2021-08-23 13:30:57 |
| [goml](https://github.com/cdipaolo/goml) | 1,214 | 115 | 7 | On-line Machine Learning in Go (and so much more) | 2015-06-27 05:52:01 | 2021-08-22 12:59:57 |
| [eaopt](https://pkg.go.dev/github.com/MaxHalford/eaopt) | 741 | 83 | 7 | :four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution) | 2016-01-31 00:04:52 | 2021-08-17 23:21:21 |
| [bayesian](https://github.com/jbrukh/bayesian) | 712 | 123 | 8 | Naive Bayesian Classification for Golang. | 2011-11-23 04:17:00 | 2021-08-04 04:45:08 |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 688 | 88 | 35 | Ensembles of decision trees in go/golang. | 2012-10-22 17:38:16 | 2021-08-24 17:47:50 |
| [gobrain](https://github.com/goml/gobrain) | 485 | 55 | 1 | Neural Networks written in go | 2014-04-29 13:32:36 | 2021-08-13 21:33:02 |
| [ocrserver](https://ocr-example.herokuapp.com/) | 394 | 93 | 1 | A simple OCR API server, seriously easy to be deployed by Docker, on Heroku as well | 2015-11-15 07:57:42 | 2021-08-19 05:56:45 |
| [onnx-go](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 339 | 34 | 21 | onnx-go gives the ability to import a pre-trained neural network within Go without being linked to a framework or library. | 2018-08-28 07:39:20 | 2021-08-17 10:46:47 |
| [go-deep](https://github.com/patrikeh/go-deep) | 324 | 32 | 7 | Artificial Neural Network | 2017-12-09 15:10:06 | 2021-08-16 04:01:01 |
| [regommend](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 291 | 26 | 0 | Recommendation engine for Go | 2014-02-05 17:00:49 | 2021-07-15 03:22:47 |
| [go-galib](https://github.com/thoj/go-galib) | 186 | 42 | 0 | Genetic Algorithms library written in Go / golang | 2009-11-30 10:46:58 | 2021-08-01 23:31:37 |
| [goptuna](https://pkg.go.dev/github.com/c-bata/goptuna) | 185 | 10 | 11 | A hyperparameter optimization framework, inspired by Optuna. | 2019-07-24 12:03:05 | 2021-07-29 16:33:51 |
| [goRecommend](https://pkg.go.dev/github.com/c-bata/goptuna) | 175 | 19 | 0 | Collaborative Filtering (CF) Algorithms in Go!  | 2014-07-16 05:32:23 | 2021-07-16 22:20:23 |
| [shield](https://github.com/eaigner/shield) | 142 | 29 | 5 | Bayesian text classifier with flexible tokenizers and storage backends for Go | 2013-04-10 19:38:16 | 2021-07-29 13:21:43 |
| [goga](https://github.com/tomcraven/goga) | 111 | 11 | 0 | Golang Genetic Algorithm | 2015-10-20 12:50:51 | 2021-08-11 14:44:33 |
| [go-fann](https://github.com/vksnk/go-fann) | 103 | 21 | 2 | Go bindings for FANN, library for artificial neural networks | 2011-03-10 21:08:27 | 2021-04-11 03:02:00 |
| [gonet](https://pkg.go.dev/github.com/dathoangnd/gonet) | 68 | 5 | 0 | Neural Network for Go. | 2020-01-11 18:27:28 | 2021-08-01 23:35:20 |
| [libsvm](https://pkg.go.dev/github.com/otiai10/gosseract) | 67 | 12 | 1 | libsvm go version | 2012-07-31 07:57:47 | 2021-01-13 19:35:07 |
| [goscore](https://gorse.io) | 65 | 20 | 3 | Go Scoring API for PMML | 2017-08-19 11:08:39 | 2021-05-17 05:16:58 |
| [neural-go](https://github.com/schuyler/neural-go) | 61 | 15 | 1 | A multilayer perceptron network implemented in Go, with training via backpropagation. | 2011-10-17 09:31:33 | 2021-07-29 15:53:28 |
| [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 60 | 5 | 2 | Fast, simple sklearn-like feature processing for Go | 2020-12-18 13:09:18 | 2021-08-23 16:18:31 |
| [go-pr](https://github.com/daviddengcn/go-pr) | 59 | 12 | 0 | Pattern recognition package in Go lang. | 2013-06-07 02:36:20 | 2021-08-02 18:31:53 |
| [neat](https://github.com/jinyeom/neat) | 59 | 12 | 4 | NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go | 2016-11-17 04:23:14 | 2021-08-20 15:31:46 |
| [fonet](https://github.com/Fontinalis/fonet) | 45 | 9 | 2 | fonet is a deep neural network package for Go. | 2017-10-03 15:57:15 | 2021-08-01 23:30:05 |
| [golinear](https://github.com/danieldk/golinear) | 42 | 12 | 0 | liblinear bindings for Go | 2013-04-05 15:37:01 | 2021-08-02 18:34:03 |
| [Varis](https://github.com/Xamber/Varis) | 37 | 6 | 0 | Golang Neural Network  | 2017-10-10 08:43:27 | 2021-08-01 23:37:20 |
| [godist](https://github.com/e-dard/godist) | 29 | 6 | 0 | Probability distributions and associated methods in Go | 2014-09-05 09:48:51 | 2021-08-01 23:33:41 |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 27 | 6 | 0 | k-modes and k-prototypes clustering algorithms implementation in Go | 2017-10-04 12:24:52 | 2021-08-10 03:46:02 |
| [probab](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 16 | 4 | 3 | Automatically exported from code.google.com/p/probab | 2015-09-14 12:07:52 | 2021-06-02 23:42:13 |
| [evoli](https://github.com/khezen/evoli) | 15 | 6 | 21 | Genetic Algorithm and Particle Swarm Optimization | 2015-06-12 06:58:30 | 2021-03-15 12:19:59 |
| [gomind](https://github.com/surenderthakran/gomind) | 15 | 1 | 7 | A simplistic Neural Network Library in Go | 2017-10-19 03:48:51 | 2021-08-17 23:18:40 |
| [ddt](https://github.com/sgrodriguez/ddt) | 13 | 1 | 0 | Golang Dynamic Decision Tree | 2020-05-20 13:51:42 | 2021-08-20 15:31:51 |
| [randomForest](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 13 | 2 | 1 | Random Forest implementation in golang | 2018-10-25 07:05:29 | 2021-08-16 14:38:31 |
</details>

### Messaging
Libraries that implement messaging systems.

<sup>*Last Update: 2021-10-13 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [sarama](https://shopify.github.io/sarama) | 7,714 | 1,340 | 191 | Sarama is a Go library for Apache Kafka 0.8, and up. | 2013-07-05 18:52:38 | 2021-10-12 09:59:58 |
| [gorush](https://github.com/appleboy/gorush) | 5,775 | 647 | 43 | A push notification server written in Go (Golang). | 2016-03-22 07:15:20 | 2021-10-12 23:46:40 |
| [machinery](https://github.com/RichardKnop/machinery) | 5,647 | 704 | 171 | Machinery is an asynchronous task queue/job queue based on distributed message passing. | 2015-04-05 19:46:34 | 2021-10-13 01:31:17 |
| [centrifugo](https://centrifugal.dev) | 5,462 | 459 | 10 | Scalable real-time messaging server in a language-agnostic way. Set up once and forever. | 2015-03-31 20:26:49 | 2021-10-12 18:29:59 |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 4,289 | 696 | 86 | socket.io library for golang, a realtime application framework. | 2013-07-13 13:04:38 | 2021-10-12 16:28:10 |
| [nats.go](https://nats.io) | 3,582 | 460 | 42 | Golang client for NATS, the cloud native messaging system. | 2012-08-15 12:54:59 | 2021-10-12 16:32:06 |
| [benthos](https://www.benthos.dev) | 3,443 | 320 | 140 | Fancy stream processing made operationally mundane | 2016-03-22 01:18:48 | 2021-10-12 13:05:03 |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 2,929 | 450 | 173 | Confluent's Apache Kafka Golang client | 2016-07-12 22:23:34 | 2021-10-13 01:49:44 |
| [apns2](https://github.com/sideshow/apns2) | 2,560 | 283 | 24 | ⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol. | 2016-01-05 00:56:53 | 2021-10-12 12:22:47 |
| [mercure](https://mercure.rocks) | 2,558 | 180 | 11 | Server-sent live updates: protocol and reference implementation | 2018-07-14 13:47:14 | 2021-10-12 12:39:08 |
| [melody](https://github.com/olahol/melody) | 2,252 | 269 | 25 | :notes: Minimalist websocket framework for Go | 2015-05-13 20:38:32 | 2021-10-11 22:18:03 |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 2,005 | 555 | 4 | Golang push server cluster | 2013-12-27 08:56:10 | 2021-10-08 07:37:41 |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1,984 | 380 | 22 | The official Go package for NSQ | 2013-08-29 01:18:32 | 2021-10-12 17:34:15 |
| [asynq](https://github.com/hibiken/asynq) | 1,631 | 143 | 18 | Asynq: simple, reliable, and efficient distributed task queue in Go | 2019-11-15 05:17:55 | 2021-10-12 16:01:20 |
| [uniqush-push](http://uniqush.org) | 1,280 | 198 | 72 | Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices. | 2011-08-29 08:42:37 | 2021-10-11 15:59:40 |
| [Beaver](https://clivern.github.io/Beaver/) | 1,173 | 63 | 4 | 💨 A real time messaging system to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. | 2018-10-20 21:10:43 | 2021-10-11 16:48:13 |
| [EventBus](https://github.com/asaskevich/EventBus) | 999 | 119 | 17 | [Go] Lightweight eventbus with async compatibility for Go | 2014-12-19 16:38:39 | 2021-10-12 12:26:17 |
| [zmq4](http://uniqush.org) | 930 | 151 | 42 | A Go interface to ZeroMQ version 4 | 2013-10-18 11:48:51 | 2021-10-11 20:25:06 |
| [gollum](http://gollum.readthedocs.org/en/latest/) | 906 | 73 | 20 | An n:m message multiplexer written in Go | 2015-06-20 21:51:20 | 2021-10-09 18:31:07 |
| [dbus](https://github.com/godbus/dbus) | 622 | 154 | 48 | Native Go bindings for D-Bus | 2014-03-27 19:07:41 | 2021-10-07 18:15:13 |
| [golongpoll](https://github.com/jcuga/golongpoll) | 576 | 45 | 1 | golang long polling library.  Makes web pub-sub easy via HTTP long-poll servers and clients :smiley: :coffee: :computer: | 2015-11-02 00:32:56 | 2021-10-08 02:44:07 |
| [mangos](https://github.com/nanomsg/mangos) | 447 | 60 | 27 | mangos is a pure Golang implementation of nanomsg's "Scalablilty Protocols" | 2018-10-12 17:35:46 | 2021-10-12 02:10:36 |
| [emitter](https://github.com/olebedev/emitter) | 410 | 32 | 4 | Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins | 2015-11-10 20:56:36 | 2021-10-12 16:54:04 |
| [glue](https://github.com/desertbit/glue) | 385 | 29 | 6 | Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) | 2015-06-07 10:21:15 | 2021-09-12 12:17:32 |
| [pubsub](https://github.com/cskr/pubsub) | 359 | 55 | 1 | A simple pubsub package for go. | 2012-04-01 06:31:43 | 2021-10-06 19:07:20 |
| [bus](https://pkg.go.dev/github.com/mustafaturan/bus/v3?tab=doc) | 224 | 15 | 0 | 🔊Minimalist message bus implementation for internal communication with zero-allocation magic on Emit | 2019-04-27 06:41:53 | 2021-10-02 00:41:00 |
| [message-bus](http://rafallorenz.com/message-bus) | 194 | 30 | 2 | Go simple async message bus | 2017-10-04 09:18:34 | 2021-10-12 01:35:06 |
| [rabtap](https://github.com/jandelgado/rabtap) | 189 | 13 | 6 | RabbitMQ wire tap and swiss army knife | 2017-11-11 11:32:39 | 2021-10-10 19:51:51 |
| [guble](https://github.com/smancke/guble) | 148 | 20 | 5 | websocket based messaging server written in golang | 2015-11-15 20:32:42 | 2021-09-07 13:09:18 |
| [hub](https://github.com/leandro-lugaresi/hub) | 111 | 10 | 2 | :incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications | 2018-04-13 23:47:13 | 2021-09-30 05:07:46 |
| [oplog](https://github.com/dailymotion/oplog) | 108 | 11 | 2 | A generic oplog/replication system for microservices | 2014-11-06 09:11:15 | 2021-10-10 07:33:11 |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 92 | 22 | 6 | A tiny wrapper over amqp exchanges and queues 🚌 ✨ | 2017-05-07 08:51:11 | 2021-10-04 10:26:25 |
| [drone-line](https://github.com/appleboy/drone-line) | 76 | 16 | 0 | Sending line notifications using a binary, docker or Drone CI. | 2016-09-13 05:21:44 | 2021-09-17 01:36:19 |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 69 | 12 | 2 | A tiny wrapper around NSQ topic and channel :rocket: | 2017-01-15 22:05:13 | 2021-08-30 01:48:37 |
| [go-mq](https://github.com/cheshir/go-mq) | 69 | 12 | 3 | Declare AMQP entities like queues, producers, and consumers in a declarative way. Can be used to work with RabbitMQ. | 2017-06-19 16:16:30 | 2021-10-05 04:56:38 |
| [redisqueue](https://godoc.org/github.com/robinjoseph08/redisqueue) | 64 | 18 | 3 | redisqueue provides a producer and consumer of a queue that uses Redis streams | 2019-07-07 04:36:54 | 2021-09-25 02:45:26 |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 62 | 9 | 1 | RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue | 2016-10-04 21:07:48 | 2021-06-02 05:57:13 |
| [commander](https://github.com/jeroenrinzema/commander) | 57 | 3 | 2 | Build event-driven and event streaming applications with ease | 2018-04-20 12:30:51 | 2021-10-12 06:49:07 |
| [go-notify](https://github.com/TheCreeper/go-notify) | 55 | 9 | 1 | Package notify provides an implementation of the Gnome DBus Notifications Specification. | 2015-03-01 19:21:44 | 2021-09-05 00:31:04 |
| [go-res](https://github.com/jirenius/go-res) | 51 | 6 | 6 | RES Service protocol library for Go | 2018-07-15 09:10:11 | 2021-10-12 10:14:11 |
| [event](https://github.com/agoalofalife/event) | 41 | 8 | 0 | The implementation of the pattern observer | 2017-07-02 12:19:56 | 2021-07-07 02:57:49 |
| [hare](https://github.com/leozz37/hare) | 30 | 4 | 0 | 🐇  Easy to use socket lib for Golang and CLI tool | 2020-12-01 22:30:27 | 2021-09-21 13:55:39 |
| [ami](https://github.com/kak-tus/ami) | 21 | 6 | 0 | Go client to reliable queues based on Redis Cluster Streams | 2018-10-27 10:38:16 | 2021-07-24 13:55:05 |
| [gosd](https://github.com/alexsniffin/gosd) | 18 | 2 | 0 | A library for scheduling when to dispatch a message to a channel | 2020-05-17 23:19:51 | 2021-05-24 02:42:11 |
| [go-vitotrol](https://godoc.org/github.com/maxatome/go-vitotrol) | 16 | 3 | 1 | golang client library to Viessmann Vitotrol web service | 2016-11-03 19:59:43 | 2021-02-19 21:40:43 |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 14 | 0 | 0 | RabbitMQ Reconnection client | 2019-01-14 16:05:44 | 2021-10-09 03:50:00 |
| [jazz](https://github.com/socifi/jazz) | 13 | 1 | 1 | Abstraction layer for simple rabbitMQ connection, messaging and administration | 2018-10-22 12:28:15 | 2021-08-10 09:13:37 |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 9 | 2 | 0 | Gaurun Client written in Go | 2017-06-29 02:50:51 | 2021-06-12 19:53:09 |
</details>

### Microsoft Office


<sup>*Last Update: 2021-10-21 09:25:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [unioffice](https://unidoc.io/unioffice/) | 3,023 | 336 | 27 | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents | 2017-08-29 01:25:48 | 2021-10-20 23:34:23 |
</details>

### Microsoft Office - Microsoft Excel
Libraries for working with Microsoft Excel.

<sup>*Last Update: 2021-07-29 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [excelize](https://github.com/qax-os/excelize) | 9,140 | 958 | 77 | Golang library for reading and writing Microsoft Excel™ (XLSX) files. | 2016-08-29 12:32:12 | 2021-07-29 02:17:50 |
| [xlsx](https://github.com/tealeg/xlsx) | 5,071 | 768 | 47 | Go (golang) library for reading and writing XLSX files.  | 2011-06-28 15:20:28 | 2021-07-28 09:43:44 |
| [xlsx](https://github.com/plandem/xlsx) | 137 | 19 | 11 | Fast and reliable way to work with Microsoft Excel™ [xlsx] files in Golang | 2017-08-26 23:11:38 | 2021-07-22 15:15:07 |
| [go-excel](https://github.com/szyhf/go-excel) | 128 | 22 | 2 | A simple and light excel file reader to read a standard excel as a table faster | 一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。 | 2017-09-03 11:51:58 | 2021-07-27 14:03:02 |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 16 | 4 | 1 | Golang bindings for libxlsxwriter for writing XLSX files | 2017-03-13 04:15:17 | 2021-07-16 03:58:58 |
</details>

### Miscellaneous - Dependency Injection
Libraries for working with dependency injection.

<sup>*Last Update: 2021-08-23 09:25:15*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fx](https://go.uber.org/fx) | 2,159 | 143 | 36 | A dependency injection based application framework for Go. | 2016-10-27 00:25:00 | 2021-08-22 12:21:22 |
| [dig](https://go.uber.org/dig) | 2,038 | 125 | 22 | A reflection based dependency injection toolkit for Go. | 2017-03-21 23:55:50 | 2021-08-22 19:58:41 |
| [container](https://github.com/golobby/container) | 225 | 14 | 2 | A lightweight yet powerful IoC dependency injection container for Go projects | 2019-09-23 16:12:50 | 2021-08-21 21:22:18 |
| [dingo](https://go.uber.org/dig) | 107 | 6 | 9 | Go Dependency Injection Framework | 2018-10-29 08:55:18 | 2021-08-20 09:17:49 |
| [di](https://github.com/goava/di) | 102 | 8 | 1 | 🛠 A full-featured dependency injection container for go programming language. | 2020-02-03 19:06:39 | 2021-08-19 10:42:31 |
| [di](https://pkg.go.dev/github.com/goioc/di/?tab=doc) | 90 | 3 | 0 | Simple and yet powerful Dependency Injection for Go | 2020-06-11 12:28:06 | 2021-08-18 14:34:36 |
| [alice](https://godoc.org/github.com/magic003/alice) | 44 | 3 | 0 | An additive dependency injection container for Golang. | 2017-04-08 16:25:21 | 2021-02-01 22:07:46 |
| [wire](https://github.com/Fs02/wire) | 33 | 7 | 1 | Strict Runtime Dependency Injection for Golang | 2018-07-05 10:42:24 | 2021-08-22 07:00:21 |
| [linker](https://github.com/logrange/linker) | 32 | 5 | 0 | Dependency Injection and Inversion of Control package | 2018-12-04 23:56:34 | 2021-05-06 05:34:51 |
| [gocontainer](https://github.com/vardius/gocontainer) | 14 | 1 | 0 | Simple Dependency Injection Container | 2019-06-06 08:18:07 | 2020-12-10 08:58:38 |
| [kinit](https://github.com/go-kata/kinit) | 6 | 0 | 0 | GO Dependency Injection | 2021-01-24 13:41:41 | 2021-08-01 02:23:02 |
| [nject](https://github.com/BlueOwlOpenSource/nject) | 3 | 1 | 2 | Go dependency injection: nject & npoint | 2018-10-31 18:15:43 | 2021-07-04 12:55:45 |
</details>

### Miscellaneous - Project Layout
Unofficial set of patterns for structuring projects.

<sup>*Last Update: 2021-09-08 15:02:39*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 26,211 | 2,822 | 56 | Standard Go Project Layout | 2017-09-09 16:33:26 | 2021-09-08 07:31:59 |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 1,042 | 103 | 16 | Modern Go Application example | 2018-09-14 12:19:02 | 2021-09-08 02:40:53 |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 471 | 107 | 4 | A Go project template | 2016-12-18 18:22:26 | 2021-09-07 10:28:50 |
| [seed](https://github.com/golang-templates/seed) | 201 | 16 | 0 | Go application GitHub repository template. | 2020-04-30 21:31:36 | 2021-09-06 08:39:41 |
| [scaffold](https://github.com/catchplay/scaffold) | 99 | 20 | 2 | Generate scaffold project layout for Go. | 2018-12-11 10:36:03 | 2021-08-06 00:56:30 |
| [go-todo-backend](https://github.com/Fs02/go-todo-backend) | 94 | 10 | 0 | Go Todo Backend example using modular project layout for product microservice. | 2020-06-25 14:28:50 | 2021-09-08 04:37:05 |
| [go-sample](https://github.com/zitryss/go-sample) | 91 | 20 | 0 | Go Project Sample Layout | 2019-01-24 23:41:46 | 2021-08-31 02:45:35 |
| [gobase](https://github.com/wajox/gobase) | 10 | 3 | 0 | This is a simple skeleton for golang application | 2020-12-15 16:54:20 | 2021-08-20 04:15:42 |
| [inizio](https://github.com/insidieux/inizio) | 8 | 0 | 1 | Golang project standard layout generator | 2021-03-02 20:59:22 | 2021-07-19 21:23:36 |
</details>

### Miscellaneous - Strings
Libraries for working with strings.

<sup>*Last Update: 2023-09-30 09:09:47*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xstrings](https://github.com/huandu/xstrings) | 1,251 | 95 | 0 | Implements string functions widely used in other languages but absent in Go. | 2015-01-06 07:25:26 | 2023-09-29 06:38:26 |
| [strutil](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 197 | 24 | 1 | String utilities for Go | 2018-08-16 06:56:15 | 2023-09-27 08:53:39 |
| [stringy](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 191 | 17 | 1 | Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package. | 2020-04-03 03:34:10 | 2023-09-16 06:40:56 |
</details>

### Miscellaneous - Uncategorized
These libraries were placed here because none of the other categories seemed to fit.

<sup>*Last Update: 2021-07-29 09:25:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 6,561 | 1,101 | 126 | psutil for golang | 2014-04-18 07:35:28 | 2021-07-29 02:17:34 |
| [archiver](https://godoc.org/github.com/mholt/archiver) | 3,272 | 307 | 73 | Easily create & extract archives, and compress & decompress files of various formats | 2016-04-08 22:46:55 | 2021-07-28 05:33:53 |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 1,960 | 109 | 0 | Random fake data generator written in go | 2015-04-24 04:45:59 | 2021-07-28 14:14:49 |
| [gatus](https://status.twinnation.org/) | 1,610 | 92 | 30 | ⛑ Gatus - Automated service health dashboard | 2019-09-04 02:35:40 | 2021-07-29 01:52:17 |
| [gosms](https://github.com/haxpax/gosms) | 1,339 | 131 | 4 | :mailbox_closed: Your own local SMS gateway in Go | 2015-01-23 19:25:55 | 2021-07-22 17:10:15 |
| [go-resiliency](https://godoc.org/github.com/eapache/go-resiliency) | 1,230 | 104 | 3 | Resiliency patterns for golang | 2014-11-29 04:11:32 | 2021-07-28 07:58:03 |
| [base64Captcha](https://captcha.mojotv.cn) | 1,149 | 194 | 9 | captcha of base64 image string | 2017-12-12 12:17:07 | 2021-07-28 03:34:48 |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 968 | 135 | 3 | a generic object pool for golang | 2015-12-28 14:26:23 | 2021-07-27 02:22:51 |
| [llvm](https://llir.github.io/document/) | 735 | 50 | 20 | Library for interacting with LLVM IR in pure Go. | 2014-09-19 11:18:44 | 2021-07-29 00:12:49 |
| [shortid](https://github.com/teris-io/shortid) | 687 | 56 | 0 | Super short, fully unique, non-sequential and URL friendly Ids | 2016-01-04 01:17:10 | 2021-07-26 23:12:03 |
| [ghorg](https://github.com/gabrie30/ghorg) | 465 | 63 | 19 | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Bitbucket, and more | 2018-03-29 02:53:05 | 2021-07-28 17:32:44 |
| [health](https://github.com/dimiro1/health) | 423 | 42 | 4 | An easy to use, extensible health check library for Go applications. | 2016-03-08 23:04:43 | 2021-07-17 17:28:14 |
| [go-conv](https://github.com/cstockton/go-conv) | 370 | 15 | 0 | Fast conversions across various Go types with a simple API. | 2016-10-11 07:41:41 | 2021-06-18 13:19:32 |
| [banner](https://github.com/dimiro1/banner) | 348 | 19 | 2 | An easy way to add useful startup banners into your Go applications | 2016-03-25 21:28:44 | 2021-07-09 20:38:21 |
| [gountries](https://github.com/pariz/gountries) | 311 | 51 | 19 | Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data. | 2016-01-13 08:04:18 | 2021-07-26 01:14:09 |
| [stateless](https://github.com/qmuntal/stateless) | 283 | 15 | 8 | Go library for creating state machines | 2019-09-11 08:19:18 | 2021-07-27 21:17:34 |
| [ffmt](https://github.com/go-ffmt/ffmt) | 233 | 16 | 2 | Golang beautify data display for Humans | 2015-02-14 15:19:45 | 2021-07-28 13:51:24 |
| [shoutrrr](https://containrrr.dev/shoutrrr/) | 220 | 32 | 11 | Notification library for gophers and their furry friends. | 2019-04-11 06:49:34 | 2021-07-29 00:18:19 |
| [lk](https://github.com/hyperboloide/lk) | 210 | 34 | 1 | Simple licensing library for golang. | 2016-07-14 16:06:07 | 2021-07-05 07:19:32 |
| [antch](https://github.com/antchfx/antch) | 209 | 40 | 4 | Antch, a fast, powerful and extensible web crawling & scraping framework for Go | 2017-09-28 05:44:17 | 2021-07-28 16:56:36 |
| [battery](https://github.com/distatus/battery) | 190 | 25 | 8 | cross-platform, normalized battery information library | 2016-03-12 23:03:40 | 2021-07-28 03:16:09 |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 190 | 26 | 1 | An simple, easily extensible and concurrent health-check library for Go services | 2017-08-18 12:48:40 | 2021-07-26 19:37:25 |
| [bitio](https://github.com/icza/bitio) | 162 | 21 | 1 | Optimized bit-level Reader and Writer for Go. | 2016-05-31 10:02:30 | 2021-07-28 22:26:54 |
| [stats](https://github.com/go-playground/stats) | 154 | 18 | 1 | :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... | 2015-09-14 20:20:20 | 2021-07-26 10:04:49 |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 145 | 23 | 7 | Go bindings for unarr (decompression library for RAR, TAR, ZIP and 7z archives) | 2015-11-01 09:38:37 | 2021-07-27 19:03:10 |
| [turtle](https://github.com/hackebrot/turtle) | 123 | 10 | 1 | Emojis for Go 😄🐢🚀 | 2017-09-08 22:25:32 | 2021-07-24 15:40:31 |
| [gommit](https://github.com/antham/gommit) | 94 | 2 | 1 | Enforce git message commit consistency | 2016-08-30 11:10:11 | 2021-07-23 06:19:53 |
| [gotoprom](https://github.com/cabify/gotoprom) | 90 | 1 | 0 | Type-safe Prometheus metrics builder library for golang | 2018-10-10 16:07:33 | 2021-06-09 15:01:42 |
| [indigo](https://github.com/osamingo/indigo) | 86 | 11 | 1 | A distributed unique ID generator of using Sonyflake and encoded by Base58 | 2016-08-31 14:17:45 | 2021-07-03 07:28:16 |
| [captcha](https://pkg.go.dev/github.com/steambap/captcha) | 79 | 14 | 0 | :sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation | 2017-09-12 06:52:15 | 2021-07-28 13:27:03 |
| [morse](https://github.com/alwindoss/morse) | 68 | 11 | 3 | Morse Code Library in Go | 2018-08-15 05:31:31 | 2021-06-13 18:40:07 |
| [persian](https://github.com/mavihq/persian) | 56 | 8 | 1 | Some utilities for Persian language in Go (Golang) | 2017-10-16 16:16:56 | 2021-07-04 10:30:04 |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 52 | 9 | 0 | HTTP service to generate PDF from Json requests | 2015-11-30 19:27:26 | 2021-05-26 01:21:01 |
| [xkg](https://godoc.org/github.com/go-xkg/xkg) | 51 | 5 | 1 | User level X Keyboard Grabber | 2015-01-05 01:04:43 | 2021-07-13 17:03:22 |
| [faker](https://github.com/pioz/faker) | 41 | 2 | 0 | Random fake data and struct generator for Go. | 2020-07-22 20:09:46 | 2021-07-25 04:53:09 |
| [browscap_go](http://browscap.org/) | 37 | 22 | 6 | GoLang Library for Browser Capabilities Project | 2014-09-18 04:47:42 | 2021-02-19 11:55:57 |
| [datacounter](https://github.com/miolini/datacounter) | 36 | 5 | 2 | Golang counters for readers/writers | 2015-10-14 19:15:50 | 2020-12-23 16:28:56 |
| [autoflags](http://pkg.go.dev/github.com/artyom/autoflags) | 34 | 2 | 0 | Populate go command line app flags from config struct | 2014-05-15 19:00:29 | 2021-07-04 12:55:48 |
| [sandid](https://pkg.go.dev/github.com/aofei/sandid) | 30 | 5 | 0 | Every grain of sand on Earth has its own ID. | 2018-06-12 01:24:14 | 2021-04-23 00:12:32 |
| [url-shortener](https://github.com/pantrif/url-shortener) | 30 | 6 | 0 | A golang URL Shortener | 2018-06-04 05:57:45 | 2021-07-26 10:04:54 |
| [gosh](https://github.com/osamingo/gosh) | 26 | 1 | 0 | Provide Go Statistics Handler, Struct, Measure Method | 2018-05-25 08:55:55 | 2021-07-11 02:11:29 |
| [xdg](https://github.com/rkoesters/xdg) | 26 | 7 | 1 | FreeDesktop.org (xdg) Specs implemented in Go | 2013-12-15 09:51:51 | 2021-07-18 10:03:57 |
| [metrics](https://github.com/pascaldekloe/metrics) | 20 | 4 | 1 | atomic measures + Prometheus exposition library | 2019-01-29 09:39:18 | 2021-06-09 15:02:13 |
| [shellwords](https://pkg.go.dev/github.com/aofei/sandid) | 15 | 2 | 0 | A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. | 2017-09-28 09:08:28 | 2021-06-21 00:37:38 |
| [anagent](https://github.com/mudler/anagent) | 13 | 3 | 0 | Minimalistic, pluggable Golang evloop/timer handler with dependency-injection | 2017-12-29 17:16:25 | 2021-06-22 11:45:30 |
| [avgRating](https://github.com/kirillDanshin/avgRating) | 10 | 1 | 0 | Calculate average score and rating based on Wilson Score Equation | 2017-08-05 19:04:30 | 2020-01-09 09:02:53 |
| [hostutils](https://github.com/Wing924/hostutils) | 9 | 3 | 0 | A golang library for packing and unpacking hosts list | 2017-09-26 03:47:32 | 2020-04-23 10:54:09 |
| [numa](https://github.com/lrita/numa) | 6 | 2 | 0 | NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. | 2018-12-10 09:59:13 | 2021-04-14 12:02:34 |
| [generators](https://github.com/azr/generators) | 4 | 1 | 0 | #golang generator | 2016-02-29 14:29:02 | 2019-06-30 00:44:30 |
</details>

### Natural Language Processing
Libraries for working with human languages.

<sup>*Last Update: 2021-08-31 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prose](https://github.com/jdkato/prose) | 2,786 | 132 | 16 | :book: A Golang library for text processing, including tokenization, part-of-speech tagging, and named-entity extraction. | 2017-02-17 17:08:22 | 2021-08-29 20:22:28 |
| [go-i18n](https://github.com/nicksnyder/go-i18n) | 1,742 | 180 | 15 | Translate your Go program into multiple languages. | 2012-01-14 21:44:37 | 2021-08-29 11:28:24 |
| [gse](https://github.com/go-ego/gse) | 1,629 | 133 | 6 | Go efficient text segmentation and NLP; support english, chinese, japanese and other. Go 语言高性能分词 | 2017-06-23 15:42:35 | 2021-08-30 15:55:44 |
| [gojieba](https://github.com/yanyiwu/gojieba) | 1,583 | 232 | 46 | "结巴"中文分词的Golang版本 | 2015-09-12 01:30:44 | 2021-08-30 06:10:14 |
| [when](https://github.com/olebedev/when) | 1,126 | 64 | 14 | A natural language date/time parser with pluggable rules | 2016-12-27 13:11:46 | 2021-08-29 15:27:52 |
| [go-pinyin](https://godoc.org/github.com/mozillazg/go-pinyin) | 1,006 | 152 | 12 | 汉字转拼音 | 2014-11-09 14:04:33 | 2021-08-27 23:38:22 |
| [spago](https://github.com/nlpodyssey/spago) | 953 | 42 | 10 | Self-contained Machine Learning and Natural Language Processing library in Go | 2020-01-05 20:39:29 | 2021-08-27 19:15:56 |
| [kagome](https://github.com/ikawaha/kagome) | 580 | 39 | 2 | Self-contained Japanese Morphological Analyzer written in pure Go | 2014-06-26 04:38:13 | 2021-08-30 17:33:36 |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 504 | 50 | 11 | Natural language detection library for Go | 2017-02-20 17:32:01 | 2021-08-26 19:11:12 |
| [nlp](https://github.com/shixzie/nlp) | 369 | 28 | 3 | [UNMANTEINED] Extract values from strings and fill your structs with nlp. | 2017-01-25 07:19:03 | 2021-06-05 13:26:18 |
| [nlp](https://github.com/james-bowman/nlp) | 336 | 40 | 3 | Selected Machine Learning algorithms for natural language processing and semantic analysis in Golang | 2017-03-15 08:28:05 | 2021-08-20 16:35:47 |
| [sentences](https://sentences-231000.appspot.com/) | 309 | 30 | 2 | A multilingual command line sentence tokenizer in Golang | 2015-08-07 01:08:20 | 2021-08-28 19:37:16 |
| [getlang](https://github.com/rylans/getlang) | 119 | 16 | 4 | Natural language detection package in pure Go | 2018-03-01 21:27:30 | 2021-07-27 14:35:18 |
| [go-unidecode](https://godoc.org/github.com/mozillazg/go-unidecode) | 89 | 13 | 3 | ASCII transliterations of Unicode text. | 2016-07-08 13:15:10 | 2021-07-06 08:27:20 |
| [go-nlp](https://github.com/nuance/go-nlp) | 89 | 10 | 0 | Utilities for working with discrete probability distributions and other tools useful for doing NLP work | 2011-05-02 06:43:36 | 2021-01-30 05:58:16 |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 82 | 16 | 4 | A Go port of the Rapid Automatic Keyword Extraction algorithm (RAKE) | 2016-12-17 13:36:25 | 2021-08-10 03:42:24 |
| [gounidecode](https://github.com/fiam/gounidecode) | 73 | 19 | 2 | Unicode transliterator for #golang | 2012-05-01 11:59:34 | 2021-05-09 09:43:14 |
| [textcat](https://github.com/pebbe/textcat) | 65 | 8 | 1 | A Go package for n-gram based text categorization, with support for utf-8 and raw text | 2012-09-21 15:04:45 | 2021-06-21 14:35:11 |
| [segment](https://github.com/blevesearch/segment) | 64 | 11 | 5 | A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 | 2014-10-16 19:24:26 | 2021-08-09 04:03:56 |
| [go-stem](svn://go-stem) | 64 | 15 | 1 | Word Stemming in Go | 2011-09-23 19:07:23 | 2021-08-08 10:04:07 |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 59 | 13 | 0 | Chinese word splitting algorithm MMSEG in GO | 2012-04-18 04:06:21 | 2020-08-05 02:49:45 |
| [stemmer](https://godoc.org/github.com/dchest/stemmer) | 49 | 4 | 0 | Stemmer packages for Go programming language. Includes English, German and Dutch stemmers. | 2011-03-21 02:08:12 | 2021-04-29 22:02:55 |
| [go2vec](https://godoc.org/github.com/mozillazg/go-unidecode) | 41 | 4 | 0 | Read and use word2vec vectors in Go | 2015-01-27 12:02:04 | 2020-10-14 13:06:04 |
| [porter2](http://zhen.org/blog/generating-porter2-fsm-for-fun-and-performance/) | 40 | 5 | 1 | High Performance Porter2 Stemmer | 2015-01-21 07:30:32 | 2021-04-30 13:16:52 |
| [petrovich](https://github.com/striker2000/petrovich) | 34 | 3 | 0 | Golang port of Petrovich - an inflector for Russian anthroponyms. | 2016-12-26 22:50:38 | 2021-06-23 11:55:26 |
| [address](https://pkg.go.dev/github.com/bojanz/address) | 30 | 0 | 0 | Address handling for Go. | 2020-10-07 18:15:27 | 2021-07-19 21:23:38 |
| [snowball](https://github.com/goodsign/snowball) | 28 | 3 | 0 | Cgo binding for Snowball C library | 2012-12-11 12:42:19 | 2021-08-08 10:04:22 |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 5 | 2 | Golang implementation of the Paice/Husk Stemming Algorithm | 2012-09-29 16:06:58 | 2019-11-28 05:32:37 |
| [mystem](https://github.com/dveselov/mystem) | 26 | 6 | 0 | CGo bindings to Yandex.Mystem | 2016-08-30 14:55:39 | 2021-07-05 11:21:11 |
| [go-localize](https://github.com/m1/go-localize) | 25 | 8 | 1 | i18n (Internationalization and localization) engine written in Go, used for translating locale strings.  | 2019-12-23 12:02:51 | 2021-06-24 15:17:10 |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 23 | 2 | 0 | Transliterate Cyrillic → Latin in every possible way | 2020-04-27 09:29:40 | 2021-08-10 13:39:08 |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 19 | 4 | 0 | Go bindings for the snowball libstemmer library including porter 2 | 2012-08-06 19:31:05 | 2021-08-10 21:12:14 |
| [icu](https://github.com/goodsign/icu) | 19 | 4 | 2 | Cgo binding for icu4c library | 2012-12-11 13:09:41 | 2020-10-29 13:45:37 |
| [govader](https://github.com/jonreiter/govader) | 15 | 2 | 0 | vader sentiment analysis in go | 2020-01-19 10:06:15 | 2021-06-22 02:34:30 |
| [transliterator](https://github.com/alexsergivan/transliterator) | 12 | 5 | 1 | Golang text Transliterator (i.e München -> Muenchen) | 2020-04-17 14:19:55 | 2021-06-17 21:02:09 |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 12 | 3 | 0 | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) | 2018-10-11 03:22:36 | 2021-06-10 07:32:49 |
| [shamoji](https://github.com/osamingo/shamoji) | 12 | 0 | 0 | The shamoji (杓文字) is a word filtering package | 2017-07-23 06:38:42 | 2021-01-14 18:13:19 |
| [detectlanguage-go](https://detectlanguage.com) | 10 | 0 | 0 | Detect Language API Go Client | 2019-12-14 23:30:44 | 2021-05-08 08:46:00 |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 6 | 0 | Cgo binding for libtextcat C library | 2012-12-10 21:21:47 | 2018-07-29 06:39:58 |
| [porter](https://github.com/a2800276/porter) | 8 | 0 | 0 | porter stemmer | 2013-09-17 11:10:16 | 2018-07-28 05:12:38 |
| [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) | 7 | 0 | 0 | 💬 Sentiment analyzer library using SentiWordnet in Go | 2020-04-21 09:09:28 | 2021-06-10 04:27:05 |
</details>

### Networking
Libraries for working with various layers of the network.

<sup>*Last Update: 2021-07-29 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 15,650 | 1,283 | 34 | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http | 2015-10-18 22:19:57 | 2021-07-26 21:17:16 |
| [kcptun](https://github.com/xtaci/kcptun) | 12,503 | 2,421 | 70 | A Stable & Secure Tunnel based on KCP with N:M multiplexing and FEC. Available for ARM, MIPS, 386 and AMD64。KCPプロトコルに基づく安全なトンネル。KCP 프로토콜을 기반으로 하는 보안 터널입니다。 | 2016-02-26 09:54:46 | 2021-07-27 00:23:39 |
| [webrtc](https://pion.ly) | 7,572 | 941 | 68 | Pure Go implementation of the WebRTC API | 2018-05-18 23:10:05 | 2021-07-28 18:53:34 |
| [dns](https://miek.nl/2014/august/16/go-dns-package) | 5,673 | 896 | 12 | DNS library in Go | 2010-08-03 21:56:23 | 2021-07-27 01:53:23 |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 5,603 | 729 | 87 | A QUIC implementation in pure go | 2016-04-06 20:16:27 | 2021-07-26 22:46:23 |
| [gnet](https://gnet.host) | 4,904 | 570 | 23 | 🚀 gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。 | 2019-02-24 03:48:45 | 2021-07-26 12:47:35 |
| [gopacket](https://github.com/google/gopacket) | 4,254 | 836 | 183 | Provides packet processing capabilities for Go | 2015-03-16 20:46:00 | 2021-07-27 00:59:52 |
| [httplab](https://github.com/gchaincl/httplab) | 3,708 | 122 | 12 | The interactive web server | 2017-02-08 17:13:19 | 2021-07-26 15:06:15 |
| [kcp-go](https://github.com/xtaci/kcp-go) | 3,029 | 548 | 26 |  A Crypto-Secure, Production-Grade Reliable-UDP Library for golang with FEC  | 2015-06-16 06:15:55 | 2021-07-26 14:31:17 |
| [gobgp](https://osrg.github.io/gobgp/) | 2,248 | 498 | 107 | BGP implemented in the Go Programming Language | 2014-09-14 01:51:58 | 2021-07-25 11:49:25 |
| [ssh](https://godoc.org/github.com/gliderlabs/ssh) | 2,092 | 247 | 31 | Easy SSH servers in Golang | 2016-10-03 21:53:44 | 2021-07-26 17:24:13 |
| [fortio](https://fortio.org) | 1,983 | 164 | 74 | Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats. | 2017-10-10 01:01:39 | 2021-07-26 17:07:56 |
| [water](https://github.com/songgao/water) | 1,269 | 196 | 20 | A simple TUN/TAP library written in native Go. | 2013-03-25 20:06:52 | 2021-07-20 15:44:26 |
| [go-getter](https://github.com/hashicorp/go-getter) | 1,204 | 149 | 97 | Package for downloading things from a string URL using a variety of protocols. | 2015-10-12 23:17:07 | 2021-07-22 15:13:13 |
| [gev](https://github.com/Allenxuxu/gev) | 1,201 | 143 | 5 | 🚀Gev is a lightweight, fast non-blocking TCP network library / websocket server based on Reactor mode. Support custom protocols to quickly and easily build high-performance servers.  | 2019-09-01 12:16:18 | 2021-07-26 01:39:01 |
| [nff-go](https://github.com/intel-go/nff-go) | 1,074 | 127 | 61 | NFF-Go -Network Function Framework for GO (former YANFF) | 2017-03-29 17:07:29 | 2021-07-23 08:14:21 |
| [sftp](https://github.com/pkg/sftp) | 1,028 | 289 | 18 | SFTP support for the go.crypto/ssh package | 2013-11-05 04:36:00 | 2021-07-27 01:37:55 |
| [grab](https://github.com/cavaliercoder/grab) | 879 | 108 | 25 | A download manager package for Go | 2016-01-05 12:46:35 | 2021-07-26 19:24:45 |
| [ftp](https://github.com/jlaffaye/ftp) | 818 | 277 | 4 | FTP client package for Go | 2011-05-06 18:31:51 | 2021-07-23 03:57:46 |
| [mdns](https://github.com/hashicorp/mdns) | 784 | 172 | 32 | Simple mDNS client/server library in Golang | 2014-01-29 19:39:18 | 2021-07-22 08:37:50 |
| [gosnmp](https://github.com/gosnmp/gosnmp) | 740 | 243 | 25 | An SNMP library written in Go | 2012-08-27 05:59:24 | 2021-07-26 19:22:18 |
| [vssh](https://github.com/yahoo/vssh) | 728 | 53 | 0 | Go Library to Execute Commands Over SSH at Scale | 2020-06-09 16:19:22 | 2021-07-26 03:17:11 |
| [lhttp](https://github.com/fanux/lhttp) | 624 | 126 | 6 | go websocket, a better way to buid your IM server | 2015-12-29 01:13:36 | 2021-07-09 11:47:41 |
| [cidranger](https://github.com/yl2chen/cidranger) | 622 | 71 | 6 | Fast IP to CIDR lookup in Golang | 2017-08-21 05:50:14 | 2021-07-23 05:19:05 |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 494 | 35 | 6 | Pure-Go library for cross-platform local peer discovery using UDP multicast :woman: :repeat: :woman: | 2018-04-22 23:59:37 | 2021-07-26 20:13:45 |
| [gotcp](https://github.com/gansidui/gotcp) | 486 | 157 | 0 | A Go package for quickly building tcp servers | 2014-04-13 14:54:01 | 2021-07-23 09:37:43 |
| [stun](https://github.com/gortc/stun) | 467 | 48 | 4 | Fast RFC 5389 STUN implementation in go | 2016-04-24 17:46:38 | 2021-07-18 02:07:53 |
| [go-stun](https://github.com/ccding/go-stun) | 448 | 83 | 2 | A go implementation of the STUN client (RFC 3489 and RFC 5389) | 2013-08-17 22:16:33 | 2021-07-24 02:46:45 |
| [gopcap](https://github.com/akrennmair/gopcap) | 426 | 136 | 12 | A simple wrapper around libpcap for the Go programming language | 2009-11-19 10:13:48 | 2021-07-19 21:23:37 |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 413 | 77 | 3 | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.x and V5 in golang | 2018-09-16 11:46:17 | 2021-07-26 02:53:06 |
| [raw](https://github.com/mdlayher/raw) | 398 | 69 | 13 | Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. | 2015-07-06 16:11:47 | 2021-07-20 16:58:16 |
| [tcp_server](https://github.com/firstrow/tcp_server) | 384 | 134 | 4 | golang tcp server | 2014-10-13 20:38:42 | 2021-07-22 17:16:15 |
| [gaio](https://github.com/xtaci/gaio) | 372 | 39 | 12 | High performance async-io(proactor) networking for Golang。golangのための高性能非同期io(proactor)ネットワーキング | 2019-12-20 05:19:00 | 2021-07-27 01:36:19 |
| [winrm](https://pion.ly) | 317 | 86 | 24 | Command-line tool and library for Windows remote command execution in Go | 2013-12-30 18:29:15 | 2021-07-17 18:11:01 |
| [arp](https://tools.ietf.org/html/rfc826) | 258 | 42 | 3 | Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. | 2015-07-06 18:50:34 | 2021-07-19 03:06:32 |
| [ftpserverlib](https://github.com/fclairamb/ftpserverlib) | 254 | 64 | 3 | golang ftp server library | 2016-09-25 12:05:29 | 2021-07-25 20:23:59 |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 245 | 32 | 7 | A library to simplify writing applications using TCP sockets to stream protobuff messages | 2015-06-29 19:07:31 | 2021-07-22 17:13:41 |
| [ethernet](https://en.wikipedia.org/wiki/Ethernet_frame) | 226 | 30 | 0 | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. MIT Licensed. | 2015-07-03 00:15:18 | 2021-07-20 16:57:13 |
| [nbio](https://github.com/lesismal/nbio) | 218 | 23 | 0 | High-performance, non-blocking, event-driven, easy-to-use networking framework written in Go, support tls/http1.x/websocket. | 2020-01-25 11:46:54 | 2021-07-25 19:11:57 |
| [gnxi](https://github.com/google/gnxi) | 185 | 88 | 13 | gNXI Tools - gRPC Network Management/Operations Interface Tools | 2017-09-26 08:19:41 | 2021-07-25 19:04:51 |
| [jazigo](https://github.com/udhos/jazigo) | 167 | 14 | 3 | Jazigo is a tool written in Go for retrieving configuration for multiple devices, similar to rancid, fetchconfig, oxidized, Sweet. | 2016-06-07 19:53:53 | 2021-07-17 00:41:00 |
| [utp](https://github.com/anacrolix/go-libutp) | 159 | 33 | 4 | Use anacrolix/go-libutp instead | 2015-03-20 04:39:22 | 2021-07-08 07:43:12 |
| [canopus](https://github.com/zubairhamed/canopus) | 145 | 39 | 43 | CoAP Client/Server implementing RFC 7252 for the Go Language | 2015-02-24 04:12:20 | 2021-07-14 08:38:24 |
| [sslb](https://godoc.org/github.com/gliderlabs/ssh) | 130 | 25 | 10 | Golang Super Simple Load Balance | 2015-10-18 21:31:09 | 2021-02-11 17:27:49 |
| [xtcp](https://github.com/xfxdev/xtcp) | 122 | 27 | 0 | A TCP Server Framework with graceful shutdown, custom protocol. | 2016-03-31 16:50:14 | 2021-07-28 16:19:21 |
| [dhcp6](https://tools.ietf.org/html/rfc3315) | 70 | 17 | 2 | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. | 2015-05-22 04:13:30 | 2021-04-30 17:24:35 |
| [ether](https://github.com/songgao/ether) | 69 | 5 | 0 | A Go package for sending and receiving ethernet frames. Currently supporting Linux, Freebsd, and OS X. | 2014-05-21 03:46:30 | 2021-02-28 14:18:18 |
| [packet](https://github.com/aerogo/packet) | 56 | 13 | 1 | :package: Send network packets over a TCP or UDP connection. | 2017-10-29 05:46:44 | 2021-04-07 01:43:57 |
| [linkio](https://github.com/ian-kent/linkio) | 49 | 5 | 0 | Simulate network link speed | 2014-12-24 10:50:03 | 2021-04-20 21:41:59 |
| [iplib](https://github.com/c-robinson/iplib) | 46 | 5 | 0 | A library  for working with IP addresses and networks in Go | 2019-05-06 06:23:41 | 2021-07-18 02:06:03 |
| [portproxy](https://github.com/aybabtme/portproxy) | 45 | 10 | 0 | TCP proxy, highjacks HTTP to allow CORS | 2014-12-13 02:57:36 | 2021-05-29 02:46:21 |
| [go-powerdns](https://pkg.go.dev/github.com/joeig/go-powerdns/v2) | 36 | 11 | 0 | Go PowerDNS 4.x API Client | 2018-06-21 21:37:33 | 2021-06-24 17:50:11 |
| [graval](https://github.com/koofr/graval) | 27 | 5 | 0 | An experimental go FTP server framework | 2014-04-22 19:17:18 | 2021-07-18 02:02:32 |
| [publicip](https://github.com/polera/publicip) | 23 | 5 | 0 | Go pkg for returning your public facing IP address. | 2016-12-28 19:31:07 | 2021-05-20 14:01:48 |
| [panoptes-stream](https://github.com/yahoo/panoptes-stream) | 23 | 3 | 1 | A cloud native distributed streaming network telemetry. | 2020-10-09 04:26:26 | 2021-06-09 13:06:57 |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 20 | 5 | 0 | GoHooks make it easy to send and consume secured web-hooks from a Go application | 2015-11-16 06:48:41 | 2021-06-26 20:43:16 |
| [gohooks](https://github.com/averageflow/gohooks) | 12 | 1 | 0 | GoHooks make it easy to send and consume secured web-hooks from a Go application | 2020-10-30 17:20:36 | 2021-07-16 09:56:57 |
| [goshark](https://github.com/sunwxg/goshark) | 11 | 2 | 0 | A simple wrapper around libpcap for the Go programming language | 2015-11-01 07:23:09 | 2021-04-27 05:55:24 |
| [llb](https://github.com/kirillDanshin/llb) | 11 | 0 | 0 | Simulate network link speed | 2016-02-21 06:30:17 | 2020-09-30 19:25:46 |
| [tspool](https://github.com/two/tspool) | 10 | 1 | 0 | tcp server pool | 2018-10-27 01:05:03 | 2021-07-18 04:02:29 |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 9 | 1 | 0 | HTTP proxy handler and dialer | 2018-07-18 09:42:34 | 2021-03-02 01:56:42 |
</details>

### Networking - HTTP Clients
Libraries for making HTTP requests.

<sup>*Last Update: 2021-07-19 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [resty](https://github.com/go-resty/resty) | 4,542 | 378 | 40 | Simple HTTP and REST client library for Go | 2015-08-28 17:48:47 | 2021-07-17 21:55:35 |
| [heimdall](http://gojek.tech) | 1,984 | 162 | 31 | An enhanced HTTP client for Go | 2018-01-19 09:32:26 | 2021-07-18 16:37:19 |
| [grequests](https://github.com/levigross/grequests) | 1,803 | 106 | 28 | A Go "clone" of the great and famous Requests library | 2015-06-11 16:41:48 | 2021-07-14 14:45:16 |
| [sling](https://github.com/dghubble/sling) | 1,335 | 105 | 0 | A Go HTTP client library for creating and sending API requests | 2015-04-02 08:42:52 | 2021-07-14 09:36:55 |
| [gentleman](https://pkg.go.dev/github.com/h2non/gentleman?tab=doc) | 905 | 49 | 18 | Plugin-driven, extensible HTTP client toolkit for Go | 2016-02-21 23:00:24 | 2021-07-15 01:09:33 |
| [pester](https://github.com/sethgrid/pester) | 556 | 61 | 4 | Go (golang) http calls with retries and backoff  | 2015-05-20 13:50:49 | 2021-07-17 16:19:30 |
| [request](https://pkg.go.dev/github.com/monaco-io/request?tab=doc) | 129 | 17 | 0 | go request, go http client | 2020-03-25 06:24:18 | 2021-07-17 17:28:22 |
| [rq](https://github.com/ddo/rq) | 39 | 2 | 1 | A nicer interface for golang stdlib HTTP client | 2017-12-26 10:48:27 | 2021-06-15 12:31:58 |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 24 | 5 | 0 | An enhanced http client for Golang | 2019-12-14 11:22:19 | 2021-07-14 14:39:17 |
| [httpretry](https://github.com/ybbus/httpretry) | 14 | 1 | 0 | Enriches the standard go http client with retry functionality. | 2020-02-05 10:17:42 | 2021-05-20 19:41:31 |
</details>

### ORM
Libraries that implement Object-Relational Mapping or datamapping techniques.

<sup>*Last Update: 2021-09-03 09:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gorm](https://gorm.io) | 25,043 | 2,841 | 123 | The fantastic ORM library for Golang, aims to be developer friendly | 2013-10-25 08:31:38 | 2021-09-03 01:34:58 |
| [ent](https://entgo.io) | 8,278 | 432 | 161 | An entity framework for Go | 2019-06-12 22:53:55 | 2021-09-03 02:05:30 |
| [pg](https://pg.uptrace.dev/) | 4,806 | 353 | 99 | Golang ORM with focus on PostgreSQL features and performance | 2013-04-24 12:31:41 | 2021-09-02 23:37:57 |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 4,209 | 398 | 108 | Generate a Go ORM tailored to your database schema. | 2016-02-21 06:18:25 | 2021-09-02 21:22:16 |
| [gorp](https://github.com/go-gorp/gorp) | 3,519 | 373 | 135 | Go Relational Persistence - an ORM-ish library for Go | 2012-01-04 19:50:09 | 2021-08-31 14:08:53 |
| [db](https://upper.io/) | 2,713 | 190 | 127 | Data access layer for PostgreSQL, CockroachDB, MySQL, SQLite and MongoDB with ORM-like features. | 2013-10-23 02:04:36 | 2021-09-02 06:26:48 |
| [gormt](https://xxjwxc.github.io/post/gormt/) | 1,395 | 236 | 31 | database to golang struct | 2019-05-05 13:10:26 | 2021-09-03 01:30:27 |
| [reform](https://gopkg.in/reform.v1) | 1,175 | 51 | 69 | A better ORM for Go, based on non-empty interfaces and code generation. | 2016-02-25 09:41:09 | 2021-09-02 11:49:49 |
| [pop](https://github.com/gobuffalo/pop) | 1,114 | 220 | 120 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2021-09-02 08:59:46 |
| [go-queryset](https://github.com/jirfag/go-queryset) | 633 | 59 | 19 | 100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood. | 2017-09-03 17:29:30 | 2021-08-27 07:18:50 |
| [go-sqlbuilder](https://pkg.go.dev/github.com/huandu/go-sqlbuilder) | 625 | 62 | 3 | A flexible and powerful SQL string builder library plus a zero-config ORM. | 2017-12-27 16:37:48 | 2021-09-01 12:43:22 |
| [qbs](https://github.com/coocood/qbs) | 547 | 104 | 9 | QBS stands for Query By Struct. A Go ORM. | 2013-02-02 05:40:59 | 2021-08-23 06:42:23 |
| [rel](https://go-rel.github.io/) | 409 | 39 | 17 | :gem: Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API | 2019-10-06 07:08:01 | 2021-09-02 11:18:54 |
| [zoom](https://github.com/albrow/zoom) | 278 | 24 | 2 | A blazing-fast datastore and querying engine for Go built on Redis. | 2013-07-17 00:32:34 | 2021-07-19 07:28:29 |
| [beego](beego.me) | 237 | 70 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2021-09-02 18:23:54 |
| [grimoire](https://fs02.github.io/grimoire) | 150 | 15 | 0 | Database access layer for golang | 2018-03-05 16:52:20 | 2021-08-22 07:06:58 |
| [gosql](https://github.com/rushteam/gosql) | 144 | 15 | 2 | golang orm and sql builder | 2020-04-27 09:16:29 | 2021-08-30 08:13:14 |
| [go-store](https://github.com/gosuri/go-store) | 102 | 8 | 1 | A simple and fast Redis backed key-value store library for Go | 2015-03-22 12:07:29 | 2021-06-05 22:34:23 |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 24 | 6 | 3 | Simple Go ORM for Google/Firebase Cloud Firestore | 2018-12-04 14:53:53 | 2021-08-17 10:39:57 |
| [lore](https://github.com/abrahambotros/lore) | 8 | 1 | 0 | Light Object-Relational Environment (LORE) provides a simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go | 2017-04-29 03:57:15 | 2021-08-07 04:49:01 |
| [marlow](https://github.com/marlow/marlow) | 6 | 0 | 0 | persistence layer code generation for golang | 2020-08-11 13:34:00 | 2021-05-22 15:11:30 |
</details>

### OpenGL
Libraries for using OpenGL in Go.

<sup>*Last Update: 2021-07-17 19:25:24*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glfw](https://github.com/go-gl/glfw) | 1,118 | 136 | 13 | Go bindings for GLFW 3 | 2013-05-19 06:38:45 | 2021-07-17 05:15:01 |
| [gl](https://github.com/go-gl/gl) | 818 | 56 | 11 | Go bindings for OpenGL (generated via glow) | 2015-02-22 03:29:45 | 2021-07-16 07:20:09 |
| [mathgl](https://github.com/go-gl/mathgl) | 383 | 50 | 9 | A pure Go 3D math library. | 2013-02-13 14:18:55 | 2021-07-14 11:28:20 |
| [gl](https://github.com/goxjs/gl) | 150 | 16 | 9 | Go cross-platform OpenGL bindings. | 2015-05-18 08:10:15 | 2021-06-04 02:46:58 |
| [glfw](https://github.com/goxjs/glfw) | 70 | 15 | 10 | Go cross-platform glfw library for creating an OpenGL context and receiving events. | 2014-12-27 22:40:24 | 2021-04-22 20:14:17 |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 2 | 2 | 0 | go-glmatrix is a golang version of glMatrix, which is "designed to perform vector and matrix operations stupidly fast". | 2020-07-02 13:40:40 | 2021-03-17 16:22:40 |
</details>

### Package Management - Official
Official experimental tooling for package management

<sup>*Last Update: 2021-08-19 09:25:26*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [dep](https://golang.github.io/dep/) | 13,145 | 1,107 | 0 | Go dependency management tool experiment (deprecated) | 2016-10-07 00:04:51 | 2021-08-18 18:35:51 |
</details>

### Package Management - Unofficial
Unofficial libraries for package and dependency management

<sup>*Last Update: 2021-08-19 09:25:18*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glide](https://glide.sh) | 8,115 | 541 | 422 | Package Management for Golang | 2014-07-09 06:02:50 | 2021-08-18 18:13:55 |
| [godep](http://godoc.org/github.com/tools/godep) | 5,609 | 486 | 79 | dependency tool for go | 2013-05-01 07:55:35 | 2021-08-16 02:51:35 |
| [govendor](https://blog.golang.org/migrating-to-go-modules) | 5,001 | 411 | 122 | Use Go Modules. | 2015-04-12 15:26:40 | 2021-08-17 02:43:12 |
| [gopm](https://github.com/gpmgo/gopm) | 2,496 | 215 | 0 | Go Package Manager (gopm) is a package manager and build tool for Go. | 2013-05-15 14:53:29 | 2021-08-18 22:34:26 |
| [gom](https://github.com/mattn/gom) | 1,395 | 104 | 14 | Go Manager - bundle for go | 2013-09-11 02:08:59 | 2021-08-08 20:57:42 |
| [gpm](https://github.com/pote/gpm) | 1,197 | 51 | 15 | Barebones dependency manager for Go. | 2013-09-05 02:24:02 | 2021-07-22 02:52:24 |
| [goop](https://github.com/petejkim/goop) | 778 | 46 | 29 | A simple dependency manager for Go (golang), inspired by Bundler. | 2014-06-18 01:55:24 | 2021-07-18 17:50:48 |
| [modgv](https://github.com/lucasepe/modgv) | 386 | 11 | 1 | Converts 'go mod graph' output into Graphviz's DOT language | 2020-09-12 16:23:46 | 2021-08-16 13:35:29 |
| [nut](https://github.com/jingweno/nut) | 239 | 11 | 14 | Vendor Go dependencies | 2015-01-23 14:46:32 | 2021-07-13 14:35:54 |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 215 | 6 | 3 | Barebones dependency manager for Go. | 2013-07-19 15:20:47 | 2020-12-15 17:40:11 |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 131 | 27 | 0 | maven plugin to automate GoSDK load and build of projects | 2016-03-24 06:47:08 | 2021-08-13 22:22:35 |
| [VenGO](https://github.com/DamnWidget/VenGO) | 120 | 10 | 3 | Create and manage Isolated Virtual Environments for Go | 2014-10-17 15:19:03 | 2021-04-21 01:54:12 |
| [gop](https://github.com/lunny/gop) | 49 | 6 | 10 | Moved to https://gitea.com/lunny/gop | 2017-02-18 04:33:48 | 2021-01-14 20:53:39 |
</details>

### Performance


<sup>*Last Update: 2021-07-19 09:25:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [jaeger](https://www.jaegertracing.io/) | 13,861 | 1,641 | 381 | CNCF Jaeger, a Distributed Tracing Platform | 2016-04-15 18:49:02 | 2021-07-19 02:14:58 |
| [pixie](https://px.dev/) | 1,684 | 75 | 67 | Instant Kubernetes-Native Application Observability | 2020-02-27 00:22:45 | 2021-07-17 19:20:35 |
| [profile](https://px.dev/) | 1,504 | 103 | 8 | Simple profiling for Go | 2014-10-22 01:35:18 | 2021-07-18 10:56:38 |
| [statsviz](https://github.com/arl/statsviz) | 1,399 | 51 | 6 | :rocket: Instant live visualization of your Go application runtime statistics (GC, MemStats, etc.) in the browser | 2020-08-14 00:00:41 | 2021-07-16 08:38:44 |
| [tracer](https://github.com/kamilsk/tracer) | 47 | 1 | 11 | 🧶 Dead simple, lightweight tracing. | 2019-06-22 13:23:27 | 2021-07-11 13:21:00 |
</details>

### Query Language


<sup>*Last Update: 2021-07-17 19:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [graphql](https://github.com/graphql-go/graphql) | 7,817 | 681 | 170 | An implementation of GraphQL for Go / Golang | 2015-07-19 12:25:43 | 2021-07-17 07:32:28 |
| [gqlgen](https://gqlgen.com) | 6,264 | 674 | 239 | go generate based graphql server library | 2018-02-11 04:54:11 | 2021-07-17 08:05:20 |
| [graphql-go](https://github.com/graph-gophers/graphql-go) | 3,804 | 410 | 86 | GraphQL server with a focus on ease of use | 2016-10-18 13:57:24 | 2021-07-17 07:06:09 |
| [gojsonq](https://github.com/thedevsaddam/gojsonq/wiki) | 1,702 | 99 | 11 | A simple Go package to Query over JSON/YAML/XML/CSV Data  | 2018-05-19 16:15:18 | 2021-07-14 15:33:50 |
| [dasel](https://daseldocs.tomwright.me) | 951 | 18 | 6 | Query, update and convert data structures from the command line. Comparable to jq/yq but supports JSON, TOML, YAML, XML and CSV with zero runtime dependencies. | 2020-09-22 10:33:56 | 2021-07-17 01:40:58 |
| [jsonql](https://github.com/elgs/jsonql) | 245 | 36 | 5 | JSON query expression library in Golang. | 2015-12-29 11:24:04 | 2021-07-14 02:47:25 |
| [rql](https://github.com/a8m/rql) | 216 | 24 | 11 | Resource Query Language for REST | 2018-06-05 18:37:29 | 2021-07-17 10:20:58 |
| [graphql](https://github.com/tmc/graphql) | 53 | 6 | 3 | graphql parser + utilities | 2015-04-18 21:05:52 | 2021-05-27 16:58:59 |
| [jsonslice](https://github.com/bhmj/jsonslice) | 52 | 4 | 4 | json slicer | 2018-05-02 00:33:15 | 2021-04-03 11:04:51 |
| [api-fu](https://github.com/ccbrown/api-fu) | 33 | 0 | 2 | A collection of Go packages for creating robust GraphQL APIs | 2019-07-30 05:18:43 | 2021-06-25 09:30:51 |
| [straf](https://github.com/SonicRoshan/straf) | 29 | 3 | 0 | Convert Golang Struct To GraphQL Object On The Fly | 2019-08-16 13:31:39 | 2021-07-15 13:50:19 |
| [rest-query-parser](https://github.com/timsolov/rest-query-parser) | 24 | 4 | 1 | Query Parser for REST | 2020-02-10 17:58:42 | 2021-07-14 08:43:37 |
| [jsonpath](https://github.com/AsaiYusuke/jsonpath) | 5 | 1 | 0 | A query library for retrieving part of JSON based on JSONPath syntax. | 2020-11-29 05:37:26 | 2021-04-27 16:37:54 |
| [gws](https://github.com/Zaba505/gws) | 4 | 1 | 2 | A WebSocket client and server for GraphQL | 2020-06-08 19:51:36 | 2020-11-16 02:47:54 |
</details>

### Resource Embedding


<sup>*Last Update: 2021-08-23 09:25:12*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [packr](https://github.com/gobuffalo/packr) | 3,257 | 180 | 66 | The simple and easy way to embed static files into Go binaries. | 2017-03-15 22:24:53 | 2021-08-23 01:47:30 |
| [statik](https://github.com/rakyll/statik) | 3,238 | 206 | 33 | Embed files into a Go executable | 2014-02-04 14:54:51 | 2021-08-19 08:53:42 |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2,235 | 141 | 36 | go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. | 2013-10-23 21:29:34 | 2021-08-21 05:04:22 |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 930 | 80 | 32 | Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. | 2015-05-18 13:03:02 | 2021-08-20 22:31:11 |
| [esc](http://godoc.org/github.com/mjibson/esc) | 602 | 66 | 11 | A simple file embedder for Go | 2014-01-26 05:08:04 | 2021-08-22 13:33:12 |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 588 | 50 | 9 | a better customizable tool to embed files in go; also update embedded files remotely without restarting the server | 2016-01-23 20:19:33 | 2021-08-06 06:55:34 |
| [go-resources](https://github.com/omeid/go-resources) | 172 | 18 | 3 | Unfancy resources embedding for Go with out of box http.FileSystem support. | 2015-02-21 15:40:17 | 2021-06-14 12:38:59 |
| [statics](https://github.com/go-playground/statics) | 63 | 4 | 0 | :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks | 2015-10-07 11:49:52 | 2021-07-19 11:37:53 |
| [templify](https://github.com/wlbr/templify) | 26 | 4 | 1 | A tool to be used with 'go generate' to embed external template files into Go code. | 2016-05-22 16:42:47 | 2021-08-16 20:22:53 |
| [rebed](https://github.com/soypat/rebed) | 18 | 2 | 0 | Recreates directory and files from embedded filesystem using Go 1.16 embed.FS type. | 2021-02-17 18:19:49 | 2021-08-16 21:07:29 |
| [mule](https://github.com/wlbr/mule) | 9 | 1 | 1 | mule is a tool to be used with 'go generate' to embed external resources files into Go code. | 2020-01-17 10:56:00 | 2021-08-16 20:23:31 |
| [debme](https://github.com/leaanthony/debme) | 8 | 1 | 0 | embed.FS wrapper providing additional functionality | 2021-04-16 00:25:13 | 2021-08-09 05:05:31 |
</details>

### Science and Data Analysis
Libraries for scientific computing and data analyzing.

<sup>*Last Update: 2021-09-15 09:25:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gonum](https://www.gonum.org/) | 5,169 | 410 | 218 | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more | 2017-03-25 14:54:38 | 2021-09-14 19:08:38 |
| [stats](https://github.com/montanaflynn/stats) | 2,040 | 131 | 15 | A well tested and comprehensive Golang statistics library package with no dependencies. | 2014-12-16 03:25:19 | 2021-09-14 11:46:54 |
| [plot](https://github.com/gonum/plot) | 1,985 | 181 | 90 | A repository for plotting and visualizing data | 2013-07-23 07:01:13 | 2021-09-13 17:05:55 |
| [gosl](https://github.com/cpmech/gosl) | 1,604 | 145 | 0 | Linear algebra, eigenvalues, FFT, Bessel, elliptic, orthogonal polys, geometry, NURBS, numerical quadrature, 3D transfinite interpolation, random numbers, Mersenne twister, probability distributions, optimisation, differential equations. | 2015-02-09 23:00:38 | 2021-09-12 06:48:13 |
| [streamtools](http://nytlabs.github.io/streamtools/) | 1,313 | 109 | 47 | tools for working with streams of data | 2013-07-05 18:58:45 | 2021-08-28 20:06:23 |
| [go-dsp](http://godoc.org/github.com/mjibson/go-dsp) | 748 | 76 | 6 | Digital Signal Processing for Go | 2011-11-02 06:28:41 | 2021-09-14 04:16:36 |
| [chart](https://github.com/vdobler/chart) | 689 | 103 | 6 | Provide basic charts in go | 2011-06-27 12:19:42 | 2021-09-14 02:20:00 |
| [goraph](https://github.com/gyuho/goraph) | 648 | 74 | 6 | Package goraph implements graph data structure and algorithms. | 2014-02-27 03:15:55 | 2021-09-11 07:18:01 |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 609 | 60 | 7 | DataFrames for Go: For statistics, machine-learning, and data manipulation/exploration | 2018-10-01 12:19:31 | 2021-09-12 23:31:28 |
| [graph](https://yourbasic.org/golang/your-basic-func/) | 477 | 48 | 2 | Graph algorithms and data structures | 2017-04-27 18:43:54 | 2021-09-13 18:54:44 |
| [orb](https://github.com/paulmach/orb) | 414 | 54 | 13 | Types and utilities for working with 2d geometry in Golang | 2016-03-28 01:19:01 | 2021-09-12 16:11:12 |
| [ewma](https://github.com/VividCortex/ewma) | 338 | 25 | 4 | Exponentially Weighted Moving Average algorithms for Go. | 2013-07-05 21:33:25 | 2021-09-12 02:26:36 |
| [calendarheatmap](https://calendarheatmap.io) | 326 | 11 | 8 | Calendar heatmap inspired by GitHub contribution activity | 2020-07-01 18:30:48 | 2021-09-13 19:44:22 |
| [gohistogram](http://godoc.org/github.com/VividCortex/gohistogram) | 157 | 28 | 1 | Streaming approximate histograms in Go | 2013-07-02 12:53:22 | 2021-08-20 17:37:14 |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 143 | 17 | 5 | :wink: :cyclone: :strawberry: TextRank implementation in Golang with extendable features (summarization, phrase extraction) and multithreading (goroutine). | 2018-01-09 19:36:17 | 2021-08-27 11:56:31 |
| [sparse](https://github.com/james-bowman/sparse) | 121 | 19 | 4 | Sparse matrix formats for linear algebra supporting scientific and machine learning applications | 2017-05-16 13:54:36 | 2021-08-22 17:08:48 |
| [go-estimate](https://github.com/milosgajdos/go-estimate) | 76 | 6 | 2 | State estimation and filtering algorithms in Go | 2018-11-04 22:32:52 | 2021-08-22 18:33:16 |
| [pagerank](https://github.com/alixaxel/pagerank) | 72 | 17 | 3 | Weighted PageRank implementation in Go | 2015-08-06 01:33:34 | 2021-09-13 05:14:55 |
| [geom](https://github.com/skelterjohn/geom) | 48 | 16 | 1 | 2d geometry for golang | 2011-06-07 17:49:11 | 2021-07-24 18:47:42 |
| [evaler](https://github.com/soniah/evaler) | 45 | 12 | 5 | Implements a simple floating point arithmetic expression evaluator in Go (golang). | 2012-09-04 23:37:58 | 2021-08-19 22:58:15 |
| [goent](https://github.com/kzahedi/goent) | 25 | 1 | 0 | GO Implementation of Entropy Measures | 2017-08-08 05:37:12 | 2021-05-10 16:11:38 |
| [triangolatte](https://github.com/tchayen/triangolatte) | 24 | 1 | 4 | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. | 2018-07-18 21:17:09 | 2021-08-04 11:32:59 |
| [decimal](https://github.com/db47h/decimal) | 23 | 0 | 0 | An arbitrary-precision decimal floating-point arithmetic package for Go | 2020-05-27 15:23:59 | 2021-08-13 00:54:49 |
| [piecewiselinear](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 20 | 1 | 0 | tiny linear interpolation library for go (factored out from https://github.com/sgreben/yeetgif) | 2018-10-21 13:19:44 | 2021-09-03 23:11:02 |
| [GoStats](https://github.com/OGFris/GoStats) | 16 | 0 | 0 | GoStats is a go library for math statistics mostly used in ML domains, it covers most of the statistical measures functions. | 2018-07-22 20:55:16 | 2021-08-06 07:28:59 |
| [godesim](https://github.com/soypat/godesim) | 15 | 0 | 1 | ODE system solver made simple. For IVPs (initial value problems). | 2020-12-16 01:02:26 | 2021-09-07 15:32:37 |
| [PiHex](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 14 | 2 | 0 | PiHex Library, written in Go, generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000. | 2016-07-22 11:21:37 | 2021-03-18 18:52:26 |
| [ode](https://yourbasic.org/golang/your-basic-func/) | 13 | 0 | 1 | An ordinary differential equation solving library in golang. | 2016-11-11 22:40:21 | 2021-01-26 11:20:02 |
| [assocentity](https://pkg.go.dev/github.com/ndabAP/assocentity/v8?tab=doc) | 5 | 1 | 6 | Package assocentity returns the average distance from words to a given entity | 2018-12-21 07:17:09 | 2021-08-06 17:15:44 |
| [rootfinding](https://github.com/khezen/rootfinding) | 5 | 0 | 0 | root-finding library | 2018-10-30 22:31:48 | 2021-05-31 18:37:41 |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2 | Automatically exported from code.google.com/p/go-gt | 2015-09-14 12:05:37 | 2019-09-05 08:47:43 |
| [bradleyterry](https://pkg.go.dev/github.com/ndabAP/assocentity/v8?tab=doc) | 4 | 0 | 0 | Package to do Bradley-Terry Model pairwise compairsons | 2019-04-30 00:28:13 | 2020-09-08 12:32:53 |
</details>

### Security
Libraries that are used to help make your application more secure.

<sup>*Last Update: 2021-08-19 09:25:07*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [lego](https://go-acme.github.io/lego/) | 4,747 | 631 | 119 | Let's Encrypt client and ACME library written in Go | 2015-06-08 00:36:41 | 2021-08-17 20:06:51 |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2,652 | 368 | 14 | Cameradar hacks its way into RTSP videosurveillance cameras | 2016-05-20 11:35:41 | 2021-08-18 18:45:33 |
| [crypto](https://golang.org/x/crypto) | 2,224 | 1,230 | 44 | [mirror] Go supplementary cryptography libraries | 2014-12-04 04:02:55 | 2021-08-18 14:26:45 |
| [memguard](https://github.com/awnumar/memguard) | 1,972 | 90 | 4 | Secure software enclave for storage of sensitive information in memory. | 2017-04-22 07:40:40 | 2021-08-15 17:28:09 |
| [acmetool](https://hlandau.github.io/acmetool/) | 1,867 | 118 | 67 | :lock: acmetool, an automatic certificate acquisition tool for ACME (Let's Encrypt) | 2015-11-15 01:56:02 | 2021-08-13 23:20:22 |
| [secure](https://github.com/unrolled/secure) | 1,804 | 114 | 3 | HTTP middleware for Go that facilitates some quick security wins. | 2014-05-20 19:46:28 | 2021-08-17 02:45:48 |
| [themis](https://www.cossacklabs.com/themis) | 1,327 | 112 | 9 | Easy to use cryptographic framework for data protection: secure messaging with forward secrecy and secure data storage. Has unified APIs across 14 platforms. | 2015-05-06 13:25:25 | 2021-08-16 20:38:30 |
| [acra](https://www.cossacklabs.com/acra/) | 785 | 91 | 7 | Database security suite. Database proxy with field-level encryption, search through encrypted data, SQL injections prevention, intrusion detection, honeypots. Supports client-side and proxy-side ("transparent") encryption. SQL, NoSQL. | 2016-11-14 16:23:25 | 2021-08-14 12:00:46 |
| [nacl](https://godoc.org/github.com/kevinburke/nacl) | 508 | 29 | 3 | Pure Go implementation of the NaCL set of API's | 2017-07-20 19:07:19 | 2021-08-17 10:50:27 |
| [firewalld-rest](https://pkg.go.dev/github.com/prashantgupta24/firewalld-rest) | 306 | 12 | 1 | A rest application to update firewalld rules on a linux server  | 2020-06-12 20:16:33 | 2021-08-18 05:20:56 |
| [ssh-vault](https://ssh-vault.com) | 303 | 22 | 4 | 🌰  encrypt/decrypt using ssh keys | 2016-09-29 14:46:30 | 2021-08-13 09:31:46 |
| [badactor](https://badactor.org) | 300 | 15 | 1 | BadActor.org An in-memory application driven jailer written in Go | 2014-12-12 20:05:20 | 2021-08-15 01:47:12 |
| [go-password-validator](https://qvault.io/2020/10/15/how-to-correctly-validate-passwords-most-websites-do-it-wrong/) | 287 | 20 | 0 | Validate the Strength of a Password in Go | 2020-10-14 15:52:14 | 2021-08-18 20:52:41 |
| [optimus-go](https://github.com/pjebs/optimus-go) | 283 | 19 | 0 | ID hashing and Obfuscation using Knuth's Algorithm | 2015-05-05 10:12:38 | 2021-08-17 10:14:20 |
| [passlib](https://github.com/hlandau/passlib) | 248 | 26 | 1 | :key: Idiotproof golang password validation library inspired by Python's passlib | 2014-12-21 17:45:52 | 2021-08-11 08:18:43 |
| [go-yara](https://github.com/hillu/go-yara) | 219 | 80 | 5 | Go bindings for YARA | 2015-01-25 01:01:11 | 2021-08-08 04:23:42 |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 172 | 20 | 4 | A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go 🔑 | 2015-04-14 06:52:21 | 2021-07-13 00:27:25 |
| [argon2pw](https://github.com/raja/argon2pw) | 85 | 7 | 0 | Argon2 password hashing package for go with constant time hash comparison | 2018-03-13 13:56:36 | 2021-07-14 22:00:34 |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 42 | 6 | 0 | A probably paranoid Golang utility library for securely hashing and encrypting passwords based on the Dropbox method. This implementation uses Blake2b, Scrypt and XSalsa20-Poly1305 (via NaCl SecretBox) to create secure password hashes that are also encrypted using a master passphrase. | 2017-10-19 19:34:45 | 2021-08-08 03:48:29 |
| [go-generate-password](https://github.com/m1/go-generate-password) | 26 | 2 | 0 | Password generator written in Go | 2019-11-14 17:57:19 | 2021-08-15 09:25:37 |
| [certificates](https://github.com/mvmaasakkers/certificates) | 21 | 3 | 0 | An opinionated helper for generating tls certificates | 2019-03-04 07:20:36 | 2021-07-09 07:31:10 |
| [secureio](https://github.com/xaionaro-go/secureio) | 18 | 1 | 1 | An easy-to-use XChaCha20-encryption wrapper for io.ReadWriteCloser (even lossy UDP) using ECDH key exchange algorithm, ED25519 signatures and Blake3+Poly1305 checksums/message-authentication for Go (golang). Also a multiplexer. | 2018-12-25 14:20:59 | 2021-07-19 08:47:58 |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 14 | 5 | 1 | goArgonPass is a Argon2 Password utility package for Go using the crypto library package Argon2 designed to be compatible with Passlib for Python and Argon2 PHP. Argon2 was the winner of the most recent Password Hashing Competition. This is designed for use anywhere password hashing and verification might be needed and is intended to replace implementations using bcrypt or Scrypt. | 2018-05-30 01:32:10 | 2021-02-01 13:35:19 |
| [argon2-hashing](https://www.cossacklabs.com/acra/) | 13 | 1 | 0 | A light package for generating and comparing password hashing with argon2 in Go | 2019-01-02 20:41:02 | 2021-08-02 13:27:00 |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 11 | 2 | 0 | A layer of abstraction the around acme/autocert certificate manager (Golang) | 2019-04-02 17:35:38 | 2021-06-10 17:50:47 |
</details>

### Serialization
Libraries and tools for binary serialization.

<sup>*Last Update: 2021-09-07 09:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go](http://jsoniter.com/migrate-from-go-std.html) | 9,763 | 796 | 156 | A high-performance 100% compatible drop-in replacement of "encoding/json" | 2016-11-30 00:30:24 | 2021-09-06 16:30:20 |
| [protobuf](https://github.com/golang/protobuf) | 7,884 | 1,437 | 56 | Go support for Google's protocol buffers | 2014-11-23 23:07:23 | 2021-09-07 01:53:29 |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 4,826 | 501 | 40 | Go library for decoding generic map values into native Go structures and vice versa. | 2013-05-20 05:24:34 | 2021-09-06 06:14:49 |
| [protobuf](https://github.com/gogo/protobuf) | 4,794 | 626 | 211 | [Looking for new ownership] Protocol Buffers for Go with Gadgets | 2014-12-03 11:27:10 | 2021-09-06 21:18:48 |
| [go](https://github.com/ugorji/go) | 1,577 | 262 | 3 | idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] | 2013-05-30 02:13:13 | 2021-09-03 08:54:19 |
| [colfer](https://github.com/pascaldekloe/colfer) | 618 | 48 | 12 | binary serialization format | 2015-09-05 16:42:41 | 2021-09-04 12:15:50 |
| [csvutil](https://github.com/jszwec/csvutil) | 567 | 33 | 2 | csvutil provides fast and idiomatic mapping between CSV and Go (golang) values. | 2017-10-30 04:09:48 | 2021-09-06 16:49:23 |
| [cbor](https://github.com/fxamacker/cbor) | 314 | 26 | 9 | CBOR codec (in Go) with CBOR tags, Go struct tags (toarray/keyasint/omitempty), float64/32/16, big.Int, and fuzz tested billions of execs for reliable RFC 7049 & RFC 8949.  | 2019-05-15 21:22:15 | 2021-09-06 20:12:13 |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 279 | 18 | 1 | Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. | 2013-11-07 20:28:12 | 2021-06-14 12:15:27 |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 151 | 41 | 3 | PHP session encoder/decoder written in Go | 2012-12-23 14:04:25 | 2021-08-12 13:22:37 |
| [structomap](https://godoc.org/github.com/danhper/structomap) | 125 | 10 | 0 | Easily and dynamically generate maps from Go static structures | 2015-05-13 08:54:11 | 2021-08-27 20:58:29 |
| [bambam](https://github.com/glycerine/bambam) | 62 | 10 | 3 | auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto | 2014-09-17 14:39:12 | 2021-03-04 20:35:06 |
| [asn1](https://github.com/Logicalis/asn1) | 47 | 22 | 6 | Asn.1 BER and DER encoding library for golang. | 2016-02-29 13:00:25 | 2020-08-21 06:42:00 |
| [binstruct](https://github.com/ghostiam/binstruct) | 33 | 5 | 0 | Golang binary decoder for mapping data into the structure | 2018-10-23 15:42:22 | 2021-08-23 08:00:01 |
| [elastic](https://github.com/epiclabs-io/elastic) | 15 | 1 | 1 | Converts go types no matter what | 2020-02-25 19:55:00 | 2021-05-20 08:50:22 |
| [fwencoder](https://github.com/o1egl/fwencoder) | 14 | 3 | 0 | Fixed width file parser (encoder/decoder) in GO (golang) | 2017-12-25 12:55:29 | 2020-09-10 10:26:54 |
| [pletter](https://github.com/vimeda/pletter) | 14 | 1 | 3 | A standard way to wrap a proto message | 2019-07-09 14:02:08 | 2021-05-21 04:32:24 |
| [bel](https://github.com/csweichel/bel) | 13 | 5 | 2 | Generate TypeScript interfaces from Go structs/interfaces - useful for JSON RPC | 2019-02-20 20:48:24 | 2021-08-04 07:31:18 |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 5 | 0 | 0 | A Go package for encode/decode fixed-width data | 2019-08-11 03:42:24 | 2020-12-07 08:51:25 |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 2 | 0 | 0 | go-lctree provides a CLI and Go primitives to serialize and deserialize LeetCode binary trees (e.g. "[5,4,7,3,null,2,null,-1,null,9]"). | 2020-05-04 05:39:46 | 2020-12-24 05:48:12 |
| [unitpacking](https://github.com/recolude/unitpacking) | 2 | 0 | 0 | A library for storing unit vectors in a representation that lends itself to saving space on disk. | 2021-01-17 22:31:41 | 2021-07-06 08:31:13 |
</details>

### Server Applications


<sup>*Last Update: 2023-10-02 20:50:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [caddy](https://caddyserver.com) | 49,948 | 3,818 | 141 | Fast and extensible multi-platform HTTP/1-2-3 web server with automatic HTTPS | 2015-01-13 19:45:03 | 2023-10-02 13:22:44 |
| [etcd](https://etcd.io) | 44,595 | 9,493 | 236 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2023-10-02 13:37:40 |
| [minio](https://min.io/download) | 40,981 | 4,932 | 41 | High Performance Object Storage for AI | 2015-01-14 19:23:58 | 2023-10-02 13:11:17 |
| [consul](https://www.consul.io) | 26,945 | 4,388 | 1,246 | Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. | 2013-11-04 22:15:27 | 2023-10-02 10:39:09 |
| [nsq](https://nsq.io) | 23,802 | 2,890 | 65 | A realtime distributed messaging platform | 2012-05-12 14:37:08 | 2023-10-02 12:28:11 |
| [roadrunner](https://roadrunner.dev) | 7,388 | 400 | 52 | 🤯 High-performance PHP application server, process manager written in Go and powered with plugins | 2017-12-26 16:13:10 | 2023-10-02 09:52:08 |
| [sftpgo](https://github.com/drakkan/sftpgo) | 6,815 | 586 | 18 | Fully featured and highly configurable SFTP server with optional HTTP/S, FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-07-20 10:18:31 | 2023-10-02 08:42:20 |
| [devd](https://github.com/cortesi/devd) | 3,350 | 150 | 25 |  A local webserver for developers | 2015-09-27 22:43:00 | 2023-09-25 12:24:40 |
| [flipt](https://flipt.io) | 2,838 | 143 | 18 | An open source, self-hosted feature flag solution | 2016-11-05 00:09:07 | 2023-10-01 16:09:03 |
| [fider](https://fider.io) | 2,426 | 584 | 47 | Open platform to collect and prioritize feedback | 2017-01-17 22:55:19 | 2023-10-02 11:24:48 |
| [algernon](https://algernon.roboticoverlords.org) | 2,297 | 121 | 22 | Small self-contained pure-Go web server with Lua, Teal, Markdown, HTTP/2, QUIC, Redis and PostgreSQL support | 2015-03-10 11:25:30 | 2023-09-30 15:51:46 |
| [flagr](https://openflagr.github.io/flagr) | 2,234 | 184 | 2 | Flagr is a feature flagging, A/B testing and dynamic configuration microservice | 2017-10-03 19:07:32 | 2023-10-01 08:36:46 |
| [trickster](https://trickstercache.org) | 1,909 | 171 | 57 | Open Source HTTP Reverse Proxy Cache and Time Series Dashboard Accelerator | 2018-03-29 20:31:44 | 2023-09-29 03:35:29 |
| [discovery](https://github.com/bilibili/discovery) | 1,760 | 399 | 26 | A registry for resilient mid-tier load balancing and failover. | 2018-04-20 12:57:50 | 2023-09-27 09:28:48 |
| [jackal](https://github.com/ortuman/jackal) | 1,429 | 132 | 22 | 💬 Instant messaging server for the Extensible Messaging and Presence Protocol (XMPP). | 2017-11-13 18:17:48 | 2023-10-02 00:29:32 |
| [go-feature-flag](https://gofeatureflag.org) | 832 | 70 | 9 | GO Feature Flag is a simple, complete and lightweight self-hosted feature flag solution 100% Open Source. 🎛️ | 2020-12-11 13:19:17 | 2023-10-02 00:56:11 |
| [dudeldu](https://github.com/krotik/dudeldu) | 141 | 17 | 0 | A simple SHOUTcast server. | 2016-09-07 19:11:04 | 2023-09-12 19:18:45 |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 100 | 13 | 25 | Simple Reverse Proxy with Caching, written in Go, using Redis. | 2020-11-12 15:10:40 | 2023-09-18 15:18:38 |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 82 | 16 | 40 | Reverse proxy with automatically obtains TLS certificates from Let's Encrypt | 2019-04-12 05:39:43 | 2023-09-18 15:18:41 |
| [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) | 75 | 49 | 5 | Prometheus remote write proxy that adds Cortex/Mimir tenant ID based on metric labels | 2020-10-06 16:52:25 | 2023-09-28 06:44:51 |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 53 | 13 | 2 | Stream database events from PostgreSQL to Kafka | 2019-04-28 21:55:31 | 2023-09-18 15:18:50 |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 38 | 5 | 0 | Turn Nginx logs into Prometheus metrics | 2018-10-23 09:10:27 | 2023-09-18 15:18:48 |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 35 | 5 | 5 | Fully featured and highly configurable SFTP server with optional HTTP/S, FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-12-18 12:48:14 | 2023-09-18 15:19:35 |
| [protoxy](https://github.com/camgraff/protoxy) | 32 | 5 | 0 | A proxy server than converts JSON request bodies to protocol buffers | 2020-09-03 23:24:34 | 2023-09-28 09:34:44 |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 2 | 2 | 0 | Service for relaying Riemann events to Riemann/Carbon destinations | 2019-04-23 14:17:12 | 2022-09-27 09:23:35 |
</details>

### Stream Processing
Libraries and tools for stream processing and reactive programming.

<sup>*Last Update: 2021-10-21 09:25:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-streams](https://pkg.go.dev/github.com/reugn/go-streams) | 732 | 55 | 2 | A lightweight stream processing library for Go | 2019-04-30 17:28:15 | 2021-10-20 11:54:11 |
| [machine](https://pkg.go.dev/github.com/whitaker-io/machine) | 99 | 6 | 6 | Machine is a workflow/pipeline library for processing data | 2020-10-13 04:24:19 | 2021-10-14 20:09:31 |
| [stream](https://github.com/youthlin/stream) | 45 | 2 | 0 | Go Stream, like Java 8 Stream. | 2020-11-12 03:52:50 | 2021-10-20 16:57:44 |
</details>

### Template Engines
Libraries and tools for templating and lexing.

<sup>*Last Update: 2023-09-30 16:55:33*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gofpdf](http://godoc.org/github.com/jung-kurt/gofpdf) | 4,190 | 760 | 56 | A PDF document generator with high level support for text, drawing and images | 2015-03-13 11:57:30 | 2023-09-28 02:38:07 |
| [sprig](http://masterminds.github.io/sprig/) | 3,755 | 421 | 126 | Useful template functions for Go templates. | 2013-11-22 01:20:40 | 2023-09-29 20:34:00 |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 2,868 | 149 | 37 | Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template | 2016-03-06 21:42:01 | 2023-09-29 22:13:47 |
| [pongo2](https://www.schlachter.tech/pongo2) | 2,652 | 284 | 61 | Django-syntax like template-engine for Go | 2013-08-23 01:00:08 | 2023-09-27 19:30:49 |
| [hero](https://shiyanhui.github.io/hero) | 1,544 | 107 | 28 | A handy, fast and powerful go template engine. | 2017-01-15 13:31:50 | 2023-09-26 14:27:53 |
| [maroto](https://maroto.io) | 1,218 | 153 | 19 | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. | 2019-05-20 23:27:47 | 2023-09-30 09:34:06 |
| [mustache](https://github.com/hoisie/mustache) | 1,089 | 234 | 33 | The mustache template language in Go | 2009-12-30 21:05:05 | 2023-09-27 19:38:49 |
| [jet](https://shiyanhui.github.io/hero) | 1,080 | 100 | 22 | Jet  template engine | 2016-03-31 16:53:36 | 2023-09-27 20:32:00 |
| [amber](https://github.com/eknkc/amber) | 908 | 61 | 24 | Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade | 2012-10-31 20:27:24 | 2023-09-25 08:48:51 |
| [gorazor](https://github.com/sipin/gorazor) | 830 | 90 | 2 | Razor view engine for go | 2014-05-01 05:30:31 | 2023-09-25 10:58:54 |
| [ace](https://github.com/yosssi/ace) | 827 | 71 | 30 | HTML template engine for Go | 2014-07-13 13:39:19 | 2023-09-12 09:48:49 |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 770 | 77 | 12 | Simple and fast template engine for Go | 2015-08-19 12:44:22 | 2023-09-27 10:19:34 |
| [ego](https://github.com/benbjohnson/ego) | 563 | 42 | 10 | An ERB-style templating language for Go. | 2014-02-23 18:14:41 | 2023-09-08 21:30:20 |
| [raymond](https://github.com/aymerick/raymond) | 552 | 92 | 19 | Handlebars for golang | 2015-04-22 13:07:59 | 2023-09-27 18:05:17 |
| [goview](https://github.com/foolin/goview) | 371 | 34 | 15 | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. | 2019-04-14 11:22:41 | 2023-09-15 06:27:47 |
| [liquid](https://godoc.org/github.com/osteele/liquid) | 233 | 47 | 26 | A Liquid template engine in Go | 2017-06-26 14:39:52 | 2023-09-25 08:49:30 |
| [soy](https://github.com/robfig/soy) | 170 | 42 | 7 | Go implementation for Soy templates (Google Closure templates) | 2013-12-15 01:14:48 | 2023-09-16 19:15:14 |
| [kasia.go](https://github.com/ziutek/kasia.go) | 74 | 9 | 2 | Templating system for HTML and other text documents - go implementation | 2010-12-07 10:46:05 | 2022-03-15 21:35:36 |
| [velvet](http://masterminds.github.io/sprig/) | 72 | 10 | 2 | A sweet velvety templating package | 2016-12-29 16:46:57 | 2023-03-01 11:33:12 |
| [extemplate](https://git.sr.ht/~dvko/extemplate) | 53 | 17 | 1 | Wrapper package for Go's template/html to allow for easy file-based template inheritance. | 2018-08-10 20:34:19 | 2023-06-14 00:13:05 |
| [gospin](https://github.com/m1/gospin) | 47 | 7 | 3 | Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations | 2019-02-22 17:04:51 | 2023-09-19 17:42:46 |
| [damsel](https://github.com/dskinner/damsel) | 23 | 6 | 1 | Package damsel provides html outlining via css-selectors and common template functionality. | 2012-05-02 23:06:48 | 2023-03-17 14:34:28 |
</details>

### Testing - Fail injection


<sup>*Last Update: 2021-08-19 09:25:26*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [failpoint](https://github.com/pingcap/failpoint) | 599 | 54 | 11 | An implementation of failpoints for Golang. | 2019-04-02 07:48:18 | 2021-08-17 10:45:50 |
</details>

### Testing - Fuzzing and delta-debugging, reducing, shrinking


<sup>*Last Update: 2021-10-13 09:25:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 4,155 | 241 | 52 | Randomized testing for Go | 2015-04-15 13:07:50 | 2021-10-13 00:50:59 |
| [gofuzz](https://github.com/google/gofuzz) | 1,160 | 108 | 12 | Fuzz testing for go. | 2014-07-31 16:21:29 | 2021-10-12 15:50:57 |
| [tavor](https://github.com/zimmski/tavor) | 233 | 9 | 53 | A generic fuzzing and delta-debugging framework | 2014-05-18 14:59:14 | 2021-10-01 18:11:25 |
</details>

### Testing - Mock


<sup>*Last Update: 2021-07-17 19:25:09*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mock](https://pkg.go.dev/github.com/h2non/gock) | 5,858 | 443 | 32 | GoMock is a mocking framework for the Go programming language. | 2015-06-12 17:15:11 | 2021-07-17 03:50:39 |
| [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) | 3,651 | 290 | 51 | Sql mock driver for golang to test database interactions | 2014-02-07 07:59:29 | 2021-07-16 16:33:07 |
| [hoverfly](https://hoverfly.io) | 1,759 | 173 | 33 | Lightweight service virtualization/API simulation tool for developers and testers | 2015-11-30 16:36:31 | 2021-07-12 04:02:00 |
| [gock](https://pkg.go.dev/github.com/h2non/gock) | 1,435 | 77 | 28 | HTTP traffic mocking and testing made easy in Go ༼ʘ̚ل͜ʘ̚༽ | 2016-03-02 16:20:26 | 2021-07-16 06:35:52 |
| [httpmock](http://godoc.org/github.com/jarcoal/httpmock) | 1,104 | 82 | 1 | HTTP mocking for Golang | 2014-02-24 16:47:59 | 2021-07-13 10:17:14 |
| [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) | 545 | 68 | 14 | A tool for generating self-contained, type-safe test doubles in go | 2014-05-21 00:12:54 | 2021-07-16 15:33:06 |
| [minimock](https://github.com/gojuno/minimock) | 371 | 24 | 9 | Powerful mock generation tool for Go programming language | 2016-08-03 16:01:35 | 2021-07-17 03:29:37 |
| [go-txdb](https://github.com/DATA-DOG/go-txdb) | 369 | 33 | 4 | Immutable transaction isolated sql driver for golang | 2015-07-08 07:34:53 | 2021-07-14 19:03:31 |
| [govcr](https://github.com/seborama/govcr) | 98 | 13 | 4 | HTTP mock for Golang: record and replay HTTP/HTTPS interactions for offline testing | 2016-07-10 17:47:41 | 2021-05-17 01:46:44 |
| [timex](https://github.com/cabify/timex) | 55 | 2 | 0 | A test-friendly replacement for golang's time package | 2020-01-02 18:06:48 | 2021-06-29 15:31:48 |
| [mockhttp](https://github.com/tv42/mockhttp) | 22 | 5 | 0 | Mock object for Go http.ResponseWriter | 2011-06-11 16:03:01 | 2020-08-05 04:12:58 |
| [go-localstack](https://github.com/elgohr/go-localstack) | 12 | 3 | 1 | Go Wrapper for using localstack | 2020-03-18 07:13:02 | 2021-07-16 19:35:17 |
| [mockit](https://github.com/pasdam/mockit) | 8 | 1 | 1 | Library that make mocking of Go functions/methods easy | 2020-01-01 08:46:09 | 2021-06-22 02:44:40 |
</details>

### Testing - Selenium and browser control tools


<sup>*Last Update: 2023-09-30 09:09:47*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [chromedp](https://github.com/chromedp/chromedp) | 6,978 | 577 | 44 | A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. | 2017-01-24 14:54:30 | 2021-11-25 03:28:48 |
| [selenoid](https://aerokube.com/selenoid/latest/) | 2,417 | 320 | 263 | Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary. | 2016-08-22 09:11:16 | 2023-09-29 07:25:21 |
| [rod](https://go-rod.github.io) | 1,937 | 130 | 75 | A Devtools driver for web automation and scraping | 2020-01-21 20:09:45 | 2021-11-24 08:14:08 |
| [cdp](https://github.com/mafredri/cdp) | 581 | 40 | 12 | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 2017-03-12 10:25:41 | 2021-11-22 19:11:49 |
| [playwright-go](https://mxschmitt.github.io/playwright-go/) | 563 | 56 | 13 | Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API. | 2020-08-16 12:46:14 | 2021-11-24 22:06:03 |
| [ggr](https://aerokube.com/ggr/latest/) | 275 | 58 | 15 | A lightweight load balancer used to create big Selenium clusters | 2016-06-16 15:33:24 | 2021-11-22 19:39:19 |
</details>

### Testing - Testing Frameworks


<sup>*Last Update: 2021-10-21 09:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [testify](https://github.com/stretchr/testify) | 14,566 | 1,135 | 296 | A toolkit with common assertions and mocks that plays nicely with the standard library | 2012-10-16 16:43:17 | 2021-10-21 01:57:46 |
| [goconvey](http://goconvey.co) | 6,710 | 484 | 145 | Go testing in the browser. Integrates with `go test`. Write behavioral tests in Go. | 2013-08-21 04:52:28 | 2021-10-21 02:17:18 |
| [ginkgo](http://onsi.github.io/ginkgo/) | 5,189 | 464 | 51 | BDD Testing Framework for Go | 2013-08-23 20:53:05 | 2021-10-21 02:18:11 |
| [go-cmp](https://github.com/google/go-cmp) | 2,536 | 156 | 15 | Package for comparing Go values in tests | 2017-07-07 19:28:22 | 2021-10-20 20:16:16 |
| [httpexpect](https://github.com/gavv/httpexpect) | 1,798 | 141 | 9 | End-to-end HTTP and REST API testing for Go. | 2016-04-29 17:05:20 | 2021-10-20 19:59:50 |
| [godog](https://github.com/cucumber/godog) | 1,485 | 156 | 44 | Cucumber for golang | 2015-06-10 13:16:31 | 2021-10-20 23:03:57 |
| [gomega](http://onsi.github.io/gomega/) | 1,482 | 225 | 21 | Ginkgo's Preferred Matcher Library | 2013-08-23 20:54:42 | 2021-10-21 02:07:07 |
| [goblin](https://github.com/franela/goblin) | 816 | 73 | 19 | Minimal and Beautiful Go testing framework | 2013-09-19 02:34:24 | 2021-10-18 20:13:45 |
| [go-vcr](https://go-testdeep.zetta.rocks/) | 783 | 52 | 4 | Record and replay your HTTP interactions for fast, deterministic and accurate tests | 2015-12-14 12:52:17 | 2021-10-17 00:53:42 |
| [baloo](https://godoc.org/github.com/h2non/baloo) | 718 | 28 | 8 | Expressive end-to-end HTTP API testing made easy in Go | 2016-05-29 21:40:58 | 2021-10-18 08:07:49 |
| [testfixtures](https://pkg.go.dev/github.com/go-testfixtures/testfixtures/v3?tab=doc) | 708 | 49 | 13 | Ruby on Rails like test fixtures for Go. Write tests against a real database | 2016-04-05 11:33:28 | 2021-10-19 11:57:55 |
| [gnomock](https://github.com/orlangure/gnomock) | 540 | 24 | 7 | Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻 | 2020-01-31 14:50:52 | 2021-10-20 16:11:21 |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 474 | 36 | 34 | Mutation testing for Go source code | 2014-12-26 22:23:44 | 2021-10-18 17:31:46 |
| [goc](https://github.com/qiniu/goc) | 436 | 60 | 24 | A Comprehensive Coverage Testing System for The Go Programming Language | 2020-05-07 03:46:25 | 2021-10-20 04:05:52 |
| [apitest](https://apitest.dev) | 404 | 37 | 0 | A simple and extensible behavioural testing library for Go. You can use api test to simplify REST API, HTTP handler and e2e tests. | 2018-12-26 22:27:19 | 2021-10-20 02:52:41 |
| [gofight](https://github.com/appleboy/gofight) | 396 | 39 | 6 | Testing API Handler written in Golang. | 2016-03-29 00:13:21 | 2021-10-18 14:55:00 |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 292 | 21 | 3 | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test | 2019-11-16 23:49:40 | 2021-10-18 06:46:54 |
| [frisby](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 269 | 26 | 13 | API testing framework inspired by frisby-js | 2015-09-15 14:35:58 | 2021-10-06 05:53:15 |
| [gotest.tools](https://gotest.tools) | 256 | 34 | 22 | A collection of packages to augment the go testing package and support common patterns. | 2017-08-08 21:28:54 | 2021-10-12 11:31:58 |
| [go-testdeep](https://go-testdeep.zetta.rocks/) | 228 | 9 | 2 | Extremely flexible golang deep comparison, extends the go testing package, tests HTTP APIs and provides tests suite | 2018-05-26 15:03:28 | 2021-10-17 03:14:54 |
| [go-carpet](https://github.com/msoap/go-carpet) | 215 | 7 | 2 | go-carpet - show test coverage in terminal for Go source files | 2016-02-28 12:02:51 | 2021-09-19 12:37:21 |
| [endly](https://github.com/viant/endly) | 198 | 23 | 0 | End to end functional test and automation framework | 2017-08-28 20:24:43 | 2021-10-20 20:49:23 |
| [charlatan](https://github.com/percolate/charlatan) | 194 | 8 | 2 | Go Interface Mocking Tool | 2017-10-06 21:55:14 | 2021-10-07 06:02:37 |
| [commander](https://github.com/commander-cli/commander) | 188 | 12 | 23 | Test your command line interfaces on windows, linux and osx and nodes viá ssh and docker | 2019-02-22 16:35:16 | 2021-10-17 06:34:17 |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 183 | 24 | 11 | Simple Go snapshot testing | 2017-08-07 18:30:05 | 2021-10-19 11:07:25 |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 132 | 11 | 1 | Clean database for testing, inspired by database_cleaner for Ruby | 2017-01-17 18:18:40 | 2021-09-27 22:19:56 |
| [gospec](https://github.com/luontola/gospec) | 114 | 16 | 3 | Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] | 2009-11-24 13:59:26 | 2021-02-11 17:23:11 |
| [wstest](https://github.com/posener/wstest) | 87 | 12 | 1 | go websocket client for unit testing of a websocket handler | 2017-03-31 21:06:18 | 2021-10-18 00:03:12 |
| [gocrest](https://github.com/corbym/gocrest) | 82 | 4 | 2 | GoCrest - Hamcrest-like matchers for Go | 2017-12-23 23:27:00 | 2021-10-13 13:57:59 |
| [testcase](https://github.com/adamluzsi/testcase) | 81 | 4 | 0 | testcase is an opinionated testing framework based on BDD principles. | 2019-04-22 21:20:51 | 2021-10-20 21:27:20 |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 71 | 11 | 4 | A Go test assertion library for verifying that two representations of JSON are semantically equal | 2018-10-26 20:31:01 | 2021-10-13 14:27:23 |
| [restit](https://github.com/go-restit/restit) | 54 | 4 | 4 | A Go library help testing your RESTful API application | 2014-06-25 10:25:46 | 2021-08-27 19:24:43 |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 6 | 1 | A BDD library for Go | 2009-11-20 06:34:29 | 2021-02-11 17:23:22 |
| [go-hit](https://github.com/Eun/go-hit) | 52 | 1 | 5 | http integration test framework | 2019-06-04 16:28:23 | 2021-10-15 02:18:06 |
| [covergates](https://covergates.com) | 45 | 8 | 11 | The portal gates to coverage reports | 2020-05-29 04:02:01 | 2021-09-26 20:00:23 |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 41 | 3 | 0 | Library created for testing JSON against patterns. | 2019-01-27 20:19:06 | 2021-07-20 10:08:34 |
| [dsunit](https://github.com/viant/dsunit) | 39 | 6 | 0 | Datastore Testibility | 2016-06-13 20:20:52 | 2021-10-17 19:42:26 |
| [assert](https://github.com/go-playground/assert) | 35 | 9 | 2 | :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions | 2015-07-20 17:53:45 | 2021-09-03 23:20:30 |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 27 | 4 | 2 | Hamcrest matchers for the Go programming language | 2010-12-22 04:49:44 | 2021-10-12 09:24:22 |
| [flute](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 16 | 0 | 3 | Golang HTTP client testing framework | 2019-07-06 04:32:03 | 2021-10-20 08:56:22 |
| [schema](https://github.com/jgroeneveld/schema) | 15 | 0 | 0 | Quick and easy expression matching for JSON schemas used in requests and responses | 2015-08-13 13:36:54 | 2021-03-01 05:46:30 |
| [testsql](https://github.com/zhulongcheng/testsql) | 12 | 1 | 3 | Generate test data from SQL files before testing and clear it after finished. | 2018-09-22 12:03:50 | 2021-10-12 09:25:43 |
| [gogiven](https://github.com/corbym/gogiven) | 11 | 2 | 4 | gogiven - BDD testing framework for go that generates readable output directly from source code | 2017-12-31 22:33:37 | 2021-10-12 09:23:57 |
| [biff](https://github.com/fulldump/biff) | 10 | 1 | 0 | Bifurcation Framework for testing and use cases | 2018-03-28 18:35:53 | 2021-07-18 09:37:31 |
| [badio](https://github.com/cavaliergopher/badio) | 10 | 1 | 0 | Extensions to Go's testing/iotest package | 2016-02-11 10:29:25 | 2021-09-30 02:21:09 |
| [gosuite](https://github.com/pavlo/gosuite) | 10 | 3 | 1 | Test suites support for standard Go1.7 "testing" by leveraging Subtests feature | 2016-10-15 19:28:14 | 2021-06-01 17:30:40 |
| [test](https://github.com/tvastar/test) | 6 | 0 | 0 | test utilities for golang | 2019-03-23 21:47:36 | 2021-08-17 10:46:53 |
| [stop-and-go](https://github.com/elgohr/stop-and-go) | 6 | 3 | 0 | Testing helper for concurrency | 2020-11-06 09:04:58 | 2021-08-28 20:30:48 |
| [trial](https://github.com/jgroeneveld/trial) | 5 | 0 | 0 | A simple assertion library for go | 2015-06-18 09:01:30 | 2021-04-18 16:15:16 |
| [tt](https://github.com/vcaesar/tt) | 5 | 0 | 0 | Simple and colorful test tools | 2018-04-03 11:47:21 | 2021-09-24 15:53:40 |
</details>

### Text Processing - Specific Formats


<sup>*Last Update: 2021-09-13 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [colly](http://go-colly.org/) | 14,790 | 1,265 | 127 | Elegant Scraper and Crawler Framework for Golang | 2017-09-29 14:08:49 | 2021-09-12 22:45:48 |
| [goquery](https://godoc.org/astuart.co/goq) | 10,583 | 795 | 4 | A little like that j-thing, only in Go. | 2012-08-29 02:14:59 | 2021-09-12 15:17:39 |
| [blackfriday](https://github.com/russross/blackfriday) | 4,785 | 580 | 199 | Blackfriday: a markdown processor for Go | 2011-05-27 22:28:58 | 2021-09-13 00:46:03 |
| [sh](https://pkg.go.dev/mvdan.cc/sh/v3) | 4,010 | 203 | 51 | A shell parser, formatter, and interpreter with bash support; includes shfmt | 2016-01-16 08:39:09 | 2021-09-12 17:18:41 |
| [toml](https://github.com/BurntSushi/toml) | 3,602 | 462 | 20 | TOML parser for Golang with reflection. | 2013-02-26 05:05:48 | 2021-09-12 07:07:33 |
| [go-humanize](https://pkg.go.dev/github.com/dustin/go-humanize) | 2,796 | 186 | 31 | Go Humans! (formatters for units to human friendly sizes) | 2012-01-13 03:48:55 | 2021-09-12 01:51:57 |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 2,053 | 129 | 16 | bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS | 2013-11-20 22:15:49 | 2021-09-11 03:37:55 |
| [gofeed](https://github.com/mmcdole/gofeed) | 1,708 | 158 | 42 | Parse RSS, Atom and JSON feeds in Go | 2016-01-23 02:44:34 | 2021-09-12 22:45:54 |
| [inject](https://godoc.org/github.com/facebookgo/inject) | 1,350 | 117 | 9 | Package inject provides a reflect based injector. | 2013-10-21 01:51:46 | 2021-09-09 09:13:30 |
| [go-toml](https://github.com/pelletier/go-toml) | 1,091 | 153 | 31 | Go library for the TOML file format | 2013-02-24 17:45:51 | 2021-09-09 20:54:18 |
| [commonregex](https://github.com/mingrammer/commonregex) | 783 | 57 | 2 | 🍫 A collection of common regular expressions for Go | 2017-03-23 14:33:18 | 2021-09-03 21:46:12 |
| [slug](https://pkg.go.dev/mvdan.cc/sh/v3) | 675 | 71 | 5 | URL-friendly slugify with multiple languages support. | 2014-03-31 06:24:51 | 2021-09-12 08:57:17 |
| [mxj](https://github.com/clbanning/mxj) | 482 | 86 | 0 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2014-02-03 13:39:16 | 2021-09-04 02:18:53 |
| [dataflowkit](https://dataflowkit.com) | 480 | 64 | 0 | Extract structured data from web sites. Web sites scraping.   | 2017-02-09 15:08:15 | 2021-09-04 06:48:17 |
| [gographviz](https://github.com/awalterschulze/gographviz) | 445 | 67 | 9 | Parses the Graphviz DOT language in golang | 2015-03-14 18:27:00 | 2021-09-08 10:28:20 |
| [gommon](https://github.com/labstack/gommon) | 408 | 91 | 16 | Common packages for Go | 2015-03-12 22:35:57 | 2021-09-08 12:50:51 |
| [htmlquery](https://github.com/antchfx/xpath) | 398 | 51 | 9 | htmlquery is golang XPath package for HTML query. | 2017-12-05 01:08:41 | 2021-09-12 19:19:49 |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 385 | 66 | 7 | wcwidth for golang | 2013-06-21 04:56:50 | 2021-09-10 02:03:17 |
| [omniparser](https://github.com/jf-tech/omniparser) | 376 | 19 | 0 | omniparser: a native Golang ETL streaming parser and transform library for CSV, JSON, XML, EDI, text, etc. | 2020-08-16 22:22:21 | 2021-09-13 01:06:42 |
| [gotext](https://github.com/zhshch2002/gospider) | 304 | 32 | 5 | Go (Golang) GNU gettext utilities package  | 2016-06-19 20:14:43 | 2021-09-05 14:19:16 |
| [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) | 251 | 39 | 6 | ⚙️ Convert HTML to Markdown. Even works with entire websites and can be extended through rules. | 2018-05-15 13:26:26 | 2021-09-10 14:49:05 |
| [goq](https://godoc.org/astuart.co/goq) | 200 | 18 | 2 | A declarative struct-tag-based HTML unmarshaling or scraping package for Go built on top of the goquery library | 2017-02-20 02:54:40 | 2021-09-12 04:56:43 |
| [goribot](https://github.com/zhshch2002/gospider) | 198 | 30 | 1 | [Crawler/Scraper for Golang]🕷A lightweight distributed friendly Golang crawler framework.一个轻量的分布式友好的 Golang 爬虫框架。 | 2019-09-08 10:39:47 | 2021-08-31 03:46:59 |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 154 | 54 | 3 | A NMEA parser library in pure Go | 2015-07-22 08:55:54 | 2021-09-10 09:19:43 |
| [github_flavored_markdown](https://github.com/shurcooL/github_flavored_markdown) | 132 | 32 | 9 | GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links. | 2015-05-16 04:09:07 | 2021-09-10 19:40:28 |
| [sdp](https://github.com/gortc/sdp) | 112 | 30 | 5 | RFC 4566 SDP implementation in go | 2016-05-13 14:35:11 | 2021-09-12 09:22:41 |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 96 | 8 | 0 | Zero-width character detection and removal for Go | 2018-06-18 13:55:09 | 2021-08-06 07:44:53 |
| [podcast](https://github.com/eduncan911/podcast) | 95 | 26 | 5 | iTunes and RSS 2.0 Podcast Generator in Golang | 2017-02-02 12:45:04 | 2021-08-28 19:10:01 |
| [editorconfig-core-go](https://godoc.org/github.com/hscells/doi) | 87 | 29 | 4 | EditorConfig Core written in Go | 2016-07-05 03:50:41 | 2021-09-12 17:22:15 |
| [align](https://github.com/Guitarbum722/align) | 70 | 6 | 0 | A general purpose application and library for aligning text. | 2017-04-29 23:22:22 | 2021-09-12 16:16:48 |
| [go-slugify](https://godoc.org/github.com/mozillazg/go-slugify) | 68 | 5 | 1 | Pretty Slug. | 2016-07-16 11:55:15 | 2021-08-10 20:53:12 |
| [genex](https://namegrep.com/) | 60 | 5 | 0 | Genex package for Go | 2015-03-09 19:24:16 | 2020-12-10 08:56:04 |
| [go-vcard](https://github.com/pelletier/go-toml) | 60 | 18 | 4 | A Go library to parse and format vCard | 2017-03-21 08:30:36 | 2021-08-16 15:20:58 |
| [goregen](https://godoc.org/github.com/zach-klippenstein/goregen) | 58 | 6 | 4 | randexp for Go. | 2014-12-27 00:19:39 | 2021-08-13 17:56:34 |
| [did](https://w3c-ccg.github.io/did-spec) | 55 | 13 | 4 | A golang package to work with Decentralized Identifiers (DIDs)  | 2018-11-02 17:49:14 | 2021-09-07 13:11:55 |
| [go-fixedwidth](http://godoc.org/github.com/ianlopshire/go-fixedwidth) | 54 | 21 | 9 | Encoding and decoding for fixed-width formatted data | 2017-11-15 21:05:44 | 2021-06-15 20:23:55 |
| [guesslanguage](https://github.com/zhshch2002/gospider) | 53 | 2 | 1 | Guess the natural language of a text in Go | 2014-12-16 10:58:47 | 2021-06-10 08:36:39 |
| [allot](https://github.com/sbstjn/allot) | 49 | 6 | 1 | Parse placeholder and wildcard text commands | 2016-10-16 15:49:08 | 2021-09-03 02:59:20 |
| [pagser](https://github.com/foolin/pagser) | 40 | 3 | 2 | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler | 2020-04-19 09:22:00 | 2021-09-08 08:41:08 |
| [gonameparts](https://github.com/polera/gonameparts) | 33 | 3 | 2 | Takes a full name and splits it into individual name parts | 2015-05-17 05:20:17 | 2020-11-25 21:46:02 |
| [slugify](https://pkg.go.dev/mvdan.cc/sh/v3) | 27 | 2 | 0 | A Go slugify application that handles string | 2015-04-13 01:54:30 | 2021-02-17 19:13:25 |
| [normalize](https://github.com/avito-tech/normalize) | 19 | 1 | 0 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2021-03-22 09:25:14 | 2021-06-27 23:46:57 |
| [codetree](https://github.com/aerogo/codetree) | 17 | 3 | 0 | :evergreen_tree: Parses indented code and returns a tree structure. | 2016-11-26 02:50:38 | 2021-02-06 17:49:49 |
| [enca](https://godoc.org/github.com/hscells/doi) | 11 | 3 | 0 | Minimal cgo bindings for libenca | 2014-12-17 04:55:16 | 2020-12-16 16:44:00 |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 2 | 0 | A syndication feed parser for Atom 1.0 and RSS 2.0 in Go | 2017-04-07 09:30:55 | 2019-10-28 23:12:30 |
| [doi](https://godoc.org/github.com/hscells/doi) | 6 | 0 | 0 | Parse and check doi objects in go. | 2017-08-02 05:58:01 | 2021-07-13 11:45:49 |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 6 | 1 | 0 | Converter from BBCode to HTML | 2016-04-15 14:35:38 | 2020-08-30 14:37:50 |
| [ltsv](https://github.com/Wing924/ltsv) | 5 | 0 | 0 | High performance LTSV (Labeled Tab Separeted Value) reader for Go. | 2019-05-12 06:11:04 | 2020-04-28 14:07:48 |
| [encoding](https://github.com/ake-persson/encoding) | 4 | 1 | 1 | Go package provides a generic interface to encoders and decoders | 2018-04-06 20:48:00 | 2021-07-06 03:56:31 |
| [go-wildcard](https://github.com/pelletier/go-toml) | 3 | 0 | 0 | A Go library to parse and format vCard | 2021-03-28 16:31:41 | 2021-08-30 12:05:34 |
| [go-output-format](https://github.com/drewstinnett/go-output-format) | 1 | 0 | 0 | Output go objects in standard formats, such as YAML, JSON, etc | 2021-04-08 20:48:17 | 2021-07-06 08:25:19 |
</details>

### Text Processing - Utility


<sup>*Last Update: 2023-10-02 20:50:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xurls](https://tysug.net) | 1,076 | 113 | 1 | Extract urls from text | 2015-01-12 01:28:46 | 2023-10-01 20:58:37 |
| [gotabulate](https://github.com/bndr/gotabulate) | 302 | 29 | 5 | Gotabulate - Easily pretty-print your tabular data with Go | 2014-08-21 07:44:28 | 2023-08-25 02:59:22 |
| [radix](https://github.com/yourbasic/radix) | 186 | 11 | 0 | A fast string sorting algorithm (MSD radix sort) | 2017-06-09 14:38:58 | 2023-08-23 11:19:55 |
| [regroup](https://github.com/oriser/regroup) | 135 | 12 | 2 | Match regex group into go struct using struct tags and automatic parsing | 2020-09-08 19:04:42 | 2023-10-02 07:15:05 |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 58 | 8 | 4 | A sanitization-based swear filter for Go. | 2018-09-09 00:07:26 | 2023-09-08 03:32:12 |
| [parth](https://github.com/codemodus/parth) | 43 | 6 | 0 | Path parsing for segment unmarshaling and slicing. | 2015-04-06 22:53:59 | 2023-01-08 16:11:52 |
| [xj2go](https://tysug.net) | 33 | 8 | 0 | Convert xml and json to go struct | 2017-09-19 13:20:57 | 2023-09-15 15:57:48 |
| [tagify](https://www.zoomio.org/tagify) | 33 | 2 | 3 | Tagify produces a set of tags from a given source. Source can be either an HTML page, a Markdown document or a plain text. Supports English, Russian, Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean languages. | 2018-03-20 10:30:11 | 2023-10-02 06:58:10 |
| [kace](https://github.com/codemodus/kace) | 20 | 3 | 1 | Common case conversions covering common initialisms. | 2015-06-04 20:36:49 | 2023-08-01 23:38:06 |
| [TySug](https://tysug.net) | 17 | 3 | 2 | A project around helping to prevent typing typos. TySug (Typo Suggestions) suggests alternative words with respect to keyboard layouts | 2018-06-05 19:46:29 | 2023-08-03 19:29:18 |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 9 | 5 | 1 | A string argument parser that understands quotes and backslashes | 2016-02-24 00:53:38 | 2023-06-15 15:58:40 |
| [textwrap](https://www.zoomio.org/tagify) | 5 | 4 | 1 | Port of Python's "textwrap" module to Go | 2019-07-26 17:57:55 | 2023-08-21 08:56:49 |
</details>

### Third-party APIs
Libraries for accessing third party APIs.

<sup>*Last Update: 2023-09-30 07:16:00*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-github](https://pkg.go.dev/github.com/google/go-github/v55/github) | 9,690 | 2,024 | 63 | Go library for accessing the GitHub v3 API | 2013-05-24 16:42:58 | 2023-09-29 13:24:14 |
| [aws-sdk-go](http://aws.amazon.com/sdk-for-go/) | 8,424 | 2,126 | 50 | AWS SDK for the Go programming language. | 2014-12-05 05:29:41 | 2023-09-29 10:34:01 |
| [slack](https://pkg.go.dev/github.com/slack-go/slack) | 4,411 | 1,125 | 103 | Slack API in Go - community-maintained fork created by the original author, @nlopes | 2015-01-24 14:19:00 | 2023-09-28 10:27:18 |
| [discordgo](https://github.com/bwmarrin/discordgo) | 4,262 | 883 | 131 |  (Golang) Go bindings for Discord | 2015-11-01 20:51:01 | 2023-09-29 11:31:33 |
| [google-api-go-client](https://pkg.go.dev/google.golang.org/api) | 3,578 | 1,127 | 26 | Auto-generated Google APIs for Go. | 2014-11-24 21:45:36 | 2023-09-29 13:31:13 |
| [google-cloud-go](https://cloud.google.com/go/docs/reference) | 3,368 | 1,237 | 256 | Google Cloud Client Libraries for Go. | 2014-05-09 11:11:58 | 2023-09-29 13:23:29 |
| [minio-go](https://docs.min.io/docs/golang-client-quickstart-guide.html) | 2,061 | 620 | 14 | MinIO Go client SDK for S3 compatible object storage | 2015-05-02 02:36:46 | 2023-09-29 15:06:54 |
| [stripe-go](https://stripe.com) | 1,861 | 473 | 9 | Go library for the Stripe API.     | 2014-06-05 23:38:14 | 2023-09-27 10:16:30 |
| [go-twitter](https://dev.twitter.com/rest/public) | 1,591 | 343 | 7 | Go Twitter REST and Streaming API v1.1 | 2015-04-11 23:26:07 | 2023-09-28 10:28:20 |
| [go-jira](https://pkg.go.dev/github.com/andygrunwald/go-jira?tab=doc) | 1,365 | 462 | 160 | Go client library for Atlassian Jira | 2015-08-20 15:02:46 | 2023-09-27 13:11:12 |
| [facebook](https://pkg.go.dev/github.com/huandu/facebook) | 1,198 | 574 | 1 | A Facebook Graph API SDK For Go. | 2012-07-28 19:05:56 | 2023-09-29 10:04:25 |
| [anaconda](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 1,137 | 252 | 73 | A Go client library for the Twitter 1.1 API | 2013-03-04 22:46:07 | 2023-09-10 07:24:54 |
| [githubv4](https://github.com/shurcooL/githubv4) | 1,034 | 89 | 41 | Package githubv4 is a client library for accessing GitHub GraphQL API v4 (https://docs.github.com/en/graphql). | 2017-05-27 05:05:31 | 2023-09-15 19:58:47 |
| [webhooks](https://github.com/go-playground/webhooks) | 868 | 226 | 31 | :fishing_pole_and_fish: Webhook receiver for GitHub, Bitbucket, GitLab, Gogs | 2015-10-25 17:38:13 | 2023-09-26 16:43:45 |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 709 | 129 | 16 | Scrape the Twitter frontend API without authentication with Golang. | 2018-11-29 15:31:50 | 2023-09-28 08:12:01 |
| [paypal](https://github.com/plutov/paypal) | 603 | 261 | 2 | Golang client for PayPal REST API | 2015-10-14 04:57:49 | 2023-09-26 18:50:11 |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 480 | 63 | 6 | Go library to access geocoding and reverse geocoding APIs | 2014-12-04 08:18:31 | 2023-09-29 09:17:15 |
| [ethrpc](https://github.com/onrik/ethrpc) | 257 | 104 | 11 | Golang client for ethereum json rpc api | 2017-01-24 09:47:00 | 2023-08-17 04:34:39 |
| [trello](https://github.com/adlio/trello) | 213 | 73 | 7 | Trello API wrapper for Go | 2016-09-24 04:36:10 | 2023-09-18 14:49:05 |
| [go-marathon](https://github.com/gambol99/go-marathon) | 200 | 129 | 27 | A GO API library for working with Marathon | 2015-02-11 13:25:26 | 2023-09-01 03:20:16 |
| [wit-go](https://github.com/wit-ai/wit-go) | 147 | 30 | 1 | Go client for wit.ai HTTP API | 2018-08-20 07:18:40 | 2023-09-23 19:01:55 |
| [go-trending](http://godoc.org/github.com/andygrunwald/go-trending) | 138 | 17 | 0 | Go library for accessing trending repositories and developers at Github. | 2015-07-04 08:06:48 | 2023-09-19 03:18:26 |
| [medium-sdk-go](https://medium.com) | 138 | 20 | 6 | A Golang SDK for Medium's OAuth2 API | 2015-09-26 23:45:46 | 2023-07-30 04:40:22 |
| [pushover](https://github.com/gregdel/pushover) | 137 | 12 | 1 | Go wrapper for the Pushover API | 2015-02-19 15:30:05 | 2023-09-26 14:27:38 |
| [gostorm](https://github.com/jsgilmore/gostorm) | 130 | 20 | 5 | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. | 2013-07-22 12:43:41 | 2023-02-13 11:11:46 |
| [gosip](https://go.spflow.com) | 119 | 28 | 8 | ⚡️ SharePoint SDK for Go | 2019-01-26 08:48:48 | 2023-09-19 13:40:59 |
| [simples3](https://github.com/rhnvrm/simples3) | 116 | 25 | 2 | Simple no frills AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK) | 2018-12-06 10:24:21 | 2023-09-28 10:56:21 |
| [hipchat](https://github.com/daneharrigan/hipchat) | 110 | 37 | 1 | A golang package to communicate with HipChat over XMPP | 2013-04-28 02:16:21 | 2023-07-09 08:35:31 |
| [hipchat](https://github.com/andybons/hipchat) | 105 | 21 | 0 | This project implements a Go client library for the Hipchat API. | 2012-10-20 18:34:06 | 2023-09-25 21:28:16 |
| [cachet](https://github.com/andygrunwald/cachet) | 91 | 13 | 1 | Go(lang) client library for Cachet (open source status page system). | 2015-10-31 12:30:07 | 2022-09-27 09:40:01 |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 85 | 21 | 1 | This is a Golang wrapper for working with TMDb API. It aims to support version 3. | 2019-01-11 22:59:33 | 2023-09-05 01:44:18 |
| [igdb](https://api.igdb.com/) | 79 | 16 | 3 | Go client for the Internet Game Database API | 2017-08-24 08:31:53 | 2023-09-25 15:17:30 |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 73 | 18 | 1 | Golang scraper to get data from Google Play Store | 2019-09-20 14:03:01 | 2023-09-18 14:43:47 |
| [gogtrends](https://github.com/groovili/gogtrends) | 73 | 28 | 3 | Unofficial Google Trends API for Go | 2018-12-27 13:50:34 | 2023-09-18 14:43:17 |
| [go-unsplash](https://unsplash.com) | 72 | 14 | 8 | Go Client for the Unsplash API  | 2017-01-19 07:04:04 | 2023-06-06 05:29:44 |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 70 | 18 | 1 | Go module to work with Postman Collections | 2019-11-16 12:13:32 | 2023-09-07 17:58:35 |
| [airtable](https://github.com/mehanizm/airtable) | 64 | 16 | 1 | Simple golang airtable API wrapper | 2020-04-12 10:05:07 | 2023-09-08 18:05:49 |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 64 | 50 | 5 | Go library for interacting with CircleCI | 2015-08-14 21:19:36 | 2023-02-25 15:44:07 |
| [mixpanel](https://github.com/dukex/mixpanel) | 60 | 32 | 6 | Golang Mixpanel Client | 2014-05-20 03:50:34 | 2023-06-22 07:09:11 |
| [ynab.go](https://godoc.org/github.com/brunomvsouza/ynab.go) | 60 | 22 | 3 | Go client for the YNAB API. Unofficial. It covers 100% of the resources made available by the YNAB API. | 2018-07-13 11:10:54 | 2023-09-15 21:06:23 |
| [go-amazon-product-advertising-api](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 57 | 15 | 3 | Go Client Library for Amazon Product Advertising API | 2016-11-15 15:37:32 | 2023-08-01 21:59:10 |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 57 | 11 | 13 | Client library for UptimeRobot v2 API | 2018-05-29 10:27:19 | 2023-09-18 14:47:47 |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 55 | 15 | 8 | DEPRECATED: please use https://github.com/Clarifai/clarifai-go-grpc | 2015-09-28 23:33:59 | 2023-08-17 19:56:10 |
| [megos](https://godoc.org/github.com/andygrunwald/megos) | 54 | 11 | 0 | Go(lang) client library for accessing information of an Apache Mesos cluster. | 2015-10-02 14:29:20 | 2023-06-19 00:39:00 |
| [fcm](https://pkg.go.dev/github.com/huandu/facebook) | 51 | 15 | 2 | Firebase Cloud Messaging for application servers implemented using the Go programming language. | 2017-01-06 08:30:57 | 2023-08-27 19:26:42 |
| [go-xkcd](https://pkg.go.dev/github.com/nishanths/go-xkcd/v2) | 51 | 5 | 0 | xkcd.com API client in Go | 2016-02-26 05:14:31 | 2022-12-24 04:57:10 |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 51 | 24 | 6 | a Go (Golang) MusicBrainz WS2 client library - work in progress | 2014-09-10 16:42:33 | 2023-08-19 09:56:35 |
| [gads](https://pkg.go.dev/github.com/huandu/facebook) | 50 | 58 | 8 | Google Adwords API for Go | 2014-01-20 02:22:15 | 2022-09-27 09:40:27 |
| [go-spotify](https://pkg.go.dev/github.com/slack-go/slack) | 48 | 8 | 1 | Go library for the Spotify Web API | 2014-10-30 02:52:04 | 2023-09-26 20:47:37 |
| [golyrics](https://github.com/mamal72/golyrics) | 40 | 2 | 0 | A simple Go package to fetch lyrics from Wikia | 2016-11-18 04:40:37 | 2022-12-24 05:00:10 |
| [patreon-go](https://github.com/mxpv/patreon-go) | 37 | 19 | 1 | Patreon Go API client | 2017-08-06 21:15:14 | 2023-09-18 14:44:21 |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 36 | 2 | 2 | Go library for accessing the MyAnimeList API: https://myanimelist.net/apiconfig/references/api/v2 | 2015-05-03 10:07:05 | 2023-03-30 10:42:12 |
| [translate](https://github.com/nuveo/translate) | 33 | 6 | 0 | Go online translation package | 2015-07-13 15:42:13 | 2023-01-28 07:01:11 |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 32 | 6 | 3 | Golang client for LastPass | 2019-07-11 14:26:39 | 2023-08-18 23:39:40 |
| [gami](https://gitlab.com/bit4bit/gami) | 32 | 30 | 1 | GO - Asterisk AMI Interface | 2014-05-14 16:11:37 | 2023-08-10 19:51:18 |
| [gcm](https://gitlab.com/bit4bit/gami) | 31 | 4 | 0 | Google Cloud Messaging for application servers implemented using the Go programming language. | 2015-11-09 16:16:25 | 2022-09-27 09:40:32 |
| [go-steam](https://pkg.go.dev/github.com/slack-go/slack) | 30 | 6 | 2 | Go library for querying Source servers | 2014-11-23 16:34:56 | 2023-05-30 02:15:11 |
| [go-imgur](https://github.com/koffeinsource/go-imgur) | 25 | 9 | 3 | Go library to use the imgur.com API | 2016-03-30 22:05:35 | 2023-05-28 07:32:55 |
| [go-shopify](https://github.com/rapito/go-shopify) | 25 | 6 | 2 | Simple Shopify API for the Go Programming Language | 2014-10-28 02:53:25 | 2022-12-07 01:15:06 |
| [device-check-go](https://github.com/rinchsan/device-check-go) | 22 | 6 | 0 | :iphone: iOS DeviceCheck SDK for Go - query and modify the per-device bits | 2019-04-11 13:09:11 | 2023-09-18 14:38:32 |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 21 | 2 | 5 | A golang client for the Twitch v3 API - public APIs only (for now) | 2016-06-28 20:54:34 | 2021-05-10 11:16:29 |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 20 | 9 | 2 | Go client library for interacting with Coinpaprika's API | 2018-09-25 07:34:50 | 2023-02-02 19:34:12 |
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
| [libgoffi](https://github.com/noctarius/libgoffi) | 9 | 1 | 0 | libgoffi - libffi adapter library for Go | 2019-08-03 17:05:34 | 2022-10-08 19:34:36 |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 9 | 0 | 0 | Go bindings for RRDA https://github.com/fcambus/rrda | 2014-09-15 21:06:16 | 2023-01-28 08:52:34 |
| [gumblr](https://github.com/mattcunningham/gumblr) | 9 | 7 | 0 | A Go Wrapper for the Tumblr v2 API | 2015-07-09 23:13:51 | 2023-08-02 19:07:05 |
| [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | 9 | 2 | 0 | This is RAWG SDK GO. This library contains methods for interacting with RAWG API. | 2020-10-16 15:31:37 | 2023-09-19 14:08:59 |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 8 | 1 | 0 | Go client library for the SPTrans Olho Vivo API. :bus: | 2017-09-11 01:21:28 | 2022-12-03 19:02:18 |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 8 | 2 | 0 | :dancers: Go Chronos 3.x REST API Client | 2017-10-23 12:19:01 | 2023-05-24 01:27:43 |
| [go-google-email-audit-api](https://godoc.org/github.com/ngs/go-google-email-audit-api/emailaudit) | 8 | 5 | 0 | Go Client Library for G Suite Email Audit API | 2016-10-24 02:34:29 | 2023-07-18 07:57:25 |
| [go-zooz](https://github.com/gojuno/go-zooz) | 7 | 7 | 1 | Zooz API client for Go | 2017-07-04 09:28:23 | 2022-09-27 09:45:40 |
| [kanka](https://kanka.io/en-US/docs/1.0) | 3 | 4 | 2 | Go client for the Kanka API | 2019-12-26 00:07:57 | 2021-12-10 23:36:57 |
| [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) | 3 | 2 | 0 | Golang SDK for AppStore Connect API (Unofficial) | 2020-06-11 10:05:56 | 2022-09-27 09:39:52 |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 2 | 1 | 0 | A TripAdvisor API wrapper for Golang. | 2019-04-15 18:12:11 | 2022-09-27 09:45:21 |
| [vl-go](https://github.com/verifid/vl-go) | 2 | 1 | 0 | Go client library around the VerifID identity verification layer API. | 2019-02-09 12:46:53 | 2022-09-27 09:45:31 |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 2 | 1 | 0 | This is the official Playlyfe Golang Sdk | 2015-05-25 09:34:47 | 2022-09-27 09:44:46 |
</details>

### UUID
Libraries for working with UUIDs.

<sup>*Last Update: 2023-10-02 20:50:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [uuid](https://github.com/google/uuid) | 4,622 | 381 | 28 | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. | 2016-02-12 22:17:59 | 2023-10-02 07:45:59 |
| [ulid](https://github.com/oklog/ulid) | 3,731 | 145 | 3 | Universally Unique Lexicographically Sortable Identifier (ULID) in Go | 2016-12-06 15:26:52 | 2023-10-02 12:44:36 |
| [uuid](https://gof.rs/projects/uuid/) | 1,422 | 103 | 7 | A UUID package originally forked from github.com/satori/go.uuid | 2018-07-13 02:13:28 | 2023-10-02 05:38:00 |
| [wuid](https://github.com/edwingeng/wuid) | 501 | 48 | 1 | An extremely fast globally unique number generator. | 2018-01-27 01:16:28 | 2023-08-27 12:15:00 |
| [sno](https://pkg.go.dev/github.com/muyo/sno?tab=doc) | 87 | 5 | 1 | Compact, sortable and fast unique IDs with embedded metadata. | 2019-05-26 22:05:26 | 2023-09-08 05:25:42 |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 58 | 7 | 0 | A tiny and fast Go unique string generator | 2019-07-02 12:15:56 | 2023-07-10 12:46:11 |
| [Goid](https://github.com/jakenvac/Goid) | 38 | 4 | 1 | A UUIDv4 generation package written in go | 2017-05-19 10:40:45 | 2023-01-07 04:19:58 |
| [gouid](https://github.com/twharmon/gouid) | 21 | 4 | 0 | Fast, dependable universally unique ids | 2020-10-08 19:54:41 | 2023-07-15 15:39:02 |
| [uuid](https://github.com/agext/uuid) | 16 | 5 | 0 | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. | 2016-02-03 03:02:51 | 2023-08-11 14:41:54 |
| [GoFlake](https://github.com/Hart87/GoFlake) | 7 | 0 | 0 | A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake. | 2021-05-03 14:44:19 | 2021-06-30 15:29:33 |
</details>

### Utilities
General utilities and tools to make your life easier.

<sup>*Last Update: 2021-10-07 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fzf](https://github.com/junegunn/fzf) | 39,400 | 1,705 | 281 | :cherry_blossom: A command-line fuzzy finder | 2013-10-23 16:04:23 | 2021-09-29 01:55:50 |
| [hub](https://hub.github.com/) | 21,264 | 2,249 | 271 | A command-line tool that makes git easier to use with GitHub. | 2009-12-05 22:15:25 | 2021-10-05 01:59:08 |
| [ctop](https://ctop.sh) | 11,964 | 471 | 66 | Top-like interface for container metrics | 2016-12-27 02:25:57 | 2021-09-28 19:25:40 |
| [sqlx](http://jmoiron.github.io/sqlx/) | 10,885 | 851 | 279 | general purpose extensions to golang's database/sql | 2013-01-28 19:40:00 | 2021-10-07 00:33:18 |
| [wuzz](https://github.com/asciimoo/wuzz) | 9,782 | 388 | 37 | Interactive cli tool for HTTP inspection | 2017-01-30 21:22:00 | 2021-10-05 14:13:11 |
| [goreleaser](https://goreleaser.com) | 8,777 | 585 | 28 | Deliver Go binaries as fast and easily as possible | 2016-12-21 17:13:39 | 2021-10-04 19:25:26 |
| [usql](https://github.com/xo/usql) | 6,785 | 244 | 55 | Universal command-line interface for SQL databases | 2017-03-02 13:03:21 | 2021-10-04 20:10:22 |
| [peco](https://github.com/peco/peco) | 6,529 | 216 | 41 | Simplistic interactive filtering tool | 2014-06-06 06:06:32 | 2021-10-05 01:59:22 |
| [godropbox](https://github.com/dropbox/godropbox) | 3,972 | 420 | 5 | Common libraries for writing Go services/applications. | 2014-06-22 23:09:29 | 2021-09-27 02:32:46 |
| [hystrix-go](https://github.com/afex/hystrix-go) | 3,391 | 391 | 54 | Netflix's Hystrix latency and fault tolerance library, for Go  | 2013-12-15 08:51:23 | 2021-10-03 00:30:44 |
| [goreporter](https://github.com/qax-os/goreporter) | 2,916 | 258 | 29 | A Golang tool that does static analysis, unit testing, code review and generate code quality report. | 2017-03-27 08:46:38 | 2021-09-28 09:28:52 |
| [go-funk](https://github.com/thoas/go-funk) | 2,908 | 169 | 5 | A modern Go utility library which provides helpers (map, find, contains, filter, ...) | 2016-12-30 13:55:15 | 2021-09-28 19:21:54 |
| [minify](https://go.tacodewolff.nl/minify) | 2,769 | 169 | 8 | Go minifiers for web formats | 2014-05-21 09:03:48 | 2021-10-04 19:59:36 |
| [panicparse](https://maruel.ca) | 2,732 | 81 | 5 | Crash your app in style (Golang) | 2015-02-02 02:14:41 | 2021-10-03 20:34:43 |
| [mc](https://min.io/download) | 1,914 | 345 | 36 | MinIO Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. | 2015-01-16 02:56:51 | 2021-10-04 19:19:52 |
| [storm](https://github.com/asdine/storm) | 1,788 | 127 | 65 | Simple and powerful toolkit for BoltDB | 2016-01-10 12:55:59 | 2021-10-03 21:18:12 |
| [mergo](https://github.com/imdario/mergo) | 1,721 | 200 | 19 | Mergo: merging Go structs and maps since 2013. | 2013-03-11 22:51:11 | 2021-10-03 10:05:40 |
| [spinner](https://github.com/briandowns/spinner) | 1,584 | 98 | 14 | Go (golang) package with 80 configurable terminal spinner/progress indicators. | 2014-12-13 00:36:19 | 2021-10-05 22:42:12 |
| [mole](https://davrodpin.github.io/mole/) | 1,503 | 82 | 14 | CLI application to create ssh tunnels focused on resiliency and user experience. | 2018-10-04 02:38:00 | 2021-10-03 01:28:41 |
| [filetype](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 1,409 | 128 | 26 | Fast, dependency-free Go package to infer binary file types based on the magic numbers header signature | 2015-09-24 09:15:51 | 2021-09-27 10:41:28 |
| [boilr](https://github.com/tmrts/boilr) | 1,390 | 107 | 44 | :zap: boilerplate template manager that generates files or directories from template repositories | 2015-12-19 16:57:26 | 2021-09-28 13:37:08 |
| [jump](http://gsamokovarov.com/jump) | 1,246 | 48 | 1 | Jump helps you navigate faster by learning your habits. ✌️ | 2015-08-16 22:07:17 | 2021-10-01 11:39:01 |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 961 | 106 | 18 | Circuit Breakers in Go | 2014-07-17 22:41:33 | 2021-09-26 08:11:08 |
| [gtm](https://github.com/git-time-metric/gtm) | 880 | 48 | 47 | Simple, seamless, lightweight time tracking for Git | 2016-06-19 21:17:04 | 2021-09-28 09:21:17 |
| [cli](https://create-go.app/wiki) | 856 | 91 | 1 | ✨ Create a new production-ready project with backend, frontend and deploy automation by running one CLI command! | 2019-12-30 22:08:38 | 2021-09-27 21:47:54 |
| [immortal](https://immortal.run) | 718 | 50 | 1 | ⭕  A *nix cross-platform (OS agnostic) supervisor | 2016-06-30 17:02:27 | 2021-09-08 18:12:56 |
| [hostctl](http://guumaster.github.io/hostctl) | 679 | 30 | 9 | Your dev tool to manage /etc/hosts like a pro! | 2020-03-14 11:29:02 | 2021-09-26 14:46:57 |
| [circuit](https://github.com/cep21/circuit) | 596 | 35 | 7 | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. | 2017-12-23 22:17:43 | 2021-09-28 05:40:29 |
| [mimetype](https://pkg.go.dev/github.com/gabriel-vasile/mimetype#pkg-overview) | 588 | 85 | 39 | A fast Golang library for media type and file extension detection, based on magic numbers | 2018-07-02 07:15:29 | 2021-09-29 11:07:18 |
| [htcat](https://github.com/htcat/htcat) | 536 | 29 | 5 | Parallel and Pipelined HTTP GET Utility | 2013-08-05 11:17:01 | 2021-10-03 23:58:32 |
| [godaemon](https://github.com/VividCortex/godaemon) | 480 | 54 | 8 | Daemonize Go applications deviously. | 2013-08-01 17:16:30 | 2021-09-17 01:36:19 |
| [ergo](https://github.com/cristianoliveira/ergo) | 480 | 47 | 16 | The management of multiple apps running over different ports made easy | 2017-08-19 18:41:56 | 2021-09-28 07:44:01 |
| [koazee](https://github.com/wesovilabs/koazee) | 473 | 25 | 15 | A StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices. | 2018-11-09 09:49:19 | 2021-09-27 19:55:38 |
| [go-dry](https://github.com/ungerik/go-dry) | 468 | 32 | 0 | DRY (don't repeat yourself) package for Go | 2014-02-28 13:49:31 | 2021-09-10 11:27:43 |
| [gopencils](https://github.com/bndr/gopencils) | 436 | 42 | 7 | Easily consume REST APIs with Go (golang) | 2014-06-23 11:41:24 | 2021-07-19 22:58:45 |
| [scany](https://github.com/georgysavva/scany) | 402 | 26 | 18 | Library for scanning data from a database into Go structs and more | 2020-07-02 11:02:58 | 2021-10-07 00:00:32 |
| [request](https://godoc.org/github.com/mozillazg/request) | 400 | 39 | 6 | A developer-friendly HTTP request library for Gopher. | 2014-12-21 04:30:42 | 2021-09-27 09:09:49 |
| [delve](https://github.com/derekparker/delve) | 390 | 82 | 0 | Delve is a debugger for the Go programming language. | 2020-02-18 18:03:33 | 2021-09-28 05:09:57 |
| [gubrak](https://pkg.go.dev/github.com/novalagung/gubrak) | 369 | 31 | 0 | ⚙️ Golang functional utility library with syntactic sugar. It's like lodash, but for Go | 2018-03-09 11:28:05 | 2021-09-22 04:33:20 |
| [deepcopier](https://github.com/ulule/deepcopier) | 361 | 50 | 7 | simple struct copying for golang | 2015-07-24 18:01:01 | 2021-09-29 00:16:25 |
| [clockwork](https://github.com/jonboulle/clockwork) | 350 | 43 | 6 | a fake clock for golang | 2014-09-09 18:24:00 | 2021-09-22 15:40:15 |
| [go-rate](https://github.com/beefsack/go-rate) | 340 | 26 | 0 | A timed rate limiter for Go | 2014-08-25 04:42:34 | 2021-08-23 16:52:38 |
| [retry](https://pkg.go.dev/github.com/kamilsk/retry/v5) | 309 | 13 | 9 | ♻️ The most advanced interruptible mechanism to perform actions repetitively until successful. | 2016-11-02 20:20:43 | 2021-09-23 03:52:44 |
| [gohper](https://github.com/cosiner/gohper) | 254 | 48 | 0 | [UNMATAINED] common libs here. | 2015-03-23 22:46:12 | 2021-09-14 16:19:51 |
| [serve](https://syntaqx-serve.herokuapp.com/) | 249 | 15 | 4 | 🍽️ a static http server anywhere you need one. | 2019-01-10 23:31:52 | 2021-09-28 11:31:44 |
| [scan](https://github.com/blockloop/scan) | 230 | 13 | 0 | Scan database/sql rows directly to structs, slices, and primitive types | 2017-11-27 23:22:18 | 2021-10-06 17:06:31 |
| [go-trigger](https://github.com/sadlil/go-trigger) | 223 | 38 | 1 | A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. | 2015-10-19 09:42:17 | 2021-09-28 08:03:29 |
| [util](https://github.com/shomali11/util) | 217 | 34 | 0 | A collection of useful utility functions | 2017-05-24 00:21:29 | 2021-09-09 12:01:01 |
| [gotenv](https://github.com/subosito/gotenv) | 205 | 22 | 5 | Load environment variables from `.env` or `io.Reader` in Go. | 2013-08-27 12:56:47 | 2021-09-14 13:06:08 |
| [death](http://vrecan.github.io/post/golang_shutdown/) | 174 | 17 | 0 | Managing go application shutdown with signals. | 2015-03-09 03:50:40 | 2021-09-21 13:39:24 |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 171 | 8 | 0 | go-bind-plugin generates API for exported plugin symbols (-buildmode=plugin) - go1.8+ only (http://golang.org/pkg/plugin) | 2016-11-08 14:40:26 | 2021-09-14 13:01:30 |
| [toolbox](https://github.com/viant/toolbox) | 166 | 19 | 2 | Toolbox - go utility library | 2016-06-13 19:33:35 | 2021-09-23 04:34:29 |
| [rerun](https://github.com/ivpusic/rerun) | 162 | 9 | 2 | Configurable recompiling and rerunning go apps when source changes | 2014-12-10 00:29:54 | 2021-05-21 19:48:51 |
| [moldova](https://github.com/StabbyCutyou/moldova) | 161 | 5 | 0 | A lightweight templating system for generating random data | 2016-01-30 05:25:39 | 2021-06-28 17:28:48 |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 161 | 48 | 22 | go-sitemap-generator is the easiest way to generate Sitemaps in Go | 2015-10-12 16:23:13 | 2021-09-23 03:13:13 |
| [apm](https://github.com/topfreegames/apm) | 155 | 73 | 9 | APM is a process manager for Golang applications. | 2015-11-18 16:56:48 | 2021-09-28 04:48:40 |
| [robustly](https://github.com/VividCortex/robustly) | 150 | 6 | 1 | Run functions resiliently in Go, catching and restarting panics | 2013-07-08 13:27:10 | 2021-07-06 08:27:20 |
| [chyle](https://github.com/antham/chyle) | 139 | 10 | 0 | Changelog generator : use a git repository and various data sources and publish the result on external services | 2016-11-17 21:14:44 | 2021-09-22 06:53:47 |
| [onecache](https://github.com/adelowo/onecache) | 124 | 6 | 0 | One caching API, Multiple backends | 2017-04-14 21:49:15 | 2021-09-22 03:27:43 |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 123 | 14 | 0 | Pure Go bsdiff and bspatch libraries and CLI tools. | 2019-02-23 23:33:50 | 2021-08-26 12:59:07 |
| [changie](https://changie.dev) | 123 | 8 | 5 | Automated changelog tool for preparing releases with lots of customization options | 2020-12-05 19:38:33 | 2021-09-23 19:12:05 |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 119 | 11 | 0 | LiveReload server for Go [golang] | 2014-07-15 05:36:53 | 2021-06-26 23:15:07 |
| [countries](https://godoc.org/github.com/biter777/countries) | 118 | 26 | 1 | Countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO init() funcs, NO external links/files/data, NO interface{}, NO specific dependencies, Databases/JSON/GOB/XML/CSV compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts. | 2019-04-22 14:47:11 | 2021-09-21 15:49:38 |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 92 | 6 | 0 | Pattern matchings for Go. | 2018-12-11 20:11:17 | 2021-09-25 09:05:47 |
| [nostromo](https://nostromo.sh) | 89 | 4 | 11 | CLI for building powerful aliases | 2019-07-13 04:51:46 | 2021-10-03 23:01:22 |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 89 | 10 | 0 | Database client library, proxy for any master slave, master master structures. Lightweight, performant and auto balancing in mind. | 2016-12-26 04:05:09 | 2021-09-17 07:32:30 |
| [goseaweedfs](https://github.com/chrislusf/seaweedfs) | 88 | 26 | 1 | A complete Golang client for SeaweedFS | 2017-07-20 04:35:39 | 2021-10-01 06:37:32 |
| [xferspdy](https://github.com/monmohan/xferspdy) | 88 | 9 | 3 | Xferspdy provides binary diff and patch library in golang. [Mentioned in Awesome Go, https://github.com/avelino/awesome-go] | 2015-05-22 13:23:34 | 2021-09-17 06:01:06 |
| [sorty](https://github.com/jfcg/sorty) | 86 | 1 | 1 | Fast Concurrent / Parallel Sorting in Go | 2019-02-18 21:05:45 | 2021-10-05 21:07:16 |
| [go-health](https://github.com/Talento90/go-health) | 83 | 4 | 1 | :heart: Health check your applications and dependencies | 2018-02-13 18:40:54 | 2021-09-14 13:06:41 |
| [pm](https://github.com/VividCortex/pm) | 78 | 6 | 2 | Processlist manager with TCP listener | 2013-11-17 19:17:01 | 2021-04-05 15:04:27 |
| [repeat](https://github.com/ssgreg/repeat) | 77 | 4 | 0 | Go implementation of different backoff strategies useful for retrying operations and heartbeating. | 2017-11-22 07:06:47 | 2021-07-15 21:15:10 |
| [mongo-go-pagination](https://mongo-go-pagination.herokuapp.com/normal-pagination?page=1&limit=15) | 74 | 27 | 0 | Golang Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines with all information like Total records, Page, Per Page, Previous, Next, Total Page and query results. | 2020-02-04 08:23:33 | 2021-09-23 14:10:34 |
| [netbug](https://github.com/e-dard/netbug) | 69 | 4 | 0 | Package netbug provides a handler for registering profilers on your own ServeMux. | 2015-03-05 19:27:29 | 2021-05-10 07:30:05 |
| [unis](https://godoc.org/github.com/esemplastic/unis) | 67 | 3 | 2 | UNIS: A Common Architecture for String Utilities within the Go Programming Language. | 2017-05-06 05:01:03 | 2021-03-10 09:39:40 |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 67 | 6 | 1 | Powerful and versatile MIME sniffing package using pre-compiled glob patterns, magic number signatures, XML document namespaces, and tree magic for mounted volumes, generated from the XDG shared-mime-info database. | 2018-10-11 16:12:54 | 2021-10-03 04:32:05 |
| [multitick](https://github.com/VividCortex/multitick) | 66 | 1 | 1 | A multiplexor for aligned time.Time tickers in Go | 2013-12-10 16:47:26 | 2021-09-05 12:06:30 |
| [handy](https://pkg.go.dev/github.com/novalagung/gubrak) | 66 | 6 | 0 | GO Golang Utilities and helpers like validators and string formatters | 2018-06-13 13:10:07 | 2021-09-02 21:39:50 |
| [cmd](https://github.com/commander-cli/cmd) | 65 | 6 | 2 | A simple package to execute shell commands on linux, windows and osx | 2019-09-27 13:22:06 | 2021-09-14 16:26:45 |
| [goreadability](https://github.com/philipjkim/goreadability) | 60 | 7 | 2 | Webpage summary extractor using Facebook Open Graph and arc90's readability | 2016-04-20 01:40:14 | 2021-09-29 16:05:32 |
| [minquery](https://github.com/icza/minquery) | 59 | 19 | 4 | MongoDB / mgo query that supports efficient pagination (cursors to continue listing documents where we left off). | 2016-11-16 12:23:07 | 2021-07-30 06:49:35 |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 59 | 8 | 2 | Parse TODOs in your GO code | 2016-10-17 19:51:36 | 2021-09-21 16:51:07 |
| [goval](https://github.com/maja42/goval) | 57 | 5 | 0 | Expression evaluation in golang | 2018-06-17 15:43:44 | 2021-10-03 21:49:09 |
| [golog](https://github.com/mlimaloureiro/golog) | 55 | 12 | 15 | Easy and simple CLI time tracker for your tasks | 2016-01-09 15:43:47 | 2021-06-13 16:19:23 |
| [pgo](https://github.com/arthurkushman/pgo) | 55 | 12 | 2 | Go library for PHP community with convenient functions | 2018-12-26 06:59:47 | 2021-09-28 09:40:52 |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 49 | 9 | 10 | Universal copy paste service, works across different machines! | 2017-01-28 15:35:24 | 2021-04-24 11:00:07 |
| [go-lock](https://github.com/viney-shih/go-lock) | 49 | 4 | 0 | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation | 2020-04-30 11:40:21 | 2021-08-13 01:49:32 |
| [retry](https://pkg.go.dev/github.com/thedevsaddam/retry) | 49 | 4 | 0 | Simple and easy retry mechanism package for Go | 2018-02-25 19:08:03 | 2021-08-30 07:30:01 |
| [filter](https://godoc.org/github.com/gookit/filter) | 48 | 5 | 1 | ⏳ Provide filtering, sanitizing, and conversion of Golang data. 提供对Golang数据的过滤，净化，转换。 | 2018-09-26 09:11:13 | 2021-07-30 07:47:31 |
| [beyond](http://wesovilabs.github.io/beyond) | 47 | 9 | 8 | The Go library that will drive you to AOP world! | 2019-10-18 05:41:45 | 2021-08-16 02:46:25 |
| [golarm](https://github.com/msempere/golarm) | 45 | 8 | 0 | Fire alarms with system events | 2015-08-14 16:51:53 | 2021-06-13 19:53:42 |
| [limiters](https://godoc.org/github.com/mennanov/limiters) | 45 | 7 | 1 | Golang rate limiters for distributed applications | 2019-08-28 18:09:54 | 2021-09-20 11:14:56 |
| [goback](https://github.com/carlescere/goback) | 44 | 7 | 6 | Golang simple exponential backoff package. | 2015-03-13 16:09:18 | 2021-03-04 10:31:30 |
| [slice](https://github.com/psampaz/slice) | 44 | 4 | 0 | Type-safe functions for common Go slice operations | 2019-11-26 05:20:39 | 2021-09-17 01:06:33 |
| [dbt](https://github.com/nikogura/dbt) | 44 | 4 | 6 | Dynamic Binary Toolkit- A framework for running self-updating signed binaries from a central, trusted repository. | 2017-11-30 22:53:17 | 2021-09-24 01:34:15 |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 42 | 3 | 2 | Retrying made simple and easy for golang :repeat:  | 2017-06-09 16:07:37 | 2021-09-30 02:21:54 |
| [intrinsic](https://immortal.run) | 42 | 1 | 1 | Provide Golang native SIMD intrinsics on x86/amd64 platform | 2017-06-13 09:26:34 | 2020-10-15 07:47:46 |
| [gpath](https://github.com/tenntenn/gpath) | 41 | 2 | 0 | gpath is a Go package to access a field by a path using reflect pacakge | 2017-05-24 06:24:18 | 2021-10-04 23:54:18 |
| [go-httpheader](https://godoc.org/github.com/mozillazg/go-httpheader) | 36 | 7 | 1 | A Go library for encoding structs into Header fields. | 2017-06-24 11:28:06 | 2021-09-16 18:14:39 |
| [myhttp](https://github.com/inancgumus/myhttp) | 35 | 13 | 1 | Simplest HTTP GET requester for Go with timeout support | 2017-09-13 15:48:47 | 2021-10-01 22:05:49 |
| [evaluator](https://github.com/nullne/evaluator) | 33 | 8 | 0 | The management of multiple apps running over different ports made easy | 2017-04-27 18:31:46 | 2021-09-13 09:29:34 |
| [equalizer](https://pkg.go.dev/github.com/reugn/equalizer) | 33 | 0 | 0 | A rate limiters package for Go | 2019-06-14 09:25:13 | 2021-09-22 13:38:22 |
| [gostrutils](https://github.com/chrislusf/seaweedfs) | 32 | 5 | 1 | Collections of string utils I have created over the years | 2018-09-19 11:06:11 | 2021-09-11 08:16:44 |
| [rclient](https://github.com/zpatrick/rclient) | 32 | 2 | 2 | Minimalistic REST client for Go applications | 2017-02-28 01:07:25 | 2020-10-13 11:13:06 |
| [backscanner](https://github.com/icza/backscanner) | 30 | 7 | 0 | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. | 2017-10-18 07:59:07 | 2021-09-06 23:55:00 |
| [slicer](https://github.com/leaanthony/slicer) | 28 | 1 | 0 | Utility class for handling slices | 2019-01-10 09:55:25 | 2021-10-01 01:16:31 |
| [tome](https://github.com/cyruzin/tome) | 28 | 2 | 1 | Package tome was designed to paginate simple RESTful APIs. | 2019-04-12 16:49:45 | 2021-06-09 06:37:39 |
| [ugo](https://github.com/alxrm/ugo) | 25 | 4 | 0 | Simple and expressive toolbox written in Go | 2016-02-17 19:41:57 | 2021-05-12 09:46:11 |
| [generate](https://github.com/go-playground/generate) | 24 | 4 | 0 | :runner:runs go generate recursively on a specified path or environment variable and can filter by regex | 2015-11-15 01:52:04 | 2021-03-13 21:24:07 |
| [shutdown](https://github.com/ztrue/shutdown) | 22 | 3 | 0 | Golang app shutdown hooks. | 2018-11-17 17:56:03 | 2021-09-17 17:13:57 |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 22 | 6 | 1 | a small golang lib to generate placeholder images | 2014-10-12 00:50:46 | 2020-09-16 12:56:17 |
| [rerate](https://github.com/abo/rerate) | 21 | 4 | 1 | redis-based rate counter and rate limiter | 2016-05-24 08:59:00 | 2021-08-16 19:43:35 |
| [ghokin](https://github.com/antham/ghokin) | 20 | 0 | 1 | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...) | 2018-08-03 11:36:35 | 2021-09-23 07:38:31 |
| [structs](https://github.com/PumpkinSeed/structs) | 18 | 2 | 0 | Golang struct operations. | 2017-08-26 09:59:00 | 2021-08-20 05:22:02 |
| [ctxutil](https://ctop.sh) | 18 | 2 | 1 | utils for Go context | 2018-07-30 11:28:57 | 2021-06-19 06:16:11 |
| [mimesniffer](https://pkg.go.dev/github.com/aofei/mimesniffer) | 17 | 0 | 0 | A MIME type sniffer for Go. | 2018-12-20 03:40:20 | 2021-09-22 02:10:28 |
| [dlog](https://github.com/kirillDanshin/dlog) | 16 | 1 | 0 | Simple build-time controlled debug log with ability to log where the logger was called | 2016-07-04 19:59:09 | 2020-10-19 07:34:21 |
| [filler](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 16 | 3 | 0 | fill struct data easily with fill tags | 2017-04-05 08:14:04 | 2020-12-25 16:48:53 |
| [okrun](https://github.com/xta/okrun) | 15 | 2 | 0 | ok, run your gofile | 2014-10-01 06:18:56 | 2020-07-14 23:33:14 |
| [rest-go](https://github.com/edermanoel94/rest-go) | 14 | 1 | 1 | A package that provide many helpful methods for working with rest api. | 2019-07-29 18:56:08 | 2021-09-04 04:16:01 |
| [jsend](https://clevergo.tech) | 14 | 4 | 0 | :100: JSend's implementation writen in Go(golang) | 2020-01-14 04:41:36 | 2021-06-29 03:45:23 |
| [command](https://github.com/txgruppi/command) | 14 | 3 | 0 | Command pattern for Go with thread safe serial and parallel dispatcher | 2015-08-24 09:43:50 | 2021-06-19 06:15:44 |
| [go-convert](https://github.com/Eun/go-convert) | 12 | 1 | 3 | Convert a value into another type | 2019-06-07 16:56:38 | 2021-07-15 17:03:27 |
| [copy](https://github.com/gotidy/copy) | 11 | 1 | 2 | Package for fast copying structs of different types | 2020-10-09 06:59:08 | 2021-09-28 20:31:44 |
| [ptr](https://github.com/gotidy/ptr) | 11 | 2 | 0 | Contains functions for simplified creation of pointers from constants of basic types | 2019-12-25 15:29:48 | 2021-07-24 12:55:03 |
| [retry](https://github.com/shafreeck/retry) | 11 | 1 | 1 | A pretty simple library to ensure your work to be done | 2018-07-18 09:48:33 | 2020-04-16 04:37:48 |
| [go-types](https://github.com/mikekonan/go-types) | 11 | 5 | 1 | Library providing opanapi3 and Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. | 2021-04-21 11:34:25 | 2021-07-27 09:01:27 |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 10 | 0 | 0 | Problem json implementation (https://tools.ietf.org/html/rfc7807) package for go | 2019-05-16 05:42:14 | 2021-05-22 13:33:58 |
| [silk](https://github.com/chrispassas/silk) | 10 | 1 | 0 | Read Silk Flow Files | 2018-12-18 04:23:35 | 2021-08-20 00:06:54 |
| [go-countries](https://github.com/mikekonan/go-countries) | 9 | 3 | 0 | Convert a value into another type | 2020-10-27 12:56:40 | 2021-02-03 14:41:07 |
| [statiks](https://github.com/janiltonmaciel/statiks) | 8 | 0 | 0 | Fast, zero-configuration, static HTTP filer server. | 2018-06-26 23:42:33 | 2021-08-12 18:14:07 |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 7 | 0 | 0 | Slice conversion between primitive types | 2019-02-15 06:50:34 | 2021-05-01 15:59:31 |
| [go-clip](https://github.com/prashantgupta24/go-clip) | 7 | 0 | 2 | A minimalistic clipboard manager for Mac. | 2020-11-18 22:19:01 | 2021-02-10 05:41:51 |
| [nfdump](https://github.com/chrispassas/nfdump) | 7 | 0 | 0 | NFDump File Reader | 2020-04-08 01:01:22 | 2021-08-11 17:23:16 |
| [retry](https://github.com/percolate/retry) | 7 | 1 | 0 | Percolate's Go retry package | 2018-06-15 19:23:36 | 2020-08-31 05:58:43 |
| [blank](https://github.com/Henry-Sarabia/blank) | 6 | 0 | 0 | Detect blank strings or remove whitespace from strings | 2019-02-13 00:07:27 | 2021-02-02 01:05:34 |
| [go-safe](https://github.com/kenkyu392/go-safe) | 4 | 0 | 0 | This Go package provides a sandbox for the safe execution of panic-inducing programs | 2019-10-29 15:20:37 | 2021-07-31 04:33:24 |
| [lets-go](https://github.com/aplescia/lets-go) | 3 | 0 | 0 | Go module that provides common utilities for Cloud Native development | 2020-02-19 16:32:41 | 2021-07-06 08:27:20 |
| [tik](https://github.com/andy2046/tik) | 3 | 0 | 0 | hierarchical timing wheel | 2020-07-04 09:13:49 | 2021-09-29 07:35:05 |
| [bleep](https://github.com/sinhashubham95/bleep) | 2 | 0 | 0 | OS Signal Handlers in Go | 2021-01-02 05:22:08 | 2021-09-02 21:20:24 |
| [goctx](https://github.com/zerosnake0/goctx) | 2 | 1 | 0 | Get your context value faster | 2020-11-14 14:16:09 | 2021-01-24 11:56:45 |
| [olaf](https://github.com/btnguyen2k/olaf) | 2 | 0 | 0 | Twitter Snowflake implemented in Go | 2019-01-03 13:31:10 | 2020-10-31 07:30:29 |
</details>

### Validation
Libraries for validation.

<sup>*Last Update: 2023-09-30 16:56:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [validator](https://github.com/go-playground/validator) | 14,076 | 1,212 | 246 | :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving | 2015-02-12 16:32:22 | 2023-09-30 09:32:28 |
| [govalidator](https://github.com/asaskevich/govalidator) | 5,793 | 571 | 168 | [Go] Package of validators and sanitizers for strings, numerics, slices and structs | 2014-06-20 10:45:23 | 2023-09-29 12:58:04 |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 3,350 | 213 | 47 | An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags. | 2016-06-22 03:47:43 | 2023-09-28 13:00:37 |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 1,243 | 119 | 42 | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. | 2017-09-13 16:42:20 | 2023-09-27 10:34:40 |
| [validate](https://gookit.github.io/validate/) | 948 | 111 | 20 | ⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。 | 2018-07-16 08:23:49 | 2023-09-28 03:11:46 |
| [checkdigit](https://github.com/osamingo/checkdigit) | 103 | 6 | 1 | Provide check digit algorithms and calculators written in Go | 2019-04-05 09:46:36 | 2023-09-27 03:51:33 |
| [validate](https://gookit.github.io/validate/) | 91 | 21 | 0 | This package provides a framework for writing validations for Go applications. | 2018-02-10 18:25:55 | 2023-09-30 08:21:52 |
| [jio](https://github.com/faceair/jio) | 90 | 11 | 1 | jio is a json schema validator similar to joi | 2018-10-28 11:02:45 | 2023-09-27 03:52:24 |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 80 | 9 | 1 | A norms and conventions validator for Terraform | 2019-05-29 11:37:15 | 2023-09-29 13:16:57 |
| [gody](https://pkg.go.dev/github.com/guiferpa/gody) | 64 | 5 | 1 | :balloon: A lightweight struct validator for Go | 2018-11-01 21:08:16 | 2023-09-27 03:51:42 |
| [govalid](https://github.com/twharmon/govalid) | 35 | 6 | 1 | Struct validation using tags | 2019-02-17 23:25:43 | 2023-09-27 03:51:47 |
</details>

### Version Control
Libraries for version control.

<sup>*Last Update: 2021-10-13 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-git](https://pkg.go.dev/github.com/go-git/go-git/v5) | 2,693 | 294 | 233 | A highly extensible Git implementation in pure Go. | 2019-12-19 10:27:02 | 2021-10-12 19:52:23 |
| [git2go](https://github.com/libgit2/git2go) | 1,658 | 287 | 40 | Git to Go; bindings for libgit2. Like McDonald's but tastier. | 2013-03-05 19:50:43 | 2021-10-08 07:42:22 |
| [hercules](https://github.com/src-d/hercules) | 1,449 | 122 | 38 | Gaining advanced insights from Git repository history. | 2016-12-12 17:30:29 | 2021-10-12 13:32:23 |
| [gh](https://github.com/rjeczalik/gh) | 76 | 12 | 2 | Scriptable server and net/http middleware for GitHub Webhooks. | 2015-03-08 21:04:05 | 2021-09-17 13:41:09 |
| [go-vcs](https://sourcegraph.com/sourcegraph/go-vcs) | 75 | 21 | 23 | manipulate and inspect VCS repositories in Go | 2013-06-02 02:36:18 | 2021-06-09 06:01:11 |
| [hgo](https://github.com/beyang/hgo) | 13 | 3 | 0 | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. | 2014-06-18 03:54:40 | 2020-05-05 03:52:16 |
</details>

### Video
Libraries for manipulating video.

<sup>*Last Update: 2023-09-30 16:55:53*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [goav](https://github.com/giorgisio/goav) | 2,022 | 397 | 48 | Golang bindings for FFmpeg (This repository is no longer maintained) | 2015-05-21 05:31:14 | 2023-09-25 07:09:42 |
| [m3u8](http://tools.ietf.org/html/draft-pantos-http-live-streaming) | 1,073 | 294 | 60 | Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema: | 2013-02-05 22:26:30 | 2023-09-26 19:08:46 |
| [gmf](https://github.com/3d0c/gmf) | 859 | 197 | 47 | Go Media Framework | 2013-04-03 09:07:47 | 2023-09-25 13:38:40 |
| [go-astisub](https://github.com/asticode/go-astisub) | 509 | 99 | 16 | Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.) | 2016-12-16 14:47:59 | 2023-09-25 11:49:48 |
| [go-astits](https://github.com/asticode/go-astits) | 503 | 51 | 11 | Demux and mux MPEG Transport Streams (.ts) natively in GO | 2017-07-04 13:06:15 | 2023-09-28 17:44:24 |
| [libvlc-go](https://pkg.go.dev/github.com/adrg/libvlc-go/v3) | 372 | 46 | 6 | Go bindings for libVLC and high-level media player interface | 2015-01-06 14:01:50 | 2023-09-29 05:28:21 |
| [gst](https://github.com/ziutek/gst) | 168 | 48 | 9 | Go bindings for GStreamer (retired: currently I don't use/develop this package) | 2011-07-26 00:44:40 | 2023-08-27 12:31:57 |
| [go-m3u8](https://tools.ietf.org/html/rfc8216) | 106 | 23 | 2 | Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8) | 2018-11-06 02:42:27 | 2023-07-23 07:27:11 |
| [v4l](https://pkg.go.dev/github.com/korandiz/v4l) | 77 | 14 | 0 | Facade to the Video4Linux video capture interface.  | 2016-10-25 10:50:25 | 2023-09-18 15:05:05 |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 24 | 5 | 0 | golang library to read and write various subtitle formats | 2017-05-03 21:05:25 | 2023-08-17 13:54:41 |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 18 | 6 | 0 | Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files | 2018-11-02 19:09:07 | 2023-09-18 15:04:52 |
</details>

### Web Frameworks
Full stack web frameworks.

<sup>*Last Update: 2023-10-01 20:27:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gin](https://gin-gonic.com/) | 71,754 | 7,713 | 750 | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 2014-06-16 23:57:25 | 2023-10-01 13:11:35 |
| [fiber](https://gofiber.io) | 28,564 | 1,454 | 61 | ⚡️ Express inspired web framework written in Go | 2020-01-16 03:59:20 | 2023-10-01 13:06:47 |
| [echo](https://echo.labstack.com) | 26,708 | 2,206 | 78 | High performance, minimalist Go web framework | 2015-03-01 17:43:01 | 2023-10-01 13:17:28 |
| [revel](http://revel.github.io) | 12,973 | 1,412 | 98 | A high productivity, full-stack web framework for the Go language. | 2011-12-09 04:10:26 | 2023-09-30 19:28:06 |
| [buffalo](http://gobuffalo.io) | 7,846 | 597 | 25 | Rapid Web Development w/ Go | 2014-10-22 17:35:14 | 2023-10-01 04:24:49 |
| [goa](https://goa.design) | 5,300 | 580 | 5 | Design-based APIs and microservices in Go | 2014-12-05 07:17:53 | 2023-09-30 19:24:59 |
| [gizmo](https://open.nytimes.com/introducing-gizmo-aa7ea463b208) | 3,731 | 236 | 32 | A Microservice Toolkit from The New York Times | 2015-12-15 18:09:36 | 2023-09-30 18:09:14 |
| [go-json-rest](https://ant0ine.github.io/go-json-rest/) | 3,510 | 407 | 47 | A quick and easy way to setup a RESTful JSON API | 2013-02-19 03:15:45 | 2023-09-23 15:18:35 |
| [macaron](https://go-macaron.com) | 3,427 | 292 | 6 | Package macaron is a high productive and modular web framework in Go. | 2014-07-10 03:13:30 | 2023-09-24 19:24:20 |
| [utron](https://github.com/gernest/utron) | 2,223 | 166 | 8 | A lightweight MVC framework for Go(Golang) | 2015-09-16 07:55:54 | 2023-09-24 17:19:46 |
| [goyave](https://goyave.dev) | 1,362 | 58 | 3 | 🍐 Elegant Golang REST API Framework (v5 preview available) | 2019-10-21 09:44:34 | 2023-10-01 08:28:48 |
| [rest-layer]( http://rest-layer.io) | 1,228 | 116 | 35 | REST Layer, Go (golang) REST API framework | 2015-07-29 19:16:20 | 2023-09-18 12:39:50 |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 996 | 78 | 28 | A Go framework for building JSON web services inspired by Dropwizard | 2013-02-09 21:16:13 | 2023-08-16 03:14:43 |
| [tango](https://github.com/lunny/tango) | 834 | 108 | 9 | This is only a mirror and Moved to https://gitea.com/lunny/tango | 2014-12-17 03:07:09 | 2023-09-23 18:40:09 |
| [gearbox](https://gogearbox.com) | 725 | 57 | 2 | Gearbox :gear: is a web framework written in Go with a focus on high performance | 2020-04-25 01:28:37 | 2023-09-29 02:22:16 |
| [aah](https://aahframework.org) | 685 | 39 | 19 | A secure, flexible, rapid Go web framework | 2016-06-27 04:47:45 | 2023-09-25 07:55:41 |
| [beego](beego.me) | 665 | 171 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2023-09-26 15:03:56 |
| [aero](https://github.com/aerogo/aero) | 537 | 39 | 4 | :bullettrain_side: High-performance web server for Go. | 2016-11-09 13:02:13 | 2023-10-01 13:23:41 |
| [gongular](http://gondolaweb.com) | 501 | 19 | 8 | A different approach to Go web frameworks | 2016-06-22 11:52:42 | 2023-09-10 16:07:42 |
| [flamingo-commerce](https://www.flamingo.me/flamingo-commerce.html) | 449 | 71 | 31 | Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce "Portals" and connect it with the help of individual Adapters to other services.  | 2019-04-02 15:11:57 | 2023-09-30 16:17:34 |
| [air](https://pkg.go.dev/github.com/aofei/air) | 433 | 43 | 8 | An ideally refined web framework for Go. | 2016-07-20 12:09:48 | 2023-09-20 02:47:06 |
| [neo](http://ivpusic.github.io/neo/) | 418 | 44 | 6 | Go Web Framework | 2015-02-04 19:16:06 | 2023-05-19 20:43:33 |
| [confetti](https://confetti-framework.github.io/docs/) | 410 | 17 | 1 | Confetti is a web application framework with an expressive, elegant syntax. This repository contains configuration files and is intended as a template for your codebase. Download these configuration files and include them in your git repository. | 2019-11-01 23:14:21 | 2023-09-28 04:12:20 |
| [flamingo](http://www.flamingo.me) | 401 | 44 | 39 | Flamingo Framework and Core Library. Flamingo is a go based framework to build pluggable applications. Focus is on clean architecture, maintainability and operation readiness. | 2019-04-02 12:24:02 | 2023-09-26 23:55:41 |
| [mango](http://github.com/paulbellamy/mango) | 372 | 40 | 9 | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. | 2011-05-25 07:26:46 | 2023-04-17 06:21:21 |
| [gondola](http://gondolaweb.com) | 311 | 26 | 8 | The web framework for writing faster sites, faster | 2014-07-25 21:28:55 | 2023-09-25 21:25:20 |
| [uadmin](https://uadmin.io) | 297 | 54 | 18 | The web framework for Golang | 2018-10-05 09:00:17 | 2023-09-27 20:24:57 |
| [webgo](https://github.com/bnkamalesh/webgo) | 288 | 29 | 2 | A microframework to build web apps; with handler chaining, middleware support, and most of all; standard library compliant HTTP handlers(i.e. http.HandlerFunc). | 2015-12-16 07:35:02 | 2023-09-29 12:22:59 |
| [ginrpc](https://xxjwxc.github.io/post/ginrpc/) | 282 | 34 | 11 | gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具 | 2019-06-22 12:03:53 | 2023-09-18 02:51:13 |
| [golf](https://golf.readme.io/) | 268 | 29 | 6 | :golf: The Golf web framework | 2015-11-18 15:10:14 | 2023-09-28 11:35:24 |
| [hiboot](https://hiboot.netlify.app/) | 178 | 27 | 5 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2023-09-28 10:49:02 |
| [appy](https://github.com/appist/appy) | 129 | 14 | 3 | An opinionated productive web framework that helps scaling business easier. | 2019-05-27 04:48:59 | 2023-09-11 18:13:23 |
| [go-rest](http://go.pkgdoc.org/github.com/ungerik/go-rest) | 127 | 16 | 2 | A small and evil REST framework for Go | 2012-07-13 10:02:15 | 2023-03-07 16:06:33 |
| [patron](https://github.com/beatlabs/patron) | 124 | 67 | 26 | Microservice framework following best cloud practices with a focus on productivity. | 2019-01-30 13:49:54 | 2023-07-24 11:12:32 |
| [microservice](https://github.com/claygod/microservice) | 113 | 15 | 0 | This library provides a simple microservice framework based on clean architecture principles with a working example implemented. | 2016-12-15 09:07:04 | 2023-09-02 09:54:41 |
| [rux](https://pkg.go.dev/github.com/gookit/rux?tab=doc) | 94 | 15 | 2 | ⚡ Rux is an simple and fast web framework. support route group, param route binding, middleware, compatible http.Handler interface. 简单且快速的 Go api/web 框架，支持路由分组，路由参数绑定，中间件，兼容 http.Handler 接口 | 2018-08-05 06:13:57 | 2023-09-12 00:07:15 |
| [vox](https://aisk.github.io/vox/) | 80 | 7 | 8 | Simple and lightweight Go web framework inspired by koa | 2014-12-24 11:22:08 | 2023-04-08 05:26:50 |
| [golax](https://github.com/fulldump/golax) | 75 | 7 | 6 | Golax, a go implementation for the Lax framework. | 2016-01-30 19:11:39 | 2023-09-27 21:43:32 |
| [yarf](https://github.com/yarf-framework/yarf) | 67 | 8 | 2 | Yet Another REST Framework | 2015-09-02 13:56:47 | 2023-09-29 12:22:35 |
| [fireball](https://github.com/ridgelines/fireball) | 59 | 6 | 1 | Go web framework with a natural feel | 2016-07-20 05:04:54 | 2022-11-09 19:23:48 |
| [goa](https://goa-go.github.io) | 49 | 4 | 0 | Goa is a  web framework based on middleware, like koa.js. | 2019-07-26 07:12:23 | 2023-06-06 18:44:28 |
| [gotuna](https://gotuna.netlify.app) | 44 | 23 | 1 | GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server. | 2021-04-08 14:08:08 | 2023-09-11 12:03:01 |
| [goweb](https://github.com/twharmon/goweb) | 35 | 6 | 1 | Lightweight web framework based on net/http. | 2019-05-07 21:04:43 | 2023-03-16 06:12:26 |
| [api](http://resoursea.com/) | 33 | 4 | 0 | A REST framework for quickly writing resource based services in Golang. | 2015-01-24 18:45:30 | 2023-03-17 14:36:12 |
| [rex](https://github.com/goanywhere/rex) | 33 | 3 | 0 | Pleasures for Web in Golang | 2014-10-16 02:26:18 | 2022-09-27 10:03:38 |
| [banjo](https://nsheremet.pw/banjo) | 21 | 7 | 4 | BANjO is a simple web framework written in Go (golang) | 2017-12-09 13:35:31 | 2023-03-16 01:57:53 |
</details>

### Web Frameworks - Middlewares - Actual middlewares


<sup>*Last Update: 2021-10-07 09:25:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tollbooth](https://github.com/didip/tollbooth) | 2,023 | 184 | 5 | Simple middleware to rate-limit HTTP requests. | 2015-05-17 15:20:03 | 2021-10-04 17:18:58 |
| [cors](https://github.com/rs/cors) | 1,937 | 172 | 18 | Go net/http configurable handler to handle CORS requests | 2014-10-25 03:49:45 | 2021-10-06 05:41:00 |
| [limiter](https://github.com/ulule/limiter) | 1,344 | 108 | 9 | Dead simple rate limit middleware for Go. | 2015-10-02 08:12:38 | 2021-10-05 12:39:42 |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 823 | 28 | 5 | Go (golang) library for creating and consuming HTTP Server-Timing headers | 2018-02-12 03:56:02 | 2021-10-04 19:49:44 |
| [go-fault](https://github.com/github/go-fault) | 417 | 17 | 0 | Fault injection library in Go using standard http middleware | 2020-05-14 16:13:17 | 2021-10-02 01:44:30 |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 114 | 7 | 17 | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ | 2018-06-29 21:51:00 | 2021-09-17 14:49:42 |
| [xff](https://github.com/sebest/xff) | 85 | 19 | 5 | A Golang Middleware to handle X-Forwarded-For Header | 2014-12-22 10:29:05 | 2021-09-06 13:28:00 |
| [formjson](https://github.com/rs/formjson) | 36 | 1 | 0 | Go net/http handler to transparently manage posted JSON | 2015-03-19 23:52:28 | 2020-09-18 01:35:42 |
| [client-timing](https://github.com/posener/client-timing) | 19 | 4 | 1 | An HTTP client for go-server-timing middleware. Enables automatic timing propagation through HTTP calls between servers. | 2018-02-23 01:52:45 | 2021-04-21 08:17:29 |
</details>

### Web Frameworks - Middlewares - Libraries for creating HTTP middlewares


<sup>*Last Update: 2023-09-30 07:16:56*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [negroni](https://github.com/urfave/negroni) | 7,363 | 594 | 9 | Idiomatic HTTP Middleware for Golang | 2014-05-18 22:09:10 | 2023-09-25 04:55:23 |
| [alice](https://godoc.org/github.com/justinas/alice) | 2,860 | 175 | 8 | Painless middleware chaining for Go | 2014-05-25 07:27:41 | 2023-09-29 09:43:32 |
| [render](https://github.com/unrolled/render) | 1,840 | 143 | 0 | Go package for easily rendering JSON, XML, binary data, and HTML templates responses. | 2014-06-10 16:20:35 | 2023-09-29 16:27:03 |
| [stats](https://github.com/thoas/stats) | 593 | 53 | 9 | A Go middleware that stores various information about your web application (response time, status code count, etc.) | 2015-03-05 18:02:50 | 2023-08-01 19:47:46 |
| [interpose](https://github.com/carbocation/interpose) | 293 | 18 | 1 | Minimalist net/http middleware for golang | 2014-07-20 00:19:52 | 2023-08-21 11:24:56 |
| [renderer](https://github.com/thedevsaddam/renderer) | 256 | 30 | 0 | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go | 2017-11-07 18:53:49 | 2023-09-26 20:48:44 |
| [muxchain](https://github.com/stephens2424/muxchain) | 208 | 14 | 1 | Lightweight Middleware for net/http | 2014-05-03 17:14:17 | 2023-05-05 11:26:07 |
| [rye](https://github.com/InVisionApp/rye) | 102 | 15 | 0 | A tiny http middleware for Golang with added handlers for common needs. | 2016-10-06 19:51:59 | 2023-07-29 18:28:02 |
| [gores](https://github.com/alioygur/gores) | 100 | 4 | 0 | Go package that handles HTML, JSON, XML and etc. responses | 2015-12-25 12:41:01 | 2023-02-28 15:29:17 |
| [mediary](https://github.com/HereMobilityDevelopers/mediary/wiki/Reasoning) | 87 | 7 | 0 | Add interceptors to GO http.Client | 2020-03-23 18:54:56 | 2023-08-07 12:00:37 |
| [chain](https://github.com/codemodus/chain) | 63 | 3 | 0 | Composable chains of nested http.Handler instances. | 2015-05-14 19:52:58 | 2023-09-16 01:41:09 |
| [wrap](https://github.com/go-on/wrap) | 59 | 5 | 0 | Go http.Hander based middleware stack with context sharing | 2014-02-16 07:12:36 | 2020-08-28 05:29:07 |
| [catena](https://github.com/codemodus/catena) | 9 | 2 | 0 | gRPC interceptor catenation. | 2015-07-30 19:07:01 | 2022-09-27 10:06:05 |
</details>

### Web Frameworks - Routers


<sup>*Last Update: 2023-09-30 09:09:12*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mux](https://gorilla.github.io) | 19,238 | 1,807 | 17 | Package gorilla/mux is a powerful HTTP router and URL matcher for building Go web servers with 🦍 | 2012-10-02 21:32:24 | 2023-09-29 19:01:45 |
| [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) | 15,802 | 1,460 | 84 | A high performance HTTP request router that scales well | 2013-12-05 15:10:55 | 2023-09-29 17:35:09 |
| [chi](https://go-chi.io) | 15,414 | 947 | 65 | lightweight, idiomatic and composable router for building Go HTTP services | 2015-10-15 20:46:29 | 2023-09-30 00:22:28 |
| [web](https://github.com/gocraft/web) | 1,501 | 151 | 23 | Go Router + Middleware. Your Contexts. | 2013-11-16 20:48:20 | 2023-09-24 14:43:27 |
| [bone](http://go-zoo.github.io/bone) | 1,292 | 90 | 3 | Lightning Fast HTTP Multiplexer | 2014-11-19 02:16:36 | 2023-09-27 05:13:02 |
| [goji](https://goji.io) | 946 | 72 | 6 | Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) | 2015-11-16 00:52:41 | 2023-09-23 20:30:10 |
| [fasthttprouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 875 | 92 | 19 | A high performance fasthttp request router that scales well | 2015-12-13 09:32:30 | 2023-09-23 20:29:33 |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 606 | 56 | 7 | High-speed, flexible tree-based HTTP router for Go. | 2014-05-14 20:10:20 | 2023-09-19 09:39:45 |
| [gorouter](https://xujiajun.cn/gorouter) | 532 | 87 | 0 | xujiajun/gorouter is a simple and fast HTTP router for Go. It is easy to build RESTful APIs and your web framework. | 2018-01-29 09:28:28 | 2023-08-25 19:53:08 |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 445 | 52 | 11 | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. | 2015-10-27 01:03:14 | 2023-09-05 12:52:58 |
| [lars](https://github.com/go-playground/lars) | 390 | 25 | 1 | :rotating_light: Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. | 2015-12-24 17:28:45 | 2023-06-14 04:55:34 |
| [siesta](https://github.com/VividCortex/siesta) | 351 | 15 | 0 | Composable framework for writing HTTP handlers in Go. | 2014-09-23 13:55:56 | 2023-04-07 21:48:35 |
| [vestigo](https://github.com/husobee/vestigo) | 269 | 31 | 13 | Echo Inspired Stand Alone URL Router | 2015-09-22 03:08:03 | 2023-06-13 00:23:19 |
| [router](https://github.com/gowww/router) | 166 | 14 | 0 | ⚡️ A lightning fast HTTP router | 2017-05-25 10:29:27 | 2023-09-11 15:16:57 |
| [gorouter](https://rafallorenz.com/gorouter) | 150 | 16 | 10 | Go Server/API micro framework, HTTP request router, multiplexer, mux | 2016-07-14 13:13:34 | 2023-09-15 16:49:01 |
| [pure](https://github.com/go-playground/pure) | 142 | 12 | 0 | :non-potable_water: Is a lightweight  HTTP router that sticks to the std "net/http" implementation | 2016-09-23 19:57:58 | 2023-09-04 21:19:45 |
| [alien](https://github.com/gernest/alien) | 127 | 14 | 3 | A lightweight and  fast http router from outer space | 2016-01-30 23:23:10 | 2023-09-28 10:32:23 |
| [violetear](http://violetear.org) | 107 | 11 | 1 | Go HTTP router | 2015-06-19 16:49:41 | 2023-09-25 21:28:11 |
| [Bxog](http://go-zoo.github.io/bone) | 104 | 8 | 0 | Bxog is a simple and fast HTTP router for Go (HTTP request multiplexer). | 2016-05-19 12:20:08 | 2023-06-28 11:21:56 |
| [xmux](http://violetear.org) | 97 | 11 | 2 | xmux is a httprouter fork on top of xhandler (net/context aware) | 2015-12-14 19:01:05 | 2023-08-05 16:14:30 |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 54 | 7 | 0 | :bell: A simple Go router | 2019-02-21 13:13:52 | 2023-02-20 06:49:25 |
| [fastrouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 22 | 5 | 0 | FastRouter is a fast, flexible HTTP router written in Go. | 2017-11-01 08:52:52 | 2022-09-27 10:08:01 |
| [route](https://goroute.github.io) | 8 | 2 | 1 | Go Route - Simple yet powerful HTTP request multiplexer | 2019-07-06 18:47:38 | 2023-04-05 19:56:17 |
</details>

### WebAssembly


<sup>*Last Update: 2021-08-23 09:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tinygo](https://tinygo.org) | 8,438 | 447 | 294 | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. | 2018-06-07 16:39:19 | 2021-08-23 01:50:30 |
| [dom](https://github.com/dennwc/dom) | 434 | 52 | 11 | DOM library for Go and WASM | 2018-06-30 18:37:35 | 2021-08-20 11:17:05 |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 140 | 9 | 5 | Library to use HTML5 Canvas  from Go-WASM, with all drawing within go code | 2019-05-05 14:05:55 | 2021-08-14 03:32:07 |
| [webapi](https://github.com/gowebapi/webapi) | 86 | 8 | 2 | Go Lang Web Assembly bindings for DOM, HTML etc | 2019-02-08 05:58:35 | 2021-07-15 15:53:33 |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 82 | 9 | 0 | Run WASM tests inside your browser | 2018-07-14 18:42:24 | 2021-07-21 10:18:04 |
| [vert](https://github.com/norunners/vert) | 50 | 6 | 2 | WebAssembly interop between Go and JS values. | 2018-03-25 17:26:47 | 2021-08-09 15:48:06 |
</details>

### Windows


<sup>*Last Update: 2023-10-02 20:50:44*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-ole](https://github.com/go-ole/go-ole) | 1,025 | 213 | 72 | win32 ole implementation for golang | 2011-01-21 12:45:20 | 2023-09-26 11:08:50 |
| [d3d9](https://github.com/gonutz/d3d9) | 147 | 13 | 1 | Direct3D9 wrapper for Go. | 2015-12-12 21:24:38 | 2023-09-02 13:30:59 |
| [gosddl](https://github.com/MonaxGT/gosddl) | 10 | 2 | 0 | GoSDDL converter | 2018-12-04 08:36:11 | 2023-04-23 09:37:32 |
</details>

### XML
Libraries and tools for manipulating XML.

<sup>*Last Update: 2021-09-07 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [zek](https://github.com/miku/zek) | 494 | 42 | 8 | Generate a Go struct from XML. | 2017-11-23 19:03:11 | 2021-09-06 19:43:01 |
| [xpath](https://github.com/antchfx/xpath) | 413 | 61 | 10 | XPath package for Golang, supports HTML, XML, JSON document query. | 2016-10-09 05:51:24 | 2021-09-03 06:21:56 |
| [xquery](https://github.com/antchfx/xpath) | 154 | 26 | 0 | Extract data or evaluate value from HTML/XML documents using XPath | 2016-10-09 05:54:10 | 2021-08-16 09:40:30 |
| [xml2map](https://github.com/sbabiv/xml2map) | 34 | 6 | 1 | XML to MAP converter written Golang | 2018-08-06 17:51:46 | 2021-08-22 15:50:01 |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 20 | 3 | 1 | xmlwriter is a pure-Go library providing procedural XML generation based on libxml2's xmlwriter module | 2017-04-11 04:43:26 | 2021-07-29 01:17:10 |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 8 | 8 | Compare ANY markup documents. | 2016-10-25 22:09:12 | 2021-04-23 04:41:47 |
</details>

