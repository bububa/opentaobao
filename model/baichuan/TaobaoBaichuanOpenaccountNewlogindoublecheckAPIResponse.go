package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaichuanopenaccountnewlogindoublecheckAPIResponse 百川新登录二次验证 API返回值
// taobao.baichuan.openaccount.newlogindoublecheck
//
// 百川新登录二次验证
type TaobaobaichuanopenaccountnewlogindoublecheckAPIResponse struct {
	model.CommonResponse
	TaobaobaichuanopenaccountnewlogindoublecheckAPIResponseModel
}

// TaobaobaichuanopenaccountnewlogindoublecheckAPIResponseModel is 百川新登录二次验证 成功返回结果
type TaobaobaichuanopenaccountnewlogindoublecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_newlogindoublecheck_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}
