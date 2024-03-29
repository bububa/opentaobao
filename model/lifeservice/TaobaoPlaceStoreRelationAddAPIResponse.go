package lifeservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPlaceStoreRelationAddAPIResponse 门店关系新增 API返回值
// taobao.place.store.relation.add
//
// 新增授权用户的门店关系信息
type TaobaoPlaceStoreRelationAddAPIResponse struct {
	model.CommonResponse
	TaobaoPlaceStoreRelationAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreRelationAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPlaceStoreRelationAddAPIResponseModel).Reset()
}

// TaobaoPlaceStoreRelationAddAPIResponseModel is 门店关系新增 成功返回结果
type TaobaoPlaceStoreRelationAddAPIResponseModel struct {
	XMLName xml.Name `xml:"place_store_relation_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果：true成功；false失败
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// 个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPlaceStoreRelationAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
	m.TotalNum = 0
	m.Failure = false
}

var poolTaobaoPlaceStoreRelationAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPlaceStoreRelationAddAPIResponse)
	},
}

// GetTaobaoPlaceStoreRelationAddAPIResponse 从 sync.Pool 获取 TaobaoPlaceStoreRelationAddAPIResponse
func GetTaobaoPlaceStoreRelationAddAPIResponse() *TaobaoPlaceStoreRelationAddAPIResponse {
	return poolTaobaoPlaceStoreRelationAddAPIResponse.Get().(*TaobaoPlaceStoreRelationAddAPIResponse)
}

// ReleaseTaobaoPlaceStoreRelationAddAPIResponse 将 TaobaoPlaceStoreRelationAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPlaceStoreRelationAddAPIResponse(v *TaobaoPlaceStoreRelationAddAPIResponse) {
	v.Reset()
	poolTaobaoPlaceStoreRelationAddAPIResponse.Put(v)
}
