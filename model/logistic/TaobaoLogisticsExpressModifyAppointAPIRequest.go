package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoLogisticsExpressModifyAppointAPIRequest
快递改约api API请求
taobao.logistics.express.modify.appoint

商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期 */
type TaobaoLogisticsExpressModifyAppointAPIRequest struct {
	model.Params
	// 改约请求对象
	_expressModifyAppointTopRequest *ExpressModifyAppointTopRequestDto
}

// New
