package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeRelationQueryAPIResponse 查询单个订阅关系 API返回值
// taobao.baichuan.item.subscribe.relation.query
//
// 查询单个订阅关系
type TaobaoBaichuanItemSubscribeRelationQueryAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemSubscribeRelationQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeRelationQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemSubscribeRelationQueryAPIResponseModel).Reset()
}

// TaobaoBaichuanItemSubscribeRelationQueryAPIResponseModel is 查询单个订阅关系 成功返回结果
type TaobaoBaichuanItemSubscribeRelationQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_item_subscribe_relation_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemSubscribeRelationQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeRelationQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemSubscribeRelationQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeRelationQueryAPIResponse)
	},
}

// GetTaobaoBaichuanItemSubscribeRelationQueryAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeRelationQueryAPIResponse
func GetTaobaoBaichuanItemSubscribeRelationQueryAPIResponse() *TaobaoBaichuanItemSubscribeRelationQueryAPIResponse {
	return poolTaobaoBaichuanItemSubscribeRelationQueryAPIResponse.Get().(*TaobaoBaichuanItemSubscribeRelationQueryAPIResponse)
}

// ReleaseTaobaoBaichuanItemSubscribeRelationQueryAPIResponse 将 TaobaoBaichuanItemSubscribeRelationQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeRelationQueryAPIResponse(v *TaobaoBaichuanItemSubscribeRelationQueryAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeRelationQueryAPIResponse.Put(v)
}
