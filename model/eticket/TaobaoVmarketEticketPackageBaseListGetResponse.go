package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id，获取关联的所有包 API返回值 
taobao.vmarket.eticket.package.base.list.get

根据卖家id，获取关联的所有包
*/
type TaobaoVmarketEticketPackageBaseListGetAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketPackageBaseListGetResponse
}

// 根据卖家id，获取关联的所有包 成功返回结果
type TaobaoVmarketEticketPackageBaseListGetResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_package_base_list_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果
    Result   *PackageResult `json:"result,omitempty" xml:"result,omitempty"`
}
