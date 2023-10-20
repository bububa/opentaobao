package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsgeneralconsignAPIResponse 一般进口发货 API返回值
// taobao.wlb.imports.general.consign
//
// 将订单信息发送到菜鸟海外转运仓；
// 业务规则：
// 1）交易订单为待发货状态。
// 2）单笔订单多个商品，交易金额不能大于1000人民币。
type TaobaowlbimportsgeneralconsignAPIResponse struct {
	model.CommonResponse
	TaobaowlbimportsgeneralconsignAPIResponseModel
}

// TaobaowlbimportsgeneralconsignAPIResponseModel is 一般进口发货 成功返回结果
type TaobaowlbimportsgeneralconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_general_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发货成功后的物流订单编号
	LgorderCode string `json:"lgorder_code,omitempty" xml:"lgorder_code,omitempty"`
	// 业务错误描述
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 业务错误编码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 是否发货成功,true:成功，false：失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
