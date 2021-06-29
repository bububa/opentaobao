package cityretail

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城零售完美履约转仓 APIResponse
taobao.cityretail.wmfl.convert.warehouse

同城零售完美履约转仓
*/
type TaobaoCityretailWmflConvertWarehouseAPIResponse struct {
    model.CommonResponse
    TaobaoCityretailWmflConvertWarehouseResponse
}

type TaobaoCityretailWmflConvertWarehouseResponse struct {
    XMLName xml.Name `xml:"cityretail_wmfl_convert_warehouse_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 服务出参
    
    Result   *WorkResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
