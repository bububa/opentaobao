package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminContentShowSetexemptauditAPIRequest
迎客松节目设置免审开关 API请求
yunos.tvpubadmin.content.show.setexemptaudit

迎客松节目设置免审开关 */
type YunosTvpubadminContentShowSetexemptauditAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
	// 牌照id：1CIBN，2WASU
	_license int64
	// 牌照免审：1-开启节目免审，2-关闭节目免审
	_exemptAudit int64
}

// New
