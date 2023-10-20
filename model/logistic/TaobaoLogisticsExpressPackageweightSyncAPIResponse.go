package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressPackageweightSyncAPIResponse TMS包裹重量回传 API返回值
// taobao.logistics.express.packageweight.sync
//
// TMS包裹重量回传
type TaobaoLogisticsExpressPackageweightSyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressPackageweightSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressPackageweightSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressPackageweightSyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressPackageweightSyncAPIResponseModel is TMS包裹重量回传 成功返回结果
type TaobaoLogisticsExpressPackageweightSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_packageweight_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码描述
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误码标识
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 校验成功或者异常
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否可重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressPackageweightSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrorMessage = ""
	m.BizErrorCode = ""
	m.Suc = false
	m.Retry = false
}

var poolTaobaoLogisticsExpressPackageweightSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressPackageweightSyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressPackageweightSyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressPackageweightSyncAPIResponse
func GetTaobaoLogisticsExpressPackageweightSyncAPIResponse() *TaobaoLogisticsExpressPackageweightSyncAPIResponse {
	return poolTaobaoLogisticsExpressPackageweightSyncAPIResponse.Get().(*TaobaoLogisticsExpressPackageweightSyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressPackageweightSyncAPIResponse 将 TaobaoLogisticsExpressPackageweightSyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressPackageweightSyncAPIResponse(v *TaobaoLogisticsExpressPackageweightSyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressPackageweightSyncAPIResponse.Put(v)
}
