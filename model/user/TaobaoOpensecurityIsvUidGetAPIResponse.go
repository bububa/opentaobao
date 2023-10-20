package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpensecurityIsvUidGetAPIResponse 获取open security uid for isv API返回值
// taobao.opensecurity.isv.uid.get
//
// 根据 open_uid 获取 open_uid_isv 用于同一个 isv的多个app间数据关联
type TaobaoOpensecurityIsvUidGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpensecurityIsvUidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpensecurityIsvUidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpensecurityIsvUidGetAPIResponseModel).Reset()
}

// TaobaoOpensecurityIsvUidGetAPIResponseModel is 获取open security uid for isv 成功返回结果
type TaobaoOpensecurityIsvUidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"opensecurity_isv_uid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// open security tbUserId for ISV，淘宝账号对ISV级别的唯一open security ID，用于同一个ISV多个Appkey间数据共享。
	OpenUidIsv string `json:"open_uid_isv,omitempty" xml:"open_uid_isv,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpensecurityIsvUidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUidIsv = ""
}

var poolTaobaoOpensecurityIsvUidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpensecurityIsvUidGetAPIResponse)
	},
}

// GetTaobaoOpensecurityIsvUidGetAPIResponse 从 sync.Pool 获取 TaobaoOpensecurityIsvUidGetAPIResponse
func GetTaobaoOpensecurityIsvUidGetAPIResponse() *TaobaoOpensecurityIsvUidGetAPIResponse {
	return poolTaobaoOpensecurityIsvUidGetAPIResponse.Get().(*TaobaoOpensecurityIsvUidGetAPIResponse)
}

// ReleaseTaobaoOpensecurityIsvUidGetAPIResponse 将 TaobaoOpensecurityIsvUidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpensecurityIsvUidGetAPIResponse(v *TaobaoOpensecurityIsvUidGetAPIResponse) {
	v.Reset()
	poolTaobaoOpensecurityIsvUidGetAPIResponse.Put(v)
}
