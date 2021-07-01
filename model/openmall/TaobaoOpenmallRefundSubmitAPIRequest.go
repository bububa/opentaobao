package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenmallRefundSubmitAPIRequest
提交OpenMall退款单物流 API请求
taobao.openmall.refund.submit

提交OpenMall退款单物流 */
type TaobaoOpenmallRefundSubmitAPIRequest struct {
	model.Params
	// 渠道
	_distributor string
	// 物流公司编码
	_logisticsCompanyCode string
	// 物流公司名称
	_logisticsCompanyName string
	// 快递单号
	_logisticsNo string
	// 退款单ID
	_refundId int64
}

// New
