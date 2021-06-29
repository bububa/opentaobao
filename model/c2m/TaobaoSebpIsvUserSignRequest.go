package c2m

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘小铺三方签约同步 API请求
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

// 初始化TaobaoSebpIsvUserSignRequest对象
func NewTaobaoSebpIsvUserSignRequest() *TaobaoSebpIsvUserSignRequest{
    return &TaobaoSebpIsvUserSignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvUserSignRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.user.sign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvUserSignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserName Setter
// 淘宝账号
func (r *TaobaoSebpIsvUserSignRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

// UserName Getter
func (r TaobaoSebpIsvUserSignRequest) GetUserName() string {
    return r.userName
}
// Identity Setter
// 身份证
func (r *TaobaoSebpIsvUserSignRequest) SetIdentity(identity string) error {
    r.identity = identity
    r.Set("identity", identity)
    return nil
}

// Identity Getter
func (r TaobaoSebpIsvUserSignRequest) GetIdentity() string {
    return r.identity
}
// EndTime Setter
// 到期日期
func (r *TaobaoSebpIsvUserSignRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSebpIsvUserSignRequest) GetEndTime() string {
    return r.endTime
}
// StartTime Setter
// 签约日期
func (r *TaobaoSebpIsvUserSignRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSebpIsvUserSignRequest) GetStartTime() string {
    return r.startTime
}
