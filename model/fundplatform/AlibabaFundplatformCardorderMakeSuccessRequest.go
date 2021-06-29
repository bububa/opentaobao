package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知制卡成功 API请求
alibaba.fundplatform.cardorder.make.success

当外部业务方调用资金平台异步制卡接口后，资金平台制卡成功后通知异步通知业务方
*/
type AlibabaFundplatformCardorderMakeSuccessRequest struct {
    model.Params
    // 入参对象
    _request   *AlibabaFundplatformCardorderMakeSuccessStruct
}

// 初始化AlibabaFundplatformCardorderMakeSuccessRequest对象
func NewAlibabaFundplatformCardorderMakeSuccessRequest() *AlibabaFundplatformCardorderMakeSuccessRequest{
    return &AlibabaFundplatformCardorderMakeSuccessRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.make.success"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 入参对象
func (r *AlibabaFundplatformCardorderMakeSuccessRequest) SetRequest(_request *AlibabaFundplatformCardorderMakeSuccessStruct) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r AlibabaFundplatformCardorderMakeSuccessRequest) GetRequest() *AlibabaFundplatformCardorderMakeSuccessStruct {
    return r._request
}
