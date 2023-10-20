package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse 批量得到推广组 API返回值
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponseModel is 批量得到推广组 成功返回结果
type TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroupsbyadgroupids_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Adgroups *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroups = nil
}

var poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse
func GetTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse() *TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse {
	return poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse.Get().(*TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse 将 TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse(v *TaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupsbyadgroupidsGetAPIResponse.Put(v)
}
