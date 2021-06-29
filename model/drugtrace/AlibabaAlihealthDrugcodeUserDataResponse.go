package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
西安杨森同步用户行为接口 APIResponse
alibaba.alihealth.drugcode.user.data

西安杨森同步用户行为接口
*/
type AlibabaAlihealthDrugcodeUserDataAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugcodeUserDataResponse
}

type AlibabaAlihealthDrugcodeUserDataResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drugcode_user_data_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugcodeUserDataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
