package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
3PL直邮获取资源列表 APIResponse
taobao.wlb.import.threepl.resource.get

获取3pl直邮的发货可用资源
*/
type TaobaoWlbImportThreeplResourceGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbImportThreeplResourceGetResponse `json:"taobao_wlb_import_threepl_resource_get_response,omitempty"`
}

type TaobaoWlbImportThreeplResourceGetResponse struct {

    // result
    Result   *TopResult `json:"result,omitempty"`

}
