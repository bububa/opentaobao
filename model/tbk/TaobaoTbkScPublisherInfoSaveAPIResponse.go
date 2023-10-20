package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScPublisherInfoSaveAPIResponse 淘宝客-公用-私域用户备案 API返回值
// taobao.tbk.sc.publisher.info.save
//
// 通过入参渠道管理或会员运营管理的邀请码，生成渠道id或会员运营id，完成渠道或会员的备案。
type TaobaoTbkScPublisherInfoSaveAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScPublisherInfoSaveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScPublisherInfoSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScPublisherInfoSaveAPIResponseModel).Reset()
}

// TaobaoTbkScPublisherInfoSaveAPIResponseModel is 淘宝客-公用-私域用户备案 成功返回结果
type TaobaoTbkScPublisherInfoSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_publisher_info_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TaobaoTbkScPublisherInfoSaveData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScPublisherInfoSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoTbkScPublisherInfoSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScPublisherInfoSaveAPIResponse)
	},
}

// GetTaobaoTbkScPublisherInfoSaveAPIResponse 从 sync.Pool 获取 TaobaoTbkScPublisherInfoSaveAPIResponse
func GetTaobaoTbkScPublisherInfoSaveAPIResponse() *TaobaoTbkScPublisherInfoSaveAPIResponse {
	return poolTaobaoTbkScPublisherInfoSaveAPIResponse.Get().(*TaobaoTbkScPublisherInfoSaveAPIResponse)
}

// ReleaseTaobaoTbkScPublisherInfoSaveAPIResponse 将 TaobaoTbkScPublisherInfoSaveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScPublisherInfoSaveAPIResponse(v *TaobaoTbkScPublisherInfoSaveAPIResponse) {
	v.Reset()
	poolTaobaoTbkScPublisherInfoSaveAPIResponse.Put(v)
}
