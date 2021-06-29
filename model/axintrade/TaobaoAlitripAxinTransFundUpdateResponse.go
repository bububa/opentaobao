package axintrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资金单接口 APIResponse
taobao.alitrip.axin.trans.fund.update

阿信供销平台-修改资金单接口
*/
type TaobaoAlitripAxinTransFundUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripAxinTransFundUpdateResponse
}

type TaobaoAlitripAxinTransFundUpdateResponse struct {
    XMLName xml.Name `xml:"alitrip_axin_trans_fund_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoAlitripAxinTransFundUpdateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
