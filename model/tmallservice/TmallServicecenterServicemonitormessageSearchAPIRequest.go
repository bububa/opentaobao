package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicemonitormessageSearchAPIRequest
根据时间段查询服务商的服务预警消息列表(15分钟内) API请求
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内) */
type TmallServicecenterServicemonitormessageSearchAPIRequest struct {
	model.Params
	// 开始时间long值
	_start int64
	// 结束时间long值，与start相差15分钟
	_end int64
}

// New
