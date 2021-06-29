package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
回滚实例化应用 API返回值 
taobao.miniapp.template.rollback

将实例化小程序回滚到指定版本
*/
type TaobaoMiniappTemplateRollbackAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappTemplateRollbackResponse
}

// 回滚实例化应用 成功返回结果
type TaobaoMiniappTemplateRollbackResponse struct {
    XMLName xml.Name `xml:"miniapp_template_rollback_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 分端回滚结果
    OnlineResults   []MiniappInstanceAppOnlineDTO `json:"online_results,omitempty" xml:"online_results>miniapp_instance_app_online_dto,omitempty"`
    // 基本信息
    AppInfo   *MiniAppEntityTemplateDTO `json:"app_info,omitempty" xml:"app_info,omitempty"`
}
