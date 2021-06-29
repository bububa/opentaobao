package youkuott

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dvb ca卡替换 API请求
youku.ott.dvb.card.change

dvb 更换ca卡
*/
type YoukuOttDvbCardChangeRequest struct {
    model.Params
    // 老卡id
    oldCardId   string
    // 新卡id
    newCardId   string
    // 广电公司code(目前没用）
    cableCompanyCode   string
}

// 初始化YoukuOttDvbCardChangeRequest对象
func NewYoukuOttDvbCardChangeRequest() *YoukuOttDvbCardChangeRequest{
    return &YoukuOttDvbCardChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttDvbCardChangeRequest) GetApiMethodName() string {
    return "youku.ott.dvb.card.change"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttDvbCardChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OldCardId Setter
// 老卡id
func (r *YoukuOttDvbCardChangeRequest) SetOldCardId(oldCardId string) error {
    r.oldCardId = oldCardId
    r.Set("old_card_id", oldCardId)
    return nil
}

// OldCardId Getter
func (r YoukuOttDvbCardChangeRequest) GetOldCardId() string {
    return r.oldCardId
}
// NewCardId Setter
// 新卡id
func (r *YoukuOttDvbCardChangeRequest) SetNewCardId(newCardId string) error {
    r.newCardId = newCardId
    r.Set("new_card_id", newCardId)
    return nil
}

// NewCardId Getter
func (r YoukuOttDvbCardChangeRequest) GetNewCardId() string {
    return r.newCardId
}
// CableCompanyCode Setter
// 广电公司code(目前没用）
func (r *YoukuOttDvbCardChangeRequest) SetCableCompanyCode(cableCompanyCode string) error {
    r.cableCompanyCode = cableCompanyCode
    r.Set("cable_company_code", cableCompanyCode)
    return nil
}

// CableCompanyCode Getter
func (r YoukuOttDvbCardChangeRequest) GetCableCompanyCode() string {
    return r.cableCompanyCode
}
