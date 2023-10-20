package yunosappstore

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAppstorePadHpApplistAPIResponse 查询HpPad appList API返回值
// yunos.appstore.pad.hp.applist
//
// 提供hp pad应用群数据
type YunosAppstorePadHpApplistAPIResponse struct {
	model.CommonResponse
	YunosAppstorePadHpApplistAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAppstorePadHpApplistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAppstorePadHpApplistAPIResponseModel).Reset()
}

// YunosAppstorePadHpApplistAPIResponseModel is 查询HpPad appList 成功返回结果
type YunosAppstorePadHpApplistAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_appstore_pad_hp_applist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []YunosAppstorePadHpApplistResult `json:"results,omitempty" xml:"results>yunos_appstore_pad_hp_applist_result,omitempty"`
}

// Reset 清空结构体
func (m *YunosAppstorePadHpApplistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolYunosAppstorePadHpApplistAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAppstorePadHpApplistAPIResponse)
	},
}

// GetYunosAppstorePadHpApplistAPIResponse 从 sync.Pool 获取 YunosAppstorePadHpApplistAPIResponse
func GetYunosAppstorePadHpApplistAPIResponse() *YunosAppstorePadHpApplistAPIResponse {
	return poolYunosAppstorePadHpApplistAPIResponse.Get().(*YunosAppstorePadHpApplistAPIResponse)
}

// ReleaseYunosAppstorePadHpApplistAPIResponse 将 YunosAppstorePadHpApplistAPIResponse 保存到 sync.Pool
func ReleaseYunosAppstorePadHpApplistAPIResponse(v *YunosAppstorePadHpApplistAPIResponse) {
	v.Reset()
	poolYunosAppstorePadHpApplistAPIResponse.Put(v)
}
