package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自提配置修改 API返回值 
taobao.omniorder.store.collectconfig.update

修改门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel
}

// 门店自提配置修改 成功返回结果
type TaobaoOmniorderStoreCollectconfigUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"omniorder_store_collectconfig_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniorderStoreCollectconfigUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
