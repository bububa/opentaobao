package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
下线实例化应用 API返回值 
taobao.miniapp.template.offlineapp

对指定的实例化小程序进行下线,需要指定clients和app_version
*/
type TaobaoMiniappTemplateOfflineappAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappTemplateOfflineappAPIResponseModel
}

// 下线实例化应用 成功返回结果
type TaobaoMiniappTemplateOfflineappAPIResponseModel struct {
    XMLName xml.Name `xml:"miniapp_template_offlineapp_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 各端的下线结果
    OfflineResults   []MiniappInstanceAppOfflineDto `json:"offline_results,omitempty" xml:"offline_results>miniapp_instance_app_offline_dto,omitempty"`
    // 下线的appId
    AppId   string `json:"app_id,omitempty" xml:"app_id,omitempty"`
}
