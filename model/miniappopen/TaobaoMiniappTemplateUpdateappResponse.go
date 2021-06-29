package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新实例化应用 APIResponse
taobao.miniapp.template.updateapp

商家应用c端模板实例化小程序更新，生成新的版本，但不会自动上线新版本
*/
type TaobaoMiniappTemplateUpdateappAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappTemplateUpdateappResponse
}

type TaobaoMiniappTemplateUpdateappResponse struct {
    XMLName xml.Name `xml:"miniapp_template_updateapp_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 小程序描述
    
    AppDescription   string `json:"app_description,omitempty" xml:"app_description,omitempty"`

    
    // 小程序icon
    
    AppIcon   string `json:"app_icon,omitempty" xml:"app_icon,omitempty"`

    
    // 小程序名称
    
    AppName   string `json:"app_name,omitempty" xml:"app_name,omitempty"`

    
    // top appkey
    
    Appkey   string `json:"appkey,omitempty" xml:"appkey,omitempty"`

    
    // 小程序app_id
    
    AppId   string `json:"app_id,omitempty" xml:"app_id,omitempty"`

    
    // 当前新生成的预览版本的链接，仅当前商家有权限预览。
    
    PreViewUrl   string `json:"pre_view_url,omitempty" xml:"pre_view_url,omitempty"`

    
    // 当前新生成的预览版本号
    
    AppVersion   string `json:"app_version,omitempty" xml:"app_version,omitempty"`

    
    // 小程序简称
    
    AppAlias   string `json:"app_alias,omitempty" xml:"app_alias,omitempty"`

    
}
