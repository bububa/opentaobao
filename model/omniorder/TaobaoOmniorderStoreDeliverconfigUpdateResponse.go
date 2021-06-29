package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改门店发货配置内容 API返回值 
taobao.omniorder.store.deliverconfig.update

修改门店发货配置内容
*/
type TaobaoOmniorderStoreDeliverconfigUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreDeliverconfigUpdateResponse
}

// 修改门店发货配置内容 成功返回结果
type TaobaoOmniorderStoreDeliverconfigUpdateResponse struct {
    XMLName xml.Name `xml:"omniorder_store_deliverconfig_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniorderStoreDeliverconfigUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}