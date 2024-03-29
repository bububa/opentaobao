package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolExcludeskucodeAPIRequest 商品池排除商品【品类优惠使用】 API请求
// alibaba.hm.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
type AlibabaHmMarketingItempoolExcludeskucodeAPIRequest struct {
	model.Params
	// 商品对象
	_param0 *ItemPoolSku
	// 活动基本信息
	_param1 *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolExcludeskucodeRequest 初始化AlibabaHmMarketingItempoolExcludeskucodeAPIRequest对象
func NewAlibabaHmMarketingItempoolExcludeskucodeRequest() *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest {
	return &AlibabaHmMarketingItempoolExcludeskucodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) Reset() {
	r._param0 = nil
	r._param1 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.excludeskucode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 商品对象
func (r *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) SetParam0(_param0 *ItemPoolSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) GetParam0() *ItemPoolSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 活动基本信息
func (r *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

var poolAlibabaHmMarketingItempoolExcludeskucodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolExcludeskucodeRequest()
	},
}

// GetAlibabaHmMarketingItempoolExcludeskucodeRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolExcludeskucodeAPIRequest
func GetAlibabaHmMarketingItempoolExcludeskucodeAPIRequest() *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest {
	return poolAlibabaHmMarketingItempoolExcludeskucodeAPIRequest.Get().(*AlibabaHmMarketingItempoolExcludeskucodeAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolExcludeskucodeAPIRequest 将 AlibabaHmMarketingItempoolExcludeskucodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolExcludeskucodeAPIRequest(v *AlibabaHmMarketingItempoolExcludeskucodeAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolExcludeskucodeAPIRequest.Put(v)
}
