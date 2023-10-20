package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadatabankopenoneservicegetdataAPIResponse 瓴羊DaaS消费者运营CGP取数接口 API返回值
// alibaba.databank.open.oneservice.getdata
//
// 瓴羊DaaS消费者运营CGP取数接口
type AlibabadatabankopenoneservicegetdataAPIResponse struct {
	model.CommonResponse
	AlibabadatabankopenoneservicegetdataAPIResponseModel
}

// AlibabadatabankopenoneservicegetdataAPIResponseModel is 瓴羊DaaS消费者运营CGP取数接口 成功返回结果
type AlibabadatabankopenoneservicegetdataAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_databank_open_oneservice_getdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 指标结果
	Data []string `json:"data,omitempty" xml:"data>string,omitempty"`
	// 空
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 请求成功
	CodeClass string `json:"code_class,omitempty" xml:"code_class,omitempty"`
	// 请求成功
	Errcode int64 `json:"errcode,omitempty" xml:"errcode,omitempty"`
}
