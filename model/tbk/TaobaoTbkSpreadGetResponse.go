package tbk

import (
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
    Response *TaobaoTbkSpreadGetResponse `json:"taobao_tbk_spread_get_response,omitempty"`
}

type TaobaoTbkSpreadGetResponse struct {

    // 传播形式对象列表
    Results   []TbkSpread `json:"results,omitempty"`

    // totalResults
    TotalResults   int64 `json:"total_results,omitempty"`

}
