package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
switchstatus.update API返回值 
taobao.omniorder.store.switchstatus.update

变更门店发货、门店自提状态
*/
type TaobaoOmniorderStoreSwitchstatusUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel
}

// switchstatus.update 成功返回结果
type TaobaoOmniorderStoreSwitchstatusUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"omniorder_store_switchstatus_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniorderStoreSwitchstatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
