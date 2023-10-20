package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasSendStatusAPIResponse 淘宝客-推广者-红包领取状态查询 API返回值
// taobao.tbk.dg.vegas.send.status
//
// 淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包（淘客接入前需签署协议 https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin）
type TaobaoTbkDgVegasSendStatusAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasSendStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasSendStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgVegasSendStatusAPIResponseModel).Reset()
}

// TaobaoTbkDgVegasSendStatusAPIResponseModel is 淘宝客-推广者-红包领取状态查询 成功返回结果
type TaobaoTbkDgVegasSendStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_send_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回结果封装对象
	Data *TaobaoTbkDgVegasSendStatusData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasSendStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Data = nil
}

var poolTaobaoTbkDgVegasSendStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasSendStatusAPIResponse)
	},
}

// GetTaobaoTbkDgVegasSendStatusAPIResponse 从 sync.Pool 获取 TaobaoTbkDgVegasSendStatusAPIResponse
func GetTaobaoTbkDgVegasSendStatusAPIResponse() *TaobaoTbkDgVegasSendStatusAPIResponse {
	return poolTaobaoTbkDgVegasSendStatusAPIResponse.Get().(*TaobaoTbkDgVegasSendStatusAPIResponse)
}

// ReleaseTaobaoTbkDgVegasSendStatusAPIResponse 将 TaobaoTbkDgVegasSendStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgVegasSendStatusAPIResponse(v *TaobaoTbkDgVegasSendStatusAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgVegasSendStatusAPIResponse.Put(v)
}
