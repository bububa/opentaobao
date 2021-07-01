package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量写入商品信息接口 API返回值 
taobao.uscesl.iteminfo.batch.put

电子架签批量写入商品数据，用于电子价签展示
*/
type TaobaoUsceslIteminfoBatchPutAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslIteminfoBatchPutAPIResponseModel
}

// 批量写入商品信息接口 成功返回结果
type TaobaoUsceslIteminfoBatchPutAPIResponseModel struct {
    XMLName xml.Name `xml:"uscesl_iteminfo_batch_put_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
