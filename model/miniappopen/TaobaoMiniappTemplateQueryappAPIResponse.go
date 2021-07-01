package miniappopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询实例化应用版本 API返回值 
taobao.miniapp.template.queryapp

根据模板id和商家信息，查询实例化小程序版本查询
*/
type TaobaoMiniappTemplateQueryappAPIResponse struct {
    model.CommonResponse
    TaobaoMiniappTemplateQueryappAPIResponseModel
}

// 查询实例化应用版本 成功返回结果
type TaobaoMiniappTemplateQueryappAPIResponseModel struct {
    XMLName xml.Name `xml:"miniapp_template_queryapp_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 所有版本信息
    AllVersionInfos   []MiniappInstanceAppAllVersionsDto `json:"all_version_infos,omitempty" xml:"all_version_infos>miniapp_instance_app_all_versions_dto,omitempty"`
}
