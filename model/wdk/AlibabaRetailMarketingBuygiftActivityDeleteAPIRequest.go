package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest 删除单品买赠活动 API请求
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
type AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest struct {
	model.Params
	// 删除活动参数
	_param *ItemDiscountActivityOperateRequest
}

// NewAlibabaRetailMarketingBuygiftActivityDeleteRequest 初始化AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest对象
func NewAlibabaRetailMarketingBuygiftActivityDeleteRequest() *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest {
	return &AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.marketing.buygift.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 删除活动参数
func (r *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) SetParam(_param *ItemDiscountActivityOperateRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) GetParam() *ItemDiscountActivityOperateRequest {
	return r._param
}

var poolAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailMarketingBuygiftActivityDeleteRequest()
	},
}

// GetAlibabaRetailMarketingBuygiftActivityDeleteRequest 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest
func GetAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest() *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest {
	return poolAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest.Get().(*AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest 将 AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest(v *AlibabaRetailMarketingBuygiftActivityDeleteAPIRequest) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityDeleteAPIRequest.Put(v)
}
