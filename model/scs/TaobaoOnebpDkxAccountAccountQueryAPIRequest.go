package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxAccountAccountQueryAPIRequest 万相台账号余额查询 API请求
// taobao.onebp.dkx.account.account.query
//
// 万相台账号余额查询
type TaobaoOnebpDkxAccountAccountQueryAPIRequest struct {
	model.Params
}

// NewTaobaoOnebpDkxAccountAccountQueryRequest 初始化TaobaoOnebpDkxAccountAccountQueryAPIRequest对象
func NewTaobaoOnebpDkxAccountAccountQueryRequest() *TaobaoOnebpDkxAccountAccountQueryAPIRequest {
	return &TaobaoOnebpDkxAccountAccountQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOnebpDkxAccountAccountQueryAPIRequest) GetApiMethodName() string {
	return "taobao.onebp.dkx.account.account.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOnebpDkxAccountAccountQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOnebpDkxAccountAccountQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}
