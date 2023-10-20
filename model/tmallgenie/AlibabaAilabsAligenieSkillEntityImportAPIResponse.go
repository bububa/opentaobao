package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieSkillEntityImportAPIResponse 实体动态更新 API返回值
// alibaba.ailabs.aligenie.skill.entity.import
//
// 根据用户上传的实体信息，进行制定技能实体的动态变更
type AlibabaAilabsAligenieSkillEntityImportAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieSkillEntityImportAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieSkillEntityImportAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieSkillEntityImportAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieSkillEntityImportAPIResponseModel is 实体动态更新 成功返回结果
type AlibabaAilabsAligenieSkillEntityImportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_skill_entity_import_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieSkillEntityImportAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsAligenieSkillEntityImportAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieSkillEntityImportAPIResponse)
	},
}

// GetAlibabaAilabsAligenieSkillEntityImportAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieSkillEntityImportAPIResponse
func GetAlibabaAilabsAligenieSkillEntityImportAPIResponse() *AlibabaAilabsAligenieSkillEntityImportAPIResponse {
	return poolAlibabaAilabsAligenieSkillEntityImportAPIResponse.Get().(*AlibabaAilabsAligenieSkillEntityImportAPIResponse)
}

// ReleaseAlibabaAilabsAligenieSkillEntityImportAPIResponse 将 AlibabaAilabsAligenieSkillEntityImportAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieSkillEntityImportAPIResponse(v *AlibabaAilabsAligenieSkillEntityImportAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieSkillEntityImportAPIResponse.Put(v)
}
