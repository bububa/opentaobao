package larkiot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLarkPosBasedataGetworkstationAPIResponse 根据影城id工作站和macId获取工作站 API返回值
// taobao.lark.pos.basedata.getworkstation
//
// 获取单独工作站
type TaobaoLarkPosBasedataGetworkstationAPIResponse struct {
	model.CommonResponse
	TaobaoLarkPosBasedataGetworkstationAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLarkPosBasedataGetworkstationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLarkPosBasedataGetworkstationAPIResponseModel).Reset()
}

// TaobaoLarkPosBasedataGetworkstationAPIResponseModel is 根据影城id工作站和macId获取工作站 成功返回结果
type TaobaoLarkPosBasedataGetworkstationAPIResponseModel struct {
	XMLName xml.Name `xml:"lark_pos_basedata_getworkstation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLarkPosBasedataGetworkstationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTaobaoLarkPosBasedataGetworkstationAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLarkPosBasedataGetworkstationAPIResponse)
	},
}

// GetTaobaoLarkPosBasedataGetworkstationAPIResponse 从 sync.Pool 获取 TaobaoLarkPosBasedataGetworkstationAPIResponse
func GetTaobaoLarkPosBasedataGetworkstationAPIResponse() *TaobaoLarkPosBasedataGetworkstationAPIResponse {
	return poolTaobaoLarkPosBasedataGetworkstationAPIResponse.Get().(*TaobaoLarkPosBasedataGetworkstationAPIResponse)
}

// ReleaseTaobaoLarkPosBasedataGetworkstationAPIResponse 将 TaobaoLarkPosBasedataGetworkstationAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLarkPosBasedataGetworkstationAPIResponse(v *TaobaoLarkPosBasedataGetworkstationAPIResponse) {
	v.Reset()
	poolTaobaoLarkPosBasedataGetworkstationAPIResponse.Put(v)
}
