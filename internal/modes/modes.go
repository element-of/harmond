/*
Package modes contains a bunch of constants and lookup tables that are
a pain to use without this file.

This package only has modes defined for Elemental-IRCd as that is the
daemon this is intended to link against. If you are forking this
daemon and want to support linking against other upstream irc daemons,
this is probably the file you want to change. Changing this file
without good reason in any major deployment may make support difficult
or impossible.
*/
package modes

// Converted from a python file

// PropFlag is a normal channel mode flag.
type PropFlag int

// SetFlag is an integer channel setting flag for parametric modes.
type SetFlag int

// PrefixFlag is an integer channel mode flag that cause a user prefix.
type PrefixFlag int

// UserFlag is a user mode flag
type UserFlag int

// ListMode is a channel list-like mode
type ListMode int

//go:generate stringer -type=PropFlag
//go:generate stringer -type=SetFlag
//go:generate stringer -type=PrefixFlag
//go:generate stringer -type=UserFlag
//go:generate stringer -type=ListMode

// Kinds of modes
const (
	ModeList       = iota // List-like mode
	ModeKey               // Key-like mode
	ModeParametric        // Paramteric channel mode
	ModeProp              // Channel property
	ModePrefix            // Channel prefix mode
)

// Channel access levels
const (
	ChflPeon   PrefixFlag = 1 << iota // No channel status
	ChflVoice                         // Voiced
	ChflHalfop                        // Channel halfop
	ChflChanop                        // Channel operator
	ChflAdmin                         // Channel admin
	ChflOwner                         // Channel owner
)

// Channel set flags for parametric modes
const (
	SetNone         SetFlag = 1 << iota // No settings
	SetKey                              // Old +k, channel is keyed
	SetForward                          // Old +f, channel has a forward channel
	SetLimit                            // Old +l, channel has a limit
	SetJoinThrottle                     // Old +j, channel has a join throttle
)

// Channel propreties
const (
	PropNone       PropFlag = 1 << iota // No properties
	PropMute                            // Old +m, mute
	PropPrivate                         // Old +p, private channel
	PropInvite                          // Old +i, invite only
	PropTopicrest                       // Old +t, only ops can set topic
	PropInternal                        // Old +n, only users in channel can send to it
	PropSecret                          // Old +s, only users in channel know it exists
	PropNoCTCP                          // Old +C, no CTCP messages
	PropNoAction                        // Old +D, no CTCP Action messages
	PropNoKicks                         // Old +E, operators cannot kick
	PropNoCaps                          // OLD +G, ALL CAPITAL LETTER MESSAGES ARE BLOCKED
	PropNoRejoin                        // Old +J, no immediate Rejoin after KICK
	PropLargelist                       // Old +L, larger channel lists
	PropNoOperKick                      // Old +M, staff cannot be kicked
	PropOperonly                        // Old +O, only opers may join
	PropPermanent                       // Old +P, channel persists without users
	PropDisforward                      // Old +Q, channel may not be forwarded to
	PropNoNotice                        // Old +T, channel may not be Notice'd to
	PropNoColor                         // Old +c, channel Color codes are stripped
	PropNoNicks                         // Old +d, nick changes are forbidden when in channel
	PropFreeinvite                      // Old +g, invite is freely usable
	PropHidebans                        // Old +u, ban list is hidden without the proper STATUS
	PropOpmod                           // Old +z, channel messages blocked by something are sent to ops
	PropFreefwd                         // Old +F, free forwarding
	PropNoRepeat                        // Old +K, no repeating messages
)

// User properties
const (
	UPropNone           UserFlag = 1 << iota // No user properties
	UPropInvisible                           // Old +i, invisible client
	UPropCallerid                            // Old +g, "caller id"
	UPropIrcop                               // Old +o, user is an IRC operator
	UPropCloaked                             // Old +x, user has a cloaked IP address
	UPropAdmin                               // Old +a, user is an IRC administrator
	UPropOverride                            // Old +p, implicit chanop access
	UPropNoCTCP                              // Old +C, prevents receiving CTCP messages other than Action (/me)
	UPropDeaf                                // Old +D, ignoes all channel messages
	UPropDisforward                          // Old +Q, prevents channel forwarding
	UPropRegpm                               // Old +R, requires people to be registered with services to pm
	UPropSoftcall                            // Old +G, Soft caller ID, caller id exempting common channels
	UPropNoinvite                            // Old +V, prevents user from getting invites
	UPropNostalk                             // Old +I, doesn't show channel list in whois
	UPropSSLClient                           // Old +Z, client is connected over SSL
	UPropNetworkService                      // Old +S, client is a network service with all of the associated powers
)

//Channel lists
const (
	ListNone   ListMode = iota
	ListBan             // Old +b, channel bans
	ListQuiet           // Old +q, channel quiets
	ListExcept          // Old +e, channel ban exceptions
	ListInvex           // Old +I, channel invite execptions
)

// This is a handy lookup table from channel mode letters to bitmasks.
var Chanmodes = []map[string]interface{}{
	map[string]interface{}{
		"q": ListQuiet,
		"b": ListBan,
		"e": ListExcept,
		"I": ListInvex,
	},
	map[string]interface{}{
		"k": SetKey,
	},
	map[string]interface{}{
		"j": SetJoinThrottle,
		"f": SetJoinThrottle,
		"l": SetLimit,
	},
	map[string]interface{}{
		"C": PropNoCTCP,
		"D": PropNoAction,
		"E": PropNoKicks,
		"G": PropNoCaps,
		"J": PropNoRejoin,
		"F": PropFreefwd,
		"K": PropNoRepeat,
		"L": PropLargelist,
		"M": PropNoOperKick,
		"O": PropOperonly,
		"P": PropPermanent,
		"Q": PropDisforward,
		"T": PropNoNotice,
		"c": PropNoColor,
		"d": PropNoNicks,
		"g": PropFreeinvite,
		"i": PropInvite,
		"m": PropMute,
		"n": PropInternal,
		"p": PropPrivate,
		"s": PropSecret,
		"t": PropTopicrest,
		"u": PropHidebans,
		"z": PropOpmod,
	},
	map[string]interface{}{
		"h": ChflHalfop,
		"o": ChflChanop,
		"v": ChflVoice,
		"a": ChflAdmin,
		"y": ChflOwner,
	},
}

// This is a handy lookup table from bitmask flags to mode letters.
var Revchanmodes = []map[interface{}]string{
	map[interface{}]string{
		ListQuiet:  "q",
		ListBan:    "b",
		ListExcept: "e",
		ListInvex:  "I",
	},
	map[interface{}]string{
		SetKey: "k",
	},
	map[interface{}]string{
		SetJoinThrottle: "j",
		SetForward:      "f",
		SetLimit:        "l",
	},
	map[interface{}]string{
		PropNoCTCP:     "C",
		PropNoAction:   "D",
		PropNoKicks:    "E",
		PropNoCaps:     "G",
		PropNoRejoin:   "J",
		PropFreefwd:    "F",
		PropNoRepeat:   "K",
		PropLargelist:  "L",
		PropNoOperKick: "M",
		PropOperonly:   "O",
		PropPermanent:  "P",
		PropDisforward: "Q",
		PropNoNotice:   "T",
		PropNoColor:    "c",
		PropNoNicks:    "d",
		PropFreeinvite: "g",
		PropInvite:     "i",
		PropMute:       "m",
		PropInternal:   "n",
		PropPrivate:    "p",
		PropSecret:     "s",
		PropTopicrest:  "t",
		PropHidebans:   "u",
		PropOpmod:      "z",
	},
	map[interface{}]string{
		ChflHalfop: "h",
		ChflChanop: "o",
		ChflVoice:  "v",
		ChflAdmin:  "a",
		ChflOwner:  "y",
	},
}

// This is a handy lookup table from user mode flags to bitmasks.
var Umodes = map[string]interface{}{
	"i": UPropInvisible,
	"g": UPropCallerid,
	"o": UPropIrcop,
	"a": UPropAdmin,
	"p": UPropOverride,
	"C": UPropNoCTCP,
	"D": UPropDeaf,
	"Q": UPropDisforward,
	"R": UPropRegpm,
	"G": UPropSoftcall,
	"V": UPropNoinvite,
	"I": UPropNostalk,
	"Z": UPropSSLClient,
	"S": UPropNetworkService,
}

// This is a handy lookup table for channel prefixes to bitmask flags.
var Prefixes = map[string]interface{}{
	"+": ChflVoice,
	"%": ChflHalfop,
	"@": ChflChanop,
	"!": ChflAdmin,
	"~": ChflOwner,
}

// This is a handy lookup table for status mode flags to human readable strings.
var Statuses = map[PrefixFlag]string{
	ChflVoice:  "voiced",
	ChflHalfop: "half-operator",
	ChflChanop: "channel operator",
	ChflAdmin:  "channel administrator",
	ChflOwner:  "channel owner",
}
