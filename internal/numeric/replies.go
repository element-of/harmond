package numeric

// The list of all numerics the ircd uses.
//
// This is scraped from the ircd source code directly.

type Response string

const (
	RplWelcome            Response = "001"
	RplYourhost           Response = "002"
	RplCreated            Response = "003"
	RplMyinfo             Response = "004"
	RplIsupport           Response = "005"
	RplSnomask            Response = "008"
	RplRedir              Response = "010"
	RplMap                Response = "015"
	RplMapend             Response = "017"
	RplSavenick           Response = "043"
	RplTracelink          Response = "200"
	RplTraceconnecting    Response = "201"
	RplTracehandshake     Response = "202"
	RplTraceunknown       Response = "203"
	RplTraceoperator      Response = "204"
	RplTraceuser          Response = "205"
	RplTraceserver        Response = "206"
	RplTracenewtype       Response = "208"
	RplTraceclass         Response = "209"
	RplStatscommands      Response = "212"
	RplStatscline         Response = "213"
	RplStatsiline         Response = "215"
	RplStatskline         Response = "216"
	RplStatsqline         Response = "217"
	RplStatsyline         Response = "218"
	RplEndofstats         Response = "219"
	RplStatspline         Response = "220"
	RplUmodeis            Response = "221"
	RplStatsdline         Response = "225"
	RplStatslline         Response = "241"
	RplStatsuptime        Response = "242"
	RplStatsoline         Response = "243"
	RplStatshline         Response = "244"
	RplStatsxline         Response = "247"
	RplStatsuline         Response = "248"
	RplStatsconn          Response = "250"
	RplLuserclient        Response = "251"
	RplLuserop            Response = "252"
	RplLuserunknown       Response = "253"
	RplLuserchannels      Response = "254"
	RplLuserme            Response = "255"
	RplAdminme            Response = "256"
	RplAdminloc1          Response = "257"
	RplAdminloc2          Response = "258"
	RplAdminemail         Response = "259"
	RplEndoftrace         Response = "262"
	RplLoad2HI            Response = "263"
	RplLocalusers         Response = "265"
	RplGlobalusers        Response = "266"
	RplPrivs              Response = "270"
	RplWhoiscertfp        Response = "276"
	RplAcceptlist         Response = "281"
	RplEndofaccept        Response = "282"
	RplAway               Response = "301"
	RplUserhost           Response = "302"
	RplIson               Response = "303"
	RplUnaway             Response = "305"
	RplNowaway            Response = "306"
	RplWhoisuser          Response = "311"
	RplWhoisserver        Response = "312"
	RplWhoisoperator      Response = "313"
	RplWhowasuser         Response = "314"
	RplEndofwho           Response = "315"
	RplWhoisidle          Response = "317"
	RplEndofwhois         Response = "318"
	RplWhoischannels      Response = "319"
	RplListstart          Response = "321"
	RplList               Response = "322"
	RplListend            Response = "323"
	RplChannelmodeis      Response = "324"
	RplChannelmlockis     Response = "325"
	RplCreationtime       Response = "329"
	RplWhoisloggedin      Response = "330"
	RplNotopic            Response = "331"
	RplTopic              Response = "332"
	RplTopicwhotime       Response = "333"
	RplWhoisbot           Response = "335"
	RplWhoisactually      Response = "338"
	RplInviting           Response = "341"
	RplInvexlist          Response = "346"
	RplEndofinvexlist     Response = "347"
	RplExceptlist         Response = "348"
	RplEndofexceptlist    Response = "349"
	RplVersion            Response = "351"
	RplWhoreply           Response = "352"
	RplNamreply           Response = "353"
	RplWhowasreal         Response = "360"
	RplClosing            Response = "362"
	RplCloseend           Response = "363"
	RplLinks              Response = "364"
	RplEndoflinks         Response = "365"
	RplEndofnames         Response = "366"
	RplBanlist            Response = "367"
	RplEndofbanlist       Response = "368"
	RplEndofwhowas        Response = "369"
	RplInfo               Response = "371"
	RplMotd               Response = "372"
	RplEndofinfo          Response = "374"
	RplMotdstart          Response = "375"
	RplEndofmotd          Response = "376"
	RplWhoishost          Response = "378"
	RplWhoismodes         Response = "379"
	RplYoureoper          Response = "381"
	RplRehashing          Response = "382"
	RplRsachallenge       Response = "386"
	RplTime               Response = "391"
	ErrNosuchnick         Response = "401"
	ErrNosuchserver       Response = "402"
	ErrNosuchchannel      Response = "403"
	ErrCannotsendtochan   Response = "404"
	ErrToomanychannels    Response = "405"
	ErrWasnosuchnick      Response = "406"
	ErrToomanytargets     Response = "407"
	ErrNoorigin           Response = "409"
	ErrInvalidcapcmd      Response = "410"
	ErrNorecipient        Response = "411"
	ErrNotexttosend       Response = "412"
	ErrNotoplevel         Response = "413"
	ErrWildtoplevel       Response = "414"
	ErrToomanymatches     Response = "416"
	ErrUnknowncommand     Response = "421"
	ErrNomotd             Response = "422"
	ErrNonicknamegiven    Response = "431"
	ErrErroneusnickname   Response = "432"
	ErrNicknameinuse      Response = "433"
	ErrBannickchange      Response = "435"
	ErrNickcollision      Response = "436"
	ErrUnavailresource    Response = "437"
	ErrNicktoofast        Response = "438"
	ErrServicesdown       Response = "440"
	ErrUsernotinchannel   Response = "441"
	ErrNotonchannel       Response = "442"
	ErrUseronchannel      Response = "443"
	ErrNoinvite           Response = "447"
	ErrNonick             Response = "449"
	ErrNotregistered      Response = "451"
	ErrAcceptfull         Response = "456"
	ErrAcceptexist        Response = "457"
	ErrAcceptnot          Response = "458"
	ErrNeedmoreparams     Response = "461"
	ErrAlreadyregistred   Response = "462"
	ErrPasswdmismatch     Response = "464"
	ErrYourebannedcreep   Response = "465"
	ErrLinkchannel        Response = "470"
	ErrChannelisfull      Response = "471"
	ErrUnknownmode        Response = "472"
	ErrInviteonlychan     Response = "473"
	ErrBannedfromchan     Response = "474"
	ErrBadchannelkey      Response = "475"
	ErrNeedreggednick     Response = "477"
	ErrBanlistfull        Response = "478"
	ErrBadchanname        Response = "479"
	ErrThrottle           Response = "480"
	ErrNoprivileges       Response = "481"
	ErrChanprivsneeded    Response = "482"
	ErrCantkillserver     Response = "483"
	ErrIschanservice      Response = "484"
	ErrNononreg           Response = "486"
	ErrVoiceneeded        Response = "489"
	ErrNooperhost         Response = "491"
	ErrNoctcp             Response = "492"
	ErrOwnmode            Response = "494"
	ErrKicknorejoin       Response = "495"
	ErrUmodeunknownflag   Response = "501"
	ErrUsersdontmatch     Response = "502"
	ErrUsernotonserv      Response = "504"
	ErrWrongpong          Response = "513"
	ErrDisabled           Response = "517"
	ErrNokick             Response = "519"
	ErrHelpnotfound       Response = "524"
	RplWhoissecure        Response = "671"
	RplWhoiswebirc        Response = "672"
	RplModlist            Response = "702"
	RplEndofmodlist       Response = "703"
	RplHelpstart          Response = "704"
	RplHelptxt            Response = "705"
	RplEndofhelp          Response = "706"
	ErrTargchange         Response = "707"
	RplEtracefull         Response = "708"
	RplEtrace             Response = "709"
	RplKnock              Response = "710"
	RplKnockdlvr          Response = "711"
	ErrToomanyknock       Response = "712"
	ErrChanopen           Response = "713"
	ErrKnockonchan        Response = "714"
	ErrKnockdisabled      Response = "715"
	ErrTargumodeg         Response = "716"
	RplTargnotify         Response = "717"
	RplUmodegmsg          Response = "718"
	RplOmotdstart         Response = "720"
	RplOmotd              Response = "721"
	RplEndofomotd         Response = "722"
	ErrNoprivs            Response = "723"
	RplTestline           Response = "725"
	RplNotestline         Response = "726"
	RplTestmaskgecos      Response = "727"
	RplQuietlist          Response = "728"
	RplEndofquietlist     Response = "729"
	RplMononline          Response = "730"
	RplMonoffline         Response = "731"
	RplMonlist            Response = "732"
	RplEndofmonlist       Response = "733"
	ErrMonlistfull        Response = "734"
	ErrNocommonchan       Response = "737"
	RplRsachallenge2      Response = "740"
	RplEndofrsachallenge2 Response = "741"
	ErrMlockrestricted    Response = "742"
	RplScanmatched        Response = "750"
	RplScanumodes         Response = "751"
	RplLoggedin           Response = "900"
	RplLoggedout          Response = "901"
	ErrNicklocked         Response = "902"
	RplSaslsuccess        Response = "903"
	ErrSaslfail           Response = "904"
	ErrSasltoolong        Response = "905"
	ErrSaslaborted        Response = "906"
	ErrSaslalready        Response = "907"
)
