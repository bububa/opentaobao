package shop

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新店铺基本信息 APIResponse
taobao.shop.update

目前只支持标题、公告和描述的更新
*/
type TaobaoShopUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoShopUpdateResponse
}

type TaobaoShopUpdateResponse struct {
    XMLName xml.Name `xml:"shop_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 店铺信息
    
    Shop   *Shop `json:"shop,omitempty" xml:"shop,omitempty"`

    
}
