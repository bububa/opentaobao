package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺三方签约同步 APIRequest
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息
*/
type TaobaoSebpIsvUserSignRequest struct {
    model.Params

    // 淘宝账号
    userName   string 

    // 身份证
    identity   string 

    // 到期日期
    endTime   string 

    // 签约日期
    startTime   string 

}

func NewTaobaoSebpIsvUserSignRequest() *TaobaoSebpIsvUserSignRequest{
    return &TaobaoSebpIsvUserSignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSebpIsvUserSignRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.user.sign"
}

func (r TaobaoSebpIsvUserSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSebpIsvUserSignRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r TaobaoSebpIsvUserSignRequest) GetUserName() string {
    return r.userName
}

func (r *TaobaoSebpIsvUserSignRequest) SetIdentity(identity string) error {
    r.identity = identity
    r.Set("identity", identity)
    return nil
}

func (r TaobaoSebpIsvUserSignRequest) GetIdentity() string {
    return r.identity
}

func (r *TaobaoSebpIsvUserSignRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TaobaoSebpIsvUserSignRequest) GetEndTime() string {
    return r.endTime
}

func (r *TaobaoSebpIsvUserSignRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSebpIsvUserSignRequest) GetStartTime() string {
    return r.startTime
}

