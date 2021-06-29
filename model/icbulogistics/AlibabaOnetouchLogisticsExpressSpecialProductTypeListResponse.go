package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商品类型配置项 API返回值 
alibaba.onetouch.logistics.express.special.product.type.list

获取商品类型配置项
*/
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressSpecialProductTypeListResponse
}

// 获取商品类型配置项 成功返回结果
type AlibabaOnetouchLogisticsExpressSpecialProductTypeListResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_special_product_type_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaOnetouchLogisticsExpressSpecialProductTypeListResult `json:"result,omitempty" xml:"result,omitempty"`
}
