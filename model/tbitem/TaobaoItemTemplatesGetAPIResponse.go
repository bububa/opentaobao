package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemTemplatesGetAPIResponse 获取用户宝贝详情页模板名称 API返回值
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetAPIResponse struct {
	model.CommonResponse
	TaobaoItemTemplatesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemTemplatesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemTemplatesGetAPIResponseModel).Reset()
}

// TaobaoItemTemplatesGetAPIResponseModel is 获取用户宝贝详情页模板名称 成功返回结果
type TaobaoItemTemplatesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"item_templates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店）
	ItemTemplateList []ItemTemplate `json:"item_template_list,omitempty" xml:"item_template_list>item_template,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemTemplatesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemTemplateList = m.ItemTemplateList[:0]
}

var poolTaobaoItemTemplatesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemTemplatesGetAPIResponse)
	},
}

// GetTaobaoItemTemplatesGetAPIResponse 从 sync.Pool 获取 TaobaoItemTemplatesGetAPIResponse
func GetTaobaoItemTemplatesGetAPIResponse() *TaobaoItemTemplatesGetAPIResponse {
	return poolTaobaoItemTemplatesGetAPIResponse.Get().(*TaobaoItemTemplatesGetAPIResponse)
}

// ReleaseTaobaoItemTemplatesGetAPIResponse 将 TaobaoItemTemplatesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemTemplatesGetAPIResponse(v *TaobaoItemTemplatesGetAPIResponse) {
	v.Reset()
	poolTaobaoItemTemplatesGetAPIResponse.Put(v)
}
