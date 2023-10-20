package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupAddAPIResponse 创建一个推广组 API返回值
// taobao.simba.adgroup.add
//
// 创建一个推广组
type TaobaoSimbaAdgroupAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupAddAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupAddAPIResponseModel is 创建一个推广组 成功返回结果
type TaobaoSimbaAdgroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增加的推广组
	Adgroup *ADGroup `json:"adgroup,omitempty" xml:"adgroup,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroup = nil
}

var poolTaobaoSimbaAdgroupAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupAddAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupAddAPIResponse
func GetTaobaoSimbaAdgroupAddAPIResponse() *TaobaoSimbaAdgroupAddAPIResponse {
	return poolTaobaoSimbaAdgroupAddAPIResponse.Get().(*TaobaoSimbaAdgroupAddAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupAddAPIResponse 将 TaobaoSimbaAdgroupAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupAddAPIResponse(v *TaobaoSimbaAdgroupAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupAddAPIResponse.Put(v)
}
