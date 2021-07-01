package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeShippingaddressUpdateAPIRequest
更改交易的收货地址 API请求
taobao.trade.shippingaddress.update

只能更新一笔交易里面的买家收货地址 <br/>只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 <br/>更新后的发货地址可以通过taobao.trade.fullinfo.get查到 <br/>参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节） */
type TaobaoTradeShippingaddressUpdateAPIRequest struct {
	model.Params
	// 交易编号。
	_tid int64
	// 收货人全名。最大长度为50个字节。
	_receiverName string
	// 固定电话。最大长度为30个字节。
	_receiverPhone string
	// 移动电话。最大长度为11个字节。
	_receiverMobile string
	// 省份。最大长度为32个字节。如：浙江
	_receiverState string
	// 城市。最大长度为32个字节。如：杭州
	_receiverCity string
	// 区/县。最大长度为32个字节。如：西湖区
	_receiverDistrict string
	// 收货地址。最大长度为228个字节。
	_receiverAddress string
	// 邮政编码。必须由6个数字组成。
	_receiverZip string
	// 四级地址。最大长度为32个字节。如：五常街道
	_receiverTown string
}

// New
