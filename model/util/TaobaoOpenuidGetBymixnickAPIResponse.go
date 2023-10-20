package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetBymixnickAPIResponse 通过mixnick转换openuid API返回值
// taobao.openuid.get.bymixnick
//
// 通过mixnick转换openuid
type TaobaoOpenuidGetBymixnickAPIResponse struct {
	model.CommonResponse
	TaobaoOpenuidGetBymixnickAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetBymixnickAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenuidGetBymixnickAPIResponseModel).Reset()
}

// TaobaoOpenuidGetBymixnickAPIResponseModel is 通过mixnick转换openuid 成功返回结果
type TaobaoOpenuidGetBymixnickAPIResponseModel struct {
	XMLName xml.Name `xml:"openuid_get_bymixnick_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// OpenUID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetBymixnickAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUid = ""
}

var poolTaobaoOpenuidGetBymixnickAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenuidGetBymixnickAPIResponse)
	},
}

// GetTaobaoOpenuidGetBymixnickAPIResponse 从 sync.Pool 获取 TaobaoOpenuidGetBymixnickAPIResponse
func GetTaobaoOpenuidGetBymixnickAPIResponse() *TaobaoOpenuidGetBymixnickAPIResponse {
	return poolTaobaoOpenuidGetBymixnickAPIResponse.Get().(*TaobaoOpenuidGetBymixnickAPIResponse)
}

// ReleaseTaobaoOpenuidGetBymixnickAPIResponse 将 TaobaoOpenuidGetBymixnickAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenuidGetBymixnickAPIResponse(v *TaobaoOpenuidGetBymixnickAPIResponse) {
	v.Reset()
	poolTaobaoOpenuidGetBymixnickAPIResponse.Put(v)
}
