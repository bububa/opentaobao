package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelmemberderbystatesyncAPIRequest 德比侧同步卡、券状态接口 API请求
// taobao.xhotel.member.derby.state.sync
//
// 德比侧同步卡、券状态接口
type TaobaoxhotelmemberderbystatesyncAPIRequest struct {
	model.Params
	// 入参
	_param *StateSynchronizeParam
}

// NewTaobaoxhotelmemberderbystatesyncRequest 初始化TaobaoxhotelmemberderbystatesyncAPIRequest对象
func NewTaobaoxhotelmemberderbystatesyncRequest() *TaobaoxhotelmemberderbystatesyncAPIRequest {
	return &TaobaoxhotelmemberderbystatesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelmemberderbystatesyncAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.derby.state.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelmemberderbystatesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelmemberderbystatesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoxhotelmemberderbystatesyncAPIRequest) SetParam(_param *StateSynchronizeParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoxhotelmemberderbystatesyncAPIRequest) GetParam() *StateSynchronizeParam {
	return r._param
}
