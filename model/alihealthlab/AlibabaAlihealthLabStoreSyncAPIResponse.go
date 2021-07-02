package alihealthlab

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthLabStoreSyncAPIResponse 阿里健康检验检测业务，isv门店同步到健康 API返回值
// alibaba.alihealth.lab.store.sync
//
// 阿里健康检验检测业务，isv门店同步到健康。支持门店的上线、下线操作
type AlibabaAlihealthLabStoreSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthLabStoreSyncAPIResponseModel
}

// AlibabaAlihealthLabStoreSyncAPIResponseModel is 阿里健康检验检测业务，isv门店同步到健康 成功返回结果
type AlibabaAlihealthLabStoreSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_lab_store_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// SUCCESS - 成功，FAIL - 失败，UNKNOWN - 未知
	ResultStatus string `json:"result_status,omitempty" xml:"result_status,omitempty"`
	// 可读的错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
