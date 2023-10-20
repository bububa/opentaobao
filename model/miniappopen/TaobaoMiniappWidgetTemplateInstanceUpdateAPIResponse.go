package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappwidgettemplateinstanceupdateAPIResponse 小部件实例化版本更新 API返回值
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
type TaobaominiappwidgettemplateinstanceupdateAPIResponse struct {
	model.CommonResponse
	TaobaominiappwidgettemplateinstanceupdateAPIResponseModel
}

// TaobaominiappwidgettemplateinstanceupdateAPIResponseModel is 小部件实例化版本更新 成功返回结果
type TaobaominiappwidgettemplateinstanceupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_widget_template_instance_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *TaobaominiappwidgettemplateinstanceupdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
