package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairQueryitemAPIRequest 换购商品查询 API请求
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabaHmMarketingItempoolStairQueryitemAPIRequest struct {
	model.Params
	// 换购商品查询参数
	_param0 *ActivitySkuQuery
}

// NewAlibabaHmMarketingItempoolStairQueryitemRequest 初始化AlibabaHmMarketingItempoolStairQueryitemAPIRequest对象
func NewAlibabaHmMarketingItempoolStairQueryitemRequest() *AlibabaHmMarketingItempoolStairQueryitemAPIRequest {
	return &AlibabaHmMarketingItempoolStairQueryitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolStairQueryitemAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.stair.queryitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 换购商品查询参数
func (r *AlibabaHmMarketingItempoolStairQueryitemAPIRequest) SetParam0(_param0 *ActivitySkuQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaHmMarketingItempoolStairQueryitemAPIRequest) GetParam0() *ActivitySkuQuery {
	return r._param0
}

var poolAlibabaHmMarketingItempoolStairQueryitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolStairQueryitemRequest()
	},
}

// GetAlibabaHmMarketingItempoolStairQueryitemRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolStairQueryitemAPIRequest
func GetAlibabaHmMarketingItempoolStairQueryitemAPIRequest() *AlibabaHmMarketingItempoolStairQueryitemAPIRequest {
	return poolAlibabaHmMarketingItempoolStairQueryitemAPIRequest.Get().(*AlibabaHmMarketingItempoolStairQueryitemAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolStairQueryitemAPIRequest 将 AlibabaHmMarketingItempoolStairQueryitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolStairQueryitemAPIRequest(v *AlibabaHmMarketingItempoolStairQueryitemAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolStairQueryitemAPIRequest.Put(v)
}
