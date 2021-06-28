package guoguo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的运单号回传接口 APIResponse
cainiao.guoguo.backup.graborder.submitmailno

快递公司回传订单号和运单号给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderSubmitmailnoAPIResponse struct {
    model.CommonResponse
    CainiaoGuoguoBackupGraborderSubmitmailnoResponse
}

type CainiaoGuoguoBackupGraborderSubmitmailnoResponse struct {
    XMLName xml.Name `xml:"cainiao_guoguo_backup_graborder_submitmailno_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回result对象
    
    Result   *CainiaoGuoguoBackupGraborderSubmitmailnoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
