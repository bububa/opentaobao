package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRpReturngoodsAgreeAPIRequest
卖家同意退货 API请求
taobao.rp.returngoods.agree

卖家同意退货，支持淘宝和天猫的订单。 */
type TaobaoRpReturngoodsAgreeAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 卖家姓名，淘宝退款为必填项。
	_name string
	// 卖家提供的退货地址，淘宝退款为必填项。
	_address string
	// 卖家提供的退货地址的邮编，淘宝退款为必填项。
	_post string
	// 卖家座机，淘宝退款为必填项。
	_tel string
	// 卖家手机，淘宝退款为必填项。
	_mobile string
	// 卖家退货留言，天猫退款为必填项。
	_remark string
	// 售中：onsale，售后：aftersale，天猫退款为必填项。
	_refundPhase string
	// 退款版本号，天猫退款为必填项。
	_refundVersion int64
	// 卖家收货地址编号，天猫淘宝退款都为必填项。
	_sellerAddressId int64
	// 邮费承担方，买家承担值为1，卖家承担值为0
	_postFeeBearRole int64
	// 是否虚拟退货，可选项
	_virtualReturnGoods bool
}

// New
