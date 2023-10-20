package cmns

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosServiceCmnsCoaMessageresultGetAPIResponse CMNS消息发送到达查询 API返回值
// yunos.service.cmns.coa.messageresult.get
//
// CMNS消息发送到达查询,根据消息ID查询，仅能查询该appKey所发送的消息
type YunosServiceCmnsCoaMessageresultGetAPIResponse struct {
	model.CommonResponse
	YunosServiceCmnsCoaMessageresultGetAPIResponseModel
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaMessageresultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosServiceCmnsCoaMessageresultGetAPIResponseModel).Reset()
}

// YunosServiceCmnsCoaMessageresultGetAPIResponseModel is CMNS消息发送到达查询 成功返回结果
type YunosServiceCmnsCoaMessageresultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_service_cmns_coa_messageresult_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口查询出错提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 200表示查询成功
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 具体的消息返回值
	Data *MessageResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YunosServiceCmnsCoaMessageresultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Status = 0
	m.Data = nil
}

var poolYunosServiceCmnsCoaMessageresultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosServiceCmnsCoaMessageresultGetAPIResponse)
	},
}

// GetYunosServiceCmnsCoaMessageresultGetAPIResponse 从 sync.Pool 获取 YunosServiceCmnsCoaMessageresultGetAPIResponse
func GetYunosServiceCmnsCoaMessageresultGetAPIResponse() *YunosServiceCmnsCoaMessageresultGetAPIResponse {
	return poolYunosServiceCmnsCoaMessageresultGetAPIResponse.Get().(*YunosServiceCmnsCoaMessageresultGetAPIResponse)
}

// ReleaseYunosServiceCmnsCoaMessageresultGetAPIResponse 将 YunosServiceCmnsCoaMessageresultGetAPIResponse 保存到 sync.Pool
func ReleaseYunosServiceCmnsCoaMessageresultGetAPIResponse(v *YunosServiceCmnsCoaMessageresultGetAPIResponse) {
	v.Reset()
	poolYunosServiceCmnsCoaMessageresultGetAPIResponse.Put(v)
}
