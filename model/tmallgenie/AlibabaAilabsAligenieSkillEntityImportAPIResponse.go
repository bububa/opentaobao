package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实体动态更新 API返回值 
alibaba.ailabs.aligenie.skill.entity.import

根据用户上传的实体信息，进行制定技能实体的动态变更
*/
type AlibabaAilabsAligenieSkillEntityImportAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieSkillEntityImportAPIResponseModel
}

// 实体动态更新 成功返回结果
type AlibabaAilabsAligenieSkillEntityImportAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_skill_entity_import_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回包装类
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
