package jstinteractive

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务素材配置接口 API返回值 
taobao.jst.interactive.assets.config

任务素材配置接口
*/
type TaobaoJstInteractiveAssetsConfigAPIResponse struct {
    model.CommonResponse
    TaobaoJstInteractiveAssetsConfigResponse
}

// 任务素材配置接口 成功返回结果
type TaobaoJstInteractiveAssetsConfigResponse struct {
    XMLName xml.Name `xml:"jst_interactive_assets_config_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
