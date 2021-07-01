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
type TaobaoSebpIsvUserSignAPIRequest struct {
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

// 初始化TaobaoSebpIsvUserSignAPIRequest对象
func NewTaobaoSebpIsvUserSignRequest() *TaobaoSebpIsvUserSignAPIRequest{
    return &TaobaoSebpIsvUserSignAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvUserSignAPIRequest) GetApiMethodName() string {
    return "taobao.sebp.isv.user.sign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvUserSignAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserName Setter
// 淘宝账号
func (r *TaobaoSebpIsvUserSignAPIRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetUserName() string {
    return r._userName
}
// Identity Setter
// 身份证
func (r *TaobaoSebpIsvUserSignAPIRequest) SetIdentity(_identity string) error {
    r._identity = _identity
    r.Set("identity", _identity)
    return nil
}

// Identity Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetIdentity() string {
    return r._identity
}
// EndTime Setter
// 到期日期
func (r *TaobaoSebpIsvUserSignAPIRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetEndTime() string {
    return r._endTime
}
// StartTime Setter
// 签约日期
func (r *TaobaoSebpIsvUserSignAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetStartTime() string {
    return r._startTime
}
