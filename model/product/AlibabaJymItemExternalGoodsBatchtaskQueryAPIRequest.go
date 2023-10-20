package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest 交易猫外部商家查询商品批量任务接口 API请求
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
type AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest struct {
	model.Params
	// 查询商品批量任务接口入参
	_goodsBatchTaskQuery *GoodsBatchTaskQueryDto
}

// NewAlibabajymitemexternalgoodsbatchtaskqueryRequest 初始化AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchtaskqueryRequest() *AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest {
	return &AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batchtask.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsBatchTaskQuery is GoodsBatchTaskQuery Setter
// 查询商品批量任务接口入参
func (r *AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest) SetGoodsBatchTaskQuery(_goodsBatchTaskQuery *GoodsBatchTaskQueryDto) error {
	r._goodsBatchTaskQuery = _goodsBatchTaskQuery
	r.Set("goods_batch_task_query", _goodsBatchTaskQuery)
	return nil
}

// GetGoodsBatchTaskQuery GoodsBatchTaskQuery Getter
func (r AlibabajymitemexternalgoodsbatchtaskqueryAPIRequest) GetGoodsBatchTaskQuery() *GoodsBatchTaskQueryDto {
	return r._goodsBatchTaskQuery
}
