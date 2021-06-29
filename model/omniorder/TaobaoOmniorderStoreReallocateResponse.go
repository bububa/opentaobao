package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
rellocate APIResponse
taobao.omniorder.store.reallocate

门店发货提供改派接口
*/
type TaobaoOmniorderStoreReallocateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreReallocateResponse
}

type TaobaoOmniorderStoreReallocateResponse struct {
    XMLName xml.Name `xml:"omniorder_store_reallocate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreReallocateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
