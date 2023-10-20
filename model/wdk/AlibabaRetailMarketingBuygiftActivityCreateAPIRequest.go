package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityCreateAPIRequest 创建买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.create
//
// 同城供给买赠活动创建
type AlibabaRetailMarketingBuygiftActivityCreateAPIRequest struct {
	model.Params
	// 创建活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivityCreateRequest 初始化AlibabaRetailMarketingBuygiftActivityCreateAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityCreateRequest() *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动参数
func (r *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivityCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivityCreateRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivityCreateRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityCreateAPIRequest
func GetAlibabaRetailMarketingBuygiftActivityCreateAPIRequest() *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivityCreateAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivityCreateAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityCreateAPIRequest 将 AlibabaRetailMarketingBuygiftActivityCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityCreateAPIRequest(v *AlibabaRetailMarketingBuygiftActivityCreateAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityCreateAPIRequest.Put(v)
}
