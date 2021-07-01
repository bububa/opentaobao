package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreItemdetailGetAPIRequest
查询商品的详情信息 API请求
alibaba.nlife.store.itemdetail.get

查询零售加平台上单个商品的详情信息 */
type AlibabaNlifeStoreItemdetailGetAPIRequest struct {
	model.Params
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
	_storeIdType string
	// 门店ID
	_storeId string
	// 商品在外部商家的编码(与item_id不能同时为空)
	_outerId string
	// 商品Item的ID(与outer_id不能同时为空)
	_itemId int64
	// skuId列表-可查询指定的sku
	_skuIdList []int64
	// 商品来源类型: 0-线上商品; 1-商户导入的线下商品. 如果为空则默认值为0
	_itemType *model.File
	// 商家对商品的自用编码
	_code string
}

// New
