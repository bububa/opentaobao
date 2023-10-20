package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityQueryAPIRequest 查询单品买赠活动【同城零售】 API请求
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
type AlibabaRetailMarketingBuygiftActivityQueryAPIRequest struct {
	model.Params
	// 买赠活动查询入参
	_param0 *BuyGiftActivityQueryRequest
}

// NewAlibabaRetailMarketingBuygiftActivityQueryRequest 初始化AlibabaRetailMarketingBuygiftActivityQueryAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityQueryRequest() *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 买赠活动查询入参
func (r *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) SetParam0(_param0 *BuyGiftActivityQueryRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) GetParam0() *BuyGiftActivityQueryRequest {
	return r._param0
}

var poolAlibabaRetailMarketingBuygiftActivityQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivityQueryRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivityQueryRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityQueryAPIRequest
func GetAlibabaRetailMarketingBuygiftActivityQueryAPIRequest() *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivityQueryAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivityQueryAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityQueryAPIRequest 将 AlibabaRetailMarketingBuygiftActivityQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityQueryAPIRequest(v *AlibabaRetailMarketingBuygiftActivityQueryAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityQueryAPIRequest.Put(v)
}
