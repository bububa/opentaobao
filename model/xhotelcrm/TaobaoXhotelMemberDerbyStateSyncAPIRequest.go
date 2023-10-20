package xhotelcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberDerbyStateSyncAPIRequest 德比侧同步卡、券状态接口 API请求
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
type TaobaoXhotelMemberDerbyStateSyncAPIRequest struct {
	model.Params
	// 入参
	_param *StateSynchronizeParam
}

// NewTaobaoXhotelMemberDerbyStateSyncRequest 初始化TaobaoXhotelMemberDerbyStateSyncAPIRequest对象
func NewTaobaoXhotelMemberDerbyStateSyncRequest() *TaobaoXhotelMemberDerbyStateSyncAPIRequest {
	return &TaobaoXhotelMemberDerbyStateSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMemberDerbyStateSyncAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMemberDerbyStateSyncAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.derby.state.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMemberDerbyStateSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMemberDerbyStateSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoXhotelMemberDerbyStateSyncAPIRequest) SetParam(_param *StateSynchronizeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelMemberDerbyStateSyncAPIRequest) GetParam() *StateSynchronizeParam {
	return r._param
}

var poolTaobaoXhotelMemberDerbyStateSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMemberDerbyStateSyncRequest()
	},
}

// GetTaobaoXhotelMemberDerbyStateSyncRequest 从 sync.Pool 获取 TaobaoXhotelMemberDerbyStateSyncAPIRequest
func GetTaobaoXhotelMemberDerbyStateSyncAPIRequest() *TaobaoXhotelMemberDerbyStateSyncAPIRequest {
	return poolTaobaoXhotelMemberDerbyStateSyncAPIRequest.Get().(*TaobaoXhotelMemberDerbyStateSyncAPIRequest)
}

// ReleaseTaobaoXhotelMemberDerbyStateSyncAPIRequest 将 TaobaoXhotelMemberDerbyStateSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMemberDerbyStateSyncAPIRequest(v *TaobaoXhotelMemberDerbyStateSyncAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMemberDerbyStateSyncAPIRequest.Put(v)
}
