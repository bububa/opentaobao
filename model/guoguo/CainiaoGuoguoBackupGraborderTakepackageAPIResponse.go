package guoguo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoguoguobackupgrabordertakepackageAPIResponse 兜底派送订单的揽件接口 API返回值
// cainiao.guoguo.backup.graborder.takepackage
//
// 快递公司回传订单号和四位取件码给菜鸟裹裹
type CainiaoguoguobackupgrabordertakepackageAPIResponse struct {
	model.CommonResponse
	CainiaoguoguobackupgrabordertakepackageAPIResponseModel
}

// CainiaoguoguobackupgrabordertakepackageAPIResponseModel is 兜底派送订单的揽件接口 成功返回结果
type CainiaoguoguobackupgrabordertakepackageAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_takepackage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *CainiaoguoguobackupgrabordertakepackageResult `json:"result,omitempty" xml:"result,omitempty"`
}
