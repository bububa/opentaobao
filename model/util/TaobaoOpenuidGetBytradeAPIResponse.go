package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetBytradeAPIResponse 通过订单获取对应买家的openUID API返回值
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoOpenuidGetBytradeAPIResponse struct {
	model.CommonResponse
	TaobaoOpenuidGetBytradeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetBytradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenuidGetBytradeAPIResponseModel).Reset()
}

// TaobaoOpenuidGetBytradeAPIResponseModel is 通过订单获取对应买家的openUID 成功返回结果
type TaobaoOpenuidGetBytradeAPIResponseModel struct {
	XMLName xml.Name `xml:"openuid_get_bytrade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前交易tid对应买家的openuid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetBytradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUid = ""
}

var poolTaobaoOpenuidGetBytradeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenuidGetBytradeAPIResponse)
	},
}

// GetTaobaoOpenuidGetBytradeAPIResponse 从 sync.Pool 获取 TaobaoOpenuidGetBytradeAPIResponse
func GetTaobaoOpenuidGetBytradeAPIResponse() *TaobaoOpenuidGetBytradeAPIResponse {
	return poolTaobaoOpenuidGetBytradeAPIResponse.Get().(*TaobaoOpenuidGetBytradeAPIResponse)
}

// ReleaseTaobaoOpenuidGetBytradeAPIResponse 将 TaobaoOpenuidGetBytradeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenuidGetBytradeAPIResponse(v *TaobaoOpenuidGetBytradeAPIResponse) {
	v.Reset()
	poolTaobaoOpenuidGetBytradeAPIResponse.Put(v)
}
