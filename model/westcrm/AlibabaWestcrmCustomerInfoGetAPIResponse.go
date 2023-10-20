package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawestcrmcustomerinfogetAPIResponse 会员信息查询接口 API返回值
// alibaba.westcrm.customer.info.get
//
// 会员信息查询接口
type AlibabawestcrmcustomerinfogetAPIResponse struct {
	model.CommonResponse
	AlibabawestcrmcustomerinfogetAPIResponseModel
}

// AlibabawestcrmcustomerinfogetAPIResponseModel is 会员信息查询接口 成功返回结果
type AlibabawestcrmcustomerinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_westcrm_customer_info_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
