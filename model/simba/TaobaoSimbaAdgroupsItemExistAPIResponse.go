package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsItemExistAPIResponse 商品是否推广 API返回值
// taobao.simba.adgroups.item.exist
//
// 判断在一个推广计划中是否已经推广了一个商品
type TaobaoSimbaAdgroupsItemExistAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupsItemExistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsItemExistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupsItemExistAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupsItemExistAPIResponseModel is 商品是否推广 成功返回结果
type TaobaoSimbaAdgroupsItemExistAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroups_item_exist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true表示已经被推广，false表示没有被推广
	Exist bool `json:"exist,omitempty" xml:"exist,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsItemExistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Exist = false
}

var poolTaobaoSimbaAdgroupsItemExistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupsItemExistAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupsItemExistAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupsItemExistAPIResponse
func GetTaobaoSimbaAdgroupsItemExistAPIResponse() *TaobaoSimbaAdgroupsItemExistAPIResponse {
	return poolTaobaoSimbaAdgroupsItemExistAPIResponse.Get().(*TaobaoSimbaAdgroupsItemExistAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupsItemExistAPIResponse 将 TaobaoSimbaAdgroupsItemExistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupsItemExistAPIResponse(v *TaobaoSimbaAdgroupsItemExistAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupsItemExistAPIResponse.Put(v)
}
