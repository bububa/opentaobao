package baichuan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
百川TOKEN 登录 API请求
taobao.baichuan.openaccount.loginbytoken

百川TOKEN 登录
*/
type TaobaoBaichuanOpenaccountLoginbytokenRequest struct {
    model.Params
    // name
    name   string
}

// 初始化TaobaoBaichuanOpenaccountLoginbytokenRequest对象
func NewTaobaoBaichuanOpenaccountLoginbytokenRequest() *TaobaoBaichuanOpenaccountLoginbytokenRequest{
    return &TaobaoBaichuanOpenaccountLoginbytokenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetApiMethodName() string {
    return "taobao.baichuan.openaccount.loginbytoken"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// name
func (r *TaobaoBaichuanOpenaccountLoginbytokenRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoBaichuanOpenaccountLoginbytokenRequest) GetName() string {
    return r.name
}
