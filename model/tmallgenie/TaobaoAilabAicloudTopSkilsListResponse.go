package tmallgenie

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取硬件平台设备下挂载的技能列表 APIResponse
taobao.ailab.aicloud.top.skils.list

提供给在硬件平台接入设备的技能列表
*/
type TaobaoAilabAicloudTopSkilsListAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopSkilsListResponse
}

type TaobaoAilabAicloudTopSkilsListResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
