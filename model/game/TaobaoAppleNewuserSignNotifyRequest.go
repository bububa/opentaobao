package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约通知接口 APIRequest
taobao.apple.newuser.sign.notify

用户付款成功后，资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyRequest struct {
    model.Params

    // 结果code
    resultCode   string 

    // 结果信息说明
    resultMsg   string 

    // 业务参数
    mainData   *AppleTopNewSignNotifyDo 

}

func NewTaobaoAppleNewuserSignNotifyRequest() *TaobaoAppleNewuserSignNotifyRequest{
    return &TaobaoAppleNewuserSignNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleNewuserSignNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.newuser.sign.notify"
}

func (r TaobaoAppleNewuserSignNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleNewuserSignNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyRequest) GetResultCode() string {
    return r.resultCode
}

func (r *TaobaoAppleNewuserSignNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleNewuserSignNotifyRequest) SetMainData(mainData *AppleTopNewSignNotifyDo) error {
    r.mainData = mainData
    r.Set("main_data", mainData)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyRequest) GetMainData() *AppleTopNewSignNotifyDo {
    return r.mainData
}

