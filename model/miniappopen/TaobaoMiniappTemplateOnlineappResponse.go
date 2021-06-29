package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上线实例化应用 APIResponse
taobao.miniapp.template.onlineapp

将指定的预览版本发布上线，预览版本号由构建实例化或更新实例化接口返回。
*/
type TaobaoMiniappTemplateOnlineappAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappTemplateOnlineappResponse
}

type TaobaoMiniappTemplateOnlineappResponse struct {
    XMLName xml.Name `xml:"miniapp_template_onlineapp_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分端上线结果
    
    OnlineResults   []MiniappInstanceAppOnlineDto `json:"online_results,omitempty" xml:"online_results>miniapp_instance_app_online_dto,omitempty"`
    
    
    // 基本信息
    
    AppInfo   *MiniAppEntityTemplateDto `json:"app_info,omitempty" xml:"app_info,omitempty"`

    
}
