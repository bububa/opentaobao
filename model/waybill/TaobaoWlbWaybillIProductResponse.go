package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询物流商产品类型接口 APIResponse
taobao.wlb.waybill.i.product

商家可以查询物流商的产品类型和服务能力。
*/
type TaobaoWlbWaybillIProductAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_waybill_i_product_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品类型返回
    
    ProductTypes   []WaybillProductType `json:"product_types,omitempty" xml:"