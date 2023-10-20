package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest 商品池活动新增商品 API请求
// alibaba.retail.marketing.itempool.activity.sku.add
//
// 新增或更新商品池活动商品信息【同城零售】
type AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest struct {
	model.Params
	// 入参
	_param *ItemPoolActivityElementOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivitySkuAddRequest 初始化AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivitySkuAddRequest() *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest {
	return &AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) SetParam(_param *ItemPoolActivityElementOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) GetParam() *ItemPoolActivityElementOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolActivitySkuAddRequest()
	},
}

// GetAlibabaRetailMarketingItempoolActivitySkuAddRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest
func GetAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest() *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest {
	return poolAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest.Get().(*AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest 将 AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest(v *AlibabaRetailMarketingItempoolActivitySkuAddAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySkuAddAPIRequest.Put(v)
}
