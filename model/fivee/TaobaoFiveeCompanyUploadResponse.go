package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传商信息接口 APIResponse
taobao.fivee.company.upload

资质共享平台上传资质证照
*/
type TaobaoFiveeCompanyUploadAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeCompanyUploadResponse
}

type TaobaoFiveeCompanyUploadResponse struct {
    XMLName xml.Name `xml:"fivee_company_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // code
    
    CodeT   int64 `json:"code_t,omitempty" xml:"code_t,omitempty"`

    
    // 返回素材id
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 是否成功
    
    SuccessT   bool `json:"success_t,omitempty" xml:"success_t,omitempty"`

    
}
