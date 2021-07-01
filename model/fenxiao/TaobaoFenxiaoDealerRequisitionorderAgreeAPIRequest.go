package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest
供应商/分销商通过采购申请/经销采购单申请 API请求
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核 */
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest struct {
	model.Params
	// 采购申请/经销采购单编号
	_dealerOrderId int64
}

// New
