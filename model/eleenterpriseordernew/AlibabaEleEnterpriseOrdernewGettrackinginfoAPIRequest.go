package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest
订单配送信息跟踪 API请求
alibaba.ele.enterprise.ordernew.gettrackinginfo

订单配送信息跟踪 */
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 用户手机号
	_phone string
}

// New
