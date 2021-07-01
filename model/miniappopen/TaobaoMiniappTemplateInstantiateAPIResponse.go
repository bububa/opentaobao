package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappTemplateInstantiateAPIResponse
构建实例化应用 API返回值
taobao.miniapp.template.instantiate

实例化saas化的小程序 */
type TaobaoMiniappTemplateInstantiateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappTemplateInstantiateAPIResponseModel
}

// TaobaoMiniappTemplateInstantiateAPIResponseModel is 构建实例化应用 成功返回结果
type TaobaoMiniappTemplateInstantiateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_template_instantiate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 小程序app_id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 小程序名称按平台规则自动生成。在授权弹窗标题、「关于」页面展示名称。
	AppName string `json:"app_name,omitempty" xml:"app_name,omitempty"`
	// 小程序描述
	AppDescription string `json:"app_description,omitempty" xml:"app_description,omitempty"`
	// 小程序icon
	AppIcon string `json:"app_icon,omitempty" xml:"app_icon,omitempty"`
	// 当前新生成的预览版本的链接，仅当前商家有权限预览。
	PreViewUrl string `json:"pre_view_url,omitempty" xml:"pre_view_url,omitempty"`
	// 当前新生成的预览版本号
	AppVersion string `json:"app_version,omitempty" xml:"app_version,omitempty"`
	// 小程序简称。在小程序Loading动画、首页标题、「更多」菜单标题上优先展示简称。
	AppAlias string `json:"app_alias,omitempty" xml:"app_alias,omitempty"`
}
