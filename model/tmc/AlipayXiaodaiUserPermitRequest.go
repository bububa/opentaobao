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
type AlipayXiaodaiUserPermitRequest struct {
    model.Params
    // 用户数字ID
    userId   int64
}

// 初始化AlipayXiaodaiUserPermitRequest对象
func NewAlipayXiaodaiUserPermitRequest() *AlipayXiaodaiUserPermitRequest{
    return &AlipayXiaodaiUserPermitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlipayXiaodaiUserPermitRequest) GetApiMethodName() string {
    return "alipay.xiaodai.user.permit"
}

// IRequest interface 方法, 获取API参数
func (r AlipayXiaodaiUserPermitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户数字ID
func (r *AlipayXiaodaiUserPermitRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlipayXiaodaiUserPermitRequest) GetUserId() int64 {
    return r.userId
}
