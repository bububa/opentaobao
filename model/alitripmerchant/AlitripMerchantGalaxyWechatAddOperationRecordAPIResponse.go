package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse 用户领取会员卡记录接口 API返回值
// alitrip.merchant.galaxy.wechat.add.operation.record
//
// 用户领取会员卡记录接口
type AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyWechatAddOperationRecordAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyWechatAddOperationRecordAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyWechatAddOperationRecordAPIResponseModel is 用户领取会员卡记录接口 成功返回结果
type AlitripMerchantGalaxyWechatAddOperationRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_wechat_add_operation_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *AlitripMerchantGalaxyWechatAddOperationRecordResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyWechatAddOperationRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse)
	},
}

// GetAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse
func GetAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse() *AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse {
	return poolAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse.Get().(*AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse)
}

// ReleaseAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse 将 AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse(v *AlitripMerchantGalaxyWechatAddOperationRecordAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatAddOperationRecordAPIResponse.Put(v)
}
