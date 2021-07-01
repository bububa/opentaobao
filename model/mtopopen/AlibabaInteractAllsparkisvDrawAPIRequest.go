package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractAllsparkisvDrawAPIRequest
allspark提供抽奖tida接口对应鉴权接口 API请求
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用 */
type AlibabaInteractAllsparkisvDrawAPIRequest struct {
	model.Params
	// ddd
	_test string
	// dd
	_ddd string
}

// New
