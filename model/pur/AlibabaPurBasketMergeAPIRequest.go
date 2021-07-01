package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPurBasketMergeAPIRequest
合并购物车 API请求
alibaba.pur.basket.merge

采购商城接入第三方商家合并购物车接口服务 */
type AlibabaPurBasketMergeAPIRequest struct {
	model.Params
	// 合并购物车入参
	_paramMallMergeCartRequestDTO *MallMergeCartRequestDto
}

// New
