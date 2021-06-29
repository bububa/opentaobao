package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动无线端抽奖接口鉴权 API请求
alibaba.interact.wireless.draw

双11到店互动无线端mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 坯子
*/
type AlibabaInteractWirelessDrawRequest struct {
    model.Params
}

// 初始化AlibabaInteractWirelessDrawRequest对象
func NewAlibabaInteractWirelessDrawRequest() *AlibabaInteractWirelessDrawRequest{
    return &AlibabaInteractWirelessDrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractWirelessDrawRequest) GetApiMethodName() string {
    return "alibaba.interact.wireless.draw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractWirelessDrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
