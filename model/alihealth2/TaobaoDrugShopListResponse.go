package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家外卖店列表 APIResponse
taobao.drug.shop.list

查询卖家外卖店列表
*/
type TaobaoDrugShopListAPIResponse struct {
    model.CommonResponse
    TaobaoDrugShopListResponse
}

type TaobaoDrugShopListResponse struct {
    XMLName xml.Name `xml:"drug_shop_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 数据结果集
    
    Result   *TakeoutShopPage `json:"result,omitempty" xml:"result,omitempty"`

    
}
