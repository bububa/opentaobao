package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletgradeAPIResponse 获取流量档位 API返回值
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
type AlibabaaliqinflowwalletgradeAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinflowwalletgradeAPIResponseModel
}

// AlibabaaliqinflowwalletgradeAPIResponseModel is 获取流量档位 成功返回结果
type AlibabaaliqinflowwalletgradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_grade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 档位
	Grade string `json:"grade,omitempty" xml:"grade,omitempty"`
}
