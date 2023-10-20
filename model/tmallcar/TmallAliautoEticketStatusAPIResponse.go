package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoEticketStatusAPIResponse 查询电子凭证状态 API返回值
// tmall.aliauto.eticket.status
//
// 查询天猫汽车二轮车行业门店电子凭证状态
type TmallAliautoEticketStatusAPIResponse struct {
	model.CommonResponse
	TmallAliautoEticketStatusAPIResponseModel
}

// TmallAliautoEticketStatusAPIResponseModel is 查询电子凭证状态 成功返回结果
type TmallAliautoEticketStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_eticket_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *EticketInfoDto `json:"data,omitempty" xml:"data,omitempty"`
}
