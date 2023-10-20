package tbk

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdReportGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScUcrowdReportGetAPIResponseModel).Reset()
}

// TaobaoTbkScUcrowdReportGetAPIResponseModel is 淘宝客-服务商-人群推广效果 成功返回结果
type TaobaoTbkScUcrowdReportGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_ucrowd_report_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScUcrowdReportGetRpcResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScUcrowdReportGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScUcrowdReportGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdReportGetAPIResponse)
	},
}

// GetTaobaoTbkScUcrowdReportGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScUcrowdReportGetAPIResponse
func GetTaobaoTbkScUcrowdReportGetAPIResponse() *TaobaoTbkScUcrowdReportGetAPIResponse {
	return poolTaobaoTbkScUcrowdReportGetAPIResponse.Get().(*TaobaoTbkScUcrowdReportGetAPIResponse)
}

// ReleaseTaobaoTbkScUcrowdReportGetAPIResponse 将 TaobaoTbkScUcrowdReportGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScUcrowdReportGetAPIResponse(v *TaobaoTbkScUcrowdReportGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScUcrowdReportGetAPIResponse.Put(v)
}
