package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStorerelatesubGetAPIRequest
门店和子门店关系查找 API请求
taobao.place.storerelatesub.get

门店和子门店关系查找 */
type TaobaoPlaceStorerelatesubGetAPIRequest struct {
	model.Params
	// 门店Id
	_storeId int64
	// 查询语句
	_query string
	// 第几页
	_pageNo int64
	// 页大小
	_pageSize int64
}

// New
