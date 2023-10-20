package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadatabankopenoneservicedatareadyAPIResponse 瓴羊DaaS消费者增长CGP查询DataReady API返回值
// alibaba.databank.open.oneservice.dataready
//
// 瓴羊DaaS消费者增长CGP取数接口
type AlibabadatabankopenoneservicedatareadyAPIResponse struct {
	model.CommonResponse
	AlibabadatabankopenoneservicedatareadyAPIResponseModel
}

// AlibabadatabankopenoneservicedatareadyAPIResponseModel is 瓴羊DaaS消费者增长CGP查询DataReady 成功返回结果
type AlibabadatabankopenoneservicedatareadyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_databank_open_oneservice_dataready_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的对象
	Data []DateRangeModel `json:"data,omitempty" xml:"data>date_range_model,omitempty"`
	// 查询成功
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 查询成功
	CodeClass string `json:"code_class,omitempty" xml:"code_class,omitempty"`
	// 请求成功
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
}
