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
    uuid   string
    // 系统信息
    degAccessToken   string
    // 奖品唯一标识
    sequenceNo   string
    // 活动名称
    activity   string
    // 奖品名称
    title   string
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
func (r *TaobaoDegoperationCreateqrcodeRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

// Uuid Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetUuid() string {
    return r.uuid
}
// DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationCreateqrcodeRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

// DegAccessToken Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetDegAccessToken() string {
    return r.degAccessToken
}
// SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCreateqrcodeRequest) SetSequenceNo(sequenceNo string) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

// SequenceNo Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetSequenceNo() string {
    return r.sequenceNo
}
// Activity Setter
// 活动名称
func (r *TaobaoDegoperationCreateqrcodeRequest) SetActivity(activity string) error {
    r.activity = activity
    r.Set("activity", activity)
    return nil
}

// Activity Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetActivity() string {
    return r.activity
}
// Title Setter
// 奖品名称
func (r *TaobaoDegoperationCreateqrcodeRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoDegoperationCreateqrcodeRequest) GetTitle() string {
    return r.title
}
