package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewCancelAPIRequest
订单取消 API请求
alibaba.ele.enterprise.ordernew.cancel

订单取消 */
type AlibabaEleEnterpriseOrdernewCancelAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
	// 取消原因(取消时提供)
	_reason string
}

// New
