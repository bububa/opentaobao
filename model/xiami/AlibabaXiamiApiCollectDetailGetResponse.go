package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐精选集详情接口 APIResponse
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口
*/
type AlibabaXiamiApiCollectDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaXiamiApiCollectDetailGetResponse `json:"alibaba_xiami_api_collect_detail_get_response,omitempty"` 
    AlibabaXiamiApiCollectDetailGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiCollectDetailGetResponse struct {

    // 精选集资料和对应歌曲列表
    
    Data  *struct {
        CollectDetail  *CollectDetail `json:"collect_detail,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaXiamiApiCollectDetailGetResponse struct {

    // 精选集资料和对应歌曲列表
    Data   *CollectDetail `json:"data,omitempty"`

}
