package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopIdListConverterAPIResponse 将淘宝openId或者设备id/用户id转换为openId API返回值
// taobao.ailab.aicloud.top.id.list.converter
//
// 将淘宝openId或者设备id/用户id转换为openId
type TaobaoAilabAicloudTopIdListConverterAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopIdListConverterAPIResponseModel
}

// TaobaoAilabAicloudTopIdListConverterAPIResponseModel is 将淘宝openId或者设备id/用户id转换为openId 成功返回结果
type TaobaoAilabAicloudTopIdListConverterAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_id_list_converter_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回体
	Result *TaobaoAilabAicloudTopIdListConverterResult `json:"result,omitempty" xml:"result,omitempty"`
}
