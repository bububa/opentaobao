package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoservicereceiptgetAPIResponse isv查询服务工单详情 API返回值
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
type TmallaliautoservicereceiptgetAPIResponse struct {
	model.CommonResponse
	TmallaliautoservicereceiptgetAPIResponseModel
}

// TmallaliautoservicereceiptgetAPIResponseModel is isv查询服务工单详情 成功返回结果
type TmallaliautoservicereceiptgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_service_receipt_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
