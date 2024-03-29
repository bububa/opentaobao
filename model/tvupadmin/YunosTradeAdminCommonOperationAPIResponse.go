package tvupadmin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTradeAdminCommonOperationAPIResponse 交易迎客松通用服务接口 API返回值
// yunos.trade.admin.common.operation
//
// 迎客松交易相关通用接口
type YunosTradeAdminCommonOperationAPIResponse struct {
	model.CommonResponse
	YunosTradeAdminCommonOperationAPIResponseModel
}

// Reset 清空结构体
func (m *YunosTradeAdminCommonOperationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosTradeAdminCommonOperationAPIResponseModel).Reset()
}

// YunosTradeAdminCommonOperationAPIResponseModel is 交易迎客松通用服务接口 成功返回结果
type YunosTradeAdminCommonOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_trade_admin_common_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *YunosTradeAdminCommonOperationTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosTradeAdminCommonOperationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosTradeAdminCommonOperationAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosTradeAdminCommonOperationAPIResponse)
	},
}

// GetYunosTradeAdminCommonOperationAPIResponse 从 sync.Pool 获取 YunosTradeAdminCommonOperationAPIResponse
func GetYunosTradeAdminCommonOperationAPIResponse() *YunosTradeAdminCommonOperationAPIResponse {
	return poolYunosTradeAdminCommonOperationAPIResponse.Get().(*YunosTradeAdminCommonOperationAPIResponse)
}

// ReleaseYunosTradeAdminCommonOperationAPIResponse 将 YunosTradeAdminCommonOperationAPIResponse 保存到 sync.Pool
func ReleaseYunosTradeAdminCommonOperationAPIResponse(v *YunosTradeAdminCommonOperationAPIResponse) {
	v.Reset()
	poolYunosTradeAdminCommonOperationAPIResponse.Put(v)
}
