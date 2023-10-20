package guoguo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse 兜底派送订单的运单号回传接口 API返回值
// cainiao.guoguo.backup.graborder.submitmailno
//
// 快递公司回传订单号和运单号给菜鸟裹裹
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse struct {
	model.CommonResponse
	CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponseModel).Reset()
}

// CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponseModel is 兜底派送订单的运单号回传接口 成功返回结果
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_submitmailno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回result对象
	Result *CainiaoGuoguoBackupGraborderSubmitmailnoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse)
	},
}

// GetCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse 从 sync.Pool 获取 CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse
func GetCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse() *CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse {
	return poolCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse.Get().(*CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse)
}

// ReleaseCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse 将 CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse(v *CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse) {
	v.Reset()
	poolCainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse.Put(v)
}
