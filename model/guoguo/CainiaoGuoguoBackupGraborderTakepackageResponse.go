package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的揽件接口 APIResponse
cainiao.guoguo.backup.graborder.takepackage

快递公司回传订单号和四位取件码给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderTakepackageAPIResponse struct {
    model.CommonResponse
    CainiaoGuoguoBackupGraborderTakepackageResponse
}

type CainiaoGuoguoBackupGraborderTakepackageResponse struct {
    XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_takepackage_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *CainiaoGuoguoBackupGraborderTakepackageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
