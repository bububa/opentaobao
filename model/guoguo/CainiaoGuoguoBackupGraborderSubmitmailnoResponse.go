package guoguo

import (
    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的运单号回传接口 APIResponse
cainiao.guoguo.backup.graborder.submitmailno

快递公司回传订单号和运单号给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoGuoguoBackupGraborderSubmitmailnoResponse `json:"cainiao_guoguo_backup_graborder_submitmailno_response,omitempty"` 
    CainiaoGuoguoBackupGraborderSubmitmailnoResponse
}

/* model for simplify = false
type CainiaoGuoguoBackupGraborderSubmitmailnoResponse struct {

    // 返回result对象
    
    Result  *struct {
        CainiaoGuoguoBackupGraborderSubmitmailnoResult  *CainiaoGuoguoBackupGraborderSubmitmailnoResult `json:"cainiao_guoguo_backup_graborder_submitmailno_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type CainiaoGuoguoBackupGraborderSubmitmailnoResponse struct {

    // 返回result对象
    Result   *CainiaoGuoguoBackupGraborderSubmitmailnoResult `json:"result,omitempty"`

}
