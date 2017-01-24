harmond
=======

## Goals

- Implement a secure, fast, stable and easy to maintain IRC daemon
- Implement a better bot API for services, webhooks and network specific features
- Websocket support
- HTTP long-polling support
- Let's Encrypt support
- Use a private TLS CA for federation validation
- RAFT clustering for sharing persistent state
- use pub-sub for making things simpler

## Commands

### ACCEPT

```
ACCEPT <parameter>

ACCEPT allows you to control who can send you a NOTICE or PRIVMSG
while you have user mode +g enabled.

Accepted users can also message you if you have user mode +R
enabled, even if they are not logged in with services.

For +g: /QUOTE ACCEPT <nick>   -- Add a permitted nickname
        /QUOTE ACCEPT -<nick>  -- Remove a permitted nickname
        /QUOTE ACCEPT *        -- List the present permitted nicknames
```

#### Event

```go
type AcceptAdd struct {
    Nick string
}

type AcceptRemove struct {
    Nick string
}

type AcceptList struct{}
```

#### Effects

User accept list is modified

### ADMIN

```
ADMIN [server]

With no arguments, ADMIN shows the information that was set by the
administrator of the server. This information can take any form that
will fit in three lines of text but is usually a list of contacts
for the persons that run the server.

With a second argument, the administrative information for the
specified server is displayed.

See also: stats
```

#### Event

```go
type Admin struct{
    Server *string
}
```

#### Effects

None

### AWAY

```
AWAY :[MSG]

Without an argument, it will set you back.  With an argument,
it will set you as AWAY with the specified message.
```

#### Event

```go
type AwaySetReason struct {
    Reason string
}

type AwayClearReason struct{}
```

### CHGHOST

```
CHGHOST <target> <host>

Will attempt to change <target> hostname to <host>

- Requires Oper Priv: oper:chghost
```

#### Event

```go
type Chghost struct {
    Target, Host string
}
```

#### Effects

Changes target user's vhost to the given host.

### DLINE

```
DLINE [duration] <ip> :[reason] [| oper reason]

Adds a DLINE to the database which will deny any
connections from the IP address of the banned client.
The banned client will receive a message saying
they are banned with reason [reason].

Duration is optional, and is in minutes. If specified,
the DLINE will not be saved in the database.

If an oper reason is added (the pipe must be specified
to seperate the fields) this will be added into the
database but will not be shown to the user when they
are given the kline reason.

DLINE [duration] <ip> ON irc.server :[reason] [| oper reason]
will dline the user on irc.server if irc.server accepts
remote dlines. irc.server can contain wildcards.

- Requires Oper Priv: oper:kline
```

#### Event

```go
type Dline struct {
    Duration time.Duration
    IP, Reason, OperReason string
}
```

#### Effects

Adds a dline to the database and removes any matching users.

### HELP

```
HELP [topic]

HELP displays the contents of the help
file for topic requested.  If no topic is
requested, it will perform the equivalent
to HELP index.
```

#### Event

```go
type Help struct {
    Topic string
}
```

#### Effects

None

### INVITE

```
INVITE <nickname> <channel>

INVITE sends a notice to the user that you have
asked him/her to come to the specified channel.
If the channel is +i, +j, +l or +r, this will
allow the user to bypass these modes.
```

#### Event

```go
type Invite struct {
    Invitee, ChannelName string
}
```

#### Effects

Invitee is allowed to join the given channel name even though they otherwise
would not have been able to.

### JOIN

```
JOIN <#channel> [key]

The JOIN command allows you to enter a public chat area known as
a channel. Network wide channels are proceeded by a '#', while
a local server channel is proceeded by an '&'. More than one
channel may be specified, separated with commas (no spaces).

If the channel has a key set, the 2nd argument must be
given to enter. This allows channels to be password protected.

See also: part, list
```

#### Event

```go
type Join struct {
    Channel, Key string
}
```

#### Effects

Adds user to channel and announces the join where applicable

### KICK

```
KICK <channel> <nick> :[msg]

The KICK command will remove the specified user
from the specified channel, using the optional
kick message.  You must be a channel operator to
use this command.
```

#### Event

```go
type Kick struct {
    Channel, Nick, Message string
}
```

#### Effects

Removes a user from a channel with an optional message to be shown on removal.

### KILL

```
KILL <nick> <reason>

Disconnects user <nick> from the IRC server he/she
is connected to with reason <reason>.

- Requires Oper Priv: oper:local_kill
- Requires Oper Priv: oper:global_kill for users not on your IRC server
```

#### Event

```go
type Kill struct {
    Nick, Reason string
}
```

### KLINE

```
KLINE <user@host> :[reason] [| oper reason]

Adds a KLINE to the database which will ban the
specified user from using this server.  The banned
client will receive a message saying he/she is banned
with reason [reason].

If an oper reason is added (the pipe must be specified
to seperate the fields) this will be added into the
database but will not be shown to the user when they
are given the kline reason.

KLINE <user@ip.ip.ip.ip> :[reason] [| oper reason]
will kline the user at the unresolved ip.
ip.ip.ip.ip can be in CIDR form i.e. 192.168.0.0/24
or 192.168.0.* (which is converted to CIDR form internally)

For a temporary KLINE, length of kline is given in
minutes as the first parameter i.e.
KLINE 10 <user@host> :cool off for 10 minutes

KLINE [duration] <user@host> ON irc.server :[reason] [| oper reason]
will kline the user on irc.server if irc.server accepts
remote klines. irc.server can contain wildcards.

- Requires Oper Priv: oper:kline
```

#### Event

```go
type Kline struct {
  Duration time.Duration
  User, Host, Reason, OperReason string
}
```

#### Effects

Sets a kline in the database and kills any matching users.

### KNOCK

```
KNOCK <channel>

KNOCK requests access to a channel that
for some reason is not open.

KNOCK cannot be used if you are banned, the
channel is +p, or it is open.
```

#### Event

```go
type Knock struct {
  Channel string
}
```

#### Effects

Sends a "channel has been knocked" action to all channel halfops and up.

### LIST

```
LIST [#channel]|[modifiers]

Without any arguments, LIST will give an entire list of all
channels which are not set as secret (+s). The list will be in
the form:

  <#channel> <amount of users> :[topic]

If an argument supplied is a channel name, LIST will give just
the statistics for the given channel.

Modifiers are also supported, seperated by a comma:
  <n - List channels with less than n users
  >n - List channels with more than n users
  C<n - List channels created in the last n minutes
  C>n - List channels older than n minutes
  T<n - List channels whose topics have changed in the
        last n minutes
  T>n - List channels whose topics were last changed
        more than n minutes ago

eg LIST <100,>20
```

#### Event

```go
type ListChannel struct {
  ChannelName string
}

type ListModifiers struct {
  Modifiers string
}
```

#### Effects

None

### MODE

#### Event

```go

```

### MONITOR

```
MONITOR <action> [nick[,nick]*]

Manages the online-notification list (similar to WATCH elsewhere). The
<action> must be a single character, one of:

  +   adds the given list of nicknames to the monitor list, returns
      each given nickname's status as RPL_MONONLINE or RPL_MONOFFLINE
      numerics

  -   removes the given list of nicknames from the monitor list, does
      not return anything

  C   clears the monitor list, does not return anything

  L   returns the current monitor list as RPL_MONLIST numerics,
      terminated with RPL_ENDOFMONLIST

  S   returns status of each monitored nickname, as RPL_MONONLINE or
      RPL_MONOFFLINE numerics

For example:

  MONITOR + jilles,kaniini,tomaw

RPL_MONONLINE numerics return a comma-separated list of nick!user@host
items. RPL_MONOFFLINE and RPL_MONLIST numerics return a comma-separated
list of nicknames.
```

#### Event

```go
type Monitor struct {
  Nicks []string
}

type MonitorAdd Monitor
type MonitorRemove Monitor
type MonitorClear struct{}
type MonitorList struct{}
type MonitorStatus struct{}
```

#### Effects

MonitorAdd changes a client's monitor list
MonitorRemove changes a client's monitor list

### MOTD

```
MOTD

MOTD will display the message of the day.
```

#### Event

```go
type Motd struct{}
```

### NAMES

```
NAMES <channel>

With the #channel argument, it displays the nicks on that channel,
also respecting the +i flag of each client. If the channel specified
is a channel that the issuing client is currently in, all nicks are
listed in similar fashion to when the user first joins a channel.

See also: join
```

#### Event

```go
type Names struct {
  Channel string
}
```

#### Effects

None

### NICK

```
NICK <nickname>

When first connected to the IRC server, NICK is required to
set the client's nickname.

NICK will also change the client's nickname once a connection
has been established.
```

#### Event

```go
type NickPrereg struct{
  Nick string
}

type Nick struct{
  Nick string
}
```

#### Effects

NickPrereg changes a pre-registration client's nickname
Nick changes a client's nickname

### NOTICE

```
NOTICE <nick|channel> :message

NOTICE will send a notice message to the
user or channel specified.
```

#### Event

```go
type NoticeUser struct {
  User, Message string
}

type NoticeChannel struct {
  Channel, Message string
}
```

#### Effects

NoticeUser sends the given user a NOTICE
NoticeChannel sends the given channel a NOTICE

### PART

```
PART <#channel> :[part message]

PART requires at least a channel argument to be given. It will
exit the client from the specified channel. More than one
channel may be specified, separated with commas (no spaces).

An optional part message may be given to be displayed to the
channel.

See also: join
```

#### Event

```go
type PartChannel struct {
  Channel string
  Reason *string
}
```

#### Effects

PartChannel makes the associated user part the given channel listing Reason if given.

### PASS

```
PASS <password>

PASS is used during registration to access
a password protected auth {} block.
```

#### Event

```go
type Pass struct {
  Password string
}
```

#### Effects

Pass will trigger account password checking during registration, and log the user into
their services account automatically.

### PING

```
PING <data>

PING will request a PONG from the target. The given data
will be returned.
```

#### Event

```go
type Ping struct {
  Data *string
}
```

#### Effects

Ping replies with a Pong and the same data given as an argument.

### PRIVMSG

```
PRIVMSG <nick|channel> :message

PRIVMSG will send a standard message to the
user or channel specified.
```

#### Event

```go
type MessageUser struct {
  User, Message string
}

type MessageChannel struct {
  Channel, Message string
}
```

#### Effects

MessageUser sends the given user a PRIVMSG
MessageChannel sends the given channel a PRIVMSG

### QUIT

```
QUIT :[quit message]

QUIT closes this connection to the IRC server.

If a registered user sends a QUIT, they will still be present
in chatrooms, scrollback will be colllected for all non-secret
or private channels to be replayed on reconnect.
```

#### Event

```go
type Quit struct {
  Message *string
}
```

#### Effects

Quit closes the local socket connection.

### RESV

```
RESV [time] <channel|nick> :<reason>

Reserves a channel or nickname from use.  If [time] is not specified this
is added to the database, otherwise is temporary for [time] minutes.

Nick resvs accept the same wildcard chars as xlines.
Channel resvs only use exact string comparisons.
```

#### Event
```go
type ResvNick struct {
  Dur time.Duration // 0 = forever
  Nick string
  PublicReason string
  PrivateReason *string
}

type ResvChannel struct {
  Dur time.Duration // 0 = forever
  Channel string
  PublicReason string
  PrivateReason *string
}
```

#### Effects

ResvNick blocks the usage of the given nickname.
ResvChannel blocks the usage of the given channel.

### TESTGECOS

```
TESTGECOS <gecos>

Looks for matching xlines for the given gecos.
```

### TESTLINE

```
TESTLINE [[nick!]user@]host

Looks up given mask, looking for any matching I/K/D lines.
If username is not specified, it will look up "dummy@host".
If nickname is specified it will also search for RESVs.

This command will not perform dns lookups on a host, for best
results you must testline a host and its IP form.

TESTLINE <#channel>

Shows whether the channel is reserved or not.
```

### TESTMASK

```
TESTMASK <[nick!]user@host> [:gecos]

Will test the given nick!user@host gecos mask, reporting how many local
and global clients match the given mask.  Supports using CIDR ip masks
as a host.
```

### TOPIC

```
TOPIC <#channel> :[new topic]

With only a channel argument, TOPIC shows the current topic of
the specified channel.

With a second argument, it changes the topic on that channel to
<new topic>.  If the channel is +t, only chanops may change the
topic.

See also: cmode
```

### UNDLINE

```
UNDLINE <ip>

Will attempt to undline the given <ip>

- Requires Oper Priv: oper:unkline
```

### UNKLINE

```
UNKLINE <user@host>

Will attempt to unkline the given <user@host>
Will unkline a temporary kline.

UNKLINE <user@host> ON irc.server will unkline
the user on irc.server if irc.server accepts
remote unklines.

- Requires Oper Priv: oper:unkline
```

### UNREJECT

```
UNREJECT <ip>

Removes an IP address from the reject cache.  IP
addresses are added to the reject cache if they are
rejected (e.g. connect and are K-lined) several
times in a short period of time.
```

### UNRESV

```
UNRESV <channel|nick>

-- Remove a RESV on a channel or nick
Will attempt to remove the resv for the given
channel/nick.
```

### UNXLINE

```
UNXLINE <gecos>

Will attempt to unxline the given <gecos>


UNXLINE <gecos> ON <server>

Will attempt to unxline the given <gecos> on <server>.

- Requires Oper Priv: oper:xline
```

### USER

```
USER <username> <unused> <unused> :<real name/gecos>

USER is used during registration to set your gecos
and to set your username if the server cannot get
a valid ident response.  The second and third fields
are not used, but there must be something in them.
The reason is backwards compatibility.
```

#### Event

```go
type User struct {
    Username string
    Realname string
}
```

#### Effects

User usually is run after Nick to finish client registration.

### USERS

```
USERS [remoteserver]

USERS will display the local and global current
and maximum user statistics for the specified
server, or the local server if there was no
parameter.
```

### VERSION

```
VERSION [servername]

VERSION will display the server version of the specified
server, or the local server if there was no parameter.
```

### WHO

```
WHO <#channel|nick>

The WHO command displays information about a user,
such as their GECOS information, their user@host,
whether they are an IRC operator or not, etc.  A
sample WHO result from a command issued like
"WHO pokey" may look something like this:

#lamers ~pokey ppp.example.net irc.example.com pokey H :0 Jim Jones

Clients often reorder the fields; the order in the
IRC protocol is described here.

The first field indicates the last channel the user
has joined.  The second is the username and the third
is the host.  The fourth field is the server the user
is on.  The fifth is the user's nickname.  The sixth
field describes status information about the user.
The possible combinations for this field are listed
below:

H       -       The user is not away.
G       -       The user is set away.
*       -       The user is an IRC operator.
@       -       The user is a channel op in the channel listed
                in the first field.
+       -       The user is voiced in the channel listed.

The final field displays the number of server hops and
the user's GECOS information.

This command may be executed on a channel, such as
"WHO #lamers".  The output will consist of WHO
listings for each user on the channel.  If you are
not on the channel, it must not have cmode +s set
and users with umode +i are not shown.

If the parameter is not a nickname or a channel, users
with matching nickname, username, host, server or
GECOS information are shown.  The wildcards * and ?
can be used.  Users with umode +i set that are not
on the same channel as you are not shown.

A second parameter of a lowercase letter o ensures
only IRC operators are displayed.

The second parameter may also contain a format
specification starting with a percent sign.
This causes the output to use numeric 354,
with the selected fields:

t       -       Query type. Outputs the given number in each reply.
c       -       Channel.
u       -       Username.
i       -       IP address.
h       -       Host.
s       -       Server.
n       -       Nickname.
f       -       Status.
d       -       Hop count.
l       -       Idle time or 0 for users on other servers.
a       -       Services account name or 0 if none.
r       -       GECOS information.

"WHO #lamers %tuhnf,42" would generate a brief listing
of channel members and include the number 42 in each
line.

See also: whois, userhost, cmode, umode
```

### WHOIS

```
WHOIS [nick] nick

WHOIS will display detailed user information for
the specified nick.  If the first parameter is
specified, WHOIS will display information from
the specified server, or the server that the
user is on.  This is how to remotely see
idle time and away status.
```

### WHOWAS

```
WHOWAS <nick>

WHOWAS will show you the last known host and whois
information for the specified nick.  Depending on the
number of times they have connected to the network, there
may be more than one listing for a specific user.

The WHOWAS data will expire after time.
```

### XLINE

```
XLINE [time] <gecos> :<reason>

Bans by gecos (aka 'real name') field.  If [time] is not specified
this is added to the database, otherwise is temporary for [time]
minutes.

Eg. /quote xline eggdrop?bot :no bots

The <gecos> field contains certain special characters:
  ? - Match any single character
  * - Many any characters
  @ - Match any letter [A-Za-z]
  # - Match any digit [0-9]

To use a literal one of these characters, escape it with '\'.  A
literal '\' character must also be escaped.  You may also insert \s
which will be converted into a space.

If the <gecos> field is purely numeric (ie "123") then the time
field must be specified.  "0" must be used to denote a permanent
numeric XLINE.

XLINE [time] <gecos> ON <server> :<reason>

Will attempt to set the XLINE on <server> if <server> accepts
remote xlines.

- Requires Oper Priv: oper:xline
```
