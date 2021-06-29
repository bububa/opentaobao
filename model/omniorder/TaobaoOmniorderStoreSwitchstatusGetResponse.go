package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.get APIResponse
taobao.omniorder.store.switchstatus.get

查询门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSwitchstatusGetResponse
}

type TaobaoOmniorderStoreSwitchstatusGetResponse struct {
    XMLName xml.Name `xml:"omniorder_store_switchstatus_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreSwitchstatusGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
