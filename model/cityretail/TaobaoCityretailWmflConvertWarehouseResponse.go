package cityretail

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同城零售完美履约转仓 APIResponse
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓
*/
type TaobaoCityretailWmflConvertWarehouseAPIResponse struct {
    model.CommonResponse
    Response *TaobaoCityretailWmflConvertWarehouseResponse `json:"taobao_cityretail_wmfl_convert_warehouse_response,omitempty"`
}

type TaobaoCityretailWmflConvertWarehouseResponse struct {

    // 服务出参
    Result   *WorkResult `json:"result,omitempty"`

}
