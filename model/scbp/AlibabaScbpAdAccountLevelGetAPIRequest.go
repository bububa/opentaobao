package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadaccountlevelgetAPIRequest 查询推广账户等级 API请求
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
type AlibabascbpadaccountlevelgetAPIRequest struct {
	model.Params
}

// NewAlibabascbpadaccountlevelgetRequest 初始化AlibabascbpadaccountlevelgetAPIRequest对象
func NewAlibabascbpadaccountlevelgetRequest() *AlibabascbpadaccountlevelgetAPIRequest {
	return &AlibabascbpadaccountlevelgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadaccountlevelgetAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.account.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadaccountlevelgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadaccountlevelgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
