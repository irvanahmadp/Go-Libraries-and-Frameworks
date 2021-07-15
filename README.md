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

<sup>*Last Update: 2021-07-01 08:37:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oto](https://github.com/hajimehoshi/oto) | 826 | 69 | 33 | ♪ A low-level library to play sound on multiple platforms ♪ | 2017-05-04 12:16:30 | 2021-06-30 07:10:18 |
| [portaudio](https://github.com/gordonklaus/portaudio) | 453 | 71 | 9 | Go bindings for the PortAudio audio I/O library | 2015-09-16 07:59:23 | 2021-06-30 20:30:02 |
| [music-theory](https://gopkg.in/music-theory.v0) | 342 | 35 | 8 | Go models of Note, Scale, Chord and Key | 2016-03-17 03:50:17 | 2021-06-15 04:26:32 |
| [waveform](https://github.com/mdlayher/waveform) | 338 | 28 | 4 | Go package capable of generating waveform images from audio streams. MIT Licensed. | 2014-09-13 18:07:36 | 2021-06-20 22:19:43 |
| [portmidi](https://github.com/rakyll/portmidi) | 258 | 49 | 13 | Go bindings for libportmidi | 2013-11-10 14:24:56 | 2021-06-25 08:20:41 |
| [id3v2](https://pkg.go.dev/github.com/bogem/id3v2) | 190 | 30 | 14 | 🎵 ID3 decoding and encoding library for Go | 2016-05-15 18:36:53 | 2021-05-21 04:31:20 |
| [flac](https://github.com/mewkiz/flac) | 158 | 30 | 10 | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 2012-11-01 20:13:58 | 2021-06-23 04:06:29 |
| [malgo](https://pkg.go.dev/github.com/bogem/id3v2) | 150 | 23 | 1 | Mini audio library | 2017-11-09 18:27:52 | 2021-06-26 08:50:01 |
| [mix](https://gopkg.in/mix.v0) | 138 | 22 | 11 | Sequence-based Go-native audio mixer for music apps | 2016-01-03 15:53:57 | 2021-06-14 00:15:40 |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 137 | 9 | 4 | Go tools for audio processing & creation 🎶 | 2020-07-05 01:35:41 | 2021-06-20 22:33:21 |
| [gaad](https://github.com/Comcast/gaad) | 84 | 12 | 4 | GAAD (Go Advanced Audio Decoder) | 2016-07-11 14:19:16 | 2021-06-09 06:01:06 |
| [minimp3](https://github.com/tosone/minimp3) | 55 | 9 | 0 | Decode mp3 base on https://github.com/lieff/minimp3 | 2018-01-26 09:10:31 | 2021-06-09 06:01:11 |
| [vorbis](https://github.com/mccoyst/vorbis) | 28 | 4 | 2 | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) | 2013-07-12 02:45:39 | 2020-10-04 21:02:12 |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 11 | 6 | 0 | Go Bindings for libsamplerate | 2016-11-20 21:19:31 | 2021-06-22 03:49:24 |
</details>

### Authentication and OAuth
Libraries for implementing authentications schemes.

<sup>*Last Update: 2021-07-02 09:03:58*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [oauth2](https://golang.org/x/oauth2) | 3,702 | 765 | 140 | Go OAuth2 | 2014-04-14 15:07:35 | 2021-07-02 01:39:05 |
| [goth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 3,217 | 393 | 59 | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 2014-10-14 20:38:12 | 2021-07-01 14:48:25 |
| [authboss](https://github.com/volatiletech/authboss) | 2,694 | 170 | 30 | The boss of http auth. | 2015-01-03 05:12:02 | 2021-06-28 10:57:04 |
| [loginsrv](https://github.com/tarent/loginsrv) | 1,791 | 155 | 29 | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 2016-11-11 12:11:21 | 2021-07-01 22:31:03 |
| [go-jose](https://github.com/square/go-jose) | 1,778 | 315 | 49 | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 2014-11-14 18:27:31 | 2021-07-01 00:24:34 |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1,772 | 271 | 28 | A standalone, specification-compliant,  OAuth2 server written in Golang. | 2015-11-01 13:30:09 | 2021-06-29 11:51:26 |
| [osin](https://golang.org/x/oauth2) | 1,673 | 370 | 3 | Golang OAuth2 server library | 2013-09-10 19:52:00 | 2021-06-27 08:53:34 |
| [gologin](https://github.com/dghubble/gologin) | 1,404 | 111 | 0 | Go login handlers for authentication providers (OAuth1, OAuth2) | 2015-06-23 04:40:52 | 2021-07-01 13:25:42 |
| [gorbac](https://github.com/mikespook/gorbac) | 1,187 | 140 | 2 | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 2013-12-26 10:00:41 | 2021-06-28 18:21:30 |
| [scs](https://github.com/alexedwards/scs) | 892 | 84 | 18 | HTTP Session Management for Go | 2016-08-08 16:42:05 | 2021-06-25 20:09:29 |
| [paseto](https://github.com/o1egl/paseto) | 519 | 22 | 4 | Platform-Agnostic Security Tokens implementation in GO (Golang) | 2018-01-23 05:27:39 | 2021-07-01 19:22:26 |
| [permissions2](https://github.com/xyproto/permissions2) | 430 | 33 | 0 |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 2014-11-19 12:23:37 | 2021-06-18 13:01:24 |
| [go-guardian](https://github.com/shaj13/go-guardian) | 273 | 22 | 6 | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 2020-05-14 12:15:56 | 2021-07-01 20:25:30 |
| [jwt](https://github.com/cristalhq/jwt) | 272 | 23 | 3 | Safe, simple and fast JSON Web Tokens for Go | 2019-07-20 18:14:58 | 2021-07-01 06:40:45 |
| [jwt](https://github.com/pascaldekloe/jwt) | 257 | 15 | 0 | JSON Web Token library | 2018-03-21 11:59:53 | 2021-07-01 06:40:45 |
| [jeff](https://github.com/abraithwaite/jeff) | 227 | 11 | 1 | 🍍Jeff provides the simplest way to manage web sessions in Go. | 2018-08-02 19:31:23 | 2021-07-01 20:52:05 |
| [httpauth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 206 | 21 | 4 | HTTP Authentication middlewares | 2014-05-26 22:53:57 | 2021-06-09 22:15:48 |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 204 | 35 | 3 | This package provides json web token (jwt) middleware for goLang http servers | 2016-07-05 23:31:43 | 2021-05-24 01:31:29 |
| [branca](https://branca.io) | 153 | 16 | 1 | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. | 2018-01-09 15:27:31 | 2021-06-23 01:35:45 |
| [sessionup](https://github.com/swithek/sessionup) | 112 | 5 | 3 | Straightforward HTTP session management | 2019-07-23 18:55:21 | 2021-06-08 14:33:23 |
| [session](https://github.com/icza/session) | 105 | 12 | 4 | Go session management for web servers (including support for Google App Engine - GAE). | 2016-02-08 09:07:07 | 2021-06-24 17:57:35 |
| [jwt](https://github.com/robbert229/jwt) | 90 | 21 | 9 | This is an implementation of JWT in golang! | 2016-06-05 22:01:37 | 2021-06-13 21:57:30 |
| [sjwt](https://godoc.org/github.com/brianvoe/sjwt) | 90 | 5 | 0 | Simple JWT Golang | 2019-06-20 04:06:21 | 2021-06-25 20:36:28 |
| [rbac](https://github.com/zpatrick/rbac) | 79 | 11 | 0 | Minimalistic RBAC package for Go applications | 2018-08-02 00:11:04 | 2021-05-31 10:33:51 |
| [sessions](https://github.com/adam-hanna/sessions) | 59 | 6 | 2 | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 2017-04-29 01:09:28 | 2021-04-23 00:07:10 |
| [securecookie](https://github.com/chmike/securecookie) | 51 | 7 | 1 | Fast, secure and efficient secure cookie encoder/decoder  | 2017-09-03 14:40:29 | 2021-07-01 12:20:10 |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 21 | 1 | 0 | Golang library for providing a canonical representation of email address. | 2020-08-21 23:13:04 | 2021-06-27 08:36:10 |
| [otpgo](https://github.com/jltorresm/otpgo) | 20 | 0 | 1 | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 2020-08-19 13:20:23 | 2021-06-20 07:17:30 |
| [scope](https://github.com/SonicRoshan/scope) | 12 | 3 | 0 | Easily Manage OAuth2 Scopes In Go | 2019-09-23 10:48:14 | 2021-04-25 07:45:30 |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 1 | 0 | A driver for the SessionGate Redis module - easy session management using the Go language. | 2017-10-20 03:39:11 | 2020-08-18 23:01:11 |
| [cookiestxt](https://casbin.org/) | 6 | 2 | 0 | cookiestxt implement parser of cookies txt format | 2017-10-09 11:27:19 | 2021-03-08 11:45:58 |
| [casbin](https://casbin.org/) | 1 | 0 | 0 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2021-05-29 04:09:46 | 2021-06-23 04:53:51 |
</details>

### Bot Building
Libraries for building and working with bots.

<sup>*Last Update: 2021-07-05 02:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [olivia](https://olivia-ai.org) | 2,940 | 292 | 18 | 💁‍♀️Your new best friend powered by an artificial neural network | 2018-06-05 18:19:31 | 2021-07-02 01:30:19 |
| [telegram-bot-api](https://go-telegram-bot-api.dev) | 2,935 | 489 | 48 | Golang bindings for the Telegram Bot API | 2015-06-25 05:33:57 | 2021-07-04 10:54:00 |
| [telebot](https://github.com/tucnak/telebot) | 1,964 | 271 | 31 | Telebot is a Telegram bot framework in Go. | 2015-06-25 19:27:50 | 2021-07-04 11:17:31 |
| [kelp](https://kelpbot.io) | 700 | 155 | 153 | Kelp is a free and open-source trading bot for the Stellar DEX and 100+ centralized exchanges | 2018-08-08 23:31:18 | 2021-07-03 20:19:55 |
| [bot](https://github.com/go-chat-bot/bot) | 683 | 169 | 10 | IRC, Slack, Telegram and RocketChat bot written in go | 2015-09-22 16:41:13 | 2021-06-30 16:18:38 |
| [golang-crypto-trading-bot](https://github.com/saniales/golang-crypto-trading-bot) | 567 | 137 | 10 | A golang implementation of a console-based trading bot for cryptocurrency exchanges | 2017-05-14 22:11:41 | 2021-07-04 08:08:04 |
| [slacker](https://github.com/shomali11/slacker) | 522 | 68 | 4 | Slack Bot Framework | 2017-05-20 01:41:20 | 2021-07-03 02:20:57 |
| [joe](https://joe-bot.net) | 424 | 24 | 5 | A general-purpose bot library inspired by Hubot but written in Go.   :robot: | 2019-03-03 11:19:18 | 2021-07-01 06:23:36 |
| [tbot](https://yanzay.github.io/tbot-doc/) | 314 | 43 | 0 | Go library for Telegram Bot API | 2015-09-11 16:19:25 | 2021-06-30 05:45:41 |
| [go-sarah](https://pkg.go.dev/github.com/oklahomer/go-sarah/v4) | 195 | 12 | 0 | Simple yet customizable bot framework written in Go. | 2016-11-06 10:04:43 | 2021-07-03 02:29:41 |
| [go-twitch-irc](https://github.com/gempir/go-twitch-irc) | 174 | 36 | 11 | go irc client for twitch.tv | 2017-03-23 21:31:35 | 2021-06-16 18:55:41 |
| [tenyks](http://tenyks.io) | 169 | 18 | 12 | The Tenyks IRC bot. | 2012-08-26 02:02:24 | 2021-04-16 01:57:30 |
| [hanu](https://sbstjn.com/host-golang-slackbot-on-heroku-with-hanu.html) | 135 | 20 | 2 | Golang Framework for writing Slack bots | 2016-09-16 07:10:42 | 2021-06-15 05:09:35 |
| [go-tgbot](https://github.com/olebedev/go-tgbot) | 107 | 4 | 2 | Golang  telegram bot API wrapper, session-based router and middleware | 2016-12-11 06:06:32 | 2021-05-28 08:08:33 |
| [margelet](https://kelpbot.io) | 63 | 10 | 0 | Telegram Bot Framework for Go | 2015-11-21 13:02:17 | 2021-04-27 09:11:42 |
| [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) | 50 | 5 | 8 | A Discord bot for managing ephemeral roles based upon voice channel member presence. | 2017-12-19 15:20:30 | 2021-06-22 06:13:16 |
| [slackscot](https://github.com/alexandre-normand/slackscot) | 45 | 8 | 1 | Slack bot core/framework written in Go with support for reactions to message updates/deletes | 2015-10-22 04:54:55 | 2021-06-27 12:21:56 |
| [slack-bot](https://github.com/innogames/slack-bot) | 40 | 13 | 7 | Ready to use Slack bot for lazy developers: start Jenkins jobs, watch Jira tickets, watch pull requests... | 2019-07-19 07:49:06 | 2021-07-03 19:18:48 |
| [govkbot](https://github.com/nikepan/govkbot) | 36 | 2 | 1 | VK bot package for Go | 2016-07-11 22:09:54 | 2021-04-08 06:24:17 |
| [micha](https://github.com/onrik/micha) | 17 | 3 | 0 | Client lib for Telegram bot api | 2016-04-14 12:09:44 | 2021-06-16 16:32:41 |
| [echotron](https://github.com/NicoNex/echotron) | 13 | 2 | 0 | Library for telegram bots written in pure go. | 2019-07-22 17:31:49 | 2021-06-10 07:27:49 |
</details>

### Build Automation
Libraries and tools helping with build automation.

<sup>*Last Update: 2021-07-01 10:25:15*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [realize](https://github.com/oxequa/realize) | 4,044 | 211 | 68 | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 2016-07-12 08:07:25 | 2021-06-30 15:00:12 |
| [task](https://taskfile.dev) | 3,591 | 230 | 90 | A task runner / simpler Make alternative written in Go | 2017-02-27 00:46:04 | 2021-06-30 23:34:06 |
| [mmake](https://github.com/tj/mmake) | 1,577 | 43 | 11 | Modern Make  | 2017-02-15 22:01:21 | 2021-06-29 16:34:19 |
| [goyek](https://github.com/goyek/goyek) | 234 | 15 | 2 | Create build pipelines in Go  | 2020-10-11 13:20:55 | 2021-06-28 02:00:26 |
| [taskctl](https://github.com/taskctl/taskctl) | 114 | 10 | 7 | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 2019-11-12 13:19:09 | 2021-06-25 05:41:16 |
| [1build](https://1build.gitbook.io) | 95 | 25 | 31 | Frictionless way of managing project-specific commands | 2019-04-23 17:05:38 | 2021-06-24 17:57:32 |
| [gilbert](https://go-gilbert.github.io/) | 90 | 4 | 0 | Build system and task runner for Go projects | 2019-01-30 09:02:31 | 2021-06-28 12:18:35 |
| [gaper](https://github.com/maxcnunes/gaper) | 50 | 3 | 7 | Builds and restarts a Go project when it crashes or some watched file changes | 2018-06-16 02:46:38 | 2021-06-28 12:18:04 |
| [anko](https://github.com/GuilhermeCaruso/anko) | 16 | 0 | 0 | :crystal_ball: Simple application watcher | 2021-03-02 14:08:42 | 2021-06-18 10:08:52 |
</details>

### CSS Preprocessors
Libraries for preprocessing CSS files.

<sup>*Last Update: 2021-06-28 15:13:41*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gcss](https://github.com/yosssi/gcss) | 442 | 33 | 8 | Pure Go CSS Preprocessor | 2014-09-04 14:38:20 | 2021-05-26 15:17:44 |
| [go-libsass](http://godoc.org/github.com/wellington/go-libsass) | 179 | 22 | 13 | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 2015-04-19 15:09:47 | 2021-06-23 15:09:17 |
</details>

### Command Line - Advanced Console UIs
Libraries for building Console Applications and Console User Interfaces.

<sup>*Last Update: 2021-07-05 02:25:05*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [termui](https://github.com/gizak/termui) | 11,051 | 704 | 83 | Golang terminal dashboard | 2015-02-03 14:09:27 | 2021-07-04 10:33:12 |
| [gocui](https://github.com/jroimartin/gocui) | 7,229 | 482 | 70 | Minimalist Go package aimed at creating Console User Interfaces. | 2014-01-04 02:50:20 | 2021-07-02 00:26:43 |
| [termbox-go](http://godoc.org/github.com/nsf/termbox-go) | 4,080 | 355 | 41 | Pure Go termbox implementation | 2012-01-12 21:03:03 | 2021-06-30 01:05:58 |
| [go-prompt](https://godoc.org/github.com/c-bata/go-prompt) | 4,001 | 243 | 82 | Building powerful interactive prompts in Go, inspired by python-prompt-toolkit. | 2017-08-14 16:02:09 | 2021-07-02 07:37:37 |
| [progressbar](https://pkg.go.dev/github.com/schollz/progressbar/v3?tab=doc) | 1,921 | 103 | 11 | A really basic thread-safe progress bar for Golang applications | 2017-10-26 18:28:10 | 2021-07-02 10:28:04 |
| [uiprogress](https://github.com/gosuri/uiprogress) | 1,805 | 120 | 25 | A go library to render progress bars in terminal applications | 2015-11-17 00:59:24 | 2021-07-01 06:21:36 |
| [termdash](https://github.com/mum4k/termdash) | 1,724 | 89 | 41 | Terminal based dashboard. | 2018-03-24 12:01:49 | 2021-07-04 10:09:07 |
| [asciigraph](https://github.com/guptarohit/asciigraph) | 1,691 | 61 | 5 | Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. | 2018-06-17 10:37:16 | 2021-06-25 00:09:02 |
| [pterm](https://pterm.sh) | 1,393 | 50 | 21 | ✨ #PTerm is a modern go module to beautify console output. Featuring charts, progressbars, tables, trees, and many more 🚀 It's completely configurable and 100% cross-platform compatible. | 2020-09-17 15:52:59 | 2021-07-02 05:53:33 |
| [mpb](https://github.com/vbauerster/mpb) | 1,316 | 78 | 2 | multi progress bar for Go cli applications | 2016-12-14 11:56:29 | 2021-07-02 05:58:16 |
| [uilive](https://github.com/gosuri/uilive) | 1,314 | 65 | 9 | uilive is a go library for updating terminal output in realtime | 2015-11-16 06:13:10 | 2021-06-24 17:52:31 |
| [aurora](https://github.com/logrusorgru/aurora) | 1,110 | 50 | 4 | Golang ultimate ANSI-colors that supports Printf/Sprintf methods | 2016-11-06 22:37:12 | 2021-06-30 02:08:05 |
| [color](https://pkg.go.dev/github.com/gookit/color?tab=overview) | 864 | 57 | 1 | 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows. GO CLI 控制台颜色渲染工具库，支持16色，256色，RGB色彩渲染输出，使用类似于 Print/Sprintf，兼容并支持 Windows 环境的色彩渲染 | 2018-07-01 07:28:17 | 2021-07-02 01:47:44 |
| [uitable](https://github.com/gosuri/uitable) | 613 | 25 | 3 | A go library to improve readability in terminal apps using tabular data | 2015-11-13 21:59:21 | 2021-06-24 17:54:38 |
| [go-colorable](http://godoc.org/github.com/mattn/go-colorable) | 535 | 67 | 10 | Another Text Attribute Manupulator | 2014-07-30 02:38:06 | 2021-06-29 16:30:58 |
| [go-isatty](http://godoc.org/github.com/mattn/go-isatty) | 529 | 70 | 4 | Change the color of console text. | 2014-04-01 01:53:09 | 2021-07-02 01:48:02 |
| [gommon](https://github.com/labstack/gommon) | 397 | 85 | 15 | Common packages for Go | 2015-03-12 22:35:57 | 2021-06-24 17:56:00 |
| [chalk](https://github.com/ttacon/chalk) | 375 | 18 | 4 | Intuitive package for prettifying terminal/console output. http://godoc.org/github.com/ttacon/chalk | 2014-07-18 19:38:58 | 2021-06-28 17:29:40 |
| [simpletable](https://github.com/alexeyco/simpletable) | 317 | 24 | 2 | Simple tables in terminal with Go | 2017-03-29 07:27:23 | 2021-06-29 16:28:48 |
| [tabby](https://github.com/cheynewallace/tabby) | 292 | 11 | 2 | A tiny library for super simple Golang tables | 2018-12-17 23:35:39 | 2021-06-29 16:29:01 |
| [go-colortext](http://godoc.org/github.com/mattn/go-colorable) | 207 | 19 | 4 | Change the color of console text. | 2013-01-23 03:38:54 | 2021-07-01 15:34:18 |
| [yacspin](https://github.com/theckman/yacspin) | 183 | 4 | 0 | Yet Another CLi Spinner; providing over 70 easy to use and customizable terminal spinners for multiple OSes | 2019-12-29 07:41:23 | 2021-07-03 18:53:00 |
| [box-cli-maker](https://github.com/Delta456/box-cli-maker) | 135 | 2 | 0 | Make Highly Customized Boxes for your CLI | 2020-05-01 07:23:56 | 2021-06-26 11:29:44 |
| [cfmt](https://github.com/mingrammer/cfmt) | 78 | 5 | 1 | :art: Contextual fmt inspired by bootstrap color classes | 2018-03-15 19:04:27 | 2021-03-09 07:48:50 |
| [tabular](https://github.com/InVisionApp/tabular) | 55 | 5 | 0 | Tabular simplifies printing ASCII tables from command line utilities | 2018-04-23 21:17:03 | 2021-05-18 22:23:53 |
| [ctc](https://github.com/wzshiming/ctc) | 33 | 2 | 0 | Console Text Colors - The non-invasive cross-platform terminal color library does not need to modify the Print method | 2018-04-27 18:07:42 | 2021-06-06 09:01:15 |
| [cfmt](https://github.com/i582/cfmt) | 26 | 0 | 0 | Small library for simple and convenient formatted stylized output to the console. | 2020-11-13 20:29:45 | 2021-06-29 16:13:46 |
| [colourize](https://github.com/TreyBastian/colourize) | 24 | 3 | 0 | An ANSI colour terminal package for Go | 2015-05-11 11:49:39 | 2021-03-30 22:37:59 |
| [marker](https://github.com/cyucelen/marker) | 19 | 13 | 4 |  🖍️ Marker is the easiest way to match and mark strings for colorful terminal outputs! | 2019-08-28 15:44:08 | 2021-05-30 14:15:46 |
| [go-ataman](https://github.com/workanator/go-ataman) | 9 | 1 | 0 | Another Text Attribute Manupulator | 2017-05-17 19:04:57 | 2020-12-23 05:34:36 |
| [table](https://github.com/tomlazar/table) | 8 | 0 | 0 | pretty colorfull tables in go with less effort | 2020-09-22 05:42:34 | 2021-05-19 20:56:16 |
</details>

### Command Line - Standard CLI
Libraries for building standard or basic Command Line applications.

<sup>*Last Update: 2021-06-22 14:35:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cobra](https://cobra.dev) | 21,968 | 1,850 | 321 | A Commander for modern Go CLI interactions | 2013-09-03 20:40:26 | 2021-06-22 07:03:10 |
| [cli](https://github.com/urfave/cli) | 15,983 | 1,400 | 82 | A simple, fast, and fun package for building command line apps in Go | 2013-07-13 19:32:06 | 2021-06-22 06:12:12 |
| [kingpin](https://github.com/alecthomas/kingpin) | 3,064 | 231 | 25 | CONTRIBUTIONS ONLY: A Go (golang) command line and flag parser | 2014-05-14 20:09:04 | 2021-06-21 22:43:40 |
| [dnote](https://www.getdnote.com) | 2,099 | 85 | 58 | A simple command line notebook for programmers | 2017-03-30 23:07:25 | 2021-06-19 03:28:43 |
| [go-flags](http://godoc.org/github.com/jessevdk/go-flags) | 1,963 | 239 | 29 | go command line option parser | 2012-08-31 13:57:58 | 2021-06-21 22:52:12 |
| [pflag](https://ops.city) | 1,470 | 245 | 104 | Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. | 2013-08-30 14:53:31 | 2021-06-22 06:46:20 |
| [cli](https://github.com/mitchellh/cli) | 1,338 | 104 | 7 | A Go library for implementing command-line interfaces. | 2013-11-03 06:47:54 | 2021-06-19 18:22:24 |
| [go-arg](https://github.com/alexflint/go-arg) | 1,177 | 66 | 6 | Struct-based argument parsing in Go | 2015-11-01 01:30:06 | 2021-06-18 16:06:46 |
| [liner](https://github.com/peterh/liner) | 789 | 103 | 12 | Pure Go line editor with history, inspired by linenoise | 2012-08-15 16:34:55 | 2021-06-15 03:49:53 |
| [complete](https://github.com/posener/complete) | 776 | 60 | 20 | bash completion written in go + bash completion for go command | 2017-05-05 21:34:07 | 2021-06-18 12:06:55 |
| [mow.cli](https://github.com/jawher/mow.cli) | 740 | 50 | 26 | A versatile library for building CLI applications in Go | 2014-12-18 19:34:20 | 2021-06-21 08:31:24 |
| [flaggy](https://github.com/integrii/flaggy) | 718 | 22 | 12 | Idiomatic Go input parsing with subcommands, positional values, and flags at any position. No required project or package layout and no external dependencies. | 2018-03-05 05:55:05 | 2021-06-14 08:15:02 |
| [ops](https://ops.city) | 602 | 68 | 129 | ops - build and run nanos unikernels | 2018-09-10 17:57:47 | 2021-06-21 19:17:24 |
| [cli](https://github.com/mkideal/cli) | 578 | 38 | 3 | CLI - A package for building command line app with go | 2016-02-26 16:45:29 | 2021-06-11 08:51:31 |
| [argparse](https://github.com/akamensky/argparse) | 335 | 42 | 10 | Argparse for golang. Just because `flag` sucks | 2017-11-24 06:42:20 | 2021-06-19 15:20:15 |
| [climax](http://git.io/climax) | 186 | 17 | 6 | Climax is an alternative CLI with the human face | 2015-11-03 21:04:57 | 2021-06-04 17:16:47 |
| [hiboot](https://hiboot.netlify.app/) | 160 | 28 | 3 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2021-06-12 12:38:29 |
| [commandeer](https://github.com/jaffee/commandeer) | 147 | 15 | 3 | Automatically sets up command line flags based on struct fields and tags. | 2017-10-12 02:51:05 | 2021-06-16 20:17:11 |
| [wmenu](https://github.com/dixonwille/wmenu) | 137 | 18 | 1 | An easy to use menu structure for cli applications that prompts users to make choices. | 2016-04-20 13:09:44 | 2021-06-11 10:14:48 |
| [sflags](https://ops.city) | 129 | 24 | 9 | Generate flags by parsing structures | 2016-12-04 14:49:27 | 2021-06-09 06:45:51 |
| [flag](https://github.com/cosiner/flag) | 115 | 6 | 1 | Flag is a simple but powerful command line option parsing library for Go support infinite level subcommand | 2016-10-05 16:49:41 | 2021-05-29 20:22:59 |
| [clif](https://github.com/ukautz/clif) | 108 | 11 | 3 | Another CLI framework for Go. It works on my machine. | 2015-05-30 18:30:08 | 2021-05-25 03:55:56 |
| [job](https://github.com/liujianping/job) | 101 | 7 | 1 | JOB, make your short-term command as a long-term job. 将命令行规划成任务的工具 | 2019-04-09 11:14:51 | 2021-06-09 03:31:02 |
| [cli](https://github.com/teris-io/cli) | 91 | 9 | 2 | Simple and complete API for building command line applications in Go | 2017-05-24 23:07:07 | 2021-05-26 03:10:37 |
| [env](https://github.com/codingconcepts/env) | 78 | 5 | 1 | Tag-based environment configuration for structs | 2017-06-14 20:01:55 | 2021-06-15 05:47:16 |
| [cmdr](https://hedzr.github.io/cmdr-docs/) | 75 | 7 | 1 | Golang library with POSIX-compliant command-line UI (CLI) and Hierarchical-configuration. Better substitute for stdlib flag. | 2019-05-15 09:58:02 | 2021-06-17 06:44:46 |
| [clir](https://github.com/leaanthony/clir) | 66 | 9 | 5 | A Simple and Clear CLI library. Dependency free. | 2019-11-18 19:52:00 | 2021-05-27 07:23:31 |
| [gocmd](https://github.com/devfacet/gocmd) | 52 | 6 | 1 | A Go library for building command line applications. | 2018-01-08 04:52:02 | 2021-05-29 20:35:23 |
| [wlog](https://github.com/dixonwille/wlog) | 51 | 6 | 0 | A simple logging interface that supports cross-platform color and concurrency. | 2016-04-13 16:47:40 | 2021-04-25 04:09:14 |
| [strumt](https://github.com/antham/strumt) | 45 | 4 | 0 | Strumt is a library to create prompt chain | 2017-06-19 19:33:16 | 2021-05-21 04:31:26 |
| [flagvar](https://pkg.go.dev/github.com/sgreben/flagvar?tab=doc) | 36 | 1 | 0 | A collection of CLI argument types for the Go `flag` package. | 2018-05-18 18:45:16 | 2020-09-18 12:21:54 |
| [go-getoptions](https://github.com/DavidGamba/go-getoptions) | 35 | 5 | 0 | Fully featured Go (golang) command line option parser with built-in auto-completion support. | 2015-12-18 02:21:14 | 2021-06-08 17:56:32 |
| [cmd](https://github.com/posener/cmd) | 33 | 1 | 0 | The standard library flag package with its missing features | 2019-10-29 00:32:11 | 2021-06-11 12:03:34 |
| [argv](https://github.com/cosiner/argv) | 30 | 6 | 1 | Argparse for golang. Just because `flag` sucks | 2017-01-22 10:37:21 | 2021-03-30 02:32:52 |
| [go-commander](http://yitsushi.github.io/go-commander/) | 22 | 4 | 1 | Go library to simplify CLI workflow | 2016-10-10 10:09:41 | 2021-05-01 01:52:24 |
| [sand](https://ops.city) | 15 | 1 | 0 | Package for creating interpreters | 2018-11-18 22:44:41 | 2021-05-10 22:00:31 |
| [ts](https://github.com/liujianping/ts) | 13 | 1 | 0 | timestamp convert & compare tool. 时间戳转换与对比工具 | 2019-06-25 10:21:13 | 2021-02-13 09:18:33 |
</details>

### Configuration
Libraries for configuration parsing.

<sup>*Last Update: 2021-06-29 16:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [viper](https://github.com/spf13/viper) | 16,021 | 1,404 | 452 | Go configuration with fangs | 2014-04-02 14:33:33 | 2021-06-29 09:10:29 |
| [godotenv](http://godoc.org/github.com/joho/godotenv) | 3,935 | 235 | 50 | A Go port of Ruby's dotenv library (Loads environment variables from `.env`.) | 2013-07-30 07:45:19 | 2021-06-29 08:38:36 |
| [envconfig](https://josh.blog/2017/04/go-configure) | 3,652 | 299 | 46 | Golang library for managing configuration data from environment variables | 2013-11-06 17:01:55 | 2021-06-29 01:55:15 |
| [ini](https://ini.unknwon.io) | 2,528 | 313 | 33 | Package ini provides INI file read and write functionality in Go. | 2014-12-18 07:36:37 | 2021-06-29 03:12:18 |
| [env](http://carlosbecker.com/posts/env-structs-golang) | 1,922 | 147 | 3 | Simple lib to parse environment variables to structs | 2015-07-28 02:14:37 | 2021-06-28 19:47:22 |
| [konfig](https://github.com/lalamove/konfig) | 609 | 40 | 4 | Composable, observable and performant config handling for Go for the distributed processing era | 2019-01-18 17:03:03 | 2021-06-27 08:36:03 |
| [koanf](https://github.com/knadh/koanf) | 521 | 46 | 7 | Light weight, extensible, configuration management library for Go. Built in support for JSON, TOML, YAML, env, command line, file, S3 etc. Alternative to viper. | 2019-06-18 06:34:05 | 2021-06-28 19:24:40 |
| [confita](https://github.com/heetch/confita) | 397 | 38 | 19 | Load configuration in cascade from multiple backends into a struct | 2017-12-21 10:49:18 | 2021-06-28 19:55:55 |
| [cleanenv](https://github.com/ilyakaznacheev/cleanenv) | 298 | 33 | 19 | ✨Clean and minimalistic environment configuration reader for Golang | 2019-07-12 15:28:52 | 2021-06-25 00:47:26 |
| [config](https://github.com/JeremyLoy/config) | 257 | 10 | 1 | 12 factor configuration as a typesafe struct in as little as two function calls | 2019-04-02 13:41:22 | 2021-06-29 07:28:18 |
| [store](https://github.com/tucnak/store) | 253 | 16 | 2 | A dead simple configuration manager for Go applications | 2015-10-03 19:17:28 | 2021-05-25 06:32:04 |
| [config](https://pkg.go.dev/github.com/gookit/config/v2) | 251 | 32 | 1 | 📝 Go config manage(load,get,set). support JSON, YAML, TOML, INI, HCL, ENV and Flags. Multi file load, data override merge, parse ENV var. Go应用配置加载管理，支持多种格式，多文件加载，远程文件加载，支持数据合并，解析环境变量名 | 2018-07-07 08:11:39 | 2021-06-28 14:08:16 |
| [config](https://github.com/olebedev/config) | 240 | 43 | 4 | JSON or YAML configuration wrapper with convenient access methods. | 2014-04-21 15:09:39 | 2021-06-11 14:40:05 |
| [hjson-go](https://hjson.github.io/) | 239 | 26 | 3 | Hjson for Go | 2016-08-05 22:59:18 | 2021-06-24 14:59:02 |
| [aconfig](https://github.com/cristalhq/aconfig) | 228 | 14 | 7 | Simple, useful and opinionated config loader. | 2020-06-26 19:43:20 | 2021-06-24 07:31:17 |
| [config](https://josh.blog/2017/04/go-configure) | 205 | 11 | 1 | 🛠 A configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP | 2017-04-02 18:37:05 | 2021-05-01 01:16:15 |
| [envconfig](https://godoc.org/github.com/tomazk/envcfg) | 201 | 21 | 0 | Small library to read your configuration from environment variables | 2015-04-21 23:37:17 | 2021-05-19 07:28:37 |
| [fig](https://github.com/kkyr/fig) | 155 | 9 | 2 | A minimalist Go configuration library | 2020-01-16 18:43:19 | 2021-06-18 12:26:22 |
| [gcfg](https://gopkg.in/gcfg.v1) | 150 | 47 | 8 | read INI-style configuration files into Go structs; supports user-defined types and subsections | 2015-08-17 14:40:55 | 2021-05-18 08:34:01 |
| [goconfig](https://pkg.go.dev/github.com/gosidekick/goconfig?tab=doc) | 148 | 20 | 7 | goconfig uses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. | 2016-12-18 11:22:41 | 2021-06-03 22:49:50 |
| [config](https://github.com/golobby/config) | 146 | 11 | 1 | A lightweight yet powerful config package for Go projects | 2019-10-15 22:51:19 | 2021-06-21 05:49:33 |
| [xdg](https://pkg.go.dev/github.com/adrg/xdg) | 107 | 6 | 0 | Go implementation of the XDG Base Directory Specification and XDG user directories | 2014-08-22 08:23:40 | 2021-06-23 07:11:17 |
| [envh](https://github.com/antham/envh) | 96 | 1 | 0 | Go helpers to manage environment variables | 2017-01-12 11:25:48 | 2021-05-21 04:31:29 |
| [envcfg](https://godoc.org/github.com/tomazk/envcfg) | 93 | 6 | 0 | Un-marshaling environment variables to Go structs | 2014-11-29 11:43:53 | 2020-12-06 17:12:56 |
| [onion](https://github.com/goraz/onion) | 91 | 8 | 8 | Layer based configuration for golang | 2015-07-22 14:28:21 | 2020-12-13 20:13:08 |
| [harvester](https://github.com/beatlabs/harvester) | 84 | 18 | 3 | Harvest configuration, watch and notify subscriber | 2019-04-09 07:37:19 | 2021-06-17 11:28:48 |
| [configuro](https://medium.com/better-programming/designing-cloud-native-configuration-framework-eefb0b3793cb) | 74 | 9 | 0 | An opinionated configuration loading framework for Containerized and Cloud-Native applications. | 2020-04-09 22:10:34 | 2021-06-10 12:57:35 |
| [xdg](https://github.com/OpenPeeDeeP/xdg) | 63 | 5 | 1 | A cross platform package that follows the XDG Standard | 2017-07-20 15:58:29 | 2021-06-18 00:46:25 |
| [gofigure](https://github.com/ian-kent/gofigure) | 59 | 9 | 1 | Go configuration made easy! | 2014-11-25 00:12:40 | 2020-08-31 14:19:02 |
| [configure](https://github.com/paked/configure) | 54 | 9 | 2 | Configure is a Go package that gives you easy configuration of your project through redundancy | 2015-06-14 07:46:56 | 2021-05-27 13:28:12 |
| [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) | 40 | 7 | 0 | Go package that interfaces with AWS System Manager | 2019-01-24 09:01:19 | 2021-06-17 21:44:36 |
| [configuration](https://github.com/BoRuDar/configuration) | 39 | 6 | 0 | Library for setting values to structs' fields from env, flags, files or default tag | 2019-11-27 17:58:49 | 2021-06-03 17:33:29 |
| [ingo](https://github.com/schachmat/ingo) | 35 | 7 | 0 | persistent storage for flags in go | 2016-02-07 22:57:40 | 2021-06-15 10:31:14 |
| [gone](https://github.com/One-com/gone) | 35 | 5 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2021-05-24 14:22:10 |
| [go-up](https://github.com/ufoscout/go-up) | 32 | 5 | 1 | go-up! A simple configuration library with recursive placeholders resolution and no magic. | 2018-02-18 09:50:00 | 2021-06-11 20:28:29 |
| [mini](https://github.com/sasbury/mini) | 28 | 6 | 1 | A golang package for parsing ini-style configuration files | 2015-04-29 23:52:36 | 2020-12-08 09:30:24 |
| [hocon](https://hjson.github.io/) | 26 | 5 | 1 | go implementation of lightbend's HOCON configuration library https://github.com/lightbend/config | 2020-03-01 18:20:12 | 2021-06-19 19:17:48 |
| [genv](https://github.com/sakirsensoy/genv) | 23 | 1 | 0 | Genv is a library for Go (golang) that makes it easy to read and use environment variables in your projects. It also allows environment variables to be loaded from the .env file. | 2019-07-15 10:25:57 | 2021-06-03 18:30:51 |
| [conflate](https://github.com/the4thamigo-uk/conflate) | 19 | 2 | 0 | Library providing routines to merge and validate JSON, YAML and/or TOML files | 2018-02-01 19:06:15 | 2021-05-20 08:54:46 |
| [go-ssm-config](http://godoc.org/github.com/ianlopshire/go-ssm-config) | 10 | 5 | 4 | Go utility for loading configuration parameters from AWS SSM (Parameter Store) | 2019-12-02 18:47:38 | 2021-05-20 10:56:58 |
| [envconf](https://godoc.org/github.com/tomazk/envcfg) | 10 | 3 | 0 | Configure Go applications from the environment | 2014-10-26 12:12:26 | 2020-08-31 14:12:53 |
| [go-ini](https://github.com/subpop/go-ini) | 5 | 1 | 1 | automatic mirror of https://git.sr.ht/~spc/go-ini | 2019-09-11 18:38:20 | 2021-06-09 06:01:09 |
| [swap](https://github.com/oblq/swap) | 4 | 0 | 0 | Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). | 2020-04-12 23:28:19 | 2021-02-25 07:52:33 |
| [typenv](https://github.com/diegomarangoni/typenv) | 4 | 0 | 0 | Go minimalist typed environment variables library | 2020-06-30 18:26:09 | 2021-02-08 22:09:22 |
| [env](https://github.com/nasermirzaei89/env) | 4 | 0 | 0 | Golang Get Environment Variables Package | 2019-07-24 06:37:13 | 2021-05-06 07:08:59 |
| [gonfig](https://github.com/miladabc/gonfig) | 2 | 0 | 0 | Tag based configuration loader from different providers | 2021-01-21 13:44:44 | 2021-03-07 21:29:39 |
</details>

### Continuous Integration
Tools for help with continuous integration.

<sup>*Last Update: 2021-07-13 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [drone](https://drone.io) | 23,511 | 2,295 | 52 | Drone is a Container-Native, Continuous Delivery Platform | 2014-02-07 07:54:44 | 2021-07-13 02:20:06 |
| [cds](https://ovh.github.io/cds/) | 3,471 | 327 | 165 | Enterprise-Grade Continuous Delivery & DevOps Automation Open Source Platform | 2016-10-11 08:28:23 | 2021-07-11 22:29:24 |
| [goveralls](https://github.com/mattn/goveralls) | 692 | 129 | 11 | A tool for testing, building, signing, and publishing binaries. | 2013-04-17 10:58:40 | 2021-07-05 20:47:57 |
| [overalls](https://github.com/go-playground/overalls) | 106 | 26 | 3 | :jeans:Multi-Package go project coverprofile for tools like goveralls | 2015-07-30 11:30:11 | 2021-03-06 16:05:33 |
| [duci](https://github.com/duck8823/duci) | 66 | 3 | 5 | The simple ci server  | 2018-04-01 01:51:02 | 2021-07-01 15:48:22 |
| [gomason](https://github.com/nikogura/gomason) | 50 | 6 | 2 | A tool for testing, building, signing, and publishing binaries. | 2017-11-18 00:59:11 | 2021-07-08 22:23:01 |
| [roveralls](https://github.com/lawrencewoodman/roveralls) | 14 | 4 | 0 | A Go recursive coverage testing tool | 2016-06-26 07:45:32 | 2020-08-20 00:07:58 |
</details>

### Data Structures
Generic datastructures and algorithms in Go.

<sup>*Last Update: 2021-06-29 14:34:28*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gods](https://github.com/emirpasic/gods) | 10,127 | 1,194 | 50 | GoDS (Go Data Structures). Containers (Sets, Lists, Stacks, Maps, Trees), Sets (HashSet, TreeSet, LinkedHashSet), Lists (ArrayList, SinglyLinkedList, DoublyLinkedList), Stacks (LinkedListStack, ArrayStack), Maps (HashMap, TreeMap, HashBidiMap, TreeBidiMap, LinkedHashMap), Trees (RedBlackTree, AVLTree, BTree, BinaryHeap), Comparators, Iterators, Enumerables, Sort, JSON | 2015-03-04 14:19:52 | 2021-06-28 21:18:20 |
| [go-datastructures](https://github.com/Workiva/go-datastructures) | 6,078 | 710 | 25 | A collection of useful, performant, and threadsafe Go datastructures. | 2014-10-29 13:55:17 | 2021-06-28 06:45:39 |
| [golang-set](https://github.com/deckarep/golang-set) | 1,981 | 170 | 12 | A simple set type for the Go language. Trusted by Docker, 1Password, Ethereum and Hashicorp. | 2013-07-03 21:52:01 | 2021-06-25 17:21:22 |
| [gota](https://github.com/go-gota/gota) | 1,634 | 173 | 46 | Gota: DataFrames and data wrangling in Go (Golang) | 2016-02-06 17:23:25 | 2021-06-28 22:21:21 |
| [BoomFilters](https://github.com/tylertreat/BoomFilters) | 1,357 | 95 | 9 | Probabilistic data structures for processing continuous, unbounded streams. | 2015-02-06 02:01:26 | 2021-06-25 09:37:00 |
| [roaring](http://roaringbitmap.org/) | 1,235 | 136 | 61 | Roaring bitmaps in Go (golang) | 2014-07-10 20:14:34 | 2021-06-28 21:42:01 |
| [bloom](https://github.com/bits-and-blooms/bloom) | 1,223 | 167 | 3 | Go package implementing Bloom filters | 2011-05-21 14:18:41 | 2021-06-28 09:51:36 |
| [gocache](https://vincent.composieux.fr/article/i-wrote-gocache-a-complete-and-extensible-go-cache-library/) | 879 | 78 | 8 | ☔️ A complete Go cache library that brings you multiple ways of managing your caches | 2019-10-05 08:13:54 | 2021-06-29 01:25:36 |
| [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) | 800 | 64 | 10 | Cuckoo Filter: Practically Better Than Bloom | 2015-06-28 23:22:09 | 2021-06-28 09:51:45 |
| [hyperloglog](https://github.com/axiomhq/hyperloglog) | 749 | 51 | 3 | HyperLogLog with lots of sugar (Sparse, LogLog-Beta bias correction and TailCut space reduction) | 2017-06-18 11:18:12 | 2021-06-23 20:31:34 |
| [bitset](https://github.com/bits-and-blooms/bitset) | 697 | 120 | 3 | Go package implementing bitsets | 2011-05-11 03:33:44 | 2021-06-28 01:22:38 |
| [trie](https://github.com/derekparker/trie) | 538 | 87 | 10 | Data structure and relevant algorithms for extremely fast prefix/fuzzy string searching. | 2014-03-06 22:01:49 | 2021-06-23 13:35:06 |
| [algorithms](https://github.com/shady831213/algorithms) | 531 | 89 | 0 | CLRS study. Codes are written with golang. | 2018-01-31 09:27:56 | 2021-06-27 14:16:50 |
| [go-geoindex](https://github.com/hailocab/go-geoindex) | 330 | 40 | 3 | Go native library for fast point tracking and K-Nearest queries | 2015-01-22 12:26:17 | 2021-06-21 14:56:53 |
| [gostl](https://github.com/liyue201/gostl) | 284 | 59 | 1 | Data structure and algorithm library for go, designed to provide functions similar to C++ STL | 2019-10-12 01:10:24 | 2021-06-22 16:39:15 |
| [go-edlib](https://github.com/hbollon/go-edlib) | 268 | 11 | 0 | Golang string comparison and edit distance algorithms library, featuring : Levenshtein, LCS, Hamming, Damerau levenshtein (OSA and Adjacent transpositions algorithms), Jaro-Winkler, Cosine, etc... | 2020-08-18 09:30:59 | 2021-06-27 15:29:53 |
| [merkletree](https://github.com/cbergoon/merkletree) | 251 | 70 | 6 | A Merkle Tree implementation written in Go. | 2017-04-12 02:50:11 | 2021-06-22 14:40:31 |
| [ttlcache](https://github.com/ReneKroon/ttlcache) | 235 | 40 | 0 | An in-memory string-interface{} map with various expiration options for golang | 2014-12-13 01:55:40 | 2021-06-19 04:50:35 |
| [hilbert](https://github.com/google/hilbert) | 233 | 33 | 2 | Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. | 2015-08-06 15:50:00 | 2021-05-28 14:00:08 |
| [goskiplist](https://github.com/ryszard/goskiplist) | 226 | 55 | 6 | A skip list implementation in Go | 2012-05-09 05:44:59 | 2021-06-18 11:09:38 |
| [deque](https://github.com/gammazero/deque) | 222 | 18 | 0 | Fast ring-buffer deque (double-ended queue) | 2018-04-24 02:57:55 | 2021-06-23 09:17:39 |
| [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | 180 | 30 | 0 | Adaptive Radix Trees implemented in Go | 2016-04-01 01:40:40 | 2021-06-20 06:44:20 |
| [binpacker](https://github.com/zhuangsirui/binpacker) | 164 | 27 | 2 | A binary stream packer and unpacker | 2016-02-02 10:06:11 | 2021-06-24 06:57:38 |
| [skiplist](https://github.com/MauriceGit/skiplist) | 159 | 21 | 2 | A Go library for an efficient implementation of a skip list: https://godoc.org/github.com/MauriceGit/skiplist | 2018-06-23 16:01:51 | 2021-06-16 20:58:09 |
| [cuckoo-filter](https://github.com/linvon/cuckoo-filter) | 156 | 11 | 0 | Cuckoo Filter go implement, better than Bloom Filter, configurable and space optimized  布谷鸟过滤器的Go实现，优于布隆过滤器，可以定制化过滤器参数，并进行了空间优化 | 2021-02-19 12:27:43 | 2021-06-28 03:01:36 |
| [levenshtein](https://github.com/agnivade/levenshtein) | 144 | 11 | 1 | Go implementation to calculate Levenshtein Distance. | 2014-07-30 14:03:55 | 2021-06-26 03:54:57 |
| [bloom](http://zhen.org/blog/benchmarking-bloom-filters-and-hash-functions-in-go/) | 144 | 15 | 1 | Bloom filters implemented in Go. | 2013-09-03 02:27:35 | 2021-06-04 02:57:34 |
| [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) | 139 | 9 | 1 | Go concurrent-safe, goroutine-safe, thread-safe queue | 2019-01-10 21:21:23 | 2021-06-22 16:24:57 |
| [iter](https://github.com/disksing/iter) | 136 | 7 | 1 | Go implementation of C++ STL iterators and algorithms. | 2019-10-20 09:29:49 | 2021-06-09 05:10:43 |
| [ring](https://pkg.go.dev/github.com/tannerryan/ring) | 119 | 13 | 2 | Package ring provides a high performance and thread safe Go implementation of a bloom filter. | 2019-01-27 04:02:20 | 2021-04-17 13:36:17 |
| [go-rquad](https://github.com/arl/go-rquad) | 115 | 3 | 0 | :pushpin: State of the art point location and neighbour finding algorithms for region quadtrees, in Go | 2016-09-12 21:46:37 | 2021-05-29 12:21:29 |
| [encoding](http://zhen.org/blog/benchmarking-integer-compression-in-go/) | 106 | 13 | 1 | Integer Compression Libraries for Go | 2013-09-20 19:29:57 | 2021-06-21 05:03:21 |
| [bit](https://yourbasic.org/algorithms/your-basic-int/#simple-sets) | 103 | 16 | 0 | Bitset data structure | 2017-05-03 19:05:35 | 2021-06-03 06:27:17 |
| [remember-go](https://github.com/rocketlaunchr/remember-go) | 99 | 6 | 1 | Cache Slow Database Queries | 2019-04-04 20:24:25 | 2021-06-22 06:09:00 |
| [conjungo](https://github.com/InVisionApp/conjungo) | 94 | 12 | 10 | A small flexible merge library in go | 2016-12-29 23:50:38 | 2021-06-02 07:01:48 |
| [skiplist](https://github.com/gansidui/skiplist) | 72 | 18 | 1 | skiplist for golang | 2014-11-18 16:29:53 | 2021-03-05 09:16:43 |
| [go-mcache](https://pkg.go.dev/github.com/OrlovEvgeny/go-mcache?tab=doc) | 67 | 9 | 0 | Fast in-memory key:value store/cache with TTL | 2018-04-14 23:31:21 | 2021-06-03 03:53:28 |
| [bloom](https://yourbasic.org/algorithms/bloom-filter/) | 61 | 8 | 0 | Probabilistic set data structure | 2017-05-06 19:57:47 | 2021-06-04 02:57:57 |
| [levenshtein](https://github.com/agext/levenshtein) | 57 | 5 | 0 | Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. | 2016-04-08 00:14:31 | 2021-02-24 01:05:36 |
| [count-min-log](https://github.com/seiflotfy/count-min-log) | 53 | 3 | 0 | Go implementation of Count-Min-Log | 2015-08-16 22:31:36 | 2021-03-31 12:17:31 |
| [crunch](https://github.com/superwhiskers/crunch) | 44 | 6 | 0 | take bytes out of things easily ✨🍪 | 2019-02-27 03:56:52 | 2021-06-21 05:46:50 |
| [nan](https://github.com/kak-tus/nan) | 36 | 5 | 0 | Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers | 2020-05-05 20:20:54 | 2021-04-08 12:15:20 |
| [concurrent-writer](https://github.com/free/concurrent-writer) | 34 | 5 | 0 | Highly concurrent drop-in replacement for bufio.Writer | 2017-09-18 15:29:59 | 2021-06-27 10:06:17 |
| [hide](https://godoc.org/github.com/yaa110/goterator) | 34 | 3 | 0 | ID type with marshalling to/from hash to prevent sending IDs to clients. | 2019-01-16 13:54:17 | 2021-06-20 09:54:05 |
| [goset](https://github.com/zoumo/goset) | 33 | 11 | 0 | Set is a useful collection but there is no built-in implementation in Go lang. | 2017-08-25 09:21:30 | 2021-06-02 19:49:07 |
| [pipeline](https://godoc.org/github.com/hyfather/pipeline) | 30 | 3 | 0 | Pipelines using goroutines | 2018-04-25 00:11:36 | 2021-05-27 16:36:48 |
| [deque](https://github.com/edwingeng/deque) | 28 | 1 | 0 | A highly optimized double-ended queue | 2019-02-01 03:32:28 | 2021-06-04 11:23:02 |
| [typ](https://github.com/gurukami/typ) | 27 | 1 | 0 | Null Types, Safe primitive type conversion and fetching value from complex structures. | 2019-03-03 05:34:23 | 2021-06-05 05:15:58 |
| [timedmap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 26 | 3 | 0 | A thread safe map which has expiring key-value pairs. | 2019-01-30 12:55:37 | 2021-06-13 03:39:14 |
| [null](https://github.com/emvi/null) | 19 | 1 | 0 | Nullable Go types that can be marshalled/unmarshalled to/from JSON. | 2018-07-04 21:18:45 | 2021-05-21 08:13:05 |
| [go-ef](https://github.com/amallia/go-ef) | 18 | 5 | 0 | A Go implementation of the Elias-Fano encoding | 2017-09-22 01:47:16 | 2021-05-05 17:59:40 |
| [dict](https://github.com/srfrog/dict) | 18 | 1 | 0 | Python-like dictionaries for Go | 2019-04-23 02:04:25 | 2020-12-08 21:19:15 |
| [mspm](https://github.com/BlackRabbitt/mspm) | 15 | 1 | 0 | Multi-String Pattern Matching Algorithm Using TrieHashNode | 2018-05-17 18:59:44 | 2021-02-13 15:48:41 |
| [cmap](https://github.com/lrita/cmap) | 15 | 0 | 0 | a thread-safe concurrent map for go | 2019-11-26 03:54:59 | 2021-06-04 02:59:28 |
| [set](https://github.com/StudioSol/set) | 14 | 6 | 1 | A simple Set data structure implementation in Go (Golang) using LinkedHashMap. | 2018-07-20 21:53:37 | 2021-06-01 08:48:00 |
| [ptrie](https://godoc.org/github.com/hyfather/pipeline) | 14 | 3 | 0 | A prefix tree implementation in go  | 2019-05-20 14:13:05 | 2021-04-23 12:43:50 |
| [treap](https://pkg.go.dev/github.com/zekroTJA/timedmap) | 11 | 2 | 0 | golang persistent immutable treap sorted sets | 2018-09-16 01:38:03 | 2021-04-28 09:21:17 |
| [parapipe](https://github.com/nazar256/parapipe) | 10 | 0 | 0 | Paralleling pipeline | 2021-04-09 06:49:56 | 2021-06-21 04:46:05 |
| [gofal](https://github.com/xxjwxc/gofal) | 9 | 0 | 0 | fractional api base on golang . golang math tools fractional molecular denominator 分数计算 分子 分母 运算 | 2019-08-05 07:37:55 | 2020-10-08 14:54:15 |
| [parsefields](https://github.com/MonaxGT/parsefields) | 6 | 0 | 0 | Tools for parse JSON-like logs for collecting unique fields and events | 2019-04-12 22:15:10 | 2021-05-15 04:37:38 |
| [dsu](https://github.com/ihebu/dsu) | 5 | 0 | 0 | Disjoint Set data structure implementation in Go | 2021-04-27 16:35:38 | 2021-06-21 04:48:16 |
| [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) | 5 | 0 | 0 | Ordered-concurrently a library for parallel processing with ordered output in Go. Process work concurrently / in parallel and returns output in a channel in the order of input. It is useful in concurrently / parallelly processing items in a queue, and get output in the order provided by the queue. | 2021-02-28 17:56:05 | 2021-06-09 09:42:22 |
| [bloomfilter](https://github.com/OldPanda/bloomfilter) | 4 | 0 | 0 | Yet another Bloomfilter implementation in Go, compatible with Java's Guava library | 2021-01-01 01:28:04 | 2021-06-20 07:24:23 |
| [goterator](https://godoc.org/github.com/yaa110/goterator) | 3 | 0 | 0 | Lazy iterator implementation for Golang | 2020-08-12 19:47:57 | 2020-12-02 04:17:41 |
| [slices](https://github.com/srfrog/slices) | 3 | 0 | 0 | Functions that operate on slices. Similar to functions from package strings or package bytes that have been adapted to work with slices. | 2020-07-02 23:17:34 | 2021-05-08 03:37:15 |
</details>

### Database - Database schema migration


<sup>*Last Update: 2021-06-22 09:01:53*</sup>
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

### Database - Database tools


<sup>*Last Update: 2021-06-25 12:39:54*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [vitess](http://vitess.io) | 12,090 | 1,542 | 682 | Vitess is a database clustering system for horizontal scaling of MySQL. | 2013-06-27 21:20:28 | 2021-06-24 17:01:25 |
| [pgweb](http://sosedoff.github.io/pgweb) | 6,907 | 545 | 53 | Cross-platform client for PostgreSQL databases | 2014-10-09 01:41:32 | 2021-06-24 07:07:07 |
| [kingshard](https://github.com/flike/kingshard) | 5,711 | 1,143 | 153 | A high-performance MySQL proxy | 2015-07-04 02:22:32 | 2021-06-23 02:56:52 |
| [orchestrator](https://github.com/openark/orchestrator) | 4,117 | 703 | 341 | MySQL replication topology management and HA | 2016-11-30 13:44:24 | 2021-06-25 02:51:57 |
| [go-mysql-elasticsearch](https://github.com/go-mysql-org/go-mysql-elasticsearch) | 3,503 | 683 | 197 | Sync MySQL data into elasticsearch  | 2015-01-15 09:54:18 | 2021-06-24 11:39:22 |
| [go-mysql](https://github.com/go-mysql-org/go-mysql) | 3,107 | 700 | 146 | a powerful mysql toolset with Go | 2014-02-21 01:56:45 | 2021-06-24 17:47:50 |
| [prest](https://prestd.com) | 2,698 | 194 | 138 | pREST (PostgreSQL REST), simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new | 2016-11-22 05:17:05 | 2021-06-25 01:29:56 |
| [chproxy](https://github.com/Vertamedia/chproxy) | 636 | 136 | 36 | ClickHouse http proxy and load balancer | 2017-09-18 13:09:23 | 2021-06-25 03:36:40 |
| [pg_timetable](https://www.cybertec-postgresql.com/en/products/pg_timetable/) | 409 | 24 | 2 | pg_timetable: Advanced scheduling for PostgreSQL | 2018-12-19 10:19:51 | 2021-06-24 16:55:54 |
| [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) | 286 | 60 | 14 | Collects many small inserts to ClickHouse and send in big inserts | 2017-04-29 10:38:41 | 2021-06-25 03:35:34 |
| [myreplication](https://github.com/2tvenom/myreplication) | 173 | 45 | 5 | Golang MySql binary log replication listener | 2015-02-04 20:59:49 | 2021-06-01 07:46:11 |
| [octillery](https://github.com/blastrain/octillery) | 140 | 18 | 6 | Go package for sharding databases ( Supports every ORM or raw SQL ) | 2018-11-26 10:39:35 | 2021-06-24 02:16:54 |
| [dbbench](https://github.com/sj14/dbbench) | 59 | 6 | 12 | 🏋️ dbbench is a simple database benchmarking tool which supports several databases and own scripts | 2018-11-24 13:21:18 | 2021-06-24 07:16:51 |
| [datagen](https://github.com/codingconcepts/datagen) | 33 | 4 | 0 | A fast data generator that's multi-table aware and supports multi-row DML. | 2019-04-18 19:58:01 | 2021-06-06 20:04:05 |
| [prep](http://sosedoff.github.io/pgweb) | 28 | 2 | 0 | Prep finds all SQL statements in a Go package and instruments db connection with prepared statements | 2017-12-11 23:47:38 | 2021-05-03 10:39:05 |
| [rwdb](https://prestd.com) | 12 | 0 | 0 | Database wrapper that manage read write connections | 2017-10-04 03:55:29 | 2021-05-29 03:21:37 |
</details>

### Database - Databases implemented in Go


<sup>*Last Update: 2021-06-23 15:12:40*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prometheus](https://prometheus.io/) | 37,287 | 6,066 | 471 | The Prometheus monitoring system and time series database. | 2012-11-24 11:14:12 | 2021-06-23 08:03:32 |
| [tidb](https://pingcap.com) | 28,188 | 4,438 | 1,928 | TiDB is an open source distributed HTAP database compatible with the MySQL protocol  | 2015-09-06 04:01:52 | 2021-06-23 08:05:28 |
| [influxdb](https://influxdata.com) | 21,653 | 2,977 | 1,378 | Scalable datastore for metrics, events, and real-time analytics | 2013-09-26 14:31:10 | 2021-06-23 07:49:47 |
| [cockroach](https://www.cockroachlabs.com) | 20,763 | 2,720 | 3,901 | CockroachDB - the open source, cloud-native distributed SQL database. | 2014-02-06 00:18:47 | 2021-06-23 06:25:20 |
| [dgraph](https://dgraph.io) | 16,237 | 1,178 | 85 | Native GraphQL Database with graph backend | 2015-08-25 07:15:56 | 2021-06-23 07:39:26 |
| [groupcache](https://github.com/golang/groupcache) | 10,369 | 1,165 | 35 | groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. | 2013-07-22 21:55:07 | 2021-06-23 07:37:15 |
| [badger](https://blog.dgraph.io/post/badger/) | 9,361 | 777 | 5 | Fast key-value DB in Go. | 2017-01-26 05:09:49 | 2021-06-22 23:20:20 |
| [rqlite](http://www.rqlite.com) | 8,529 | 440 | 52 | The lightweight, distributed relational database built on SQLite | 2014-08-23 04:31:18 | 2021-06-23 07:31:27 |
| [go-cache](https://patrickmn.com/projects/go-cache/) | 5,092 | 641 | 51 | An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. | 2012-01-02 13:07:13 | 2021-06-22 11:50:37 |
| [bigcache](http://allegro.tech/2016/03/writing-fast-cache-service-in-go.html) | 4,930 | 416 | 48 | Efficient cache for gigabytes of data written in Go. | 2016-03-23 07:18:52 | 2021-06-23 06:00:36 |
| [bbolt](https://go.etcd.io/bbolt) | 4,536 | 344 | 104 | An embedded key/value database for Go. | 2017-06-17 01:42:09 | 2021-06-22 13:54:41 |
| [VictoriaMetrics](https://victoriametrics.com/) | 4,439 | 384 | 276 | VictoriaMetrics: fast, cost-effective monitoring solution and time series database | 2018-09-30 09:58:01 | 2021-06-23 05:08:53 |
| [goleveldb](https://patrickmn.com/projects/go-cache/) | 4,370 | 656 | 70 | LevelDB key/value database in Go. | 2013-01-23 04:08:58 | 2021-06-23 07:02:06 |
| [ledisdb](https://ledisdb.io) | 3,657 | 416 | 0 | A high performance NoSQL Database Server powered by Go | 2014-04-30 00:43:09 | 2021-06-23 08:04:07 |
| [buntdb](https://github.com/tidwall/buntdb) | 3,292 | 233 | 3 | BuntDB is an embeddable, in-memory key/value database for Go with custom indexing and geospatial support | 2016-07-19 22:11:40 | 2021-06-22 12:53:19 |
| [immudb](https://codenotary.com/technologies/immudb) | 2,878 | 119 | 56 | immudb - world’s fastest immutable database | 2019-11-07 08:22:16 | 2021-06-21 16:51:13 |
| [tiedot](https://github.com/HouzuoGuo/tiedot) | 2,596 | 258 | 26 | A rudimentary implementation of a basic document (NoSQL) database in Go | 2013-05-26 10:03:49 | 2021-06-22 18:59:36 |
| [nutsdb](https://xujiajun.cn/nutsdb/) | 1,662 | 141 | 30 | A simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. | 2018-12-07 07:03:38 | 2021-06-21 15:40:46 |
| [gcache](https://github.com/bluele/gcache) | 1,570 | 191 | 16 | An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC | 2015-01-24 18:17:07 | 2021-06-23 05:14:12 |
| [cache2go](https://github.com/muesli/cache2go) | 1,561 | 437 | 23 | Concurrency-safe Go caching library with expiration capabilities and access counters | 2013-11-11 03:45:02 | 2021-06-23 07:36:53 |
| [CovenantSQL](https://developers.covenantsql.io) | 1,190 | 131 | 25 | A decentralized, trusted, high performance, SQL database with blockchain features | 2018-04-11 09:52:58 | 2021-06-21 10:39:47 |
| [fastcache](https://github.com/VictoriaMetrics/fastcache) | 1,165 | 88 | 24 | Fast thread-safe inmemory cache for big number of entries in Go. Minimizes GC overhead | 2018-11-22 22:50:13 | 2021-06-22 16:26:44 |
| [rosedb](https://space.bilibili.com/26194591) | 1,079 | 160 | 7 | 🚀A fast, stable and embedded k-v database in pure Golang, supports string, list, hash, set, sorted set. 一个 Go 语言实现的快速、稳定、内嵌的 k-v 数据库。 | 2020-12-06 07:02:48 | 2021-06-23 06:59:49 |
| [diskv](http://godoc.org/github.com/peterbourgon/diskv) | 1,066 | 89 | 7 | A disk-backed key-value store. | 2012-03-21 16:44:32 | 2021-06-20 10:22:43 |
| [eliasdb](https://github.com/krotik/eliasdb) | 814 | 41 | 10 | EliasDB a graph-based database. | 2016-08-13 13:53:28 | 2021-06-21 23:07:01 |
| [moss](https://github.com/couchbase/moss) | 811 | 48 | 46 | moss - a simple, fast, ordered, persistable, key-val storage library for golang | 2016-02-06 20:27:22 | 2021-06-15 00:42:21 |
| [pogreb](https://github.com/akrylysov/pogreb) | 797 | 56 | 8 | Embedded key-value store for read-heavy workloads written in Go | 2018-01-06 23:16:36 | 2021-06-22 15:35:01 |
| [databunker](https://databunker.org/) | 779 | 31 | 0 | Secure vault for customer records built to comply with GDPR | 2019-12-08 21:55:55 | 2021-06-23 07:44:17 |
| [bitcask](https://prologic.github.io/bitcask) | 759 | 70 | 10 | 🔑A high performance Key/Value store written in Go with a predictable read/write performance and high throughput. Uses a Bitcask on-disk layout (LSM+WAL) similar to Riak. | 2019-03-12 13:57:35 | 2021-06-22 16:39:37 |
| [levigo](https://github.com/jmhodges/levigo) | 396 | 81 | 3 | levigo is a Go wrapper for LevelDB | 2012-01-17 08:17:54 | 2021-05-28 05:09:18 |
| [pudge](https://github.com/recoilme/pudge) | 289 | 22 | 0 | Fast and simple key/value store written using Go's standard library | 2018-11-20 10:11:53 | 2021-06-21 00:10:48 |
| [vasto](https://github.com/chrislusf/vasto) | 214 | 23 | 4 | A distributed key-value store. On Disk. Able to grow or shrink without service interruption. | 2018-01-16 05:16:57 | 2021-06-18 01:58:08 |
| [kivik](https://github.com/go-kivik/kivik) | 210 | 29 | 10 | Kivik provides a common interface to CouchDB or CouchDB-like databases for Go and GopherJS. | 2017-02-09 14:14:54 | 2021-06-21 15:47:43 |
| [piladb](https://www.piladb.org) | 186 | 19 | 9 | Lightweight RESTful database engine based on stack data structures | 2015-09-08 23:12:22 | 2021-04-15 21:45:35 |
| [golang-scribble](https://github.com/nanobox-io/golang-scribble) | 127 | 13 | 1 | A tiny Golang JSON database | 2018-06-21 22:13:33 | 2021-06-21 15:45:21 |
| [slowpoke](https://github.com/recoilme/slowpoke) | 98 | 8 | 0 | Low-level key/value store in pure Go.  | 2018-02-19 09:22:37 | 2020-11-15 06:42:52 |
| [cache](https://github.com/akyoto/cache) | 89 | 8 | 0 | :handbag: Cache arbitrary data with an expiration time. | 2019-05-11 12:42:45 | 2021-06-20 22:33:37 |
| [bcache](https://github.com/iwanbk/bcache) | 66 | 10 | 3 | Eventually consistent distributed in-memory  cache Go library | 2018-12-26 15:45:16 | 2021-05-07 05:37:53 |
| [unitdb](https://github.com/unit-io/unitdb) | 63 | 7 | 0 | Fast specialized time-series database for IoT, real-time internet connected devices and AI analytics. | 2019-08-29 18:21:27 | 2021-05-21 16:10:59 |
| [couchcache](https://github.com/codingsince1985/couchcache) | 53 | 4 | 0 | A RESTful caching micro-service in Go backed by Couchbase | 2015-04-05 07:13:05 | 2021-03-03 04:44:10 |
| [hare](https://github.com/jameycribbs/hare) | 41 | 3 | 1 | Hare is a nimble little database management system for Go. | 2016-10-05 20:05:45 | 2021-06-21 15:45:05 |
| [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) | 37 | 4 | 2 | golang bigcache with clustering as a library. | 2017-12-18 07:48:07 | 2021-03-12 09:43:04 |
| [coffer](https://github.com/claygod/coffer) | 25 | 1 | 0 | Simply ACID* key-value database. At the medium or even low latency it tries to provide greater throughput without losing the ACID properties of the database. The database provides the ability to create record headers at own discretion and use them as transactions. The maximum size of stored data is limited by the size of the computer's RAM. | 2019-05-13 18:30:23 | 2021-05-11 17:35:16 |
| [tempdb](https://github.com/rafaeljesus/tempdb) | 15 | 1 | 0 | Key-value store for temporary items :memo: | 2017-03-17 18:03:42 | 2021-04-26 04:59:34 |
| [ttlcache](https://github.com/cheshir/ttlcache) | 3 | 1 | 0 | Simple in-memory key-value storage with TTL for each record. | 2021-01-06 19:24:26 | 2021-05-24 05:31:48 |
</details>

### Database - SQL query builder
libraries for building and using SQL

<sup>*Last Update: 2021-06-28 15:13:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [squirrel](https://github.com/Masterminds/squirrel) | 3,992 | 311 | 38 | Fluent SQL generation for golang | 2014-01-18 05:29:58 | 2021-06-27 20:42:18 |
| [xo](https://github.com/xo/xo) | 2,801 | 254 | 24 | Command line tool to generate idiomatic Go code for SQL databases supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server | 2016-02-05 10:22:20 | 2021-06-26 11:01:19 |
| [gendry](https://github.com/didi/gendry) | 1,230 | 151 | 5 | a golang library for sql builder | 2017-12-01 08:15:43 | 2021-06-27 15:59:08 |
| [goqu](http://doug-martin.github.io/goqu/) | 1,082 | 100 | 41 | SQL builder and query library for golang | 2015-02-21 01:06:00 | 2021-06-27 16:33:14 |
| [dotsql](https://github.com/gchaincl/dotsql) | 590 | 45 | 6 | A Golang library for using SQL. | 2014-11-20 12:14:39 | 2021-06-23 08:25:29 |
| [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) | 536 | 48 | 36 | A Go (golang) package that enhances the standard database/sql package by providing powerful data retrieval methods as well as DB-agnostic query building capabilities. | 2015-12-10 22:39:26 | 2021-06-27 07:09:47 |
| [jet](https://github.com/go-jet/jet) | 426 | 27 | 16 | Type safe SQL builder with code generation and automatic query result data mapping | 2019-03-02 11:06:23 | 2021-06-22 19:59:43 |
| [dbq](https://github.com/rocketlaunchr/dbq) | 298 | 15 | 1 | Zero boilerplate database operations for Go | 2019-07-11 02:17:33 | 2021-06-28 02:57:56 |
| [sqrl](https://github.com/elgris/sqrl) | 224 | 27 | 8 | Fluent SQL generation for golang | 2014-06-25 10:03:06 | 2021-06-10 08:14:18 |
| [sqlingo](https://github.com/lqs/sqlingo) | 154 | 11 | 0 | 💥 A lightweight DSL & ORM which helps you to write SQL in Go. | 2018-11-18 14:11:03 | 2021-06-27 05:52:22 |
| [go-structured-query](https://bokwoon95.github.io/sq/) | 105 | 4 | 2 | Type safe SQL query builder and struct mapper for Go | 2020-05-30 14:07:30 | 2021-05-31 10:15:41 |
| [igor](https://github.com/galeone/igor) | 85 | 3 | 0 | igor is an abstraction layer for PostgreSQL with a gorm like syntax. | 2016-03-10 14:45:08 | 2021-06-27 05:51:37 |
| [go-hasql](https://golang.yandex) | 83 | 4 | 2 | Go library for accessing multi-host SQL database installations | 2020-08-19 09:56:00 | 2021-05-07 22:05:04 |
| [godbal](https://github.com/xujiajun/godbal) | 52 | 27 | 0 | Database Abstraction Layer (dbal) for Go. Support SQL builder and get result easily  (now only support mysql) | 2018-02-28 05:47:42 | 2021-04-22 06:11:02 |
| [buildsqlx](https://github.com/arthurkushman/buildsqlx) | 36 | 3 | 8 | Go database query builder library for PostgreSQL | 2019-08-18 08:18:21 | 2021-06-14 12:28:34 |
| [qry](https://github.com/HnH/qry) | 20 | 3 | 1 | Write your SQL queries in raw files with all benefits of modern IDEs, use them in an easy way inside your application with all the profit of compile time constants | 2019-08-20 09:01:00 | 2021-02-03 14:40:44 |
| [sqlf](https://github.com/leporo/sqlf) | 18 | 1 | 0 | Fast SQL query builder for Go | 2019-07-20 07:03:27 | 2021-04-29 15:03:08 |
| [gosql](https://twharmon.gitbook.io/gosql/) | 15 | 0 | 0 | SQL query builder for Go | 2020-01-08 17:13:09 | 2020-12-25 02:43:58 |
| [mpath-go](https://github.com/spacetab-io/mpath-go) | 8 | 0 | 0 | Golang package for MPTT (Modified Preorder Tree Traversal) - materialized path realisation. | 2020-01-09 15:04:45 | 2021-06-17 02:48:46 |
| [ormlite](https://github.com/pupizoid/ormlite) | 0 | 0 | 2 | Lightweight package containing some ORM-like features and helpers for sqlite databases. | 2018-06-28 13:42:09 | 2021-01-13 15:25:06 |
</details>

### Database Drivers - Multiple Backends.


<sup>*Last Update: 2021-07-01 08:37:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [cayley](https://cayley.io) | 13,843 | 1,252 | 89 | An open-source graph database | 2014-06-05 18:49:41 | 2021-06-29 03:49:31 |
| [gokv](https://github.com/philippgille/gokv) | 343 | 34 | 30 | Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more) | 2018-10-08 18:55:22 | 2021-06-30 22:45:42 |
| [cachego](https://github.com/faabiosr/cachego) | 156 | 9 | 1 | Golang Cache component - Multiple drivers | 2016-10-05 18:10:03 | 2021-06-23 17:52:44 |
| [dsc](https://github.com/viant/dsc) | 22 | 6 | 0 | Datastore Connectivity in go | 2016-06-13 20:18:10 | 2021-04-17 04:50:27 |
</details>

### Database Drivers - NoSQL Databases


<sup>*Last Update: 2021-06-27 02:00:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [redis](https://redis.uptrace.dev/) | 11,796 | 1,503 | 116 | Type-safe Redis client for Golang | 2012-07-25 13:01:39 | 2021-06-26 06:59:40 |
| [redigo](https://github.com/gomodule/redigo) | 8,461 | 1,174 | 17 | Go client for Redis | 2012-04-14 04:31:58 | 2021-06-25 17:49:19 |
| [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) | 5,831 | 658 | 7 | The Go driver for MongoDB | 2017-02-08 17:18:02 | 2021-06-25 22:11:26 |
| [gocql](http://gocql.github.io/) | 2,102 | 515 | 156 | Package gocql implements a fast and robust Cassandra client for the Go programming language. | 2012-08-26 15:42:42 | 2021-06-25 17:46:15 |
| [mgo](https://github.com/globalsign/mgo) | 1,901 | 233 | 63 | The MongoDB driver for Go | 2017-04-13 11:14:04 | 2021-06-25 04:08:29 |
| [rethinkdb-go](https://github.com/rethinkdb/rethinkdb-go) | 1,565 | 176 | 15 | Go language driver for RethinkDB | 2013-09-12 13:56:27 | 2021-06-23 12:03:34 |
| [gomemcache](https://github.com/bradfitz/gomemcache) | 1,352 | 375 | 45 | Go Memcached client library #golang | 2011-06-28 19:29:12 | 2021-06-24 12:07:37 |
| [qmgo](https://github.com/qiniu/qmgo) | 521 | 58 | 25 | Qmgo - The Go driver for MongoDB. It‘s based on official mongo-go-driver but easier to use like Mgo. | 2020-08-04 09:06:00 | 2021-06-22 03:03:52 |
| [redeo](https://github.com/bsm/redeo) | 397 | 27 | 3 | High-performance framework for building redis-protocol compatible TCP servers/services | 2014-03-06 08:46:18 | 2021-06-21 09:37:53 |
| [neoism](https://github.com/jmcvetta/neoism) | 375 | 57 | 16 | Neo4j client for Golang | 2012-07-12 07:42:33 | 2021-06-08 21:35:52 |
| [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) | 365 | 164 | 21 | Aerospike Client Go  | 2014-07-26 02:56:21 | 2021-06-03 14:34:08 |
| [gocb](http://blog.couchbase.com/2015/september/go-sdk-1.0-ga) | 326 | 92 | 0 | The Couchbase Go SDK | 2015-01-15 20:01:32 | 2021-06-14 16:53:32 |
| [mgm](https://github.com/Kamva/mgm) | 323 | 31 | 2 | Mongo Go Models (mgm) is a fast and simple MongoDB ODM for Go (based on official Mongo Go Driver) | 2019-12-27 14:40:51 | 2021-06-25 19:53:31 |
| [go-couchbase](https://godoc.org/github.com/couchbase/go-couchbase) | 312 | 87 | 41 | Couchbase client in Go | 2012-01-19 22:52:08 | 2021-06-22 14:57:10 |
| [go-rejson](https://github.com/nitishm/go-rejson) | 168 | 27 | 6 | Golang client for redislabs' ReJSON module with support for multilple redis clients (redigo, go-redis) | 2018-04-23 00:51:05 | 2021-06-20 08:42:55 |
| [godis](https://github.com/piaohao/godis) | 91 | 15 | 0 | redis client implement by golang, inspired by jedis. | 2019-06-14 03:14:22 | 2021-06-18 12:55:16 |
| [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) | 76 | 15 | 0 | Neo4j REST Client in golang | 2011-06-04 16:08:35 | 2020-08-28 21:16:55 |
| [arangolite](https://github.com/solher/arangolite) | 69 | 19 | 5 | Lightweight Golang driver for ArangoDB | 2015-10-04 17:27:34 | 2021-03-10 17:27:16 |
| [go-pilosa](https://www.pilosa.com/) | 46 | 20 | 12 | Go client library for Pilosa | 2016-09-30 21:37:10 | 2021-06-13 07:36:01 |
| [goforestdb](https://github.com/couchbase/goforestdb) | 30 | 4 | 7 | Go bindings for ForestDB | 2014-05-14 15:36:12 | 2021-02-25 12:12:32 |
| [neo4j](https://github.com/cihangir/neo4j) | 26 | 7 | 8 | Neo4j Rest API Client for Go lang | 2013-05-18 08:54:01 | 2020-08-28 21:16:50 |
| [goriak](https://godoc.org/gopkg.in/zegl/goriak.v3) | 26 | 5 | 5 | goriak - Go language driver for Riak KV | 2016-10-05 16:48:17 | 2021-05-21 04:31:42 |
| [xredis](https://github.com/shomali11/xredis) | 15 | 2 | 0 | Go Redis Client | 2017-06-14 00:19:26 | 2021-06-02 07:51:37 |
| [godscache](https://github.com/defcronyke/godscache) | 8 | 1 | 0 | An unofficial Google Cloud Platform Go Datastore wrapper that adds caching using memcached. For App Engine Flexible, Compute Engine, Kubernetes Engine, and more. | 2018-05-08 20:19:39 | 2021-02-06 20:45:22 |
| [asc](https://github.com/viant/asc) | 6 | 1 | 0 | Datastore Connectivity for Aerospike for go | 2016-06-13 20:22:31 | 2020-08-28 21:15:29 |
| [gocosmos](http://blog.couchbase.com/2015/september/go-sdk-1.0-ga) | 3 | 1 | 0 | Go driver for Azure CosmosDB SQL API | 2020-12-06 07:03:43 | 2021-05-18 10:58:47 |
</details>

### Database Drivers - Relational Databases


<sup>*Last Update: 2021-06-28 15:13:35*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql) | 10,998 | 1,941 | 77 | Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package | 2012-12-09 20:33:55 | 2021-06-28 01:55:27 |
| [pq](https://pkg.go.dev/github.com/lib/pq) | 6,620 | 807 | 285 | Pure Go Postgres driver for database/sql | 2012-03-12 18:50:22 | 2021-06-28 08:09:45 |
| [go-sqlite3](http://mattn.github.io/go-sqlite3) | 4,913 | 852 | 103 | sqlite3 driver for go using database/sql | 2011-11-11 12:36:50 | 2021-06-28 01:58:49 |
| [pgx](https://github.com/jackc/pgx) | 4,209 | 417 | 153 | PostgreSQL driver and toolkit for Go | 2013-03-30 19:06:26 | 2021-06-28 02:24:28 |
| [go-mssqldb](https://github.com/denisenkom/go-mssqldb) | 1,409 | 359 | 134 | Microsoft SQL server driver written in go language | 2013-12-16 00:10:47 | 2021-06-28 06:37:11 |
| [go-oci8](https://mattn.kaoriya.net/) | 551 | 200 | 6 | Oracle driver for Go using database/sql | 2012-02-29 12:19:16 | 2021-06-26 10:16:40 |
| [godror](https://github.com/godror/godror) | 260 | 44 | 5 | GO DRiver for ORacle DB | 2019-11-21 21:23:17 | 2021-06-23 05:35:29 |
| [firebirdsql](https://github.com/nakagami/firebirdsql) | 151 | 44 | 8 | Firebird RDBMS sql driver for Go (golang) | 2013-08-27 13:09:14 | 2021-06-26 00:19:07 |
| [go-adodb](http://mattn.kaoriya.net/) | 120 | 30 | 17 | Microsoft ActiveX Object DataBase driver for go that using exp/sql | 2011-11-14 04:32:50 | 2021-06-09 05:54:10 |
| [gofreetds](https://github.com/minus5/gofreetds) | 104 | 43 | 18 | Go Sql Server database driver. | 2012-12-06 17:29:26 | 2021-06-17 11:58:24 |
| [calcite-avatica-go](https://github.com/apache/calcite-avatica-go) | 74 | 19 | 1 | Mirror of Apache Calcite - Avatica Go SQL Driver | 2017-08-08 07:00:08 | 2021-06-24 17:57:25 |
| [sqinn-go](https://github.com/cvilsmeier/sqinn-go) | 67 | 7 | 0 | SQLite with pure Go | 2020-06-06 20:37:12 | 2021-06-27 02:46:31 |
| [bgc](https://github.com/viant/bgc) | 15 | 4 | 0 | Datastore Connectivity for BigQuery in go | 2016-06-13 20:24:26 | 2020-09-06 06:58:35 |
| [pig](https://github.com/alexeyco/pig) | 5 | 0 | 0 | Simple pgx wrapper to execute and scan query results | 2021-04-15 15:33:23 | 2021-06-13 02:33:54 |
</details>

### Database Drivers - Search and Analytic Databases


<sup>*Last Update: 2021-06-22 09:02:00*</sup>
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

### Date and Time
Libraries for working with dates and times.

<sup>*Last Update: 2021-06-26 08:30:56*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [now](https://github.com/jinzhu/now) | 3,118 | 186 | 5 | Now is a time toolkit for golang | 2013-11-18 10:55:30 | 2021-06-25 14:51:41 |
| [dateparse](https://github.com/araddon/dateparse) | 1,441 | 103 | 33 | GoLang Parse many date strings without knowing format in advance. | 2014-04-21 02:55:48 | 2021-06-25 13:49:17 |
| [carbon](https://github.com/uniplaces/carbon) | 597 | 45 | 11 | Carbon for Golang, an extension for Time | 2016-08-03 10:55:52 | 2021-06-24 02:55:52 |
| [durafmt](https://github.com/hako/durafmt) | 389 | 41 | 5 | :clock8:  Better time duration formatting in Go!  | 2016-05-20 21:49:33 | 2021-06-22 13:50:53 |
| [timeutil](https://github.com/leekchan/timeutil) | 186 | 12 | 1 | timeutil - useful extensions (Timedelta, Strftime, ...) to the golang's time package | 2015-08-02 01:32:06 | 2021-06-03 03:54:54 |
| [gostradamus](https://github.com/bykof/gostradamus) | 153 | 2 | 1 | Gostradamus: Better DateTimes for Go 🕰️ | 2020-04-07 12:29:21 | 2021-05-27 19:45:04 |
| [go-persian-calendar](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 97 | 13 | 3 | The implementation of Persian (Solar Hijri) Calendar in Go | 2016-01-31 18:40:23 | 2021-06-16 06:42:33 |
| [iso8601](https://github.com/relvacode/iso8601) | 89 | 5 | 1 | A fast ISO8601 date parser for Go | 2017-04-25 15:54:18 | 2021-06-18 10:18:41 |
| [timespan](https://github.com/SaidinWoT/timespan) | 74 | 9 | 3 | Golang package to manipulate time intervals. | 2014-10-07 03:57:02 | 2020-10-20 14:32:08 |
| [date](https://godoc.org/github.com/rickb777/date) | 72 | 15 | 2 | A Go package for working with dates | 2015-11-23 22:58:07 | 2021-06-22 15:28:28 |
| [feiertage](https://github.com/wlbr/feiertage) | 39 | 4 | 1 | Gesetzliche Feiertage und mehr in Deutschland und Österreich (Bank holidays/public holidays in Austria and Germany) | 2015-11-04 14:19:27 | 2021-06-20 19:44:24 |
| [go-sunrise](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 35 | 5 | 0 | Go package for calculating the sunrise and sunset times for a given location | 2017-06-15 20:49:41 | 2021-06-03 14:15:31 |
| [go-str2duration](https://pkg.go.dev/github.com/yaa110/go-persian-calendar) | 25 | 1 | 1 | Convert string to duration in golang | 2020-02-02 06:04:07 | 2021-06-22 14:03:36 |
| [kair](https://github.com/GuilhermeCaruso/kair) | 19 | 5 | 0 | :clock1: Date and Time - Golang Formatting Library | 2018-10-03 15:44:07 | 2021-04-23 00:08:00 |
| [nulltime](https://github.com/kirillDanshin/nulltime) | 11 | 2 | 0 | Now is a time toolkit for golang | 2016-03-06 01:44:58 | 2020-08-28 21:19:40 |
| [cronrange](https://github.com/1set/cronrange) | 10 | 3 | 0 | time range expression in cron style | 2019-11-10 01:30:45 | 2021-04-20 15:24:22 |
| [tuesday](https://godoc.org/github.com/osteele/tuesday) | 9 | 1 | 1 | Ruby-compatible strftime for golang | 2017-08-10 20:46:26 | 2021-06-19 03:38:21 |
| [strftime](https://github.com/awoodbeck/strftime) | 7 | 2 | 0 | C99-compatible strftime formatter for use with Go time.Time instances. | 2018-02-10 00:35:46 | 2020-11-22 03:01:25 |
| [go-week](https://github.com/stoewer/go-week) | 5 | 3 | 2 | A Go package to work with ISO 8601 week dates | 2018-02-23 07:02:37 | 2021-04-27 06:46:04 |
</details>

### Distributed Systems
Packages that help with building Distributed Systems.

<sup>*Last Update: 2021-06-26 03:25:32*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [etcd](https://etcd.io) | 36,319 | 7,758 | 188 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2021-06-25 13:04:18 |
| [kit](https://gokit.io) | 20,473 | 2,120 | 91 | A standard library for microservices. | 2015-02-03 00:01:19 | 2021-06-25 09:32:14 |
| [go-micro](https://go-micro.dev) | 16,198 | 1,772 | 23 | Go Micro is a framework for distributed systems development | 2015-01-13 23:30:18 | 2021-06-25 13:04:32 |
| [grpc-go](https://grpc.io) | 13,954 | 3,031 | 96 | The Go language implementation of gRPC. HTTP/2 based RPC | 2014-12-08 18:59:34 | 2021-06-25 11:46:58 |
| [micro](https://micro.mu) | 10,166 | 891 | 80 | Micro is a distributed OS built for the Cloud | 2015-01-16 22:35:14 | 2021-06-25 12:17:20 |
| [nats-server](https://nats.io) | 9,490 | 916 | 143 | High-Performance server for NATS.io, the cloud and edge native messaging system. | 2012-10-29 16:12:24 | 2021-06-25 08:11:02 |
| [go-zero](https://go-zero.dev) | 9,138 | 1,129 | 71 | go-zero is a web and rpc framework written in Go. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. | 2020-08-07 15:37:57 | 2021-06-25 13:06:36 |
| [rpcx](https://rpcx.io) | 5,756 | 900 | 10 | Best microservices framework in Go, like alibaba Dubbo, but with more features, Scale easily. Try it. Test it. If you feel it's better, use it! 𝐉𝐚𝐯𝐚有𝐝𝐮𝐛𝐛𝐨, 𝐆𝐨𝐥𝐚𝐧𝐠有𝐫𝐩𝐜𝐱! | 2016-05-18 09:34:05 | 2021-06-25 08:30:27 |
| [raft](https://godoc.org/cirello.io/pglock) | 4,821 | 652 | 22 | Golang implementation of the Raft consensus protocol | 2013-11-05 00:41:20 | 2021-06-25 12:36:16 |
| [lura](https://luraproject.org) | 4,254 | 405 | 53 | Ultra performant API Gateway with middlewares. A project hosted at The Linux Foundation | 2016-11-04 18:37:13 | 2021-06-25 04:55:29 |
| [tendermint](https://tendermint.com/) | 4,171 | 1,327 | 343 | ⟁ Tendermint Core (BFT Consensus) in Go | 2014-05-14 23:21:35 | 2021-06-25 13:14:43 |
| [torrent](https://github.com/anacrolix/torrent) | 3,921 | 481 | 43 | Full-featured BitTorrent client package and utilities | 2015-01-08 21:10:42 | 2021-06-25 11:42:49 |
| [dragonboat](https://github.com/lni/dragonboat) | 3,689 | 373 | 12 | A feature complete and high performance multi-group Raft library in Go.   | 2018-12-23 07:02:04 | 2021-06-25 13:07:39 |
| [glow](https://github.com/chrislusf/glow) | 2,959 | 228 | 11 | Glow is an easy-to-use distributed computation system written in Go, similar to Hadoop Map Reduce, Spark, Flink, Storm, etc. I am also working on another similar pure Go system, https://github.com/chrislusf/gleam , which is more flexible and more performant. | 2015-06-14 00:33:48 | 2021-06-24 03:19:10 |
| [emitter](https://emitter.io) | 2,910 | 264 | 20 | High performance, distributed and low latency publish-subscribe platform. | 2016-10-29 08:52:21 | 2021-06-24 15:29:46 |
| [gleam](https://github.com/chrislusf/gleam) | 2,837 | 256 | 36 | Fast, efficient, and scalable distributed map/reduce system, DAG execution, in memory or on disk, written in pure Go, runs standalone or distributedly. | 2016-08-26 08:44:48 | 2021-06-25 05:05:42 |
| [liftbridge](https://liftbridge.io) | 2,102 | 91 | 35 | Lightweight, fault-tolerant message streams. | 2017-10-13 19:50:26 | 2021-06-23 14:34:25 |
| [hprose-golang](https://github.com/hprose/hprose-golang) | 1,166 | 203 | 20 | Hprose is a cross-language RPC. This project is Hprose for Golang. | 2014-02-14 03:16:43 | 2021-06-18 09:52:03 |
| [ringpop-go](http://www.uber.com) | 687 | 61 | 26 | Scalable, fault-tolerant application-layer sharding for Go applications | 2015-06-05 22:48:53 | 2021-06-24 15:34:03 |
| [gorpc](https://github.com/valyala/gorpc) | 636 | 93 | 15 | Simple, fast and scalable golang rpc library for high load | 2014-11-20 17:02:37 | 2021-06-21 14:34:15 |
| [rain](https://put.io) | 613 | 37 | 0 | 🌧 BitTorrent client and library in Go | 2014-05-21 09:17:24 | 2021-06-24 11:47:40 |
| [go-health](https://github.com/InVisionApp/go-health) | 595 | 37 | 9 | Library for enabling asynchronous health checks in your service | 2017-11-29 21:00:07 | 2021-06-24 12:04:32 |
| [resgate](https://resgate.io) | 493 | 35 | 15 | A Realtime API Gateway used with NATS to build REST, real time, and RPC APIs, where all your clients are synchronized seamlessly. | 2018-02-22 12:06:26 | 2021-06-22 16:11:24 |
| [redislock](https://put.io) | 430 | 60 | 0 | Simplified distributed locking implementation using Redis | 2019-06-24 11:10:10 | 2021-06-25 07:50:41 |
| [go-sundheit](https://pdu.pub) | 408 | 22 | 4 | A library built to provide support for defining service health for golang services. It allows you to register async health checks for your dependencies and the service itself, provides a health endpoint that exposes their status, and health metrics. | 2019-04-08 12:54:01 | 2021-06-25 07:30:27 |
| [digota](http://digota.com) | 393 | 64 | 10 | ecommerce microservice | 2017-08-14 12:01:37 | 2021-06-21 22:41:30 |
| [consistent](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 384 | 44 | 3 | Consistent hashing with bounded loads in Golang | 2018-03-25 15:38:27 | 2021-06-06 10:45:04 |
| [sleuth](http://ursiform.github.io/sleuth/) | 336 | 21 | 0 | A Go library for master-less peer-to-peer autodiscovery and RPC between HTTP services | 2016-04-23 14:21:41 | 2021-06-22 18:21:52 |
| [go-jump](https://github.com/dgryski/go-jump) | 327 | 27 | 1 | go-jump: Jump consistent hashing | 2014-06-15 22:12:04 | 2021-06-09 21:37:50 |
| [dht](https://github.com/anacrolix/dht) | 189 | 41 | 3 | dht is used by anacrolix/torrent, and is intended for use as a library in other projects both torrent related and otherwise | 2016-12-14 00:34:42 | 2021-06-12 18:06:17 |
| [arpc](https://github.com/lesismal/arpc) | 182 | 18 | 0 | More effective network communication, two-way calling, notify and broadcast supported. | 2020-05-19 11:30:05 | 2021-06-18 20:26:34 |
| [jsonrpc](https://github.com/ybbus/jsonrpc) | 178 | 65 | 2 | A simple go implementation of json rpc 2.0 client over http | 2016-11-10 11:27:55 | 2021-06-23 03:08:31 |
| [jsonrpc](https://github.com/osamingo/jsonrpc) | 149 | 14 | 4 | The jsonrpc package helps implement of JSON-RPC 2.0 | 2016-10-28 13:36:59 | 2021-06-21 15:40:00 |
| [celeriac.v1](https://github.com/svcavallar/celeriac.v1) | 67 | 8 | 1 | Golang client library for adding support for interacting and monitoring Celery workers, tasks and events. | 2015-10-10 07:27:33 | 2021-06-13 12:55:46 |
| [dynamolock](https://godoc.org/cirello.io/dynamolock) | 67 | 33 | 0 | [Mirror] DynamoDB Lock Client for Go | 2018-07-08 11:13:00 | 2021-06-14 16:56:45 |
| [doublejump](https://github.com/edwingeng/doublejump) | 60 | 12 | 0 | A revamped Google's jump consistent hash | 2018-06-26 16:04:50 | 2021-04-17 21:27:36 |
| [dot](https://github.com/dotchain/dot) | 57 | 1 | 0 | distributed data sync with operational transformation/transforms  | 2017-12-18 01:08:12 | 2021-06-13 11:31:16 |
| [semaphore](https://jexia.github.io/semaphore/) | 55 | 12 | 9 | Take control of your data, connect with anything, and expose it anywhere through protocols such as HTTP, GraphQL, and gRPC. | 2020-02-05 16:39:39 | 2021-06-04 08:27:09 |
| [outboxer](https://github.com/italolelis/outboxer) | 53 | 7 | 9 | A library that implements the outboxer pattern in go | 2019-02-01 09:50:13 | 2021-06-08 08:28:12 |
| [flowgraph](https://emitter.io) | 41 | 4 | 0 | Flowgraph package for scalable asynchronous system development | 2018-08-29 21:45:26 | 2021-06-13 11:33:04 |
| [drmaa](https://github.com/dgruber/drmaa) | 32 | 15 | 0 | Compute cluster (HPC) job submission library for Go (#golang) based on the open DRMAA standard. | 2013-03-17 12:58:02 | 2021-05-11 03:56:31 |
| [go-pdu](https://pdu.pub) | 28 | 3 | 0 | Parallel Digital Universe - A decentralized identity-based social network | 2018-10-08 08:13:22 | 2021-06-24 07:58:57 |
| [pglock](https://godoc.org/cirello.io/pglock) | 28 | 6 | 0 | [Mirror] PostgreSQL Lock Client for Go | 2018-12-17 17:43:41 | 2021-05-27 19:38:47 |
| [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) | 21 | 1 | 2 | MySQL Backed Locking Primitive | 2020-06-06 16:30:07 | 2021-04-09 09:33:30 |
| [dynatomic](https://github.com/tylfin/dynatomic) | 14 | 1 | 0 | Dynatomic is a library for using dynamodb as an atomic counter | 2019-02-08 17:45:14 | 2020-12-25 16:56:32 |
| [micro](https://github.com/gmsec/micro) | 13 | 4 | 0 | A Go distributed systems development framework | 2020-05-03 01:16:16 | 2021-05-23 07:59:32 |
| [consistenthash](https://research.googleblog.com/2017/04/consistent-hashing-with-bounded-loads.html) | 9 | 1 | 0 | A Go library that implements Consistent Hashing | 2020-04-22 16:01:25 | 2021-04-09 18:54:48 |
</details>

### Dynamic DNS
Tools for updating dynamic DNS records.

<sup>*Last Update: 2021-07-15 09:25:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [godns](https://xiaozhou.net/godns-project-2014-05-18.html) | 839 | 158 | 8 | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 2014-05-11 11:49:17 | 2021-07-14 09:20:13 |
| [ddns](https://github.com/skibish/ddns) | 193 | 19 | 0 | Personal DDNS client with Digital Ocean Networking DNS as backend. | 2017-03-13 21:02:27 | 2021-07-01 21:29:36 |
</details>

### Email
Libraries and tools that implement email creation and sending.

<sup>*Last Update: 2021-06-28 16:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [MailHog](https://github.com/mailhog/MailHog) | 8,668 | 651 | 186 | Web and API based SMTP testing | 2014-04-16 22:28:49 | 2021-06-28 09:12:46 |
| [hermes](https://github.com/matcornic/hermes) | 2,291 | 182 | 25 | Golang package that generates clean, responsive HTML e-mails for sending transactional mail | 2017-03-25 18:25:36 | 2021-06-27 03:30:33 |
| [email](https://github.com/jordan-wright/email) | 1,766 | 245 | 47 | Robust and flexible email library for Go | 2013-12-12 20:11:59 | 2021-06-28 01:05:13 |
| [go-imap](https://github.com/emersion/go-imap) | 1,297 | 182 | 60 |  :inbox_tray: An IMAP library for clients and servers | 2016-04-26 17:59:18 | 2021-06-28 08:39:58 |
| [sendgrid-go](https://sendgrid.com) | 736 | 230 | 12 | The Official Twilio SendGrid Led, Community Driven Golang API Library | 2013-09-12 03:31:13 | 2021-06-25 15:51:09 |
| [mailgun-go](https://mailchain.xyz) | 531 | 118 | 7 | Go library for sending mail with the Mailgun API. | 2014-02-28 00:28:44 | 2021-06-24 09:51:30 |
| [chasquid](https://blitiri.com.ar/p/chasquid/) | 445 | 27 | 1 | SMTP (email) server with a focus on simplicity, security, and ease of operation [mirror] | 2016-11-03 01:28:05 | 2021-06-25 11:20:34 |
| [email-verifier](https://github.com/AfterShip/email-verifier) | 248 | 26 | 0 | :white_check_mark: A Go library for email verification without sending any emails. | 2020-12-18 08:47:28 | 2021-06-24 14:53:03 |
| [go-message](https://github.com/emersion/go-message) | 209 | 63 | 19 | :envelope: A streaming Go library for the Internet Message Format and mail messages | 2016-12-31 09:31:52 | 2021-06-27 19:13:50 |
| [hectane](https://github.com/hectane/hectane) | 207 | 25 | 16 | Lightweight SMTP client written in Go | 2015-08-28 01:36:47 | 2021-06-26 06:27:58 |
| [douceur](https://github.com/aymerick/douceur) | 194 | 33 | 9 | A simple CSS parser and inliner in Go | 2015-04-09 10:21:26 | 2021-06-02 19:06:59 |
| [go-simple-mail](https://github.com/xhit/go-simple-mail) | 165 | 34 | 7 | Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP. | 2019-09-15 05:38:54 | 2021-06-24 11:59:36 |
| [mailchain](https://mailchain.xyz) | 82 | 42 | 45 | Using Mailchain, blockchain users can now send and receive rich-media HTML messages with attachments via a blockchain address. | 2019-04-11 17:37:31 | 2021-06-20 12:45:11 |
| [go-premailer](https://github.com/vanng822/go-premailer) | 71 | 12 | 3 | Inline styling for html mail in golang | 2015-02-16 22:19:18 | 2021-05-23 17:07:47 |
| [go-dkim](https://github.com/toorop/go-dkim) | 69 | 29 | 4 | DKIM package for golang | 2015-04-29 15:38:27 | 2021-05-27 11:39:05 |
| [smtp](https://sendgrid.com) | 63 | 22 | 6 | MailHog SMTP Protocol | 2014-12-24 16:13:49 | 2021-05-26 06:49:05 |
| [go-email-validator](https://github.com/go-email-validator/go-email-validator) | 11 | 3 | 4 | 📧 Golang Email address validator | 2020-12-10 18:27:20 | 2021-06-13 02:30:45 |
</details>

### Embeddable Scripting Languages
Embedding other languages inside your go code.

<sup>*Last Update: 2021-06-28 15:13:29*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopher-lua](https://github.com/yuin/gopher-lua) | 4,148 | 466 | 87 | GopherLua: VM and compiler for Lua in Go | 2015-02-15 13:23:37 | 2021-06-28 06:56:25 |
| [tengo](https://tengolang.com) | 2,324 | 132 | 43 | A fast script language for Go | 2019-01-09 07:17:17 | 2021-06-28 07:48:22 |
| [goja](https://github.com/dop251/goja) | 2,251 | 187 | 19 | ECMAScript/JavaScript engine in pure Go | 2016-11-04 22:04:06 | 2021-06-27 15:26:25 |
| [go-lua](https://github.com/Shopify/go-lua) | 2,109 | 154 | 38 | A Lua VM in Go | 2013-12-20 17:29:43 | 2021-06-24 13:07:05 |
| [expr](https://github.com/antonmedv/expr) | 1,824 | 161 | 29 | Expression language for Go | 2018-07-14 15:57:34 | 2021-06-27 06:01:50 |
| [go-python](https://github.com/sbinet/go-python) | 1,277 | 124 | 26 | naive go bindings to the CPython C-API | 2012-07-09 15:43:31 | 2021-06-28 00:26:27 |
| [anko](http://play-anko.appspot.com/) | 1,140 | 108 | 19 | Scriptable interpreter written in golang | 2014-03-28 07:29:40 | 2021-06-27 19:05:10 |
| [cel-go](https://opensource.google.com/projects/cel) | 820 | 91 | 30 | Fast, portable, non-Turing complete expression evaluation with gradual typing (Go) | 2018-03-09 22:57:58 | 2021-06-26 18:49:55 |
| [go-php](https://github.com/deuill/go-php) | 800 | 91 | 20 | PHP bindings for the Go programming language (Golang) | 2015-09-17 21:23:52 | 2021-06-23 15:10:25 |
| [go-duktape](https://github.com/olebedev/go-duktape) | 764 | 86 | 7 | Duktape JavaScript engine bindings for Go | 2015-01-08 05:09:05 | 2021-06-15 02:55:07 |
| [golua](https://github.com/aarzilli/golua) | 545 | 154 | 11 | Go bindings for Lua C API - in progress | 2010-12-06 21:39:53 | 2021-06-24 11:22:48 |
| [gisp](https://docs.gentee.org) | 461 | 32 | 1 | Simple LISP in Go | 2014-01-11 14:05:43 | 2021-05-19 07:31:40 |
| [gval](https://github.com/PaesslerAG/gval) | 338 | 43 | 3 | Expression evaluation in golang | 2017-09-27 08:32:49 | 2021-06-25 16:07:23 |
| [gentee](https://docs.gentee.org) | 75 | 8 | 0 | Gentee - script programming language for automation. It uses VM and compiler written in Go (Golang). | 2018-01-14 15:49:05 | 2021-05-26 09:06:58 |
| [binder](https://github.com/alexeyco/binder) | 51 | 7 | 0 | High level go to Lua binder. Write less, do more. | 2017-04-02 17:14:52 | 2021-06-03 01:31:08 |
| [purl](https://github.com/ian-kent/purl) | 31 | 2 | 2 | Perl, but fluffy like a cat! | 2014-11-29 19:06:01 | 2021-04-30 14:25:41 |
| [ngaro](https://github.com/db47h/ngaro) | 20 | 1 | 1 | An embeddable implementation of the Ngaro Virtual Machine for Go programs | 2016-08-09 15:23:50 | 2020-04-06 11:44:00 |
| [ecal](https://github.com/krotik/ecal) | 10 | 0 | 0 | A simple embeddable scripting language which supports concurrent event processing. | 2020-11-30 15:58:56 | 2021-05-23 09:52:38 |
</details>

### Error Handling
Libraries for handling errors.

<sup>*Last Update: 2021-06-25 14:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [errors](https://godoc.org/github.com/pkg/errors) | 7,015 | 530 | 41 | Simple error handling primitives | 2015-12-27 12:05:38 | 2021-06-24 14:46:36 |
| [go-multierror](https://github.com/hashicorp/go-multierror) | 1,291 | 85 | 12 | A Go (golang) package for representing a list of errors as a single error. | 2014-12-15 20:12:26 | 2021-06-24 09:55:19 |
| [eris](https://github.com/rotisserie/eris) | 815 | 19 | 0 | eris provides a better way to handle, trace, and log errors in Go 🎆 | 2019-09-07 16:50:33 | 2021-06-25 06:23:24 |
| [errorx](https://github.com/joomcode/errorx) | 748 | 25 | 4 | A comprehensive error handling library for Go | 2018-08-17 08:02:10 | 2021-06-17 15:47:01 |
| [tracerr](https://github.com/ztrue/tracerr) | 660 | 24 | 1 | Golang errors with stack trace and source fragments. | 2019-02-06 18:57:46 | 2021-06-18 16:57:15 |
| [errlog](https://github.com/snwfdhmp/errlog) | 396 | 16 | 1 | Reduce debugging time while programming Go. Use static and stack-trace analysis to determine which func call causes the error. | 2019-02-16 23:19:05 | 2021-06-14 15:23:44 |
| [emperror](https://emperror.dev/emperror) | 216 | 11 | 5 | The Emperor takes care of all errors personally | 2017-06-13 00:24:28 | 2021-06-22 03:36:37 |
| [errors](https://emperror.dev/errors) | 96 | 4 | 6 | Drop-in replacement for the standard library errors package and github.com/pkg/errors | 2019-07-09 13:02:52 | 2021-06-04 10:46:01 |
| [errors](https://github.com/bnkamalesh/errors) | 23 | 3 | 0 | A drop-in replacement for Go errors, with some added sugar! Unwrap user-friendly messages, HTTP status code, easy wrapping with multiple error types. | 2020-07-17 18:57:04 | 2021-06-19 13:37:32 |
| [falcon](https://github.com/SonicRoshan/falcon) | 6 | 0 | 0 | A Simple Yet Highly Powerful Package For Error Handling | 2019-09-09 12:49:43 | 2020-07-29 05:20:41 |
| [errors](https://github.com/neuronlabs/errors) | 3 | 0 | 0 | Simple golang error handling with classification primitives. | 2019-07-26 00:15:36 | 2020-01-23 14:57:00 |
| [errors](https://github.com/PumpkinSeed/errors) | 2 | 0 | 0 | Simple and efficient error package  | 2020-01-08 21:12:51 | 2020-02-08 06:50:05 |
</details>

### File Handling
Libraries for handling files and file systems.

<sup>*Last Update: 2021-07-01 08:37:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [afero](https://github.com/spf13/afero) | 3,836 | 344 | 83 | A FileSystem Abstraction System for Go | 2014-10-28 14:19:05 | 2021-06-30 15:25:34 |
| [pdfcpu](https://pdfcpu.io) | 2,438 | 185 | 41 | A PDF processor written in Go. | 2017-06-18 17:27:38 | 2021-06-30 08:40:04 |
| [notify](https://github.com/rjeczalik/notify) | 666 | 98 | 38 | File system event notification library on steroids. | 2014-09-08 16:09:34 | 2021-06-30 19:27:51 |
| [copy](https://pkg.go.dev/github.com/otiai10/copy) | 319 | 76 | 8 | Go copy directory recursively | 2017-09-01 03:18:56 | 2021-06-25 11:54:07 |
| [bigfile](https://learnku.com/docs/bigfile) | 191 | 37 | 1 | Bigfile -- a file transfer system that supports http, rpc and ftp protocol   https://bigfile.site   | 2019-07-15 10:43:50 | 2021-05-20 13:04:55 |
| [afs](https://github.com/viant/afs) | 148 | 12 | 0 | Abstract File Storage | 2019-08-19 18:43:38 | 2021-06-09 14:47:02 |
| [vfs](https://github.com/C2FO/vfs) | 110 | 10 | 8 | Pluggable, extensible virtual file system for Go | 2017-08-01 18:06:14 | 2021-06-24 16:51:43 |
| [go-csv-tag](https://github.com/artonge/go-csv-tag) | 86 | 19 | 0 | Read csv file from go using tags | 2017-06-18 15:31:16 | 2021-06-17 02:42:08 |
| [opc](https://github.com/qmuntal/opc) | 67 | 5 | 0 | Go implementation of the Open Packaging Conventions (OPC) | 2018-11-06 14:49:06 | 2021-06-09 11:29:01 |
| [go-exiftool](https://github.com/barasher/go-exiftool) | 64 | 15 | 3 | Golang wrapper for Exiftool : extract as much metadata as possible (EXIF, ...) from files (pictures, pdf, office documents, ...) | 2019-05-12 20:34:09 | 2021-06-30 08:03:31 |
| [skywalker](https://github.com/dixonwille/skywalker) | 64 | 6 | 1 | A package to allow one to concurrently go through a filesystem with ease | 2017-08-01 20:08:25 | 2021-04-18 23:49:38 |
| [tarfs](https://github.com/posener/tarfs) | 47 | 6 | 1 | An implementation of the FileSystem interface for tar files. | 2017-03-10 22:13:11 | 2021-06-30 11:02:38 |
| [checksum](https://github.com/codingsince1985/checksum) | 34 | 10 | 0 | Compute message digest for large files in Go | 2014-11-05 09:37:00 | 2021-06-24 01:14:10 |
| [baraka](https://github.com/xis/baraka) | 30 | 5 | 1 | a tool for handling file uploads simple | 2020-07-12 21:56:50 | 2021-06-29 11:37:57 |
| [go-gtfs](https://github.com/artonge/go-gtfs) | 27 | 14 | 0 | Load GTFS files in golang | 2017-07-09 09:30:31 | 2021-06-28 13:24:51 |
| [flop](https://github.com/homedepot/flop) | 24 | 5 | 1 | Go file operations library chasing GNU APIs. | 2019-03-01 13:41:39 | 2021-02-19 21:53:29 |
| [parquet](https://github.com/parsyl/parquet) | 21 | 1 | 0 | A library for reading and writing parquet files. | 2019-01-29 21:52:30 | 2021-06-18 18:48:59 |
| [gut](https://github.com/1set/gut) | 17 | 3 | 13 | 🍱 yet another collection of go utilities & tools | 2019-10-05 23:47:24 | 2021-05-02 15:31:54 |
| [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) | 14 | 4 | 1 | copy files for humans | 2018-10-16 07:08:24 | 2021-03-12 20:07:53 |
| [todotxt](https://github.com/1set/todotxt) | 8 | 0 | 0 | Parser for todo.txt files in Go ✅ | 2020-11-06 17:41:59 | 2021-06-11 22:20:24 |
| [higgs](https://github.com/dastoori/higgs) | 6 | 1 | 1 | A tiny cross-platform Go library to hide/unhide files and directories | 2020-12-13 18:33:10 | 2021-06-27 05:11:45 |
</details>

### Financial
Packages for accounting and finance.

<sup>*Last Update: 2021-06-25 14:25:15*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [decimal](https://github.com/shopspring/decimal) | 3,094 | 375 | 59 | Arbitrary-precision fixed-point decimal numbers in go | 2015-02-25 20:12:57 | 2021-06-25 06:41:19 |
| [go-money](https://github.com/Rhymond/go-money) | 953 | 82 | 20 | Go implementation of Fowler's Money pattern | 2017-03-20 16:23:54 | 2021-06-24 06:05:31 |
| [accounting](https://github.com/leekchan/accounting) | 659 | 50 | 6 | money and currency formatting for golang | 2015-08-10 13:23:56 | 2021-06-24 18:20:57 |
| [go-finance](https://github.com/FlashBoys/go-finance) | 534 | 49 | 4 | :warning: Deprecrated in favor of https://github.com/piquette/finance-go  | 2016-02-28 00:37:46 | 2021-06-23 04:00:30 |
| [techan](https://godoc.org/github.com/sdcoffey/techan) | 466 | 79 | 18 | Technical Analysis Library for Golang | 2017-03-08 03:04:08 | 2021-06-23 21:07:51 |
| [currency](https://pkg.go.dev/github.com/bojanz/currency) | 260 | 10 | 3 | Currency handling for Go. | 2020-04-16 15:34:39 | 2021-06-23 07:08:12 |
| [orderbook](https://github.com/i25959341/orderbook) | 209 | 68 | 5 | Matching Engine for Limit Order Book in Golang | 2018-04-24 18:05:26 | 2021-06-18 06:10:34 |
| [go-finance](https://github.com/alpeb/go-finance) | 99 | 17 | 0 | Go library containing a collection of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. | 2017-06-01 15:58:33 | 2021-06-18 06:12:09 |
| [transaction](https://github.com/claygod/transaction) | 91 | 13 | 0 | Embedded database for accounts transactions. | 2017-10-11 13:50:30 | 2021-06-10 11:42:24 |
| [ofxgo](https://github.com/aclindsa/ofxgo) | 88 | 19 | 0 | Golang library for querying and parsing OFX | 2015-11-08 13:56:53 | 2021-05-23 10:05:11 |
| [vat](https://github.com/dannyvankooten/vat) | 82 | 11 | 4 | Go package for dealing with EU VAT. Does VAT number validation & rates retrieval. | 2016-06-18 16:10:09 | 2021-05-12 01:44:47 |
| [sleet](https://github.com/BoltApp/sleet) | 58 | 7 | 6 | Payment abstraction library - one interface for multiple payment processors ( inspired by Ruby's ActiveMerchant ) | 2019-11-13 21:56:58 | 2021-06-22 02:07:27 |
| [go-finnhub](https://github.com/m1/go-finnhub) | 56 | 12 | 0 | Simple and easy to use client for stock market, forex and crypto data from finnhub.io written in Go. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges | 2020-01-13 20:47:13 | 2021-05-04 02:48:46 |
| [currency](https://github.com/bnkamalesh/currency) | 39 | 5 | 0 | A currency computations package. | 2017-05-09 06:06:38 | 2021-04-01 19:00:35 |
| [fastme](https://github.com/newity/fastme) | 24 | 4 | 0 | Arbitrary-precision fixed-point decimal numbers in go | 2020-10-29 13:57:10 | 2021-05-31 12:29:50 |
| [go-finance](https://www.yellowduck.be) | 5 | 2 | 0 | Finance related Go functions (e.g. exchange rates, VAT number checking, …) | 2019-09-30 06:49:07 | 2021-03-08 22:46:44 |
</details>

### Forms
Libraries for working with forms.

<sup>*Last Update: 2021-06-25 14:25:12*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [nosurf](http://godoc.org/github.com/justinas/nosurf) | 1,165 | 94 | 9 | CSRF protection middleware for Go. | 2013-08-22 17:47:34 | 2021-06-23 06:02:11 |
| [binding](http://mholt.github.io/binding) | 780 | 78 | 8 | Reflectionless data binding for Go's net/http (not actively maintained) | 2014-05-20 23:35:14 | 2021-06-21 14:33:41 |
| [csrf](https://github.com/gorilla/csrf) | 670 | 98 | 1 | gorilla/csrf provides Cross Site Request Forgery (CSRF) prevention middleware for Go web applications & services 🔒 | 2015-08-03 00:35:16 | 2021-06-24 17:54:52 |
| [form](https://github.com/go-playground/form) | 474 | 28 | 10 | :steam_locomotive: Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. | 2016-05-26 13:26:40 | 2021-06-24 17:55:09 |
| [conform](https://github.com/leebenson/conform) | 221 | 29 | 0 | Trims, sanitizes & scrubs data based on struct tags (go, golang) | 2016-01-05 18:00:06 | 2021-06-22 21:45:54 |
| [formam](https://github.com/monoculum/formam) | 161 | 14 | 2 | a package for decode form's values into struct in Go | 2014-10-25 00:23:30 | 2021-06-13 02:30:20 |
| [forms](https://github.com/albrow/forms) | 121 | 16 | 2 | A lightweight go library for parsing form data or json from an http.Request. | 2014-08-07 16:11:30 | 2021-06-13 02:29:50 |
| [qs](https://github.com/sonh/qs) | 57 | 0 | 0 | Go module for encoding structs into URL query parameters | 2020-10-02 09:50:35 | 2021-06-09 20:01:55 |
| [bind](https://github.com/robfig/bind) | 25 | 3 | 0 |  | 2014-08-06 00:13:10 | 2020-08-28 21:26:49 |
| [queryparam](https://github.com/TomWright/queryparam) | 9 | 5 | 0 | Go package to easily convert a URL's query parameters/values into usable struct values of the correct types. | 2018-06-14 10:23:05 | 2021-04-08 21:18:27 |
</details>

### Functional
Packages to support functional programming in Go.

<sup>*Last Update: 2021-07-13 09:25:20*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-underscore](https://github.com/tobyhede/go-underscore) | 1,200 | 65 | 4 |  Helpfully Functional Go -  A useful collection of Go utilities. Designed for programmer happiness.  | 2014-07-02 10:27:16 | 2021-07-10 04:00:28 |
| [fpGo](https://github.com/TeaEntityLab/fpGo) | 179 | 12 | 0 | Monad, Functional Programming features for Golang | 2018-05-24 09:08:45 | 2021-07-08 19:13:34 |
| [fuego](https://github.com/seborama/fuego) | 90 | 8 | 0 | Functional Experiment in Golang | 2018-11-05 22:24:09 | 2021-07-08 19:13:33 |
| [gofp](https://github.com/rbrahul/gofp) | 80 | 3 | 0 | A super simple Lodash like utility library with essential functions that empowers the development in Go | 2021-02-19 00:01:39 | 2021-07-11 14:09:27 |
</details>

### Game Development
Awesome game development libraries.

<sup>*Last Update: 2021-06-25 14:25:06*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ebiten](https://ebiten.org) | 4,686 | 313 | 212 | A dead simple 2D game library for Go | 2013-06-16 15:13:01 | 2021-06-25 06:48:42 |
| [leaf](https://github.com/name5566/leaf) | 4,035 | 1,077 | 9 | A game server framework in Go (golang) | 2014-08-04 12:40:08 | 2021-06-25 03:55:46 |
| [pixel](https://github.com/faiface/pixel) | 3,519 | 205 | 41 | A hand-crafted 2D game library in Go | 2016-11-19 11:15:34 | 2021-06-24 17:40:52 |
| [goworld](https://github.com/xiaonanln/goworld) | 1,862 | 347 | 16 | Scalable Distributed Game Server Engine with Hot Swapping in Golang | 2017-06-03 15:02:46 | 2021-06-25 03:31:49 |
| [nano](https://github.com/lonng/nano) | 1,721 | 287 | 16 | Lightweight, facility, high performance golang based game server framework | 2017-08-02 06:05:14 | 2021-06-25 00:41:41 |
| [go-sdl2](https://godoc.org/github.com/veandco/go-sdl2) | 1,557 | 193 | 51 | SDL2 binding for Go | 2013-06-05 18:30:03 | 2021-06-24 15:21:07 |
| [engine](http://g3n.rocks) | 1,479 | 151 | 68 | Go 3D Game Engine | 2017-03-07 18:25:09 | 2021-06-24 13:05:19 |
| [engo](https://engoengine.github.io) | 1,378 | 114 | 51 | Engo is an open-source 2D game engine written in Go. | 2014-11-12 05:50:03 | 2021-06-21 09:41:39 |
| [termloop](https://github.com/JoelOtter/termloop) | 1,210 | 74 | 5 | Terminal-based game engine for Go, built on top of Termbox | 2015-05-23 17:12:34 | 2021-06-09 08:26:56 |
| [gonet](https://github.com/xtaci/gonet) | 1,142 | 300 | 0 | A Game Server Skeleton in golang. | 2013-04-11 02:18:23 | 2021-06-25 03:33:49 |
| [pitaya](https://github.com/topfreegames/pitaya) | 1,059 | 227 | 22 | Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. | 2018-03-19 19:40:36 | 2021-06-25 06:56:21 |
| [oak](https://github.com/oakmound/oak) | 893 | 58 | 15 | A pure Go game engine | 2017-07-15 16:24:27 | 2021-06-23 14:44:22 |
| [raylib-go](https://github.com/gen2brain/raylib-go) | 612 | 68 | 6 | Go bindings for raylib, a simple and easy-to-use library to enjoy videogames programming. | 2017-01-27 08:31:45 | 2021-06-24 01:51:14 |
| [engine](https://azul3d.org) | 489 | 45 | 82 | Azul3D - A 3D game engine written in Go! | 2016-02-29 04:54:44 | 2021-06-21 20:37:27 |
| [go-astar](http://g3n.rocks) | 430 | 58 | 1 | Go implementation of the A* search algorithm | 2014-05-28 02:00:03 | 2021-06-22 01:42:56 |
| [go3d](https://github.com/ungerik/go3d) | 200 | 36 | 2 | A performance oriented 2D/3D math package for Go | 2011-06-27 13:02:26 | 2021-06-24 06:18:57 |
| [prototype](https://github.com/gonutz/prototype) | 56 | 4 | 1 | Simple 2D game prototyping framework. | 2015-03-04 09:24:39 | 2021-06-12 19:53:11 |
| [tile](https://github.com/kelindar/tile) | 24 | 1 | 0 | Tile is a 2D grid engine, built with data and cache friendly ways, includes pathfinding and observers. | 2020-08-19 13:23:18 | 2021-06-16 20:36:45 |
</details>

### Generation and Generics
Tools to enhance the language with features like generics via code generation.

<sup>*Last Update: 2021-06-25 10:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-linq](https://godoc.org/github.com/ahmetb/go-linq) | 2,543 | 178 | 6 | .NET LINQ capabilities in Go | 2013-12-19 03:05:00 | 2021-06-24 03:21:01 |
| [jennifer](https://github.com/dave/jennifer) | 2,057 | 104 | 16 | Jennifer is a code generator for Go | 2016-12-04 20:57:38 | 2021-06-24 04:07:17 |
| [gen](http://clipperhouse.com/gen/overview/) | 1,295 | 82 | 32 | Type-driven code generation for Go | 2013-10-13 20:26:36 | 2021-06-24 06:32:30 |
| [goderive](https://github.com/awalterschulze/goderive) | 890 | 35 | 14 | Code Generation for Functional Programming, Concurrency and Generics in Golang | 2017-02-10 21:46:49 | 2021-06-23 15:53:39 |
| [gowrap](https://github.com/hexdigest/gowrap) | 516 | 43 | 2 | GoWrap is a command line tool for generating decorators for Go interfaces | 2018-09-15 09:20:42 | 2021-06-22 22:51:05 |
| [interfaces](https://github.com/rjeczalik/interfaces) | 300 | 20 | 5 | Code generation tools for Go. | 2015-12-06 00:04:50 | 2021-06-15 11:36:03 |
| [go-enum](https://github.com/abice/go-enum) | 209 | 26 | 2 | An enum generator for go | 2017-08-10 22:07:31 | 2021-06-24 14:47:50 |
| [pkgreflect](https://github.com/ungerik/pkgreflect) | 96 | 15 | 0 | A Go preprocessor for package scoped reflection | 2012-09-03 07:53:00 | 2021-05-23 02:33:23 |
| [efaceconv](https://github.com/t0pep0/efaceconv) | 48 | 9 | 1 |  | 2016-11-18 11:38:54 | 2021-06-13 20:28:39 |
| [gotype](https://github.com/wzshiming/gotype) | 37 | 5 | 0 | Golang source code parsing, usage like reflect package | 2017-12-05 04:09:47 | 2021-06-12 18:01:28 |
| [GENERIS](https://github.com/senselogic/GENERIS) | 27 | 0 | 0 | Versatile Go code generator. | 2019-03-10 19:33:31 | 2021-06-13 20:29:04 |
| [go-xray](https://godoc.org/github.com/ahmetb/go-linq) | 17 | 1 | 0 | Helpers for making the use of reflection easier | 2019-10-01 08:40:51 | 2021-03-16 01:44:08 |
| [typeregistry](https://github.com/xiaoxin01/typeregistry) | 11 | 0 | 0 | create type dynamically in Golang | 2020-01-14 15:50:38 | 2021-06-14 08:05:42 |
</details>

### Geographic
Geographic tools and servers

<sup>*Last Update: 2021-06-28 16:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tile38](https://tile38.com) | 7,483 | 449 | 122 | Real-time Geospatial and Geofencing | 2016-03-04 23:07:44 | 2021-06-27 22:12:47 |
| [geo](https://github.com/golang/geo) | 1,230 | 138 | 9 | S2 geometry library in Go | 2014-12-03 23:02:15 | 2021-06-24 21:57:59 |
| [mbtileserver](https://github.com/consbio/mbtileserver) | 246 | 42 | 9 | Basic Go server for mbtiles | 2014-11-01 04:12:14 | 2021-06-27 00:58:11 |
| [osm](https://github.com/paulmach/osm) | 167 | 27 | 2 | General purpose library for reading, writing and working with OpenStreetMap data | 2016-02-02 00:59:03 | 2021-06-15 04:25:04 |
| [wgs84](https://github.com/wroge/wgs84) | 62 | 3 | 0 | A pure Go package for coordinate transformations. | 2019-06-08 17:17:59 | 2021-06-24 10:30:51 |
| [geoserver](https://github.com/hishamkaram/geoserver) | 55 | 11 | 3 | geoserver is a Go library for manipulating a GeoServer instance via the GeoServer REST API. | 2018-03-26 21:36:49 | 2021-06-25 07:14:25 |
| [gismanager](https://github.com/hishamkaram/gismanager) | 38 | 6 | 0 | Publish Your GIS Data(Vector Data) to PostGIS and Geoserver | 2018-09-29 12:51:37 | 2021-03-02 15:54:38 |
| [pbf](https://github.com/maguro/pbf) | 26 | 3 | 1 | OpenStreetMap PBF golang parser | 2017-09-18 23:13:18 | 2021-05-26 02:13:09 |
| [s2-geojson](https://github.com/pantrif/s2-geojson) | 13 | 3 | 1 | Draw a polygon on the map or paste a geoJSON and explore how the s2.RegionCoverer covers it with S2 cells depending on the min and max levels | 2020-03-27 09:47:32 | 2021-06-13 06:12:30 |
</details>

### Go Compilers
Tools for compiling Go to other languages.

<sup>*Last Update: 2021-06-24 16:25:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopherjs](https://github.com/gopherjs/gopherjs) | 10,257 | 474 | 212 | A compiler from Go to JavaScript for running Go code in a browser | 2013-08-27 22:23:58 | 2021-06-24 07:17:09 |
| [tardisgo](http://tardisgo.github.io) | 406 | 25 | 4 | Golang->Haxe->CPP/CSharp/Java/JavaScript transpiler   | 2014-01-08 11:07:33 | 2021-05-27 06:23:26 |
| [c4go](https://github.com/Konstantin8105/c4go) | 267 | 30 | 23 | Transpiling C code to Go code | 2018-03-28 06:24:57 | 2021-06-23 14:47:26 |
| [f4go](https://github.com/Konstantin8105/f4go) | 24 | 5 | 4 | Transpiling fortran code to golang code | 2018-07-08 17:05:43 | 2021-06-15 21:53:09 |
</details>

### Goroutines
Tools for managing and working with Goroutines.

<sup>*Last Update: 2021-06-29 14:34:35*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ants](https://ants.andypan.me) | 5,864 | 732 | 16 | 🐜🐜🐜 ants is a high-performance and low-cost goroutine pool in Go, inspired by fasthttp./ ants 是一个高性能且低损耗的 goroutine 池。 | 2018-05-19 01:13:38 | 2021-06-29 06:51:57 |
| [goworker](https://www.goworker.org) | 2,568 | 238 | 32 | goworker is a Go-based background worker that runs 10 to 100,000* times faster than Ruby-based workers. | 2013-07-22 17:04:27 | 2021-06-29 02:22:30 |
| [tunny](https://github.com/Jeffail/tunny) | 2,401 | 210 | 0 | A goroutine pool for Go | 2014-04-02 16:14:58 | 2021-06-29 06:49:28 |
| [grpool](https://godoc.org/github.com/ivpusic/grpool) | 637 | 93 | 4 | Lightweight Goroutine pool | 2015-07-22 00:15:04 | 2021-06-29 03:46:42 |
| [pool](https://github.com/go-playground/pool) | 617 | 58 | 4 | :speedboat: a limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation | 2015-10-28 16:36:08 | 2021-06-23 09:32:12 |
| [workerpool](https://github.com/gammazero/workerpool) | 569 | 72 | 4 | Concurrency limiting goroutine pool | 2016-05-17 14:32:06 | 2021-06-27 13:55:24 |
| [gowp](https://xxjwxc.github.io/post/gowp/) | 300 | 50 | 0 | golang worker pool , Concurrency limiting goroutine pool | 2019-09-14 11:43:50 | 2021-06-28 09:18:59 |
| [pond](https://github.com/alitto/pond) | 224 | 8 | 0 | Minimalistic and High-performance goroutine worker pool written in Go | 2020-03-21 14:56:33 | 2021-06-27 23:13:12 |
| [go-floc](https://github.com/workanator/go-floc) | 202 | 15 | 0 | Floc: Orchestrate goroutines with ease. | 2017-07-03 07:34:06 | 2021-06-15 02:27:08 |
| [go-flow](https://github.com/kamildrazkiewicz/go-flow) | 167 | 16 | 1 | Simply way to control goroutines execution order based on dependencies | 2016-09-25 14:46:09 | 2021-06-11 03:22:35 |
| [semaphore](https://github.com/marusama/semaphore) | 115 | 7 | 0 | Fast resizable golang semaphore primitive | 2017-11-22 14:00:58 | 2021-06-15 08:31:02 |
| [go-workers](https://github.com/catmullet/go-workers/wiki/Getting-Started) | 108 | 5 | 0 | 👷 Library for safely running groups of workers concurrently or consecutively that require input and output through channels | 2020-10-06 15:39:43 | 2021-06-22 14:02:04 |
| [artifex](https://github.com/mborders/artifex) | 107 | 8 | 0 | Simple in-memory job queue for Golang using worker-based dispatching | 2018-10-31 19:34:31 | 2021-06-18 08:30:12 |
| [breaker](https://pkg.go.dev/github.com/kamilsk/breaker) | 97 | 5 | 7 | 🚧 Flexible mechanism to make execution flow interruptible. | 2019-02-15 15:08:24 | 2021-06-27 05:54:10 |
| [semaphore](https://github.com/kamilsk/semaphore) | 85 | 10 | 6 | 🚦 Semaphore pattern implementation with timeout of lock/unlock operations. | 2016-10-08 11:48:12 | 2021-04-18 23:49:16 |
| [async](https://github.com/StudioSol/async) | 83 | 9 | 2 | A safe way to execute functions asynchronously, recovering them in case of panic. It also provides an error stack aiming to facilitate fail causes discovery. | 2017-06-30 17:08:33 | 2021-06-25 01:12:56 |
| [errgroup](https://github.com/neilotoole/errgroup) | 80 | 6 | 4 | errgroup with goroutine worker limits | 2020-06-26 06:07:39 | 2021-06-27 15:01:23 |
| [gpool](https://github.com/sherifabdlnaby/gpool) | 79 | 3 | 0 | gpool - a generic context-aware resizable goroutines pool to bound concurrency based on semaphore.  | 2018-12-03 04:23:35 | 2021-06-13 10:24:46 |
| [worker-pool](https://github.com/vardius/worker-pool) | 75 | 12 | 0 | Go simple async worker pool | 2017-10-04 09:18:31 | 2021-06-27 07:24:49 |
| [cyclicbarrier](https://github.com/marusama/cyclicbarrier) | 67 | 8 | 0 | CyclicBarrier golang implementation | 2018-01-11 10:38:46 | 2021-06-11 06:33:33 |
| [threadpool](https://github.com/shettyh/threadpool) | 60 | 14 | 0 | Golang simple thread pool implementation | 2017-09-06 18:45:39 | 2021-06-06 08:02:33 |
| [gollback](https://github.com/vardius/gollback) | 58 | 4 | 0 | Go asynchronous simple function utilities, for managing execution of closures and callbacks | 2019-05-11 05:56:37 | 2021-06-24 09:44:46 |
| [Hunch](https://github.com/AaronJan/Hunch) | 54 | 4 | 0 | Hunch provides functions like: All, First, Retry, Waterfall etc., that makes asynchronous flow control more intuitive. | 2019-06-05 13:21:04 | 2021-06-25 10:01:43 |
| [routine](https://github.com/x-mod/routine) | 41 | 3 | 0 | go routine control, abstraction of the Main and some useful Executors.如果你不会管理Goroutine的话，用它 | 2019-03-04 12:25:23 | 2021-06-08 09:53:07 |
| [kyoo](https://github.com/dirkaholic/kyoo) | 32 | 0 | 0 | Unlimited job queue for go, using a pool of concurrent workers processing the job queue entries | 2020-01-06 20:35:11 | 2021-02-07 16:41:12 |
| [parallel-fn](https://github.com/rafaeljesus/parallel-fn) | 30 | 1 | 0 | Run functions in parallel :comet: | 2017-06-18 09:47:54 | 2021-05-23 15:19:32 |
| [nursery](https://github.com/arunsworld/nursery) | 30 | 3 | 1 | Structured Concurrency in Go | 2019-11-23 19:26:02 | 2021-05-31 08:32:27 |
| [async](https://pkg.go.dev/github.com/reugn/async) | 23 | 1 | 0 | Alternative sync library for Go | 2019-12-28 09:48:40 | 2021-06-09 09:44:48 |
| [oversight](https://cirello.io/oversight) | 22 | 3 | 0 | [Mirror] Erlang-like supervisor trees | 2018-11-09 14:46:48 | 2021-04-29 15:26:23 |
| [goccm](https://github.com/zenthangplus/goccm) | 21 | 3 | 1 | Limits the number of goroutines that are allowed to run concurrently | 2019-08-16 02:26:53 | 2021-06-06 16:15:27 |
| [go-waitgroup](https://www.yellowduck.be) | 20 | 1 | 0 | A sync.WaitGroup with error handling and concurrency control | 2018-08-08 16:12:35 | 2021-06-19 10:14:26 |
| [stl](https://github.com/ssgreg/stl) | 18 | 3 | 0 | Software Transactional Locks | 2018-06-19 10:50:11 | 2020-11-19 10:08:47 |
| [go-trylock](https://github.com/subchen/go-trylock) | 18 | 6 | 0 | TryLock support on read-write lock for Golang | 2018-04-26 06:02:47 | 2021-05-21 04:31:56 |
| [channelify](https://github.com/ddelizia/channelify) | 12 | 1 | 1 | Make functions return a channel for parallel processing via go routines. | 2020-10-05 13:12:48 | 2021-05-16 18:07:10 |
| [gohive](https://github.com/loveleshsharma/gohive) | 12 | 1 | 1 | 🐝 A Highly Performant and easy to use goroutine pool for Go | 2019-05-31 10:44:24 | 2021-06-21 17:28:15 |
| [conexec](https://github.com/ITcathyh/conexec) | 10 | 1 | 0 | A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking. | 2019-12-24 07:35:11 | 2021-04-22 08:24:00 |
| [queue](https://github.com/AnikHasibul/queue) | 9 | 1 | 0 | package queue gives you a queue group accessibility. Helps you to limit goroutines, wait for the end of the all goroutines and much more. | 2018-12-21 07:15:00 | 2020-11-02 15:19:05 |
| [gowl](https://github.com/hamed-yousefi/gowl) | 9 | 2 | 5 | Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. | 2021-04-12 19:15:53 | 2021-06-10 23:49:20 |
| [hands](https://github.com/duanckham/hands) | 7 | 1 | 1 | Hands is a process controller used to control the execution and return strategies of multiple goroutines. | 2020-04-04 11:04:11 | 2021-06-24 01:19:23 |
| [go-tools](https://github.com/nikhilsaraf/go-tools) | 5 | 2 | 0 | A collection of tools for Golang | 2018-11-14 02:53:08 | 2019-05-04 06:52:40 |
| [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) | 4 | 1 | 0 | Make functions return a channel for parallel processing via go routines. | 2020-11-22 02:35:52 | 2020-12-04 21:15:03 |
</details>

### Images
Libraries for manipulating images.

<sup>*Last Update: 2021-07-01 10:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocv](https://gocv.io) | 4,122 | 608 | 167 | Go package for computer vision using OpenCV 4 and beyond. | 2017-09-18 21:54:17 | 2021-07-01 00:48:37 |
| [imaging](https://github.com/disintegration/imaging) | 3,799 | 319 | 10 | Imaging is a simple image processing package for Go | 2012-12-06 20:21:21 | 2021-06-30 21:14:01 |
| [imaginary](https://fly.io/launch/github/h2non/imaginary) | 3,708 | 351 | 93 | Fast, simple, scalable, Docker-ready HTTP microservice for high-level image processing | 2015-03-04 18:51:40 | 2021-06-30 20:53:12 |
| [bild](https://github.com/anthonynsimon/bild) | 3,235 | 166 | 12 | Image processing algorithms in pure Go | 2016-08-01 15:54:29 | 2021-06-30 16:06:01 |
| [ln](https://godoc.org/github.com/fogleman/ln/ln) | 2,914 | 108 | 12 | 3D line art engine. | 2016-01-10 04:28:10 | 2021-06-28 01:17:31 |
| [gg](https://godoc.org/github.com/fogleman/gg) | 2,909 | 207 | 51 | Go Graphics - 2D rendering in Go with a simple API. | 2016-02-18 21:05:08 | 2021-06-30 08:34:21 |
| [resize](https://github.com/nfnt/resize) | 2,655 | 260 | 9 | Pure golang image resizing  | 2012-08-02 19:48:26 | 2021-06-28 12:06:10 |
| [pt](http://bit.ly/1E7rSoi) | 1,953 | 112 | 8 | A path tracer written in Go. | 2015-01-23 19:39:29 | 2021-06-28 01:21:21 |
| [svgo](https://github.com/ajstarks/svgo) | 1,680 | 140 | 6 | Go Language Library for SVG generation | 2010-03-05 23:24:10 | 2021-06-28 01:24:31 |
| [smartcrop](https://github.com/muesli/smartcrop) | 1,524 | 105 | 7 | smartcrop finds good image crops for arbitrary crop sizes | 2014-04-07 22:40:03 | 2021-06-28 01:23:37 |
| [picfit](http://bit.ly/1E7rSoi) | 1,488 | 126 | 16 | An image resizing server written in Go | 2014-12-06 17:30:45 | 2021-06-29 21:44:20 |
| [bimg](https://pkg.go.dev/github.com/h2non/bimg?tab=doc) | 1,484 | 251 | 119 | Go package for fast high-level image processing powered by libvips C library | 2015-03-17 14:14:02 | 2021-07-01 00:24:21 |
| [gift](https://godoc.org/github.com/fogleman/gg) | 1,449 | 104 | 2 | Go Image Filtering Toolkit | 2014-07-12 18:47:40 | 2021-06-28 15:58:30 |
| [imagick](https://godoc.org/github.com/gographics/imagick/imagick) | 1,321 | 156 | 11 | Go binding to ImageMagick's MagickWand C API | 2013-04-30 17:31:48 | 2021-06-28 18:21:30 |
| [go-opencv](https://github.com/go-opencv/go-opencv) | 1,234 | 196 | 45 | Go bindings for OpenCV / 2.x API in gocv / 1.x API in opencv | 2013-12-09 09:43:26 | 2021-06-05 07:04:50 |
| [geopattern](https://github.com/pravj/geopattern) | 1,126 | 58 | 3 | :triangular_ruler: Create beautiful generative image patterns from a string in golang. | 2014-10-22 17:26:30 | 2021-06-28 01:04:36 |
| [stegify](https://github.com/DimitarPetrov/stegify) | 936 | 106 | 0 | 🔍 Go tool for LSB steganography, capable of hiding any file within an image. | 2018-11-29 16:45:58 | 2021-06-28 01:24:10 |
| [canvas](https://github.com/tdewolff/canvas) | 731 | 38 | 8 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2017-05-20 18:10:51 | 2021-06-29 18:42:25 |
| [image2ascii](https://github.com/qeesung/image2ascii) | 551 | 47 | 3 | :foggy: Convert image to ASCII | 2018-10-20 05:06:25 | 2021-06-28 01:14:09 |
| [draft](https://github.com/lucasepe/draft) | 494 | 19 | 0 | Generate High Level Cloud Architecture diagrams using YAML syntax. | 2020-06-05 16:11:40 | 2021-07-01 03:18:48 |
| [govips](https://github.com/davidbyttow/govips) | 486 | 106 | 12 | A lightning fast image processing and resizing library for Go | 2016-12-25 04:32:56 | 2021-06-24 06:32:46 |
| [govatar](https://github.com/o1egl/govatar) | 442 | 24 | 0 | Avatar generation library for GO language | 2016-01-18 12:12:28 | 2021-06-28 01:12:10 |
| [mort](https://github.com/aldor007/mort) | 436 | 15 | 6 | Storage and image processing server written in Go | 2017-11-19 13:37:58 | 2021-06-28 01:19:25 |
| [goimagehash](https://github.com/corona10/goimagehash) | 413 | 44 | 8 | Go Perceptual image hashing package | 2017-07-28 17:15:58 | 2021-06-30 09:52:23 |
| [go-nude](https://github.com/koyachi/go-nude) | 333 | 38 | 2 | Nudity detection with Go. | 2014-05-02 08:32:29 | 2021-06-28 01:07:41 |
| [rez](https://github.com/bamiaux/rez) | 200 | 16 | 1 | Image resizing in pure Go and SIMD | 2014-01-16 21:16:15 | 2021-06-28 01:22:33 |
| [darkroom](https://www.gojek.io/darkroom/) | 164 | 32 | 7 | Cairo in Go: vector to raster, SVG, PDF, EPS, WASM, OpenGL, Gio, etc. | 2019-07-01 10:17:08 | 2021-06-16 02:11:52 |
| [mergi](http://mergi.io/) | 139 | 18 | 2 | go library for image programming (merge, crop, resize, watermark, animate, ease, transit) | 2018-09-24 03:40:47 | 2021-06-28 01:18:06 |
| [img](hawx.me/code/img) | 137 | 9 | 1 | A selection of image manipulation tools | 2012-07-28 19:57:47 | 2021-06-28 01:16:41 |
| [gltf](https://www.khronos.org/gltf/) | 111 | 17 | 2 | :eyeglasses: Go library for [d]encoding glTF 2.0 files | 2019-01-15 17:43:54 | 2021-06-27 10:10:17 |
| [go-cairo](https://github.com/ungerik/go-cairo) | 106 | 28 | 0 | Go binding for the cairo graphics library | 2012-08-22 18:27:01 | 2021-06-28 01:06:50 |
| [steganography](https://github.com/auyer/steganography) | 101 | 17 | 0 | Pure Golang Library that allows simple LSB steganography on images | 2018-05-21 17:27:36 | 2021-06-14 10:48:50 |
| [cameron](https://pkg.go.dev/github.com/aofei/cameron) | 69 | 7 | 1 | An avatar generator for Go. | 2018-05-05 22:13:11 | 2021-06-06 02:26:18 |
| [go-gd](https://github.com/bolknote/go-gd) | 53 | 15 | 1 | Go bingings for GD (http://www.boutell.com/gd/) | 2011-05-12 06:33:54 | 2021-06-29 11:02:01 |
| [gridder](https://github.com/shomali11/gridder) | 42 | 3 | 0 | A Grid based 2D Graphics library | 2020-04-10 00:13:10 | 2021-06-28 01:13:55 |
| [goimghdr](https://github.com/corona10/goimghdr) | 36 | 2 | 0 | The imghdr module determines the type of image contained in a file for go | 2018-02-25 09:34:44 | 2021-06-28 01:11:03 |
| [tga](https://github.com/ftrvxmtrx/tga) | 28 | 10 | 1 | Go package for decoding and encoding TARGA image format | 2012-10-08 01:09:20 | 2021-06-28 01:25:28 |
| [go-webcolors](https://github.com/jyotiska/go-webcolors) | 25 | 4 | 0 | Port of webcolors library from Python to Go | 2014-04-24 14:41:22 | 2021-06-28 01:09:17 |
| [webp-server](https://github.com/mehdipourfar/webp-server) | 16 | 2 | 0 | Simple and minimal image server capable of storing, resizing, converting and caching images. | 2020-11-22 12:03:12 | 2021-06-28 08:04:37 |
| [mpo](https://donatstudios.com/MPO-to-JPEG-Stereo) | 7 | 2 | 1 | JPEG-MPO Decoder / Converter Library and CLI Tool | 2015-04-14 22:37:59 | 2021-06-28 01:20:05 |
</details>

### IoT (Internet of Things)
Libraries for programming devices of the IoT.

<sup>*Last Update: 2021-06-25 12:39:49*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gobot](https://gobot.io) | 7,204 | 904 | 163 | Golang framework for robotics, drones, and the Internet of Things (IoT) | 2013-09-21 14:09:19 | 2021-06-25 03:06:28 |
| [flogo](http://flogo.io) | 1,767 | 239 | 153 | Project Flogo is an open source ecosystem of opinionated  event-driven capabilities to simplify building efficient & modern serverless functions, microservices & edge apps. | 2016-07-10 02:57:43 | 2021-06-24 13:13:41 |
| [periph](https://periph.io) | 1,662 | 175 | 42 | Go·Hardware·Lean | 2016-10-13 16:53:51 | 2021-06-22 17:19:07 |
| [mainflux](https://www.mainflux.io) | 1,444 | 434 | 92 | Industrial IoT Messaging and Device Management Platform | 2015-07-06 20:31:50 | 2021-06-24 01:48:02 |
| [gatt](http://flogo.io) | 978 | 260 | 51 | Gatt is a Go package for building Bluetooth Low Energy peripherals | 2014-04-23 13:45:27 | 2021-06-18 01:24:00 |
| [heedy](https://heedy.org) | 276 | 27 | 7 | An aggregator for personal metrics, and an extensible analysis engine | 2015-01-16 19:44:21 | 2021-06-22 20:36:43 |
| [devices](https://github.com/goiot/devices) | 241 | 24 | 9 | Suite of libraries for IoT devices (written in Go), experimental for x/exp/io | 2016-05-30 08:07:02 | 2021-05-25 19:15:47 |
| [sensorbee](http://sensorbee.io/) | 203 | 33 | 39 | Lightweight stream processing engine for IoT | 2016-02-19 07:49:56 | 2021-05-19 21:00:03 |
| [huego](https://github.com/amimof/huego) | 182 | 28 | 3 | An extensive Philips Hue client library for Go with an emphasis on simplicity | 2017-05-16 05:31:45 | 2021-06-11 14:38:11 |
| [iot](https://github.com/vaelen/iot) | 52 | 7 | 0 | A Go client for Google IoT Core | 2018-03-08 06:51:51 | 2021-04-30 18:44:01 |
| [eywa](https://github.com/xcodersun/eywa) | 48 | 11 | 9 | Make IoT a lot more fun with data.  | 2016-02-20 17:02:16 | 2021-03-19 19:55:56 |
</details>

### JSON
Libraries for working with JSON.

<sup>*Last Update: 2021-06-24 16:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gjson](https://github.com/tidwall/gjson) | 8,505 | 577 | 25 | Get JSON values quickly - JSON parser for Go | 2016-08-11 03:08:47 | 2021-06-24 08:12:47 |
| [json-to-go](https://mholt.github.io/json-to-go/) | 3,123 | 367 | 12 | Translates JSON into a Go type in your browser instantly (original) | 2014-01-21 18:11:13 | 2021-06-24 09:22:11 |
| [gojson](https://github.com/ChimeraCoder/gojson) | 2,369 | 176 | 39 | Automatically generate Go (golang) struct definitions from example JSON | 2012-12-27 19:10:50 | 2021-06-21 14:32:35 |
| [fastjson](https://github.com/valyala/fastjson) | 1,252 | 67 | 31 | Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection | 2018-05-28 21:41:47 | 2021-06-22 12:29:00 |
| [kazaam](https://github.com/qntfy/kazaam) | 192 | 41 | 24 | Arbitrary transformations of JSON in Golang | 2016-07-19 14:19:03 | 2021-06-11 12:18:35 |
| [gojq](https://godoc.org/github.com/nicklaw5/go-respond) | 170 | 19 | 1 | JSON query in Golang | 2015-12-30 09:02:13 | 2021-05-19 19:33:24 |
| [jsondiff](https://pkg.go.dev/github.com/wI2L/jsondiff) | 118 | 10 | 0 | JSON diff library for Go based on RFC6902 (JSON Patch) | 2020-11-28 19:05:16 | 2021-06-07 08:36:16 |
| [jettison](https://pkg.go.dev/github.com/wI2L/jettison) | 108 | 6 | 0 | Fast and flexible JSON encoder for Go | 2019-08-30 13:28:03 | 2021-06-20 21:26:53 |
| [jsongo](https://pkg.go.dev/github.com/wI2L/jsondiff) | 96 | 13 | 1 | Fluent API to make it easier to create Json objects. | 2015-08-07 23:23:17 | 2021-05-01 10:53:24 |
| [gjo](https://github.com/skanehira/gjo) | 94 | 11 | 1 | Small utility to create JSON objects | 2019-02-23 01:54:21 | 2021-05-21 04:32:02 |
| [jaydiff](https://github.com/yazgazan/jaydiff) | 85 | 7 | 2 | A JSON diff utility | 2017-04-24 16:05:35 | 2021-05-13 13:34:51 |
| [json2go](https://m-zajac.github.io/json2go) | 82 | 12 | 1 | Create go type representation from json | 2017-06-10 23:55:07 | 2021-06-16 14:20:55 |
| [jsonf](https://pkg.go.dev/github.com/wI2L/jsondiff) | 61 | 9 | 0 | Console JSON formatter with query feature | 2015-05-25 04:53:32 | 2021-04-24 08:25:31 |
| [ajson](https://github.com/spyzhov/ajson) | 61 | 7 | 5 | Abstract JSON for golang with JSONPath support  | 2019-03-07 20:47:38 | 2021-06-22 06:02:34 |
| [mp](https://github.com/sanbornm/mp) | 44 | 4 | 1 | Simple Email Parser | 2014-06-15 21:14:39 | 2020-10-22 17:21:09 |
| [go-respond](https://godoc.org/github.com/nicklaw5/go-respond) | 41 | 5 | 0 | A Go package for handling common HTTP JSON responses. | 2017-03-12 21:00:54 | 2021-03-09 14:00:56 |
| [json-to-proto.github.io](https://json-to-proto.github.io/) | 40 | 5 | 0 | convert JSON to Protocol Buffers online in your browser instantly | 2020-04-18 20:42:45 | 2021-06-14 02:09:52 |
| [go-jsonerror](https://github.com/ddymko/go-jsonerror) | 10 | 1 | 0 | Small package which wraps error responses to follow jsonapi.org | 2018-10-18 15:03:45 | 2020-10-02 04:02:28 |
| [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) | 9 | 2 | 0 | Go bindings based on the JSON API errors reference | 2016-07-08 10:08:58 | 2020-08-30 14:39:42 |
| [jsonhal](https://github.com/RichardKnop/jsonhal) | 9 | 4 | 1 | A simple Go package to make custom structs marshal into HAL compatible JSON responses. | 2016-01-15 11:38:40 | 2020-05-26 19:24:02 |
| [ask](https://github.com/simonnilsson/ask) | 7 | 0 | 0 | A Go package that provides a simple way of accessing nested properties in maps and slices. | 2020-09-13 13:53:31 | 2021-05-29 00:44:25 |
| [ej](https://github.com/lucassscaravelli/ej) | 7 | 0 | 0 | Write and read JSON from different sources in one line | 2020-01-04 17:39:35 | 2021-04-27 22:00:09 |
| [dynjson](https://github.com/cocoonspace/dynjson) | 6 | 1 | 0 | Client-customizable JSON formats for dynamic APIs | 2020-05-06 07:10:02 | 2021-06-09 06:01:06 |
| [epoch](https://github.com/vtopc/epoch) | 5 | 1 | 0 | Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from built-in time.Time type in JSON | 2019-12-15 12:54:37 | 2021-06-09 06:01:11 |
| [jzon](https://github.com/zerosnake0/jzon) | 4 | 0 | 0 | A golang json library inspired by jsoniter | 2019-11-12 10:42:41 | 2021-05-24 05:31:48 |
| [mapslice-json](https://github.com/mickep76/mapslice-json) | 4 | 1 | 0 | Go MapSlice for ordered marshal/ unmarshal of maps in JSON | 2020-02-19 11:01:48 | 2021-04-20 04:09:25 |
| [jsonic](https://github.com/sinhashubham95/jsonic) | 2 | 0 | 0 | All you need with JSON | 2021-01-09 06:21:59 | 2021-03-05 04:01:32 |
</details>

### Job Scheduler
Libraries for scheduling jobs.

<sup>*Last Update: 2021-06-24 08:56:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gocron](https://github.com/go-co-op/gocron) | 889 | 78 | 10 | Easy and fluent Go cron scheduling. This is a fork from https://github.com/jasonlvhit/gocron | 2020-03-20 15:33:05 | 2021-06-24 00:15:13 |
| [gron](https://github.com/roylee0704/gron) | 862 | 52 | 8 | gron, Cron Jobs in Go. | 2016-06-04 08:02:22 | 2021-06-17 09:41:34 |
| [jobrunner](https://github.com/bamzi/jobrunner) | 853 | 73 | 10 | Framework for performing work asynchronously, outside of the request flow | 2015-10-21 04:17:01 | 2021-06-20 23:52:04 |
| [jobs](https://github.com/albrow/jobs) | 482 | 39 | 17 | A persistent and flexible background jobs library for go. | 2015-02-09 22:13:29 | 2021-05-31 17:24:28 |
| [scheduler](https://github.com/carlescere/scheduler) | 368 | 49 | 6 | Job scheduling made easy. | 2015-02-03 17:10:23 | 2021-06-16 20:42:40 |
| [go-cron](https://github.com/rk/go-cron) | 202 | 16 | 0 | A simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. | 2011-04-15 14:50:49 | 2021-06-13 02:29:27 |
| [go-quartz](https://pkg.go.dev/github.com/reugn/go-quartz/quartz) | 144 | 12 | 0 | Simple, zero-dependency scheduling library for Go | 2019-04-14 18:57:51 | 2021-06-20 09:39:03 |
| [clockwerk](https://github.com/onatm/clockwerk) | 107 | 10 | 0 | Job Scheduling Library | 2017-04-09 23:10:48 | 2021-06-18 08:30:15 |
| [leprechaun](https://github.com/kilgaloon/leprechaun) | 82 | 11 | 13 | You had one job, or more then one, which can be done in steps | 2018-04-08 13:44:04 | 2021-04-04 05:43:31 |
| [tasks](https://github.com/madflojo/tasks) | 45 | 3 | 1 | Package tasks is an easy to use in-process scheduler for recurring tasks in Go | 2019-12-24 18:26:18 | 2021-06-08 09:12:47 |
| [clockwork](https://github.com/whiteShtef/clockwork) | 28 | 12 | 2 | Job Scheduling Library | 2020-02-21 01:25:57 | 2021-06-19 03:47:34 |
| [cronticker](https://github.com/krayzpipes/cronticker) | 1 | 0 | 0 | Golang ticker that works with Cron scheduling. | 2020-11-28 20:59:38 | 2021-01-02 01:57:07 |
</details>

### Logging
Libraries for generating and working with log files.

<sup>*Last Update: 2021-06-23 16:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [logrus](https://github.com/sirupsen/logrus) | 18,080 | 1,961 | 119 | Structured, pluggable logging for Go. | 2013-10-16 19:08:55 | 2021-06-23 08:41:11 |
| [zap](https://godoc.org/go.uber.org/zap) | 12,885 | 972 | 90 | Blazing fast, structured, leveled logging in Go. | 2016-02-18 19:52:56 | 2021-06-23 07:41:52 |
| [zerolog](https://github.com/rs/zerolog) | 4,825 | 292 | 68 | Zero Allocation JSON Logger | 2017-05-12 05:24:39 | 2021-06-23 07:36:33 |
| [go-spew](https://github.com/davecgh/go-spew) | 4,465 | 290 | 55 | Implements a deep pretty printer for Go data structures to aid in debugging | 2013-01-09 05:18:22 | 2021-06-22 17:51:21 |
| [glog](https://github.com/golang/glog) | 2,734 | 774 | 8 | Leveled execution logs for Go | 2013-07-16 04:33:04 | 2021-06-21 14:24:36 |
| [lumberjack](https://github.com/natefinch/lumberjack) | 2,596 | 348 | 51 | lumberjack is a log rolling package for Go | 2014-06-14 11:55:47 | 2021-06-23 07:51:36 |
| [tail](https://github.com/hpcloud/tail) | 2,117 | 426 | 70 | Go package for reading from continously updated files (tail -f) | 2013-02-05 00:28:03 | 2021-06-22 13:10:10 |
| [seelog](https://github.com/cihub/seelog) | 1,543 | 242 | 39 | Seelog is a native Go logging library that provides flexible asynchronous dispatching, filtering, and formatting. | 2011-11-17 09:43:15 | 2021-06-20 03:57:09 |
| [log](https://github.com/apex/log) | 1,144 | 95 | 34 | Structured logging package for Go. | 2015-12-21 20:27:48 | 2021-06-21 08:08:33 |
| [log15](https://godoc.org/github.com/inconshreveable/log15) | 1,012 | 137 | 43 | Structured, composable logging for Go | 2014-05-20 00:11:52 | 2021-06-21 08:08:28 |
| [onelog](https://github.com/francoispqt/onelog) | 394 | 14 | 1 | Dead simple, super fast, zero allocation and modular logger for Golang | 2018-05-06 14:32:10 | 2021-06-12 08:42:40 |
| [log](https://github.com/phuslu/log) | 375 | 23 | 3 | Structured Logging Made Easy | 2019-07-07 09:40:38 | 2021-06-22 18:09:50 |
| [logxi](https://logur.dev/logur) | 347 | 39 | 24 | A 12-factor app logger built for performance and happy development | 2015-03-01 22:13:45 | 2021-05-08 17:22:18 |
| [logutils](https://logur.dev/logur) | 292 | 31 | 1 | Utilities for slightly better logging in Go (Golang). | 2013-10-09 07:31:15 | 2021-03-24 13:59:46 |
| [log](https://github.com/go-playground/log) | 275 | 21 | 1 | :green_book: Simple, configurable and scalable Structured Logging for Go. | 2016-02-07 16:17:48 | 2021-05-06 17:20:44 |
| [go-logger](https://github.com/apsdehal/go-logger) | 263 | 49 | 2 | Simple logger for Go programs. Allows custom formats for messages. | 2014-09-26 04:57:06 | 2021-06-18 03:21:21 |
| [httpretty](https://asciinema.org/a/297429) | 228 | 6 | 1 | Package httpretty prints the HTTP requests you make with Go pretty on your terminal. | 2020-01-24 18:17:16 | 2021-06-10 04:17:21 |
| [sqldb-logger](https://github.com/simukti/sqldb-logger) | 180 | 5 | 4 | A logger for Go SQL database driver without modify existing *sql.DB stdlib usage. | 2019-11-02 17:28:03 | 2021-06-13 10:47:37 |
| [rollingwriter](https://github.com/arthurkiller/rollingwriter) | 176 | 26 | 8 | Rolling writer is an IO util for auto rolling write in go. | 2017-02-12 12:05:26 | 2021-06-20 14:26:49 |
| [logger](http://godoc.org/github.com/azer/logger) | 148 | 15 | 0 | Minimalistic logging library for Go. | 2014-09-30 06:45:09 | 2021-03-29 21:01:13 |
| [xlog](https://github.com/rs/xlog) | 135 | 12 | 3 | xlog is a logger for net/context aware HTTP applications | 2015-10-22 09:26:45 | 2021-02-20 02:54:42 |
| [logur](https://logur.dev/logur) | 126 | 8 | 8 | Logur is an opinionated collection of logging best practices | 2018-12-09 16:43:11 | 2021-06-10 16:44:50 |
| [glg](https://github.com/kpango/glg) | 114 | 11 | 1 | Simple and blazing fast lockfree logging library for golang | 2017-06-21 13:26:16 | 2021-06-19 04:03:48 |
| [ozzo-log](https://github.com/go-ozzo/ozzo-log) | 114 | 31 | 9 | A Go (golang) package providing high-performance asynchronous logging, message filtering by severity and category, and multiple message targets. | 2015-10-22 22:29:02 | 2021-05-17 07:01:57 |
| [logvoyage](https://github.com/firstrow/logvoyage) | 88 | 10 | 9 | LogVoyage - logging SaaS written in GoLang | 2015-03-29 11:05:09 | 2020-09-04 14:11:03 |
| [log](https://github.com/alexcesaro/log) | 45 | 3 | 1 | Logging packages for Go | 2014-04-19 14:31:56 | 2020-08-12 17:36:34 |
| [go-cronowriter](https://github.com/utahta/go-cronowriter) | 41 | 7 | 3 | Time based rotating file writer | 2017-02-04 09:02:55 | 2021-05-31 11:45:51 |
| [gologger](https://github.com/sadlil/gologger) | 39 | 10 | 2 | The Simplest and worst logging library ever written | 2015-09-02 08:52:26 | 2021-01-19 03:46:18 |
| [logex](https://github.com/chzyer/logex) | 37 | 8 | 2 | An golang log lib, supports tracking and level, wrap by standard log lib | 2014-10-10 06:38:39 | 2021-03-29 21:33:46 |
| [go-log](https://github.com/ian-kent/go-log) | 37 | 17 | 3 | A logger, for Go | 2014-05-02 00:34:09 | 2021-02-21 19:02:39 |
| [gone](https://github.com/One-com/gone) | 35 | 5 | 0 | Golang packages for writing small daemons and servers. | 2016-09-05 09:39:11 | 2021-05-24 14:22:10 |
| [go-log](https://github.com/siddontang/go-log) | 28 | 14 | 1 | a golang log lib supports level and multi handlers | 2014-05-18 03:41:55 | 2020-11-11 08:54:54 |
| [journald](https://asciinema.org/a/297429) | 26 | 1 | 0 | Go implementation of systemd Journal's native API for logging | 2017-08-23 07:06:09 | 2021-03-05 18:33:49 |
| [logrusly](https://github.com/sebest/logrusly) | 26 | 14 | 2 | Loggly Hooks for GO Logrus logger | 2014-09-11 23:27:11 | 2020-08-17 21:36:09 |
| [distillog](https://github.com/amoghe/distillog) | 26 | 5 | 0 | Logging, distilled | 2015-10-12 16:32:21 | 2021-06-21 15:05:58 |
| [log](https://github.com/teris-io/log) | 24 | 2 | 0 | Structured log interface | 2017-10-28 19:57:55 | 2021-01-03 07:59:28 |
| [mlog](https://github.com/jbrodriguez/mlog) | 23 | 17 | 1 | A simple logging module for go, with a rotating file feature and console logging. | 2014-10-20 15:06:26 | 2021-03-21 07:13:54 |
| [gomol](https://github.com/aphistic/gomol) | 17 | 0 | 3 | Gomol is a library for structured, multiple-output logging for Go with extensible logging outputs | 2015-08-30 15:51:46 | 2021-04-21 04:52:21 |
| [zkits-logger](https://github.com/edoger/zkits-logger) | 15 | 0 | 0 | A powerful zero-dependency json logger. | 2020-03-31 14:23:40 | 2021-06-17 03:37:04 |
| [glo](https://github.com/lajosbencz/glo) | 14 | 0 | 0 | Logging library for Golang | 2019-01-19 22:10:42 | 2021-06-04 05:15:29 |
| [logrusiowriter](https://github.com/cabify/logrusiowriter) | 12 | 0 | 0 | io.Writer implementation using logrus logger | 2019-08-09 08:58:25 | 2021-04-28 10:20:48 |
| [go-log](https://github.com/subchen/go-log) | 11 | 5 | 0 | Simple and configurable Logging in Go, with level, formatters and writers | 2017-05-07 08:09:24 | 2021-02-21 03:47:33 |
| [logmatic](http://godoc.org/github.com/azer/logger) | 10 | 3 | 1 | Colorized logger for Golang with dynamic log level configuration | 2018-11-07 01:52:45 | 2021-03-10 18:40:34 |
| [logdump](https://github.com/ewwwwwqm/logdump) | 9 | 2 | 0 | Package for multi-level logging | 2017-01-13 15:34:31 | 2019-07-22 17:02:07 |
| [logo](https://github.com/mbndr/logo) | 9 | 1 | 0 | Golang logger to different configurable writers. | 2017-02-07 18:02:55 | 2021-02-09 06:29:03 |
| [log](https://github.com/aerogo/log) | 8 | 0 | 0 | :memo: Logging with multiple output targets. | 2017-06-10 09:54:08 | 2020-11-19 15:43:56 |
| [go-log](https://github.com/pieterclaerhout/go-log) | 8 | 3 | 0 | A logging library with strack traces, object dumping and optional timestamps | 2019-10-01 08:55:38 | 2021-02-21 07:39:35 |
| [xlog](https://github.com/xfxdev/xlog) | 6 | 3 | 0 | plugin architecture and flexible log system for golang | 2016-05-05 16:47:45 | 2019-09-26 11:33:58 |
| [kemba](https://pkg.go.dev/github.com/clok/kemba?tab=overview) | 5 | 1 | 0 | A tiny debug logging tool. Ideal for CLI tools and command applications. Inspired by https://github.com/visionmedia/debug | 2020-07-13 03:10:54 | 2021-05-21 04:32:04 |
| [yell](https://github.com/jfcg/yell) | 1 | 0 | 0 | Yet another minimalistic logging library | 2021-02-07 16:07:27 | 2021-05-21 13:00:39 |
</details>

### Machine Learning
Libraries for Machine Learning.

<sup>*Last Update: 2021-06-25 08:44:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [golearn](https://github.com/sjwhitworth/golearn) | 7,874 | 1,101 | 68 | Machine Learning for Go | 2013-12-26 13:06:14 | 2021-06-24 15:55:37 |
| [gorgonia](https://github.com/gorgonia/gorgonia) | 4,061 | 353 | 78 | Gorgonia is a library that helps facilitate machine learning in Go. | 2016-09-14 23:19:43 | 2021-06-25 01:21:57 |
| [tfgo](https://pgaleone.eu/tensorflow/go/2017/05/29/understanding-tensorflow-using-go/) | 1,738 | 132 | 5 | Tensorflow + Go, the gopher way | 2017-05-23 13:27:39 | 2021-06-23 19:40:02 |
| [gosseract](https://pkg.go.dev/github.com/otiai10/gosseract) | 1,477 | 184 | 13 | Go package for OCR (Optical Character Recognition), by using Tesseract C++ library | 2013-10-11 07:27:53 | 2021-06-24 04:55:17 |
| [gorse](https://gorse.io) | 1,284 | 144 | 12 | An open source recommender system service written in Go | 2018-08-14 11:01:09 | 2021-06-24 23:09:43 |
| [goml](https://github.com/cdipaolo/goml) | 1,193 | 113 | 5 | On-line Machine Learning in Go (and so much more) | 2015-06-27 05:52:01 | 2021-06-24 15:58:42 |
| [eaopt](https://pkg.go.dev/github.com/MaxHalford/eaopt) | 734 | 82 | 6 | :four_leaf_clover: Evolutionary optimization library for Go (genetic algorithm, partical swarm optimization, differential evolution) | 2016-01-31 00:04:52 | 2021-06-22 01:52:06 |
| [bayesian](https://github.com/jbrukh/bayesian) | 706 | 124 | 8 | Naive Bayesian Classification for Golang. | 2011-11-23 04:17:00 | 2021-06-14 21:27:03 |
| [CloudForest](https://github.com/ryanbressler/CloudForest) | 686 | 88 | 35 | Ensembles of decision trees in go/golang. | 2012-10-22 17:38:16 | 2021-06-21 08:38:17 |
| [gobrain](https://github.com/goml/gobrain) | 480 | 55 | 1 | Neural Networks written in go | 2014-04-29 13:32:36 | 2021-06-24 02:01:23 |
| [ocrserver](https://ocr-example.herokuapp.com/) | 380 | 90 | 0 | A simple OCR API server, seriously easy to be deployed by Docker, on Heroku as well | 2015-11-15 07:57:42 | 2021-06-22 02:53:28 |
| [onnx-go](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 330 | 32 | 21 | onnx-go gives the ability to import a pre-trained neural network within Go without being linked to a framework or library. | 2018-08-28 07:39:20 | 2021-06-12 02:06:52 |
| [go-deep](https://github.com/patrikeh/go-deep) | 315 | 31 | 7 | Artificial Neural Network | 2017-12-09 15:10:06 | 2021-05-27 19:45:03 |
| [regommend](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 290 | 26 | 0 | Recommendation engine for Go | 2014-02-05 17:00:49 | 2021-06-21 14:57:05 |
| [go-galib](https://github.com/thoj/go-galib) | 185 | 42 | 0 | Genetic Algorithms library written in Go / golang | 2009-11-30 10:46:58 | 2021-06-02 23:43:19 |
| [goptuna](https://pkg.go.dev/github.com/c-bata/goptuna) | 180 | 10 | 11 | A hyperparameter optimization framework, inspired by Optuna. | 2019-07-24 12:03:05 | 2021-06-11 18:10:32 |
| [goRecommend](https://pkg.go.dev/github.com/c-bata/goptuna) | 173 | 19 | 0 | Collaborative Filtering (CF) Algorithms in Go!  | 2014-07-16 05:32:23 | 2021-06-08 13:28:06 |
| [shield](https://github.com/eaigner/shield) | 141 | 29 | 5 | Bayesian text classifier with flexible tokenizers and storage backends for Go | 2013-04-10 19:38:16 | 2021-06-12 12:32:27 |
| [goga](https://github.com/tomcraven/goga) | 107 | 11 | 0 | Golang Genetic Algorithm | 2015-10-20 12:50:51 | 2021-06-06 18:01:17 |
| [go-fann](https://github.com/vksnk/go-fann) | 103 | 21 | 2 | Go bindings for FANN, library for artificial neural networks | 2011-03-10 21:08:27 | 2021-04-11 03:02:00 |
| [libsvm](https://pkg.go.dev/github.com/otiai10/gosseract) | 67 | 12 | 1 | libsvm go version | 2012-07-31 07:57:47 | 2021-01-13 19:35:07 |
| [gonet](https://pkg.go.dev/github.com/dathoangnd/gonet) | 66 | 5 | 0 | Neural Network for Go. | 2020-01-11 18:27:28 | 2021-03-15 23:06:06 |
| [goscore](https://gorse.io) | 65 | 20 | 3 | Go Scoring API for PMML | 2017-08-19 11:08:39 | 2021-05-17 05:16:58 |
| [neural-go](https://github.com/schuyler/neural-go) | 62 | 15 | 1 | A multilayer perceptron network implemented in Go, with training via backpropagation. | 2011-10-17 09:31:33 | 2021-06-23 08:27:03 |
| [go-pr](https://github.com/daviddengcn/go-pr) | 58 | 12 | 0 | Pattern recognition package in Go lang. | 2013-06-07 02:36:20 | 2020-11-22 22:15:28 |
| [neat](https://github.com/jinyeom/neat) | 58 | 12 | 4 | NEAT (NeuroEvolution of Augmenting Topologies) implemented in Go | 2016-11-17 04:23:14 | 2020-09-07 20:06:39 |
| [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) | 56 | 5 | 0 | Fast, simple sklearn-like feature processing for Go | 2020-12-18 13:09:18 | 2021-06-22 01:49:37 |
| [fonet](https://github.com/Fontinalis/fonet) | 44 | 9 | 2 | fonet is a deep neural network package for Go. | 2017-10-03 15:57:15 | 2021-06-24 02:00:48 |
| [golinear](https://github.com/danieldk/golinear) | 41 | 12 | 0 | liblinear bindings for Go | 2013-04-05 15:37:01 | 2021-03-15 23:05:51 |
| [Varis](https://github.com/Xamber/Varis) | 35 | 6 | 0 | Golang Neural Network  | 2017-10-10 08:43:27 | 2021-05-27 11:07:07 |
| [godist](https://github.com/e-dard/godist) | 28 | 6 | 0 | Probability distributions and associated methods in Go | 2014-09-05 09:48:51 | 2021-05-03 13:34:11 |
| [go-cluster](https://github.com/e-XpertSolutions/go-cluster) | 28 | 6 | 0 | k-modes and k-prototypes clustering algorithms implementation in Go | 2017-10-04 12:24:52 | 2021-03-16 20:24:47 |
| [probab](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 16 | 4 | 3 | Automatically exported from code.google.com/p/probab | 2015-09-14 12:07:52 | 2021-06-02 23:42:13 |
| [evoli](https://github.com/khezen/evoli) | 15 | 6 | 21 | Genetic Algorithm and Particle Swarm Optimization | 2015-06-12 06:58:30 | 2021-03-15 12:19:59 |
| [ddt](https://github.com/sgrodriguez/ddt) | 12 | 1 | 0 | Golang Dynamic Decision Tree | 2020-05-20 13:51:42 | 2021-06-18 17:26:09 |
| [gomind](https://github.com/surenderthakran/gomind) | 11 | 1 | 7 | A simplistic Neural Network Library in Go | 2017-10-19 03:48:51 | 2021-05-15 17:23:23 |
| [randomForest](https://blog.owulveryck.info/2019/04/03/from-a-project-to-a-product-the-state-of-onnx-go.html) | 10 | 1 | 0 | Random Forest implementation in golang | 2018-10-25 07:05:29 | 2021-05-21 04:32:07 |
</details>

### Messaging
Libraries that implement messaging systems.

<sup>*Last Update: 2021-07-01 16:25:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [sarama](https://shopify.github.io/sarama) | 7,371 | 1,283 | 185 | Sarama is a Go library for Apache Kafka 0.8, and up. | 2013-07-05 18:52:38 | 2021-07-01 03:54:33 |
| [gorush](https://github.com/appleboy/gorush) | 5,547 | 619 | 47 | A push notification server written in Go (Golang). | 2016-03-22 07:15:20 | 2021-06-30 06:17:45 |
| [machinery](https://github.com/RichardKnop/machinery) | 5,361 | 671 | 151 | Machinery is an asynchronous task queue/job queue based on distributed message passing. | 2015-04-05 19:46:34 | 2021-07-01 06:33:35 |
| [centrifugo](https://centrifugal.github.io/centrifugo/) | 5,120 | 442 | 11 | Scalable real-time messaging server in language-agnostic way. Set up once and forever. | 2015-03-31 20:26:49 | 2021-07-01 06:39:05 |
| [go-socket.io](https://github.com/googollee/go-socket.io) | 4,143 | 680 | 76 | socket.io library for golang, a realtime application framework. | 2013-07-13 13:04:38 | 2021-07-01 03:19:47 |
| [nats.go](https://nats.io) | 3,414 | 438 | 47 | Golang client for NATS, the cloud native messaging system. | 2012-08-15 12:54:59 | 2021-07-01 04:12:11 |
| [benthos](https://www.benthos.dev) | 3,196 | 286 | 117 | Declarative stream processing for mundane tasks and data engineering | 2016-03-22 01:18:48 | 2021-07-01 08:21:02 |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | 2,751 | 425 | 154 | Confluent's Apache Kafka Golang client | 2016-07-12 22:23:34 | 2021-07-01 04:48:30 |
| [apns2](https://github.com/sideshow/apns2) | 2,512 | 271 | 22 | ⚡ HTTP/2 Apple Push Notification Service (APNs) push provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps, using the APNs HTTP/2 protocol. | 2016-01-05 00:56:53 | 2021-06-28 06:44:39 |
| [mercure](https://mercure.rocks) | 2,468 | 178 | 14 | Server-sent live updates: protocol and reference implementation | 2018-07-14 13:47:14 | 2021-06-30 22:37:25 |
| [melody](https://github.com/olahol/melody) | 2,168 | 263 | 25 | :notes: Minimalist websocket framework for Go | 2015-05-13 20:38:32 | 2021-06-29 16:50:49 |
| [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) | 1,983 | 553 | 4 | Golang push server cluster | 2013-12-27 08:56:10 | 2021-06-30 01:35:55 |
| [go-nsq](https://github.com/nsqio/go-nsq) | 1,930 | 365 | 27 | The official Go package for NSQ | 2013-08-29 01:18:32 | 2021-07-01 08:10:40 |
| [asynq](https://github.com/hibiken/asynq) | 1,320 | 100 | 13 | Asynq: simple, reliable, and efficient distributed task queue in Go | 2019-11-15 05:17:55 | 2021-07-01 08:20:46 |
| [uniqush-push](http://uniqush.org) | 1,259 | 194 | 71 | Uniqush is a free and open source software system which provides a unified push service for server side notification to apps on mobile devices. | 2011-08-29 08:42:37 | 2021-07-01 08:20:24 |
| [Beaver](https://github.com/Clivern/Beaver) | 1,107 | 60 | 17 | 💨 A real time messaging system to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. | 2018-10-20 21:10:43 | 2021-06-28 01:15:19 |
| [EventBus](https://github.com/asaskevich/EventBus) | 939 | 113 | 17 | [Go] Lightweight eventbus with async compatibility for Go | 2014-12-19 16:38:39 | 2021-06-30 22:44:50 |
| [zmq4](http://uniqush.org) | 917 | 147 | 40 | A Go interface to ZeroMQ version 4 | 2013-10-18 11:48:51 | 2021-06-23 05:15:21 |
| [gollum](http://gollum.readthedocs.org/en/latest/) | 897 | 73 | 21 | An n:m message multiplexer written in Go | 2015-06-20 21:51:20 | 2021-06-22 15:07:43 |
| [dbus](https://github.com/godbus/dbus) | 592 | 142 | 42 | Native Go bindings for D-Bus | 2014-03-27 19:07:41 | 2021-06-28 12:25:12 |
| [golongpoll](https://github.com/jcuga/golongpoll) | 559 | 43 | 0 | golang long polling library.  Makes web pub-sub easy via HTTP long-poll servers and clients :smiley: :coffee: :computer: | 2015-11-02 00:32:56 | 2021-06-29 08:38:45 |
| [mangos](https://github.com/nanomsg/mangos) | 416 | 57 | 27 | mangos is a pure Golang implementation of nanomsg's "Scalablilty Protocols" | 2018-10-12 17:35:46 | 2021-07-01 09:04:04 |
| [emitter](https://github.com/olebedev/emitter) | 406 | 31 | 4 | Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins | 2015-11-10 20:56:36 | 2021-06-01 03:05:05 |
| [glue](https://github.com/desertbit/glue) | 379 | 29 | 6 | Glue - Robust Go and Javascript Socket Library (Alternative to Socket.io) | 2015-06-07 10:21:15 | 2021-06-30 14:45:45 |
| [pubsub](https://github.com/cskr/pubsub) | 351 | 53 | 1 | A simple pubsub package for go. | 2012-04-01 06:31:43 | 2021-06-17 15:05:49 |
| [bus](https://pkg.go.dev/github.com/mustafaturan/bus/v3?tab=doc) | 210 | 14 | 0 | 🔊Minimalist message bus implementation for internal communication with zero-allocation magic on Emit | 2019-04-27 06:41:53 | 2021-06-12 16:52:54 |
| [rabtap](https://github.com/jandelgado/rabtap) | 184 | 14 | 2 | RabbitMQ wire tap and swiss army knife | 2017-11-11 11:32:39 | 2021-06-22 11:23:06 |
| [message-bus](http://rafallorenz.com/message-bus) | 180 | 29 | 2 | Go simple async message bus | 2017-10-04 09:18:34 | 2021-06-25 07:18:14 |
| [guble](https://github.com/smancke/guble) | 148 | 20 | 5 | websocket based messaging server written in golang | 2015-11-15 20:32:42 | 2021-06-03 12:06:35 |
| [oplog](https://github.com/dailymotion/oplog) | 106 | 11 | 2 | A generic oplog/replication system for microservices | 2014-11-06 09:11:15 | 2021-01-31 13:50:35 |
| [hub](https://github.com/leandro-lugaresi/hub) | 105 | 9 | 2 | :incoming_envelope: A fast Message/Event Hub using publish/subscribe pattern with support for topics like* rabbitMQ exchanges for Go applications | 2018-04-13 23:47:13 | 2021-07-01 01:05:50 |
| [rabbus](https://github.com/rafaeljesus/rabbus) | 89 | 22 | 6 | A tiny wrapper over amqp exchanges and queues 🚌 ✨ | 2017-05-07 08:51:11 | 2021-06-18 12:02:07 |
| [drone-line](https://github.com/appleboy/drone-line) | 75 | 16 | 0 | Sending line notifications using a binary, docker or Drone CI. | 2016-09-13 05:21:44 | 2021-06-18 00:53:32 |
| [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) | 67 | 12 | 2 | A tiny wrapper around NSQ topic and channel :rocket: | 2017-01-15 22:05:13 | 2021-03-24 06:52:22 |
| [go-mq](https://github.com/cheshir/go-mq) | 65 | 12 | 3 | Declare AMQP entities like queues, producers, and consumers in a declarative way. Can be used to work with RabbitMQ. | 2017-06-19 16:16:30 | 2021-06-28 08:06:50 |
| [RapidMQ](https://github.com/sybrexsys/RapidMQ) | 62 | 9 | 1 | RapidMQ is a pure, extremely productive, lightweight and reliable library for managing of the local messages queue | 2016-10-04 21:07:48 | 2021-06-02 05:57:13 |
| [redisqueue](https://godoc.org/github.com/robinjoseph08/redisqueue) | 58 | 14 | 1 | redisqueue provides a producer and consumer of a queue that uses Redis streams | 2019-07-07 04:36:54 | 2021-06-23 03:02:01 |
| [commander](https://github.com/jeroenrinzema/commander) | 55 | 3 | 2 | Build event-driven and event streaming applications with ease | 2018-04-20 12:30:51 | 2021-06-07 10:38:24 |
| [go-notify](https://github.com/TheCreeper/go-notify) | 53 | 9 | 1 | Package notify provides an implementation of the Gnome DBus Notifications Specification. | 2015-03-01 19:21:44 | 2021-05-28 09:03:01 |
| [go-res](https://github.com/jirenius/go-res) | 49 | 6 | 7 | RES Service protocol library for Go | 2018-07-15 09:10:11 | 2021-06-25 02:10:20 |
| [event](https://github.com/agoalofalife/event) | 39 | 7 | 0 | The implementation of the pattern observer | 2017-07-02 12:19:56 | 2021-06-09 02:32:03 |
| [hare](https://github.com/leozz37/hare) | 28 | 4 | 0 | 🐇  Easy to use socket lib for Golang | 2020-12-01 22:30:27 | 2021-05-23 07:11:09 |
| [ami](https://github.com/kak-tus/ami) | 20 | 6 | 0 | Go client to reliable queues based on Redis Cluster Streams | 2018-10-27 10:38:16 | 2021-06-05 17:25:39 |
| [gosd](https://github.com/alexsniffin/gosd) | 18 | 1 | 0 | A library for scheduling when to dispatch a message to a channel | 2020-05-17 23:19:51 | 2021-05-24 02:42:11 |
| [go-vitotrol](https://godoc.org/github.com/maxatome/go-vitotrol) | 16 | 3 | 1 | golang client library to Viessmann Vitotrol web service | 2016-11-03 19:59:43 | 2021-02-19 21:40:43 |
| [rmqconn](https://github.com/sbabiv/rmqconn) | 15 | 0 | 0 | RabbitMQ Reconnection client | 2019-01-14 16:05:44 | 2021-06-21 04:16:31 |
| [jazz](https://github.com/socifi/jazz) | 12 | 1 | 1 | Abstraction layer for simple rabbitMQ connection, messaging and administration | 2018-10-22 12:28:15 | 2021-05-31 21:49:08 |
| [gaurun-client](https://github.com/osamingo/gaurun-client) | 9 | 2 | 1 | Gaurun Client written in Go | 2017-06-29 02:50:51 | 2021-06-12 19:53:09 |
</details>

### Microsoft Office


<sup>*Last Update: 2021-07-13 09:25:22*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [unioffice](https://unidoc.io/unioffice/) | 2,905 | 325 | 26 | Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents | 2017-08-29 01:25:48 | 2021-07-12 11:01:52 |
</details>

### Microsoft Office - Microsoft Excel
Libraries for working with Microsoft Excel.

<sup>*Last Update: 2021-06-23 15:12:59*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [excelize](https://github.com/360EntSecGroup-Skylar/excelize) | 8,900 | 929 | 77 | Golang library for reading and writing Microsoft Excel™ (XLSX) files. | 2016-08-29 12:32:12 | 2021-06-23 07:45:10 |
| [xlsx](https://github.com/tealeg/xlsx) | 5,017 | 763 | 37 | Go (golang) library for reading and writing XLSX files.  | 2011-06-28 15:20:28 | 2021-06-22 07:40:20 |
| [xlsx](https://github.com/plandem/xlsx) | 136 | 19 | 11 | Fast and reliable way to work with Microsoft Excel™ [xlsx] files in Golang | 2017-08-26 23:11:38 | 2021-06-09 19:38:33 |
| [go-excel](https://github.com/szyhf/go-excel) | 120 | 20 | 1 | A simple and light excel file reader to read a standard excel as a table faster | 一个轻量级的Excel数据读取库，用一种更`关系数据库`的方式解析Excel。 | 2017-09-03 11:51:58 | 2021-06-16 08:20:03 |
| [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) | 15 | 3 | 1 | Golang bindings for libxlsxwriter for writing XLSX files | 2017-03-13 04:15:17 | 2020-12-31 07:08:18 |
</details>

### Miscellaneous - Dependency Injection
Libraries for working with dependency injection.

<sup>*Last Update: 2021-06-25 08:44:33*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fx](https://go.uber.org/fx) | 2,054 | 137 | 34 | A dependency injection based application framework for Go. | 2016-10-27 00:25:00 | 2021-06-24 12:11:10 |
| [dig](https://go.uber.org/dig) | 1,944 | 119 | 20 | A reflection based dependency injection toolkit for Go. | 2017-03-21 23:55:50 | 2021-06-24 14:27:04 |
| [container](https://github.com/golobby/container) | 203 | 14 | 3 | A lightweight yet powerful IoC dependency injection container for Go projects | 2019-09-23 16:12:50 | 2021-06-24 19:47:56 |
| [dingo](https://go.uber.org/dig) | 105 | 6 | 9 | Go Dependency Injection Framework | 2018-10-29 08:55:18 | 2021-05-26 13:40:49 |
| [di](https://github.com/goava/di) | 91 | 8 | 2 | 🛠 A full-featured dependency injection container for go programming language. | 2020-02-03 19:06:39 | 2021-06-24 16:44:30 |
| [di](https://pkg.go.dev/github.com/goioc/di/?tab=doc) | 82 | 2 | 0 | Simple and yet powerful Dependency Injection for Go | 2020-06-11 12:28:06 | 2021-06-18 02:14:22 |
| [alice](https://godoc.org/github.com/magic003/alice) | 44 | 3 | 0 | An additive dependency injection container for Golang. | 2017-04-08 16:25:21 | 2021-02-01 22:07:46 |
| [linker](https://github.com/logrange/linker) | 32 | 5 | 0 | Dependency Injection and Inversion of Control package | 2018-12-04 23:56:34 | 2021-05-06 05:34:51 |
| [wire](https://github.com/Fs02/wire) | 32 | 7 | 1 | Strict Runtime Dependency Injection for Golang | 2018-07-05 10:42:24 | 2021-02-14 21:50:05 |
| [gocontainer](https://github.com/vardius/gocontainer) | 14 | 1 | 0 | Simple Dependency Injection Container | 2019-06-06 08:18:07 | 2020-12-10 08:58:38 |
| [kinit](https://github.com/go-kata/kinit) | 5 | 0 | 0 | GO Dependency Injection | 2021-01-24 13:41:41 | 2021-06-12 14:24:51 |
| [nject](https://github.com/BlueOwlOpenSource/nject) | 4 | 1 | 0 | Go dependency injection: nject & npoint | 2018-10-31 18:15:43 | 2021-05-21 04:32:12 |
</details>

### Miscellaneous - Project Layout
Unofficial set of patterns for structuring projects.

<sup>*Last Update: 2021-06-26 08:31:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [project-layout](https://github.com/golang-standards/project-layout) | 24,506 | 2,594 | 57 | Standard Go Project Layout | 2017-09-09 16:33:26 | 2021-06-25 23:01:29 |
| [modern-go-application](https://github.com/sagikazarmark/modern-go-application) | 946 | 95 | 16 | Modern Go Application example | 2018-09-14 12:19:02 | 2021-06-23 12:28:50 |
| [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) | 457 | 99 | 4 | A Go project template | 2016-12-18 18:22:26 | 2021-06-24 04:50:23 |
| [seed](https://github.com/golang-templates/seed) | 178 | 13 | 0 | Go application GitHub repository template. | 2020-04-30 21:31:36 | 2021-06-23 22:37:59 |
| [scaffold](https://github.com/catchplay/scaffold) | 96 | 21 | 1 | Generate scaffold project layout for Go. | 2018-12-11 10:36:03 | 2021-06-21 15:06:01 |
| [go-sample](https://github.com/zitryss/go-sample) | 86 | 20 | 0 | Go Project Sample Layout | 2019-01-24 23:41:46 | 2021-06-13 02:29:07 |
| [go-todo-backend](https://github.com/Fs02/go-todo-backend) | 78 | 10 | 0 | Go Todo Backend example using modular project layout for product microservice. | 2020-06-25 14:28:50 | 2021-06-24 19:52:02 |
| [gobase](https://github.com/wajox/gobase) | 9 | 3 | 0 | This is a simple skeleton for golang application | 2020-12-15 16:54:20 | 2021-06-09 16:03:16 |
| [inizio](https://github.com/insidieux/inizio) | 9 | 0 | 1 | Golang project standard layout generator | 2021-03-02 20:59:22 | 2021-05-21 04:32:12 |
</details>

### Miscellaneous - Strings
Libraries for working with strings.

<sup>*Last Update: 2021-06-23 09:17:37*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xstrings](https://github.com/huandu/xstrings) | 899 | 61 | 0 | Implements string functions widely used in other languages but absent in Go. | 2015-01-06 07:25:26 | 2021-06-21 14:33:31 |
| [strutil](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 121 | 12 | 1 | String utilities for Go | 2018-08-16 06:56:15 | 2021-06-04 00:46:11 |
| [stringy](https://pkg.go.dev/github.com/gobeam/Stringy?tab=doc) | 67 | 5 | 0 | Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionalities with help of by Stringy package. | 2020-04-03 03:34:10 | 2021-06-06 00:05:41 |
</details>

### Miscellaneous - Uncategorized
These libraries were placed here because none of the other categories seemed to fit.

<sup>*Last Update: 2021-06-23 09:17:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gopsutil](https://github.com/shirou/gopsutil) | 6,431 | 1,082 | 117 | psutil for golang | 2014-04-18 07:35:28 | 2021-06-22 15:26:37 |
| [archiver](https://godoc.org/github.com/mholt/archiver) | 3,242 | 303 | 72 | Easily create & extract archives, and compress & decompress files of various formats | 2016-04-08 22:46:55 | 2021-06-22 20:27:41 |
| [gofakeit](https://github.com/brianvoe/gofakeit) | 1,906 | 106 | 0 | Random fake data generator written in go | 2015-04-24 04:45:59 | 2021-06-22 12:38:28 |
| [gatus](https://status.twinnation.org/) | 1,536 | 87 | 24 | ⛑ Gatus - Automated service health dashboard | 2019-09-04 02:35:40 | 2021-06-22 18:18:02 |
| [gosms](https://github.com/haxpax/gosms) | 1,336 | 131 | 4 | :mailbox_closed: Your own local SMS gateway in Go | 2015-01-23 19:25:55 | 2021-06-21 16:26:26 |
| [go-resiliency](https://godoc.org/github.com/eapache/go-resiliency) | 1,203 | 101 | 1 | Resiliency patterns for golang | 2014-11-29 04:11:32 | 2021-06-21 09:31:11 |
| [base64Captcha](https://captcha.mojotv.cn) | 1,121 | 188 | 9 | captcha of base64 image string | 2017-12-12 12:17:07 | 2021-06-22 08:56:44 |
| [go-commons-pool](https://github.com/jolestar/go-commons-pool) | 962 | 132 | 3 | a generic object pool for golang | 2015-12-28 14:26:23 | 2021-06-22 03:43:27 |
| [llvm](https://github.com/llir/llvm) | 713 | 49 | 20 | Library for interacting with LLVM IR in pure Go. | 2014-09-19 11:18:44 | 2021-06-21 13:20:03 |
| [shortid](https://github.com/teris-io/shortid) | 678 | 54 | 0 | Super short, fully unique, non-sequential and URL friendly Ids | 2016-01-04 01:17:10 | 2021-06-14 10:37:32 |
| [ghorg](https://github.com/gabrie30/ghorg) | 445 | 63 | 18 | Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Bitbucket, and more | 2018-03-29 02:53:05 | 2021-06-21 20:54:41 |
| [health](https://github.com/dimiro1/health) | 423 | 42 | 4 | An easy to use, extensible health check library for Go applications. | 2016-03-08 23:04:43 | 2021-06-18 03:08:52 |
| [go-conv](https://github.com/cstockton/go-conv) | 370 | 15 | 0 | Fast conversions across various Go types with a simple API. | 2016-10-11 07:41:41 | 2021-06-18 13:19:32 |
| [banner](https://github.com/dimiro1/banner) | 346 | 19 | 2 | An easy way to add useful startup banners into your Go applications | 2016-03-25 21:28:44 | 2021-06-22 08:05:47 |
| [gountries](https://github.com/pariz/gountries) | 311 | 51 | 19 | Gountries provides: Countries (ISO-3166-1), Country Subdivisions(ISO-3166-2), Currencies (ISO 4217), Geo Coordinates(ISO-6709) as well as translations, country borders and other stuff exposed as struct data. | 2016-01-13 08:04:18 | 2021-06-15 23:39:06 |
| [stateless](https://github.com/qmuntal/stateless) | 274 | 14 | 6 | Go library for creating state machines | 2019-09-11 08:19:18 | 2021-06-21 01:12:23 |
| [ffmt](https://github.com/go-ffmt/ffmt) | 230 | 16 | 2 | Golang beautify data display for Humans | 2015-02-14 15:19:45 | 2021-06-21 02:14:30 |
| [shoutrrr](https://containrrr.dev/shoutrrr/) | 217 | 31 | 15 | Notification library for gophers and their furry friends. | 2019-04-11 06:49:34 | 2021-06-18 16:08:37 |
| [lk](https://github.com/hyperboloide/lk) | 208 | 34 | 1 | Simple licensing library for golang. | 2016-07-14 16:06:07 | 2021-06-22 07:39:57 |
| [antch](https://github.com/antchfx/antch) | 207 | 39 | 4 | Antch, a fast, powerful and extensible web crawling & scraping framework for Go | 2017-09-28 05:44:17 | 2021-06-11 18:52:33 |
| [battery](https://github.com/distatus/battery) | 188 | 25 | 8 | cross-platform, normalized battery information library | 2016-03-12 23:03:40 | 2021-05-30 11:59:46 |
| [healthcheck](https://github.com/etherlabsio/healthcheck) | 169 | 26 | 1 | An simple, easily extensible and concurrent health-check library for Go services | 2017-08-18 12:48:40 | 2021-06-20 16:53:41 |
| [bitio](https://github.com/icza/bitio) | 156 | 21 | 1 | Optimized bit-level Reader and Writer for Go. | 2016-05-31 10:02:30 | 2021-06-19 12:03:49 |
| [stats](https://github.com/go-playground/stats) | 152 | 18 | 1 | :chart_with_upwards_trend: Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... | 2015-09-14 20:20:20 | 2021-06-17 21:00:26 |
| [go-unarr](https://github.com/gen2brain/go-unarr) | 137 | 21 | 7 | Go bindings for unarr (decompression library for RAR, TAR, ZIP and 7z archives) | 2015-11-01 09:38:37 | 2021-06-03 02:05:08 |
| [turtle](https://github.com/hackebrot/turtle) | 122 | 10 | 2 | Emojis for Go 😄🐢🚀 | 2017-09-08 22:25:32 | 2021-06-11 18:31:53 |
| [gommit](https://github.com/antham/gommit) | 95 | 2 | 1 | Enforce git message commit consistency | 2016-08-30 11:10:11 | 2021-06-16 07:39:29 |
| [gotoprom](https://github.com/cabify/gotoprom) | 90 | 1 | 0 | Type-safe Prometheus metrics builder library for golang | 2018-10-10 16:07:33 | 2021-06-09 15:01:42 |
| [indigo](https://github.com/osamingo/indigo) | 84 | 11 | 1 | A distributed unique ID generator of using Sonyflake and encoded by Base58 | 2016-08-31 14:17:45 | 2021-05-26 08:43:23 |
| [captcha](https://pkg.go.dev/github.com/steambap/captcha) | 79 | 14 | 0 | :sunglasses:Package captcha provides an easy to use, unopinionated API for captcha generation | 2017-09-12 06:52:15 | 2021-06-22 22:16:57 |
| [morse](https://github.com/alwindoss/morse) | 68 | 11 | 3 | Morse Code Library in Go | 2018-08-15 05:31:31 | 2021-06-13 18:40:07 |
| [persian](https://github.com/mavihq/persian) | 54 | 8 | 1 | Some utilities for Persian language in Go (Golang) | 2017-10-16 16:16:56 | 2021-06-17 05:22:04 |
| [pdfgen](https://github.com/hyperboloide/pdfgen) | 52 | 9 | 0 | HTTP service to generate PDF from Json requests | 2015-11-30 19:27:26 | 2021-05-26 01:21:01 |
| [xkg](https://godoc.org/github.com/go-xkg/xkg) | 51 | 5 | 1 | User level X Keyboard Grabber | 2015-01-05 01:04:43 | 2021-06-13 17:00:39 |
| [faker](https://github.com/pioz/faker) | 39 | 2 | 0 | Random fake data and struct generator for Go. | 2020-07-22 20:09:46 | 2021-06-18 13:19:21 |
| [browscap_go](http://browscap.org/) | 37 | 22 | 6 | GoLang Library for Browser Capabilities Project | 2014-09-18 04:47:42 | 2021-02-19 11:55:57 |
| [datacounter](https://github.com/miolini/datacounter) | 36 | 5 | 2 | Golang counters for readers/writers | 2015-10-14 19:15:50 | 2020-12-23 16:28:56 |
| [autoflags](http://pkg.go.dev/github.com/artyom/autoflags) | 35 | 2 | 0 | Populate go command line app flags from config struct | 2014-05-15 19:00:29 | 2021-05-21 04:32:13 |
| [sandid](https://pkg.go.dev/github.com/aofei/sandid) | 30 | 4 | 0 | Every grain of sand on Earth has its own ID. | 2018-06-12 01:24:14 | 2021-04-23 00:12:32 |
| [url-shortener](https://github.com/pantrif/url-shortener) | 28 | 6 | 0 | A golang URL Shortener | 2018-06-04 05:57:45 | 2021-06-16 05:23:55 |
| [gosh](https://github.com/osamingo/gosh) | 25 | 1 | 0 | Provide Go Statistics Handler, Struct, Measure Method | 2018-05-25 08:55:55 | 2021-02-10 12:48:54 |
| [xdg](https://github.com/rkoesters/xdg) | 25 | 7 | 1 | FreeDesktop.org (xdg) Specs implemented in Go | 2013-12-15 09:51:51 | 2021-05-28 21:05:39 |
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

<sup>*Last Update: 2021-06-25 10:25:07*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [prose](https://github.com/jdkato/prose) | 2,762 | 131 | 15 | :book: A Golang library for text processing, including tokenization, part-of-speech tagging, and named-entity extraction. | 2017-02-17 17:08:22 | 2021-06-24 22:55:32 |
| [go-i18n](https://github.com/nicksnyder/go-i18n) | 1,684 | 174 | 15 | Translate your Go program into multiple languages. | 2012-01-14 21:44:37 | 2021-06-24 11:50:29 |
| [gse](https://github.com/go-ego/gse) | 1,589 | 128 | 7 | Go efficient text segmentation and NLP; support english, chinese, japanese and other. Go 语言高性能分词 | 2017-06-23 15:42:35 | 2021-06-24 12:27:18 |
| [gojieba](https://github.com/yanyiwu/gojieba) | 1,528 | 217 | 46 | "结巴"中文分词的Golang版本 | 2015-09-12 01:30:44 | 2021-06-24 12:24:01 |
| [when](https://github.com/olebedev/when) | 1,119 | 62 | 14 | A natural language date/time parser with pluggable rules | 2016-12-27 13:11:46 | 2021-06-22 21:16:49 |
| [go-pinyin](https://godoc.org/github.com/mozillazg/go-pinyin) | 968 | 151 | 13 | 汉字转拼音 | 2014-11-09 14:04:33 | 2021-06-23 06:52:05 |
| [spago](https://github.com/nlpodyssey/spago) | 929 | 40 | 10 | Self-contained Machine Learning and Natural Language Processing library in Go | 2020-01-05 20:39:29 | 2021-06-22 01:52:55 |
| [kagome](https://github.com/ikawaha/kagome) | 569 | 40 | 2 | Self-contained Japanese Morphological Analyzer written in pure Go | 2014-06-26 04:38:13 | 2021-06-11 08:18:40 |
| [whatlanggo](https://github.com/abadojack/whatlanggo) | 497 | 48 | 11 | Natural language detection library for Go | 2017-02-20 17:32:01 | 2021-06-22 08:43:36 |
| [nlp](https://github.com/shixzie/nlp) | 369 | 28 | 3 | [UNMANTEINED] Extract values from strings and fill your structs with nlp. | 2017-01-25 07:19:03 | 2021-06-05 13:26:18 |
| [nlp](https://github.com/james-bowman/nlp) | 329 | 38 | 2 | Selected Machine Learning algorithms for natural language processing and semantic analysis in Golang | 2017-03-15 08:28:05 | 2021-06-25 03:06:19 |
| [sentences](https://sentences-231000.appspot.com/) | 302 | 28 | 2 | A multilingual command line sentence tokenizer in Golang | 2015-08-07 01:08:20 | 2021-06-22 10:36:47 |
| [getlang](https://github.com/rylans/getlang) | 116 | 15 | 4 | Natural language detection package in pure Go | 2018-03-01 21:27:30 | 2021-06-04 01:55:28 |
| [go-unidecode](https://godoc.org/github.com/mozillazg/go-unidecode) | 90 | 12 | 3 | ASCII transliterations of Unicode text. | 2016-07-08 13:15:10 | 2021-06-01 16:58:25 |
| [go-nlp](https://github.com/nuance/go-nlp) | 89 | 10 | 0 | Utilities for working with discrete probability distributions and other tools useful for doing NLP work | 2011-05-02 06:43:36 | 2021-01-30 05:58:16 |
| [RAKE.Go](https://github.com/afjoseph/RAKE.Go) | 82 | 15 | 4 | A Go port of the Rapid Automatic Keyword Extraction algorithm (RAKE) | 2016-12-17 13:36:25 | 2021-06-11 18:24:05 |
| [gounidecode](https://github.com/fiam/gounidecode) | 73 | 19 | 2 | Unicode transliterator for #golang | 2012-05-01 11:59:34 | 2021-05-09 09:43:14 |
| [textcat](https://github.com/pebbe/textcat) | 65 | 8 | 1 | A Go package for n-gram based text categorization, with support for utf-8 and raw text | 2012-09-21 15:04:45 | 2021-06-21 14:35:11 |
| [segment](https://github.com/blevesearch/segment) | 63 | 9 | 5 | A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 | 2014-10-16 19:24:26 | 2021-06-09 06:51:13 |
| [go-stem](svn://go-stem) | 61 | 15 | 1 | Word Stemming in Go | 2011-09-23 19:07:23 | 2021-03-21 13:10:33 |
| [MMSEGO](https://github.com/awsong/MMSEGO) | 59 | 13 | 0 | Chinese word splitting algorithm MMSEG in GO | 2012-04-18 04:06:21 | 2020-08-05 02:49:45 |
| [stemmer](https://godoc.org/github.com/dchest/stemmer) | 49 | 3 | 0 | Stemmer packages for Go programming language. Includes English, German and Dutch stemmers. | 2011-03-21 02:08:12 | 2021-04-29 22:02:55 |
| [go2vec](https://godoc.org/github.com/mozillazg/go-unidecode) | 41 | 4 | 0 | Read and use word2vec vectors in Go | 2015-01-27 12:02:04 | 2020-10-14 13:06:04 |
| [porter2](http://zhen.org/blog/generating-porter2-fsm-for-fun-and-performance/) | 40 | 5 | 1 | High Performance Porter2 Stemmer | 2015-01-21 07:30:32 | 2021-04-30 13:16:52 |
| [petrovich](https://github.com/striker2000/petrovich) | 34 | 3 | 0 | Golang port of Petrovich - an inflector for Russian anthroponyms. | 2016-12-26 22:50:38 | 2021-06-23 11:55:26 |
| [address](https://pkg.go.dev/github.com/bojanz/address) | 29 | 0 | 0 | Address handling for Go. | 2020-10-07 18:15:27 | 2021-06-13 02:26:43 |
| [snowball](https://github.com/goodsign/snowball) | 27 | 3 | 0 | Cgo binding for Snowball C library | 2012-12-11 12:42:19 | 2021-03-27 19:57:20 |
| [paicehusk](https://github.com/rookii/paicehusk) | 26 | 5 | 2 | Golang implementation of the Paice/Husk Stemming Algorithm | 2012-09-29 16:06:58 | 2019-11-28 05:32:37 |
| [go-localize](https://github.com/m1/go-localize) | 25 | 5 | 0 | i18n (Internationalization and localization) engine written in Go, used for translating locale strings.  | 2019-12-23 12:02:51 | 2021-06-24 15:17:10 |
| [mystem](https://github.com/dveselov/mystem) | 25 | 6 | 0 | CGo bindings to Yandex.Mystem | 2016-08-30 14:55:39 | 2020-08-31 10:50:04 |
| [iuliia-go](https://github.com/mehanizm/iuliia-go) | 23 | 2 | 0 | Transliterate Cyrillic → Latin in every possible way | 2020-04-27 09:29:40 | 2021-06-15 16:27:25 |
| [icu](https://github.com/goodsign/icu) | 19 | 4 | 2 | Cgo binding for icu4c library | 2012-12-11 13:09:41 | 2020-10-29 13:45:37 |
| [golibstemmer](https://github.com/rjohnsondev/golibstemmer) | 18 | 4 | 0 | Go bindings for the snowball libstemmer library including porter 2 | 2012-08-06 19:31:05 | 2020-09-14 19:40:20 |
| [govader](https://github.com/jonreiter/govader) | 15 | 2 | 0 | vader sentiment analysis in go | 2020-01-19 10:06:15 | 2021-06-22 02:34:30 |
| [transliterator](https://github.com/alexsergivan/transliterator) | 12 | 5 | 1 | Golang text Transliterator (i.e München -> Muenchen) | 2020-04-17 14:19:55 | 2021-06-17 21:02:09 |
| [shamoji](https://github.com/osamingo/shamoji) | 12 | 0 | 0 | The shamoji (杓文字) is a word filtering package | 2017-07-23 06:38:42 | 2021-01-14 18:13:19 |
| [gotokenizer](https://github.com/xujiajun/gotokenizer) | 12 | 3 | 0 | A tokenizer based on the dictionary and Bigram language models for Go. (Now only support chinese segmentation) | 2018-10-11 03:22:36 | 2021-06-10 07:32:49 |
| [detectlanguage-go](https://detectlanguage.com) | 10 | 0 | 0 | Detect Language API Go Client | 2019-12-14 23:30:44 | 2021-05-08 08:46:00 |
| [libtextcat](https://github.com/goodsign/libtextcat) | 10 | 6 | 0 | Cgo binding for libtextcat C library | 2012-12-10 21:21:47 | 2018-07-29 06:39:58 |
| [porter](https://github.com/a2800276/porter) | 8 | 0 | 0 | porter stemmer | 2013-09-17 11:10:16 | 2018-07-28 05:12:38 |
| [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) | 7 | 0 | 0 | 💬 Sentiment analyzer library using SentiWordnet in Go | 2020-04-21 09:09:28 | 2021-06-10 04:27:05 |
</details>

### Networking
Libraries for working with various layers of the network.

<sup>*Last Update: 2021-06-23 09:17:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fasthttp](https://github.com/valyala/fasthttp) | 15,359 | 1,263 | 35 | Fast HTTP package for Go. Tuned for high performance. Zero memory allocations in hot paths. Up to 10x faster than net/http | 2015-10-18 22:19:57 | 2021-06-22 09:18:24 |
| [kcptun](https://github.com/xtaci/kcptun) | 12,483 | 2,416 | 69 | A Stable & Secure Tunnel based on KCP with N:M multiplexing and FEC. Available for ARM, MIPS, 386 and AMD64。KCPプロトコルに基づく安全なトンネル。KCP 프로토콜을 기반으로 하는 보안 터널입니다。 | 2016-02-26 09:54:46 | 2021-06-22 01:26:19 |
| [webrtc](https://pion.ly) | 7,340 | 914 | 70 | Pure Go implementation of the WebRTC API | 2018-05-18 23:10:05 | 2021-06-22 17:01:56 |
| [dns](https://miek.nl/2014/august/16/go-dns-package) | 5,621 | 881 | 7 | DNS library in Go | 2010-08-03 21:56:23 | 2021-06-22 07:41:34 |
| [quic-go](https://github.com/lucas-clemente/quic-go) | 5,484 | 699 | 85 | A QUIC implementation in pure go | 2016-04-06 20:16:27 | 2021-06-22 06:51:39 |
| [gnet](https://gnet.host) | 4,733 | 545 | 23 | 🚀 gnet is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go./ gnet 是一个高性能、轻量级、非阻塞的事件驱动 Go 网络框架。 | 2019-02-24 03:48:45 | 2021-06-22 08:43:52 |
| [gopacket](https://github.com/google/gopacket) | 4,208 | 817 | 175 | Provides packet processing capabilities for Go | 2015-03-16 20:46:00 | 2021-06-22 06:42:35 |
| [httplab](https://github.com/gchaincl/httplab) | 3,700 | 121 | 12 | The interactive web server | 2017-02-08 17:13:19 | 2021-06-21 20:17:58 |
| [kcp-go](https://github.com/xtaci/kcp-go) | 2,992 | 537 | 27 |  A Crypto-Secure, Production-Grade Reliable-UDP Library for golang with FEC  | 2015-06-16 06:15:55 | 2021-06-19 03:19:22 |
| [gobgp](https://osrg.github.io/gobgp/) | 2,229 | 490 | 107 | BGP implemented in the Go Programming Language | 2014-09-14 01:51:58 | 2021-06-22 07:32:21 |
| [ssh](https://godoc.org/github.com/gliderlabs/ssh) | 2,057 | 242 | 36 | Easy SSH servers in Golang | 2016-10-03 21:53:44 | 2021-06-20 02:04:34 |
| [fortio](https://fortio.org) | 1,941 | 162 | 74 | Fortio load testing library, command line tool, advanced echo server and web UI in go (golang). Allows to specify a set query-per-second load and record latency histograms and other useful stats. | 2017-10-10 01:01:39 | 2021-06-22 06:51:24 |
| [water](https://github.com/songgao/water) | 1,258 | 194 | 19 | A simple TUN/TAP library written in native Go. | 2013-03-25 20:06:52 | 2021-06-22 06:06:02 |
| [go-getter](https://github.com/hashicorp/go-getter) | 1,187 | 147 | 96 | Package for downloading things from a string URL using a variety of protocols. | 2015-10-12 23:17:07 | 2021-06-18 07:51:05 |
| [gev](https://github.com/Allenxuxu/gev) | 1,161 | 138 | 5 | 🚀Gev is a lightweight, fast non-blocking TCP network library / websocket server based on Reactor mode. Support custom protocols to quickly and easily build high-performance servers.  | 2019-09-01 12:16:18 | 2021-06-21 13:12:03 |
| [nff-go](https://github.com/intel-go/nff-go) | 1,065 | 125 | 61 | NFF-Go -Network Function Framework for GO (former YANFF) | 2017-03-29 17:07:29 | 2021-06-16 17:38:32 |
| [sftp](https://github.com/pkg/sftp) | 1,010 | 285 | 17 | SFTP support for the go.crypto/ssh package | 2013-11-05 04:36:00 | 2021-06-21 22:25:57 |
| [grab](https://github.com/cavaliercoder/grab) | 860 | 106 | 25 | A download manager package for Go | 2016-01-05 12:46:35 | 2021-06-20 12:46:04 |
| [ftp](https://github.com/jlaffaye/ftp) | 808 | 273 | 8 | FTP client package for Go | 2011-05-06 18:31:51 | 2021-06-22 06:20:19 |
| [mdns](https://github.com/hashicorp/mdns) | 776 | 171 | 32 | Simple mDNS client/server library in Golang | 2014-01-29 19:39:18 | 2021-06-13 12:06:51 |
| [gosnmp](https://github.com/gosnmp/gosnmp) | 732 | 244 | 22 | An SNMP library written in Go | 2012-08-27 05:59:24 | 2021-06-19 07:30:34 |
| [vssh](https://github.com/yahoo/vssh) | 715 | 53 | 1 | Go Library to Execute Commands Over SSH at Scale | 2020-06-09 16:19:22 | 2021-06-18 15:19:25 |
| [lhttp](https://github.com/fanux/lhttp) | 623 | 125 | 6 | go websocket, a better way to buid your IM server | 2015-12-29 01:13:36 | 2021-06-11 07:47:35 |
| [cidranger](https://github.com/yl2chen/cidranger) | 617 | 69 | 6 | Fast IP to CIDR lookup in Golang | 2017-08-21 05:50:14 | 2021-06-21 06:45:43 |
| [peerdiscovery](https://github.com/schollz/peerdiscovery) | 491 | 35 | 5 | Pure-Go library for cross-platform local peer discovery using UDP multicast :woman: :repeat: :woman: | 2018-04-22 23:59:37 | 2021-06-11 19:07:41 |
| [gotcp](https://github.com/gansidui/gotcp) | 483 | 156 | 0 | A Go package for quickly building tcp servers | 2014-04-13 14:54:01 | 2021-06-21 18:43:35 |
| [stun](https://github.com/gortc/stun) | 466 | 48 | 4 | Fast RFC 5389 STUN implementation in go | 2016-04-24 17:46:38 | 2021-06-15 07:29:21 |
| [go-stun](https://github.com/ccding/go-stun) | 447 | 83 | 2 | A go implementation of the STUN client (RFC 3489 and RFC 5389) | 2013-08-17 22:16:33 | 2021-06-10 09:35:48 |
| [gopcap](https://github.com/akrennmair/gopcap) | 424 | 136 | 12 | A simple wrapper around libpcap for the Go programming language | 2009-11-19 10:13:48 | 2021-06-19 13:02:57 |
| [raw](https://github.com/mdlayher/raw) | 397 | 69 | 13 | Package raw enables reading and writing data at the device driver level for a network interface.  MIT Licensed. | 2015-07-06 16:11:47 | 2021-06-14 21:48:08 |
| [tcp_server](https://github.com/firstrow/tcp_server) | 381 | 133 | 4 | golang tcp server | 2014-10-13 20:38:42 | 2021-06-15 07:10:15 |
| [gmqtt](https://github.com/DrmagicE/gmqtt) | 376 | 73 | 5 | Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1 and V5 in golang | 2018-09-16 11:46:17 | 2021-06-19 16:52:17 |
| [gaio](https://github.com/xtaci/gaio) | 358 | 37 | 11 | High performance async-io(proactor) networking for Golang。golangのための高性能非同期io(proactor)ネットワーキング | 2019-12-20 05:19:00 | 2021-06-22 08:43:28 |
| [winrm](https://pion.ly) | 313 | 85 | 23 | Command-line tool and library for Windows remote command execution in Go | 2013-12-30 18:29:15 | 2021-06-18 13:17:45 |
| [arp](https://tools.ietf.org/html/rfc826) | 256 | 42 | 3 | Package arp implements the ARP protocol, as described in RFC 826. MIT Licensed. | 2015-07-06 18:50:34 | 2021-05-11 16:55:07 |
| [ftpserverlib](https://github.com/fclairamb/ftpserverlib) | 251 | 62 | 2 | golang ftp server library | 2016-09-25 12:05:29 | 2021-06-16 18:03:57 |
| [buffstreams](https://github.com/StabbyCutyou/buffstreams) | 244 | 32 | 7 | A library to simplify writing applications using TCP sockets to stream protobuff messages | 2015-06-29 19:07:31 | 2021-05-12 12:13:22 |
| [ethernet](https://en.wikipedia.org/wiki/Ethernet_frame) | 224 | 30 | 0 | Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. MIT Licensed. | 2015-07-03 00:15:18 | 2021-06-14 21:53:03 |
| [gnxi](https://github.com/google/gnxi) | 183 | 87 | 13 | gNXI Tools - gRPC Network Management/Operations Interface Tools | 2017-09-26 08:19:41 | 2021-06-10 02:22:35 |
| [nbio](https://github.com/lesismal/nbio) | 169 | 19 | 2 | High-performance, non-blocking, event-driven, easy-to-use networking framework written in Go, support TLS/HTTP 1.X/Websocket. | 2020-01-25 11:46:54 | 2021-06-21 02:13:28 |
| [jazigo](https://github.com/udhos/jazigo) | 166 | 14 | 3 | Jazigo is a tool written in Go for retrieving configuration for multiple devices, similar to rancid, fetchconfig, oxidized, Sweet. | 2016-06-07 19:53:53 | 2021-05-16 13:32:17 |
| [utp](https://github.com/anacrolix/go-libutp) | 158 | 33 | 4 | Use anacrolix/go-libutp instead | 2015-03-20 04:39:22 | 2021-05-04 15:11:32 |
| [canopus](https://github.com/zubairhamed/canopus) | 144 | 39 | 43 | CoAP Client/Server implementing RFC 7252 for the Go Language | 2015-02-24 04:12:20 | 2021-04-11 14:53:25 |
| [sslb](https://godoc.org/github.com/gliderlabs/ssh) | 130 | 25 | 10 | Golang Super Simple Load Balance | 2015-10-18 21:31:09 | 2021-02-11 17:27:49 |
| [xtcp](https://github.com/xfxdev/xtcp) | 119 | 25 | 0 | A TCP Server Framework with graceful shutdown, custom protocol. | 2016-03-31 16:50:14 | 2021-06-10 13:56:29 |
| [dhcp6](https://tools.ietf.org/html/rfc3315) | 70 | 16 | 2 | Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. MIT Licensed. | 2015-05-22 04:13:30 | 2021-04-30 17:24:35 |
| [ether](https://github.com/songgao/ether) | 69 | 5 | 0 | A Go package for sending and receiving ethernet frames. Currently supporting Linux, Freebsd, and OS X. | 2014-05-21 03:46:30 | 2021-02-28 14:18:18 |
| [packet](https://github.com/aerogo/packet) | 56 | 13 | 1 | :package: Send network packets over a TCP or UDP connection. | 2017-10-29 05:46:44 | 2021-04-07 01:43:57 |
| [linkio](https://github.com/ian-kent/linkio) | 49 | 5 | 0 | Simulate network link speed | 2014-12-24 10:50:03 | 2021-04-20 21:41:59 |
| [portproxy](https://github.com/aybabtme/portproxy) | 45 | 10 | 0 | TCP proxy, highjacks HTTP to allow CORS | 2014-12-13 02:57:36 | 2021-05-29 02:46:21 |
| [iplib](https://github.com/c-robinson/iplib) | 45 | 5 | 0 | A library  for working with IP addresses and networks in Go | 2019-05-06 06:23:41 | 2021-06-08 18:12:15 |
| [go-powerdns](https://pkg.go.dev/github.com/joeig/go-powerdns/v2) | 36 | 11 | 1 | Go PowerDNS 4.x API Client | 2018-06-21 21:37:33 | 2021-06-17 09:04:25 |
| [graval](https://github.com/koofr/graval) | 26 | 5 | 0 | An experimental go FTP server framework | 2014-04-22 19:17:18 | 2020-10-02 13:42:17 |
| [publicip](https://github.com/polera/publicip) | 23 | 5 | 0 | Go pkg for returning your public facing IP address. | 2016-12-28 19:31:07 | 2021-05-20 14:01:48 |
| [panoptes-stream](https://github.com/yahoo/panoptes-stream) | 23 | 3 | 0 | A cloud native distributed streaming network telemetry. | 2020-10-09 04:26:26 | 2021-06-09 13:06:57 |
| [golibwireshark](https://github.com/sunwxg/golibwireshark) | 19 | 4 | 0 | GoHooks make it easy to send and consume secured web-hooks from a Go application | 2015-11-16 06:48:41 | 2021-02-06 07:31:26 |
| [gohooks](https://github.com/averageflow/gohooks) | 12 | 1 | 0 | GoHooks make it easy to send and consume secured web-hooks from a Go application | 2020-10-30 17:20:36 | 2021-01-28 02:19:56 |
| [goshark](https://github.com/sunwxg/goshark) | 11 | 2 | 0 | A simple wrapper around libpcap for the Go programming language | 2015-11-01 07:23:09 | 2021-04-27 05:55:24 |
| [llb](https://github.com/kirillDanshin/llb) | 11 | 0 | 0 | Simulate network link speed | 2016-02-21 06:30:17 | 2020-09-30 19:25:46 |
| [tspool](https://github.com/two/tspool) | 9 | 1 | 0 | tcp server pool | 2018-10-27 01:05:03 | 2021-03-04 06:34:07 |
| [httpproxy](https://github.com/wzshiming/httpproxy) | 9 | 1 | 0 | HTTP proxy handler and dialer | 2018-07-18 09:42:34 | 2021-03-02 01:56:42 |
</details>

### Networking - HTTP Clients
Libraries for making HTTP requests.

<sup>*Last Update: 2021-06-22 14:34:58*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [resty](https://github.com/go-resty/resty) | 4,382 | 363 | 33 | Simple HTTP and REST client library for Go | 2015-08-28 17:48:47 | 2021-06-22 07:25:55 |
| [heimdall](http://gojek.tech) | 1,957 | 159 | 29 | An enhanced HTTP client for Go | 2018-01-19 09:32:26 | 2021-06-21 11:30:04 |
| [grequests](https://github.com/levigross/grequests) | 1,792 | 105 | 28 | A Go "clone" of the great and famous Requests library | 2015-06-11 16:41:48 | 2021-06-22 01:26:40 |
| [sling](https://github.com/dghubble/sling) | 1,329 | 105 | 0 | A Go HTTP client library for creating and sending API requests | 2015-04-02 08:42:52 | 2021-06-16 15:42:48 |
| [gentleman](https://pkg.go.dev/github.com/h2non/gentleman?tab=doc) | 902 | 49 | 19 | Plugin-driven, extensible HTTP client toolkit for Go | 2016-02-21 23:00:24 | 2021-06-16 15:42:58 |
| [pester](https://github.com/sethgrid/pester) | 553 | 61 | 4 | Go (golang) http calls with retries and backoff  | 2015-05-20 13:50:49 | 2021-06-21 16:34:44 |
| [request](https://pkg.go.dev/github.com/monaco-io/request?tab=doc) | 123 | 17 | 0 | go request, go http client | 2020-03-25 06:24:18 | 2021-06-18 16:14:30 |
| [rq](https://github.com/ddo/rq) | 39 | 2 | 1 | A nicer interface for golang stdlib HTTP client | 2017-12-26 10:48:27 | 2021-06-15 12:31:58 |
| [go-http-client](https://github.com/bozd4g/go-http-client) | 25 | 4 | 0 | An enhanced http client for Golang | 2019-12-14 11:22:19 | 2021-05-21 05:19:24 |
| [httpretry](https://github.com/ybbus/httpretry) | 14 | 1 | 0 | Enriches the standard go http client with retry functionality. | 2020-02-05 10:17:42 | 2021-05-20 19:41:31 |
</details>

### ORM
Libraries that implement Object-Relational Mapping or datamapping techniques.

<sup>*Last Update: 2021-06-25 12:39:40*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gorm](https://gorm.io) | 24,240 | 2,763 | 99 | The fantastic ORM library for Golang, aims to be developer friendly | 2013-10-25 08:31:38 | 2021-06-25 03:38:53 |
| [ent](https://entgo.io) | 7,482 | 384 | 126 | An entity framework for Go | 2019-06-12 22:53:55 | 2021-06-25 03:58:16 |
| [pg](https://pg.uptrace.dev/) | 4,684 | 348 | 100 | Golang ORM with focus on PostgreSQL features and performance | 2013-04-24 12:31:41 | 2021-06-25 03:30:33 |
| [sqlboiler](https://github.com/volatiletech/sqlboiler) | 4,029 | 387 | 102 | Generate a Go ORM tailored to your database schema. | 2016-02-21 06:18:25 | 2021-06-25 05:31:24 |
| [gorp](https://github.com/go-gorp/gorp) | 3,498 | 373 | 135 | Go Relational Persistence - an ORM-ish library for Go | 2012-01-04 19:50:09 | 2021-06-23 15:04:59 |
| [db](https://upper.io/) | 2,592 | 182 | 127 | Data access layer for PostgreSQL, CockroachDB, MySQL, SQLite and MongoDB with ORM-like features. | 2013-10-23 02:04:36 | 2021-06-25 03:57:52 |
| [gormt](https://xxjwxc.github.io/post/gormt/) | 1,264 | 214 | 26 | database to golang struct | 2019-05-05 13:10:26 | 2021-06-25 02:33:20 |
| [reform](https://gopkg.in/reform.v1) | 1,148 | 51 | 71 | A better ORM for Go, based on non-empty interfaces and code generation. | 2016-02-25 09:41:09 | 2021-06-24 15:04:34 |
| [pop](https://github.com/gobuffalo/pop) | 1,077 | 217 | 119 | A Tasty Treat For All Your Database Needs | 2018-02-07 21:13:46 | 2021-06-24 17:52:53 |
| [go-queryset](https://github.com/jirfag/go-queryset) | 620 | 56 | 17 | 100% type-safe ORM for Go (Golang) with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support. GORM under the hood. | 2017-09-03 17:29:30 | 2021-06-17 04:28:12 |
| [go-sqlbuilder](https://pkg.go.dev/github.com/huandu/go-sqlbuilder) | 593 | 61 | 1 | A flexible and powerful SQL string builder library plus a zero-config ORM. | 2017-12-27 16:37:48 | 2021-06-13 15:32:26 |
| [qbs](https://github.com/coocood/qbs) | 550 | 103 | 9 | QBS stands for Query By Struct. A Go ORM. | 2013-02-02 05:40:59 | 2021-05-05 01:33:06 |
| [rel](https://go-rel.github.io/) | 365 | 33 | 12 | :gem: Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API | 2019-10-06 07:08:01 | 2021-06-23 09:41:07 |
| [zoom](https://github.com/albrow/zoom) | 276 | 24 | 2 | A blazing-fast datastore and querying engine for Go built on Redis. | 2013-07-17 00:32:34 | 2021-05-28 10:28:40 |
| [beego](beego.me) | 190 | 51 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2021-06-25 02:31:47 |
| [grimoire](https://fs02.github.io/grimoire) | 150 | 15 | 0 | Database access layer for golang | 2018-03-05 16:52:20 | 2021-06-14 20:06:55 |
| [gosql](https://github.com/rushteam/gosql) | 142 | 15 | 2 | golang orm and sql builder | 2020-04-27 09:16:29 | 2021-06-21 07:03:38 |
| [go-store](https://github.com/gosuri/go-store) | 102 | 8 | 1 | A simple and fast Redis backed key-value store library for Go | 2015-03-22 12:07:29 | 2021-06-05 22:34:23 |
| [go-firestorm](https://github.com/jschoedt/go-firestorm) | 24 | 6 | 3 | Simple Go ORM for Google/Firebase Cloud Firestore | 2018-12-04 14:53:53 | 2021-05-21 05:19:28 |
| [lore](https://github.com/abrahambotros/lore) | 7 | 1 | 0 | Light Object-Relational Environment (LORE) provides a simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go | 2017-04-29 03:57:15 | 2021-04-20 15:13:47 |
| [marlow](https://github.com/marlow/marlow) | 6 | 0 | 0 | persistence layer code generation for golang | 2020-08-11 13:34:00 | 2021-05-22 15:11:30 |
</details>

### OpenGL
Libraries for using OpenGL in Go.

<sup>*Last Update: 2021-06-22 14:34:55*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glfw](https://github.com/go-gl/glfw) | 1,108 | 136 | 12 | Go bindings for GLFW 3 | 2013-05-19 06:38:45 | 2021-06-15 22:38:17 |
| [gl](https://github.com/go-gl/gl) | 814 | 55 | 11 | Go bindings for OpenGL (generated via glow) | 2015-02-22 03:29:45 | 2021-06-10 21:54:18 |
| [mathgl](https://github.com/go-gl/mathgl) | 379 | 49 | 9 | A pure Go 3D math library. | 2013-02-13 14:18:55 | 2021-06-14 02:50:14 |
| [gl](https://github.com/goxjs/gl) | 150 | 16 | 9 | Go cross-platform OpenGL bindings. | 2015-05-18 08:10:15 | 2021-06-04 02:46:58 |
| [glfw](https://github.com/goxjs/glfw) | 70 | 15 | 10 | Go cross-platform glfw library for creating an OpenGL context and receiving events. | 2014-12-27 22:40:24 | 2021-04-22 20:14:17 |
| [go-glmatrix](https://github.com/technohippy/go-glmatrix) | 2 | 2 | 0 | go-glmatrix is a golang version of glMatrix, which is "designed to perform vector and matrix operations stupidly fast". | 2020-07-02 13:40:40 | 2021-03-17 16:22:40 |
</details>

### Package Management - Official
Official experimental tooling for package management

<sup>*Last Update: 2021-06-28 15:13:43*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [dep](https://golang.github.io/dep/) | 13,162 | 1,104 | 0 | Go dependency management tool experiment (deprecated) | 2016-10-07 00:04:51 | 2021-06-27 23:49:32 |
</details>

### Package Management - Unofficial
Unofficial libraries for package and dependency management

<sup>*Last Update: 2021-06-24 16:25:11*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [glide](https://glide.sh) | 8,099 | 536 | 421 | Package Management for Golang | 2014-07-09 06:02:50 | 2021-06-19 10:08:33 |
| [godep](http://godoc.org/github.com/tools/godep) | 5,614 | 486 | 79 | dependency tool for go | 2013-05-01 07:55:35 | 2021-06-22 03:44:36 |
| [govendor](https://blog.golang.org/migrating-to-go-modules) | 5,004 | 411 | 122 | Use Go Modules. | 2015-04-12 15:26:40 | 2021-06-23 22:25:08 |
| [gopm](https://github.com/gpmgo/gopm) | 2,493 | 215 | 0 | Go Package Manager (gopm) is a package manager and build tool for Go. | 2013-05-15 14:53:29 | 2021-06-18 15:12:52 |
| [gom](https://github.com/mattn/gom) | 1,390 | 104 | 14 | Go Manager - bundle for go | 2013-09-11 02:08:59 | 2021-06-23 22:25:33 |
| [gpm](https://github.com/pote/gpm) | 1,196 | 51 | 15 | Barebones dependency manager for Go. | 2013-09-05 02:24:02 | 2021-06-21 14:56:48 |
| [goop](https://github.com/petejkim/goop) | 779 | 46 | 29 | A simple dependency manager for Go (golang), inspired by Bundler. | 2014-06-18 01:55:24 | 2021-04-05 08:27:03 |
| [modgv](https://github.com/lucasepe/modgv) | 379 | 11 | 1 | Converts 'go mod graph' output into Graphviz's DOT language | 2020-09-12 16:23:46 | 2021-06-22 23:53:44 |
| [nut](https://github.com/jingweno/nut) | 241 | 11 | 14 | Vendor Go dependencies | 2015-01-23 14:46:32 | 2021-02-11 17:30:39 |
| [johnny-deps](https://github.com/VividCortex/johnny-deps) | 215 | 6 | 3 | Barebones dependency manager for Go. | 2013-07-19 15:20:47 | 2020-12-15 17:40:11 |
| [mvn-golang](https://github.com/raydac/mvn-golang) | 127 | 25 | 0 | maven plugin to automate GoSDK load and build of projects | 2016-03-24 06:47:08 | 2021-06-22 11:57:45 |
| [VenGO](https://github.com/DamnWidget/VenGO) | 120 | 10 | 3 | Create and manage Isolated Virtual Environments for Go | 2014-10-17 15:19:03 | 2021-04-21 01:54:12 |
| [gop](https://github.com/lunny/gop) | 49 | 6 | 10 | Moved to https://gitea.com/lunny/gop | 2017-02-18 04:33:48 | 2021-01-14 20:53:39 |
</details>

### Performance


<sup>*Last Update: 2021-06-23 15:12:57*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [jaeger](https://www.jaegertracing.io/) | 13,682 | 1,617 | 385 | CNCF Jaeger, a Distributed Tracing Platform | 2016-04-15 18:49:02 | 2021-06-23 06:31:00 |
| [pixie](https://px.dev/) | 1,629 | 71 | 52 | Instant Kubernetes-Native Application Observability | 2020-02-27 00:22:45 | 2021-06-23 02:55:58 |
| [profile](https://px.dev/) | 1,491 | 101 | 6 | Simple profiling for Go | 2014-10-22 01:35:18 | 2021-06-23 01:17:43 |
| [statsviz](https://github.com/arl/statsviz) | 1,378 | 49 | 6 | :rocket: Instant live visualization of your Go application runtime statistics (GC, MemStats, etc.) in the browser | 2020-08-14 00:00:41 | 2021-06-19 16:57:15 |
| [tracer](https://github.com/kamilsk/tracer) | 45 | 1 | 11 | 🧶 Dead simple, lightweight tracing. | 2019-06-22 13:23:27 | 2021-05-28 09:42:51 |
</details>

### Query Language


<sup>*Last Update: 2021-06-22 09:01:41*</sup>
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

### Resource Embedding


<sup>*Last Update: 2021-06-24 16:25:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [packr](https://github.com/gobuffalo/packr) | 3,241 | 175 | 62 | The simple and easy way to embed static files into Go binaries. | 2017-03-15 22:24:53 | 2021-06-24 08:58:11 |
| [statik](https://github.com/rakyll/statik) | 3,206 | 200 | 33 | Embed files into a Go executable | 2014-02-04 14:54:51 | 2021-06-22 10:52:08 |
| [go.rice](https://github.com/GeertJohan/go.rice) | 2,218 | 135 | 35 | go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. | 2013-10-23 21:29:34 | 2021-06-21 00:42:10 |
| [vfsgen](https://github.com/shurcooL/vfsgen) | 931 | 79 | 32 | Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. | 2015-05-18 13:03:02 | 2021-06-18 15:30:57 |
| [esc](http://godoc.org/github.com/mjibson/esc) | 597 | 66 | 11 | A simple file embedder for Go | 2014-01-26 05:08:04 | 2021-06-19 03:43:39 |
| [fileb0x](https://github.com/UnnoTed/fileb0x) | 585 | 48 | 8 | a better customizable tool to embed files in go; also update embedded files remotely without restarting the server | 2016-01-23 20:19:33 | 2021-06-02 12:58:17 |
| [go-resources](https://github.com/omeid/go-resources) | 172 | 18 | 3 | Unfancy resources embedding for Go with out of box http.FileSystem support. | 2015-02-21 15:40:17 | 2021-06-14 12:38:59 |
| [statics](https://github.com/go-playground/statics) | 62 | 4 | 0 | :file_folder: Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks | 2015-10-07 11:49:52 | 2020-11-10 12:17:01 |
| [templify](https://github.com/wlbr/templify) | 26 | 4 | 1 | A tool to be used with 'go generate' to embed external template files into Go code. | 2016-05-22 16:42:47 | 2021-02-07 19:39:17 |
| [rebed](https://github.com/soypat/rebed) | 15 | 1 | 0 | Recreates directory and files from embedded filesystem using Go 1.16 embed.FS type. | 2021-02-17 18:19:49 | 2021-06-09 06:01:08 |
| [mule](https://github.com/wlbr/mule) | 9 | 1 | 1 | mule is a tool to be used with 'go generate' to embed external resources files into Go code. | 2020-01-17 10:56:00 | 2021-01-05 11:20:02 |
| [debme](https://github.com/leaanthony/debme) | 7 | 1 | 0 | embed.FS wrapper providing additional functionality | 2021-04-16 00:25:13 | 2021-06-07 21:17:48 |
</details>

### Science and Data Analysis
Libraries for scientific computing and data analyzing.

<sup>*Last Update: 2021-06-28 16:25:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gonum](https://www.gonum.org/) | 4,967 | 400 | 212 | Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more | 2017-03-25 14:54:38 | 2021-06-28 01:15:44 |
| [stats](https://github.com/montanaflynn/stats) | 2,008 | 131 | 14 | A well tested and comprehensive Golang statistics library package with no dependencies. | 2014-12-16 03:25:19 | 2021-06-28 09:06:29 |
| [plot](https://github.com/gonum/plot) | 1,919 | 174 | 98 | A repository for plotting and visualizing data | 2013-07-23 07:01:13 | 2021-06-25 22:52:37 |
| [gosl](https://github.com/cpmech/gosl) | 1,587 | 144 | 0 | Linear algebra, eigenvalues, FFT, Bessel, elliptic, orthogonal polys, geometry, NURBS, numerical quadrature, 3D transfinite interpolation, random numbers, Mersenne twister, probability distributions, optimisation, differential equations. | 2015-02-09 23:00:38 | 2021-06-28 08:07:32 |
| [streamtools](http://nytlabs.github.io/streamtools/) | 1,310 | 108 | 47 | tools for working with streams of data | 2013-07-05 18:58:45 | 2021-06-21 14:55:00 |
| [go-dsp](http://godoc.org/github.com/mjibson/go-dsp) | 737 | 75 | 6 | Digital Signal Processing for Go | 2011-11-02 06:28:41 | 2021-06-26 19:38:12 |
| [chart](https://github.com/vdobler/chart) | 679 | 101 | 6 | Provide basic charts in go | 2011-06-27 12:19:42 | 2021-06-21 14:34:02 |
| [goraph](https://github.com/gyuho/goraph) | 641 | 74 | 6 | Package goraph implements graph data structure and algorithms. | 2014-02-27 03:15:55 | 2021-06-17 07:54:36 |
| [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) | 558 | 55 | 5 | DataFrames for Go: For statistics, machine-learning, and data manipulation/exploration | 2018-10-01 12:19:31 | 2021-06-24 17:04:32 |
| [graph](https://yourbasic.org/golang/your-basic-func/) | 455 | 45 | 2 | Graph algorithms and data structures | 2017-04-27 18:43:54 | 2021-06-28 02:58:45 |
| [orb](https://github.com/paulmach/orb) | 401 | 51 | 11 | Types and utilities for working with 2d geometry in Golang | 2016-03-28 01:19:01 | 2021-06-20 23:54:07 |
| [ewma](https://github.com/VividCortex/ewma) | 335 | 23 | 2 | Exponentially Weighted Moving Average algorithms for Go. | 2013-07-05 21:33:25 | 2021-06-26 22:08:41 |
| [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) | 262 | 7 | 2 | Calendar heatmap inspired by GitHub contribution activity | 2020-07-01 18:30:48 | 2021-06-28 09:20:00 |
| [gohistogram](http://godoc.org/github.com/VividCortex/gohistogram) | 153 | 28 | 1 | Streaming approximate histograms in Go | 2013-07-02 12:53:22 | 2021-06-05 08:24:50 |
| [TextRank](https://github.com/DavidBelicza/TextRank) | 142 | 17 | 4 | :wink: :cyclone: :strawberry: TextRank implementation in Golang with extendable features (summarization, phrase extraction) and multithreading (goroutine). | 2018-01-09 19:36:17 | 2021-06-25 13:04:42 |
| [sparse](https://github.com/james-bowman/sparse) | 116 | 19 | 4 | Sparse matrix formats for linear algebra supporting scientific and machine learning applications | 2017-05-16 13:54:36 | 2021-06-23 22:27:46 |
| [go-estimate](https://github.com/milosgajdos/go-estimate) | 76 | 6 | 1 | State estimation and filtering algorithms in Go | 2018-11-04 22:32:52 | 2021-06-11 06:53:45 |
| [pagerank](https://github.com/alixaxel/pagerank) | 70 | 16 | 3 | Weighted PageRank implementation in Go | 2015-08-06 01:33:34 | 2021-06-23 10:19:33 |
| [geom](https://github.com/skelterjohn/geom) | 47 | 16 | 1 | 2d geometry for golang | 2011-06-07 17:49:11 | 2021-04-07 08:17:36 |
| [evaler](https://github.com/soniah/evaler) | 44 | 12 | 5 | Implements a simple floating point arithmetic expression evaluator in Go (golang). | 2012-09-04 23:37:58 | 2021-04-01 11:17:53 |
| [goent](https://github.com/kzahedi/goent) | 25 | 1 | 0 | GO Implementation of Entropy Measures | 2017-08-08 05:37:12 | 2021-05-10 16:11:38 |
| [triangolatte](https://github.com/tchayen/triangolatte) | 23 | 1 | 4 | 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. | 2018-07-18 21:17:09 | 2021-06-16 04:40:28 |
| [decimal](https://github.com/db47h/decimal) | 22 | 0 | 0 | An arbitrary-precision decimal floating-point arithmetic package for Go | 2020-05-27 15:23:59 | 2021-06-05 08:18:20 |
| [piecewiselinear](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 18 | 1 | 0 | tiny linear interpolation library for go (factored out from https://github.com/sgreben/yeetgif) | 2018-10-21 13:19:44 | 2021-06-19 03:56:31 |
| [godesim](https://github.com/soypat/godesim) | 16 | 0 | 1 | ODE system solver made simple. For IVPs (initial value problems). | 2020-12-16 01:02:26 | 2021-06-16 03:50:02 |
| [GoStats](https://github.com/OGFris/GoStats) | 15 | 0 | 0 | GoStats is a go library for math statistics mostly used in ML domains, it covers most of the statistical measures functions. | 2018-07-22 20:55:16 | 2021-04-26 15:21:10 |
| [PiHex](https://pkg.go.dev/github.com/sgreben/piecewiselinear) | 14 | 2 | 0 | PiHex Library, written in Go, generates a hexadecimal number sequence in the number Pi in the range from 0 to 10,000,000. | 2016-07-22 11:21:37 | 2021-03-18 18:52:26 |
| [ode](https://yourbasic.org/golang/your-basic-func/) | 13 | 0 | 1 | An ordinary differential equation solving library in golang. | 2016-11-11 22:40:21 | 2021-01-26 11:20:02 |
| [assocentity](https://pkg.go.dev/github.com/ndabAP/assocentity/v8?tab=doc) | 6 | 1 | 6 | Package assocentity returns the average distance from words to a given entity | 2018-12-21 07:17:09 | 2020-10-27 12:49:42 |
| [rootfinding](https://github.com/khezen/rootfinding) | 5 | 0 | 0 | root-finding library | 2018-10-30 22:31:48 | 2021-05-31 18:37:41 |
| [go-gt](https://github.com/ThePaw/go-gt) | 5 | 0 | 2 | Automatically exported from code.google.com/p/go-gt | 2015-09-14 12:05:37 | 2019-09-05 08:47:43 |
| [bradleyterry](https://pkg.go.dev/github.com/ndabAP/assocentity/v8?tab=doc) | 4 | 0 | 0 | Package to do Bradley-Terry Model pairwise compairsons | 2019-04-30 00:28:13 | 2020-09-08 12:32:53 |
</details>

### Security
Libraries that are used to help make your application more secure.

<sup>*Last Update: 2021-06-24 08:56:08*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [lego](https://go-acme.github.io/lego/) | 4,657 | 619 | 116 | Let's Encrypt client and ACME library written in Go | 2015-06-08 00:36:41 | 2021-06-23 21:29:14 |
| [cameradar](https://github.com/Ullaakut/cameradar) | 2,595 | 354 | 13 | Cameradar hacks its way into RTSP videosurveillance cameras | 2016-05-20 11:35:41 | 2021-06-23 23:01:17 |
| [crypto](https://golang.org/x/crypto) | 2,198 | 1,216 | 42 | [mirror] Go supplementary cryptography libraries | 2014-12-04 04:02:55 | 2021-06-21 01:49:15 |
| [memguard](https://github.com/awnumar/memguard) | 1,956 | 89 | 4 | Secure software enclave for storage of sensitive information in memory. | 2017-04-22 07:40:40 | 2021-06-19 07:23:19 |
| [acmetool](https://hlandau.github.io/acmetool/) | 1,854 | 117 | 64 | :lock: acmetool, an automatic certificate acquisition tool for ACME (Let's Encrypt) | 2015-11-15 01:56:02 | 2021-06-20 09:01:52 |
| [secure](https://github.com/unrolled/secure) | 1,764 | 111 | 3 | HTTP middleware for Go that facilitates some quick security wins. | 2014-05-20 19:46:28 | 2021-06-22 14:15:33 |
| [themis](https://www.cossacklabs.com/themis) | 1,300 | 110 | 9 | Easy to use cryptographic framework for data protection: secure messaging with forward secrecy and secure data storage. Has unified APIs across 14 platforms. | 2015-05-06 13:25:25 | 2021-06-22 07:24:06 |
| [acra](https://www.cossacklabs.com/acra/) | 768 | 89 | 7 | Database security suite. Database proxy with field-level encryption, search through encrypted data, SQL injections prevention, intrusion detection, honeypots. Supports client-side and proxy-side ("transparent") encryption. SQL, NoSQL. | 2016-11-14 16:23:25 | 2021-06-23 19:44:49 |
| [nacl](https://godoc.org/github.com/kevinburke/nacl) | 506 | 28 | 3 | Pure Go implementation of the NaCL set of API's | 2017-07-20 19:07:19 | 2021-06-14 07:43:03 |
| [firewalld-rest](https://pkg.go.dev/github.com/prashantgupta24/firewalld-rest) | 306 | 12 | 1 | A rest application to update firewalld rules on a linux server  | 2020-06-12 20:16:33 | 2021-03-26 06:41:02 |
| [badactor](https://badactor.org) | 297 | 14 | 1 | BadActor.org An in-memory application driven jailer written in Go | 2014-12-12 20:05:20 | 2021-06-07 23:16:59 |
| [ssh-vault](https://ssh-vault.com) | 290 | 20 | 8 | 🌰  encrypt/decrypt using ssh keys | 2016-09-29 14:46:30 | 2021-06-19 22:23:59 |
| [optimus-go](https://github.com/pjebs/optimus-go) | 278 | 18 | 0 | ID hashing and Obfuscation using Knuth's Algorithm | 2015-05-05 10:12:38 | 2021-06-23 03:06:21 |
| [go-password-validator](https://qvault.io/2020/10/15/how-to-correctly-validate-passwords-most-websites-do-it-wrong/) | 278 | 19 | 0 | Validate the Strength of a Password in Go | 2020-10-14 15:52:14 | 2021-06-17 14:03:10 |
| [passlib](https://github.com/hlandau/passlib) | 242 | 26 | 1 | :key: Idiotproof golang password validation library inspired by Python's passlib | 2014-12-21 17:45:52 | 2021-05-24 05:31:48 |
| [go-yara](https://github.com/hillu/go-yara) | 211 | 79 | 4 | Go bindings for YARA | 2015-01-25 01:01:11 | 2021-06-16 07:44:48 |
| [simple-scrypt](https://github.com/elithrar/simple-scrypt) | 171 | 19 | 4 | A convenience library for generating, comparing and inspecting password hashes using the scrypt KDF in Go 🔑 | 2015-04-14 06:52:21 | 2021-05-21 04:32:23 |
| [argon2pw](https://github.com/raja/argon2pw) | 86 | 7 | 0 | Argon2 password hashing package for go with constant time hash comparison | 2018-03-13 13:56:36 | 2021-06-05 14:35:16 |
| [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) | 42 | 6 | 0 | A probably paranoid Golang utility library for securely hashing and encrypting passwords based on the Dropbox method. This implementation uses Blake2b, Scrypt and XSalsa20-Poly1305 (via NaCl SecretBox) to create secure password hashes that are also encrypted using a master passphrase. | 2017-10-19 19:34:45 | 2021-05-04 18:20:06 |
| [certificates](https://github.com/mvmaasakkers/certificates) | 20 | 3 | 0 | An opinionated helper for generating tls certificates | 2019-03-04 07:20:36 | 2020-12-09 19:50:01 |
| [go-generate-password](https://github.com/m1/go-generate-password) | 19 | 2 | 1 | Password generator written in Go | 2019-11-14 17:57:19 | 2021-06-03 14:24:14 |
| [secureio](https://github.com/xaionaro-go/secureio) | 16 | 1 | 1 | An easy-to-use XChaCha20-encryption wrapper for io.ReadWriteCloser (even lossy UDP) using ECDH key exchange algorithm, ED25519 signatures and Blake3+Poly1305 checksums/message-authentication for Go (golang). Also a multiplexer. | 2018-12-25 14:20:59 | 2021-05-25 04:45:16 |
| [goArgonPass](https://github.com/dwin/goArgonPass) | 14 | 5 | 1 | goArgonPass is a Argon2 Password utility package for Go using the crypto library package Argon2 designed to be compatible with Passlib for Python and Argon2 PHP. Argon2 was the winner of the most recent Password Hashing Competition. This is designed for use anywhere password hashing and verification might be needed and is intended to replace implementations using bcrypt or Scrypt. | 2018-05-30 01:32:10 | 2021-02-01 13:35:19 |
| [argon2-hashing](https://www.cossacklabs.com/acra/) | 12 | 1 | 0 | A light package for generating and comparing password hashing with argon2 in Go | 2019-01-02 20:41:02 | 2021-05-15 23:48:59 |
| [sslmgr](https://github.com/adrianosela/sslmgr) | 11 | 2 | 0 | A layer of abstraction the around acme/autocert certificate manager (Golang) | 2019-04-02 17:35:38 | 2021-06-10 17:50:47 |
</details>

### Serialization
Libraries and tools for binary serialization.

<sup>*Last Update: 2021-06-25 14:25:05*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go](http://jsoniter.com/migrate-from-go-std.html) | 9,439 | 769 | 143 | A high-performance 100% compatible drop-in replacement of "encoding/json" | 2016-11-30 00:30:24 | 2021-06-24 01:39:25 |
| [protobuf](https://github.com/golang/protobuf) | 7,729 | 1,413 | 51 | Go support for Google's protocol buffers | 2014-11-23 23:07:23 | 2021-06-23 14:27:39 |
| [protobuf](https://github.com/gogo/protobuf) | 4,667 | 608 | 207 | [Looking for new ownership] Protocol Buffers for Go with Gadgets | 2014-12-03 11:27:10 | 2021-06-23 06:48:20 |
| [mapstructure](https://github.com/mitchellh/mapstructure) | 4,584 | 482 | 40 | Go library for decoding generic map values into native Go structures and vice versa. | 2013-05-20 05:24:34 | 2021-06-23 15:32:52 |
| [go](https://github.com/ugorji/go) | 1,545 | 259 | 1 | idiomatic codec and rpc lib for msgpack, cbor, json, etc. msgpack.org[Go] | 2013-05-30 02:13:13 | 2021-06-23 07:19:52 |
| [colfer](https://github.com/pascaldekloe/colfer) | 606 | 48 | 12 | binary serialization format | 2015-09-05 16:42:41 | 2021-06-05 14:01:53 |
| [csvutil](https://github.com/jszwec/csvutil) | 540 | 29 | 0 | csvutil provides fast and idiomatic mapping between CSV and Go (golang) values. | 2017-10-30 04:09:48 | 2021-06-23 10:39:47 |
| [cbor](https://github.com/fxamacker/cbor) | 284 | 21 | 9 | CBOR codec (in Go) with CBOR tags, Go struct tags (toarray/keyasint/omitempty), float64/32/16, big.Int, and fuzz tested billions of execs for reliable RFC 7049 & RFC 8949.  | 2019-05-15 21:22:15 | 2021-06-23 13:10:09 |
| [go-capnproto](https://github.com/glycerine/go-capnproto) | 279 | 18 | 1 | Cap'n Proto library and parser for go. This is go-capnproto-1.0, and does not have rpc. See https://github.com/zombiezen/go-capnproto2 for 2.0 which has rpc and capabilities. | 2013-11-07 20:28:12 | 2021-06-14 12:15:27 |
| [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) | 149 | 41 | 3 | PHP session encoder/decoder written in Go | 2012-12-23 14:04:25 | 2021-06-17 08:52:21 |
| [structomap](https://godoc.org/github.com/danhper/structomap) | 122 | 10 | 0 | Easily and dynamically generate maps from Go static structures | 2015-05-13 08:54:11 | 2021-03-01 12:52:38 |
| [bambam](https://github.com/glycerine/bambam) | 62 | 10 | 3 | auto-generate capnproto schema from your golang source files. Depends on go-capnproto-1.0 at https://github.com/glycerine/go-capnproto | 2014-09-17 14:39:12 | 2021-03-04 20:35:06 |
| [asn1](https://github.com/Logicalis/asn1) | 47 | 21 | 6 | Asn.1 BER and DER encoding library for golang. | 2016-02-29 13:00:25 | 2020-08-21 06:42:00 |
| [binstruct](https://github.com/ghostiam/binstruct) | 31 | 5 | 0 | Golang binary decoder for mapping data into the structure | 2018-10-23 15:42:22 | 2021-05-31 21:50:02 |
| [elastic](https://github.com/epiclabs-io/elastic) | 15 | 1 | 1 | Converts go types no matter what | 2020-02-25 19:55:00 | 2021-05-20 08:50:22 |
| [fwencoder](https://github.com/o1egl/fwencoder) | 14 | 3 | 0 | Fixed width file parser (encoder/decoder) in GO (golang) | 2017-12-25 12:55:29 | 2020-09-10 10:26:54 |
| [pletter](https://github.com/vimeda/pletter) | 14 | 1 | 3 | A standard way to wrap a proto message | 2019-07-09 14:02:08 | 2021-05-21 04:32:24 |
| [bel](https://github.com/csweichel/bel) | 12 | 4 | 2 | Generate TypeScript interfaces from Go structs/interfaces - useful for JSON RPC | 2019-02-20 20:48:24 | 2021-03-06 17:17:44 |
| [fixedwidth](https://github.com/huydang284/fixedwidth) | 5 | 0 | 0 | A Go package for encode/decode fixed-width data | 2019-08-11 03:42:24 | 2020-12-07 08:51:25 |
| [go-lctree](https://github.com/sbourlon/go-lctree) | 2 | 0 | 0 | go-lctree provides a CLI and Go primitives to serialize and deserialize LeetCode binary trees (e.g. "[5,4,7,3,null,2,null,-1,null,9]"). | 2020-05-04 05:39:46 | 2020-12-24 05:48:12 |
| [unitpacking](https://github.com/recolude/unitpacking) | 2 | 0 | 0 | A library for storing unit vectors in a representation that lends itself to saving space on disk. | 2021-01-17 22:31:41 | 2021-05-21 04:32:24 |
</details>

### Server Applications


<sup>*Last Update: 2021-07-15 09:25:05*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [etcd](https://etcd.io) | 36,552 | 7,814 | 214 | Distributed reliable key-value store for the most critical data of a distributed system | 2013-07-06 21:57:21 | 2021-07-15 01:35:16 |
| [caddy](https://caddyserver.com) | 33,952 | 2,740 | 73 | Fast, multi-platform web server with automatic HTTPS | 2015-01-13 19:45:03 | 2021-07-14 23:16:45 |
| [minio](https://min.io/download) | 28,451 | 3,122 | 35 | High Performance, Kubernetes Native Object Storage | 2015-01-14 19:23:58 | 2021-07-15 01:27:33 |
| [consul](https://www.consul.io) | 22,505 | 3,779 | 907 | Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. | 2013-11-04 22:15:27 | 2021-07-15 02:20:33 |
| [nsq](https://nsq.io) | 19,937 | 2,585 | 62 | A realtime distributed messaging platform | 2012-05-12 14:37:08 | 2021-07-14 23:04:45 |
| [roadrunner](https://roadrunner.dev/) | 5,727 | 308 | 41 | High-performance PHP application server, load-balancer and process manager written in Golang | 2017-12-26 16:13:10 | 2021-07-14 18:24:38 |
| [devd](https://github.com/cortesi/devd) | 3,142 | 138 | 21 |  A local webserver for developers | 2015-09-27 22:43:00 | 2021-07-14 21:14:52 |
| [sftpgo](https://github.com/drakkan/sftpgo) | 2,757 | 209 | 7 | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-07-20 10:18:31 | 2021-07-14 20:25:35 |
| [algernon](https://algernon.roboticoverlords.org) | 1,824 | 98 | 11 | :tophat: Small self-contained pure-Go web server with Lua, Markdown, HTTP/2, QUIC, Redis and PostgreSQL support | 2015-03-10 11:25:30 | 2021-07-13 23:18:40 |
| [flagr](https://checkr.github.io/flagr) | 1,618 | 142 | 63 | Flagr is a feature flagging, A/B testing and dynamic configuration microservice | 2017-10-03 19:07:32 | 2021-07-13 07:14:48 |
| [fider](https://getfider.com) | 1,604 | 445 | 31 | Open platform to collect and prioritize product feedback | 2017-01-17 22:55:19 | 2021-07-14 04:19:57 |
| [flipt](https://flipt.io) | 1,521 | 72 | 12 | An open-source, on-prem feature flag solution | 2016-11-05 00:09:07 | 2021-07-15 01:15:09 |
| [trickster](https://trickstercache.org) | 1,452 | 139 | 19 | Open Source HTTP Reverse Proxy Cache and Time Series Dashboard Accelerator | 2018-03-29 20:31:44 | 2021-07-14 15:32:45 |
| [discovery](https://github.com/bilibili/discovery) | 1,437 | 336 | 18 | A registry for resilient mid-tier load balancing and failover. | 2018-04-20 12:57:50 | 2021-07-14 14:36:28 |
| [jackal](https://github.com/ortuman/jackal) | 1,103 | 88 | 10 | 💬 Instant messaging server for the Extensible Messaging and Presence Protocol (XMPP). | 2017-11-13 18:17:48 | 2021-07-13 04:03:55 |
| [go-feature-flag](https://thomaspoignant.github.io/go-feature-flag) | 341 | 10 | 9 | A simple and complete feature flag solution, without any complex backend system to install, all you need is a file as your backend. 🎛️ | 2020-12-11 13:19:17 | 2021-07-14 17:56:55 |
| [dudeldu](https://github.com/krotik/dudeldu) | 129 | 12 | 0 | A simple SHOUTcast server. | 2016-09-07 19:11:04 | 2021-07-12 07:03:25 |
| [lets-proxy2](https://github.com/rekby/lets-proxy2) | 56 | 6 | 31 | Reverse proxy with automatically obtains TLS certificates from Let's Encrypt | 2019-04-12 05:39:43 | 2021-07-03 06:26:03 |
| [psql-streamer](https://github.com/blind-oracle/psql-streamer) | 28 | 4 | 1 | Stream database events from PostgreSQL to Kafka | 2019-04-28 21:55:31 | 2021-05-01 06:49:46 |
| [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) | 22 | 3 | 0 | Turn Nginx logs into Prometheus metrics | 2018-10-23 09:10:27 | 2021-05-21 16:30:18 |
| [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) | 18 | 2 | 5 | Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support - S3, Google Cloud Storage, Azure Blob | 2019-12-18 12:48:14 | 2021-06-05 11:57:06 |
| [protoxy](https://github.com/camgraff/protoxy) | 16 | 1 | 0 | A proxy server than converts JSON request bodies to protocol buffers | 2020-09-03 23:24:34 | 2021-05-21 16:35:18 |
| [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) | 15 | 9 | 0 | Prometheus remote write proxy that adds Cortex tenant ID based on metric labels | 2020-10-06 16:52:25 | 2021-06-22 20:37:09 |
| [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) | 10 | 2 | 37 | Simple Reverse Proxy with Caching, written in Go, using Redis. | 2020-11-12 15:10:40 | 2021-07-14 20:08:07 |
| [riemann-relay](https://github.com/blind-oracle/riemann-relay) | 0 | 0 | 0 | Service for relaying Riemann events to Riemann/Carbon destinations | 2019-04-23 14:17:12 | 2019-10-29 15:00:17 |
</details>

### Stream Processing
Libraries and tools for stream processing and reactive programming.

<sup>*Last Update: 2021-07-15 09:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-streams](https://pkg.go.dev/github.com/reugn/go-streams) | 678 | 44 | 1 | A lightweight stream processing library for Go | 2019-04-30 17:28:15 | 2021-07-14 16:59:04 |
| [machine](https://pkg.go.dev/github.com/whitaker-io/machine) | 87 | 4 | 3 | Machine is a workflow/pipeline library for processing data | 2020-10-13 04:24:19 | 2021-06-24 10:56:50 |
| [stream](https://github.com/youthlin/stream) | 37 | 2 | 0 | Go Stream, like Java 8 Stream. | 2020-11-12 03:52:50 | 2021-05-31 13:41:51 |
</details>

### Template Engines
Libraries and tools for templating and lexing.

<sup>*Last Update: 2021-07-12 14:22:16*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gofpdf](http://godoc.org/github.com/jung-kurt/gofpdf) | 3,823 | 610 | 56 | A PDF document generator with high level support for text, drawing and images | 2015-03-13 11:57:30 | 2021-07-12 04:55:13 |
| [sprig](http://masterminds.github.io/sprig/) | 2,409 | 266 | 77 | Useful template functions for Go templates. | 2013-11-22 01:20:40 | 2021-07-12 07:07:39 |
| [quicktemplate](https://github.com/valyala/quicktemplate) | 2,137 | 117 | 27 | Fast, powerful, yet easy to use template engine for Go. Optimized for speed, zero memory allocations in hot paths. Up to 20x faster than html/template | 2016-03-06 21:42:01 | 2021-07-11 16:26:22 |
| [pongo2](https://www.schlachter.tech/pongo2) | 1,995 | 196 | 56 | Django-syntax like template-engine for Go | 2013-08-23 01:00:08 | 2021-07-09 20:48:10 |
| [hero](https://shiyanhui.github.io/hero) | 1,462 | 97 | 27 | A handy, fast and powerful go template engine. | 2017-01-15 13:31:50 | 2021-06-28 15:58:28 |
| [mustache](https://github.com/hoisie/mustache) | 1,013 | 187 | 32 | The mustache template language in Go | 2009-12-30 21:05:05 | 2021-05-07 07:50:12 |
| [amber](https://github.com/eknkc/amber) | 874 | 59 | 23 | Amber is an elegant templating engine for Go Programming Language, inspired from HAML and Jade | 2012-10-31 20:27:24 | 2021-06-26 05:56:45 |
| [ace](https://github.com/yosssi/ace) | 801 | 45 | 28 | HTML template engine for Go | 2014-07-13 13:39:19 | 2021-07-01 05:19:46 |
| [jet](https://shiyanhui.github.io/hero) | 798 | 80 | 17 | Jet  template engine | 2016-03-31 16:53:36 | 2021-07-08 12:42:33 |
| [gorazor](https://github.com/sipin/gorazor) | 782 | 87 | 2 | Razor view engine for go | 2014-05-01 05:30:31 | 2021-06-21 15:58:00 |
| [fasttemplate](https://github.com/valyala/fasttemplate) | 509 | 63 | 9 | Simple and fast template engine for Go | 2015-08-19 12:44:22 | 2021-07-08 20:05:23 |
| [ego](https://github.com/benbjohnson/ego) | 491 | 36 | 10 | An ERB-style templating language for Go. | 2014-02-23 18:14:41 | 2021-06-26 06:22:20 |
| [raymond](https://github.com/aymerick/raymond) | 432 | 62 | 18 | Handlebars for golang | 2015-04-22 13:07:59 | 2021-06-30 02:31:13 |
| [maroto](https://pkg.go.dev/github.com/johnfercher/maroto?tab=doc) | 408 | 72 | 18 | A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. | 2019-05-20 23:27:47 | 2021-07-11 11:17:19 |
| [goview](https://github.com/foolin/goview) | 237 | 24 | 10 | Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. | 2019-04-14 11:22:41 | 2021-07-04 12:46:26 |
| [soy](https://github.com/robfig/soy) | 155 | 38 | 6 | Go implementation for Soy templates (Google Closure templates) | 2013-12-15 01:14:48 | 2021-07-08 18:24:30 |
| [liquid](https://godoc.org/github.com/osteele/liquid) | 127 | 31 | 17 | A Liquid template engine in Go | 2017-06-26 14:39:52 | 2021-07-06 02:33:24 |
| [velvet](http://masterminds.github.io/sprig/) | 72 | 7 | 2 | A sweet velvety templating package | 2016-12-29 16:46:57 | 2021-06-25 10:31:41 |
| [kasia.go](https://github.com/ziutek/kasia.go) | 72 | 7 | 2 | Templating system for HTML and other text documents - go implementation | 2010-12-07 10:46:05 | 2021-05-05 02:28:01 |
| [extemplate](https://github.com/dannyvankooten/extemplate) | 38 | 9 | 2 | Wrapper package for Go's template/html to allow for easy file-based template inheritance. | 2018-08-10 20:34:19 | 2021-06-18 14:44:18 |
| [gospin](https://github.com/m1/gospin) | 27 | 6 | 3 | Article spinning and spintax/spinning syntax engine written in Go, useful for A/B, testing pieces of text/articles and creating more natural conversations | 2019-02-22 17:04:51 | 2021-05-21 04:32:26 |
| [damsel](https://github.com/dskinner/damsel) | 24 | 4 | 1 | Package damsel provides html outlining via css-selectors and common template functionality. | 2012-05-02 23:06:48 | 2020-09-12 23:20:49 |
</details>

### Testing - Fail injection


<sup>*Last Update: 2021-06-28 16:25:27*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [failpoint](https://github.com/pingcap/failpoint) | 590 | 52 | 10 | An implementation of failpoints for Golang. | 2019-04-02 07:48:18 | 2021-06-28 03:31:07 |
</details>

### Testing - Fuzzing and delta-debugging, reducing, shrinking


<sup>*Last Update: 2021-07-15 09:25:18*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-fuzz](https://github.com/dvyukov/go-fuzz) | 4,081 | 236 | 51 | Randomized testing for Go | 2015-04-15 13:07:50 | 2021-07-14 15:50:12 |
| [gofuzz](https://github.com/google/gofuzz) | 1,101 | 103 | 12 | Fuzz testing for go. | 2014-07-31 16:21:29 | 2021-07-13 13:52:07 |
| [tavor](https://github.com/zimmski/tavor) | 232 | 9 | 53 | A generic fuzzing and delta-debugging framework | 2014-05-18 14:59:14 | 2021-04-27 23:55:43 |
</details>

### Testing - Mock


<sup>*Last Update: 2021-06-22 09:01:46*</sup>
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


<sup>*Last Update: 2021-07-15 09:25:18*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [chromedp](https://github.com/chromedp/chromedp) | 6,463 | 539 | 31 | A faster, simpler way to drive browsers supporting the Chrome DevTools Protocol. | 2017-01-24 14:54:30 | 2021-07-12 05:29:52 |
| [selenoid](https://aerokube.com/selenoid/latest/) | 1,941 | 253 | 164 | Selenium Hub successor running browsers within containers. Scalable, immutable, self hosted Selenium-Grid on any platform with single binary. | 2016-08-22 09:11:16 | 2021-07-14 08:29:55 |
| [rod](https://go-rod.github.io) | 1,643 | 110 | 48 | A Devtools driver for web automation and scraping | 2020-01-21 20:09:45 | 2021-07-11 04:05:40 |
| [cdp](https://github.com/mafredri/cdp) | 561 | 39 | 12 | Package cdp provides type-safe bindings for the Chrome DevTools Protocol (CDP), written in the Go programming language. | 2017-03-12 10:25:41 | 2021-07-08 23:12:35 |
| [playwright-go](https://mxschmitt.github.io/playwright-go/) | 421 | 43 | 19 | Playwright for Go a browser automation library to control Chromium, Firefox and WebKit with a single API. | 2020-08-16 12:46:14 | 2021-07-11 07:40:16 |
| [ggr](https://aerokube.com/ggr/latest/) | 273 | 57 | 15 | A lightweight load balancer used to create big Selenium clusters | 2016-06-16 15:33:24 | 2021-07-09 15:17:08 |
</details>

### Testing - Testing Frameworks


<sup>*Last Update: 2021-07-02 09:03:50*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [testify](https://github.com/stretchr/testify) | 13,665 | 1,098 | 286 | A toolkit with common assertions and mocks that plays nicely with the standard library | 2012-10-16 16:43:17 | 2021-07-02 01:49:06 |
| [goconvey](http://goconvey.co) | 6,440 | 471 | 139 | Go testing in the browser. Integrates with `go test`. Write behavioral tests in Go. | 2013-08-21 04:52:28 | 2021-06-26 11:32:01 |
| [ginkgo](http://onsi.github.io/ginkgo/) | 4,871 | 442 | 58 | BDD Testing Framework for Go | 2013-08-23 20:53:05 | 2021-06-25 11:57:14 |
| [go-cmp](https://github.com/google/go-cmp) | 2,349 | 141 | 11 | Package for comparing Go values in tests | 2017-07-07 19:28:22 | 2021-06-26 04:56:06 |
| [httpexpect](https://github.com/gavv/httpexpect) | 1,713 | 133 | 9 | End-to-end HTTP and REST API testing for Go. | 2016-04-29 17:05:20 | 2021-06-28 15:56:31 |
| [gomega](http://onsi.github.io/gomega/) | 1,410 | 215 | 18 | Ginkgo's Preferred Matcher Library | 2013-08-23 20:54:42 | 2021-07-01 15:18:51 |
| [godog](https://github.com/cucumber/godog) | 1,367 | 148 | 57 | Cucumber for golang | 2015-06-10 13:16:31 | 2021-06-26 11:18:18 |
| [goblin](https://github.com/franela/goblin) | 800 | 73 | 17 | Minimal and Beautiful Go testing framework | 2013-09-19 02:34:24 | 2021-06-25 10:21:43 |
| [go-vcr](https://go-testdeep.zetta.rocks/) | 724 | 50 | 4 | Record and replay your HTTP interactions for fast, deterministic and accurate tests | 2015-12-14 12:52:17 | 2021-06-26 08:47:53 |
| [baloo](https://godoc.org/github.com/h2non/baloo) | 712 | 27 | 8 | Expressive end-to-end HTTP API testing made easy in Go | 2016-05-29 21:40:58 | 2021-06-14 14:54:20 |
| [testfixtures](https://pkg.go.dev/github.com/go-testfixtures/testfixtures/v3?tab=doc) | 677 | 51 | 13 | Ruby on Rails like test fixtures for Go. Write tests against a real database | 2016-04-05 11:33:28 | 2021-06-30 18:47:47 |
| [gnomock](https://github.com/orlangure/gnomock) | 471 | 14 | 7 | Test your code without writing mocks with ephemeral Docker containers 📦 Setup popular services with just a couple lines of code ⏱️ No bash, no yaml, only code 💻 | 2020-01-31 14:50:52 | 2021-06-25 16:53:40 |
| [go-mutesting](https://github.com/zimmski/go-mutesting) | 400 | 34 | 26 | Mutation testing for Go source code | 2014-12-26 22:23:44 | 2021-06-25 01:27:40 |
| [goc](https://github.com/qiniu/goc) | 383 | 54 | 16 | A Comprehensive Coverage Testing System for The Go Programming Language | 2020-05-07 03:46:25 | 2021-06-25 06:14:51 |
| [gofight](https://github.com/appleboy/gofight) | 373 | 35 | 8 | Testing API Handler written in Golang. | 2016-03-29 00:13:21 | 2021-06-15 18:48:22 |
| [apitest](https://apitest.dev) | 358 | 35 | 1 | A simple and extensible behavioural testing library for Go. You can use api test to simplify REST API, HTTP handler and e2e tests. | 2018-12-26 22:27:19 | 2021-06-23 13:55:50 |
| [frisby](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 264 | 26 | 13 | API testing framework inspired by frisby-js | 2015-09-15 14:35:58 | 2021-06-21 15:06:03 |
| [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) | 254 | 15 | 1 | Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test | 2019-11-16 23:49:40 | 2021-06-25 20:55:53 |
| [gotest.tools](https://gotest.tools) | 244 | 35 | 22 | A collection of packages to augment the go testing package and support common patterns. | 2017-08-08 21:28:54 | 2021-07-01 12:37:09 |
| [go-carpet](https://github.com/msoap/go-carpet) | 214 | 7 | 1 | go-carpet - show test coverage in terminal for Go source files | 2016-02-28 12:02:51 | 2021-05-21 04:32:28 |
| [charlatan](https://github.com/percolate/charlatan) | 194 | 8 | 2 | Go Interface Mocking Tool | 2017-10-06 21:55:14 | 2021-05-08 03:35:26 |
| [endly](https://github.com/viant/endly) | 189 | 20 | 0 | End to end functional test and automation framework | 2017-08-28 20:24:43 | 2021-06-22 18:23:38 |
| [commander](https://github.com/commander-cli/commander) | 186 | 12 | 22 | Test your command line interfaces on windows, linux and osx and nodes viá ssh and docker | 2019-02-22 16:35:16 | 2021-05-29 04:12:20 |
| [go-testdeep](https://go-testdeep.zetta.rocks/) | 171 | 5 | 3 | Extremely flexible golang deep comparison, extends the go testing package, tests HTTP APIs and provides tests suite | 2018-05-26 15:03:28 | 2021-06-18 11:17:45 |
| [cupaloy](https://github.com/bradleyjkemp/cupaloy) | 166 | 21 | 9 | Simple Go snapshot testing | 2017-08-07 18:30:05 | 2021-06-25 16:32:32 |
| [dbcleaner](https://github.com/khaiql/dbcleaner) | 130 | 11 | 1 | Clean database for testing, inspired by database_cleaner for Ruby | 2017-01-17 18:18:40 | 2021-06-18 21:07:13 |
| [gospec](https://github.com/luontola/gospec) | 114 | 16 | 3 | Testing framework for Go. Allows writing self-documenting tests/specifications, and executes them concurrently and safely isolated. [UNMAINTAINED] | 2009-11-24 13:59:26 | 2021-02-11 17:23:11 |
| [wstest](https://github.com/posener/wstest) | 84 | 13 | 1 | go websocket client for unit testing of a websocket handler | 2017-03-31 21:06:18 | 2021-05-24 08:41:09 |
| [gocrest](https://github.com/corbym/gocrest) | 80 | 4 | 2 | GoCrest - Hamcrest-like matchers for Go | 2017-12-23 23:27:00 | 2021-05-14 09:54:10 |
| [testcase](https://github.com/adamluzsi/testcase) | 69 | 2 | 0 | testcase is an opinionated behavior-driven-testing library | 2019-04-22 21:20:51 | 2021-07-01 08:45:15 |
| [jsonassert](https://github.com/kinbiko/jsonassert) | 68 | 11 | 4 | A Go test assertion library for verifying that two representations of JSON are semantically equal | 2018-10-26 20:31:01 | 2021-06-24 08:32:17 |
| [restit](https://github.com/go-restit/restit) | 53 | 4 | 4 | A Go library help testing your RESTful API application | 2014-06-25 10:25:46 | 2021-04-30 17:01:37 |
| [gospecify](https://github.com/stesla/gospecify) | 53 | 6 | 1 | A BDD library for Go | 2009-11-20 06:34:29 | 2021-02-11 17:23:22 |
| [gomatch](https://github.com/jfilipczyk/gomatch) | 40 | 3 | 0 | Library created for testing JSON against patterns. | 2019-01-27 20:19:06 | 2021-05-15 22:30:58 |
| [covergates](https://covergates.com) | 39 | 9 | 11 | The portal gates to coverage reports | 2020-05-29 04:02:01 | 2021-05-20 11:01:00 |
| [dsunit](https://github.com/viant/dsunit) | 38 | 6 | 0 | Datastore Testibility | 2016-06-13 20:20:52 | 2021-02-14 11:12:04 |
| [go-hit](https://github.com/Eun/go-hit) | 34 | 1 | 8 | http integration test framework | 2019-06-04 16:28:23 | 2021-06-07 07:32:59 |
| [assert](https://github.com/go-playground/assert) | 33 | 9 | 2 | :exclamation:Basic Assertion Library used along side native go testing, with building blocks for custom assertions | 2015-07-20 17:53:45 | 2021-06-22 09:01:30 |
| [hamcrest](https://github.com/rdrdr/hamcrest) | 26 | 4 | 2 | Hamcrest matchers for the Go programming language | 2010-12-22 04:49:44 | 2021-02-11 17:08:31 |
| [flute](https://pkg.go.dev/github.com/suzuki-shunsuke/flute/flute) | 16 | 0 | 4 | Golang HTTP client testing framework | 2019-07-06 04:32:03 | 2021-05-21 04:32:27 |
| [schema](https://github.com/jgroeneveld/schema) | 15 | 0 | 0 | Quick and easy expression matching for JSON schemas used in requests and responses | 2015-08-13 13:36:54 | 2021-03-01 05:46:30 |
| [gogiven](https://github.com/corbym/gogiven) | 10 | 2 | 4 | gogiven - BDD testing framework for go that generates readable output directly from source code | 2017-12-31 22:33:37 | 2021-06-03 09:08:48 |
| [testsql](https://github.com/zhulongcheng/testsql) | 10 | 1 | 3 | Generate test data from SQL files before testing and clear it after finished. | 2018-09-22 12:03:50 | 2020-08-30 23:10:29 |
| [biff](https://github.com/fulldump/biff) | 10 | 1 | 0 | Bifurcation Framework for testing and use cases | 2018-03-28 18:35:53 | 2020-08-31 12:00:00 |
| [badio](https://github.com/cavaliercoder/badio) | 10 | 1 | 0 | Extensions to Go's testing/iotest package | 2016-02-11 10:29:25 | 2020-08-31 08:36:30 |
| [gosuite](https://github.com/pavlo/gosuite) | 10 | 3 | 1 | Test suites support for standard Go1.7 "testing" by leveraging Subtests feature | 2016-10-15 19:28:14 | 2021-06-01 17:30:40 |
| [test](https://github.com/tvastar/test) | 7 | 0 | 0 | test utilities for golang | 2019-03-23 21:47:36 | 2020-01-14 15:42:38 |
| [stop-and-go](https://github.com/elgohr/stop-and-go) | 5 | 2 | 0 | Testing helper for concurrency | 2020-11-06 09:04:58 | 2021-01-31 11:58:58 |
| [trial](https://github.com/jgroeneveld/trial) | 5 | 0 | 0 | A simple assertion library for go | 2015-06-18 09:01:30 | 2021-04-18 16:15:16 |
| [tt](https://github.com/vcaesar/tt) | 5 | 0 | 0 | Simple and colorful test tools | 2018-04-03 11:47:21 | 2021-06-06 17:44:13 |
</details>

### Text Processing - Specific Formats


<sup>*Last Update: 2021-06-28 02:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [colly](http://go-colly.org/) | 14,139 | 1,211 | 115 | Elegant Scraper and Crawler Framework for Golang | 2017-09-29 14:08:49 | 2021-06-27 10:21:31 |
| [goquery](https://godoc.org/astuart.co/goq) | 10,311 | 778 | 3 | A little like that j-thing, only in Go. | 2012-08-29 02:14:59 | 2021-06-27 08:55:52 |
| [blackfriday](https://github.com/russross/blackfriday) | 4,736 | 568 | 195 | Blackfriday: a markdown processor for Go | 2011-05-27 22:28:58 | 2021-06-25 16:31:49 |
| [sh](https://pkg.go.dev/mvdan.cc/sh/v3) | 3,820 | 192 | 51 | A shell parser, formatter, and interpreter with bash support; includes shfmt | 2016-01-16 08:39:09 | 2021-06-26 10:08:51 |
| [toml](https://github.com/BurntSushi/toml) | 3,495 | 455 | 20 | TOML parser for Golang with reflection. | 2013-02-26 05:05:48 | 2021-06-27 00:42:55 |
| [go-humanize](https://pkg.go.dev/github.com/dustin/go-humanize) | 2,684 | 182 | 28 | Go Humans! (formatters for units to human friendly sizes) | 2012-01-13 03:48:55 | 2021-06-26 21:21:56 |
| [bluemonday](https://github.com/microcosm-cc/bluemonday) | 1,901 | 119 | 15 | bluemonday: a fast golang HTML sanitizer (inspired by the OWASP Java HTML Sanitizer) to scrub user generated content of XSS | 2013-11-20 22:15:49 | 2021-06-24 20:17:41 |
| [gofeed](https://github.com/mmcdole/gofeed) | 1,655 | 154 | 42 | Parse RSS, Atom and JSON feeds in Go | 2016-01-23 02:44:34 | 2021-06-26 17:55:24 |
| [inject](https://godoc.org/github.com/facebookgo/inject) | 1,329 | 115 | 9 | Package inject provides a reflect based injector. | 2013-10-21 01:51:46 | 2021-06-24 02:21:43 |
| [go-toml](https://github.com/pelletier/go-toml) | 1,049 | 148 | 28 | Go library for the TOML file format | 2013-02-24 17:45:51 | 2021-06-25 04:18:32 |
| [commonregex](https://github.com/mingrammer/commonregex) | 774 | 57 | 2 | 🍫 A collection of common regular expressions for Go | 2017-03-23 14:33:18 | 2021-06-23 07:27:53 |
| [slug](https://pkg.go.dev/mvdan.cc/sh/v3) | 654 | 69 | 11 | URL-friendly slugify with multiple languages support. | 2014-03-31 06:24:51 | 2021-06-26 04:00:22 |
| [mxj](https://github.com/clbanning/mxj) | 469 | 86 | 0 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2014-02-03 13:39:16 | 2021-06-05 07:21:45 |
| [dataflowkit](https://dataflowkit.com) | 466 | 62 | 0 | Extract structured data from web sites. Web sites scraping.   | 2017-02-09 15:08:15 | 2021-06-12 06:59:45 |
| [gographviz](https://github.com/awalterschulze/gographviz) | 434 | 65 | 9 | Parses the Graphviz DOT language in golang | 2015-03-14 18:27:00 | 2021-06-25 10:33:46 |
| [gommon](https://github.com/labstack/gommon) | 397 | 85 | 15 | Common packages for Go | 2015-03-12 22:35:57 | 2021-06-24 17:56:00 |
| [htmlquery](https://github.com/antchfx/xpath) | 372 | 46 | 7 | htmlquery is golang XPath package for HTML query. | 2017-12-05 01:08:41 | 2021-06-24 04:35:37 |
| [go-runewidth](https://github.com/mattn/go-runewidth) | 370 | 62 | 9 | wcwidth for golang | 2013-06-21 04:56:50 | 2021-06-23 14:15:47 |
| [omniparser](https://github.com/jf-tech/omniparser) | 352 | 14 | 0 | omniparser: a native Golang ETL streaming parser and transform library for CSV, JSON, XML, EDI, text, etc. | 2020-08-16 22:22:21 | 2021-06-24 15:31:33 |
| [gotext](https://github.com/zhshch2002/gospider) | 301 | 34 | 5 | Go (Golang) GNU gettext utilities package  | 2016-06-19 20:14:43 | 2021-06-17 23:23:05 |
| [goq](https://godoc.org/astuart.co/goq) | 196 | 18 | 2 | A declarative struct-tag-based HTML unmarshaling or scraping package for Go built on top of the goquery library | 2017-02-20 02:54:40 | 2021-05-25 05:36:42 |
| [goribot](https://github.com/zhshch2002/gospider) | 195 | 29 | 1 | [Crawler/Scraper for Golang]🕷A lightweight distributed friendly Golang crawler framework.一个轻量的分布式友好的 Golang 爬虫框架。 | 2019-09-08 10:39:47 | 2021-05-19 12:04:06 |
| [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) | 184 | 33 | 4 | ⚙️ Convert HTML to Markdown. Even works with entire websites and can be extended through rules. | 2018-05-15 13:26:26 | 2021-06-22 06:09:21 |
| [go-nmea](https://github.com/adrianmo/go-nmea) | 151 | 51 | 3 | A NMEA parser library in pure Go | 2015-07-22 08:55:54 | 2021-06-27 02:16:30 |
| [github_flavored_markdown](https://github.com/shurcooL/github_flavored_markdown) | 129 | 32 | 9 | GitHub Flavored Markdown renderer with fenced code block highlighting, clickable header anchor links. | 2015-05-16 04:09:07 | 2021-06-11 04:53:29 |
| [sdp](https://github.com/gortc/sdp) | 111 | 29 | 5 | RFC 4566 SDP implementation in go | 2016-05-13 14:35:11 | 2021-06-05 14:30:12 |
| [go-zero-width](https://github.com/trubitsyn/go-zero-width) | 96 | 8 | 0 | Zero-width character detection and removal for Go | 2018-06-18 13:55:09 | 2021-05-26 03:02:08 |
| [podcast](https://github.com/eduncan911/podcast) | 94 | 26 | 5 | iTunes and RSS 2.0 Podcast Generator in Golang | 2017-02-02 12:45:04 | 2021-06-12 11:02:17 |
| [editorconfig-core-go](https://godoc.org/github.com/hscells/doi) | 85 | 28 | 4 | EditorConfig Core written in Go | 2016-07-05 03:50:41 | 2021-05-30 13:07:53 |
| [align](https://github.com/Guitarbum722/align) | 70 | 5 | 0 | A general purpose application and library for aligning text. | 2017-04-29 23:22:22 | 2021-06-19 06:08:20 |
| [go-slugify](https://godoc.org/github.com/mozillazg/go-slugify) | 67 | 5 | 1 | Pretty Slug. | 2016-07-16 11:55:15 | 2021-04-30 21:15:15 |
| [genex](https://namegrep.com/) | 60 | 5 | 0 | Genex package for Go | 2015-03-09 19:24:16 | 2020-12-10 08:56:04 |
| [go-vcard](https://github.com/pelletier/go-toml) | 58 | 19 | 4 | A Go library to parse and format vCard | 2017-03-21 08:30:36 | 2021-06-19 05:38:04 |
| [goregen](https://godoc.org/github.com/zach-klippenstein/goregen) | 57 | 5 | 4 | randexp for Go. | 2014-12-27 00:19:39 | 2021-06-06 05:03:43 |
| [go-fixedwidth](http://godoc.org/github.com/ianlopshire/go-fixedwidth) | 54 | 20 | 8 | Encoding and decoding for fixed-width formatted data | 2017-11-15 21:05:44 | 2021-06-15 20:23:55 |
| [guesslanguage](https://github.com/zhshch2002/gospider) | 53 | 2 | 1 | Guess the natural language of a text in Go | 2014-12-16 10:58:47 | 2021-06-10 08:36:39 |
| [did](https://w3c-ccg.github.io/did-spec) | 52 | 13 | 4 | A golang package to work with Decentralized Identifiers (DIDs)  | 2018-11-02 17:49:14 | 2021-06-23 07:40:00 |
| [allot](https://github.com/sbstjn/allot) | 49 | 6 | 2 | Parse placeholder and wildcard text commands | 2016-10-16 15:49:08 | 2021-05-21 04:32:31 |
| [pagser](https://github.com/foolin/pagser) | 36 | 2 | 2 | Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler | 2020-04-19 09:22:00 | 2021-06-20 14:14:58 |
| [gonameparts](https://github.com/polera/gonameparts) | 33 | 3 | 2 | Takes a full name and splits it into individual name parts | 2015-05-17 05:20:17 | 2020-11-25 21:46:02 |
| [slugify](https://pkg.go.dev/mvdan.cc/sh/v3) | 27 | 2 | 0 | A Go slugify application that handles string | 2015-04-13 01:54:30 | 2021-02-17 19:13:25 |
| [normalize](https://github.com/avito-tech/normalize) | 18 | 1 | 0 | Decode / encode XML to/from map[string]interface{} (or JSON); extract values with dot-notation paths and wildcards.  Replaces x2j and j2x packages. | 2021-03-22 09:25:14 | 2021-06-25 19:29:34 |
| [codetree](https://github.com/aerogo/codetree) | 17 | 3 | 0 | :evergreen_tree: Parses indented code and returns a tree structure. | 2016-11-26 02:50:38 | 2021-02-06 17:49:49 |
| [enca](https://godoc.org/github.com/hscells/doi) | 11 | 2 | 0 | Minimal cgo bindings for libenca | 2014-12-17 04:55:16 | 2020-12-16 16:44:00 |
| [syndfeed](https://github.com/zhengchun/syndfeed) | 7 | 2 | 0 | A syndication feed parser for Atom 1.0 and RSS 2.0 in Go | 2017-04-07 09:30:55 | 2019-10-28 23:12:30 |
| [doi](https://godoc.org/github.com/hscells/doi) | 6 | 0 | 0 | Parse and check doi objects in go. | 2017-08-02 05:58:01 | 2020-12-31 02:08:59 |
| [bbConvert](https://github.com/CalebQ42/bbConvert) | 6 | 1 | 0 | Converter from BBCode to HTML | 2016-04-15 14:35:38 | 2020-08-30 14:37:50 |
| [ltsv](https://github.com/Wing924/ltsv) | 5 | 0 | 0 | High performance LTSV (Labeled Tab Separeted Value) reader for Go. | 2019-05-12 06:11:04 | 2020-04-28 14:07:48 |
| [encoding](https://github.com/mickep76/encoding) | 3 | 1 | 1 | Go package provides a generic interface to encoders and decoders | 2018-04-06 20:48:00 | 2019-11-12 13:29:44 |
| [go-output-format](https://github.com/drewstinnett/go-output-format) | 2 | 0 | 0 | Output go objects in standard formats, such as YAML, JSON, etc | 2021-04-08 20:48:17 | 2021-05-21 04:32:32 |
| [go-wildcard](https://github.com/pelletier/go-toml) | 2 | 0 | 0 | A Go library to parse and format vCard | 2021-03-28 16:31:41 | 2021-06-12 19:53:09 |
</details>

### Text Processing - Utility


<sup>*Last Update: 2021-07-15 09:25:14*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [xurls](https://tysug.net) | 799 | 93 | 1 | Extract urls from text | 2015-01-12 01:28:46 | 2021-07-13 14:47:26 |
| [gotabulate](https://github.com/bndr/gotabulate) | 267 | 27 | 6 | Gotabulate - Easily pretty-print your tabular data with Go | 2014-08-21 07:44:28 | 2021-07-05 14:37:13 |
| [radix](https://github.com/yourbasic/radix) | 173 | 9 | 0 | A fast string sorting algorithm (MSD radix sort) | 2017-06-09 14:38:58 | 2021-06-28 03:33:25 |
| [regroup](https://github.com/oriser/regroup) | 87 | 6 | 1 | Match regex group into go struct using struct tags and automatic parsing | 2020-09-08 19:04:42 | 2021-07-09 10:45:56 |
| [parth](https://github.com/codemodus/parth) | 40 | 4 | 0 | Path parsing for segment unmarshaling and slicing. | 2015-04-06 22:53:59 | 2021-03-27 05:04:42 |
| [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) | 37 | 3 | 1 | A sanitization-based swear filter for Go. | 2018-09-09 00:07:26 | 2021-07-03 21:08:29 |
| [xj2go](https://tysug.net) | 20 | 5 | 0 | Convert xml and json to go struct | 2017-09-19 13:20:57 | 2021-07-07 13:14:07 |
| [tagify](https://www.zoomio.org/tagify) | 17 | 0 | 0 | Tagify produces a set of tags from a given source. Source can be either an HTML page, a Markdown document or a plain text. Supports English, Russian, Chinese, Hindi, Spanish, Arabic, Japanese, German, Hebrew, French and Korean languages. | 2018-03-20 10:30:11 | 2021-07-04 11:15:17 |
| [kace](https://github.com/codemodus/kace) | 16 | 1 | 1 | Common case conversions covering common initialisms. | 2015-06-04 20:36:49 | 2020-09-27 22:30:36 |
| [TySug](https://tysug.net) | 10 | 1 | 0 | A project around helping to prevent typing typos. TySug (Typo Suggestions) suggests alternative words with respect to keyboard layouts | 2018-06-05 19:46:29 | 2021-06-29 10:37:35 |
| [parseargs-go](https://github.com/txgruppi/parseargs-go) | 8 | 3 | 1 | A string argument parser that understands quotes and backslashes | 2016-02-24 00:53:38 | 2020-05-16 17:37:29 |
| [textwrap](https://www.zoomio.org/tagify) | 2 | 0 | 1 | Port of Python's "textwrap" module to Go | 2019-07-26 17:57:55 | 2020-12-22 12:39:59 |
</details>

### Third-party APIs
Libraries for accessing third party APIs.

<sup>*Last Update: 2021-07-03 02:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-github](https://pkg.go.dev/github.com/google/go-github/v35/github) | 7,569 | 1,541 | 99 | Go library for accessing the GitHub API | 2013-05-24 16:42:58 | 2021-07-02 00:38:17 |
| [aws-sdk-go](http://aws.amazon.com/sdk-for-go/) | 6,988 | 1,695 | 103 | AWS SDK for the Go programming language. | 2014-12-05 05:29:41 | 2021-07-01 23:42:37 |
| [slack](https://pkg.go.dev/github.com/slack-go/slack) | 3,572 | 868 | 62 | Slack API in Go - community-maintained fork created by the original author, @nlopes | 2015-01-24 14:19:00 | 2021-07-02 08:29:18 |
| [google-api-go-client](https://pkg.go.dev/google.golang.org/api) | 2,696 | 833 | 28 | Auto-generated Google APIs for Go. | 2014-11-24 21:45:36 | 2021-07-01 11:58:32 |
| [google-cloud-go](https://pkg.go.dev/cloud.google.com/go) | 2,582 | 905 | 211 | Google Cloud Client Libraries for Go. | 2014-05-09 11:11:58 | 2021-07-02 00:42:26 |
| [discordgo](http://bwmarrin.github.io/discordgo/) | 2,135 | 416 | 83 |  (Golang) Go bindings for Discord | 2015-11-01 20:51:01 | 2021-07-01 17:41:59 |
| [stripe-go](https://stripe.com) | 1,381 | 369 | 13 | Go library for the Stripe API.     | 2014-06-05 23:38:14 | 2021-06-30 22:57:31 |
| [minio-go](https://docs.min.io/docs/golang-client-quickstart-guide.html) | 1,323 | 404 | 14 | MinIO Client SDK for Go | 2015-05-02 02:36:46 | 2021-07-02 07:56:17 |
| [go-twitter](https://dev.twitter.com/rest/public) | 1,266 | 247 | 30 | Go Twitter REST and Streaming API v1.1 | 2015-04-11 23:26:07 | 2021-06-30 00:05:40 |
| [anaconda](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 1,101 | 253 | 72 | A Go client library for the Twitter 1.1 API | 2013-03-04 22:46:07 | 2021-06-28 23:33:20 |
| [facebook](https://github.com/huandu/facebook) | 978 | 359 | 0 | A Facebook Graph API SDK For Go. | 2012-07-28 19:05:56 | 2021-06-30 16:21:47 |
| [go-jira](https://pkg.go.dev/github.com/andygrunwald/go-jira?tab=doc) | 918 | 307 | 58 | Go client library for Atlassian Jira | 2015-08-20 15:02:46 | 2021-07-01 05:04:57 |
| [githubv4](https://github.com/shurcooL/githubv4) | 787 | 59 | 32 | Package githubv4 is a client library for accessing GitHub GraphQL API v4 (https://docs.github.com/en/graphql). | 2017-05-27 05:05:31 | 2021-06-27 08:13:53 |
| [webhooks](https://github.com/go-playground/webhooks) | 630 | 152 | 26 | :fishing_pole_and_fish: Webhook receiver for GitHub, Bitbucket, GitLab, Gogs | 2015-10-25 17:38:13 | 2021-06-24 02:19:25 |
| [paypal](https://github.com/plutov/paypal) | 424 | 187 | 2 | Golang client for PayPal REST API | 2015-10-14 04:57:49 | 2021-07-01 07:01:50 |
| [geo-golang](https://github.com/codingsince1985/geo-golang) | 405 | 45 | 5 | Go library to access geocoding and reverse geocoding APIs | 2014-12-04 08:18:31 | 2021-06-29 16:07:20 |
| [ethrpc](http://bwmarrin.github.io/discordgo/) | 209 | 87 | 10 | Golang client for ethereum json rpc api | 2017-01-24 09:47:00 | 2021-06-13 10:21:59 |
| [go-marathon](https://github.com/gambol99/go-marathon) | 196 | 129 | 26 | A GO API library for working with Marathon | 2015-02-11 13:25:26 | 2021-01-22 14:11:14 |
| [trello](https://github.com/adlio/trello) | 171 | 67 | 12 | Trello API wrapper for Go | 2016-09-24 04:36:10 | 2021-06-23 22:25:36 |
| [medium-sdk-go](https://medium.com) | 130 | 20 | 6 | A Golang SDK for Medium's OAuth2 API | 2015-09-26 23:45:46 | 2021-06-22 23:20:42 |
| [gostorm](https://github.com/jsgilmore/gostorm) | 127 | 20 | 5 | GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. | 2013-07-22 12:43:41 | 2021-06-03 17:46:29 |
| [twitter-scraper](https://github.com/n0madic/twitter-scraper) | 123 | 32 | 4 | Scrape the Twitter Frontend API without authentication with Golang. | 2018-11-29 15:31:50 | 2021-07-02 11:55:57 |
| [go-trending](http://godoc.org/github.com/andygrunwald/go-trending) | 116 | 17 | 2 | Go library for accessing trending repositories and developers at Github. | 2015-07-04 08:06:48 | 2021-06-22 17:00:21 |
| [hipchat](https://github.com/daneharrigan/hipchat) | 111 | 36 | 3 | A golang package to communicate with HipChat over XMPP | 2013-04-28 02:16:21 | 2021-06-10 11:52:19 |
| [wit-go](https://github.com/wit-ai/wit-go) | 105 | 24 | 0 | Go client for wit.ai HTTP API | 2018-08-20 07:18:40 | 2021-06-25 06:02:39 |
| [hipchat](https://github.com/andybons/hipchat) | 104 | 21 | 0 | This project implements a Go client library for the Hipchat API. | 2012-10-20 18:34:06 | 2021-06-10 11:52:17 |
| [pushover](https://github.com/gregdel/pushover) | 100 | 10 | 1 | Go wrapper for the Pushover API | 2015-02-19 15:30:05 | 2021-06-26 23:51:53 |
| [cachet](https://github.com/andygrunwald/cachet) | 88 | 12 | 1 | Go(lang) client library for Cachet (open source status page system). | 2015-10-31 12:30:07 | 2021-06-22 17:03:44 |
| [igdb](https://api.igdb.com/) | 67 | 10 | 3 | Go client for the Internet Game Database API | 2017-08-24 08:31:53 | 2021-06-14 07:28:34 |
| [gosip](https://go.spflow.com) | 60 | 17 | 10 | ⚡️ SharePoint authentication, HTTP client & fluent API wrapper for Go (Golang) | 2019-01-26 08:48:48 | 2021-06-29 11:55:17 |
| [go-circleci](https://github.com/jszwedko/go-circleci) | 59 | 46 | 5 | Go library for interacting with CircleCI | 2015-08-14 21:19:36 | 2021-05-17 19:26:50 |
| [simples3](https://github.com/rhnvrm/simples3) | 56 | 8 | 0 | Simple no frills AWS S3 Golang Library using REST with V4 Signing (without AWS Go SDK) | 2018-12-06 10:24:21 | 2021-05-30 10:06:07 |
| [gogtrends](https://github.com/groovili/gogtrends) | 55 | 15 | 0 | Unofficial Google Trends API for Go | 2018-12-27 13:50:34 | 2021-06-26 20:47:16 |
| [clarifai-go](https://github.com/Clarifai/clarifai-go) | 55 | 12 | 8 | Clarifai library for Go | 2015-09-28 23:33:59 | 2021-06-18 06:11:31 |
| [megos](https://godoc.org/github.com/andygrunwald/megos) | 55 | 10 | 0 | Go(lang) client library for accessing information of an Apache Mesos cluster. | 2015-10-02 14:29:20 | 2021-06-25 15:22:05 |
| [go-unsplash](https://unsplash.com) | 53 | 9 | 8 | Go Client for the Unsplash API  | 2017-01-19 07:04:04 | 2021-06-09 06:01:06 |
| [go-amazon-product-advertising-api](https://godoc.org/github.com/ngs/go-amazon-product-advertising-api/amazon) | 49 | 15 | 3 | Go Client Library for Amazon Product Advertising API | 2016-11-15 15:37:32 | 2021-06-13 13:35:55 |
| [gads](https://github.com/emiddleton/gads) | 49 | 57 | 8 | Google Adwords API for Go | 2014-01-20 02:22:15 | 2020-10-05 08:15:45 |
| [ynab.go](https://godoc.org/go.bmvs.io/ynab) | 49 | 12 | 5 | Go client for the YNAB API. Unofficial. It covers 100% of the resources made available by the YNAB API. | 2018-07-13 11:10:54 | 2021-06-13 02:39:00 |
| [uptimerobot](https://github.com/bitfield/uptimerobot) | 46 | 10 | 12 | Client library for UptimeRobot v2 API | 2018-05-29 10:27:19 | 2021-06-29 20:50:41 |
| [go-xkcd](https://pkg.go.dev/github.com/nishanths/go-xkcd/v2) | 45 | 3 | 0 | xkcd.com API client in Go | 2016-02-26 05:14:31 | 2021-05-10 11:16:35 |
| [golang-tmdb](https://github.com/cyruzin/golang-tmdb) | 43 | 8 | 0 | This is a Golang wrapper for working with TMDb API. It aims to support version 3. | 2019-01-11 22:59:33 | 2021-06-09 18:15:11 |
| [gomusicbrainz](https://github.com/michiwend/gomusicbrainz) | 43 | 14 | 5 | a Go (Golang) MusicBrainz WS2 client library - work in progress | 2014-09-10 16:42:33 | 2021-05-08 12:29:00 |
| [fcm](https://github.com/maddevsio/fcm) | 40 | 14 | 2 | Firebase Cloud Messaging for application servers implemented using the Go programming language. | 2017-01-06 08:30:57 | 2021-06-22 14:06:40 |
| [go-spotify](https://pkg.go.dev/github.com/slack-go/slack) | 38 | 4 | 0 | Go library for the Spotify Web API | 2014-10-30 02:52:04 | 2021-06-05 16:50:55 |
| [mixpanel](https://github.com/dukex/mixpanel) | 38 | 19 | 3 | Golang Mixpanel Client | 2014-05-20 03:50:34 | 2021-05-09 15:12:04 |
| [golyrics](https://github.com/mamal72/golyrics) | 36 | 1 | 0 | A simple Go package to fetch lyrics from Wikia | 2016-11-18 04:40:37 | 2021-01-03 15:56:41 |
| [translate](https://github.com/nuveo/translate) | 32 | 6 | 0 | Go online translation package | 2015-07-13 15:42:13 | 2021-05-10 11:16:37 |
| [gami](https://gitlab.com/bit4bit/gami) | 30 | 23 | 1 | GO - Asterisk AMI Interface | 2014-05-14 16:11:37 | 2021-04-21 04:24:16 |
| [gcm](https://gitlab.com/bit4bit/gami) | 29 | 3 | 0 | Google Cloud Messaging for application servers implemented using the Go programming language. | 2015-11-09 16:16:25 | 2019-07-14 22:15:57 |
| [patreon-go](https://github.com/mxpv/patreon-go) | 26 | 12 | 1 | Patreon Go API client | 2017-08-06 21:15:14 | 2021-05-22 11:02:31 |
| [airtable](https://github.com/mehanizm/airtable) | 26 | 6 | 0 | Simple golang airtable API wrapper | 2020-04-12 10:05:07 | 2021-06-09 06:01:09 |
| [go-postman-collection](https://github.com/rbretecher/go-postman-collection) | 25 | 8 | 9 | Go module to work with Postman Collections | 2019-11-16 12:13:32 | 2021-06-10 13:29:49 |
| [go-steam](https://pkg.go.dev/github.com/slack-go/slack) | 24 | 5 | 1 | Go library for querying Source servers | 2014-11-23 16:34:56 | 2021-05-10 11:16:39 |
| [go-myanimelist](https://github.com/nstratos/go-myanimelist) | 24 | 1 | 0 | Go library for accessing the MyAnimeList API: https://myanimelist.net/apiconfig/references/api/v2 | 2015-05-03 10:07:05 | 2021-07-02 03:13:16 |
| [lastpass-go](https://github.com/ansd/lastpass-go) | 22 | 2 | 0 | Golang client for LastPass | 2019-07-11 14:26:39 | 2021-06-16 07:46:17 |
| [go-imgur](https://github.com/koffeinsource/go-imgur) | 22 | 4 | 0 | Go library to use the imgur.com API | 2016-03-30 22:05:35 | 2021-05-21 04:32:36 |
| [go-shopify](https://github.com/rapito/go-shopify) | 21 | 6 | 2 | Simple Shopify API for the Go Programming Language | 2014-10-28 02:53:25 | 2021-04-17 01:17:44 |
| [go-twitch](https://github.com/knspriggs/go-twitch) | 21 | 2 | 5 | A golang client for the Twitch v3 API - public APIs only (for now) | 2016-06-28 20:54:34 | 2021-05-10 11:16:29 |
| [brewerydb](https://github.com/naegelejd/brewerydb) | 17 | 0 | 5 | Go library for http://www.brewerydb.com/ API | 2015-04-15 02:59:41 | 2020-02-21 23:20:13 |
| [textbelt](https://www.gregd.org/) | 17 | 0 | 0 | golang library for textbelt.com | 2015-09-01 22:46:42 | 2020-07-06 15:57:19 |
| [codeship-go](https://godoc.org/github.com/codeship/codeship-go) | 16 | 7 | 1 | Go library for accessing the Codeship API v2 | 2017-09-08 16:49:59 | 2020-11-03 16:20:17 |
| [go-hacknews](https://github.com/PaulRosset/go-hacknews) | 13 | 0 | 0 | 📟  Tiny utility Go client for HackerNews API. | 2017-08-10 20:44:02 | 2021-06-05 16:51:56 |
| [google-play-scraper](https://github.com/n0madic/google-play-scraper) | 13 | 6 | 0 | Golang scraper to get data from Google Play Store | 2019-09-20 14:03:01 | 2021-05-07 15:11:41 |
| [coinpaprika-api-go-client](https://github.com/coinpaprika/coinpaprika-api-go-client) | 12 | 4 | 1 | Go client library for interacting with Coinpaprika's API | 2018-09-25 07:34:50 | 2021-02-19 14:18:11 |
| [go-google-analytics](https://github.com/chonthu/go-google-analytics) | 12 | 2 | 0 | Simple Reporting for Google Analytics | 2015-06-01 13:50:00 | 2020-03-26 14:32:00 |
| [go-aws-news](https://caleblemoine.dev/go-aws-news/) | 12 | 4 | 0 | Go app + library to fetch what's new from AWS | 2020-01-08 00:59:39 | 2021-05-09 03:47:50 |
| [device-check-go](https://github.com/rinchsan/device-check-go) | 11 | 3 | 1 | :iphone: iOS DeviceCheck SDK for Go - query and modify the per-device bits | 2019-04-11 13:09:11 | 2021-06-02 08:59:24 |
| [smitego](https://pkg.go.dev/github.com/slack-go/slack) | 10 | 0 | 0 | SmiteGo is an API wrapper for the Smite game from HiRez. It is written in Go! | 2013-12-11 02:38:19 | 2019-04-19 22:37:32 |
| [go-here](https://github.com/abdullahselek/go-here) | 9 | 4 | 0 | Go client library around the HERE location based APIs. | 2019-07-07 12:14:34 | 2021-06-14 17:20:12 |
| [gopaapi5](https://www.utekar.com/amazon-product-advertising-api-5-go-client-library-gopaapi/) | 9 | 4 | 0 | Go Client Library for Amazon's Product Advertising API 5.0 | 2020-02-15 06:21:31 | 2021-05-10 19:24:55 |
| [rrdaclient](https://github.com/Omie/rrdaclient) | 8 | 0 | 0 | Go bindings for RRDA https://github.com/fcambus/rrda | 2014-09-15 21:06:16 | 2021-05-16 09:40:44 |
| [go-sophos](https://github.com/esurdam/go-sophos) | 8 | 3 | 0 | Sophos UTM 9 REST API Client in Golang | 2018-09-05 04:37:25 | 2021-03-21 16:31:40 |
| [gomalshare](https://github.com/MonaxGT/gomalshare) | 8 | 1 | 0 | Go library MalShare API | 2019-03-01 09:33:41 | 2020-08-19 19:08:55 |
| [go-google-email-audit-api](https://godoc.org/github.com/ngs/go-google-email-audit-api/emailaudit) | 7 | 4 | 0 | Go Client Library for G Suite Email Audit API | 2016-10-24 02:34:29 | 2020-05-18 05:50:35 |
| [go-sptrans](https://github.com/sergioaugrod/go-sptrans) | 6 | 0 | 0 | Go client library for the SPTrans Olho Vivo API. :bus: | 2017-09-11 01:21:28 | 2020-10-20 23:33:30 |
| [go-openproject](https://github.com/manuelbcd/go-openproject) | 6 | 2 | 5 | Go client library for OpenProject | 2021-02-13 23:23:13 | 2021-06-13 16:22:11 |
| [gumblr](https://github.com/mattcunningham/gumblr) | 6 | 5 | 0 | A Go Wrapper for the Tumblr v2 API | 2015-07-09 23:13:51 | 2018-07-29 14:51:36 |
| [go-zooz](https://github.com/gojuno/go-zooz) | 6 | 5 | 0 | Zooz API client for Go | 2017-07-04 09:28:23 | 2020-11-12 17:54:22 |
| [go-chronos](https://github.com/axelspringer/go-chronos) | 4 | 2 | 0 | :dancers: Go Chronos 3.x REST API Client | 2017-10-23 12:19:01 | 2021-01-13 06:15:38 |
| [kanka](https://kanka.io/en-US/docs/1.0) | 3 | 3 | 2 | Go client for the Kanka API | 2019-12-26 00:07:57 | 2021-04-22 02:42:17 |
| [libgoffi](https://github.com/clevabit/libgoffi) | 3 | 0 | 0 | libgoffi - libffi adapter library for Go | 2019-08-03 17:05:34 | 2020-11-26 01:07:32 |
| [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) | 2 | 0 | 0 | This is RAWG SDK GO. This library contains methods for interacting with RAWG API. | 2020-10-16 15:31:37 | 2021-01-08 13:13:18 |
| [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) | 2 | 0 | 0 | Golang SDK for AppStore Connect API (Unofficial) | 2020-06-11 10:05:56 | 2021-02-21 19:19:35 |
| [tripadvisor-golang](https://github.com/mrbenosborne/tripadvisor-golang) | 1 | 0 | 0 | A TripAdvisor API wrapper for Golang. | 2019-04-15 18:12:11 | 2019-08-12 08:47:19 |
| [vl-go](https://github.com/verifid/vl-go) | 1 | 0 | 0 | Go client library around the VerifID identity verification layer API. | 2019-02-09 12:46:53 | 2021-05-30 19:02:05 |
| [playlyfe-go-sdk](https://github.com/playlyfe/playlyfe-go-sdk) | 1 | 0 | 0 | This is the official Playlyfe Golang Sdk | 2015-05-25 09:34:47 | 2018-07-29 20:25:20 |
</details>

### UUID
Libraries for working with UUIDs.

<sup>*Last Update: 2021-07-15 09:25:01*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [uuid](https://github.com/google/uuid) | 2,791 | 259 | 9 | Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. | 2016-02-12 22:17:59 | 2021-07-15 00:46:44 |
| [ulid](https://github.com/oklog/ulid) | 2,305 | 88 | 1 | Universally Unique Lexicographically Sortable Identifier (ULID) in Go | 2016-12-06 15:26:52 | 2021-07-13 23:50:34 |
| [uuid](https://gof.rs/projects/uuid/) | 949 | 61 | 9 | A UUID package originally forked from github.com/satori/go.uuid | 2018-07-13 02:13:28 | 2021-07-13 14:24:52 |
| [wuid](https://github.com/edwingeng/wuid) | 418 | 36 | 0 | An extremely fast UUID alternative written in golang | 2018-01-27 01:16:28 | 2021-07-13 07:38:28 |
| [sno](https://pkg.go.dev/github.com/muyo/sno?tab=doc) | 49 | 2 | 1 | Compact, sortable and fast unique IDs with embedded metadata. | 2019-05-26 22:05:26 | 2021-07-09 09:42:01 |
| [nanoid](https://github.com/aidarkhanov/nanoid) | 32 | 2 | 0 | A tiny and fast Go unique string generator | 2019-07-02 12:15:56 | 2021-06-06 04:41:36 |
| [Goid](https://github.com/JakeHL/Goid) | 30 | 1 | 1 | A UUIDv4 generation package written in go | 2017-05-19 10:40:45 | 2021-04-02 13:06:38 |
| [uuid](https://github.com/agext/uuid) | 13 | 3 | 0 | Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. | 2016-02-03 03:02:51 | 2021-07-09 09:41:36 |
| [gouid](https://github.com/twharmon/gouid) | 8 | 0 | 0 | Fast, dependable universally unique ids | 2020-10-08 19:54:41 | 2021-06-04 08:39:15 |
| [GoFlake](https://github.com/Hart87/GoFlake) | 7 | 0 | 0 | A highly scalable and serverless unique ID generator for use in distributed systems. Written in GoLang. Inspired by Twitters Snowflake. | 2021-05-03 14:44:19 | 2021-06-30 15:29:33 |
</details>

### Utilities
General utilities and tools to make your life easier.

<sup>*Last Update: 2021-07-01 08:37:02*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [fzf](https://github.com/junegunn/fzf) | 37,541 | 1,623 | 246 | :cherry_blossom: A command-line fuzzy finder | 2013-10-23 16:04:23 | 2021-06-30 00:30:59 |
| [hub](https://hub.github.com/) | 21,020 | 2,211 | 270 | A command-line tool that makes git easier to use with GitHub. | 2009-12-05 22:15:25 | 2021-06-30 03:29:33 |
| [ctop](https://ctop.sh) | 11,725 | 465 | 62 | Top-like interface for container metrics | 2016-12-27 02:25:57 | 2021-06-29 19:49:11 |
| [sqlx](http://jmoiron.github.io/sqlx/) | 10,405 | 826 | 259 | general purpose extensions to golang's database/sql | 2013-01-28 19:40:00 | 2021-06-30 15:30:32 |
| [wuzz](https://github.com/asciimoo/wuzz) | 9,679 | 373 | 36 | Interactive cli tool for HTTP inspection | 2017-01-30 21:22:00 | 2021-07-01 01:19:45 |
| [goreleaser](https://goreleaser.com) | 8,274 | 553 | 21 | Deliver Go binaries as fast and easily as possible | 2016-12-21 17:13:39 | 2021-06-29 15:58:19 |
| [usql](https://github.com/xo/usql) | 6,565 | 227 | 50 | Universal command-line interface for SQL databases | 2017-03-02 13:03:21 | 2021-06-29 09:20:08 |
| [peco](https://github.com/peco/peco) | 6,436 | 213 | 39 | Simplistic interactive filtering tool | 2014-06-06 06:06:32 | 2021-06-29 15:51:01 |
| [godropbox](https://github.com/dropbox/godropbox) | 3,945 | 416 | 5 | Common libraries for writing Go services/applications. | 2014-06-22 23:09:29 | 2021-06-29 19:13:12 |
| [hystrix-go](https://github.com/afex/hystrix-go) | 3,210 | 369 | 50 | Netflix's Hystrix latency and fault tolerance library, for Go  | 2013-12-15 08:51:23 | 2021-06-30 08:13:44 |
| [goreporter](https://github.com/360EntSecGroup-Skylar/goreporter) | 2,878 | 255 | 28 | A Golang tool that does static analysis, unit testing, code review and generate code quality report. | 2017-03-27 08:46:38 | 2021-06-29 06:52:15 |
| [go-funk](https://github.com/thoas/go-funk) | 2,703 | 155 | 5 | A modern Go utility library which provides helpers (map, find, contains, filter, ...) | 2016-12-30 13:55:15 | 2021-06-29 20:52:59 |
| [minify](https://go.tacodewolff.nl/minify) | 2,697 | 166 | 12 | Go minifiers for web formats | 2014-05-21 09:03:48 | 2021-06-30 00:50:50 |
| [panicparse](https://maruel.ca) | 2,596 | 80 | 3 | Crash your app in style (Golang) | 2015-02-02 02:14:41 | 2021-06-29 15:13:28 |
| [mc](https://min.io/download) | 1,833 | 323 | 36 | MinIO Client is a replacement for ls, cp, mkdir, diff and rsync commands for filesystems and object storage. | 2015-01-16 02:56:51 | 2021-06-29 15:48:17 |
| [storm](https://github.com/asdine/storm) | 1,734 | 122 | 65 | Simple and powerful toolkit for BoltDB | 2016-01-10 12:55:59 | 2021-06-23 08:00:29 |
| [mergo](https://github.com/imdario/mergo) | 1,581 | 195 | 18 | Mergo: merging Go structs and maps since 2013. | 2013-03-11 22:51:11 | 2021-06-30 00:35:49 |
| [spinner](https://github.com/briandowns/spinner) | 1,533 | 99 | 12 | Go (golang) package with 80 configurable terminal spinner/progress indicators. | 2014-12-13 00:36:19 | 2021-06-26 14:56:24 |
| [mole](https://davrodpin.github.io/mole/) | 1,475 | 81 | 14 | CLI application to create ssh tunnels focused on resiliency and user experience. | 2018-10-04 02:38:00 | 2021-06-20 21:58:45 |
| [boilr](https://github.com/tmrts/boilr) | 1,331 | 100 | 43 | :zap: boilerplate template manager that generates files or directories from template repositories | 2015-12-19 16:57:26 | 2021-06-24 05:07:27 |
| [filetype](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 1,321 | 120 | 26 | Fast, dependency-free Go package to infer binary file types based on the magic numbers header signature | 2015-09-24 09:15:51 | 2021-06-28 15:59:04 |
| [jump](http://gsamokovarov.com/jump) | 1,093 | 45 | 1 | Jump helps you navigate faster by learning your habits. ✌️ | 2015-08-16 22:07:17 | 2021-06-29 14:16:37 |
| [circuitbreaker](https://github.com/rubyist/circuitbreaker) | 945 | 103 | 18 | Circuit Breakers in Go | 2014-07-17 22:41:33 | 2021-06-25 15:45:04 |
| [gtm](https://github.com/git-time-metric/gtm) | 868 | 48 | 46 | Simple, seamless, lightweight time tracking for Git | 2016-06-19 21:17:04 | 2021-06-27 07:50:11 |
| [immortal](https://immortal.run) | 708 | 47 | 1 | ⭕  A *nix cross-platform (OS agnostic) supervisor | 2016-06-30 17:02:27 | 2021-06-18 07:24:09 |
| [hostctl](http://guumaster.github.io/hostctl) | 661 | 30 | 11 | Your dev tool to manage /etc/hosts like a pro! | 2020-03-14 11:29:02 | 2021-06-20 09:19:34 |
| [circuit](https://github.com/cep21/circuit) | 573 | 33 | 7 | An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. | 2017-12-23 22:17:43 | 2021-06-25 15:44:57 |
| [htcat](https://github.com/htcat/htcat) | 533 | 29 | 5 | Parallel and Pipelined HTTP GET Utility | 2013-08-05 11:17:01 | 2021-06-25 05:28:10 |
| [mimetype](https://pkg.go.dev/github.com/gabriel-vasile/mimetype#pkg-overview) | 521 | 78 | 32 | A fast Golang library for media type and file extension detection, based on magic numbers | 2018-07-02 07:15:29 | 2021-06-30 03:19:44 |
| [godaemon](https://github.com/VividCortex/godaemon) | 477 | 53 | 8 | Daemonize Go applications deviously. | 2013-08-01 17:16:30 | 2021-06-17 07:42:30 |
| [cli](https://create-go.app/wiki) | 476 | 52 | 0 | ✨ Create a new production-ready project with backend, frontend and deploy automation by running one CLI command! | 2019-12-30 22:08:38 | 2021-06-27 09:59:26 |
| [ergo](https://github.com/cristianoliveira/ergo) | 468 | 44 | 18 | The management of multiple apps running over different ports made easy | 2017-08-19 18:41:56 | 2021-06-29 17:13:08 |
| [go-dry](https://github.com/ungerik/go-dry) | 465 | 32 | 0 | DRY (don't repeat yourself) package for Go | 2014-02-28 13:49:31 | 2021-06-13 15:42:48 |
| [koazee](https://github.com/wesovilabs/koazee) | 463 | 24 | 14 | A StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices. | 2018-11-09 09:49:19 | 2021-06-16 06:53:00 |
| [gopencils](https://github.com/bndr/gopencils) | 437 | 39 | 7 | Easily consume REST APIs with Go (golang) | 2014-06-23 11:41:24 | 2021-06-21 14:34:45 |
| [request](https://godoc.org/github.com/mozillazg/request) | 397 | 40 | 6 | A developer-friendly HTTP request library for Gopher. | 2014-12-21 04:30:42 | 2021-06-22 03:49:07 |
| [deepcopier](https://github.com/ulule/deepcopier) | 352 | 49 | 6 | simple struct copying for golang | 2015-07-24 18:01:01 | 2021-06-28 11:47:18 |
| [gubrak](https://pkg.go.dev/github.com/novalagung/gubrak) | 349 | 28 | 0 | ⚙️ Golang functional utility library with syntactic sugar. It's like lodash, but for Go | 2018-03-09 11:28:05 | 2021-06-30 03:59:58 |
| [clockwork](https://github.com/jonboulle/clockwork) | 338 | 42 | 5 | a fake clock for golang | 2014-09-09 18:24:00 | 2021-06-17 08:29:49 |
| [go-rate](https://github.com/beefsack/go-rate) | 336 | 26 | 0 | A timed rate limiter for Go | 2014-08-25 04:42:34 | 2021-06-23 15:28:20 |
| [delve](https://github.com/derekparker/delve) | 329 | 74 | 0 | Delve is a debugger for the Go programming language. | 2020-02-18 18:03:33 | 2021-06-29 07:33:54 |
| [retry](https://pkg.go.dev/github.com/kamilsk/retry/v5) | 304 | 13 | 9 | ♻️ The most advanced interruptible mechanism to perform actions repetitively until successful. | 2016-11-02 20:20:43 | 2021-06-18 19:09:14 |
| [scany](https://github.com/georgysavva/scany) | 300 | 23 | 12 | Library for scanning data from a database into Go structs and more | 2020-07-02 11:02:58 | 2021-06-30 04:57:14 |
| [gohper](https://github.com/cosiner/gohper) | 253 | 46 | 0 | [UNMATAINED] common libs here. | 2015-03-23 22:46:12 | 2021-04-05 15:04:34 |
| [serve](https://syntaqx-serve.herokuapp.com/) | 245 | 15 | 4 | 🍽️ a static http server anywhere you need one. | 2019-01-10 23:31:52 | 2021-05-21 04:32:45 |
| [go-trigger](https://github.com/sadlil/go-trigger) | 217 | 37 | 1 | A Global event triggerer for golang. Defines functions as event with id string. Trigger the event anywhere from your project. | 2015-10-19 09:42:17 | 2021-06-09 04:33:45 |
| [util](https://github.com/shomali11/util) | 212 | 34 | 0 | A collection of useful utility functions | 2017-05-24 00:21:29 | 2021-06-25 03:13:25 |
| [scan](https://github.com/blockloop/scan) | 204 | 11 | 0 | Scan database/sql rows directly to structs, slices, and primitive types | 2017-11-27 23:22:18 | 2021-06-28 16:16:17 |
| [gotenv](https://github.com/subosito/gotenv) | 198 | 20 | 3 | Load environment variables from `.env` or `io.Reader` in Go. | 2013-08-27 12:56:47 | 2021-06-27 07:08:58 |
| [death](http://vrecan.github.io/post/golang_shutdown/) | 172 | 17 | 0 | Managing go application shutdown with signals. | 2015-03-09 03:50:40 | 2021-05-21 04:32:40 |
| [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) | 169 | 7 | 0 | go-bind-plugin generates API for exported plugin symbols (-buildmode=plugin) - go1.8+ only (http://golang.org/pkg/plugin) | 2016-11-08 14:40:26 | 2021-06-12 14:30:14 |
| [rerun](https://github.com/ivpusic/rerun) | 162 | 9 | 2 | Configurable recompiling and rerunning go apps when source changes | 2014-12-10 00:29:54 | 2021-05-21 19:48:51 |
| [moldova](https://github.com/StabbyCutyou/moldova) | 161 | 5 | 0 | A lightweight templating system for generating random data | 2016-01-30 05:25:39 | 2021-06-28 17:28:48 |
| [toolbox](https://github.com/viant/toolbox) | 161 | 19 | 2 | Toolbox - go utility library | 2016-06-13 19:33:35 | 2021-05-31 04:35:43 |
| [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) | 158 | 46 | 21 | go-sitemap-generator is the easiest way to generate Sitemaps in Go | 2015-10-12 16:23:13 | 2021-06-21 01:19:26 |
| [robustly](https://github.com/VividCortex/robustly) | 151 | 6 | 1 | Run functions resiliently in Go, catching and restarting panics | 2013-07-08 13:27:10 | 2021-06-09 06:27:19 |
| [apm](https://github.com/topfreegames/apm) | 149 | 72 | 9 | APM is a process manager for Golang applications. | 2015-11-18 16:56:48 | 2021-05-02 16:12:18 |
| [chyle](https://github.com/antham/chyle) | 139 | 8 | 0 | Changelog generator : use a git repository and various data sources and publish the result on external services | 2016-11-17 21:14:44 | 2021-06-16 06:38:01 |
| [onecache](https://github.com/adelowo/onecache) | 120 | 6 | 0 | One caching API, Multiple backends | 2017-04-14 21:49:15 | 2021-06-12 21:13:18 |
| [lrserver](https://github.com/jaschaephraim/lrserver) | 119 | 11 | 0 | LiveReload server for Go [golang] | 2014-07-15 05:36:53 | 2021-06-26 23:15:07 |
| [go-bsdiff](https://github.com/gabstv/go-bsdiff) | 117 | 12 | 0 | Pure Go bsdiff and bspatch libraries and CLI tools. | 2019-02-23 23:33:50 | 2021-06-19 11:26:43 |
| [countries](https://godoc.org/github.com/biter777/countries) | 97 | 23 | 3 | Countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO init() funcs, NO external links/files/data, NO interface{}, NO specific dependencies, Databases/JSON/GOB/XML/CSV compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts. | 2019-04-22 14:47:11 | 2021-06-29 08:02:09 |
| [nostromo](https://nostromo.sh) | 89 | 4 | 11 | CLI for building powerful aliases | 2019-07-13 04:51:46 | 2021-06-02 03:19:17 |
| [xferspdy](https://github.com/monmohan/xferspdy) | 87 | 9 | 3 | Xferspdy provides binary diff and patch library in golang. [Mentioned in Awesome Go, https://github.com/avelino/awesome-go] | 2015-05-22 13:23:34 | 2021-06-25 01:43:37 |
| [mssqlx](https://github.com/linxGnu/mssqlx) | 86 | 11 | 0 | Database client library, proxy for any master slave, master master structures. Lightweight, performant and auto balancing in mind. | 2016-12-26 04:05:09 | 2021-06-09 06:01:07 |
| [goseaweedfs](https://github.com/chrislusf/seaweedfs) | 85 | 25 | 1 | A complete Golang client for SeaweedFS | 2017-07-20 04:35:39 | 2021-06-17 08:28:32 |
| [go-pattern-match](https://github.com/alexpantyukhin/go-pattern-match) | 83 | 6 | 0 | Pattern matchings for Go. | 2018-12-11 20:11:17 | 2021-06-16 03:54:56 |
| [go-health](https://github.com/Talento90/go-health) | 80 | 4 | 1 | :heart: Health check your applications and dependencies | 2018-02-13 18:40:54 | 2021-06-17 20:37:31 |
| [sorty](https://github.com/jfcg/sorty) | 80 | 1 | 0 | Fast Concurrent / Parallel Sorting in Go | 2019-02-18 21:05:45 | 2021-06-22 23:23:52 |
| [pm](https://github.com/VividCortex/pm) | 78 | 6 | 2 | Processlist manager with TCP listener | 2013-11-17 19:17:01 | 2021-04-05 15:04:27 |
| [repeat](https://github.com/ssgreg/repeat) | 76 | 5 | 0 | Go implementation of different backoff strategies useful for retrying operations and heartbeating. | 2017-11-22 07:06:47 | 2021-04-18 15:14:24 |
| [netbug](https://github.com/e-dard/netbug) | 69 | 4 | 0 | Package netbug provides a handler for registering profilers on your own ServeMux. | 2015-03-05 19:27:29 | 2021-05-10 07:30:05 |
| [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination/wiki) | 67 | 24 | 1 | Golang Mongodb Pagination for official mongodb/mongo-go-driver package which supports  both normal queries and Aggregation pipelines with all information like Total records, Page, Per Page, Previous , Next, Total Page and query results. | 2020-02-04 08:23:33 | 2021-06-29 09:05:39 |
| [unis](https://godoc.org/github.com/esemplastic/unis) | 67 | 3 | 2 | UNIS: A Common Architecture for String Utilities within the Go Programming Language. | 2017-05-06 05:01:03 | 2021-03-10 09:39:40 |
| [multitick](https://github.com/VividCortex/multitick) | 66 | 1 | 1 | A multiplexor for aligned time.Time tickers in Go | 2013-12-10 16:47:26 | 2021-06-21 00:38:25 |
| [handy](https://pkg.go.dev/github.com/novalagung/gubrak) | 64 | 5 | 0 | GO Golang Utilities and helpers like validators and string formatters | 2018-06-13 13:10:07 | 2021-04-15 12:14:13 |
| [mimemagic](https://github.com/zRedShift/mimemagic) | 62 | 5 | 2 | Powerful and versatile MIME sniffing package using pre-compiled glob patterns, magic number signatures, XML document namespaces, and tree magic for mounted volumes, generated from the XDG shared-mime-info database. | 2018-10-11 16:12:54 | 2021-06-15 09:34:33 |
| [cmd](https://github.com/commander-cli/cmd) | 61 | 4 | 2 | A simple package to execute shell commands on linux, windows and osx | 2019-09-27 13:22:06 | 2021-06-28 08:49:55 |
| [changie](https://changie.dev) | 58 | 6 | 8 | Automated changelog tool for preparing releases with lots of customization options | 2020-12-05 19:38:33 | 2021-06-24 09:25:39 |
| [minquery](https://github.com/icza/minquery) | 58 | 19 | 4 | MongoDB / mgo query that supports efficient pagination (cursors to continue listing documents where we left off). | 2016-11-16 12:23:07 | 2021-04-02 11:42:09 |
| [go-astitodo](https://github.com/asticode/go-astitodo) | 58 | 8 | 2 | Parse TODOs in your GO code | 2016-10-17 19:51:36 | 2021-04-25 08:26:07 |
| [goreadability](https://github.com/philipjkim/goreadability) | 58 | 7 | 2 | Webpage summary extractor using Facebook Open Graph and arc90's readability | 2016-04-20 01:40:14 | 2021-04-16 20:59:37 |
| [golog](https://github.com/mlimaloureiro/golog) | 55 | 12 | 15 | Easy and simple CLI time tracker for your tasks | 2016-01-09 15:43:47 | 2021-06-13 16:19:23 |
| [pgo](https://github.com/arthurkushman/pgo) | 53 | 11 | 2 | Go library for PHP community with convenient functions | 2018-12-26 06:59:47 | 2021-06-09 06:38:15 |
| [goval](https://github.com/maja42/goval) | 49 | 5 | 0 | Expression evaluation in golang | 2018-06-17 15:43:44 | 2021-06-28 13:45:05 |
| [copy-pasta](https://github.com/jutkko/copy-pasta) | 49 | 9 | 10 | Universal copy paste service, works across different machines! | 2017-01-28 15:35:24 | 2021-04-24 11:00:07 |
| [retry](https://pkg.go.dev/github.com/thedevsaddam/retry) | 48 | 4 | 0 | Simple and easy retry mechanism package for Go | 2018-02-25 19:08:03 | 2021-06-18 05:42:13 |
| [filter](https://godoc.org/github.com/gookit/filter) | 47 | 4 | 1 | ⏳ Provide filtering, sanitizing, and conversion of Golang data. 提供对Golang数据的过滤，净化，转换。 | 2018-09-26 09:11:13 | 2021-06-07 11:14:47 |
| [beyond](http://wesovilabs.github.io/beyond) | 46 | 9 | 7 | The Go library that will drive you to AOP world! | 2019-10-18 05:41:45 | 2021-05-21 04:32:39 |
| [golarm](https://github.com/msempere/golarm) | 45 | 8 | 0 | Fire alarms with system events | 2015-08-14 16:51:53 | 2021-06-13 19:53:42 |
| [goback](https://github.com/carlescere/goback) | 44 | 7 | 6 | Golang simple exponential backoff package. | 2015-03-13 16:09:18 | 2021-03-04 10:31:30 |
| [go-lock](https://github.com/viney-shih/go-lock) | 42 | 2 | 0 | go-lock is a lock library implementing read-write mutex and read-write trylock without starvation | 2020-04-30 11:40:21 | 2021-05-06 14:39:12 |
| [slice](https://github.com/psampaz/slice) | 42 | 4 | 0 | Type-safe functions for common Go slice operations | 2019-11-26 05:20:39 | 2021-05-27 17:09:50 |
| [intrinsic](https://immortal.run) | 42 | 1 | 1 | Provide Golang native SIMD intrinsics on x86/amd64 platform | 2017-06-13 09:26:34 | 2020-10-15 07:47:46 |
| [dbt](https://github.com/nikogura/dbt) | 41 | 4 | 6 | Dynamic Binary Toolkit- A framework for running self-updating signed binaries from a central, trusted repository. | 2017-11-30 22:53:17 | 2021-03-23 14:38:08 |
| [gpath](https://github.com/tenntenn/gpath) | 41 | 2 | 0 | gpath is a Go package to access a field by a path using reflect pacakge | 2017-05-24 06:24:18 | 2021-05-14 06:30:49 |
| [retry-go](https://github.com/rafaeljesus/retry-go) | 40 | 2 | 2 | Retrying made simple and easy for golang :repeat:  | 2017-06-09 16:07:37 | 2021-02-11 06:35:06 |
| [go-httpheader](https://godoc.org/github.com/mozillazg/go-httpheader) | 34 | 7 | 1 | A Go library for encoding structs into Header fields. | 2017-06-24 11:28:06 | 2021-02-13 08:10:21 |
| [myhttp](https://github.com/inancgumus/myhttp) | 33 | 13 | 1 | Simplest HTTP GET requester for Go with timeout support | 2017-09-13 15:48:47 | 2019-11-21 09:45:21 |
| [gostrutils](https://github.com/ik5/gostrutils) | 32 | 5 | 1 | Collections of string utils I have created over the years | 2018-09-19 11:06:11 | 2021-06-18 05:28:34 |
| [rclient](https://github.com/zpatrick/rclient) | 32 | 2 | 2 | Minimalistic REST client for Go applications | 2017-02-28 01:07:25 | 2020-10-13 11:13:06 |
| [evaluator](https://github.com/nullne/evaluator) | 31 | 6 | 1 | The management of multiple apps running over different ports made easy | 2017-04-27 18:31:46 | 2021-03-30 21:02:32 |
| [equalizer](https://pkg.go.dev/github.com/reugn/equalizer) | 31 | 0 | 0 | A rate limiters package for Go | 2019-06-14 09:25:13 | 2021-06-13 18:46:40 |
| [limiters](https://godoc.org/github.com/mennanov/limiters) | 29 | 3 | 1 | Golang rate limiters for distributed applications | 2019-08-28 18:09:54 | 2021-06-17 08:54:42 |
| [tome](https://github.com/cyruzin/tome) | 28 | 2 | 1 | Package tome was designed to paginate simple RESTful APIs. | 2019-04-12 16:49:45 | 2021-06-09 06:37:39 |
| [ugo](https://github.com/alxrm/ugo) | 25 | 4 | 0 | Simple and expressive toolbox written in Go | 2016-02-17 19:41:57 | 2021-05-12 09:46:11 |
| [generate](https://github.com/go-playground/generate) | 24 | 4 | 0 | :runner:runs go generate recursively on a specified path or environment variable and can filter by regex | 2015-11-15 01:52:04 | 2021-03-13 21:24:07 |
| [slicer](https://github.com/leaanthony/slicer) | 23 | 1 | 0 | Utility class for handling slices | 2019-01-10 09:55:25 | 2021-02-10 17:08:04 |
| [goplaceholder](https://github.com/michiwend/goplaceholder) | 22 | 6 | 1 | a small golang lib to generate placeholder images | 2014-10-12 00:50:46 | 2020-09-16 12:56:17 |
| [rerate](https://github.com/abo/rerate) | 20 | 4 | 1 | redis-based rate counter and rate limiter | 2016-05-24 08:59:00 | 2021-06-09 05:12:08 |
| [shutdown](https://github.com/ztrue/shutdown) | 19 | 3 | 0 | Golang app shutdown hooks. | 2018-11-17 17:56:03 | 2021-06-16 06:07:01 |
| [ctxutil](https://ctop.sh) | 18 | 2 | 1 | utils for Go context | 2018-07-30 11:28:57 | 2021-06-19 06:16:11 |
| [ghokin](https://github.com/antham/ghokin) | 18 | 0 | 1 | Parallelized formatter with no external dependencies for gherkin (cucumber, behat...) | 2018-08-03 11:36:35 | 2021-06-28 07:02:00 |
| [structs](https://github.com/PumpkinSeed/structs) | 17 | 2 | 0 | Golang struct operations. | 2017-08-26 09:59:00 | 2021-04-08 10:47:04 |
| [backscanner](https://github.com/icza/backscanner) | 17 | 5 | 0 | A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. | 2017-10-18 07:59:07 | 2021-06-09 04:44:53 |
| [filler](https://pkg.go.dev/github.com/h2non/filetype?tab=doc) | 16 | 3 | 0 | fill struct data easily with fill tags | 2017-04-05 08:14:04 | 2020-12-25 16:48:53 |
| [dlog](https://github.com/kirillDanshin/dlog) | 16 | 1 | 0 | Simple build-time controlled debug log with ability to log where the logger was called | 2016-07-04 19:59:09 | 2020-10-19 07:34:21 |
| [mimesniffer](https://pkg.go.dev/github.com/aofei/mimesniffer) | 15 | 0 | 0 | A MIME type sniffer for Go. | 2018-12-20 03:40:20 | 2021-02-18 10:34:43 |
| [okrun](https://github.com/xta/okrun) | 15 | 2 | 0 | ok, run your gofile | 2014-10-01 06:18:56 | 2020-07-14 23:33:14 |
| [jsend](https://clevergo.tech) | 14 | 4 | 0 | :100: JSend's implementation writen in Go(golang) | 2020-01-14 04:41:36 | 2021-06-29 03:45:23 |
| [command](https://github.com/txgruppi/command) | 14 | 3 | 0 | Command pattern for Go with thread safe serial and parallel dispatcher | 2015-08-24 09:43:50 | 2021-06-19 06:15:44 |
| [go-convert](https://github.com/Eun/go-convert) | 13 | 1 | 3 | Convert a value into another type | 2019-06-07 16:56:38 | 2021-06-19 06:18:18 |
| [rest-go](https://github.com/edermanoel94/rest-go) | 13 | 1 | 1 | A package that provide many helpful methods for working with rest api. | 2019-07-29 18:56:08 | 2021-03-01 04:09:50 |
| [retry](https://github.com/shafreeck/retry) | 11 | 1 | 1 | A pretty simple library to ensure your work to be done | 2018-07-18 09:48:33 | 2020-04-16 04:37:48 |
| [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) | 10 | 0 | 0 | Problem json implementation (https://tools.ietf.org/html/rfc7807) package for go | 2019-05-16 05:42:14 | 2021-05-22 13:33:58 |
| [go-types](https://github.com/mikekonan/go-types) | 10 | 4 | 1 | Library providing opanapi3 and Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. | 2021-04-21 11:34:25 | 2021-06-16 13:07:08 |
| [ptr](https://github.com/gotidy/ptr) | 10 | 2 | 0 | Contains functions for simplified creation of pointers from constants of basic types | 2019-12-25 15:29:48 | 2021-06-28 17:16:27 |
| [go-countries](https://github.com/mikekonan/go-countries) | 9 | 3 | 0 | Convert a value into another type | 2020-10-27 12:56:40 | 2021-02-03 14:41:07 |
| [silk](https://github.com/chrispassas/silk) | 9 | 1 | 0 | Read Silk Flow Files | 2018-12-18 04:23:35 | 2021-04-02 08:08:30 |
| [nfdump](https://github.com/chrispassas/nfdump) | 7 | 0 | 0 | NFDump File Reader | 2020-04-08 01:01:22 | 2021-03-23 16:08:50 |
| [copy](https://github.com/gotidy/copy) | 7 | 1 | 2 | Package for fast copying structs of different types | 2020-10-09 06:59:08 | 2021-06-05 16:51:05 |
| [statiks](https://github.com/janiltonmaciel/statiks) | 7 | 0 | 0 | Fast, zero-configuration, static HTTP filer server. | 2018-06-26 23:42:33 | 2020-11-09 16:55:26 |
| [sliceconv](https://github.com/Henry-Sarabia/sliceconv) | 7 | 0 | 0 | Slice conversion between primitive types | 2019-02-15 06:50:34 | 2021-05-01 15:59:31 |
| [go-clip](https://github.com/prashantgupta24/go-clip) | 7 | 0 | 2 | A minimalistic clipboard manager for Mac. | 2020-11-18 22:19:01 | 2021-02-10 05:41:51 |
| [retry](https://github.com/percolate/retry) | 7 | 1 | 0 | Percolate's Go retry package | 2018-06-15 19:23:36 | 2020-08-31 05:58:43 |
| [blank](https://github.com/Henry-Sarabia/blank) | 6 | 0 | 0 | Detect blank strings or remove whitespace from strings | 2019-02-13 00:07:27 | 2021-02-02 01:05:34 |
| [lets-go](https://github.com/aplescia/lets-go) | 4 | 0 | 0 | Go module that provides common utilities for Cloud Native development | 2020-02-19 16:32:41 | 2021-05-21 04:32:43 |
| [go-safe](https://github.com/kenkyu392/go-safe) | 4 | 0 | 0 | This Go package provides a sandbox for the safe execution of panic-inducing programs | 2019-10-29 15:20:37 | 2021-05-21 04:32:42 |
| [goctx](https://github.com/zerosnake0/goctx) | 2 | 1 | 0 | Get your context value faster | 2020-11-14 14:16:09 | 2021-01-24 11:56:45 |
| [olaf](https://github.com/btnguyen2k/olaf) | 2 | 0 | 0 | Twitter Snowflake implemented in Go | 2019-01-03 13:31:10 | 2020-10-31 07:30:29 |
| [tik](https://github.com/andy2046/tik) | 2 | 0 | 0 | hierarchical timing wheel | 2020-07-04 09:13:49 | 2021-01-09 09:42:40 |
| [bleep](https://github.com/sinhashubham95/bleep) | 1 | 0 | 0 | OS Signal Handlers in Go | 2021-01-02 05:22:08 | 2021-01-06 03:41:44 |
</details>

### Validation
Libraries for validation.

<sup>*Last Update: 2021-07-12 14:22:32*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [validator](https://github.com/go-playground/validator) | 8,166 | 729 | 133 | :100:Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving | 2015-02-12 16:32:22 | 2021-07-12 03:32:33 |
| [govalidator](https://github.com/asaskevich/govalidator) | 4,831 | 488 | 132 | [Go] Package of validators and sanitizers for strings, numerics, slices and structs | 2014-06-20 10:45:23 | 2021-07-11 15:58:52 |
| [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) | 2,213 | 141 | 21 | An idiomatic Go (golang) validation package. Supports configurable and extensible validation rules (validators) using normal language constructs instead of error-prone struct tags. | 2016-06-22 03:47:43 | 2021-07-11 13:49:23 |
| [govalidator](https://github.com/thedevsaddam/govalidator) | 999 | 88 | 29 | Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. | 2017-09-13 16:42:20 | 2021-07-12 03:14:37 |
| [validate](https://pkg.go.dev/github.com/gookit/validate) | 443 | 65 | 8 | ⚔ Go package for data validation and filtering. support Map, Struct, Form data. Go通用的数据验证与过滤库，使用简单，内置大部分常用验证、过滤器，支持自定义验证器、自定义消息、字段翻译。 | 2018-07-16 08:23:49 | 2021-07-12 05:09:31 |
| [checkdigit](https://github.com/osamingo/checkdigit) | 78 | 2 | 0 | Provide check digit algorithms and calculators written in Go | 2019-04-05 09:46:36 | 2021-06-23 13:49:51 |
| [validate](https://pkg.go.dev/github.com/gookit/validate) | 59 | 18 | 5 | This package provides a framework for writing validations for Go applications. | 2018-02-10 18:25:55 | 2021-04-01 19:07:51 |
| [jio](https://github.com/faceair/jio) | 58 | 9 | 0 | jio is a json schema validator similar to joi | 2018-10-28 11:02:45 | 2021-05-26 03:38:40 |
| [terraform-validator](https://github.com/thazelart/terraform-validator) | 58 | 4 | 6 | A norms and conventions validator for Terraform | 2019-05-29 11:37:15 | 2021-07-04 00:28:42 |
| [gody](https://pkg.go.dev/github.com/guiferpa/gody) | 49 | 3 | 1 | :balloon: A lightweight struct validator for Go | 2018-11-01 21:08:16 | 2021-02-19 02:07:17 |
| [govalid](https://github.com/twharmon/govalid) | 22 | 3 | 0 | Struct validation using tags | 2019-02-17 23:25:43 | 2021-06-22 18:28:35 |
</details>

### Version Control
Libraries for version control.

<sup>*Last Update: 2021-07-12 14:22:30*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-git](https://pkg.go.dev/github.com/go-git/go-git/v5) | 2,458 | 257 | 195 | A highly extensible Git implementation in pure Go. | 2019-12-19 10:27:02 | 2021-07-11 21:28:53 |
| [git2go](https://github.com/libgit2/git2go) | 1,633 | 285 | 42 | Git to Go; bindings for libgit2. Like McDonald's but tastier. | 2013-03-05 19:50:43 | 2021-07-08 19:08:10 |
| [hercules](https://github.com/src-d/hercules) | 1,381 | 118 | 38 | Gaining advanced insights from Git repository history. | 2016-12-12 17:30:29 | 2021-07-11 08:36:23 |
| [gh](https://github.com/rjeczalik/gh) | 75 | 12 | 2 | Scriptable server and net/http middleware for GitHub Webhooks. | 2015-03-08 21:04:05 | 2020-05-08 16:42:07 |
| [go-vcs](https://sourcegraph.com/sourcegraph/go-vcs) | 75 | 21 | 23 | manipulate and inspect VCS repositories in Go | 2013-06-02 02:36:18 | 2021-06-09 06:01:11 |
| [hgo](https://github.com/beyang/hgo) | 13 | 3 | 0 | Hgo is a collection of Go packages providing read-access to local Mercurial repositories. | 2014-06-18 03:54:40 | 2020-05-05 03:52:16 |
</details>

### Video
Libraries for manipulating video.

<sup>*Last Update: 2021-07-12 14:22:24*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [goav](https://github.com/giorgisio/goav) | 1,648 | 296 | 44 | Golang bindings for FFmpeg | 2015-05-21 05:31:14 | 2021-07-11 13:48:51 |
| [m3u8](http://tools.ietf.org/html/draft-pantos-http-live-streaming) | 840 | 219 | 43 | Parser and generator of M3U8-playlists for Apple HLS. Library for Go language. :cinema: | 2013-02-05 22:26:30 | 2021-07-09 22:59:20 |
| [gmf](https://github.com/3d0c/gmf) | 693 | 135 | 41 | Go Media Framework | 2013-04-03 09:07:47 | 2021-07-11 13:47:30 |
| [go-astits](https://github.com/asticode/go-astits) | 369 | 33 | 6 | Demux and mux MPEG Transport Streams (.ts) natively in GO | 2017-07-04 13:06:15 | 2021-07-05 09:10:08 |
| [go-astisub](https://github.com/asticode/go-astisub) | 320 | 63 | 5 | Manipulate subtitles in GO (.srt, .ssa/.ass, .stl, .ttml, .vtt (webvtt), teletext, etc.) | 2016-12-16 14:47:59 | 2021-07-12 01:03:45 |
| [libvlc-go](https://pkg.go.dev/github.com/adrg/libvlc-go/v3) | 222 | 31 | 2 | Go bindings for libVLC and high-level media player interface | 2015-01-06 14:01:50 | 2021-07-09 23:07:14 |
| [gst](https://github.com/ziutek/gst) | 159 | 45 | 9 | Go bindings for GStreamer (retired: currently I don't use/develop this package) | 2011-07-26 00:44:40 | 2021-04-21 01:25:11 |
| [go-m3u8](https://tools.ietf.org/html/rfc8216) | 75 | 11 | 0 | Parse and generate m3u8 playlists for Apple HTTP Live Streaming (HLS) in Golang (ported from gem https://github.com/sethdeckard/m3u8) | 2018-11-06 02:42:27 | 2021-06-02 05:43:35 |
| [v4l](https://pkg.go.dev/github.com/korandiz/v4l) | 58 | 10 | 0 | Facade to the Video4Linux video capture interface.  | 2016-10-25 10:50:25 | 2021-05-16 09:56:04 |
| [libgosubs](https://github.com/wargarblgarbl/libgosubs) | 14 | 3 | 0 | golang library to read and write various subtitle formats | 2017-05-03 21:05:25 | 2021-05-20 14:04:11 |
| [go-mpd](https://github.com/unki2aut/go-mpd) | 9 | 2 | 0 | Go library for parsing and generating MPEG-DASH Media Presentation Description (MPD) files | 2018-11-02 19:09:07 | 2021-06-13 08:14:53 |
</details>

### Web Frameworks
Full stack web frameworks.

<sup>*Last Update: 2021-07-13 09:25:04*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gin](https://gin-gonic.com/) | 49,468 | 5,611 | 407 | Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin. | 2014-06-16 23:57:25 | 2021-07-13 02:07:58 |
| [echo](https://echo.labstack.com) | 20,226 | 1,786 | 57 | High performance, minimalist Go web framework | 2015-03-01 17:43:01 | 2021-07-13 01:09:01 |
| [fiber](https://gofiber.io) | 14,151 | 677 | 27 | ⚡️ Express inspired web framework written in Go | 2020-01-16 03:59:20 | 2021-07-12 23:31:54 |
| [revel](http://revel.github.io) | 12,300 | 1,397 | 101 | A high productivity, full-stack web framework for the Go language. | 2011-12-09 04:10:26 | 2021-07-12 12:23:45 |
| [buffalo](http://gobuffalo.io) | 6,302 | 486 | 82 | Rapid Web Development w/ Go | 2014-10-22 17:35:14 | 2021-07-13 02:17:43 |
| [goa](https://goa.design) | 4,304 | 458 | 41 | Design-based APIs and microservices in Go | 2014-12-05 07:17:53 | 2021-07-12 23:02:15 |
| [go-json-rest](https://ant0ine.github.io/go-json-rest/) | 3,473 | 388 | 49 | A quick and easy way to setup a RESTful JSON API | 2013-02-19 03:15:45 | 2021-07-04 07:17:06 |
| [gizmo](https://open.nytimes.com/introducing-gizmo-aa7ea463b208) | 3,445 | 229 | 28 | A Microservice Toolkit from The New York Times | 2015-12-15 18:09:36 | 2021-07-12 11:22:25 |
| [macaron](https://go-macaron.com) | 3,174 | 283 | 8 | Package macaron is a high productive and modular web framework in Go. | 2014-07-10 03:13:30 | 2021-07-12 08:14:03 |
| [utron](https://github.com/gernest/utron) | 2,198 | 158 | 9 | A lightweight MVC framework for Go(Golang) | 2015-09-16 07:55:54 | 2021-07-11 00:30:27 |
| [rest-layer]( http://rest-layer.io) | 1,084 | 92 | 33 | REST Layer, Go (golang) REST API framework | 2015-07-29 19:16:20 | 2021-07-11 14:14:05 |
| [go-tigertonic](https://github.com/rcrowley/go-tigertonic) | 995 | 74 | 28 | A Go framework for building JSON web services inspired by Dropwizard | 2013-02-09 21:16:13 | 2021-06-21 14:33:15 |
| [goyave](https://goyave.dev) | 863 | 40 | 8 | 🍐 Elegant Golang REST API Framework | 2019-10-21 09:44:34 | 2021-07-11 05:07:20 |
| [tango](https://github.com/lunny/tango) | 834 | 109 | 9 | This is only a mirror and Moved to https://gitea.com/lunny/tango | 2014-12-17 03:07:09 | 2021-05-31 10:00:01 |
| [aah](https://aahframework.org) | 652 | 35 | 16 | A secure, flexible, rapid Go web framework | 2016-06-27 04:47:45 | 2021-07-05 07:26:50 |
| [gearbox](https://gogearbox.com) | 520 | 38 | 0 | Gearbox :gear: is a web framework written in Go with a focus on high performance | 2020-04-25 01:28:37 | 2021-07-05 08:14:49 |
| [gongular](http://gondolaweb.com) | 442 | 15 | 6 | A different approach to Go web frameworks | 2016-06-22 11:52:42 | 2021-06-22 23:19:47 |
| [neo](http://ivpusic.github.io/neo/) | 410 | 40 | 7 | Go Web Framework | 2015-02-04 19:16:06 | 2021-06-13 15:37:14 |
| [air](https://pkg.go.dev/github.com/aofei/air) | 405 | 39 | 3 | An ideally refined web framework for Go. | 2016-07-20 12:09:48 | 2021-07-06 08:27:20 |
| [aero](https://github.com/aerogo/aero) | 384 | 21 | 4 | :bullettrain_side: High-performance web server for Go. | 2016-11-09 13:02:13 | 2021-07-11 00:09:15 |
| [mango](http://github.com/paulbellamy/mango) | 357 | 37 | 9 | Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. | 2011-05-25 07:26:46 | 2021-07-09 10:52:11 |
| [gondola](http://gondolaweb.com) | 308 | 22 | 8 | The web framework for writing faster sites, faster | 2014-07-25 21:28:55 | 2021-06-22 23:20:30 |
| [confetti](https://www.confetti-framework.com) | 273 | 8 | 1 | Confetti is a web application framework with an expressive, elegant syntax. This repository contains configuration files and is intended as a template for your codebase. Download these configuration files and include them in your git repository. | 2019-11-01 23:14:21 | 2021-07-10 12:50:14 |
| [golf](https://golf.readme.io/) | 249 | 26 | 4 | :golf: The Golf web framework | 2015-11-18 15:10:14 | 2021-06-06 14:57:36 |
| [flamingo](http://www.flamingo.me) | 216 | 29 | 16 | Flamingo Framework and Core Library. Flamingo is a go based framework for pluggable web projects. It is used to build scalable and maintainable (web)applications. | 2019-04-02 12:24:02 | 2021-07-12 15:39:46 |
| [beego](beego.me) | 195 | 55 | 0 | beego is an open-source, high-performance web framework for the Go programming language. | 2020-12-13 14:58:50 | 2021-07-12 11:54:21 |
| [flamingo-commerce](https://www.flamingo.me/flamingo-commerce.html) | 191 | 31 | 17 | Flexible E-Commerce Framework on top of Flamingo. Used to build E-Commerce "Portals" and connect it with the help of individual Adapters to other services.  | 2019-04-02 15:11:57 | 2021-07-12 09:34:16 |
| [ginrpc](https://xxjwxc.github.io/post/ginrpc/) | 179 | 19 | 7 | gin auto binding,grpc, and annotated route,gin 注解路由, grpc,自动参数绑定工具 | 2019-06-22 12:03:53 | 2021-07-12 22:59:56 |
| [webgo](https://github.com/bnkamalesh/webgo) | 174 | 17 | 5 | A minimal framework to build web apps; with handler chaining, middleware support; and most of all standard library compliant HTTP handlers(i.e. http.HandlerFunc). | 2015-12-16 07:35:02 | 2021-07-12 07:29:37 |
| [hiboot](https://hiboot.netlify.app/) | 162 | 28 | 5 | hiboot is a high performance web and cli application framework with dependency injection support | 2018-03-16 11:21:46 | 2021-06-26 11:52:55 |
| [uadmin](https://uadmin.io) | 138 | 30 | 11 | The web framework for Golang | 2018-10-05 09:00:17 | 2021-07-10 19:01:55 |
| [go-rest](http://go.pkgdoc.org/github.com/ungerik/go-rest) | 125 | 13 | 2 | A small and evil REST framework for Go | 2012-07-13 10:02:15 | 2021-02-11 17:43:01 |
| [appy](https://github.com/appist/appy) | 96 | 11 | 14 | An opinionated productive web framework that helps scaling business easier. | 2019-05-27 04:48:59 | 2021-07-09 13:04:17 |
| [patron](https://github.com/beatlabs/patron) | 76 | 48 | 28 | Microservice framework following best cloud practices with a focus on productivity. | 2019-01-30 13:49:54 | 2021-07-06 06:48:11 |
| [microservice](http://github.com/paulbellamy/mango) | 74 | 13 | 0 | This library provides a simple framework of microservice, which includes a configurator, a logger, metrics, and of course the handler | 2016-12-15 09:07:04 | 2021-04-15 02:41:22 |
| [vox](https://aisk.github.io/vox/) | 74 | 4 | 7 | Simple and lightweight Go web framework inspired by koa | 2014-12-24 11:22:08 | 2021-07-02 15:21:34 |
| [golax](https://github.com/fulldump/golax) | 73 | 6 | 6 | Golax, a go implementation for the Lax framework. | 2016-01-30 19:11:39 | 2020-08-31 11:59:35 |
| [rux](https://pkg.go.dev/github.com/gookit/rux?tab=doc) | 63 | 11 | 1 | ⚡ Rux is an simple and fast web framework. support middleware, compatible http.Handler interface. 简单且快速的 Go web 框架，支持中间件，兼容 http.Handler 接口 | 2018-08-05 06:13:57 | 2021-06-22 13:19:17 |
| [yarf](https://github.com/yarf-framework/yarf) | 62 | 6 | 1 | Yet Another REST Framework | 2015-09-02 13:56:47 | 2021-06-30 12:49:00 |
| [fireball](https://github.com/zpatrick/fireball) | 56 | 4 | 0 | Go web framework with a natural feel | 2016-07-20 05:04:54 | 2021-04-06 16:51:53 |
| [goa](https://goa-go.github.io) | 45 | 1 | 0 | Goa is a  web framework based on middleware, like koa.js. | 2019-07-26 07:12:23 | 2021-06-06 10:44:15 |
| [gotuna](https://gotuna.org) | 34 | 4 | 0 | GoTuna a lightweight web framework for Go with mux router, middlewares, user sessions, templates, embedded views, and static file server. | 2021-04-08 14:08:08 | 2021-07-10 12:50:34 |
| [api](http://resoursea.com/) | 31 | 2 | 0 | A REST framework for quickly writing resource based services in Golang. | 2015-01-24 18:45:30 | 2020-02-21 12:56:46 |
| [rex](https://github.com/goanywhere/rex) | 31 | 1 | 0 | Pleasures for Web in Golang | 2014-10-16 02:26:18 | 2020-12-22 09:31:27 |
| [goweb](https://github.com/twharmon/goweb) | 24 | 1 | 0 | Lightweight web framework based on net/http. | 2019-05-07 21:04:43 | 2021-05-13 03:52:22 |
| [banjo](https://nsheremet.pw/banjo) | 18 | 5 | 1 | BANjO is a simple web framework written in Go (golang) | 2017-12-09 13:35:31 | 2021-07-11 02:55:22 |
</details>

### Web Frameworks - Middlewares - Actual middlewares


<sup>*Last Update: 2021-07-01 16:25:21*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tollbooth](https://github.com/didip/tollbooth) | 1,961 | 180 | 5 | Simple middleware to rate-limit HTTP requests. | 2015-05-17 15:20:03 | 2021-06-30 01:04:12 |
| [cors](https://github.com/rs/cors) | 1,852 | 159 | 17 | Go net/http configurable handler to handle CORS requests | 2014-10-25 03:49:45 | 2021-06-30 16:27:23 |
| [limiter](https://github.com/ulule/limiter) | 1,288 | 103 | 7 | Dead simple rate limit middleware for Go. | 2015-10-02 08:12:38 | 2021-06-29 13:38:09 |
| [go-server-timing](https://github.com/mitchellh/go-server-timing) | 819 | 26 | 4 | Go (golang) library for creating and consuming HTTP Server-Timing headers | 2018-02-12 03:56:02 | 2021-07-01 05:12:43 |
| [go-fault](https://github.com/github/go-fault) | 398 | 16 | 1 | Fault injection library in Go using standard http middleware | 2020-05-14 16:13:17 | 2021-07-01 09:22:05 |
| [ln-paywall](https://github.com/philippgille/ln-paywall) | 111 | 7 | 17 | Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️ | 2018-06-29 21:51:00 | 2021-06-11 18:31:55 |
| [xff](https://github.com/sebest/xff) | 81 | 17 | 5 | A Golang Middleware to handle X-Forwarded-For Header | 2014-12-22 10:29:05 | 2021-06-30 17:00:43 |
| [formjson](https://github.com/rs/formjson) | 36 | 0 | 0 | Go net/http handler to transparently manage posted JSON | 2015-03-19 23:52:28 | 2020-09-18 01:35:42 |
| [client-timing](https://github.com/posener/client-timing) | 19 | 4 | 1 | An HTTP client for go-server-timing middleware. Enables automatic timing propagation through HTTP calls between servers. | 2018-02-23 01:52:45 | 2021-04-21 08:17:29 |
</details>

### Web Frameworks - Middlewares - Libraries for creating HTTP middlewares


<sup>*Last Update: 2021-07-12 14:22:11*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [negroni](https://github.com/urfave/negroni) | 7,020 | 563 | 9 | Idiomatic HTTP Middleware for Golang | 2014-05-18 22:09:10 | 2021-07-04 12:14:04 |
| [alice](https://godoc.org/github.com/justinas/alice) | 2,295 | 130 | 6 | Painless middleware chaining for Go | 2014-05-25 07:27:41 | 2021-07-03 12:00:43 |
| [render](https://github.com/unrolled/render) | 1,520 | 125 | 1 | Go package for easily rendering JSON, XML, binary data, and HTML templates responses. | 2014-06-10 16:20:35 | 2021-07-11 11:54:00 |
| [stats](https://github.com/thoas/stats) | 579 | 47 | 8 | A Go middleware that stores various information about your web application (response time, status code count, etc.) | 2015-03-05 18:02:50 | 2021-07-01 23:46:48 |
| [interpose](https://github.com/carbocation/interpose) | 290 | 15 | 1 | Minimalist net/http middleware for golang | 2014-07-20 00:19:52 | 2021-06-21 14:54:37 |
| [renderer](https://github.com/thedevsaddam/renderer) | 223 | 24 | 0 | Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go | 2017-11-07 18:53:49 | 2021-06-28 15:20:33 |
| [muxchain](https://github.com/stephens2424/muxchain) | 209 | 13 | 1 | Lightweight Middleware for net/http | 2014-05-03 17:14:17 | 2020-09-04 11:26:14 |
| [rye](https://github.com/InVisionApp/rye) | 97 | 12 | 0 | A tiny http middleware for Golang with added handlers for common needs. | 2016-10-06 19:51:59 | 2021-06-06 09:24:36 |
| [gores](https://github.com/alioygur/gores) | 96 | 2 | 0 | Go package that handles HTML, JSON, XML and etc. responses | 2015-12-25 12:41:01 | 2021-07-04 03:40:01 |
| [mediary](https://github.com/HereMobilityDevelopers/mediary/wiki/Reasoning) | 71 | 4 | 0 | Add interceptors to GO http.Client | 2020-03-23 18:54:56 | 2021-02-10 17:40:20 |
| [chain](https://github.com/codemodus/chain) | 63 | 3 | 0 | Composable chains of nested http.Handler instances. | 2015-05-14 19:52:58 | 2021-02-11 17:26:39 |
| [wrap](https://github.com/go-on/wrap) | 59 | 5 | 0 | Go http.Hander based middleware stack with context sharing | 2014-02-16 07:12:36 | 2020-08-28 05:29:07 |
| [catena](https://github.com/codemodus/catena) | 7 | 0 | 0 | gRPC interceptor catenation. | 2015-07-30 19:07:01 | 2018-08-25 22:06:07 |
</details>

### Web Frameworks - Routers


<sup>*Last Update: 2021-07-05 02:25:10*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [mux](http://www.gorillatoolkit.org/pkg/mux) | 14,639 | 1,365 | 15 | A powerful HTTP router and URL matcher for building Go web servers with 🦍 | 2012-10-02 21:32:24 | 2021-07-04 10:58:12 |
| [httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter) | 12,867 | 1,245 | 61 | A high performance HTTP request router that scales well | 2013-12-05 15:10:55 | 2021-07-04 07:22:24 |
| [chi](https://github.com/go-chi/chi) | 9,648 | 659 | 21 | lightweight, idiomatic and composable router for building Go HTTP services | 2015-10-15 20:46:29 | 2021-07-04 08:39:27 |
| [web](https://github.com/gocraft/web) | 1,443 | 121 | 23 | Go Router + Middleware. Your Contexts. | 2013-11-16 20:48:20 | 2021-06-24 17:51:49 |
| [bone](http://go-zoo.github.io/bone) | 1,273 | 85 | 2 | Lightning Fast HTTP Multiplexer | 2014-11-19 02:16:36 | 2021-06-27 11:38:38 |
| [fasthttprouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 873 | 92 | 19 | A high performance fasthttp request router that scales well | 2015-12-13 09:32:30 | 2021-06-17 13:16:03 |
| [goji](https://goji.io) | 864 | 62 | 6 | Goji is a minimalistic and flexible HTTP request multiplexer for Go (golang) | 2015-11-16 00:52:41 | 2021-07-01 16:40:45 |
| [gorouter](https://xujiajun.cn/gorouter) | 499 | 80 | 1 | xujiajun/gorouter is a simple and fast HTTP router for Go. It is easy to build RESTful APIs and your web framework. | 2018-01-29 09:28:28 | 2021-07-01 09:03:40 |
| [httptreemux](https://github.com/dimfeld/httptreemux) | 492 | 44 | 2 | High-speed, flexible tree-based HTTP router for Go. | 2014-05-14 20:10:20 | 2021-06-20 15:04:16 |
| [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) | 407 | 50 | 11 | An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. | 2015-10-27 01:03:14 | 2021-07-02 23:33:36 |
| [lars](https://github.com/go-playground/lars) | 383 | 23 | 1 | :rotating_light: Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. | 2015-12-24 17:28:45 | 2021-06-24 07:46:46 |
| [siesta](https://github.com/VividCortex/siesta) | 350 | 14 | 0 | Composable framework for writing HTTP handlers in Go. | 2014-09-23 13:55:56 | 2021-05-23 22:02:49 |
| [vestigo](https://github.com/husobee/vestigo) | 262 | 29 | 14 | Echo Inspired Stand Alone URL Router | 2015-09-22 03:08:03 | 2021-05-16 04:33:49 |
| [router](https://github.com/gowww/router) | 160 | 12 | 0 | ⚡️ A lightning fast HTTP router | 2017-05-25 10:29:27 | 2021-04-26 14:57:04 |
| [alien](https://github.com/gernest/alien) | 118 | 10 | 3 | A lightweight and  fast http router from outer space | 2016-01-30 23:23:10 | 2021-06-27 14:41:50 |
| [pure](https://github.com/go-playground/pure) | 114 | 9 | 0 | :non-potable_water: Is a lightweight  HTTP router that sticks to the std "net/http" implementation | 2016-09-23 19:57:58 | 2021-06-25 18:01:39 |
| [violetear](http://violetear.org) | 103 | 8 | 1 | Go HTTP router | 2015-06-19 16:49:41 | 2021-05-25 19:15:16 |
| [Bxog](http://go-zoo.github.io/bone) | 102 | 6 | 0 | Bxog is a simple and fast HTTP router for Go (HTTP request multiplexer). | 2016-05-19 12:20:08 | 2021-06-15 07:22:50 |
| [gorouter](https://rafallorenz.com/gorouter) | 95 | 14 | 3 | Go Server/API micro framework, HTTP request router, multiplexer, mux | 2016-07-14 13:13:34 | 2021-07-03 05:19:08 |
| [xmux](http://violetear.org) | 90 | 9 | 2 | xmux is a httprouter fork on top of xhandler (net/context aware) | 2015-12-14 19:01:05 | 2021-06-13 02:37:37 |
| [bellt](https://github.com/GuilhermeCaruso/bellt) | 49 | 4 | 0 | :bell: A simple Go router | 2019-02-21 13:13:52 | 2021-04-21 03:20:58 |
| [fastrouter](http://godoc.org/github.com/buaazp/fasthttprouter) | 19 | 3 | 0 | FastRouter is a fast, flexible HTTP router written in Go. | 2017-11-01 08:52:52 | 2020-12-30 15:49:44 |
| [route](https://goroute.github.io) | 7 | 1 | 1 | Go Route - Simple yet powerful HTTP request multiplexer | 2019-07-06 18:47:38 | 2020-08-31 13:36:03 |
</details>

### WebAssembly


<sup>*Last Update: 2021-06-29 16:25:17*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [tinygo](https://tinygo.org) | 8,189 | 427 | 264 | Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. | 2018-06-07 16:39:19 | 2021-06-29 08:54:46 |
| [dom](https://github.com/dennwc/dom) | 429 | 51 | 11 | DOM library for Go and WASM | 2018-06-30 18:37:35 | 2021-06-28 15:18:01 |
| [go-canvas](https://github.com/markfarnan/go-canvas) | 137 | 9 | 5 | Library to use HTML5 Canvas  from Go-WASM, with all drawing within go code | 2019-05-05 14:05:55 | 2021-06-25 12:35:37 |
| [webapi](https://github.com/gowebapi/webapi) | 82 | 8 | 2 | Go Lang Web Assembly bindings for DOM, HTML etc | 2019-02-08 05:58:35 | 2021-06-16 15:50:16 |
| [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) | 81 | 9 | 0 | Run WASM tests inside your browser | 2018-07-14 18:42:24 | 2021-06-24 22:24:19 |
| [vert](https://github.com/norunners/vert) | 48 | 6 | 2 | WebAssembly interop between Go and JS values. | 2018-03-25 17:26:47 | 2021-05-23 19:32:09 |
</details>

### Windows


<sup>*Last Update: 2021-06-25 12:40:03*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-ole](https://github.com/go-ole/go-ole) | 759 | 139 | 54 | win32 ole implementation for golang | 2011-01-21 12:45:20 | 2021-06-20 21:21:12 |
| [d3d9](https://github.com/gonutz/d3d9) | 123 | 9 | 0 | Direct3D9 wrapper for Go. | 2015-12-12 21:24:38 | 2021-06-20 12:58:01 |
| [gosddl](https://github.com/MonaxGT/gosddl) | 5 | 1 | 0 | GoSDDL converter | 2018-12-04 08:36:11 | 2021-02-14 13:03:11 |
</details>

### XML
Libraries and tools for manipulating XML.

<sup>*Last Update: 2021-07-01 10:25:19*</sup>
<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [zek](https://github.com/miku/zek) | 468 | 40 | 8 | Generate a Go struct from XML. | 2017-11-23 19:03:11 | 2021-06-29 08:00:10 |
| [xpath](https://github.com/antchfx/xpath) | 399 | 55 | 9 | XPath package for Golang, supports HTML, XML, JSON document query. | 2016-10-09 05:51:24 | 2021-06-23 04:40:20 |
| [xquery](https://github.com/antchfx/xpath) | 154 | 25 | 0 | Extract data or evaluate value from HTML/XML documents using XPath | 2016-10-09 05:54:10 | 2021-06-23 23:30:22 |
| [xml2map](https://github.com/sbabiv/xml2map) | 32 | 5 | 1 | XML to MAP converter written Golang | 2018-08-06 17:51:46 | 2021-06-09 18:08:01 |
| [xmlwriter](https://github.com/shabbyrobe/xmlwriter) | 19 | 3 | 1 | xmlwriter is a pure-Go library providing procedural XML generation based on libxml2's xmlwriter module | 2017-04-11 04:43:26 | 2021-05-26 06:49:04 |
| [XML-Comp](https://github.com/XML-Comp/XML-Comp) | 15 | 8 | 8 | Compare ANY markup documents. | 2016-10-25 22:09:12 | 2021-04-23 04:41:47 |
</details>

