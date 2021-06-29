package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中奖生成二维码 API请求
taobao.degoperation.createqrcode

用户中奖后，生成二维码图片链接
*/
type TaobaoDegoperationCreateqrcodeRequest struct {
    model.Params
    // 设备id
    _uuid   string
    // 系统信息
    _degAccessToken   string
    // 奖品唯一标识
    _sequenceNo   string
    // 活动名称
    _activity   string
    // 奖品名称
    _title   string
}

// 初始化TaobaoDegoperationCreateqrcodeRequest对象
func NewTaobaoDegoperationCreateqrcodeRequest() *TaobaoDegoperationCreateqrcodeRequest{
    return &TaobaoDegoperationCreateqrcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationCreateqrcodeRequest) GetApiMethodName() string {
    return "taobao.degoperation.createqrcode"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationCreateqrcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Uuid Setter
// 设备id
func (r *TaobaoDegoperationCreateqrcodeRequest) SetUuid(_uuid string) error {
    r._uuid = _uuid
    r.Set("uuid", _uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetUuid() string {
    return r._uuid
}
// DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationCreateqrcodeRequest) SetDegAccessToken(_degAccessToken string) error {
    r._degAccessToken = _degAccessToken
    r.Set("deg_access_token", _degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetDegAccessToken() string {
    return r._degAccessToken
}
// SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCreateqrcodeRequest) SetSequenceNo(_sequenceNo string) error {
    r._sequenceNo = _sequenceNo
    r.Set("sequence_no", _sequenceNo)
    return nil
}

// SequenceNo Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetSequenceNo() string {
    return r._sequenceNo
}
// Activity Setter
// 活动名称
func (r *TaobaoDegoperationCreateqrcodeRequest) SetActivity(_activity string) error {
    r._activity = _activity
    r.Set("activity", _activity)
    return nil
}

// Activity Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetActivity() string {
    return r._activity
}
// Title Setter
// 奖品名称
func (r *TaobaoDegoperationCreateqrcodeRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetTitle() string {
    return r._title
}
