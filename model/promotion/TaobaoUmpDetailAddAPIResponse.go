package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailAddAPIResponse 新增活动详情 API返回值
// taobao.ump.detail.add
//
// 增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。
type TaobaoUmpDetailAddAPIResponse struct {
	model.CommonResponse
	TaobaoUmpDetailAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpDetailAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpDetailAddAPIResponseModel).Reset()
}

// TaobaoUmpDetailAddAPIResponseModel is 新增活动详情 成功返回结果
type TaobaoUmpDetailAddAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动详情的id
	DetailId int64 `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpDetailAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DetailId = 0
}

var poolTaobaoUmpDetailAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpDetailAddAPIResponse)
	},
}

// GetTaobaoUmpDetailAddAPIResponse 从 sync.Pool 获取 TaobaoUmpDetailAddAPIResponse
func GetTaobaoUmpDetailAddAPIResponse() *TaobaoUmpDetailAddAPIResponse {
	return poolTaobaoUmpDetailAddAPIResponse.Get().(*TaobaoUmpDetailAddAPIResponse)
}

// ReleaseTaobaoUmpDetailAddAPIResponse 将 TaobaoUmpDetailAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpDetailAddAPIResponse(v *TaobaoUmpDetailAddAPIResponse) {
	v.Reset()
	poolTaobaoUmpDetailAddAPIResponse.Put(v)
}
