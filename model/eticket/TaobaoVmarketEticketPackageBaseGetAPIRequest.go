package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketPackageBaseGetAPIRequest
获取包基本信息 API请求
taobao.vmarket.eticket.package.base.get

获取包基本信息 */
type TaobaoVmarketEticketPackageBaseGetAPIRequest struct {
	model.Params
	// 包id
	_packageId int64
}

// New
