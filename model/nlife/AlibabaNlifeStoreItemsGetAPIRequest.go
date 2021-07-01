package nlife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaNlifeStoreItemsGetAPIRequest
获取门店的商品列表(在售|已下架|全部) API请求
alibaba.nlife.store.items.get

利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。 */
type AlibabaNlifeStoreItemsGetAPIRequest struct {
	model.Params
	// 门店类型: 零售加的门店-RETAIL_PLUS_STORE ; 商户中心门店-PLACE_STORE ;  门店设备号-STORE_DEVICE
	_storeIdType string
	// 门店ID/设备号
	_storeId string
	// 商品类目ID
	_cid int64
	// 品牌ID
	_brandId int64
	// 商品状态: ON_SALE-在售 ; OFF_SALE-已下架 ; ALL-全部
	_status string
	// 商品类型: STORE_GOODS-经销/现货 ; SUPPLIER_GOODS-代销/网直供 ; TAOKE-淘宝客 ; ALL-全部商品
	_type string
	// 商品名称(支持模糊查询)
	_title string
	// 查询开始时间
	_startModified string
	// 查询结束时间
	_endModified string
	// 分页的页码
	_pageNo int64
	// 分页时每页的数量
	_pageSize int64
	// 商品的来源：0-从零售加采购的商品；1-商户线下导入的商品
	_itemType int64
}

// New
