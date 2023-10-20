package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoYphRefundsGetAPIResponse 一盘货商家批量查询退款单信息 API返回值
// taobao.fenxiao.yph.refunds.get
//
// 一盘货商家批量查询退款单信息
type TaobaoFenxiaoYphRefundsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoYphRefundsGetAPIResponseModel
}

// TaobaoFenxiaoYphRefundsGetAPIResponseModel is 一盘货商家批量查询退款单信息 成功返回结果
type TaobaoFenxiaoYphRefundsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_yph_refunds_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款单详情列表
	RefundDetailList []RefundDetailList `json:"refund_detail_list,omitempty" xml:"refund_detail_list>refund_detail_list,omitempty"`
	// 操作时间
	OptTime string `json:"opt_time,omitempty" xml:"opt_time,omitempty"`
	// 返回结果错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回结果错误信息
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 退款单查询总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
