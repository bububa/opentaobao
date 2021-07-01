package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkStockPublishAPIRequest
五道口库存发布接口（针对外部渠道） API请求
alibaba.wdk.stock.publish

五道口库存发布接口（针对外部渠道） */
type AlibabaWdkStockPublishAPIRequest struct {
	model.Params
	// 批量参数
	_batchStockPublishDto *BatchStockPublishDto
}

// New
