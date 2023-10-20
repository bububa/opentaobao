package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityUpdateAPIRequest 更新商品池活动【同城零售】 API请求
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
type AlibabaRetailMarketingItempoolActivityUpdateAPIRequest struct {
	model.Params
	// 更新商品池活动参数
	_param *ItemPoolActivityOperateRequest
}

// NewAlibabaRetailMarketingItempoolActivityUpdateRequest 初始化AlibabaRetailMarketingItempoolActivityUpdateAPIRequest对象
func NewAlibabaRetailMarketingItempoolActivityUpdateRequest() *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest {
	return &AlibabaRetailMarketingItempoolActivityUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.itempool.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 更新商品池活动参数
func (r *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) SetParam(_param *ItemPoolActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) GetParam() *ItemPoolActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingItempoolActivityUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingItempoolActivityUpdateRequest()
	},
}

// GetAlibabaRetailMarketingItempoolActivityUpdateRequest 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityUpdateAPIRequest
func GetAlibabaRetailMarketingItempoolActivityUpdateAPIRequest() *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest {
	return poolAlibabaRetailMarketingItempoolActivityUpdateAPIRequest.Get().(*AlibabaRetailMarketingItempoolActivityUpdateAPIRequest)
}

// ReleaseAlibabaRetailMarketingItempoolActivityUpdateAPIRequest 将 AlibabaRetailMarketingItempoolActivityUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityUpdateAPIRequest(v *AlibabaRetailMarketingItempoolActivityUpdateAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityUpdateAPIRequest.Put(v)
}
