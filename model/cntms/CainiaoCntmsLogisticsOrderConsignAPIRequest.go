package cntms

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCntmsLogisticsOrderConsignAPIRequest) Reset() {
	r._content = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetApiMethodName() string {
	return "cainiao.cntms.logistics.order.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 配送发货信息
func (r *CainiaoCntmsLogisticsOrderConsignAPIRequest) SetContent(_content *CnTmsLogisticsOrderConsignContent) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r CainiaoCntmsLogisticsOrderConsignAPIRequest) GetContent() *CnTmsLogisticsOrderConsignContent {
	return r._content
}

var poolCainiaoCntmsLogisticsOrderConsignAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCntmsLogisticsOrderConsignRequest()
	},
}

// GetCainiaoCntmsLogisticsOrderConsignRequest 从 sync.Pool 获取 CainiaoCntmsLogisticsOrderConsignAPIRequest
func GetCainiaoCntmsLogisticsOrderConsignAPIRequest() *CainiaoCntmsLogisticsOrderConsignAPIRequest {
	return poolCainiaoCntmsLogisticsOrderConsignAPIRequest.Get().(*CainiaoCntmsLogisticsOrderConsignAPIRequest)
}

// ReleaseCainiaoCntmsLogisticsOrderConsignAPIRequest 将 CainiaoCntmsLogisticsOrderConsignAPIRequest 放入 sync.Pool
func ReleaseCainiaoCntmsLogisticsOrderConsignAPIRequest(v *CainiaoCntmsLogisticsOrderConsignAPIRequest) {
	v.Reset()
	poolCainiaoCntmsLogisticsOrderConsignAPIRequest.Put(v)
}
