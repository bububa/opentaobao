package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcommoninvoicelistvtwoAPIRequest 用户常用发票信息查询接口 API请求
// taobao.xhotel.commoninvoice.list.vtwo
//
// 获取用户常用发票信息接口
type TaobaoxhotelcommoninvoicelistvtwoAPIRequest struct {
	model.Params
}

// NewTaobaoxhotelcommoninvoicelistvtwoRequest 初始化TaobaoxhotelcommoninvoicelistvtwoAPIRequest对象
func NewTaobaoxhotelcommoninvoicelistvtwoRequest() *TaobaoxhotelcommoninvoicelistvtwoAPIRequest {
	return &TaobaoxhotelcommoninvoicelistvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcommoninvoicelistvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.commoninvoice.list.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcommoninvoicelistvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcommoninvoicelistvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
