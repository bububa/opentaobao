package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoRefundQueryAPIResponse 批量查询采购退款 API返回值
// taobao.fenxiao.refund.query
//
// 供应商按查询条件批量查询代销采购退款
type TaobaoFenxiaoRefundQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoRefundQueryAPIResponseModel
}

// TaobaoFenxiaoRefundQueryAPIResponseModel is 批量查询采购退款 成功返回结果
type TaobaoFenxiaoRefundQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款列表
	RefundList []RefundDetail `json:"refund_list,omitempty" xml:"refund_list>refund_detail,omitempty"`
	// 按查询条件查到的记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
