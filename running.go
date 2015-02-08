// Copyright (c) 2015 Joshua Scoggins
//
// This software is provided 'as-is', without any express or implied
// warranty. In no event will the authors be held liable for any damages
// arising from the use of this software.
//
// Permission is granted to anyone to use this software for any purpose,
// including commercial applications, and to alter it and redistribute it
// freely, subject to the following restrictions:
//
// 1. The origin of this software must not be misrepresented; you must not
//    claim that you wrote the original software. If you use this software
//    in a product, an acknowledgement in the product documentation would be
//    appreciated but is not required.
// 2. Altered source versions must be plainly marked as such, and must not be
//    misrepresented as being the original software.
// 3. This notice may not be removed or altered from any source distribution.
//

//basic functionalty for terminating programs that are meant to run forever

package neuron

import (
	"os"
	"os/signal"
	"syscall"
)

var running = true

func IsRunning() bool {
	return running
}

func StopRunning() {
	running = false
}

func StartRunning() {
	running = true
}

func StopRunningOnSignal(sig syscall.Signal) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, sig)
	go func() {
		<-signalChan
		running = false
	}()
}

func StartRunningOnSignal(sig syscall.Signal) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, sig)
	go func() {
		<-signalChan
		running = true
	}()
}
