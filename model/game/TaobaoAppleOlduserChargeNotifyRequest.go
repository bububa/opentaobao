package game

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
老用户激活并兑换通知接口 API请求
taobao.apple.olduser.charge.notify

老用户激活并兑换通知接口
*/
type TaobaoAppleOlduserChargeNotifyRequest struct {
    model.Params
    // 结果code
    _resultCode   string
    // 结果信息说明
    _resultMsg   string
    // 业务参数
    _mainData   *AppleTopOldSignNotifyDO
}

// 初始化TaobaoAppleOlduserChargeNotifyRequest对象
func NewTaobaoAppleOlduserChargeNotifyRequest() *TaobaoAppleOlduserChargeNotifyRequest{
    return &TaobaoAppleOlduserChargeNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAppleOlduserChargeNotifyRequest) GetApiMethodName() string {
    return "taobao.apple.olduser.charge.notify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAppleOlduserChargeNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ResultCode Setter
// 结果code
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultCode(_resultCode string) error {
    r._resultCode = _resultCode
    r.Set("result_code", _resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultCode() string {
    return r._resultCode
}
// ResultMsg Setter
// 结果信息说明
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultMsg(_resultMsg string) error {
    r._resultMsg = _resultMsg
    r.Set("result_msg", _resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultMsg() string {
    return r._resultMsg
}
// MainData Setter
// 业务参数
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetMainData(_mainData *AppleTopOldSignNotifyDO) error {
    r._mainData = _mainData
    r.Set("main_data", _mainData)
    return nil
}

// MainData Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetMainData() *AppleTopOldSignNotifyDO {
    return r._mainData
}
