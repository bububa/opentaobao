package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoofnselfrecycleauthAPIResponse 自助回收鉴权 API返回值
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
type TaobaoofnselfrecycleauthAPIResponse struct {
	model.CommonResponse
	TaobaoofnselfrecycleauthAPIResponseModel
}

// TaobaoofnselfrecycleauthAPIResponseModel is 自助回收鉴权 成功返回结果
type TaobaoofnselfrecycleauthAPIResponseModel struct {
	XMLName xml.Name `xml:"ofn_self_recycle_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鉴权结果
	Data *OfnSelfRecycleAuthDto `json:"data,omitempty" xml:"data,omitempty"`
}
