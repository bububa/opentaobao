package sungari

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariDisposeSubmitAPIResponse
商品商家处置提交任务 API返回值
taobao.sungari.dispose.submit

商品商家处置信息接口，提供政府部门发送处置信息给阿里 */
type TaobaoSungariDisposeSubmitAPIResponse struct {
	model.CommonResponse
	TaobaoSungariDisposeSubmitAPIResponseModel
}

// TaobaoSungariDisposeSubmitAPIResponseModel is 商品商家处置提交任务 成功返回结果
type TaobaoSungariDisposeSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"sungari_dispose_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务是否调用成功，1：成功 2：失败 11：重复提交 其他：失败
	ResuleCode int64 `json:"resule_code,omitempty" xml:"resule_code,omitempty"`
	// 返回的ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
