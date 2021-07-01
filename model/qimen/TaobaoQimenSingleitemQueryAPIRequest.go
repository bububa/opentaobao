package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenSingleitemQueryAPIRequest
商品查询接口 API请求
taobao.qimen.singleitem.query

商品查询接口 */
type TaobaoQimenSingleitemQueryAPIRequest struct {
	model.Params
	//
	_request *RequestDo
}

// New
