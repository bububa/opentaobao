package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户宝贝详情页模板名称 APIResponse
taobao.item.templates.get

查询当前登录用户的店铺的宝贝详情页的模板名称
*/
type TaobaoItemTemplatesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_templates_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店）
    
    ItemTemplateList   []ItemTemplate `json:"item_template_list,omitempty" xml:"