package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityCreateAPIRequest 创建商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
type AlibabaRetailMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivityCreateRequest 初始化AlibabaRetailMarketingItempoolActivityCreateAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityCreateRequest() *AlibabaRetailMarketingItempoolActivityCreateAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolActivityCreateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivityCreateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivityCreateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItempoolActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolActivityCreateRequest()
	},
}

// GetAlibabaRetailMarketingItempoolActivityCreateRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityCreateAPIRequest
func GetAlibabaRetailMarketingItempoolActivityCreateAPIRequest() *AlibabaRetailMarketingItempoolActivityCreateAPIRequest {
	return poolAlibabaRetailMarketingItempoolActivityCreateAPIRequest.Get().(*AlibabaRetailMarketingItempoolActivityCreateAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolActivityCreateAPIRequest 将 AlibabaRetailMarketingItempoolActivityCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityCreateAPIRequest(v *AlibabaRetailMarketingItempoolActivityCreateAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityCreateAPIRequest.Put(v)
}
