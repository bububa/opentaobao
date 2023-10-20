package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest 删除商品池活动商品【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
type AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest 初始化AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest() *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest {
	return &AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolActivitySkuDeleteRequest()
	},
}

// GetAlibabaRetailMarketingItempoolActivitySkuDeleteRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest
func GetAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest() *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest {
	return poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest.Get().(*AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest 将 AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest(v *AlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySkuDeleteAPIRequest.Put(v)
}
