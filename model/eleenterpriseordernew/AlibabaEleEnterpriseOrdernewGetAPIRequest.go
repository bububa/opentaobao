package eleenterpriseordernew

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewGetAPIRequest
查询订单详情 API请求
alibaba.ele.enterprise.ordernew.get

查询订单详情 */
type AlibabaEleEnterpriseOrdernewGetAPIRequest struct {
	model.Params
	// 饿了么订单ID
	_orderId string
	// 电话号码
	_phone string
}

// New
