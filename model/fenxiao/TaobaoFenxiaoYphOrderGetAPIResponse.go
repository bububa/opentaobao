package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoYphOrderGetAPIResponse 一盘货商家单个查询采购单信息 API返回值
// taobao.fenxiao.yph.order.get
//
// 一盘货商家单个查询采购单信息
type TaobaoFenxiaoYphOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoYphOrderGetAPIResponseModel
}

// TaobaoFenxiaoYphOrderGetAPIResponseModel is 一盘货商家单个查询采购单信息 成功返回结果
type TaobaoFenxiaoYphOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_yph_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回错误信息
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作时间
	OptTime string `json:"opt_time,omitempty" xml:"opt_time,omitempty"`
	// 返回的采购单对象
	PurchaseOrder *PurchaseOrder `json:"purchase_order,omitempty" xml:"purchase_order,omitempty"`
}
