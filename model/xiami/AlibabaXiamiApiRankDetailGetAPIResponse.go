package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
排行榜详情 API返回值 
alibaba.xiami.api.rank.detail.get

虾米排行榜详情数据
*/
type AlibabaXiamiApiRankDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiRankDetailGetAPIResponseModel
}

// 排行榜详情 成功返回结果
type AlibabaXiamiApiRankDetailGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_rank_detail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // resultObj
    ResultObj   *BillboardItemVO `json:"result_obj,omitempty" xml:"result_obj,omitempty"`
}
