package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenWarehouseinfoQueryAPIRequest
货主仓库资源查询接口 API请求
taobao.qimen.warehouseinfo.query

货主仓库资源查询 */
type TaobaoQimenWarehouseinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenWarehouseinfoQueryRequest
}

// New
