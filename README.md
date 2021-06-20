# Go Libraries
List of Go frameworks, libraries and software inspired by [go-web-framework-stars](https://github.com/mingrammer/go-web-framework-stars).

List of frameworks and libraries from [awesome-go](https://github.com/avelino/awesome-go).

### Contents
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
* [IoT (Internet of Things)](#iot-(internet-of-things))
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
## Audio and Music
Libraries for manipulating audio.

*Last Update: 2021-06-20 09:29:16*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [gosamplerate](https://github.com/dh1tw/gosamplerate) | 10 | 5 | 0 | Go Bindings for libsamplerate | 2016-11-20 21:19:31 | 2020-07-10 23:35:35 |
| [vorbis](https://github.com/mccoyst/vorbis) | 28 | 4 | 2 | A "native" ogg vorbis decoder for Go (uses inline stb_vorbis) | 2013-07-12 02:45:39 | 2020-10-04 21:02:12 |
| [minimp3](https://github.com/tosone/minimp3) | 55 | 9 | 0 | Decode mp3 base on https://github.com/lieff/minimp3 | 2018-01-26 09:10:31 | 2021-06-09 06:01:11 |
| [gaad](https://github.com/Comcast/gaad) | 84 | 12 | 4 | GAAD (Go Advanced Audio Decoder) | 2016-07-11 14:19:16 | 2021-06-09 06:01:06 |
| [GoAudio](https://github.com/DylanMeeus/GoAudio) | 136 | 9 | 4 | Go tools for audio processing & creation 🎶 | 2020-07-05 01:35:41 | 2021-06-19 10:58:02 |
| [mix](https://gopkg.in/mix.v0) | 138 | 22 | 11 | Sequence-based Go-native audio mixer for music apps | 2016-01-03 15:53:57 | 2021-06-14 00:15:40 |
| [malgo](https://pkg.go.dev/github.com/bogem/id3v2) | 149 | 23 | 1 | Mini audio library | 2017-11-09 18:27:52 | 2021-06-01 04:53:41 |
| [flac](https://github.com/mewkiz/flac) | 157 | 30 | 10 | Package flac provides access to FLAC (Free Lossless Audio Codec) streams. | 2012-11-01 20:13:58 | 2021-06-16 02:49:09 |
| [id3v2](https://pkg.go.dev/github.com/bogem/id3v2) | 190 | 31 | 14 | 🎵 ID3 decoding and encoding library for Go | 2016-05-15 18:36:53 | 2021-05-21 04:31:20 |
| [portmidi](https://github.com/rakyll/portmidi) | 256 | 49 | 12 | Go bindings for libportmidi | 2013-11-10 14:24:56 | 2021-06-10 02:27:25 |
| [waveform](https://github.com/mdlayher/waveform) | 337 | 28 | 4 | Go package capable of generating waveform images from audio streams. MIT Licensed. | 2014-09-13 18:07:36 | 2021-06-04 02:49:48 |
| [music-theory](https://gopkg.in/music-theory.v0) | 342 | 35 | 8 | Go models of Note, Scale, Chord and Key | 2016-03-17 03:50:17 | 2021-06-15 04:26:32 |
| [portaudio](https://github.com/gordonklaus/portaudio) | 451 | 71 | 9 | Go bindings for the PortAudio audio I/O library | 2015-09-16 07:59:23 | 2021-06-15 16:01:18 |
| [oto](https://github.com/hajimehoshi/oto) | 821 | 67 | 33 | ♪ A low-level library to play sound on multiple platforms ♪ | 2017-05-04 12:16:30 | 2021-06-15 13:59:34 |
</details>

## Authentication and OAuth
Libraries for implementing authentications schemes.

*Last Update: 2021-06-20 09:29:16*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [casbin](https://casbin.org/) | 1 | 0 | 0 | An authorization library that supports access control models like ACL, RBAC, ABAC in Golang | 2021-05-29 04:09:46 | 2021-05-30 16:41:50 |
| [cookiestxt](https://casbin.org/) | 6 | 2 | 0 | cookiestxt implement parser of cookies txt format | 2017-10-09 11:27:19 | 2021-03-08 11:45:58 |
| [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) | 9 | 1 | 0 | A driver for the SessionGate Redis module - easy session management using the Go language. | 2017-10-20 03:39:11 | 2020-08-18 23:01:11 |
| [scope](https://github.com/SonicRoshan/scope) | 12 | 3 | 0 | Easily Manage OAuth2 Scopes In Go | 2019-09-23 10:48:14 | 2021-04-25 07:45:30 |
| [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) | 19 | 1 | 0 | Golang library for providing a canonical representation of email address. | 2020-08-21 23:13:04 | 2021-06-18 21:02:34 |
| [otpgo](https://github.com/jltorresm/otpgo) | 19 | 0 | 1 | Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. | 2020-08-19 13:20:23 | 2021-05-23 08:24:17 |
| [securecookie](https://github.com/chmike/securecookie) | 51 | 6 | 1 | Fast, secure and efficient secure cookie encoder/decoder  | 2017-09-03 14:40:29 | 2021-06-15 16:21:04 |
| [sessions](https://github.com/adam-hanna/sessions) | 59 | 6 | 2 | A dead simple, highly performant, highly customizable sessions middleware for go http servers. | 2017-04-29 01:09:28 | 2021-04-23 00:07:10 |
| [rbac](https://github.com/zpatrick/rbac) | 79 | 11 | 0 | Minimalistic RBAC package for Go applications | 2018-08-02 00:11:04 | 2021-05-31 10:33:51 |
| [sjwt](https://godoc.org/github.com/brianvoe/sjwt) | 88 | 5 | 0 | Simple JWT Golang | 2019-06-20 04:06:21 | 2021-06-07 18:24:25 |
| [jwt](https://github.com/robbert229/jwt) | 90 | 21 | 9 | This is an implementation of JWT in golang! | 2016-06-05 22:01:37 | 2021-06-13 21:57:30 |
| [session](https://github.com/icza/session) | 107 | 12 | 4 | Go session management for web servers (including support for Google App Engine - GAE). | 2016-02-08 09:07:07 | 2021-05-02 03:01:38 |
| [sessionup](https://github.com/swithek/sessionup) | 112 | 5 | 3 | Straightforward HTTP session management | 2019-07-23 18:55:21 | 2021-06-08 14:33:23 |
| [branca](https://branca.io) | 152 | 15 | 1 | :key: Secure alternative to JWT. Authenticated Encrypted API Tokens for Go. | 2018-01-09 15:27:31 | 2021-05-22 13:49:09 |
| [jwt-auth](https://github.com/adam-hanna/jwt-auth) | 204 | 35 | 3 | This package provides json web token (jwt) middleware for goLang http servers | 2016-07-05 23:31:43 | 2021-05-24 01:31:29 |
| [httpauth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 206 | 21 | 4 | HTTP Authentication middlewares | 2014-05-26 22:53:57 | 2021-06-09 22:15:48 |
| [jeff](https://github.com/abraithwaite/jeff) | 225 | 12 | 2 | 🍍Jeff provides the simplest way to manage web sessions in Go. | 2018-08-02 19:31:23 | 2021-04-10 19:18:30 |
| [jwt](https://github.com/pascaldekloe/jwt) | 256 | 15 | 0 | JSON Web Token library | 2018-03-21 11:59:53 | 2021-06-17 22:18:46 |
| [go-guardian](https://github.com/shaj13/go-guardian) | 266 | 22 | 5 | Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication. | 2020-05-14 12:15:56 | 2021-06-18 12:54:00 |
| [jwt](https://github.com/cristalhq/jwt) | 269 | 24 | 4 | Safe, simple and fast JSON Web Tokens for Go | 2019-07-20 18:14:58 | 2021-06-18 17:10:48 |
| [permissions2](https://github.com/xyproto/permissions2) | 430 | 33 | 0 |   :closed_lock_with_key: Middleware for keeping track of users, login states and permissions | 2014-11-19 12:23:37 | 2021-06-18 13:01:24 |
| [paseto](https://github.com/o1egl/paseto) | 517 | 22 | 3 | Platform-Agnostic Security Tokens implementation in GO (Golang) | 2018-01-23 05:27:39 | 2021-06-14 12:58:41 |
| [scs](https://github.com/alexedwards/scs) | 886 | 83 | 17 | HTTP Session Management for Go | 2016-08-08 16:42:05 | 2021-06-18 09:44:43 |
| [gorbac](https://github.com/mikespook/gorbac) | 1,183 | 139 | 2 | goRBAC provides a lightweight role-based access control (RBAC) implementation in Golang. | 2013-12-26 10:00:41 | 2021-06-18 04:41:52 |
| [gologin](https://github.com/dghubble/gologin) | 1,393 | 110 | 0 | Go login handlers for authentication providers (OAuth1, OAuth2) | 2015-06-23 04:40:52 | 2021-06-19 07:59:58 |
| [osin](https://golang.org/x/oauth2) | 1,671 | 369 | 3 | Golang OAuth2 server library | 2013-09-10 19:52:00 | 2021-06-16 06:45:18 |
| [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) | 1,766 | 271 | 28 | A standalone, specification-compliant,  OAuth2 server written in Golang. | 2015-11-01 13:30:09 | 2021-06-18 14:02:09 |
| [go-jose](https://github.com/square/go-jose) | 1,775 | 315 | 49 | An implementation of JOSE standards (JWE, JWS, JWT) in Go | 2014-11-14 18:27:31 | 2021-06-18 21:51:01 |
| [loginsrv](https://github.com/tarent/loginsrv) | 1,789 | 154 | 29 | JWT login microservice with plugable backends such as OAuth2, Google, Github, htpasswd, osiam, .. | 2016-11-11 12:11:21 | 2021-06-18 20:10:47 |
| [authboss](https://github.com/volatiletech/authboss) | 2,685 | 170 | 30 | The boss of http auth. | 2015-01-03 05:12:02 | 2021-06-18 11:56:15 |
| [goth](https://blog.gobuffalo.io/goth-needs-a-new-maintainer-626cd47ca37b) | 3,205 | 392 | 58 | Package goth provides a simple, clean, and idiomatic way to write authentication packages for Go web applications. | 2014-10-14 20:38:12 | 2021-06-19 17:58:10 |
| [oauth2](https://golang.org/x/oauth2) | 3,685 | 764 | 143 | Go OAuth2 | 2014-04-14 15:07:35 | 2021-06-19 22:17:14 |
</details>

## Build Automation
Libraries and tools helping with build automation.

*Last Update: 2021-06-20 09:29:16*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [anko](https://github.com/GuilhermeCaruso/anko) | 16 | 0 | 0 | :crystal_ball: Simple application watcher | 2021-03-02 14:08:42 | 2021-06-18 10:08:52 |
| [gaper](https://github.com/maxcnunes/gaper) | 49 | 3 | 7 | Builds and restarts a Go project when it crashes or some watched file changes | 2018-06-16 02:46:38 | 2021-06-08 12:55:32 |
| [gilbert](https://go-gilbert.github.io/) | 89 | 4 | 0 | Build system and task runner for Go projects | 2019-01-30 09:02:31 | 2021-05-25 10:17:10 |
| [1build](https://1build.gitbook.io) | 96 | 25 | 31 | Frictionless way of managing project-specific commands | 2019-04-23 17:05:38 | 2021-05-26 16:00:32 |
| [taskctl](https://github.com/taskctl/taskctl) | 112 | 8 | 7 | Concurrent task runner, developer's routine tasks automation toolkit. Simple modern alternative to GNU Make 🧰 | 2019-11-12 13:19:09 | 2021-06-04 05:26:40 |
| [goyek](https://github.com/goyek/goyek) | 231 | 15 | 4 | Create build pipelines in Go  | 2020-10-11 13:20:55 | 2021-06-18 06:37:16 |
| [mmake](https://github.com/tj/mmake) | 1,575 | 43 | 11 | Modern Make  | 2017-02-15 22:01:21 | 2021-06-19 11:00:22 |
| [task](https://taskfile.dev) | 3,557 | 228 | 88 | A task runner / simpler Make alternative written in Go | 2017-02-27 00:46:04 | 2021-06-19 08:41:53 |
| [realize](https://github.com/oxequa/realize) | 4,035 | 211 | 68 | Realize is the #1 Golang Task Runner which enhance your workflow by automating the most common tasks and using the best performing Golang live reloading. | 2016-07-12 08:07:25 | 2021-06-19 22:01:58 |
</details>

## CSS Preprocessors
Libraries for preprocessing CSS files.

*Last Update: 2021-06-20 09:29:16*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [go-libsass](http://godoc.org/github.com/wellington/go-libsass) | 178 | 22 | 13 | Go wrapper for libsass, the only Sass 3.5 compiler for Go | 2015-04-19 15:09:47 | 2021-06-18 22:31:34 |
| [gcss](https://github.com/yosssi/gcss) | 442 | 33 | 8 | Pure Go CSS Preprocessor | 2014-09-04 14:38:20 | 2021-05-26 15:17:44 |
</details>

## Dynamic DNS
Tools for updating dynamic DNS records.

*Last Update: 2021-06-20 09:29:16*

<details>
  <summary>Show Table</summary>

| Project Name | Stars | Forks | Open Issues | Description | Created At | Last Update |
| ------------ | ----- | ----- | ----------- | ----------- | ---------- | ----------- |
| [ddns](https://github.com/skibish/ddns) | 191 | 19 | 0 | Personal DDNS client with Digital Ocean Networking DNS as backend. | 2017-03-13 21:02:27 | 2021-06-16 15:50:33 |
| [godns](https://xiaozhou.net/godns-project-2014-05-18.html) | 828 | 158 | 7 | A dynamic DNS client tool supports AliDNS, Cloudflare, Google Domains, DNSPod, HE.net & DuckDNS & DreamHost, etc, written in Go. | 2014-05-11 11:49:17 | 2021-06-19 08:27:19 |
</details>

