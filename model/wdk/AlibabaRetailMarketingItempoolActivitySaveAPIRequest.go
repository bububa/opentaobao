package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySaveAPIRequest 【同城零售】保存商品池活动 API请求
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
type AlibabaRetailMarketingItempoolActivitySaveAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivitySaveRequest 初始化AlibabaRetailMarketingItempoolActivitySaveAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySaveRequest() *AlibabaRetailMarketingItempoolActivitySaveAPIRequest {
	return &AlibabaRetailMarketingItempoolActivitySaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolActivitySaveAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivitySaveAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivitySaveAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItempoolActivitySaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolActivitySaveRequest()
	},
}

// GetAlibabaRetailMarketingItempoolActivitySaveRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySaveAPIRequest
func GetAlibabaRetailMarketingItempoolActivitySaveAPIRequest() *AlibabaRetailMarketingItempoolActivitySaveAPIRequest {
	return poolAlibabaRetailMarketingItempoolActivitySaveAPIRequest.Get().(*AlibabaRetailMarketingItempoolActivitySaveAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySaveAPIRequest 将 AlibabaRetailMarketingItempoolActivitySaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySaveAPIRequest(v *AlibabaRetailMarketingItempoolActivitySaveAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySaveAPIRequest.Put(v)
}
