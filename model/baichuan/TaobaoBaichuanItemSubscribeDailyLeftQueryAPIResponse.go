package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse 查询当天可添加的余量 API返回值
// taobao.baichuan.item.subscribe.daily.left.query
//
// 查询当天可添加的余量
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse struct {
	model.CommonResponse
	TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel).Reset()
}

// TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel is 查询当天可添加的余量 成功返回结果
type TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"baichuan_item_subscribe_daily_left_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoBaichuanItemSubscribeDailyLeftQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse)
	},
}

// GetTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse 从 sync.Pool 获取 TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse
func GetTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse() *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse {
	return poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse.Get().(*TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse)
}

// ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse 将 TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse(v *TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse) {
	v.Reset()
	poolTaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse.Put(v)
}
