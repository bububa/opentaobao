package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenInventorybatchQueryAPIRequest
商品单仓批次库存查询接口 API请求
taobao.qimen.inventorybatch.query

ERP 通过该接口查询指定商品的单仓批次库存 */
type TaobaoQimenInventorybatchQueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoQimenInventorybatchQueryRequest
}

// New
