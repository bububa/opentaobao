package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
家装新零售主商品同步至阿里 APIResponse
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里
*/
type TmallNrtItemMainSynchronizeAPIResponse struct {
    model.CommonResponse
    TmallNrtItemMainSynchronizeResponse
}

type TmallNrtItemMainSynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_item_main_synchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    TmallNrtItemMainSynchronize   *TmallNrtItemMainSynchronizeResultDo `json:"tmall_nrt_item_main_synchronize,omitempty" xml:"tmall_nrt_item_main_synchronize,omitempty"`

    
}
