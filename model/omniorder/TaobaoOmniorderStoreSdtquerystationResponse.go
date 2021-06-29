package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
速店通查询站点信息 APIResponse
taobao.omniorder.store.sdtquerystation

速店通查询站点信息
*/
type TaobaoOmniorderStoreSdtquerystationAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreSdtquerystationResponse
}

type TaobaoOmniorderStoreSdtquerystationResponse struct {
    XMLName xml.Name `xml:"omniorder_store_sdtquerystation_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniorderStoreSdtquerystationResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
