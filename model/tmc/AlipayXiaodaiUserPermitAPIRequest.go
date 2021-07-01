package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里金融用户授权 API请求
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口
*/
type AlipayXiaodaiUserPermitAPIRequest struct {
    model.Params
    // 用户数字ID
    _userId   int64
}

// 初始化AlipayXiaodaiUserPermitAPIRequest对象
func NewAlipayXiaodaiUserPermitRequest() *AlipayXiaodaiUserPermitAPIRequest{
    return &AlipayXiaodaiUserPermitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayXiaodaiUserPermitAPIRequest) GetApiMethodName() string {
    return "alipay.xiaodai.user.permit"
}

// IRequest interface 方法, 获取API参数
func (r AlipayXiaodaiUserPermitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户数字ID
func (r *AlipayXiaodaiUserPermitAPIRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlipayXiaodaiUserPermitAPIRequest) GetUserId() int64 {
    return r._userId
}
