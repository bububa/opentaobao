package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest 删除单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.delete
//
// 删除单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest struct {
	model.Params
	// 删除买赠活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest 初始化AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest() *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 删除买赠活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySkuDeleteRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest
func GetAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest() *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest 将 AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest(v *AlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySkuDeleteAPIRequest.Put(v)
}
