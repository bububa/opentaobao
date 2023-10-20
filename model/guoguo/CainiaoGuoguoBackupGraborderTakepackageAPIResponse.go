package guoguo

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGuoguoBackupGraborderTakepackageAPIResponse 兜底派送订单的揽件接口 API返回值
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
type CainiaoGuoguoBackupGraborderTakepackageAPIResponse struct {
	model.CommonResponse
	CainiaoGuoguoBackupGraborderTakepackageAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGuoguoBackupGraborderTakepackageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGuoguoBackupGraborderTakepackageAPIResponseModel).Reset()
}

// CainiaoGuoguoBackupGraborderTakepackageAPIResponseModel is 兜底派送订单的揽件接口 成功返回结果
type CainiaoGuoguoBackupGraborderTakepackageAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_takepackage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *CainiaoGuoguoBackupGraborderTakepackageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGuoguoBackupGraborderTakepackageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoGuoguoBackupGraborderTakepackageAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGuoguoBackupGraborderTakepackageAPIResponse)
	},
}

// GetCainiaoGuoguoBackupGraborderTakepackageAPIResponse 从 sync.Pool 获取 CainiaoGuoguoBackupGraborderTakepackageAPIResponse
func GetCainiaoGuoguoBackupGraborderTakepackageAPIResponse() *CainiaoGuoguoBackupGraborderTakepackageAPIResponse {
	return poolCainiaoGuoguoBackupGraborderTakepackageAPIResponse.Get().(*CainiaoGuoguoBackupGraborderTakepackageAPIResponse)
}

// ReleaseCainiaoGuoguoBackupGraborderTakepackageAPIResponse 将 CainiaoGuoguoBackupGraborderTakepackageAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGuoguoBackupGraborderTakepackageAPIResponse(v *CainiaoGuoguoBackupGraborderTakepackageAPIResponse) {
	v.Reset()
	poolCainiaoGuoguoBackupGraborderTakepackageAPIResponse.Put(v)
}
