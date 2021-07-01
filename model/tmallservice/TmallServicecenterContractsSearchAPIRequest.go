package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterContractsSearchAPIRequest
获取合同类的服务工单信息 API请求
tmall.servicecenter.contracts.search

获取合同类的服务工单信息 */
type TmallServicecenterContractsSearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// New
