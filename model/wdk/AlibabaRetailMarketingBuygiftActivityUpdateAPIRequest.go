package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest 更新单品买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.update
//
// 同城零售单品买赠活动更新
type AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest struct {
	model.Params
	// 更新单品买赠活动参数
	_param *BuyGiftActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivityUpdateRequest 初始化AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityUpdateRequest() *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新单品买赠活动参数
func (r *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) SetParam(_param *BuyGiftActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) GetParam() *BuyGiftActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivityUpdateRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivityUpdateRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest
func GetAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest() *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest 将 AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest(v *AlibabaRetailMarketingBuygiftActivityUpdateAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityUpdateAPIRequest.Put(v)
}
