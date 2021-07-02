package cntms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntmsLogisticsOrderConsignAPIRequest 菜鸟配商家仓库发货 API请求
// cainiao.cntms.logistics.order.consign
//
// 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaoCntmsLogisticsOrderConsignAPIRequest struct {
	model.Params
	// 配送发货信息
	_content *CnTmsLogisticsOrderConsignContent
}

// NewCainiaoCntmsLogisticsOrderConsignRequest 初始化CainiaoCntmsLogisticsOrderConsignAPIRequest对象
func NewCainiaoCntmsLogisticsOrderConsignRequest() *CainiaoCntmsLogisticsOrderConsignAPIRequest {
	return &CainiaoCntmsLogisticsOrderConsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetApiMethodName() string {
	return "cainiao.cntms.logistics.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Content Setter
// 配送发货信息
func (r *CainiaoCntmsLogisticsOrderConsignAPIRequest) SetContent(_content *CnTmsLogisticsOrderConsignContent) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// Get Content Getter
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetContent() *CnTmsLogisticsOrderConsignContent {
	return r._content
}
