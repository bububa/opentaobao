package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreQueryAPIRequest
门店信息查询接口 API请求
taobao.place.store.query

根据用户授权信息，获取用户的门店公开信息 */
type TaobaoPlaceStoreQueryAPIRequest struct {
	model.Params
	// 业务code，用于区分业务
	_bizCode string
	// 业务外部id
	_outerId string
	// 门店id
	_storeId int64
}

// New
