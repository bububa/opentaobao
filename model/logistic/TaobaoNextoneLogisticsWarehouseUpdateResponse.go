package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AG退货入仓状态写接口 APIResponse
taobao.nextone.logistics.warehouse.update

商家上传退货入仓状态给ag
*/
type TaobaoNextoneLogisticsWarehouseUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoNextoneLogisticsWarehouseUpdateResponse `json:"nextone_logistics_warehouse_update_response,omitempty"` 
    TaobaoNextoneLogisticsWarehouseUpdateResponse
}

/* model for simplify = false
type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {

    // resultData
    
    ResultData   string `json:"result_data,omitempty"`
    

    // errorInfo
    
    ErrInfo   string `json:"err_info,omitempty"`
    

    // errorCode
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // success
    
    Succeed   bool `json:"succeed,omitempty"`
    

}
*/

type TaobaoNextoneLogisticsWarehouseUpdateResponse struct {

    // resultData
    ResultData   string `json:"result_data,omitempty"`

    // errorInfo
    ErrInfo   string `json:"err_info,omitempty"`

    // errorCode
    ErrCode   string `json:"err_code,omitempty"`

    // success
    Succeed   bool `json:"succeed,omitempty"`

}
