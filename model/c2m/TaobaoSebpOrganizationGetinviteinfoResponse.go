package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构上下级关系 API返回值 
taobao.sebp.organization.getinviteinfo

机构人员获取机构上下级关系信息
*/
type TaobaoSebpOrganizationGetinviteinfoAPIResponse struct {
    model.CommonResponse
    TaobaoSebpOrganizationGetinviteinfoResponse
}

// 淘小铺机构上下级关系 成功返回结果
type TaobaoSebpOrganizationGetinviteinfoResponse struct {
    XMLName xml.Name `xml:"sebp_organization_getinviteinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoSebpOrganizationGetinviteinfoResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
