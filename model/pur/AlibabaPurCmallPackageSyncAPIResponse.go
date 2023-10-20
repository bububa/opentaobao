package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCmallPackageSyncAPIResponse 套餐同步 API返回值
// alibaba.pur.cmall.package.sync
//
// 套餐同步
type AlibabaPurCmallPackageSyncAPIResponse struct {
	model.CommonResponse
	AlibabaPurCmallPackageSyncAPIResponseModel
}

// AlibabaPurCmallPackageSyncAPIResponseModel is 套餐同步 成功返回结果
type AlibabaPurCmallPackageSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_cmall_package_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ActionResult `json:"result,omitempty" xml:"result,omitempty"`
}
