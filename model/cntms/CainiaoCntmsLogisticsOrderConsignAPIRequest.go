package cntms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocntmslogisticsorderconsignAPIRequest 菜鸟配商家仓库发货 API请求
// cainiao.cntms.logistics.order.consign
//
// 商家包装打印面单结束后，通知菜鸟包裹要发货
type CainiaocntmslogisticsorderconsignAPIRequest struct {
	model.Params
	// 配送发货信息
	_content *CnTmsLogisticsOrderConsignContent
}

// NewCainiaocntmslogisticsorderconsignRequest 初始化CainiaocntmslogisticsorderconsignAPIRequest对象
func NewCainiaocntmslogisticsorderconsignRequest() *CainiaocntmslogisticsorderconsignAPIRequest {
	return &CainiaocntmslogisticsorderconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocntmslogisticsorderconsignAPIRequest) GetApiMethodName() string {
	return "cainiao.cntms.logistics.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocntmslogisticsorderconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocntmslogisticsorderconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 配送发货信息
func (r *CainiaocntmslogisticsorderconsignAPIRequest) SetContent(_content *CnTmsLogisticsOrderConsignContent) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r CainiaocntmslogisticsorderconsignAPIRequest) GetContent() *CnTmsLogisticsOrderConsignContent {
	return r._content
}
