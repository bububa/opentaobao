package flight

import (
    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品删除 APIResponse
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口
*/
type TaobaoAlitripTotoroAuxproductDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTotoroAuxproductDeleteResponse `json:"taobao_alitrip_totoro_auxproduct_delete_response,omitempty"`
}

type TaobaoAlitripTotoroAuxproductDeleteResponse struct {

    // result
    Result   *DelAuxProductsRs `json:"result,omitempty"`

}
