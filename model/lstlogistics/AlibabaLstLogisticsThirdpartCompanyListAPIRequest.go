package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticsthirdpartcompanylistAPIRequest 供应商-异云-第三方物流公司列表 API请求
// alibaba.lst.logistics.thirdpart.company.list
//
// 异地云仓发货时，需填写的第三方物流公司列表
type AlibabalstlogisticsthirdpartcompanylistAPIRequest struct {
	model.Params
}

// NewAlibabalstlogisticsthirdpartcompanylistRequest 初始化AlibabalstlogisticsthirdpartcompanylistAPIRequest对象
func NewAlibabalstlogisticsthirdpartcompanylistRequest() *AlibabalstlogisticsthirdpartcompanylistAPIRequest {
	return &AlibabalstlogisticsthirdpartcompanylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstlogisticsthirdpartcompanylistAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.thirdpart.company.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstlogisticsthirdpartcompanylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstlogisticsthirdpartcompanylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}
