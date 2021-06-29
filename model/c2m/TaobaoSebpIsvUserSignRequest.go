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
    _userName   string
    // 身份证
    _identity   string
    // 到期日期
    _endTime   string
    // 签约日期
    _startTime   string
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
func (r *TaobaoSebpIsvUserSignRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r TaobaoSebpIsvUserSignRequest) GetUserName() string {
    return r._userName
}
// Identity Setter
// 身份证
func (r *TaobaoSebpIsvUserSignRequest) SetIdentity(_identity string) error {
    r._identity = _identity
    r.Set("identity", _identity)
    return nil
}

// Identity Getter
func (r TaobaoSebpIsvUserSignRequest) GetIdentity() string {
    return r._identity
}
// EndTime Setter
// 到期日期
func (r *TaobaoSebpIsvUserSignRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSebpIsvUserSignRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 签约日期
func (r *TaobaoSebpIsvUserSignRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSebpIsvUserSignRequest) GetStartTime() string {
    return r._startTime
}
