// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2024 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package prepare

import (
	"github.com/roddhjav/apparmor.d/pkg/prebuild"
)

type SystemdDefault struct {
	prebuild.Base
}

type SystemdEarly struct {
	prebuild.Base
}

func init() {
	RegisterTask(&SystemdDefault{
		Base: prebuild.Base{
			Keyword: "systemd-default",
			Msg:     "Configure systemd unit drop in files to a profile for some units",
		},
	})
	RegisterTask(&SystemdEarly{
		Base: prebuild.Base{
			Keyword: "systemd-early",
			Msg:     "Configure systemd unit drop in files to ensure some service start after apparmor",
		},
	})
}

func (p SystemdDefault) Apply() ([]string, error) {
	return []string{}, prebuild.SystemdDir.Join("default").CopyFS(prebuild.Root.Join("systemd"))
}

func (p SystemdEarly) Apply() ([]string, error) {
	return []string{}, prebuild.SystemdDir.Join("early").CopyFS(prebuild.Root.Join("systemd"))
}
