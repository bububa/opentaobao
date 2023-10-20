package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemtemplatesgetAPIResponse 获取用户宝贝详情页模板名称 API返回值
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoitemtemplatesgetAPIResponse struct {
	model.CommonResponse
	TaobaoitemtemplatesgetAPIResponseModel
}

// TaobaoitemtemplatesgetAPIResponseModel is 获取用户宝贝详情页模板名称 成功返回结果
type TaobaoitemtemplatesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_templates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店）
	ItemTemplateList []ItemTemplate `json:"item_template_list,omitempty" xml:"item_template_list>item_template,omitempty"`
}
