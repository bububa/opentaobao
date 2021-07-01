package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCntecItemChangeMessageAPIRequest
商品变更消息 API请求
cainiao.cntec.item.change.message

供货商商品信息变更消息 */
type CainiaoCntecItemChangeMessageAPIRequest struct {
	model.Params
	// 供应商商品变更信息
	_itemChangeMessage *SupplyItemChangeMessage
}

// New
