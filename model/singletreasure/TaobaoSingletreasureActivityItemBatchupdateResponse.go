package singletreasure

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量修改商品接口 APIResponse
taobao.singletreasure.activity.item.batchupdate

批量修改商品优惠接口
*/
type TaobaoSingletreasureActivityItemBatchupdateAPIResponse struct {
    model.CommonResponse
    TaobaoSingletreasureActivityItemBatchupdateResponse
}

type TaobaoSingletreasureActivityItemBatchupdateResponse struct {
    XMLName xml.Name `xml:"singletreasure_activity_item_batchupdate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoSingletreasureActivityItemBatchupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
