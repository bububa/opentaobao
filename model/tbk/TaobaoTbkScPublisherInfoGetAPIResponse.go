package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPublisherInfoGetAPIResponse 淘宝客-公用-私域用户备案信息查询 API返回值
// taobao.tbk.sc.publisher.info.get
//
// 查询已生成的渠道id或会员运营id的相关信息。
type TaobaoTbkScPublisherInfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScPublisherInfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScPublisherInfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScPublisherInfoGetAPIResponseModel).Reset()
}

// TaobaoTbkScPublisherInfoGetAPIResponseModel is 淘宝客-公用-私域用户备案信息查询 成功返回结果
type TaobaoTbkScPublisherInfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_publisher_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScPublisherInfoGetData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScPublisherInfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScPublisherInfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPublisherInfoGetAPIResponse)
	},
}

// GetTaobaoTbkScPublisherInfoGetAPIResponse 从 sync.Pool 获取 TaobaoTbkScPublisherInfoGetAPIResponse
func GetTaobaoTbkScPublisherInfoGetAPIResponse() *TaobaoTbkScPublisherInfoGetAPIResponse {
	return poolTaobaoTbkScPublisherInfoGetAPIResponse.Get().(*TaobaoTbkScPublisherInfoGetAPIResponse)
}

// ReleaseTaobaoTbkScPublisherInfoGetAPIResponse 将 TaobaoTbkScPublisherInfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScPublisherInfoGetAPIResponse(v *TaobaoTbkScPublisherInfoGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkScPublisherInfoGetAPIResponse.Put(v)
}
