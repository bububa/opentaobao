package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
快递改约api APIResponse
taobao.logistics.express.modify.appoint

商家通过此api操作修改物流单，交易单的收货人地址、收货人联系方式、预约配送日期
*/
type TaobaoLogisticsExpressModifyAppointAPIResponse struct {
    model.CommonResponse
    Response *TaobaoLogisticsExpressModifyAppointResponse `json:"taobao_logistics_express_modify_appoint_response,omitempty"`
}

type TaobaoLogisticsExpressModifyAppointResponse struct {

    // 调用结果
    Result   *SingleResultDto `json:"result,omitempty"`

}
