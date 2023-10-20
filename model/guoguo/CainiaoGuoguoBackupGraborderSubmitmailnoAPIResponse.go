package guoguo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguobackupgrabordersubmitmailnoAPIResponse 兜底派送订单的运单号回传接口 API返回值
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
type CainiaoguoguobackupgrabordersubmitmailnoAPIResponse struct {
	model.CommonResponse
	CainiaoguoguobackupgrabordersubmitmailnoAPIResponseModel
}

// CainiaoguoguobackupgrabordersubmitmailnoAPIResponseModel is 兜底派送订单的运单号回传接口 成功返回结果
type CainiaoguoguobackupgrabordersubmitmailnoAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_submitmailno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回result对象
	Result *CainiaoguoguobackupgrabordersubmitmailnoResult `json:"result,omitempty" xml:"result,omitempty"`
}
