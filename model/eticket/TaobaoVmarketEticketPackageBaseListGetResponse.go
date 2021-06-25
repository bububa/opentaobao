package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id，获取关联的所有包 APIResponse
taobao.vmarket.eticket.package.base.list.get

根据卖家id，获取关联的所有包
*/
type TaobaoVmarketEticketPackageBaseListGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketPackageBaseListGetResponse `json:"taobao_vmarket_eticket_package_base_list_get_response,omitempty"`
}

type TaobaoVmarketEticketPackageBaseListGetResponse struct {

    // 查询结果
    Result   *PackageResult `json:"result,omitempty"`

}
