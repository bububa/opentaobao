package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest 交易猫外部商家查询商品批量任务接口 API请求
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
type AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest struct {
	model.Params
	// 查询商品批量任务接口入参
	_goodsBatchTaskQuery *GoodsBatchTaskQueryDto
}

// NewAlibabaJymItemExternalGoodsBatchtaskQueryRequest 初始化AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchtaskQueryRequest() *AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) Reset() {
	r._goodsBatchTaskQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batchtask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsBatchTaskQuery is GoodsBatchTaskQuery Setter
// 查询商品批量任务接口入参
func (r *AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) SetGoodsBatchTaskQuery(_goodsBatchTaskQuery *GoodsBatchTaskQueryDto) error {
	r._goodsBatchTaskQuery = _goodsBatchTaskQuery
	r.Set("goods_batch_task_query", _goodsBatchTaskQuery)
	return nil
}

// GetGoodsBatchTaskQuery GoodsBatchTaskQuery Getter
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetGoodsBatchTaskQuery() *GoodsBatchTaskQueryDto {
	return r._goodsBatchTaskQuery
}

var poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaJymItemExternalGoodsBatchtaskQueryRequest()
	},
}

// GetAlibabaJymItemExternalGoodsBatchtaskQueryRequest 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest
func GetAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest() *AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest {
	return poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest.Get().(*AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest)
}

// ReleaseAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest 将 AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest(v *AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest.Put(v)
}
