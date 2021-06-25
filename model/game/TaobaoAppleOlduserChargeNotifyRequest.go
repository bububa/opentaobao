package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
老用户激活并兑换通知接口 APIRequest
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口
*/
type TaobaoAppleOlduserChargeNotifyRequest struct {
    model.Params

    // 结果code
    resultCode   string 

    // 结果信息说明
    resultMsg   string 

    // 业务参数
    mainData   *AppleTopOldSignNotifyDo 

}

func NewTaobaoAppleOlduserChargeNotifyRequest() *TaobaoAppleOlduserChargeNotifyRequest{
    return &TaobaoAppleOlduserChargeNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleOlduserChargeNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.olduser.charge.notify"
}

func (r TaobaoAppleOlduserChargeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultCode() string {
    return r.resultCode
}

func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleOlduserChargeNotifyRequest) SetMainData(mainData *AppleTopOldSignNotifyDo) error {
    r.mainData = mainData
    r.Set("main_data", mainData)
    return nil
}

func (r TaobaoAppleOlduserChargeNotifyRequest) GetMainData() *AppleTopOldSignNotifyDo {
    return r.mainData
}

