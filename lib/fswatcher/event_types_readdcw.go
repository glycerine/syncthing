// Copyright (C) 2017 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// +build windows

package fswatcher

import "github.com/zillode/notify"

func (watcher *fsWatcher) eventMask() notify.Event {
	events := notify.FileNotifyChangeFileName | notify.FileNotifyChangeDirName | notify.FileNotifyChangeSize | notify.FileNotifyChangeCreation
	if !watcher.ignorePerms {
		events |= notify.FileNotifyChangeAttributes
	}
	return events
}

const removeEventMask = notify.FileActionRemoved | notify.FileActionRenamedOldName