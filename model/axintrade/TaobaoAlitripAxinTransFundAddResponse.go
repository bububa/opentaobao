package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建资金单接口 APIResponse
taobao.alitrip.axin.trans.fund.add

创建资金单
*/
type TaobaoAlitripAxinTransFundAddAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransFundAddResponse
}

type TaobaoAlitripAxinTransFundAddResponse struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_fund_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoAlitripAxinTransFundAddResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
