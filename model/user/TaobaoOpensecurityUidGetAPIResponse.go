package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpensecurityUidGetAPIResponse 淘宝open security uid 获取接口 API返回值
// taobao.opensecurity.uid.get
//
// 根据明文 taobao user id 换取 app的 open_uid
type TaobaoOpensecurityUidGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpensecurityUidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpensecurityUidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpensecurityUidGetAPIResponseModel).Reset()
}

// TaobaoOpensecurityUidGetAPIResponseModel is 淘宝open security uid 获取接口 成功返回结果
type TaobaoOpensecurityUidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"opensecurity_uid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// open security tbUserId，淘宝用户对每个Appkey会有唯一的一个open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpensecurityUidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUid = ""
}

var poolTaobaoOpensecurityUidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpensecurityUidGetAPIResponse)
	},
}

// GetTaobaoOpensecurityUidGetAPIResponse 从 sync.Pool 获取 TaobaoOpensecurityUidGetAPIResponse
func GetTaobaoOpensecurityUidGetAPIResponse() *TaobaoOpensecurityUidGetAPIResponse {
	return poolTaobaoOpensecurityUidGetAPIResponse.Get().(*TaobaoOpensecurityUidGetAPIResponse)
}

// ReleaseTaobaoOpensecurityUidGetAPIResponse 将 TaobaoOpensecurityUidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpensecurityUidGetAPIResponse(v *TaobaoOpensecurityUidGetAPIResponse) {
	v.Reset()
	poolTaobaoOpensecurityUidGetAPIResponse.Put(v)
}
