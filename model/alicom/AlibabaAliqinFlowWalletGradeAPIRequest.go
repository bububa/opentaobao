package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取流量档位 API请求
alibaba.aliqin.flow.wallet.grade

获取直充流量档位
*/
type AlibabaAliqinFlowWalletGradeAPIRequest struct {
    model.Params
    // 手机号码
    _phoneNum   string
}

// 初始化AlibabaAliqinFlowWalletGradeAPIRequest对象
func NewAlibabaAliqinFlowWalletGradeRequest() *AlibabaAliqinFlowWalletGradeAPIRequest{
    return &AlibabaAliqinFlowWalletGradeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.grade"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PhoneNum Setter
// 手机号码
func (r *AlibabaAliqinFlowWalletGradeAPIRequest) SetPhoneNum(_phoneNum string) error {
    r._phoneNum = _phoneNum
    r.Set("phone_num", _phoneNum)
    return nil
}

// PhoneNum Getter
func (r AlibabaAliqinFlowWalletGradeAPIRequest) GetPhoneNum() string {
    return r._phoneNum
}
