package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateUpdateAPIResponse 价格推送接口（全量更新） API返回值
// taobao.xhotel.rate.update
//
// 酒店产品库rate更新
type TaobaoXhotelRateUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelRateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelRateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelRateUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelRateUpdateAPIResponseModel is 价格推送接口（全量更新） 成功返回结果
type TaobaoXhotelRateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_rate_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店商品ID-酒店RPid
	GidAndRpid string `json:"gid_and_rpid,omitempty" xml:"gid_and_rpid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelRateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GidAndRpid = ""
}

var poolTaobaoXhotelRateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelRateUpdateAPIResponse)
	},
}

// GetTaobaoXhotelRateUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelRateUpdateAPIResponse
func GetTaobaoXhotelRateUpdateAPIResponse() *TaobaoXhotelRateUpdateAPIResponse {
	return poolTaobaoXhotelRateUpdateAPIResponse.Get().(*TaobaoXhotelRateUpdateAPIResponse)
}

// ReleaseTaobaoXhotelRateUpdateAPIResponse 将 TaobaoXhotelRateUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelRateUpdateAPIResponse(v *TaobaoXhotelRateUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelRateUpdateAPIResponse.Put(v)
}
