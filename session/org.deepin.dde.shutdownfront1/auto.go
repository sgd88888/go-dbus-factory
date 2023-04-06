// Code generated by "./generator ./session/org.deepin.dde.shutdownfront1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package shutdownfront1

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type ShutdownFront interface {
	shutdownfront // interface org.deepin.dde.ShutdownFront1
	proxy.Object
}

type objectShutdownFront struct {
	interfaceShutdownfront // interface org.deepin.dde.ShutdownFront1
	proxy.ImplObject
}

func NewShutdownFront(conn *dbus.Conn) ShutdownFront {
	obj := new(objectShutdownFront)
	obj.ImplObject.Init_(conn, "org.deepin.dde.ShutdownFront1", "/org/deepin/dde/ShutdownFront1")
	return obj
}

type shutdownfront interface {
	GoHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Hibernate(flags dbus.Flags) error
	GoLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Lock(flags dbus.Flags) error
	GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Logout(flags dbus.Flags) error
	GoRestart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Restart(flags dbus.Flags) error
	GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Show(flags dbus.Flags) error
	GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Shutdown(flags dbus.Flags) error
	GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	Suspend(flags dbus.Flags) error
	GoSwitchUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	SwitchUser(flags dbus.Flags) error
	ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error)
}

type interfaceShutdownfront struct{}

func (v *interfaceShutdownfront) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceShutdownfront) GetInterfaceName_() string {
	return "org.deepin.dde.ShutdownFront1"
}

// method Hibernate

func (v *interfaceShutdownfront) GoHibernate(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Hibernate", flags, ch)
}

func (v *interfaceShutdownfront) Hibernate(flags dbus.Flags) error {
	return (<-v.GoHibernate(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Lock

func (v *interfaceShutdownfront) GoLock(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Lock", flags, ch)
}

func (v *interfaceShutdownfront) Lock(flags dbus.Flags) error {
	return (<-v.GoLock(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Logout

func (v *interfaceShutdownfront) GoLogout(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Logout", flags, ch)
}

func (v *interfaceShutdownfront) Logout(flags dbus.Flags) error {
	return (<-v.GoLogout(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Restart

func (v *interfaceShutdownfront) GoRestart(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Restart", flags, ch)
}

func (v *interfaceShutdownfront) Restart(flags dbus.Flags) error {
	return (<-v.GoRestart(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Show

func (v *interfaceShutdownfront) GoShow(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Show", flags, ch)
}

func (v *interfaceShutdownfront) Show(flags dbus.Flags) error {
	return (<-v.GoShow(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Shutdown

func (v *interfaceShutdownfront) GoShutdown(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Shutdown", flags, ch)
}

func (v *interfaceShutdownfront) Shutdown(flags dbus.Flags) error {
	return (<-v.GoShutdown(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Suspend

func (v *interfaceShutdownfront) GoSuspend(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Suspend", flags, ch)
}

func (v *interfaceShutdownfront) Suspend(flags dbus.Flags) error {
	return (<-v.GoSuspend(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SwitchUser

func (v *interfaceShutdownfront) GoSwitchUser(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SwitchUser", flags, ch)
}

func (v *interfaceShutdownfront) SwitchUser(flags dbus.Flags) error {
	return (<-v.GoSwitchUser(flags, make(chan *dbus.Call, 1)).Done).Err
}

// signal ChangKey

func (v *interfaceShutdownfront) ConnectChangKey(cb func(keyEvent string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ChangKey", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ChangKey",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keyEvent string
		err := dbus.Store(sig.Body, &keyEvent)
		if err == nil {
			cb(keyEvent)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
