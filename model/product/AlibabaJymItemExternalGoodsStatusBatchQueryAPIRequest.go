package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest 交易猫外部商家商品状态批量查询接口 API请求
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
type AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest struct {
	model.Params
	// 批量查询商品状态接口入参
	_batchGoodsStatusQuery *BatchGoodsStatusQueryDto
}

// NewAlibabaJymItemExternalGoodsStatusBatchQueryRequest 初始化AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest对象
func NewAlibabaJymItemExternalGoodsStatusBatchQueryRequest() *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest {
	return &AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.status.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBatchGoodsStatusQuery is BatchGoodsStatusQuery Setter
// 批量查询商品状态接口入参
func (r *AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) SetBatchGoodsStatusQuery(_batchGoodsStatusQuery *BatchGoodsStatusQueryDto) error {
	r._batchGoodsStatusQuery = _batchGoodsStatusQuery
	r.Set("batch_goods_status_query", _batchGoodsStatusQuery)
	return nil
}

// GetBatchGoodsStatusQuery BatchGoodsStatusQuery Getter
func (r AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest) GetBatchGoodsStatusQuery() *BatchGoodsStatusQueryDto {
	return r._batchGoodsStatusQuery
}
