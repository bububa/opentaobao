package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增电子价签商家门店接口 APIResponse
taobao.uscesl.biz.store.insert

新增电子价签商家门店接口
*/
type TaobaoUsceslBizStoreInsertAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizStoreInsertResponse
}

type TaobaoUsceslBizStoreInsertResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_store_insert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
