package degoperation

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
中奖生成二维码 APIRequest
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

func NewTaobaoDegoperationCreateqrcodeRequest() *TaobaoDegoperationCreateqrcodeRequest{
    return &TaobaoDegoperationCreateqrcodeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetApiMethodName() string {
    return "taobao.degoperation.createqrcode"
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDegoperationCreateqrcodeRequest) SetUuid(uuid string) error {
    r.uuid = uuid
    r.Set("uuid", uuid)
    return nil
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetUuid() string {
    return r.uuid
}

func (r *TaobaoDegoperationCreateqrcodeRequest) SetDegAccessToken(degAccessToken string) error {
    r.degAccessToken = degAccessToken
    r.Set("deg_access_token", degAccessToken)
    return nil
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetDegAccessToken() string {
    return r.degAccessToken
}

func (r *TaobaoDegoperationCreateqrcodeRequest) SetSequenceNo(sequenceNo string) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetSequenceNo() string {
    return r.sequenceNo
}

func (r *TaobaoDegoperationCreateqrcodeRequest) SetActivity(activity string) error {
    r.activity = activity
    r.Set("activity", activity)
    return nil
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetActivity() string {
    return r.activity
}

func (r *TaobaoDegoperationCreateqrcodeRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoDegoperationCreateqrcodeRequest) GetTitle() string {
    return r.title
}

