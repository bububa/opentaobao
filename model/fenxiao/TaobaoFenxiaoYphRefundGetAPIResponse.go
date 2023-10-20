package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoyphrefundgetAPIResponse 一盘货商家单个查询退款单信息 API返回值
// taobao.fenxiao.yph.refund.get
//
// 一盘货商家单个查询退款单信息
type TaobaofenxiaoyphrefundgetAPIResponse struct {
	model.CommonResponse
	TaobaofenxiaoyphrefundgetAPIResponseModel
}

// TaobaofenxiaoyphrefundgetAPIResponseModel is 一盘货商家单个查询退款单信息 成功返回结果
type TaobaofenxiaoyphrefundgetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_yph_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回值错误信息
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作时间
	OptTime string `json:"opt_time,omitempty" xml:"opt_time,omitempty"`
	// 退款单查询详情
	RefundDetail *RefundDetail `json:"refund_detail,omitempty" xml:"refund_detail,omitempty"`
}
