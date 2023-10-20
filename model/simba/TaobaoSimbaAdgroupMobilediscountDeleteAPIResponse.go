package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse 批量删除adgroup的移动溢价 API返回值
// taobao.simba.adgroup.mobilediscount.delete
//
// 批量删除adgroup的移动溢价
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel is 批量删除adgroup的移动溢价 成功返回结果
type TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_mobilediscount_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回成功个数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupMobilediscountDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Key = ""
	m.Message = ""
	m.Result = 0
}

var poolTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse
func GetTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse() *TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse {
	return poolTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse.Get().(*TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse 将 TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse(v *TaobaoSimbaAdgroupMobilediscountDeleteAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupMobilediscountDeleteAPIResponse.Put(v)
}
