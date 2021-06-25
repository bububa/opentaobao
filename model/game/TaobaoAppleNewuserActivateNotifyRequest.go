package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户激活通知接口 APIRequest
taobao.apple.newuser.activate.notify

资和信主动通知激活结果
*/
type TaobaoAppleNewuserActivateNotifyRequest struct {
    model.Params

    // 结果对应值，00位成功，其他为失败
    resultCode   string 

    // 处理结果中文描述
    resultMsg   string 

    // 主业务参数
    mainData   *AppleTopActivateNotifyDo 

}

func NewTaobaoAppleNewuserActivateNotifyRequest() *TaobaoAppleNewuserActivateNotifyRequest{
    return &TaobaoAppleNewuserActivateNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAppleNewuserActivateNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.newuser.activate.notify"
}

func (r TaobaoAppleNewuserActivateNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAppleNewuserActivateNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

func (r TaobaoAppleNewuserActivateNotifyRequest) GetResultCode() string {
    return r.resultCode
}

func (r *TaobaoAppleNewuserActivateNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

func (r TaobaoAppleNewuserActivateNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}

func (r *TaobaoAppleNewuserActivateNotifyRequest) SetMainData(mainData *AppleTopActivateNotifyDo) error {
    r.mainData = mainData
    r.Set("main_data", mainData)
    return nil
}

func (r TaobaoAppleNewuserActivateNotifyRequest) GetMainData() *AppleTopActivateNotifyDo {
    return r.mainData
}

