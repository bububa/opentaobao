package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlihealthDrugStoreGetAPIRequest
根据店铺id获取店铺详情 API请求
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情 */
type TaobaoAlihealthDrugStoreGetAPIRequest struct {
	model.Params
	// 店铺ID
	_shopId int64
}

// New
