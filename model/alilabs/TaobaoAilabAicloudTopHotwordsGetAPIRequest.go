package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopHotwordsGetAPIRequest
获取热词 API请求
taobao.ailab.aicloud.top.hotwords.get

获取ASR热词 */
type TaobaoAilabAicloudTopHotwordsGetAPIRequest struct {
	model.Params
	// 三方用户id
	_userId string
	// 业务类型
	_bizClass string
	// schemeKey
	_schema string
}

// New
