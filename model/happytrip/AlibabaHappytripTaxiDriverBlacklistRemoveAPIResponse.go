package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse 移除司机黑名单 API返回值
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
type AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponseModel is 移除司机黑名单 成功返回结果
type AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_blacklist_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误代码
	Errno string `json:"errno,omitempty" xml:"errno,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errno = ""
	m.Errmsg = ""
}

var poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse)
	},
}

// GetAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse
func GetAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse() *AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse {
	return poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse.Get().(*AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse)
}

// ReleaseAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse 将 AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse(v *AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse.Put(v)
}
