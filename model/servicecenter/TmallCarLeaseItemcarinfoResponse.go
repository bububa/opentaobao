package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
整车租赁商品四级车型信息 APIResponse
tmall.car.lease.itemcarinfo

整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
*/
type TmallCarLeaseItemcarinfoAPIResponse struct {
    model.CommonResponse
    TmallCarLeaseItemcarinfoResponse
}

type TmallCarLeaseItemcarinfoResponse struct {
    XMLName xml.Name `xml:"tmall_car_lease_itemcarinfo_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallCarLeaseItemcarinfoResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
