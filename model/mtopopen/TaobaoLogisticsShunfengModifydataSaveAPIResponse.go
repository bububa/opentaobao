package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsShunfengModifydataSaveAPIResponse 顺丰小程序修改配送信息回传接口 API返回值
// taobao.logistics.shunfeng.modifydata.save
//
// 顺丰小程序修改配送信息回传接口
type TaobaoLogisticsShunfengModifydataSaveAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsShunfengModifydataSaveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsShunfengModifydataSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsShunfengModifydataSaveAPIResponseModel).Reset()
}

// TaobaoLogisticsShunfengModifydataSaveAPIResponseModel is 顺丰小程序修改配送信息回传接口 成功返回结果
type TaobaoLogisticsShunfengModifydataSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_shunfeng_modifydata_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 通讯失败码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 通讯失败信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 回传返回结果对象
	ModifyResponse *ModifyResponse `json:"modify_response,omitempty" xml:"modify_response,omitempty"`
	// 通讯成功/失败
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsShunfengModifydataSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.ModifyResponse = nil
	m.ResultSuccess = false
}

var poolTaobaoLogisticsShunfengModifydataSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsShunfengModifydataSaveAPIResponse)
	},
}

// GetTaobaoLogisticsShunfengModifydataSaveAPIResponse 从 sync.Pool 获取 TaobaoLogisticsShunfengModifydataSaveAPIResponse
func GetTaobaoLogisticsShunfengModifydataSaveAPIResponse() *TaobaoLogisticsShunfengModifydataSaveAPIResponse {
	return poolTaobaoLogisticsShunfengModifydataSaveAPIResponse.Get().(*TaobaoLogisticsShunfengModifydataSaveAPIResponse)
}

// ReleaseTaobaoLogisticsShunfengModifydataSaveAPIResponse 将 TaobaoLogisticsShunfengModifydataSaveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsShunfengModifydataSaveAPIResponse(v *TaobaoLogisticsShunfengModifydataSaveAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsShunfengModifydataSaveAPIResponse.Put(v)
}
