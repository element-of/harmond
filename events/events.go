package events

import irc "gopkg.in/irc.v1"

//go:generate stringer -type=Type

// Type is the event type.
type Type int

// The entire list of events.
const (
	EvNone Type = iota
	EvAcceptAdd
	EvAcceptRemove
	EvAcceptList

	EvAdmin

	EvAwaySetReason
	EvAwayClearReason

	EvChghost

	EvDlineAdd
	EvDlineTest
	EvDlineRemove

	EvHelp
	EvInvite
	EvJoin
	EvKick
	EvKill

	EvKlineAdd
	EvKlineTest
	EvKlineRemove

	EvKnock
	EvList

	EvChannelModeAdd
	EvChannelModeQuery
	EvChannelModeRemove
	EvChannelListAdd
	EvChannelListQuery
	EvChannelListRemove

	EvMonitorAdd
	EvMonitorRemove
	EvMonitorClear
	EvMonitorQuery
	EvMonitorStatus

	EvMotd
	EvNames

	EvNick

	EvNoticeChannel
	EvNoticeUser

	EvPartChannel
	EvPass
	EvPing

	EvMessageChannel
	EvMessageUser

	EvQuit

	EvResvChannelAdd
	EvResvChannelQuery
	EvResvChannelRemove

	EvResvNickAdd
	EvResvNickQuery
	EvResvNickRemove

	EvTopic
	EvUser
	EvUsers
	EvVersion

	EvWhoChannel
	EvWhoUser

	EvWhois
	EvWhowas

	EvGecosBanAdd
	EvGecosBanQuery
	EvGecosBanRemove
)

// Event is an encapsulated event recieved from a Client.
type Event struct {
	Type  Type        `json:"type"`
	Event interface{} `json:"event"`
	Tags  irc.Tags    `json:"tags"`
}
