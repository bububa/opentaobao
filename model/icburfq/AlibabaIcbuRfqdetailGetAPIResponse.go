package icburfq

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicburfqdetailgetAPIResponse 获取RFQ详情 API返回值
// alibaba.icbu.rfqdetail.get
//
// 查看RFQ的详情信息
type AlibabaicburfqdetailgetAPIResponse struct {
	model.CommonResponse
	AlibabaicburfqdetailgetAPIResponseModel
}

// AlibabaicburfqdetailgetAPIResponseModel is 获取RFQ详情 成功返回结果
type AlibabaicburfqdetailgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_rfqdetail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *RfqRemoteServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
