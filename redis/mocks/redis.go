/*
 * Copyright (c) 2016 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

// Automatically generated by MockGen. DO NOT EDIT!
// Source: interfaces/redis.go

package mocks

import (
	redis "github.com/go-redis/redis"
	gomock "github.com/golang/mock/gomock"
	time "time"
)

// Mock of RedisClient interface
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *_MockRedisClientRecorder
}

// Recorder for MockRedisClient (not exported)
type _MockRedisClientRecorder struct {
	mock *MockRedisClient
}

func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &_MockRedisClientRecorder{mock}
	return mock
}

func (_m *MockRedisClient) EXPECT() *_MockRedisClientRecorder {
	return _m.recorder
}

func (_m *MockRedisClient) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRedisClientRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockRedisClient) HGetAll(_param0 string) *redis.StringStringMapCmd {
	ret := _m.ctrl.Call(_m, "HGetAll", _param0)
	ret0, _ := ret[0].(*redis.StringStringMapCmd)
	return ret0
}

func (_mr *_MockRedisClientRecorder) HGetAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HGetAll", arg0)
}

func (_m *MockRedisClient) HMSet(_param0 string, _param1 map[string]interface{}) *redis.StatusCmd {
	ret := _m.ctrl.Call(_m, "HMSet", _param0, _param1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

func (_mr *_MockRedisClientRecorder) HMSet(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HMSet", arg0, arg1)
}

func (_m *MockRedisClient) Ping() *redis.StatusCmd {
	ret := _m.ctrl.Call(_m, "Ping")
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

func (_mr *_MockRedisClientRecorder) Ping() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ping")
}

func (_m *MockRedisClient) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	ret := _m.ctrl.Call(_m, "SetNX", key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

func (_mr *_MockRedisClientRecorder) SetNX(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetNX", arg0, arg1, arg2)
}

func (_m *MockRedisClient) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	_s := []interface{}{script, keys}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Eval", _s...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

func (_mr *_MockRedisClientRecorder) Eval(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Eval", _s...)
}
