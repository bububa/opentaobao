package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验当前用户是否入驻了零售通门店接口 API请求
alibaba.lst.pos.open.account.checkissettled

校验当前用户是否入驻了零售通门店接口
*/
type AlibabaLstPosOpenAccountCheckissettledRequest struct {
    model.Params
    // 当前登录主账号userId
    _userId   int64
}

// 初始化AlibabaLstPosOpenAccountCheckissettledRequest对象
func NewAlibabaLstPosOpenAccountCheckissettledRequest() *AlibabaLstPosOpenAccountCheckissettledRequest{
    return &AlibabaLstPosOpenAccountCheckissettledRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.account.checkissettled"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 当前登录主账号userId
func (r *AlibabaLstPosOpenAccountCheckissettledRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetUserId() int64 {
    return r._userId
}
