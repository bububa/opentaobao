package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrOrderLogisInfoAPIRequest
区域零售订单获取取件码 API请求
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。 */
type TmallNrOrderLogisInfoAPIRequest struct {
	model.Params
	// 卖家ID
	_sellerId int64
	// 主订单号
	_mainOrderIds []int64
	// 来源标识
	_channel string
}

// New
