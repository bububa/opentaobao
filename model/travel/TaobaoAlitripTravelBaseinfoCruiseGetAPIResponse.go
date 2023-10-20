package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse 【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 API返回值
// taobao.alitrip.travel.baseinfo.cruise.get
//
// 旅行度假新商品发布时可用的扩展接口，用于获取邮轮类目相关扩展信息。
type TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelBaseinfoCruiseGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelBaseinfoCruiseGetAPIResponseModel).Reset()
}

// TaobaoAlitripTravelBaseinfoCruiseGetAPIResponseModel is 【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取 成功返回结果
type TaobaoAlitripTravelBaseinfoCruiseGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_baseinfo_cruise_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 邮轮类目扩展信息的json格式字符串
	CruiseExtInfos string `json:"cruise_ext_infos,omitempty" xml:"cruise_ext_infos,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelBaseinfoCruiseGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CruiseExtInfos = ""
}

var poolTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse)
	},
}

// GetTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse
func GetTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse() *TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse {
	return poolTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse.Get().(*TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse)
}

// ReleaseTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse 将 TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse(v *TaobaoAlitripTravelBaseinfoCruiseGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelBaseinfoCruiseGetAPIResponse.Put(v)
}
