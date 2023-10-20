package nrpos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamoscommdyposmerchandisegetmerchandiseAPIResponse 去前置机商品在线查询 API返回值
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
type AlibabamoscommdyposmerchandisegetmerchandiseAPIResponse struct {
	model.CommonResponse
	AlibabamoscommdyposmerchandisegetmerchandiseAPIResponseModel
}

// AlibabamoscommdyposmerchandisegetmerchandiseAPIResponseModel is 去前置机商品在线查询 成功返回结果
type AlibabamoscommdyposmerchandisegetmerchandiseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_commdy_posmerchandise_getmerchandise_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabamoscommdyposmerchandisegetmerchandiseResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
