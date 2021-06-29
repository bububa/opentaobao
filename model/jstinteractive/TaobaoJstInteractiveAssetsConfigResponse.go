package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务素材配置接口 APIResponse
taobao.jst.interactive.assets.config

任务素材配置接口
*/
type TaobaoJstInteractiveAssetsConfigAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveAssetsConfigResponse
}

type TaobaoJstInteractiveAssetsConfigResponse struct {
    XMLName xml.Name `xml:"jst_interactive_assets_config_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
