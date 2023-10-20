package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbStoresBaseinfoGetAPIResponse 商家按照仓的类型获取仓库接口 API返回值
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbStoresBaseinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbStoresBaseinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbStoresBaseinfoGetAPIResponseModel).Reset()
}

// TaobaoWlbStoresBaseinfoGetAPIResponseModel is 商家按照仓的类型获取仓库接口 成功返回结果
type TaobaoWlbStoresBaseinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_stores_baseinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 仓库列表
	StoreInfoList []StoreInfo `json:"store_info_list,omitempty" xml:"store_info_list>store_info,omitempty"`
	// 返回的总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbStoresBaseinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.StoreInfoList = m.StoreInfoList[:0]
	m.TotalCount = 0
}

var poolTaobaoWlbStoresBaseinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbStoresBaseinfoGetAPIResponse)
	},
}

// GetTaobaoWlbStoresBaseinfoGetAPIResponse 从 sync.Pool 获取 TaobaoWlbStoresBaseinfoGetAPIResponse
func GetTaobaoWlbStoresBaseinfoGetAPIResponse() *TaobaoWlbStoresBaseinfoGetAPIResponse {
	return poolTaobaoWlbStoresBaseinfoGetAPIResponse.Get().(*TaobaoWlbStoresBaseinfoGetAPIResponse)
}

// ReleaseTaobaoWlbStoresBaseinfoGetAPIResponse 将 TaobaoWlbStoresBaseinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbStoresBaseinfoGetAPIResponse(v *TaobaoWlbStoresBaseinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbStoresBaseinfoGetAPIResponse.Put(v)
}
