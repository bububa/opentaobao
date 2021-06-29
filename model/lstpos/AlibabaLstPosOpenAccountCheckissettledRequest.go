package lstpos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
校验当前用户是否入驻了零售通门店接口 APIRequest
alibaba.lst.pos.open.account.checkissettled

校验当前用户是否入驻了零售通门店接口
*/
type AlibabaLstPosOpenAccountCheckissettledRequest struct {
    model.Params

    // 当前登录主账号userId
    userId   int64 

}

func NewAlibabaLstPosOpenAccountCheckissettledRequest() *AlibabaLstPosOpenAccountCheckissettledRequest{
    return &AlibabaLstPosOpenAccountCheckissettledRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetApiMethodName() string {
    return "alibaba.lst.pos.open.account.checkissettled"
}

func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstPosOpenAccountCheckissettledRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaLstPosOpenAccountCheckissettledRequest) GetUserId() int64 {
    return r.userId
}

