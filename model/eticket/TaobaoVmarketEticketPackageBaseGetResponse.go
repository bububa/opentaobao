package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取包基本信息 APIResponse
taobao.vmarket.eticket.package.base.get

获取包基本信息
*/
type TaobaoVmarketEticketPackageBaseGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketPackageBaseGetResponse `json:"taobao_vmarket_eticket_package_base_get_response,omitempty"`
}

type TaobaoVmarketEticketPackageBaseGetResponse struct {

    // 查询结果
    Result   *PackageResult `json:"result,omitempty"`

}
