package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
虾米音乐精选集详情接口 APIResponse
alibaba.xiami.api.collect.detail.get

虾米音乐精选集详情接口
*/
type AlibabaXiamiApiCollectDetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaXiamiApiCollectDetailGetResponse
}

type AlibabaXiamiApiCollectDetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_xiami_api_collect_detail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 精选集资料和对应歌曲列表
    
    Data   *CollectDetail `json:"data,omitempty" xml:"data,omitempty"`

    
}
