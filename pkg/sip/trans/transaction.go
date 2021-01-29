package trans

import (
	"time"

	"github.com/jart/gosip/sip"
)

type Transactioner interface {
	Do(*Event) error
	IsTerminated() bool
}

type State int

const (
	/* STATES for invite client trans */
	ICT_PRE_CALLING State = iota
	ICT_CALLING
	ICT_PROCEEDING
	ICT_COMPLETED
	ICT_TERMINATED

	/* STATES for invite server trans */
	IST_PRE_PROCEEDING
	IST_PROCEEDING
	IST_COMPLETED
	IST_CONFIRMED
	IST_TERMINATED

	/* STATES for NON-invite client trans */
	NICT_PRE_TRYING
	NICT_TRYING
	NICT_PROCEEDING
	NICT_COMPLETED
	NICT_TERMINATED

	/* STATES for NON-invite server trans */
	NIST_PRE_TRYING
	NIST_TRYING
	NIST_PROCEEDING
	NIST_COMPLETED
	NIST_TERMINATED

	UNKNOWN_STATE /**< Max state*/
	/* STATES for dialog */
//	DIALOG_EARLY
//	DIALOG_CONFIRMED
//	DIALOG_CLOSE
)

type Type int

const (
	ICT     Type = iota /**< Invite Client (outgoing) Transaction */
	IST                 /**< Invite Server (incoming) Transaction */
	NICT                /**< Non-Invite Client (outgoing) Transaction */
	NIST                /**< Non-Invite Server (incoming) Transaction */
	UNKNOWN             /**< Invalid Transaction */
)

type EventType int

const (
	/* TIMEOUT EVENTS for ICT */
	TIMEOUT_A EventType = iota /**< Timer A */
	TIMEOUT_B                  /**< Timer B */
	TIMEOUT_D                  /**< Timer D */

	/* TIMEOUT EVENTS for NICT */
	TIMEOUT_E /**< Timer E */
	TIMEOUT_F /**< Timer F */
	TIMEOUT_K /**< Timer K */

	/* TIMEOUT EVENTS for IST */
	TIMEOUT_G /**< Timer G */
	TIMEOUT_H /**< Timer H */
	TIMEOUT_I /**< Timer I */

	/* TIMEOUT EVENTS for NIST */
	TIMEOUT_J /**< Timer J */

	/* FOR INCOMING MESSAGE */
	RCV_REQINVITE     /**< Event is an incoming INVITE request */
	RCV_REQACK        /**< Event is an incoming ACK request */
	RCV_REQUEST       /**< Event is an incoming NON-INVITE and NON-ACK request */
	RCV_STATUS_1XX    /**< Event is an incoming informational response */
	RCV_STATUS_2XX    /**< Event is an incoming 2XX response */
	RCV_STATUS_3456XX /**< Event is an incoming final response (not 2XX) */

	/* FOR OUTGOING MESSAGE */
	SND_REQINVITE     /**< Event is an outgoing INVITE request */
	SND_REQACK        /**< Event is an outgoing ACK request */
	SND_REQUEST       /**< Event is an outgoing NON-INVITE and NON-ACK request */
	SND_STATUS_1XX    /**< Event is an outgoing informational response */
	SND_STATUS_2XX    /**< Event is an outgoing 2XX response */
	SND_STATUS_3456XX /**< Event is an outgoing final response (not 2XX) */

	//	KILL_TRANSACTION /**< Event to 'kill' the trans before termination */
	UNKNOWN_EVT /**< Max event */
)

type Event struct {
	Type EventType
	Tid  string // trans id
	Msg  *sip.Msg
}

func SetIncomingEventType(msg *sip.Msg) EventType {
	if msg.IsResponse() {
		switch {
		case msg.Status >= 100 && msg.Status < 200:
			return RCV_STATUS_1XX
		case msg.Status >= 200 && msg.Status < 300:
			return RCV_STATUS_2XX
		default:
			return RCV_STATUS_3456XX
		}
	}

	switch msg.Method {
	case sip.MethodInvite:
		return RCV_REQINVITE
	case sip.MethodAck:
		return RCV_REQACK
	default:
		return RCV_REQUEST
	}
}

func SetOutcomingEventType(msg *sip.Msg) EventType {
	if msg.IsResponse() {
		switch {
		case msg.Status >= 100 && msg.Status < 200:
			return SND_STATUS_1XX
		case msg.Status >= 200 && msg.Status < 300:
			return SND_STATUS_2XX
		default:
			return SND_STATUS_3456XX
		}
	}

	switch msg.Method {
	case sip.MethodInvite:
		return SND_REQINVITE
	case sip.MethodAck:
		return SND_REQACK
	default:
		return SND_REQUEST
	}
}

// timer相关基础常量、方法等定义
const (
	T1      = 100 * time.Millisecond
	T2      = 4 * time.Second
	T4      = 5 * time.Second
	TimeA   = T1
	TimeB   = 64 * T1
	TimeD   = 32 * time.Second
	TimeE   = T1
	TimeF   = 64 * T1
	TimeG   = T1
	TimeH   = 64 * T1
	TimeI   = T4
	TimeJ   = 64 * T1
	TimeK   = T4
	Time1xx = 100 * time.Millisecond
)
