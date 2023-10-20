package tuanhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelComboStatusGetAPIResponse 酒店宝贝状态查询 API返回值
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
type TaobaoXhotelComboStatusGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelComboStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelComboStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelComboStatusGetAPIResponseModel).Reset()
}

// TaobaoXhotelComboStatusGetAPIResponseModel is 酒店宝贝状态查询 成功返回结果
type TaobaoXhotelComboStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_combo_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Data *ItemInfoListResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelComboStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoXhotelComboStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelComboStatusGetAPIResponse)
	},
}

// GetTaobaoXhotelComboStatusGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelComboStatusGetAPIResponse
func GetTaobaoXhotelComboStatusGetAPIResponse() *TaobaoXhotelComboStatusGetAPIResponse {
	return poolTaobaoXhotelComboStatusGetAPIResponse.Get().(*TaobaoXhotelComboStatusGetAPIResponse)
}

// ReleaseTaobaoXhotelComboStatusGetAPIResponse 将 TaobaoXhotelComboStatusGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelComboStatusGetAPIResponse(v *TaobaoXhotelComboStatusGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelComboStatusGetAPIResponse.Put(v)
}
