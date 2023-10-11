package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAccountIsUniversalUserAPIResponse 判断用户是否迁移新bp API返回值
// taobao.universalbp.account.is.universal.user
//
// 获取客户是否应使用新接口。对于迁移了新bp的客户，使用新接口，没有迁移的，使用老bp接口。不可错乱使用。
type TaobaoUniversalbpAccountIsUniversalUserAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAccountIsUniversalUserAPIResponseModel
}

// TaobaoUniversalbpAccountIsUniversalUserAPIResponseModel is 判断用户是否迁移新bp 成功返回结果
type TaobaoUniversalbpAccountIsUniversalUserAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_account_is_universal_user_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAccountIsUniversalUserTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
