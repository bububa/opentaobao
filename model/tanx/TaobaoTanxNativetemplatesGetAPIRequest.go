package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxNativetemplatesGetAPIRequest
批量获取本地模板信息 API请求
taobao.tanx.nativetemplates.get

根据传入的本地模板ID批量返回本地模板 */
type TaobaoTanxNativetemplatesGetAPIRequest struct {
	model.Params
	// dsp在tanx的memberid
	_memberId int64
	// dsp对应的tanx的token
	_token string
	// 1970年到现在的毫秒
	_signTime int64
	// 本地模板ID列表
	_templateIds []int64
}

// New
