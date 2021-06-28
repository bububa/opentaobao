package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮线下发货 APIResponse
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货
*/
type TaobaoWlbImportThreeplOfflineConsignAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_import_threepl_offline_consign_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopResult `json:"result,omitempty" xml:"