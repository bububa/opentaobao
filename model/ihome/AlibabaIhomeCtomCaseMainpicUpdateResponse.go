package ihome

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
方案渲染图修改 APIResponse
alibaba.ihome.ctom.case.mainpic.update

用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。
*/
type AlibabaIhomeCtomCaseMainpicUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaIhomeCtomCaseMainpicUpdateResponse
}

type AlibabaIhomeCtomCaseMainpicUpdateResponse struct {
    XMLName xml.Name `xml:"alibaba_ihome_ctom_case_mainpic_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 根据站点名称查询产品
    
    ApiResult   *AlibabaIhomeCtomCaseMainpicUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
