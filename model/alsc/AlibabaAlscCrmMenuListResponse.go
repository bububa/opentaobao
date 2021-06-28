package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取特价菜单 APIResponse
alibaba.alsc.crm.menu.list

获取特价菜单
*/
type AlibabaAlscCrmMenuListAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmMenuListResponse
}

type AlibabaAlscCrmMenuListResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_menu_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
