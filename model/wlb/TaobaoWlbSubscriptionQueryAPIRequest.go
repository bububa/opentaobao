package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbSubscriptionQueryAPIRequest
查询商家定购的所有服务 API请求
taobao.wlb.subscription.query

查询商家定购的所有服务,可通过入参状态来筛选 */
type TaobaoWlbSubscriptionQueryAPIRequest struct {
	model.Params
	// 状态 <br/>AUDITING 1-待审核; <br/>CANCEL 2-撤销 ;<br/>CHECKED 3-审核通过 ;<br/>FAILED 4-审核未通过 ;<br/>SYNCHRONIZING 5-同步中;<br/>只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。
	_status string
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// New
