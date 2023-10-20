package drug

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthnrrxqueryimageAPIResponse o2o查看处方图片 API返回值
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
type AlibabaalihealthnrrxqueryimageAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthnrrxqueryimageAPIResponseModel
}

// AlibabaalihealthnrrxqueryimageAPIResponseModel is o2o查看处方图片 成功返回结果
type AlibabaalihealthnrrxqueryimageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_rx_queryimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	ErrorInfoCode string `json:"error_info_code,omitempty" xml:"error_info_code,omitempty"`
	// 错误描述
	ErrorInfoMsg string `json:"error_info_msg,omitempty" xml:"error_info_msg,omitempty"`
	// 处方图片url，多张图片用逗号分隔
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 成功或失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
