package lstmarketing

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstMarketingQuerybyorderidAPIRequest) Reset() {
	r._mainOrderId = 0
	r._subOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.marketing.querybyorderid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单
func (r *AlibabaLstMarketingQuerybyorderidAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetSubOrderId is SubOrderId Setter
// 子订单
func (r *AlibabaLstMarketingQuerybyorderidAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlibabaLstMarketingQuerybyorderidAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}

var poolAlibabaLstMarketingQuerybyorderidAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstMarketingQuerybyorderidRequest()
	},
}

// GetAlibabaLstMarketingQuerybyorderidRequest 从 sync.Pool 获取 AlibabaLstMarketingQuerybyorderidAPIRequest
func GetAlibabaLstMarketingQuerybyorderidAPIRequest() *AlibabaLstMarketingQuerybyorderidAPIRequest {
	return poolAlibabaLstMarketingQuerybyorderidAPIRequest.Get().(*AlibabaLstMarketingQuerybyorderidAPIRequest)
}

// ReleaseAlibabaLstMarketingQuerybyorderidAPIRequest 将 AlibabaLstMarketingQuerybyorderidAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstMarketingQuerybyorderidAPIRequest(v *AlibabaLstMarketingQuerybyorderidAPIRequest) {
	v.Reset()
	poolAlibabaLstMarketingQuerybyorderidAPIRequest.Put(v)
}
