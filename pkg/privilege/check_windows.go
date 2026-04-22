//go:build windows

package privilege

import "golang.org/x/sys/windows"

func IsElevated() bool {
	token := windows.Token(0)

	adminSID, err := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid)
	if err != nil {
		return false
	}

	member, err := token.IsMember(adminSID)
	return err == nil && member
}
