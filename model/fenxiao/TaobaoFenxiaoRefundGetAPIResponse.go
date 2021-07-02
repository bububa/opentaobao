package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRefundGetAPIResponse 查询采购单退款信息 API返回值
// taobao.fenxiao.refund.get
//
// 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
type TaobaoFenxiaoRefundGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoRefundGetAPIResponseModel
}

// TaobaoFenxiaoRefundGetAPIResponseModel is 查询采购单退款信息 成功返回结果
type TaobaoFenxiaoRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款详情
	RefundDetail *RefundDetail `json:"refund_detail,omitempty" xml:"refund_detail,omitempty"`
}
