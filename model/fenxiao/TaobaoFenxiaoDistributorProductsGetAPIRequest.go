package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDistributorProductsGetAPIRequest
分销商查询产品信息 API请求
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息 */
type TaobaoFenxiaoDistributorProductsGetAPIRequest struct {
	model.Params
	// order_by
	_orderBy string
	// time_type
	_timeType string
	// 下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。
	_downloadStatus string
	// 分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营
	_tradeType string
	// 结束时间
	_endTime string
	// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
	_fields []string
	// 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
	_itemIds []int64
	// 产品线ID
	_productcatId int64
	// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
	_pids []int64
	// 开始修改时间
	_startTime string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// 供应商nick，分页查询时，必传
	_supplierNick string
}

// New
