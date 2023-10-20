package simba

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel).Reset()
}

// TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel is 获取账户可用的bizCode 成功返回结果
type TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_account_get_can_use_bizcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAccountGetCanUseBizcodeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse)
	},
}

// GetTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse
func GetTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse() *TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse {
	return poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse.Get().(*TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse)
}

// ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse 将 TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse(v *TaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpAccountGetCanUseBizcodeAPIResponse.Put(v)
}
