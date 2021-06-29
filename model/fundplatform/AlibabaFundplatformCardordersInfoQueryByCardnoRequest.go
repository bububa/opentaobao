package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过卡号查询卡信息 API请求
alibaba.fundplatform.cardorders.info.query.by.cardno

该接口由汇金实现，外部调用。通过制卡单号分页查询卡信息
*/
type AlibabaFundplatformCardordersInfoQueryByCardnoRequest struct {
    model.Params
    // 请求结构体
    _parameters   *CardMakingInfoQueryRequest
}

// 初始化AlibabaFundplatformCardordersInfoQueryByCardnoRequest对象
func NewAlibabaFundplatformCardordersInfoQueryByCardnoRequest() *AlibabaFundplatformCardordersInfoQueryByCardnoRequest{
    return &AlibabaFundplatformCardordersInfoQueryByCardnoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.info.query.by.cardno"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Parameters Setter
// 请求结构体
func (r *AlibabaFundplatformCardordersInfoQueryByCardnoRequest) SetParameters(_parameters *CardMakingInfoQueryRequest) error {
    r._parameters = _parameters
    r.Set("parameters", _parameters)
    return nil
}

// Parameters Getter
func (r AlibabaFundplatformCardordersInfoQueryByCardnoRequest) GetParameters() *CardMakingInfoQueryRequest {
    return r._parameters
}
