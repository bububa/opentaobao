package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportsResourceGetAPIRequest
获取所有服务列表 API请求
taobao.wlb.imports.resource.get

一般进口TOP接口，获取所有服务列表。 */
type TaobaoWlbImportsResourceGetAPIRequest struct {
	model.Params
	// 卖家发货地区域ID
	_fromId int64
	// 买家收货地信息
	_toAddress *ReciverAddressDo
}

// New
