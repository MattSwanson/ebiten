// Copyright 2016 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build android ios

package ebitenmobileview

import (
	"github.com/MattSwanson/ebiten/v2/internal/driver"
	"github.com/MattSwanson/ebiten/v2/internal/uidriver/mobile"
)

type position struct {
	x int
	y int
}

var (
	keys     = map[driver.Key]struct{}{}
	runes    []rune
	touches  = map[driver.TouchID]position{}
	gamepads = map[driver.GamepadID]*mobile.Gamepad{}
)

func updateInput() {
	ts := make([]*mobile.Touch, 0, len(touches))
	for id, position := range touches {
		ts = append(ts, &mobile.Touch{
			ID: id,
			X:  position.x,
			Y:  position.y,
		})
	}

	gs := make([]mobile.Gamepad, 0, len(gamepads))
	for _, g := range gamepads {
		gs = append(gs, *g)
	}

	mobile.Get().UpdateInput(keys, runes, ts, gs)
}
