package eleenterpriseordernew

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse
订单配送信息跟踪 API返回值
alibaba.ele.enterprise.ordernew.gettrackinginfo

订单配送信息跟踪 */
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponse struct {
	model.CommonResponse
	AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel
}

// AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel is 订单配送信息跟踪 成功返回结果
type AlibabaEleEnterpriseOrdernewGettrackinginfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_enterprise_ordernew_gettrackinginfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应code
	EnterpriseCode string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`
	// 响应信息
	EnterpriseMsg string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`
	// 返回信息
	EnterpriseData *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`
	// 请求id
	EnterpriseRequestid string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`
}
