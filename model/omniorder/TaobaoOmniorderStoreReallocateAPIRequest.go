package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreReallocateAPIRequest
rellocate API请求
taobao.omniorder.store.reallocate

门店发货提供改派接口 */
type TaobaoOmniorderStoreReallocateAPIRequest struct {
	model.Params
	// 主订单号
	_mainOrderId int64
	// 子订单号
	_subOrderIds []int64
	// 门店Id
	_storeId int64
	// 电商仓code
	_warehouseCode string
}

// New
