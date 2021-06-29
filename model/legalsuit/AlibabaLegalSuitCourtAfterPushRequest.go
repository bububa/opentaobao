package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新或者新增庭后信息 API请求
alibaba.legal.suit.court.after.push

供外部ISV供应商 推送庭后信息给集团诉讼
*/
type AlibabaLegalSuitCourtAfterPushRequest struct {
    model.Params
    // 庭后信息
    _afterCourtInfoModel   *AfterCourtInfoModel
}

// 初始化AlibabaLegalSuitCourtAfterPushRequest对象
func NewAlibabaLegalSuitCourtAfterPushRequest() *AlibabaLegalSuitCourtAfterPushRequest{
    return &AlibabaLegalSuitCourtAfterPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtAfterPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.after.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtAfterPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AfterCourtInfoModel Setter
// 庭后信息
func (r *AlibabaLegalSuitCourtAfterPushRequest) SetAfterCourtInfoModel(_afterCourtInfoModel *AfterCourtInfoModel) error {
    r._afterCourtInfoModel = _afterCourtInfoModel
    r.Set("after_court_info_model", _afterCourtInfoModel)
    return nil
}

// AfterCourtInfoModel Getter
func (r AlibabaLegalSuitCourtAfterPushRequest) GetAfterCourtInfoModel() *AfterCourtInfoModel {
    return r._afterCourtInfoModel
}
