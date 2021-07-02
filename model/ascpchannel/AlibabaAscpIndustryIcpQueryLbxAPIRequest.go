package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryIcpQueryLbxAPIRequest icp订单号查询lbx订单号 API请求
// alibaba.ascp.industry.icp.query.lbx
//
// 根据icp订单号查询lbx订单号
type AlibabaAscpIndustryIcpQueryLbxAPIRequest struct {
	model.Params
	// icps订单号
	_icpOrderCode string
}

// NewAlibabaAscpIndustryIcpQueryLbxRequest 初始化AlibabaAscpIndustryIcpQueryLbxAPIRequest对象
func NewAlibabaAscpIndustryIcpQueryLbxRequest() *AlibabaAscpIndustryIcpQueryLbxAPIRequest {
	return &AlibabaAscpIndustryIcpQueryLbxAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryIcpQueryLbxAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.icp.query.lbx"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryIcpQueryLbxAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IcpOrderCode Setter
// icps订单号
func (r *AlibabaAscpIndustryIcpQueryLbxAPIRequest) SetIcpOrderCode(_icpOrderCode string) error {
	r._icpOrderCode = _icpOrderCode
	r.Set("icp_order_code", _icpOrderCode)
	return nil
}

// Get IcpOrderCode Getter
func (r AlibabaAscpIndustryIcpQueryLbxAPIRequest) GetIcpOrderCode() string {
	return r._icpOrderCode
}
