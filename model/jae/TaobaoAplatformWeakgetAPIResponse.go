package jae

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAplatformWeakgetAPIResponse
活动平台弱登录接口 API返回值
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作 */
type TaobaoAplatformWeakgetAPIResponse struct {
	model.CommonResponse
	TaobaoAplatformWeakgetAPIResponseModel
}

// TaobaoAplatformWeakgetAPIResponseModel is 活动平台弱登录接口 成功返回结果
type TaobaoAplatformWeakgetAPIResponseModel struct {
	XMLName xml.Name `xml:"aplatform_weakget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoAplatformWeakgetResult `json:"result,omitempty" xml:"result,omitempty"`
}
