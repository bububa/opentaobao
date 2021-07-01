package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterAnomalyrecourseQuerybyidAPIRequest
根据一键求助id查询指定服务商的一键求助单 API请求
tmall.servicecenter.anomalyrecourse.querybyid

根据一键求助id查询指定服务商的一键求助单 */
type TmallServicecenterAnomalyrecourseQuerybyidAPIRequest struct {
	model.Params
	// 一键求助的id
	_anomalyRecourseId int64
}

// New
