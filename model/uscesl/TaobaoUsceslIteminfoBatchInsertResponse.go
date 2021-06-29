package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
按商家批量写入商品接口 APIResponse
taobao.uscesl.iteminfo.batch.insert

【电子价签】支持按照商家-门店维度批量写入商品数据
*/
type TaobaoUsceslIteminfoBatchInsertAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslIteminfoBatchInsertResponse
}

type TaobaoUsceslIteminfoBatchInsertResponse struct {
    XMLName xml.Name `xml:"uscesl_iteminfo_batch_insert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
