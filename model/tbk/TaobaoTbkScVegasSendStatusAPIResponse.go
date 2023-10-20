package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScVegasSendStatusAPIResponse 淘宝客-服务商-红包领取状态查询 API返回值
// taobao.tbk.sc.vegas.send.status
//
// 服务商使用。支持淘宝客传入用户标识的信息，查询该用户是否有当前阶段待核销的红包。接入前需签署协议：https://pub.alimama.com/fourth/protocol/common.htm?key=hangye_laxin
type TaobaoTbkScVegasSendStatusAPIResponse struct {
	model.CommonResponse
	TaobaoTbkScVegasSendStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkScVegasSendStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkScVegasSendStatusAPIResponseModel).Reset()
}

// TaobaoTbkScVegasSendStatusAPIResponseModel is 淘宝客-服务商-红包领取状态查询 成功返回结果
type TaobaoTbkScVegasSendStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_sc_vegas_send_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果描述信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回结果封装对象
	Data *TaobaoTbkScVegasSendStatusData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkScVegasSendStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.Data = nil
}

var poolTaobaoTbkScVegasSendStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScVegasSendStatusAPIResponse)
	},
}

// GetTaobaoTbkScVegasSendStatusAPIResponse 从 sync.Pool 获取 TaobaoTbkScVegasSendStatusAPIResponse
func GetTaobaoTbkScVegasSendStatusAPIResponse() *TaobaoTbkScVegasSendStatusAPIResponse {
	return poolTaobaoTbkScVegasSendStatusAPIResponse.Get().(*TaobaoTbkScVegasSendStatusAPIResponse)
}

// ReleaseTaobaoTbkScVegasSendStatusAPIResponse 将 TaobaoTbkScVegasSendStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkScVegasSendStatusAPIResponse(v *TaobaoTbkScVegasSendStatusAPIResponse) {
	v.Reset()
	poolTaobaoTbkScVegasSendStatusAPIResponse.Put(v)
}
