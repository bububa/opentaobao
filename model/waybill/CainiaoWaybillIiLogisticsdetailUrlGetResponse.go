package waybill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单物流详情授权url获取 APIResponse
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url
*/
type CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse struct {
    model.CommonResponse
    CainiaoWaybillIiLogisticsdetailUrlGetResponse
}

type CainiaoWaybillIiLogisticsdetailUrlGetResponse struct {
    XMLName xml.Name `xml:"cainiao_waybill_ii_logisticsdetail_url_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 授权访问的url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`

    
}
