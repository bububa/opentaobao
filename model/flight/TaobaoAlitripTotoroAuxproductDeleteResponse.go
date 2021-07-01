package flight

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品删除 API返回值 
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口
*/
type TaobaoAlitripTotoroAuxproductDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTotoroAuxproductDeleteResponse
}

// 廉航辅营产品删除 成功返回结果
type TaobaoAlitripTotoroAuxproductDeleteResponse struct {
    XMLName xml.Name `xml:"alitrip_totoro_auxproduct_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *DelAuxProductsRS `json:"result,omitempty" xml:"result,omitempty"`
}
