package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
西安杨森同步用户行为接口 API返回值 
alibaba.alihealth.drugcode.user.data

西安杨森同步用户行为接口
*/
type AlibabaAlihealthDrugcodeUserDataAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeUserDataResponse
}

// 西安杨森同步用户行为接口 成功返回结果
type AlibabaAlihealthDrugcodeUserDataResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_user_data_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaAlihealthDrugcodeUserDataResult `json:"result,omitempty" xml:"result,omitempty"`
}
