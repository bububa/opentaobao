package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreExtendAddAPIRequest
新增门店扩展属性 API请求
taobao.place.store.extend.add

新增授权用户的门店扩展属性 */
type TaobaoPlaceStoreExtendAddAPIRequest struct {
	model.Params
	// 门店ID
	_storeId int64
	// 扩展信息
	_etv []ExtendTypeValueTopDto
}

// New
