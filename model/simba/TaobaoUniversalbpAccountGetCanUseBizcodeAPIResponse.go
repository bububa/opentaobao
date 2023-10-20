package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse 获取账户可用的bizCode API返回值
// taobao.universalbp.account.get.can.use.bizcode
//
// 查询账户可用场景，查询场景名称和场景bizcode的对应关系。其中bizcode在几乎所有接口的context中需要传入。
type TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel
}

// TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel is 获取账户可用的bizCode 成功返回结果
type TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_account_get_can_use_bizcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAccountGetCanUseBizcodeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
