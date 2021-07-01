package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoregroupDeleteAPIRequest
删除门店库 API请求
taobao.place.storegroup.delete

删除门店库 */
type TaobaoPlaceStoregroupDeleteAPIRequest struct {
	model.Params
	// 库Id
	_id int64
}

// New
