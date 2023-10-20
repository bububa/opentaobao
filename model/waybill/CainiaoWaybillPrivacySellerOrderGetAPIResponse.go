package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillPrivacySellerOrderGetAPIResponse 隐私面单商家订单查询 API返回值
// cainiao.waybill.privacy.seller.order.get
//
// 商家查询最近100天隐私面单记录
type CainiaoWaybillPrivacySellerOrderGetAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillPrivacySellerOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillPrivacySellerOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillPrivacySellerOrderGetAPIResponseModel).Reset()
}

// CainiaoWaybillPrivacySellerOrderGetAPIResponseModel is 隐私面单商家订单查询 成功返回结果
type CainiaoWaybillPrivacySellerOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_privacy_seller_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误列表
	ErrorCodeList []string `json:"error_code_list,omitempty" xml:"error_code_list>string,omitempty"`
	// 返回值
	ResponseList []CainiaoWaybillPrivacySellerOrderGetModule `json:"response_list,omitempty" xml:"response_list>cainiao_waybill_privacy_seller_order_get_module,omitempty"`
	// 错误信息
	ErrorInfoList []string `json:"error_info_list,omitempty" xml:"error_info_list>string,omitempty"`
	// 第一个错误
	OneErrorInfo string `json:"one_error_info,omitempty" xml:"one_error_info,omitempty"`
	// objectId
	ObjectId string `json:"object_id,omitempty" xml:"object_id,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillPrivacySellerOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCodeList = m.ErrorCodeList[:0]
	m.ResponseList = m.ResponseList[:0]
	m.ErrorInfoList = m.ErrorInfoList[:0]
	m.OneErrorInfo = ""
	m.ObjectId = ""
	m.Failure = false
}

var poolCainiaoWaybillPrivacySellerOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillPrivacySellerOrderGetAPIResponse)
	},
}

// GetCainiaoWaybillPrivacySellerOrderGetAPIResponse 从 sync.Pool 获取 CainiaoWaybillPrivacySellerOrderGetAPIResponse
func GetCainiaoWaybillPrivacySellerOrderGetAPIResponse() *CainiaoWaybillPrivacySellerOrderGetAPIResponse {
	return poolCainiaoWaybillPrivacySellerOrderGetAPIResponse.Get().(*CainiaoWaybillPrivacySellerOrderGetAPIResponse)
}

// ReleaseCainiaoWaybillPrivacySellerOrderGetAPIResponse 将 CainiaoWaybillPrivacySellerOrderGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillPrivacySellerOrderGetAPIResponse(v *CainiaoWaybillPrivacySellerOrderGetAPIResponse) {
	v.Reset()
	poolCainiaoWaybillPrivacySellerOrderGetAPIResponse.Put(v)
}
