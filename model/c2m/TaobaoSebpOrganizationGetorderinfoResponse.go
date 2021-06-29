package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构订单信息 APIResponse
taobao.sebp.organization.getorderinfo

淘小铺合作机构获取机构订单信息，用于机构结算使用
*/
type TaobaoSebpOrganizationGetorderinfoAPIResponse struct {
    model.CommonResponse
    TaobaoSebpOrganizationGetorderinfoResponse
}

type TaobaoSebpOrganizationGetorderinfoResponse struct {
    XMLName xml.Name `xml:"sebp_organization_getorderinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoSebpOrganizationGetorderinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
