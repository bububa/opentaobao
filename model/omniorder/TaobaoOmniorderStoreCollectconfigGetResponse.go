package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询门店自提配置内容 API返回值 
taobao.omniorder.store.collectconfig.get

查询门店自提配置内容
*/
type TaobaoOmniorderStoreCollectconfigGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreCollectconfigGetResponse
}

// 查询门店自提配置内容 成功返回结果
type TaobaoOmniorderStoreCollectconfigGetResponse struct {
    XMLName xml.Name `xml:"omniorder_store_collectconfig_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniorderStoreCollectconfigGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
