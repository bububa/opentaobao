package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttDvbCardChangeAPIRequest dvb ca卡替换 API请求
// youku.ott.dvb.card.change
//
// dvb 更换ca卡
type YoukuOttDvbCardChangeAPIRequest struct {
	model.Params
	// 老卡id
	_oldCardId string
	// 新卡id
	_newCardId string
	// 广电公司code(目前没用）
	_cableCompanyCode string
}

// NewYoukuOttDvbCardChangeRequest 初始化YoukuOttDvbCardChangeAPIRequest对象
func NewYoukuOttDvbCardChangeRequest() *YoukuOttDvbCardChangeAPIRequest {
	return &YoukuOttDvbCardChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttDvbCardChangeAPIRequest) GetApiMethodName() string {
	return "youku.ott.dvb.card.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttDvbCardChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OldCardId Setter
// 老卡id
func (r *YoukuOttDvbCardChangeAPIRequest) SetOldCardId(_oldCardId string) error {
	r._oldCardId = _oldCardId
	r.Set("old_card_id", _oldCardId)
	return nil
}

// Get OldCardId Getter
func (r YoukuOttDvbCardChangeAPIRequest) GetOldCardId() string {
	return r._oldCardId
}

// Set is NewCardId Setter
// 新卡id
func (r *YoukuOttDvbCardChangeAPIRequest) SetNewCardId(_newCardId string) error {
	r._newCardId = _newCardId
	r.Set("new_card_id", _newCardId)
	return nil
}

// Get NewCardId Getter
func (r YoukuOttDvbCardChangeAPIRequest) GetNewCardId() string {
	return r._newCardId
}

// Set is CableCompanyCode Setter
// 广电公司code(目前没用）
func (r *YoukuOttDvbCardChangeAPIRequest) SetCableCompanyCode(_cableCompanyCode string) error {
	r._cableCompanyCode = _cableCompanyCode
	r.Set("cable_company_code", _cableCompanyCode)
	return nil
}

// Get CableCompanyCode Getter
func (r YoukuOttDvbCardChangeAPIRequest) GetCableCompanyCode() string {
	return r._cableCompanyCode
}
