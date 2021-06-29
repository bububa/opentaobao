package alilabs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取产品下挂载的技能列表 APIResponse
taobao.ailab.aicloud.top.skils.list.new

星空平台提供的获取产品下挂载的技能列表新接口
*/
type TaobaoAilabAicloudTopSkilsListNewAPIResponse struct {
    model.CommonResponse
    TaobaoAilabAicloudTopSkilsListNewResponse
}

type TaobaoAilabAicloudTopSkilsListNewResponse struct {
    XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_new_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回包装类
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
