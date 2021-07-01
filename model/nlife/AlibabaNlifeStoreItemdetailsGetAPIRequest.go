package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreItemdetailsGetAPIRequest
批量获取零售加商品详情信息 API请求
alibaba.nlife.store.itemdetails.get

批量获取零售加平台上的商品详情信息 */
type AlibabaNlifeStoreItemdetailsGetAPIRequest struct {
	model.Params
	// 门店ID/设备号
	_storeId string
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ; 门店设备号-STORE_DEVICE
	_storeIdType string
	// 查询参数list
	_itemQueryDOList []ItemQueryDOList
	// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
	_itemType *model.File
}

// New
