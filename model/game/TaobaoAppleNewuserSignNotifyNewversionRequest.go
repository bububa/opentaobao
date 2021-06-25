package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约结果通知接口v2 APIRequest
taobao.apple.newuser.sign.notify.newversion

资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyNewversionRequest struct {
    model.Params

    // 结果code
    resultCode   string 

    // 结果信息说明
    resultMsg   string 

    // 业务参数
    mainData   *AppleTopNewSignNotifyDo 

}

func NewTaobaoAppleNewuserSignNotifyNewversionRequest() *TaobaoAppleNewuserSignNotifyNewversionRequest{
    return &TaobaoAppleNewuserSignNotifyNewversionRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleNewuserSignNotifyNewversionRequest) GetApiMethodName() string {
    return "taobao.apple.newuser.sign.notify.newversion"
}

func (r TaobaoAppleNewuserSignNotifyNewversionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleNewuserSignNotifyNewversionRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyNewversionRequest) GetResultCode() string {
    return r.resultCode
}

func (r *TaobaoAppleNewuserSignNotifyNewversionRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyNewversionRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleNewuserSignNotifyNewversionRequest) SetMainData(mainData *AppleTopNewSignNotifyDo) error {
    r.mainData = mainData
    r.Set("main_data", mainData)
    return nil
}

func (r TaobaoAppleNewuserSignNotifyNewversionRequest) GetMainData() *AppleTopNewSignNotifyDo {
    return r.mainData
}

