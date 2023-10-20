package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryicpquerylbxAPIRequest icp订单号查询lbx订单号 API请求
// alibaba.ascp.industry.icp.query.lbx
//
// 根据icp订单号查询lbx订单号
type AlibabaascpindustryicpquerylbxAPIRequest struct {
	model.Params
	// icps订单号
	_icpOrderCode string
}

// NewAlibabaascpindustryicpquerylbxRequest 初始化AlibabaascpindustryicpquerylbxAPIRequest对象
func NewAlibabaascpindustryicpquerylbxRequest() *AlibabaascpindustryicpquerylbxAPIRequest {
	return &AlibabaascpindustryicpquerylbxAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustryicpquerylbxAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.icp.query.lbx"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustryicpquerylbxAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustryicpquerylbxAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIcpOrderCode is IcpOrderCode Setter
// icps订单号
func (r *AlibabaascpindustryicpquerylbxAPIRequest) SetIcpOrderCode(_icpOrderCode string) error {
	r._icpOrderCode = _icpOrderCode
	r.Set("icp_order_code", _icpOrderCode)
	return nil
}

// GetIcpOrderCode IcpOrderCode Getter
func (r AlibabaascpindustryicpquerylbxAPIRequest) GetIcpOrderCode() string {
	return r._icpOrderCode
}
