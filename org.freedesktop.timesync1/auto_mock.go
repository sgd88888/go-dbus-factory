// Code generated by "./generator org.freedesktop.timesync1"; DO NOT EDIT.

package timesync1

import "fmt"

import "github.com/stretchr/testify/mock"

import "pkg.deepin.io/lib/dbusutil/proxy"

type MockTimesync1 struct {
	mockInterfaceTimesync1 // interface org.freedesktop.timesync1.Manager
}

type mockInterfaceTimesync1 struct {
	mock.Mock
}

// property LinkNTPServers as

func (v *mockInterfaceTimesync1) LinkNTPServers() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property SystemNTPServers as

func (v *mockInterfaceTimesync1) SystemNTPServers() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property FallbackNTPServers as

func (v *mockInterfaceTimesync1) FallbackNTPServers() proxy.PropStringArray {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropStringArray)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property ServerName s

func (v *mockInterfaceTimesync1) ServerName() proxy.PropString {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropString)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property RootDistanceMaxUSec t

func (v *mockInterfaceTimesync1) RootDistanceMaxUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PollIntervalMinUSec t

func (v *mockInterfaceTimesync1) PollIntervalMinUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PollIntervalMaxUSec t

func (v *mockInterfaceTimesync1) PollIntervalMaxUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property PollIntervalUSec t

func (v *mockInterfaceTimesync1) PollIntervalUSec() proxy.PropUint64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropUint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}

// property Frequency x

func (v *mockInterfaceTimesync1) Frequency() proxy.PropInt64 {
	mockArgs := v.Called()

	ret0, ok := mockArgs.Get(0).(*proxy.MockPropInt64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0
}
