package product

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batchtask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
