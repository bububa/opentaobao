package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜详情 APIResponse
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
type AlibabaXiamiApiRankDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiRankDetailGetResponse
}

type AlibabaXiamiApiRankDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_rank_detail_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // resultObj
    
    ResultObj   *BillboardItemVO `json:"result_obj,omitempty" xml:"result_obj,omitempty"`

    
}
