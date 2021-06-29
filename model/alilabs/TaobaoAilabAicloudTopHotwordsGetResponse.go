package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取热词 APIResponse
taobao.ailab.aicloud.top.hotwords.get

获取ASR热词
*/
type TaobaoAilabAicloudTopHotwordsGetAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopHotwordsGetResponse
}

type TaobaoAilabAicloudTopHotwordsGetResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_hotwords_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // baseresult
    
    Baseresult   *BaseResult `json:"baseresult,omitempty" xml:"baseresult,omitempty"`

    
}
