package xhotelcrm

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
