package legalsuit

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开庭信息推送接口 API请求
alibaba.legal.suit.court.open.push

供ISV推送开庭信息给集团诉讼
*/
type AlibabaLegalSuitCourtOpenPushRequest struct {
    model.Params
    // 开庭信息
    _courtInfoModel   *CourtInfoModel
}

// 初始化AlibabaLegalSuitCourtOpenPushRequest对象
func NewAlibabaLegalSuitCourtOpenPushRequest() *AlibabaLegalSuitCourtOpenPushRequest{
    return &AlibabaLegalSuitCourtOpenPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtOpenPushRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.court.open.push"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtOpenPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtInfoModel Setter
// 开庭信息
func (r *AlibabaLegalSuitCourtOpenPushRequest) SetCourtInfoModel(_courtInfoModel *CourtInfoModel) error {
    r._courtInfoModel = _courtInfoModel
    r.Set("court_info_model", _courtInfoModel)
    return nil
}

// CourtInfoModel Getter
func (r AlibabaLegalSuitCourtOpenPushRequest) GetCourtInfoModel() *CourtInfoModel {
    return r._courtInfoModel
}
