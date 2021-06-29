package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
实体动态更新 APIResponse
alibaba.ailabs.aligenie.skill.entity.import

根据用户上传的实体信息，进行制定技能实体的动态变更
*/
type AlibabaAilabsAligenieSkillEntityImportAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieSkillEntityImportResponse
}

type AlibabaAilabsAligenieSkillEntityImportResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_skill_entity_import_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
