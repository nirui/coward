//  Crypto-Obscured Forwarder
//
//  Copyright (C) 2017 Rui NI <ranqus@gmail.com>
//
//  This file is part of Crypto-Obscured Forwarder.
//
//  Crypto-Obscured Forwarder is free software: you can redistribute it
//  and/or modify it under the terms of the GNU General Public License
//  as published by the Free Software Foundation, either version 3 of
//  the License, or (at your option) any later version.
//
//  Crypto-Obscured Forwarder is distributed in the hope that it will be
//  useful, but WITHOUT ANY WARRANTY; without even the implied warranty
//  of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with Crypto-Obscured Forwarder. If not, see
//  <http://www.gnu.org/licenses/>.

package network

import (
	"net"
	"time"
)

// ConnectionID is the UID of current connection
type ConnectionID string

// Connection is a network connection
type Connection interface {
	net.Conn

	ID() ConnectionID
	SetTimeout(timeout time.Duration)
	SetReadTimeout(timeout time.Duration)
	SetWriteTimeout(timeout time.Duration)
	Closed() <-chan struct{}
}

// ConnectionWrapper wraps a connection
type ConnectionWrapper func(net.Conn) Connection

// Listener is a network listenr
type Listener interface {
	Listen() (Acceptor, error)
	String() string
}

// Acceptor network acceptor
type Acceptor interface {
	Addr() net.Addr
	Accept() (Connection, error)
	Close() error
}

// Dial Dialer builder
type Dial interface {
	Dial() (Connection, error)
	String() string
}

// Dialer dials to the remote host
type Dialer interface {
	Dialer() Dial
}
