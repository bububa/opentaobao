package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodiandepositgetAPIRequest 宝点用户帐户查询（已迁移） API请求
// taobao.baodian.deposit.get
//
// 查询用户宝点帐户信息及当前宝点价格
type TaobaobaodiandepositgetAPIRequest struct {
	model.Params
}

// NewTaobaobaodiandepositgetRequest 初始化TaobaobaodiandepositgetAPIRequest对象
func NewTaobaobaodiandepositgetRequest() *TaobaobaodiandepositgetAPIRequest {
	return &TaobaobaodiandepositgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaodiandepositgetAPIRequest) GetApiMethodName() string {
	return "taobao.baodian.deposit.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaodiandepositgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaodiandepositgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
