// Code generated by MockGen. DO NOT EDIT.
// Source: redis/interfaces/redis.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	redis "github.com/go-redis/redis"
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/itsjunglexyz/extensions/v9/redis/interfaces"
)

// MockTraceWrapper is a mock of TraceWrapper interface.
type MockTraceWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockTraceWrapperMockRecorder
}

// MockTraceWrapperMockRecorder is the mock recorder for MockTraceWrapper.
type MockTraceWrapperMockRecorder struct {
	mock *MockTraceWrapper
}

// NewMockTraceWrapper creates a new mock instance.
func NewMockTraceWrapper(ctrl *gomock.Controller) *MockTraceWrapper {
	mock := &MockTraceWrapper{ctrl: ctrl}
	mock.recorder = &MockTraceWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTraceWrapper) EXPECT() *MockTraceWrapperMockRecorder {
	return m.recorder
}

// WithContext mocks base method.
func (m *MockTraceWrapper) WithContext(ctx context.Context, c interfaces.RedisClient) interfaces.RedisClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", ctx, c)
	ret0, _ := ret[0].(interfaces.RedisClient)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockTraceWrapperMockRecorder) WithContext(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockTraceWrapper)(nil).WithContext), ctx, c)
}

// MockRedisClient is a mock of RedisClient interface.
type MockRedisClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientMockRecorder
}

// MockRedisClientMockRecorder is the mock recorder for MockRedisClient.
type MockRedisClientMockRecorder struct {
	mock *MockRedisClient
}

// NewMockRedisClient creates a new mock instance.
func NewMockRedisClient(ctrl *gomock.Controller) *MockRedisClient {
	mock := &MockRedisClient{ctrl: ctrl}
	mock.recorder = &MockRedisClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClient) EXPECT() *MockRedisClientMockRecorder {
	return m.recorder
}

// BLPop mocks base method.
func (m *MockRedisClient) BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{timeout}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BLPop", varargs...)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// BLPop indicates an expected call of BLPop.
func (mr *MockRedisClientMockRecorder) BLPop(timeout interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{timeout}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BLPop", reflect.TypeOf((*MockRedisClient)(nil).BLPop), varargs...)
}

// Close mocks base method.
func (m *MockRedisClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisClient)(nil).Close))
}

// Context mocks base method.
func (m *MockRedisClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockRedisClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockRedisClient)(nil).Context))
}

// Del mocks base method.
func (m *MockRedisClient) Del(keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisClientMockRecorder) Del(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClient)(nil).Del), keys...)
}

// Eval mocks base method.
func (m *MockRedisClient) Eval(script string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{script, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Eval", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// Eval indicates an expected call of Eval.
func (mr *MockRedisClientMockRecorder) Eval(script, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{script, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eval", reflect.TypeOf((*MockRedisClient)(nil).Eval), varargs...)
}

// EvalSha mocks base method.
func (m *MockRedisClient) EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{sha1, keys}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EvalSha", varargs...)
	ret0, _ := ret[0].(*redis.Cmd)
	return ret0
}

// EvalSha indicates an expected call of EvalSha.
func (mr *MockRedisClientMockRecorder) EvalSha(sha1, keys interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sha1, keys}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvalSha", reflect.TypeOf((*MockRedisClient)(nil).EvalSha), varargs...)
}

// Exists mocks base method.
func (m *MockRedisClient) Exists(keys ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exists", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockRedisClientMockRecorder) Exists(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRedisClient)(nil).Exists), keys...)
}

// Get mocks base method.
func (m *MockRedisClient) Get(key string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisClientMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClient)(nil).Get), key)
}

// HDel mocks base method.
func (m *MockRedisClient) HDel(key string, fields ...string) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HDel", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// HDel indicates an expected call of HDel.
func (mr *MockRedisClientMockRecorder) HDel(key interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HDel", reflect.TypeOf((*MockRedisClient)(nil).HDel), varargs...)
}

// HGet mocks base method.
func (m *MockRedisClient) HGet(key, field string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGet", key, field)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// HGet indicates an expected call of HGet.
func (mr *MockRedisClientMockRecorder) HGet(key, field interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGet", reflect.TypeOf((*MockRedisClient)(nil).HGet), key, field)
}

// HGetAll mocks base method.
func (m *MockRedisClient) HGetAll(arg0 string) *redis.StringStringMapCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HGetAll", arg0)
	ret0, _ := ret[0].(*redis.StringStringMapCmd)
	return ret0
}

// HGetAll indicates an expected call of HGetAll.
func (mr *MockRedisClientMockRecorder) HGetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HGetAll", reflect.TypeOf((*MockRedisClient)(nil).HGetAll), arg0)
}

// HMGet mocks base method.
func (m *MockRedisClient) HMGet(arg0 string, arg1 ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HMGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// HMGet indicates an expected call of HMGet.
func (mr *MockRedisClientMockRecorder) HMGet(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMGet", reflect.TypeOf((*MockRedisClient)(nil).HMGet), varargs...)
}

// HMSet mocks base method.
func (m *MockRedisClient) HMSet(arg0 string, arg1 map[string]interface{}) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HMSet", arg0, arg1)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// HMSet indicates an expected call of HMSet.
func (mr *MockRedisClientMockRecorder) HMSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HMSet", reflect.TypeOf((*MockRedisClient)(nil).HMSet), arg0, arg1)
}

// HSet mocks base method.
func (m *MockRedisClient) HSet(key, field string, value interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HSet", key, field, value)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// HSet indicates an expected call of HSet.
func (mr *MockRedisClientMockRecorder) HSet(key, field, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HSet", reflect.TypeOf((*MockRedisClient)(nil).HSet), key, field, value)
}

// LRange mocks base method.
func (m *MockRedisClient) LRange(key string, start, stop int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LRange", key, start, stop)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// LRange indicates an expected call of LRange.
func (mr *MockRedisClientMockRecorder) LRange(key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LRange", reflect.TypeOf((*MockRedisClient)(nil).LRange), key, start, stop)
}

// MGet mocks base method.
func (m *MockRedisClient) MGet(keys ...string) *redis.SliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MGet", varargs...)
	ret0, _ := ret[0].(*redis.SliceCmd)
	return ret0
}

// MGet indicates an expected call of MGet.
func (mr *MockRedisClientMockRecorder) MGet(keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MGet", reflect.TypeOf((*MockRedisClient)(nil).MGet), keys...)
}

// Ping mocks base method.
func (m *MockRedisClient) Ping() *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockRedisClientMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockRedisClient)(nil).Ping))
}

// RPopLPush mocks base method.
func (m *MockRedisClient) RPopLPush(source, destination string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RPopLPush", source, destination)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// RPopLPush indicates an expected call of RPopLPush.
func (mr *MockRedisClientMockRecorder) RPopLPush(source, destination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPopLPush", reflect.TypeOf((*MockRedisClient)(nil).RPopLPush), source, destination)
}

// RPush mocks base method.
func (m *MockRedisClient) RPush(key string, values ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RPush", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// RPush indicates an expected call of RPush.
func (mr *MockRedisClientMockRecorder) RPush(key interface{}, values ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPush", reflect.TypeOf((*MockRedisClient)(nil).RPush), varargs...)
}

// SAdd mocks base method.
func (m *MockRedisClient) SAdd(key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockRedisClientMockRecorder) SAdd(key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockRedisClient)(nil).SAdd), varargs...)
}

// SCard mocks base method.
func (m *MockRedisClient) SCard(key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SCard", key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SCard indicates an expected call of SCard.
func (mr *MockRedisClientMockRecorder) SCard(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SCard", reflect.TypeOf((*MockRedisClient)(nil).SCard), key)
}

// SIsMember mocks base method.
func (m *MockRedisClient) SIsMember(key string, member interface{}) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SIsMember", key, member)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SIsMember indicates an expected call of SIsMember.
func (mr *MockRedisClientMockRecorder) SIsMember(key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SIsMember", reflect.TypeOf((*MockRedisClient)(nil).SIsMember), key, member)
}

// SMembers mocks base method.
func (m *MockRedisClient) SMembers(key string) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", key)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SMembers indicates an expected call of SMembers.
func (mr *MockRedisClientMockRecorder) SMembers(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockRedisClient)(nil).SMembers), key)
}

// SPopN mocks base method.
func (m *MockRedisClient) SPopN(key string, count int64) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SPopN", key, count)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// SPopN indicates an expected call of SPopN.
func (mr *MockRedisClientMockRecorder) SPopN(key, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SPopN", reflect.TypeOf((*MockRedisClient)(nil).SPopN), key, count)
}

// SRem mocks base method.
func (m *MockRedisClient) SRem(key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// SRem indicates an expected call of SRem.
func (mr *MockRedisClientMockRecorder) SRem(key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SRem", reflect.TypeOf((*MockRedisClient)(nil).SRem), varargs...)
}

// ScriptExists mocks base method.
func (m *MockRedisClient) ScriptExists(scripts ...string) *redis.BoolSliceCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range scripts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScriptExists", varargs...)
	ret0, _ := ret[0].(*redis.BoolSliceCmd)
	return ret0
}

// ScriptExists indicates an expected call of ScriptExists.
func (mr *MockRedisClientMockRecorder) ScriptExists(scripts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptExists", reflect.TypeOf((*MockRedisClient)(nil).ScriptExists), scripts...)
}

// ScriptLoad mocks base method.
func (m *MockRedisClient) ScriptLoad(script string) *redis.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScriptLoad", script)
	ret0, _ := ret[0].(*redis.StringCmd)
	return ret0
}

// ScriptLoad indicates an expected call of ScriptLoad.
func (mr *MockRedisClientMockRecorder) ScriptLoad(script interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScriptLoad", reflect.TypeOf((*MockRedisClient)(nil).ScriptLoad), script)
}

// Set mocks base method.
func (m *MockRedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value, expiration)
	ret0, _ := ret[0].(*redis.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisClientMockRecorder) Set(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClient)(nil).Set), key, value, expiration)
}

// SetNX mocks base method.
func (m *MockRedisClient) SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetNX", key, value, expiration)
	ret0, _ := ret[0].(*redis.BoolCmd)
	return ret0
}

// SetNX indicates an expected call of SetNX.
func (mr *MockRedisClientMockRecorder) SetNX(key, value, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNX", reflect.TypeOf((*MockRedisClient)(nil).SetNX), key, value, expiration)
}

// TTL mocks base method.
func (m *MockRedisClient) TTL(key string) *redis.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", key)
	ret0, _ := ret[0].(*redis.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockRedisClientMockRecorder) TTL(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockRedisClient)(nil).TTL), key)
}

// TxPipeline mocks base method.
func (m *MockRedisClient) TxPipeline() redis.Pipeliner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxPipeline")
	ret0, _ := ret[0].(redis.Pipeliner)
	return ret0
}

// TxPipeline indicates an expected call of TxPipeline.
func (mr *MockRedisClientMockRecorder) TxPipeline() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxPipeline", reflect.TypeOf((*MockRedisClient)(nil).TxPipeline))
}

// WithContext mocks base method.
func (m *MockRedisClient) WithContext(arg0 context.Context) *redis.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(*redis.Client)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockRedisClientMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockRedisClient)(nil).WithContext), arg0)
}

// ZAdd mocks base method.
func (m *MockRedisClient) ZAdd(key string, members ...redis.Z) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZAdd", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZAdd indicates an expected call of ZAdd.
func (mr *MockRedisClientMockRecorder) ZAdd(key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZAdd", reflect.TypeOf((*MockRedisClient)(nil).ZAdd), varargs...)
}

// ZCard mocks base method.
func (m *MockRedisClient) ZCard(key string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZCard", key)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZCard indicates an expected call of ZCard.
func (mr *MockRedisClientMockRecorder) ZCard(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZCard", reflect.TypeOf((*MockRedisClient)(nil).ZCard), key)
}

// ZRangeByScore mocks base method.
func (m *MockRedisClient) ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScore", key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRangeByScore indicates an expected call of ZRangeByScore.
func (mr *MockRedisClientMockRecorder) ZRangeByScore(key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScore", reflect.TypeOf((*MockRedisClient)(nil).ZRangeByScore), key, opt)
}

// ZRangeByScoreWithScores mocks base method.
func (m *MockRedisClient) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeByScoreWithScores", key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeByScoreWithScores indicates an expected call of ZRangeByScoreWithScores.
func (mr *MockRedisClientMockRecorder) ZRangeByScoreWithScores(key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeByScoreWithScores", reflect.TypeOf((*MockRedisClient)(nil).ZRangeByScoreWithScores), key, opt)
}

// ZRangeWithScores mocks base method.
func (m *MockRedisClient) ZRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRangeWithScores", key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRangeWithScores indicates an expected call of ZRangeWithScores.
func (mr *MockRedisClientMockRecorder) ZRangeWithScores(key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRangeWithScores", reflect.TypeOf((*MockRedisClient)(nil).ZRangeWithScores), key, start, stop)
}

// ZRank mocks base method.
func (m *MockRedisClient) ZRank(key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRank", key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRank indicates an expected call of ZRank.
func (mr *MockRedisClientMockRecorder) ZRank(key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRank", reflect.TypeOf((*MockRedisClient)(nil).ZRank), key, member)
}

// ZRem mocks base method.
func (m *MockRedisClient) ZRem(key string, members ...interface{}) *redis.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ZRem", varargs...)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRem indicates an expected call of ZRem.
func (mr *MockRedisClientMockRecorder) ZRem(key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRem", reflect.TypeOf((*MockRedisClient)(nil).ZRem), varargs...)
}

// ZRevRangeByScore mocks base method.
func (m *MockRedisClient) ZRevRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScore", key, opt)
	ret0, _ := ret[0].(*redis.StringSliceCmd)
	return ret0
}

// ZRevRangeByScore indicates an expected call of ZRevRangeByScore.
func (mr *MockRedisClientMockRecorder) ZRevRangeByScore(key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScore", reflect.TypeOf((*MockRedisClient)(nil).ZRevRangeByScore), key, opt)
}

// ZRevRangeByScoreWithScores mocks base method.
func (m *MockRedisClient) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeByScoreWithScores", key, opt)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeByScoreWithScores indicates an expected call of ZRevRangeByScoreWithScores.
func (mr *MockRedisClientMockRecorder) ZRevRangeByScoreWithScores(key, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeByScoreWithScores", reflect.TypeOf((*MockRedisClient)(nil).ZRevRangeByScoreWithScores), key, opt)
}

// ZRevRangeWithScores mocks base method.
func (m *MockRedisClient) ZRevRangeWithScores(key string, start, stop int64) *redis.ZSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRangeWithScores", key, start, stop)
	ret0, _ := ret[0].(*redis.ZSliceCmd)
	return ret0
}

// ZRevRangeWithScores indicates an expected call of ZRevRangeWithScores.
func (mr *MockRedisClientMockRecorder) ZRevRangeWithScores(key, start, stop interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRangeWithScores", reflect.TypeOf((*MockRedisClient)(nil).ZRevRangeWithScores), key, start, stop)
}

// ZRevRank mocks base method.
func (m *MockRedisClient) ZRevRank(key, member string) *redis.IntCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZRevRank", key, member)
	ret0, _ := ret[0].(*redis.IntCmd)
	return ret0
}

// ZRevRank indicates an expected call of ZRevRank.
func (mr *MockRedisClientMockRecorder) ZRevRank(key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZRevRank", reflect.TypeOf((*MockRedisClient)(nil).ZRevRank), key, member)
}

// ZScore mocks base method.
func (m *MockRedisClient) ZScore(key, member string) *redis.FloatCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ZScore", key, member)
	ret0, _ := ret[0].(*redis.FloatCmd)
	return ret0
}

// ZScore indicates an expected call of ZScore.
func (mr *MockRedisClientMockRecorder) ZScore(key, member interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ZScore", reflect.TypeOf((*MockRedisClient)(nil).ZScore), key, member)
}
