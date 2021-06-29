package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sellerNick获取互动实例列表 API请求
alibaba.interact.isvadmin.getinteractbysellernick

根据sellerNick获取互动实例列表
*/
type AlibabaInteractIsvadminGetinteractbysellernickRequest struct {
    model.Params
}

// 初始化AlibabaInteractIsvadminGetinteractbysellernickRequest对象
func NewAlibabaInteractIsvadminGetinteractbysellernickRequest() *AlibabaInteractIsvadminGetinteractbysellernickRequest{
    return &AlibabaInteractIsvadminGetinteractbysellernickRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminGetinteractbysellernickRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.getinteractbysellernick"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminGetinteractbysellernickRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
