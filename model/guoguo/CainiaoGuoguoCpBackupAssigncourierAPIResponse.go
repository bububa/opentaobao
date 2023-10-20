package guoguo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguocpbackupassigncourierAPIResponse CP兜底后指定接单的小件员 API返回值
// cainiao.guoguo.cp.backup.assigncourier
//
// CP兜底后指定接单的小件员；CP改派小件员
type CainiaoguoguocpbackupassigncourierAPIResponse struct {
	model.CommonResponse
	CainiaoguoguocpbackupassigncourierAPIResponseModel
}

// CainiaoguoguocpbackupassigncourierAPIResponseModel is CP兜底后指定接单的小件员 成功返回结果
type CainiaoguoguocpbackupassigncourierAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_cp_backup_assigncourier_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 错误信息
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
