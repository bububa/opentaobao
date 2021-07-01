package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoregroupCreateAPIRequest
商户门店库创建接口 API请求
taobao.place.storegroup.create

用于商家创建线下门店库 */
type TaobaoPlaceStoregroupCreateAPIRequest struct {
	model.Params
	// 库名
	_name string
	// 备注
	_desc string
}

// New
