package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.update APIResponse
taobao.omniorder.store.switchstatus.update

变更门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSwitchstatusUpdateResponse
}

type TaobaoOmniorderStoreSwitchstatusUpdateResponse struct {
    XMLName xml.Name `xml:"omniorder_store_switchstatus_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreSwitchstatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
