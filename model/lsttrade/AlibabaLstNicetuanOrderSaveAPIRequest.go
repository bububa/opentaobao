package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstNicetuanOrderSaveAPIRequest
十荟团订单同步至零售通 API请求
alibaba.lst.nicetuan.order.save

十荟团订单同步至零售通，十荟团单向写到零售通 */
type AlibabaLstNicetuanOrderSaveAPIRequest struct {
	model.Params
	// 订单数据
	_param *NicetuanMainOrderParam
}

// New
