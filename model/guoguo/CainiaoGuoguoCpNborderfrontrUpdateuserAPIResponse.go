package guoguo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguocpnborderfrontrupdateuserAPIResponse 小件员信息变更 API返回值
// cainiao.guoguo.cp.nborderfrontr.updateuser
//
// 小件员信息变更
type CainiaoguoguocpnborderfrontrupdateuserAPIResponse struct {
	model.CommonResponse
	CainiaoguoguocpnborderfrontrupdateuserAPIResponseModel
}

// CainiaoguoguocpnborderfrontrupdateuserAPIResponseModel is 小件员信息变更 成功返回结果
type CainiaoguoguocpnborderfrontrupdateuserAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_cp_nborderfrontr_updateuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	StatusMessage string `json:"status_message,omitempty" xml:"status_message,omitempty"`
	// errorCode
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
