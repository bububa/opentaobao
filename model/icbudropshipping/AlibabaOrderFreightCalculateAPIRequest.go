package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOrderFreightCalculateAPIRequest
阿里巴巴下单场景运费方案计算 API请求
alibaba.order.freight.calculate

icbu开展 drop shipping 业务，阿里巴巴下单场景运费方案计算
alibaba Create order scenario freight calculation */
type AlibabaOrderFreightCalculateAPIRequest struct {
	model.Params
	// {}
	_paramMultiFreightTemplateRequest *MultiFreightTemplateRequest
}

// New
