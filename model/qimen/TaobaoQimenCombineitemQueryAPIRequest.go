package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenCombineitemQueryAPIRequest
组合货品关系查询接口 API请求
taobao.qimen.combineitem.query

组合货品关系查询 */
type TaobaoQimenCombineitemQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenCombineitemQueryRequest
}

// New
