//go:build darwin
// +build darwin

package dns

import "net"

// SessionUDP holds the remote address.
type SessionUDP struct {
	raddr *net.UDPAddr
}

// RemoteAddr returns the remote network address.
func (s *SessionUDP) RemoteAddr() net.Addr { return s.raddr }

// ReadFromSessionUDP acts just like net.UDPConn.ReadFrom(), but returns a
// session struct instead of a net.UDPAddr.
func ReadFromSessionUDP(conn *net.UDPConn, b []byte) (int, *SessionUDP, error) {
	n, _, _, raddr, err := conn.ReadMsgUDP(b, nil)
	if err != nil {
		return n, nil, err
	}
	return n, &SessionUDP{raddr}, err
}

// WriteToSessionUDP acts just like net.UDPConn.WriteTo(), but uses a
// session struct instead of a net.Addr.
func WriteToSessionUDP(conn *net.UDPConn, b []byte, session *SessionUDP) (int, error) {
	n, _, err := conn.WriteMsgUDP(b, nil, session.raddr)
	return n, err
}

func setUDPSocketOptions(*net.UDPConn) error { return nil }

//lint:ignore U1000 Ignore unused function.
func parseDstFromOOB([]byte, net.IP) net.IP { return nil }
