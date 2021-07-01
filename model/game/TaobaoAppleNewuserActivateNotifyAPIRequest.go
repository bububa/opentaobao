package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户激活通知接口 API请求
taobao.apple.newuser.activate.notify

资和信主动通知激活结果
*/
type TaobaoAppleNewuserActivateNotifyAPIRequest struct {
    model.Params
    // 结果对应值，00位成功，其他为失败
    _resultCode   string
    // 处理结果中文描述
    _resultMsg   string
    // 主业务参数
    _mainData   *AppleTopActivateNotifyDo
}

// 初始化TaobaoAppleNewuserActivateNotifyAPIRequest对象
func NewTaobaoAppleNewuserActivateNotifyRequest() *TaobaoAppleNewuserActivateNotifyAPIRequest{
    return &TaobaoAppleNewuserActivateNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetApiMethodName() string {
    return "taobao.apple.newuser.activate.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResultCode Setter
// 结果对应值，00位成功，其他为失败
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetResultCode(_resultCode string) error {
    r._resultCode = _resultCode
    r.Set("result_code", _resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetResultCode() string {
    return r._resultCode
}
// ResultMsg Setter
// 处理结果中文描述
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
    r._resultMsg = _resultMsg
    r.Set("result_msg", _resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetResultMsg() string {
    return r._resultMsg
}
// MainData Setter
// 主业务参数
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetMainData(_mainData *AppleTopActivateNotifyDo) error {
    r._mainData = _mainData
    r.Set("main_data", _mainData)
    return nil
}

// MainData Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetMainData() *AppleTopActivateNotifyDo {
    return r._mainData
}
