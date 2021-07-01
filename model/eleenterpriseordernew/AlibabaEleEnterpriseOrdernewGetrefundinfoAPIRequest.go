package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest
退单和申诉 API请求
alibaba.ele.enterprise.ordernew.getrefundinfo

退单和申诉 */
type AlibabaEleEnterpriseOrdernewGetrefundinfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
}

// New
