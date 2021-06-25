package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的揽件接口 APIResponse
cainiao.guoguo.backup.graborder.takepackage

快递公司回传订单号和四位取件码给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderTakepackageAPIResponse struct {
    model.CommonResponse
    Response *CainiaoGuoguoBackupGraborderTakepackageResponse `json:"cainiao_guoguo_backup_graborder_takepackage_response,omitempty"`
}

type CainiaoGuoguoBackupGraborderTakepackageResponse struct {

    // 接口返回model
    Result   *CainiaoGuoguoBackupGraborderTakepackageResult `json:"result,omitempty"`

}
