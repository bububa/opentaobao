package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreturnpackagereportAPIResponse 退货包裹状态通知接口 API返回值
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
type TaobaoqimenreturnpackagereportAPIResponse struct {
	model.CommonResponse
	TaobaoqimenreturnpackagereportAPIResponseModel
}

// TaobaoqimenreturnpackagereportAPIResponseModel is 退货包裹状态通知接口 成功返回结果
type TaobaoqimenreturnpackagereportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_returnpackage_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenreturnpackagereportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
