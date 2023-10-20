package lstpos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenCashierSynccashierdataAPIResponse 收银快照同步接口(最多10条订单信息) API返回值
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
type AlibabaLstPosOpenCashierSynccashierdataAPIResponse struct {
	model.CommonResponse
	AlibabaLstPosOpenCashierSynccashierdataAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenCashierSynccashierdataAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstPosOpenCashierSynccashierdataAPIResponseModel).Reset()
}

// AlibabaLstPosOpenCashierSynccashierdataAPIResponseModel is 收银快照同步接口(最多10条订单信息) 成功返回结果
type AlibabaLstPosOpenCashierSynccashierdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_pos_open_cashier_synccashierdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *AlibabaLstPosOpenCashierSynccashierdataResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstPosOpenCashierSynccashierdataAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstPosOpenCashierSynccashierdataAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstPosOpenCashierSynccashierdataAPIResponse)
	},
}

// GetAlibabaLstPosOpenCashierSynccashierdataAPIResponse 从 sync.Pool 获取 AlibabaLstPosOpenCashierSynccashierdataAPIResponse
func GetAlibabaLstPosOpenCashierSynccashierdataAPIResponse() *AlibabaLstPosOpenCashierSynccashierdataAPIResponse {
	return poolAlibabaLstPosOpenCashierSynccashierdataAPIResponse.Get().(*AlibabaLstPosOpenCashierSynccashierdataAPIResponse)
}

// ReleaseAlibabaLstPosOpenCashierSynccashierdataAPIResponse 将 AlibabaLstPosOpenCashierSynccashierdataAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstPosOpenCashierSynccashierdataAPIResponse(v *AlibabaLstPosOpenCashierSynccashierdataAPIResponse) {
	v.Reset()
	poolAlibabaLstPosOpenCashierSynccashierdataAPIResponse.Put(v)
}
