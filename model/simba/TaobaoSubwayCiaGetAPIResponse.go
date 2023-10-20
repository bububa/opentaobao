package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCiaGetAPIResponse 查询单元智能出价信息 API返回值
// taobao.subway.cia.get
//
// 查询单元智能出价信息
type TaobaoSubwayCiaGetAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCiaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCiaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCiaGetAPIResponseModel).Reset()
}

// TaobaoSubwayCiaGetAPIResponseModel is 查询单元智能出价信息 成功返回结果
type TaobaoSubwayCiaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_cia_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单元智能出价信息
	Result *CiaConfig `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCiaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSubwayCiaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCiaGetAPIResponse)
	},
}

// GetTaobaoSubwayCiaGetAPIResponse 从 sync.Pool 获取 TaobaoSubwayCiaGetAPIResponse
func GetTaobaoSubwayCiaGetAPIResponse() *TaobaoSubwayCiaGetAPIResponse {
	return poolTaobaoSubwayCiaGetAPIResponse.Get().(*TaobaoSubwayCiaGetAPIResponse)
}

// ReleaseTaobaoSubwayCiaGetAPIResponse 将 TaobaoSubwayCiaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCiaGetAPIResponse(v *TaobaoSubwayCiaGetAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCiaGetAPIResponse.Put(v)
}
