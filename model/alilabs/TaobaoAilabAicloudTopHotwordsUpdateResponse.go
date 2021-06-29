package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新热词 API返回值 
taobao.ailab.aicloud.top.hotwords.update

更新ASR热词
*/
type TaobaoAilabAicloudTopHotwordsUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopHotwordsUpdateResponse
}

// 更新热词 成功返回结果
type TaobaoAilabAicloudTopHotwordsUpdateResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_hotwords_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // baseresult
    Baseresult   *BaseResult `json:"baseresult,omitempty" xml:"baseresult,omitempty"`
}
