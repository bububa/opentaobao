package tbk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdReportGetAPIResponse 淘宝客-服务商-人群推广效果 API返回值
// taobao.tbk.sc.ucrowd.report.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群的推广和转化效果数据。
type TaobaoTbkScUcrowdReportGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScUcrowdReportGetAPIResponseModel
}

// TaobaoTbkScUcrowdReportGetAPIResponseModel is 淘宝客-服务商-人群推广效果 成功返回结果
type TaobaoTbkScUcrowdReportGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_report_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdReportGetRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}
