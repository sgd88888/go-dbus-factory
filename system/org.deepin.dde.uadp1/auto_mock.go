// Code generated by "./generator ./system/org.deepin.dde.uadp1"; DO NOT EDIT.

// SPDX-FileCopyrightText: 2018 - 2023 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later
package uadp1

import "fmt"
import "github.com/godbus/dbus"

import "github.com/linuxdeepin/go-lib/dbusutil/proxy"
import "github.com/stretchr/testify/mock"

type MockUadp struct {
	MockInterfaceUadp // interface org.deepin.dde.Uadp1
	proxy.MockObject
}

type MockInterfaceUadp struct {
	mock.Mock
}

// method SetDataKey

func (v *MockInterfaceUadp) GoSetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, dataKey string, keyringKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, exePath, keyName, dataKey, keyringKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) SetDataKey(flags dbus.Flags, exePath string, keyName string, dataKey string, keyringKey string) error {
	mockArgs := v.Called(flags, exePath, keyName, dataKey, keyringKey)

	return mockArgs.Error(0)
}

// method GetDataKey

func (v *MockInterfaceUadp) GoGetDataKey(flags dbus.Flags, ch chan *dbus.Call, exePath string, keyName string, keyringKey string) *dbus.Call {
	mockArgs := v.Called(flags, ch, exePath, keyName, keyringKey)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) GetDataKey(flags dbus.Flags, exePath string, keyName string, keyringKey string) (string, error) {
	mockArgs := v.Called(flags, exePath, keyName, keyringKey)

	return mockArgs.String(0), mockArgs.Error(1)
}

// method Available

func (v *MockInterfaceUadp) GoAvailable(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) Available(flags dbus.Flags) (bool, error) {
	mockArgs := v.Called(flags)

	return mockArgs.Bool(0), mockArgs.Error(1)
}

// method Delete

func (v *MockInterfaceUadp) GoDelete(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) Delete(flags dbus.Flags, name string) error {
	mockArgs := v.Called(flags, name)

	return mockArgs.Error(0)
}

// method Get

func (v *MockInterfaceUadp) GoGet(flags dbus.Flags, ch chan *dbus.Call, name string) *dbus.Call {
	mockArgs := v.Called(flags, ch, name)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) Get(flags dbus.Flags, name string) ([]uint8, error) {
	mockArgs := v.Called(flags, name)

	ret0, ok := mockArgs.Get(0).([]uint8)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method ListName

func (v *MockInterfaceUadp) GoListName(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) ListName(flags dbus.Flags) ([]string, error) {
	mockArgs := v.Called(flags)

	ret0, ok := mockArgs.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}

// method Release

func (v *MockInterfaceUadp) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	mockArgs := v.Called(flags, ch)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) Release(flags dbus.Flags) error {
	mockArgs := v.Called(flags)

	return mockArgs.Error(0)
}

// method Set

func (v *MockInterfaceUadp) GoSet(flags dbus.Flags, ch chan *dbus.Call, name string, data []uint8) *dbus.Call {
	mockArgs := v.Called(flags, ch, name, data)

	ret, ok := mockArgs.Get(0).(*dbus.Call)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: 0 failed because object wasn't correct type: %v", mockArgs.Get(0)))
	}

	return ret
}

func (v *MockInterfaceUadp) Set(flags dbus.Flags, name string, data []uint8) error {
	mockArgs := v.Called(flags, name, data)

	return mockArgs.Error(0)
}