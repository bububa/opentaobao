package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressLogisticsBuyerFreightCalculateAPIRequest
提供给买家使用的运费计算接口 API请求
aliexpress.logistics.buyer.freight.calculate

提供给买家使用的运费计算接口 */
type AliexpressLogisticsBuyerFreightCalculateAPIRequest struct {
	model.Params
	// 运费计算请求参数
	_paramAeopFreightCalculateForBuyerDTO *AeopFreightCalculateForBuyerDto
}

// New
