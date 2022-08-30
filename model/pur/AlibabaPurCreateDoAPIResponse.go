package pur

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPurCreateDoAPIResponse top创建DO/RT接口 API返回值
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
type AlibabaPurCreateDoAPIResponse struct {
	model.CommonResponse
	AlibabaPurCreateDoAPIResponseModel
}

// AlibabaPurCreateDoAPIResponseModel is top创建DO/RT接口 成功返回结果
type AlibabaPurCreateDoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_pur_create_do_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ContentList []ResultTopDto `json:"content_list,omitempty" xml:"content_list>result_top_dto,omitempty"`
	// 无
	ErrorLevel int64 `json:"error_level,omitempty" xml:"error_level,omitempty"`
}
