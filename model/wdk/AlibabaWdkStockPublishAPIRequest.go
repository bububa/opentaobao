package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkstockpublishAPIRequest 五道口库存发布接口（针对外部渠道） API请求
// alibaba.wdk.stock.publish
//
// 五道口库存发布接口（针对外部渠道）
type AlibabawdkstockpublishAPIRequest struct {
	model.Params
	// 批量参数
	_batchStockPublishDto *BatchStockPublishDto
}

// NewAlibabawdkstockpublishRequest 初始化AlibabawdkstockpublishAPIRequest对象
func NewAlibabawdkstockpublishRequest() *AlibabawdkstockpublishAPIRequest {
	return &AlibabawdkstockpublishAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkstockpublishAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.stock.publish"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkstockpublishAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkstockpublishAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchStockPublishDto is BatchStockPublishDto Setter
// 批量参数
func (r *AlibabawdkstockpublishAPIRequest) SetBatchStockPublishDto(_batchStockPublishDto *BatchStockPublishDto) error {
	r._batchStockPublishDto = _batchStockPublishDto
	r.Set("batch_stock_publish_dto", _batchStockPublishDto)
	return nil
}

// GetBatchStockPublishDto BatchStockPublishDto Getter
func (r AlibabawdkstockpublishAPIRequest) GetBatchStockPublishDto() *BatchStockPublishDto {
	return r._batchStockPublishDto
}
