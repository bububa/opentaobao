package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterSpserviceorderQueryAPIRequest
服务单列表查询 API请求
tmall.servicecenter.spserviceorder.query

查询服务单列表 */
type TmallServicecenterSpserviceorderQueryAPIRequest struct {
	model.Params
	// 交易主订单id
	_parentBizOrderId int64
}

// New
