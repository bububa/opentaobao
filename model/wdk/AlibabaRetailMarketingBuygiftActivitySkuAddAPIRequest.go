package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest 添加单品买赠活动商品 API请求
// alibaba.retail.marketing.buygift.activity.sku.add
//
// 新增或更新单品买赠活动商品信息【同城零售】
type AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest struct {
	model.Params
	// 添加活动商品参数
	_param *BuyGiftActivitySkuOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest 初始化AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest() *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.sku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 添加活动商品参数
func (r *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) SetParam(_param *BuyGiftActivitySkuOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) GetParam() *BuyGiftActivitySkuOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivitySkuAddRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySkuAddRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest
func GetAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest() *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest 将 AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest(v *AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest.Put(v)
}
