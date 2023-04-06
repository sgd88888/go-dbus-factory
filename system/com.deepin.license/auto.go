// Code generated by "./generator ./system/com.deepin.license"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package license

import "errors"
import "fmt"
import "github.com/godbus/dbus/v5"

import "github.com/linuxdeepin/go-lib/dbusutil"
import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "unsafe"

type License interface {
	license // interface com.deepin.license.Info
	proxy.Object
}

type objectLicense struct {
	interfaceLicense // interface com.deepin.license.Info
	proxy.ImplObject
}

func NewLicense(conn *dbus.Conn) License {
	obj := new(objectLicense)
	obj.ImplObject.Init_(conn, "com.deepin.license", "/com/deepin/license/Info")
	return obj
}

type license interface {
	ConnectLicenseStateChange(cb func(operationMseeage uint32)) (dbusutil.SignalHandlerId, error)
	AuthorizationState() proxy.PropInt32
	LicenseState() proxy.PropInt32
	ActiveMethod() proxy.PropUint32
	ActiveCode() proxy.PropString
	AuthorizeUrl() proxy.PropString
	AuthorizeVersion() proxy.PropString
	ValidityEnd() proxy.PropString
	Organization() proxy.PropString
}

type interfaceLicense struct{}

func (v *interfaceLicense) GetObject_() *proxy.ImplObject {
	return (*proxy.ImplObject)(unsafe.Pointer(v))
}

func (*interfaceLicense) GetInterfaceName_() string {
	return "com.deepin.license.Info"
}

// signal LicenseStateChange

func (v *interfaceLicense) ConnectLicenseStateChange(cb func(operationMseeage uint32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LicenseStateChange", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LicenseStateChange",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var operationMseeage uint32
		err := dbus.Store(sig.Body, &operationMseeage)
		if err == nil {
			cb(operationMseeage)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property AuthorizationState i

func (v *interfaceLicense) AuthorizationState() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "AuthorizationState",
	}
}

// property LicenseState i

func (v *interfaceLicense) LicenseState() proxy.PropInt32 {
	return &proxy.ImplPropInt32{
		Impl: v,
		Name: "LicenseState",
	}
}

// property ActiveMethod u

func (v *interfaceLicense) ActiveMethod() proxy.PropUint32 {
	return &proxy.ImplPropUint32{
		Impl: v,
		Name: "ActiveMethod",
	}
}

// property ActiveCode s

func (v *interfaceLicense) ActiveCode() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ActiveCode",
	}
}

// property AuthorizeUrl s

func (v *interfaceLicense) AuthorizeUrl() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "AuthorizeUrl",
	}
}

// property AuthorizeVersion s

func (v *interfaceLicense) AuthorizeVersion() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "AuthorizeVersion",
	}
}

// property ValidityEnd s

func (v *interfaceLicense) ValidityEnd() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "ValidityEnd",
	}
}

// property Organization s

func (v *interfaceLicense) Organization() proxy.PropString {
	return &proxy.ImplPropString{
		Impl: v,
		Name: "Organization",
	}
}
