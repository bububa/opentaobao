package interactvip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractVipGetAPIResponse
会员淘气值获取 API返回值
alibaba.interact.vip.get

提供用户淘气值&用户角色身份查询 */
type AlibabaInteractVipGetAPIResponse struct {
	model.CommonResponse
	AlibabaInteractVipGetAPIResponseModel
}

// AlibabaInteractVipGetAPIResponseModel is 会员淘气值获取 成功返回结果
type AlibabaInteractVipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_vip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
