package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse 新房楼栋修改e码 API返回值
// alibaba.alihouse.newhome.project.building.ecode.update
//
// 新房楼栋修改e码
type AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponseModel is 新房楼栋修改e码 成功返回结果
type AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_building_ecode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse
func GetAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse() *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse 将 AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse(v *AlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectBuildingEcodeUpdateAPIResponse.Put(v)
}
