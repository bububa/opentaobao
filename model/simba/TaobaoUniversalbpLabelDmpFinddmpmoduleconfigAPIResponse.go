package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponse 查询dmp浮层配置 API返回值
// taobao.universalbp.label.dmp.finddmpmoduleconfig
//
// 入参账号信息，出参达摩盘相关配置信息
type TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponseModel
}

// TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponseModel is 查询dmp浮层配置 成功返回结果
type TaobaouniversalbplabeldmpfinddmpmoduleconfigAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_label_dmp_finddmpmoduleconfig_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbplabeldmpfinddmpmoduleconfigTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
