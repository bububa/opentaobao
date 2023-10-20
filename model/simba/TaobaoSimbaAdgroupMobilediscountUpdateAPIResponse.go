package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse 对推广组进行单独移动溢价 API返回值
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
type TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupMobilediscountUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupMobilediscountUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupMobilediscountUpdateAPIResponseModel is 对推广组进行单独移动溢价 成功返回结果
type TaobaoSimbaAdgroupMobilediscountUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroup_mobilediscount_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 更新成功的个数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupMobilediscountUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Key = ""
	m.Result = 0
}

var poolTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse
func GetTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse() *TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse {
	return poolTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse.Get().(*TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse 将 TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse(v *TaobaoSimbaAdgroupMobilediscountUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupMobilediscountUpdateAPIResponse.Put(v)
}
