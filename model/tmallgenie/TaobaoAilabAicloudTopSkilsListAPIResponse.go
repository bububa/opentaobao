package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopskilslistAPIResponse 获取硬件平台设备下挂载的技能列表 API返回值
// taobao.ailab.aicloud.top.skils.list
//
// 提供给在硬件平台接入设备的技能列表
type TaobaoailabaicloudtopskilslistAPIResponse struct {
	model.CommonResponse
	TaobaoailabaicloudtopskilslistAPIResponseModel
}

// TaobaoailabaicloudtopskilslistAPIResponseModel is 获取硬件平台设备下挂载的技能列表 成功返回结果
type TaobaoailabaicloudtopskilslistAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
