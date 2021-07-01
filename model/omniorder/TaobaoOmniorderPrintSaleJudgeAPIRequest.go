package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导购员判断 API请求
taobao.omniorder.print.sale.judge

用于判断当前子账号是否导购员
*/
type TaobaoOmniorderPrintSaleJudgeAPIRequest struct {
    model.Params
    // 用户子账号ID
    _subUid   int64
}

// 初始化TaobaoOmniorderPrintSaleJudgeAPIRequest对象
func NewTaobaoOmniorderPrintSaleJudgeRequest() *TaobaoOmniorderPrintSaleJudgeAPIRequest{
    return &TaobaoOmniorderPrintSaleJudgeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetApiMethodName() string {
    return "taobao.omniorder.print.sale.judge"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubUid Setter
// 用户子账号ID
func (r *TaobaoOmniorderPrintSaleJudgeAPIRequest) SetSubUid(_subUid int64) error {
    r._subUid = _subUid
    r.Set("sub_uid", _subUid)
    return nil
}

// SubUid Getter
func (r TaobaoOmniorderPrintSaleJudgeAPIRequest) GetSubUid() int64 {
    return r._subUid
}
