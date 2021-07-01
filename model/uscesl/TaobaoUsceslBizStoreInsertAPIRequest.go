package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizStoreInsertAPIRequest
新增电子价签商家门店接口 API请求
taobao.uscesl.biz.store.insert

新增电子价签商家门店接口 */
type TaobaoUsceslBizStoreInsertAPIRequest struct {
	model.Params
	// 门店名称
	_storeName string
	// 门店外部ID，要保持同一商家下的唯一性
	_storeOutId string
	// 商家标识key
	_bizBrandKey string
}

// New
