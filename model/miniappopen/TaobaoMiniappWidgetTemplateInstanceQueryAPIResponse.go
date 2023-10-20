package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappwidgettemplateinstancequeryAPIResponse 小部件实例化版本查询 API返回值
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
type TaobaominiappwidgettemplateinstancequeryAPIResponse struct {
	model.CommonResponse
	TaobaominiappwidgettemplateinstancequeryAPIResponseModel
}

// TaobaominiappwidgettemplateinstancequeryAPIResponseModel is 小部件实例化版本查询 成功返回结果
type TaobaominiappwidgettemplateinstancequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_widget_template_instance_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaominiappwidgettemplateinstancequeryResult `json:"result,omitempty" xml:"result,omitempty"`
}
