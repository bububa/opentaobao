package tbk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-长链转短链 APIResponse
taobao.tbk.spread.get

输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
现阶段只支持短连接。
*/
type TaobaoTbkSpreadGetAPIResponse struct {
    model.CommonResponse
    TaobaoTbkSpreadGetResponse
}

type TaobaoTbkSpreadGetResponse struct {
    XMLName xml.Name `xml:"tbk_spread_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 传播形式对象列表
    
    Results   []TbkSpread `json:"results,omitempty" xml:"results>tbk_spread,omitempty"`
    
    
    // totalResults
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
