package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportThreeplResourceGetAPIRequest
3PL直邮获取资源列表 API请求
taobao.wlb.import.threepl.resource.get

获取3pl直邮的发货可用资源 */
type TaobaoWlbImportThreeplResourceGetAPIRequest struct {
	model.Params
	// ONLINE或者OFFLINE
	_type string
	// 发货地区域id
	_fromId int64
	// 收件人地址
	_toAddress *AddressDto
}

// New
