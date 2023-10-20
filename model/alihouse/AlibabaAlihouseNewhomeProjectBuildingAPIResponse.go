package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingAPIResponse 楼栋同步 API返回值
// alibaba.alihouse.newhome.project.building
//
// 楼栋同步
type AlibabaAlihouseNewhomeProjectBuildingAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectBuildingAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel is 楼栋同步 成功返回结果
type AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_building_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectBuildingResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectBuildingAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectBuildingAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectBuildingAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectBuildingAPIResponse
func GetAlibabaAlihouseNewhomeProjectBuildingAPIResponse() *AlibabaAlihouseNewhomeProjectBuildingAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectBuildingAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectBuildingAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingAPIResponse 将 AlibabaAlihouseNewhomeProjectBuildingAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectBuildingAPIResponse(v *AlibabaAlihouseNewhomeProjectBuildingAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectBuildingAPIResponse.Put(v)
}
