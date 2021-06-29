package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对外音乐搜索服务 APIResponse
taobao.ailab.aicloud.top.music.search

供厂商获取音乐列表
*/
type TaobaoAilabAicloudTopMusicSearchAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopMusicSearchResponse
}

type TaobaoAilabAicloudTopMusicSearchResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_music_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
