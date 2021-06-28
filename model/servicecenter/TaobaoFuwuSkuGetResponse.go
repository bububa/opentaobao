package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取内购服务及SKU详情 APIResponse
taobao.fuwu.sku.get

通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
*/
type TaobaoFuwuSkuGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fuwu_sku_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 内购服务及SKU详情
    
    Result   *ArticleViewResult `json:"result,omitempty" xml:"