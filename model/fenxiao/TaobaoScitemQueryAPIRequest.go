package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemQueryAPIRequest
查询后端商品 API请求
taobao.scitem.query

查询后端商品 */
type TaobaoScitemQueryAPIRequest struct {
	model.Params
	// 商品名称
	_itemName string
	// 商家给商品的一个编码
	_outerCode string
	// 条形码
	_barCode string
	// ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION
	_itemType int64
	// 仓库编码
	_wmsCode string
	// 当前页码数
	_pageIndex int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// New
