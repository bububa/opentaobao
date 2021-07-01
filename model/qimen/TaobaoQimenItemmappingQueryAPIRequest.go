package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenItemmappingQueryAPIRequest
前后端商品映射查询接口 API请求
taobao.qimen.itemmapping.query

前后端商品映射查询接口 */
type TaobaoQimenItemmappingQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenItemmappingQueryRequest
}

// New
