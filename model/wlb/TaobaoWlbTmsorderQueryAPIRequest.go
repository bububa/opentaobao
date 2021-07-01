package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbTmsorderQueryAPIRequest
通过物流订单编号查询物流信息 API请求
taobao.wlb.tmsorder.query

通过物流订单编号分页查询物流信息 */
type TaobaoWlbTmsorderQueryAPIRequest struct {
	model.Params
	// 物流订单编号
	_orderCode string
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// New
