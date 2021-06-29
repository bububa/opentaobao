package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据制卡单分页查询卡信息 API请求
alibaba.fundplatform.cardorders.info.query

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryRequest struct {
    model.Params
    // 请求结构体
    _parameters   *CardMakingInfoQueryRequest
}

// 初始化AlibabaFundplatformCardordersInfoQueryRequest对象
func NewAlibabaFundplatformCardordersInfoQueryRequest() *AlibabaFundplatformCardordersInfoQueryRequest{
    return &AlibabaFundplatformCardordersInfoQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersInfoQueryRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.info.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersInfoQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameters Setter
// 请求结构体
func (r *AlibabaFundplatformCardordersInfoQueryRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
    r._parameters = _parameters
    r.Set("parameters", _parameters)
    return nil
}

// Parameters Getter
func (r AlibabaFundplatformCardordersInfoQueryRequest) GetParameters() *CardMakingInfoQueryRequest {
    return r._parameters
}
