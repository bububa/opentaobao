package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityAddAPIResponse 新增优惠活动 API返回值
// taobao.ump.activity.add
//
// 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家）
type TaobaoUmpActivityAddAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivityAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivityAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivityAddAPIResponseModel).Reset()
}

// TaobaoUmpActivityAddAPIResponseModel is 新增优惠活动 成功返回结果
type TaobaoUmpActivityAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动id
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivityAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ActId = 0
}

var poolTaobaoUmpActivityAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivityAddAPIResponse)
	},
}

// GetTaobaoUmpActivityAddAPIResponse 从 sync.Pool 获取 TaobaoUmpActivityAddAPIResponse
func GetTaobaoUmpActivityAddAPIResponse() *TaobaoUmpActivityAddAPIResponse {
	return poolTaobaoUmpActivityAddAPIResponse.Get().(*TaobaoUmpActivityAddAPIResponse)
}

// ReleaseTaobaoUmpActivityAddAPIResponse 将 TaobaoUmpActivityAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivityAddAPIResponse(v *TaobaoUmpActivityAddAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivityAddAPIResponse.Put(v)
}
