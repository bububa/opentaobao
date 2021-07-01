package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoregroupUpdateAPIRequest
门店库修改基本信息 API请求
taobao.place.storegroup.update

门店库修改基本信息 */
type TaobaoPlaceStoregroupUpdateAPIRequest struct {
	model.Params
	// 库id
	_id int64
	// 库名称
	_name string
	// 库备注
	_desc string
}

// New
