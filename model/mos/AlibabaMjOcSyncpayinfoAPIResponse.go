package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcSyncpayinfoAPIResponse 支付参考号回传 API返回值
// alibaba.mj.oc.syncpayinfo
//
// 支付参考号同步到oc
type AlibabaMjOcSyncpayinfoAPIResponse struct {
	model.CommonResponse
	AlibabaMjOcSyncpayinfoAPIResponseModel
}

// AlibabaMjOcSyncpayinfoAPIResponseModel is 支付参考号回传 成功返回结果
type AlibabaMjOcSyncpayinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_oc_syncpayinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMjOcSyncpayinfoResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
