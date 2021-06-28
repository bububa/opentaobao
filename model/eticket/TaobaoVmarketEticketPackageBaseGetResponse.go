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
    // Response *TaobaoVmarketEticketPackageBaseGetResponse `json:"vmarket_eticket_package_base_get_response,omitempty"` 
    TaobaoVmarketEticketPackageBaseGetResponse
}

/* model for simplify = false
type TaobaoVmarketEticketPackageBaseGetResponse struct {

    // 查询结果
    
    Result  *struct {
        PackageResult  *PackageResult `json:"package_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoVmarketEticketPackageBaseGetResponse struct {

    // 查询结果
    Result   *PackageResult `json:"result,omitempty"`

}
