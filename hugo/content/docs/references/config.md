---
categories:
- reference
description: Reference documentation for Homechart's configuration
title: Config
---

Homechart can be configured using [command line arguments]({{< ref "/docs/references/cli#-x-keyvalue" >}}), environment variables or a JSON/Jsonnet configuration file.  The configuration is divided into sections:

- <a href="#app">App</a>
- <a href="#cli">App</a>
- <a href="#postgresql">PostgreSQL</a>
- <a href="#smtp">SMTP</a>
- <a href="#webpush">Web Push</a>

**For command line values**, every configuration key can be set using `-x <a_config_key1>="a value" -x <a_config_key2>="another value"`, i.e. `-x cli_debug=true -x postgresql_username=homechart`.  Config values can also be set using JSON, i.e. `-x webPush='{"vapidPrivateKey": ""}'`

**For environment variables**, every configuration key can be set using `HOMECHART_section_key=a value`, i.e. `HOMECHART_cli_logLevel=debug`

**For configuration files**, they can be formatted using JSON or Jsonnet.  Homechart will look for `homechart.jsonnet` by default, ascending the directory tree to find it.  See [the Jsonnet reference]({{< ref "/docs/references/jsonnet" >}}) for more information.  **Configuration files are rendered at startup**, allowing you to use [dynamic Jsonnet functions]({{< ref "/docs/references/jsonnet#native-functions" >}}) to dynamically alter the config, i.e.:

{{< highlight json >}}
local getRecord(type, name, fallback=null) = std.native('getRecord')(type, name, fallback);
local adminEmailAddress = getRecord('txt', 'adminemail.local');

{
  app: {
    adminEmailAddresses: [
      adminEmailAddress
    ],
  },
}
{{< /highlight >}}

You can view the rendered configuration by running [`homechart show-config`]({{< ref "/docs/references/cli#show-config" >}}).

## `app`

### `adminEmailAddresses` (recommended)

List of strings, email addresses which will have admin access to Homechart for their account.

Default: []

### `baseURL` (recommended)

String, base URL for Homechart, mostly used by notifications.  

Default: "https://web.homechart.app"

### `cacheTTLMinutes`

Number, number of minutes to keep entries in cache.

Default: 15

### `demo`

Boolean, allow users to create demo logins, mostly used by web.homechart.app.  Easy way to try out Homechart.

Default: false

### `disableTasks`

Boolean, disable any background tasks (sending notifications, cleaning up old data) from running on this instance.

Default: false

### `keepCalendarEventDays`

Number, number of days to retain Calendar Events after their end date.

Default: 90

### `keepCookMealPlanDays`

Number, number of days to retain Cook Meal Plans after their scheduled date.

Default: 90

### `keepDeletedDays`

Number, number of days to keep deleted data (Notes, Recipes).

Default: 30

### `keepHealthLogDays`

Number, number of days to retain Health Logs.

Default: 90

### `keepNotesPageVersions`

Number, number of Notes Page Versions to keep.

Default: 10

### `keepPlanTasksDays`

Number, number of days to keep completed tasks.

Default: 90

### `motd`

String, informational message to display on the UI for all users.

Default: ""

### `port` {#app-port}

Number, listening port for Homechart.  Setup port forwarding to this port to expose Homechart externally.

Default: 3000

### `proxyAddress` {#app-proxyaddress}

String, Upstream IPv4 or IPv6 address of a trusted proxy.  See the [Single Sign-On (SSO) guide]({{< ref "/docs/guides/get-homechart/self-hosted/sso" >}}) for usage details.

Default: ""

### `proxyHeaderEmail` {#app-proxyheaderemail}

String, proxy header that should be associated with an account email address.  See [Single Sign-On (SSO) guide]({{< ref "/docs/guides/get-homechart/self-hosted/sso" >}}) for usage details.

Default: ""

### `proxyHeaderName` {#app-proxyheadername}

String, proxy header that should be associated with an account name.  See [Single Sign-On (SSO) guide]({{< ref "/docs/guides/get-homechart/self-hosted/sso" >}}) for usage details.

Default: ""

### `rateLimiterRate`

String, maximum number of requests over a specific time to public endpoints.  Prevents brute force attacks.  Takes the format of (number-H/M/S) where H=hour, M=minute, S=Second.  The default, 15-H, means 15 requests per hour.

Default: "15-H"

### `rollupBudgetTransactionsBalanceMonths`

Number, number of months before a Budget Transaction is rolled up into a starting balance.

Default: 48

### `rollupBudgetTransactionsSummaryMonths`

Number, number of months before a Budget Transaction is rolled up into monthly summaries.

Default: 12

### `sessionExpirationDefaultSeconds`

Number, time between non-Remember Me sessions expiring, in seconds.

Default: 3600

### `sessionExpirationRememberSeconds`

Number, time between Remember Me sessions expiring, in seconds.

Default: 7776000

### `signupDisabled` (recommended)

Boolean, disables new account signups.  Accounts can still be created/invited under the Admin > Accounts.  Self-hosted instances should enable this after they have setup their household.

Default: false

### `tlsCertificate` (recommended)

String, path to a SSL/TLS certificate file.  Should work for the domain in your [baseURL](#baseurl-recommended).  If set, along with [tlsKey](#tlsKey), Homechart will listen for HTTPS connections only.  HTTPS is necessary for certain Homechart functionality, you should enable HTTPS within Homechart or your reverse proxy.

Default: ""

### `tlsKey` (recommended)

String, path to a SSL/TLS private key file.  Should work for the domain in your [baseURL](#baseurl-recommended).  If set, along with [tlsCertificate](#tlscertificate), Homechart will listen for HTTPS connections only.  HTTPS is necessary for certain Homechart functionality, you should enable HTTPS within Homechart or your reverse proxy.

Default: ""

## `cli`

### `configPath`

String, path to the configuration file.  If a filename without a path is specified, Homechart will search parent directories for the filename and use the first one found.

**Default:** `"homechart.jsonnet"`

### `logFormat`

String, log format to use for logging: human, kv, or raw.

**Default:** `"human"`

### `logLevel`

String, log level to use for logging: none, debug, info, or error.

**Default:** `"info"`

### `noColor`

Boolean, disable colored logging.

Default: false

## `postgresql`

### `database` (required)

String, database to use when connecting to PostgreSQL.

Default: ""

### `hostname` (required) {#postgresql-hostname}

String, hostname to use when connecting to PostgreSQL.

Default: "localhost"

### `maxConnections`

Number, maximum number of open connections to PostgreSQL.

Default: 25

### `maxIdleConnections`

Number, maximum number of idle connections to PostgreSQL.

Default: 5

### `maxLifetimeMinutes`

Number, maximum number of minutes to keep a connection to PostgreSQL open.

Default: 5

### `password` (required) {#postgresql-password}

String, password to use when connecting to PostgreSQL.

Default: ""

### `port` {#postgresql-port}

Number, port to use when connecting to PostgreSQL.

Default: 5432

### `sslMode`

String, postgreSQL SSL/TLS enforcement level.

Default: "disable"

### `username` (required) {#postgresql-username}

String, username to use when connecting to PostgreSQL.

Default: ""

## `smtp`

Homechart can use a SMTP server to send notifications to your household members.

### `fromAddress`

String, email address to send from.  Required to make SMTP work.

Default: ""

### `hostname` {#smtp-hostname}

String, hostname to use when connecting to SMTP server.

Default: ""

### `noEmailDomains`

List of strings, domains that will not be verified.  Use this to automatically activate accounts for each domain listed.

Default: []

### `password` {#smtp-password}

String, password to use when connecting to SMTP server.

Default: ""

### `port` {#smtp-port}

Number, port to use when connecting to SMTP server.

Default: 587

### `replyTo`

String, email address to have users send to when replying.

Default: ""

### `username` {#smtp-username}

String, username to use when connecting to SMTP server.

Default: ""

## `webPush`

Homechart can use [Web Push](https://developer.mozilla.org/en-US/docs/Web/API/Push_API) to send push notifications from Homechart to your devices.  Homechart communicates directly to web push services provided by Apple, Google, Mozilla and other standards-compliant endpoints.  Additionally, all of the data in the push notification is encrypted between your server and the client--the web push services can't read it.

You need to generate the VAPID private and public keys to use Web Push.  This can be done from the command line, e.g.:

```shell
$ ./homechart_linux_amd64 generate-vapid
{
  "privateKey": "VEIYXV6qF_enUzycyQYdplDUgi05UM4lPh_FTzYmwX8",
  "publicKey": "BNh2dabXjc2N8mctezlEm5pd1-1m_kkVZpdNYJl5gtRtdmKNIZvA6IZwYEYSy5UmVr5N7Bt9y9qKCLTp1sc_89c"
}
```

Or using a container:

```shell
$ docker run -it --rm ghcr.io/candiddev/homechart generate-vapid
{
  "privateKey": "VEIYXV6qF_enUzycyQYdplDUgi05UM4lPh_FTzYmwX8",
  "publicKey": "BNh2dabXjc2N8mctezlEm5pd1-1m_kkVZpdNYJl5gtRtdmKNIZvA6IZwYEYSy5UmVr5N7Bt9y9qKCLTp1sc_89c"
}
```

This command will output the private and public keys you'll use in the configuration sections below.

### `vapidPrivateKey`

String, the privateKey value from running `generate-vapid`.

Default: ""

### `vapidPublicKey`

String, the publicKey value from running `generate-vapid`.

Default: ""
