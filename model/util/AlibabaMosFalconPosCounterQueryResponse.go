package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
云POS查看专柜属性 APIResponse
alibaba.mos.falcon.pos.counter.query

银泰商业获取专柜是否支持小数等属性查看
*/
type AlibabaMosFalconPosCounterQueryAPIResponse struct {
    model.CommonResponse
    AlibabaMosFalconPosCounterQueryResponse
}

type AlibabaMosFalconPosCounterQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_falcon_pos_counter_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaMosFalconPosCounterQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
