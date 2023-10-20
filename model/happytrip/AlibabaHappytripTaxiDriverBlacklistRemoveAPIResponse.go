package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxidriverblacklistremoveAPIResponse 移除司机黑名单 API返回值
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
type AlibabahappytriptaxidriverblacklistremoveAPIResponse struct {
	model.CommonResponse
	AlibabahappytriptaxidriverblacklistremoveAPIResponseModel
}

// AlibabahappytriptaxidriverblacklistremoveAPIResponseModel is 移除司机黑名单 成功返回结果
type AlibabahappytriptaxidriverblacklistremoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_driver_blacklist_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误代码
	Errno string `json:"errno,omitempty" xml:"errno,omitempty"`
	// 错误描述
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
