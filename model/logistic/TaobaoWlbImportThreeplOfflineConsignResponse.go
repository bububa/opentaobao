package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮线下发货 APIResponse
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货
*/
type TaobaoWlbImportThreeplOfflineConsignAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbImportThreeplOfflineConsignResponse `json:"taobao_wlb_import_threepl_offline_consign_response,omitempty"`
}

type TaobaoWlbImportThreeplOfflineConsignResponse struct {

    // result
    Result   *TopResult `json:"result,omitempty"`

}
