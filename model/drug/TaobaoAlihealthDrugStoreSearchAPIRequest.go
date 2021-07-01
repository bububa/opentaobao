package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlihealthDrugStoreSearchAPIRequest
药品店内搜索 API请求
taobao.alihealth.drug.store.search

提供给千牛智能客服，在阿里健康O2O店铺内搜索药品 */
type TaobaoAlihealthDrugStoreSearchAPIRequest struct {
	model.Params
	// 搜索关键字
	_keyword string
	// 每页显示数量
	_pageSize int64
	// 店铺ID
	_shopId string
	// 页码
	_pageNo int64
}

// New
