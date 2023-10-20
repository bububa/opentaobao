package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappwidgettemplateinstantiateAPIResponse 小部件实例化接口 API返回值
// taobao.miniapp.widget.template.instantiate
//
// 小部件实例化接口
type TaobaominiappwidgettemplateinstantiateAPIResponse struct {
	model.CommonResponse
	TaobaominiappwidgettemplateinstantiateAPIResponseModel
}

// TaobaominiappwidgettemplateinstantiateAPIResponseModel is 小部件实例化接口 成功返回结果
type TaobaominiappwidgettemplateinstantiateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_widget_template_instantiate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaominiappwidgettemplateinstantiateResult `json:"result,omitempty" xml:"result,omitempty"`
}
