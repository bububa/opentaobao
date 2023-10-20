package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleuserpermitAPIResponse 用户appkey授权 API返回值
// alibaba.idle.user.permit
//
// 闲鱼卖家与服务商关系绑定，用于业务数据分发/授权校验/消息通知鉴权
type AlibabaidleuserpermitAPIResponse struct {
	model.CommonResponse
	AlibabaidleuserpermitAPIResponseModel
}

// AlibabaidleuserpermitAPIResponseModel is 用户appkey授权 成功返回结果
type AlibabaidleuserpermitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_user_permit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
