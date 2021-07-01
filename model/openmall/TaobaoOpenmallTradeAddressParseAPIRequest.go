package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallTradeAddressParseAPIRequest
openmall服务地址区域码解析 API请求
taobao.openmall.trade.address.parse

openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息 */
type TaobaoOpenmallTradeAddressParseAPIRequest struct {
	model.Params
	// 需解析的地址信息，建议只传地址选择器中的省市区，街道门牌号等用户手动输入数据不传
	_rawAddress string
	// 渠道商分销者淘宝账号
	_distributor string
}

// New
