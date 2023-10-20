package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiDriverBlacklistAddAPIResponse 添加司机黑名单 API返回值
// alibaba.happytrip.taxi.driver.blacklist.add
//
// 实现用户1对1永久拉黑司机，如果不支持永久拉黑，则在自动解禁黑名单司机时需回调通知欢行
type AlibabaHappytripTaxiDriverBlacklistAddAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiDriverBlacklistAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverBlacklistAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiDriverBlacklistAddAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiDriverBlacklistAddAPIResponseModel is 添加司机黑名单 成功返回结果
type AlibabaHappytripTaxiDriverBlacklistAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_blacklist_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误代码
	Errno string `json:"errno,omitempty" xml:"errno,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiDriverBlacklistAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errno = ""
	m.Errmsg = ""
}

var poolAlibabaHappytripTaxiDriverBlacklistAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiDriverBlacklistAddAPIResponse)
	},
}

// GetAlibabaHappytripTaxiDriverBlacklistAddAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiDriverBlacklistAddAPIResponse
func GetAlibabaHappytripTaxiDriverBlacklistAddAPIResponse() *AlibabaHappytripTaxiDriverBlacklistAddAPIResponse {
	return poolAlibabaHappytripTaxiDriverBlacklistAddAPIResponse.Get().(*AlibabaHappytripTaxiDriverBlacklistAddAPIResponse)
}

// ReleaseAlibabaHappytripTaxiDriverBlacklistAddAPIResponse 将 AlibabaHappytripTaxiDriverBlacklistAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiDriverBlacklistAddAPIResponse(v *AlibabaHappytripTaxiDriverBlacklistAddAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiDriverBlacklistAddAPIResponse.Put(v)
}
