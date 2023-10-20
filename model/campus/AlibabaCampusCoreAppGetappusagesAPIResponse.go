package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusCoreAppGetappusagesAPIResponse 根据应用ID获得应用所在的园区 API返回值
// alibaba.campus.core.app.getappusages
//
// 传入应用的id,  获得用户授权的园区
type AlibabaCampusCoreAppGetappusagesAPIResponse struct {
	model.CommonResponse
	AlibabaCampusCoreAppGetappusagesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusCoreAppGetappusagesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusCoreAppGetappusagesAPIResponseModel).Reset()
}

// AlibabaCampusCoreAppGetappusagesAPIResponseModel is 根据应用ID获得应用所在的园区 成功返回结果
type AlibabaCampusCoreAppGetappusagesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_core_app_getappusages_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusCoreAppGetappusagesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusCoreAppGetappusagesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusCoreAppGetappusagesAPIResponse)
	},
}

// GetAlibabaCampusCoreAppGetappusagesAPIResponse 从 sync.Pool 获取 AlibabaCampusCoreAppGetappusagesAPIResponse
func GetAlibabaCampusCoreAppGetappusagesAPIResponse() *AlibabaCampusCoreAppGetappusagesAPIResponse {
	return poolAlibabaCampusCoreAppGetappusagesAPIResponse.Get().(*AlibabaCampusCoreAppGetappusagesAPIResponse)
}

// ReleaseAlibabaCampusCoreAppGetappusagesAPIResponse 将 AlibabaCampusCoreAppGetappusagesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusCoreAppGetappusagesAPIResponse(v *AlibabaCampusCoreAppGetappusagesAPIResponse) {
	v.Reset()
	poolAlibabaCampusCoreAppGetappusagesAPIResponse.Put(v)
}
