package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取包基本信息 APIResponse
taobao.vmarket.eticket.package.base.get

获取包基本信息
*/
type TaobaoVmarketEticketPackageBaseGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"vmarket_eticket_package_base_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *PackageResult `json:"result,omitempty" xml:"