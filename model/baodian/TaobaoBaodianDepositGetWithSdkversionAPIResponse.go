package baodian

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaodiandepositgetwithsdkversionAPIResponse 查询用户宝点信息（带sdk版本，已迁移） API返回值
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
type TaobaobaodiandepositgetwithsdkversionAPIResponse struct {
	model.CommonResponse
	TaobaobaodiandepositgetwithsdkversionAPIResponseModel
}

// TaobaobaodiandepositgetwithsdkversionAPIResponseModel is 查询用户宝点信息（带sdk版本，已迁移） 成功返回结果
type TaobaobaodiandepositgetwithsdkversionAPIResponseModel struct {
	XMLName xml.Name `xml:"baodian_deposit_get_with_sdkversion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结构体
	Result *CoinUserDepositV2 `json:"result,omitempty" xml:"result,omitempty"`
}
