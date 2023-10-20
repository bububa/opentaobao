package baodian

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaodianDepositGetWithSdkversionAPIResponse 查询用户宝点信息（带sdk版本，已迁移） API返回值
// taobao.baodian.deposit.get.with.sdkversion
//
// 获取用户宝点信息（带sdk版本，已迁移）
type TaobaoBaodianDepositGetWithSdkversionAPIResponse struct {
	model.CommonResponse
	TaobaoBaodianDepositGetWithSdkversionAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaodianDepositGetWithSdkversionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaodianDepositGetWithSdkversionAPIResponseModel).Reset()
}

// TaobaoBaodianDepositGetWithSdkversionAPIResponseModel is 查询用户宝点信息（带sdk版本，已迁移） 成功返回结果
type TaobaoBaodianDepositGetWithSdkversionAPIResponseModel struct {
	XMLName xml.Name `xml:"baodian_deposit_get_with_sdkversion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结构体
	Result *CoinUserDepositV2 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaodianDepositGetWithSdkversionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaodianDepositGetWithSdkversionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaodianDepositGetWithSdkversionAPIResponse)
	},
}

// GetTaobaoBaodianDepositGetWithSdkversionAPIResponse 从 sync.Pool 获取 TaobaoBaodianDepositGetWithSdkversionAPIResponse
func GetTaobaoBaodianDepositGetWithSdkversionAPIResponse() *TaobaoBaodianDepositGetWithSdkversionAPIResponse {
	return poolTaobaoBaodianDepositGetWithSdkversionAPIResponse.Get().(*TaobaoBaodianDepositGetWithSdkversionAPIResponse)
}

// ReleaseTaobaoBaodianDepositGetWithSdkversionAPIResponse 将 TaobaoBaodianDepositGetWithSdkversionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaodianDepositGetWithSdkversionAPIResponse(v *TaobaoBaodianDepositGetWithSdkversionAPIResponse) {
	v.Reset()
	poolTaobaoBaodianDepositGetWithSdkversionAPIResponse.Put(v)
}
