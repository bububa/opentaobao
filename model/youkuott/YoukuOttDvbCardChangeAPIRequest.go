package youkuott

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttDvbCardChangeAPIRequest
dvb ca卡替换 API请求
youku.ott.dvb.card.change

dvb 更换ca卡 */
type YoukuOttDvbCardChangeAPIRequest struct {
	model.Params
	// 老卡id
	_oldCardId string
	// 新卡id
	_newCardId string
	// 广电公司code(目前没用）
	_cableCompanyCode string
}

// New
