package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradeServiceinfoWriteAPIRequest
订单服务信息写入接口 API请求
alitrip.travel.trade.serviceinfo.write

订单服务信息写入接口 */
type AlitripTravelTradeServiceinfoWriteAPIRequest struct {
	model.Params
	// 根据模版要求传递对应的订单服务信息，当涉及多值时，用英文分号隔开";"，目前api暂时不支持文件上传，只支持链接；链接必须外网能访问
	_tipValue string
	// 对应的订单信息
	_subTcOrderId int64
}

// New
