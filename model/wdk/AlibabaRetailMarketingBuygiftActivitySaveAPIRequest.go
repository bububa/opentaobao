package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySaveAPIRequest 【同城零售】单品买赠活动保存 API请求
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
type AlibabaRetailMarketingBuygiftActivitySaveAPIRequest struct {
	model.Params
	// 保存单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivitySaveRequest 初始化AlibabaRetailMarketingBuygiftActivitySaveAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivitySaveRequest() *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivitySaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 保存单品买赠活动参数
func (r *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivitySaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivitySaveRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySaveRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySaveAPIRequest
func GetAlibabaRetailMarketingBuygiftActivitySaveAPIRequest() *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivitySaveAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivitySaveAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySaveAPIRequest 将 AlibabaRetailMarketingBuygiftActivitySaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySaveAPIRequest(v *AlibabaRetailMarketingBuygiftActivitySaveAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySaveAPIRequest.Put(v)
}
