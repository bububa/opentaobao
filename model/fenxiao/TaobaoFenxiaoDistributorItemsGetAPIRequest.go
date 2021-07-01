package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDistributorItemsGetAPIRequest
查询商品下载记录 API请求
taobao.fenxiao.distributor.items.get

供应商查询分销商商品下载记录。 */
type TaobaoFenxiaoDistributorItemsGetAPIRequest struct {
	model.Params
	// 分销商ID 。
	_distributorId int64
	// 设置开始时间。空为不设置。
	_startModified string
	// 设置结束时间,空为不设置。
	_endModified string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// 产品ID
	_productId int64
}

// New
