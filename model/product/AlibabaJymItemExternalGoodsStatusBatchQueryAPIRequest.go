package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest 交易猫外部商家商品状态批量查询接口 API请求
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
type AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest struct {
	model.Params
	// 批量查询商品状态接口入参
	_batchGoodsStatusQuery *BatchGoodsStatusQueryDto
}

// NewAlibabajymitemexternalgoodsstatusbatchqueryRequest 初始化AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest对象
func NewAlibabajymitemexternalgoodsstatusbatchqueryRequest() *AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest {
	return &AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.status.batch.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchGoodsStatusQuery is BatchGoodsStatusQuery Setter
// 批量查询商品状态接口入参
func (r *AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest) SetBatchGoodsStatusQuery(_batchGoodsStatusQuery *BatchGoodsStatusQueryDto) error {
	r._batchGoodsStatusQuery = _batchGoodsStatusQuery
	r.Set("batch_goods_status_query", _batchGoodsStatusQuery)
	return nil
}

// GetBatchGoodsStatusQuery BatchGoodsStatusQuery Getter
func (r AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest) GetBatchGoodsStatusQuery() *BatchGoodsStatusQueryDto {
	return r._batchGoodsStatusQuery
}
