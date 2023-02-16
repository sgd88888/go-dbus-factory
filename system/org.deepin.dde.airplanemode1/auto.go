// Code generated by "./generator ./system/org.deepin.dde.airplanemode1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package airplanemode1

import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type AirplaneMode interface {
	airplaneMode // interface org.deepin.dde.AirplaneMode1
	proxy.Object
}

type objectAirplaneMode struct {
	interfaceAirplaneMode // interface org.deepin.dde.AirplaneMode1
	proxy.ImplObject
}

func NewAirplaneMode(conn *dbus.Conn) AirplaneMode {
	obj := new(objectAirplaneMode)
	obj.ImplObject.Init_(conn, "org.deepin.dde.AirplaneMode1", "/org/deepin/dde/AirplaneMode1")
	return obj
}

type airplaneMode interface {
	GoDumpState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call
	DumpState(flags dbus.Flags) error
	GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	Enable(flags dbus.Flags, enabled bool) error
	GoEnableBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	EnableBluetooth(flags dbus.Flags, enabled bool) error
	GoEnableWifi(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call
	EnableWifi(flags dbus.Flags, enabled bool) error
	Enabled() proxy.PropBool
	WifiEnabled() proxy.PropBool
	BluetoothEnabled() proxy.PropBool
}

type interfaceAirplaneMode struct{}

func (v *interfaceAirplaneMode) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceAirplaneMode) GetInterfaceName_() string {
	return "org.deepin.dde.AirplaneMode1"
}

// method DumpState

func (v *interfaceAirplaneMode) GoDumpState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpState", flags, ch)
}

func (v *interfaceAirplaneMode) DumpState(flags dbus.Flags) error {
	return (<-v.GoDumpState(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method Enable

func (v *interfaceAirplaneMode) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *interfaceAirplaneMode) Enable(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method EnableBluetooth

func (v *interfaceAirplaneMode) GoEnableBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableBluetooth", flags, ch, enabled)
}

func (v *interfaceAirplaneMode) EnableBluetooth(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableBluetooth(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method EnableWifi

func (v *interfaceAirplaneMode) GoEnableWifi(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWifi", flags, ch, enabled)
}

func (v *interfaceAirplaneMode) EnableWifi(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableWifi(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// property Enabled b

func (v *interfaceAirplaneMode) Enabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property WifiEnabled b

func (v *interfaceAirplaneMode) WifiEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "WifiEnabled",
	}
}

// property BluetoothEnabled b

func (v *interfaceAirplaneMode) BluetoothEnabled() proxy.PropBool {
	return &proxy.ImplPropBool{
		Impl: v,
		Name: "BluetoothEnabled",
	}
}