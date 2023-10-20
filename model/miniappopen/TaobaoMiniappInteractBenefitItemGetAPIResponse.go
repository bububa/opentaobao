package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappinteractbenefititemgetAPIResponse 读取实物权益奖池对应绑定的专属下单商品 API返回值
// taobao.miniapp.interact.benefit.item.get
//
// 读取实物权益奖池对应绑定的专属下单商品
type TaobaominiappinteractbenefititemgetAPIResponse struct {
	model.CommonResponse
	TaobaominiappinteractbenefititemgetAPIResponseModel
}

// TaobaominiappinteractbenefititemgetAPIResponseModel is 读取实物权益奖池对应绑定的专属下单商品 成功返回结果
type TaobaominiappinteractbenefititemgetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_interact_benefit_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败时，对应的错误码
	InvokeErrCode string `json:"invoke_err_code,omitempty" xml:"invoke_err_code,omitempty"`
	// 失败时，对应的错误信息
	InvokeErrMessage string `json:"invoke_err_message,omitempty" xml:"invoke_err_message,omitempty"`
	// 对应绑定的商品id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	InvokeSuccess bool `json:"invoke_success,omitempty" xml:"invoke_success,omitempty"`
}
