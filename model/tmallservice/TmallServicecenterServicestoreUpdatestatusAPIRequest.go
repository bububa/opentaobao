package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterServicestoreUpdatestatusAPIRequest
网点/门店状态修改 API请求
tmall.servicecenter.servicestore.updatestatus

修改网点/门店状态 */
type TmallServicecenterServicestoreUpdatestatusAPIRequest struct {
	model.Params
	// 门店id
	_id int64
	// 状态。1 营业，0歇业，-1彻底关店
	_status int64
	// 业务类型。不同业务传不同的值
	_bizType string
}

// New
