package lstlogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-第三方物流公司列表 APIResponse
alibaba.lst.logistics.thirdpart.company.list

异地云仓发货时，需填写的第三方物流公司列表
*/
type AlibabaLstLogisticsThirdpartCompanyListAPIResponse struct {
    model.CommonResponse
    AlibabaLstLogisticsThirdpartCompanyListResponse
}

type AlibabaLstLogisticsThirdpartCompanyListResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_logistics_thirdpart_company_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回内容
    
    ContentList   []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
    
    
}
