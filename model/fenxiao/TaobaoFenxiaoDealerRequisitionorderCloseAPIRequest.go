package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest
供应商/分销商关闭采购申请/经销采购单 API请求
taobao.fenxiao.dealer.requisitionorder.close

供应商或分销商关闭采购申请/经销采购单 */
type TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest struct {
	model.Params
	// 经销采购单编号
	_dealerOrderId int64
	// 关闭原因：<br/>1：长时间无法联系到分销商，取消交易。<br/>2：分销商错误提交申请，取消交易。<br/>3：缺货，近段时间都无法发货。<br/>4：分销商恶意提交申请单。<br/>5：其他原因。
	_reason int64
	// 关闭详细原因，字数5-200字
	_reasonDetail string
}

// New
