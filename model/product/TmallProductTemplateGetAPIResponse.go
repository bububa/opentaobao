package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallproducttemplategetAPIResponse 产品接口 API返回值
// tmall.product.template.get
//
// 产品模板获取接口，对于非关键属性的类目，发布达尔文(监管)产品时，必须先根据类目获取产品模板。&lt;br/&gt;&lt;br/&gt;产品模板定义产品发布需要的类目属性，包括：&lt;br/&gt;&lt;br/&gt;   关键属性:关键属性可以在类目上不存在。不存在的PID，默认为输入，没有子属性。属性名称在prop_name_str中取&lt;br/&gt;   绑定属性:内容为属性ID(PID)的列表,绑定属性肯定在类目上有，对应属性的类目特征，子属性请根据PID到类目上去取&lt;br/&gt;&lt;br/&gt;   过滤属性:内容有属性ID(PID)列表，很重要的属性，filter_properties包含的属性，必须填写&lt;br/&gt;&lt;br/&gt;   如果获取不到模板，非关键属性类目是不能发布产品的&lt;br/&gt;
type TmallproducttemplategetAPIResponse struct {
	model.CommonResponse
	TmallproducttemplategetAPIResponseModel
}

// TmallproducttemplategetAPIResponseModel is 产品接口 成功返回结果
type TmallproducttemplategetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_template_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 见SpuTemplateDO说明
	Template *SpuTemplateDo `json:"template,omitempty" xml:"template,omitempty"`
}
