package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewGetstatusAPIRequest
订单状态查询接口 API请求
alibaba.ele.enterprise.ordernew.getstatus

订单状态查询接口 */
type AlibabaEleEnterpriseOrdernewGetstatusAPIRequest struct {
	model.Params
	// 订单号
	_elemeOrderId string
}

// New
