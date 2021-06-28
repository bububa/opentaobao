package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
排行榜详情 APIResponse
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
type AlibabaXiamiApiRankDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaXiamiApiRankDetailGetResponse `json:"alibaba_xiami_api_rank_detail_get_response,omitempty"` 
    AlibabaXiamiApiRankDetailGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiRankDetailGetResponse struct {

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // resultObj
    
    ResultObj  *struct {
        BillboardItemVO  *BillboardItemVO `json:"billboard_item_vo,omitempty"`
    } `json:"result_obj,omitempty"`
    

}
*/

type AlibabaXiamiApiRankDetailGetResponse struct {

    // msgCode
    MsgCode   string `json:"msg_code,omitempty"`

    // resultObj
    ResultObj   *BillboardItemVO `json:"result_obj,omitempty"`

}
