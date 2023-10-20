package deliveryvoucher

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse 回滚券 API返回值
// taobao.game.deliveryvoucher.rollbackvoucher
//
// 提货券发券接口：同步券和订单的关联信息
type TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherRollbackvoucherAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGameDeliveryvoucherRollbackvoucherAPIResponseModel).Reset()
}

// TaobaoGameDeliveryvoucherRollbackvoucherAPIResponseModel is 回滚券 成功返回结果
type TaobaoGameDeliveryvoucherRollbackvoucherAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_rollbackvoucher_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGameDeliveryvoucherRollbackvoucherAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ResultCode = ""
	m.IsSuccess = false
}

var poolTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse)
	},
}

// GetTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse 从 sync.Pool 获取 TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse
func GetTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse() *TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse {
	return poolTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse.Get().(*TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse)
}

// ReleaseTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse 将 TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse(v *TaobaoGameDeliveryvoucherRollbackvoucherAPIResponse) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherRollbackvoucherAPIResponse.Put(v)
}
