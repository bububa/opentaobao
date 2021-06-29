package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自提配置修改 APIResponse
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreCollectconfigUpdateResponse
}

type TaobaoOmniorderStoreCollectconfigUpdateResponse struct {
    XMLName xml.Name `xml:"omniorder_store_collectconfig_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreCollectconfigUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
