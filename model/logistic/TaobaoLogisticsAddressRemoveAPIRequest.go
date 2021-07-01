package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsAddressRemoveAPIRequest
删除卖家地址库 API请求
taobao.logistics.address.remove

用此接口删除卖家地址库 */
type TaobaoLogisticsAddressRemoveAPIRequest struct {
	model.Params
	// 地址库ID
	_contactId int64
}

// New
