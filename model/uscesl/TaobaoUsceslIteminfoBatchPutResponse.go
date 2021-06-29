package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量写入商品信息接口 APIResponse
taobao.uscesl.iteminfo.batch.put

电子架签批量写入商品数据，用于电子价签展示
*/
type TaobaoUsceslIteminfoBatchPutAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslIteminfoBatchPutResponse
}

type TaobaoUsceslIteminfoBatchPutResponse struct {
    XMLName xml.Name `xml:"uscesl_iteminfo_batch_put_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
