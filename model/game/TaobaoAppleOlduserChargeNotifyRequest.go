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
    resultCode   string
    // 结果信息说明
    resultMsg   string
    // 业务参数
    mainData   *AppleTopOldSignNotifyDo
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
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultCode(resultCode string) error {
    r.resultCode = resultCode
    r.Set("result_code", resultCode)
    return nil
}

// ResultCode Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultCode() string {
    return r.resultCode
}
// ResultMsg Setter
// 结果信息说明
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetResultMsg(resultMsg string) error {
    r.resultMsg = resultMsg
    r.Set("result_msg", resultMsg)
    return nil
}

// ResultMsg Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetResultMsg() string {
    return r.resultMsg
}
// MainData Setter
// 业务参数
func (r *TaobaoAppleOlduserChargeNotifyRequest) SetMainData(mainData *AppleTopOldSignNotifyDo) error {
    r.mainData = mainData
    r.Set("main_data", mainData)
    return nil
}

// MainData Getter
func (r TaobaoAppleOlduserChargeNotifyRequest) GetMainData() *AppleTopOldSignNotifyDo {
    return r.mainData
}
