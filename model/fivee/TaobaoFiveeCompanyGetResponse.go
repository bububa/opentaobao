package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商信息 APIResponse
taobao.fivee.company.get

资质共享平台查询商信息
*/
type TaobaoFiveeCompanyGetAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeCompanyGetResponse
}

type TaobaoFiveeCompanyGetResponse struct {
    XMLName xml.Name `xml:"fivee_company_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFiveeCompanyGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
