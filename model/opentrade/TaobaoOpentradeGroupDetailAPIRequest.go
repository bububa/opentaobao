package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupDetailAPIRequest
组团购场景查询团详情 API请求
taobao.opentrade.group.detail

组团购场景下，查询团详情 */
type TaobaoOpentradeGroupDetailAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
}

// New
