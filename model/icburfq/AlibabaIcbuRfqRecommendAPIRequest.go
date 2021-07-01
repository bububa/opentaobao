package icburfq

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuRfqRecommendAPIRequest
rfq推荐 API请求
alibaba.icbu.rfq.recommend

rfq推荐 */
type AlibabaIcbuRfqRecommendAPIRequest struct {
	model.Params
	// 入参数据
	_queryDto *QueryDto
}

// New
