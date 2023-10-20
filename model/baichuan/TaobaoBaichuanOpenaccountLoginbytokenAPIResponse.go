package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanOpenaccountLoginbytokenAPIResponse 百川TOKEN 登录 API返回值
// taobao.baichuan.openaccount.loginbytoken
//
// 百川TOKEN 登录
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLoginbytokenAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel).Reset()
}

// TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel is 百川TOKEN 登录 成功返回结果
type TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_openaccount_loginbytoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanOpenaccountLoginbytokenAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Name = ""
}

var poolTaobaoBaichuanOpenaccountLoginbytokenAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanOpenaccountLoginbytokenAPIResponse)
	},
}

// GetTaobaoBaichuanOpenaccountLoginbytokenAPIResponse 从 sync.Pool 获取 TaobaoBaichuanOpenaccountLoginbytokenAPIResponse
func GetTaobaoBaichuanOpenaccountLoginbytokenAPIResponse() *TaobaoBaichuanOpenaccountLoginbytokenAPIResponse {
	return poolTaobaoBaichuanOpenaccountLoginbytokenAPIResponse.Get().(*TaobaoBaichuanOpenaccountLoginbytokenAPIResponse)
}

// ReleaseTaobaoBaichuanOpenaccountLoginbytokenAPIResponse 将 TaobaoBaichuanOpenaccountLoginbytokenAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanOpenaccountLoginbytokenAPIResponse(v *TaobaoBaichuanOpenaccountLoginbytokenAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanOpenaccountLoginbytokenAPIResponse.Put(v)
}
