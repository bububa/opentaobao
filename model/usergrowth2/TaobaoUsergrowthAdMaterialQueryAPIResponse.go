package usergrowth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsergrowthAdMaterialQueryAPIResponse 素材审核结果查询 API返回值
// taobao.usergrowth.ad.material.query
//
// 素材审核结果查询
type TaobaoUsergrowthAdMaterialQueryAPIResponse struct {
	model.CommonResponse
	TaobaoUsergrowthAdMaterialQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsergrowthAdMaterialQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsergrowthAdMaterialQueryAPIResponseModel).Reset()
}

// TaobaoUsergrowthAdMaterialQueryAPIResponseModel is 素材审核结果查询 成功返回结果
type TaobaoUsergrowthAdMaterialQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"usergrowth_ad_material_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 参数错误
	ResponseCode int64 `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 返回信息
	Data *CreativeAuditResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求结果
	Successful bool `json:"successful,omitempty" xml:"successful,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsergrowthAdMaterialQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Message = ""
	m.ResponseCode = 0
	m.Data = nil
	m.Successful = false
}

var poolTaobaoUsergrowthAdMaterialQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsergrowthAdMaterialQueryAPIResponse)
	},
}

// GetTaobaoUsergrowthAdMaterialQueryAPIResponse 从 sync.Pool 获取 TaobaoUsergrowthAdMaterialQueryAPIResponse
func GetTaobaoUsergrowthAdMaterialQueryAPIResponse() *TaobaoUsergrowthAdMaterialQueryAPIResponse {
	return poolTaobaoUsergrowthAdMaterialQueryAPIResponse.Get().(*TaobaoUsergrowthAdMaterialQueryAPIResponse)
}

// ReleaseTaobaoUsergrowthAdMaterialQueryAPIResponse 将 TaobaoUsergrowthAdMaterialQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsergrowthAdMaterialQueryAPIResponse(v *TaobaoUsergrowthAdMaterialQueryAPIResponse) {
	v.Reset()
	poolTaobaoUsergrowthAdMaterialQueryAPIResponse.Put(v)
}
