package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店发货配置内容 APIResponse
taobao.omniorder.store.deliverconfig.get

查询门店发货配置内容
*/
type TaobaoOmniorderStoreDeliverconfigGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreDeliverconfigGetResponse
}

type TaobaoOmniorderStoreDeliverconfigGetResponse struct {
    XMLName xml.Name `xml:"omniorder_store_deliverconfig_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreDeliverconfigGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
