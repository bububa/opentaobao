package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫奖池鉴权接口 API请求
alibaba.interact.isvlottery.isvdraw

鉴权接口，为tida.isvdraw接口鉴权
*/
type AlibabaInteractIsvlotteryIsvdrawRequest struct {
    model.Params
}

// 初始化AlibabaInteractIsvlotteryIsvdrawRequest对象
func NewAlibabaInteractIsvlotteryIsvdrawRequest() *AlibabaInteractIsvlotteryIsvdrawRequest{
    return &AlibabaInteractIsvlotteryIsvdrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvlotteryIsvdrawRequest) GetApiMethodName() string {
    return "alibaba.interact.isvlottery.isvdraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvlotteryIsvdrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
