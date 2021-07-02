package lstmarketing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstMarketingQuerybyorderidAPIRequest 根据订单查询营销信息 API请求
// alibaba.lst.marketing.querybyorderid
//
// 根据订单查询营销信息
type AlibabaLstMarketingQuerybyorderidAPIRequest struct {
	model.Params
	// 主订单
	_mainOrderId int64
	// 子订单
	_subOrderId int64
}

// NewAlibabaLstMarketingQuerybyorderidRequest 初始化AlibabaLstMarketingQuerybyorderidAPIRequest对象
func NewAlibabaLstMarketingQuerybyorderidRequest() *AlibabaLstMarketingQuerybyorderidAPIRequest {
	return &AlibabaLstMarketingQuerybyorderidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.marketing.querybyorderid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MainOrderId Setter
// 主订单
func (r *AlibabaLstMarketingQuerybyorderidAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// Get MainOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// Set is SubOrderId Setter
// 子订单
func (r *AlibabaLstMarketingQuerybyorderidAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// Get SubOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}
