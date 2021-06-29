package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构上下级关系 APIResponse
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息
*/
type TaobaoSebpOrganizationGetinviteinfoAPIResponse struct {
    model.CommonResponse
    TaobaoSebpOrganizationGetinviteinfoResponse
}

type TaobaoSebpOrganizationGetinviteinfoResponse struct {
    XMLName xml.Name `xml:"sebp_organization_getinviteinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoSebpOrganizationGetinviteinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
