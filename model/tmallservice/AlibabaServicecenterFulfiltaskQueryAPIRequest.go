package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterFulfiltaskQueryAPIRequest
核销单查询 API请求
alibaba.servicecenter.fulfiltask.query

当系统生成核销单之后，需要派单到服务商，服务商根据核销里的服务信息和用户信息，给消费者提供服务 */
type AlibabaServicecenterFulfiltaskQueryAPIRequest struct {
	model.Params
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateStart string
	// 每页条数，默认50，最大50
	_pageSize int64
	// 核销单外部单号
	_outerId string
	// 时间段查询，核销单创建时间，时间段跨度不超过15分钟
	_gmtCreateEnd string
	// 查询第几页
	_currentPage int64
	// 核销单id列表
	_fulfilTaskIdList []int64
}

// New
