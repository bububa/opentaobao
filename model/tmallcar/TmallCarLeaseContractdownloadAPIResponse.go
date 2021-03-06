package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseContractdownloadAPIResponse 天猫开新车租后合同下载 API返回值
// tmall.car.lease.contractdownload
//
// 天猫开新车租后合同下载
type TmallCarLeaseContractdownloadAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseContractdownloadAPIResponseModel
}

// TmallCarLeaseContractdownloadAPIResponseModel is 天猫开新车租后合同下载 成功返回结果
type TmallCarLeaseContractdownloadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_contractdownload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
