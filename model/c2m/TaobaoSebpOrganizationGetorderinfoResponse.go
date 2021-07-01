package c2m

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺机构订单信息 API返回值 
taobao.sebp.organization.getorderinfo

淘小铺合作机构获取机构订单信息，用于机构结算使用
*/
type TaobaoSebpOrganizationGetorderinfoAPIResponse struct {
    model.CommonResponse
    TaobaoSebpOrganizationGetorderinfoResponse
}

// 淘小铺机构订单信息 成功返回结果
type TaobaoSebpOrganizationGetorderinfoResponse struct {
    XMLName xml.Name `xml:"sebp_organization_getorderinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
