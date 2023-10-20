package yunosappstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstoreOpenGetadsAPIResponse 获取外投广告 API返回值
// yunos.appstore.open.getads
//
// 将广告外投给外部合作伙伴
type YunosAppstoreOpenGetadsAPIResponse struct {
	model.CommonResponse
	YunosAppstoreOpenGetadsAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAppstoreOpenGetadsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAppstoreOpenGetadsAPIResponseModel).Reset()
}

// YunosAppstoreOpenGetadsAPIResponseModel is 获取外投广告 成功返回结果
type YunosAppstoreOpenGetadsAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_open_getads_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 广告集
	Ads []AdInfo `json:"ads,omitempty" xml:"ads>ad_info,omitempty"`
	// 响应消息
	Rm string `json:"rm,omitempty" xml:"rm,omitempty"`
	// 请求id
	Rid string `json:"rid,omitempty" xml:"rid,omitempty"`
	// 响应码
	Rc int64 `json:"rc,omitempty" xml:"rc,omitempty"`
}

// Reset 清空结构体
func (m *YunosAppstoreOpenGetadsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ads = m.Ads[:0]
	m.Rm = ""
	m.Rid = ""
	m.Rc = 0
}

var poolYunosAppstoreOpenGetadsAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAppstoreOpenGetadsAPIResponse)
	},
}

// GetYunosAppstoreOpenGetadsAPIResponse 从 sync.Pool 获取 YunosAppstoreOpenGetadsAPIResponse
func GetYunosAppstoreOpenGetadsAPIResponse() *YunosAppstoreOpenGetadsAPIResponse {
	return poolYunosAppstoreOpenGetadsAPIResponse.Get().(*YunosAppstoreOpenGetadsAPIResponse)
}

// ReleaseYunosAppstoreOpenGetadsAPIResponse 将 YunosAppstoreOpenGetadsAPIResponse 保存到 sync.Pool
func ReleaseYunosAppstoreOpenGetadsAPIResponse(v *YunosAppstoreOpenGetadsAPIResponse) {
	v.Reset()
	poolYunosAppstoreOpenGetadsAPIResponse.Put(v)
}
