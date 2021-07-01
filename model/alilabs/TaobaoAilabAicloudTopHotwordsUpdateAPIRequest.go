package alilabs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopHotwordsUpdateAPIRequest
更新热词 API请求
taobao.ailab.aicloud.top.hotwords.update

更新ASR热词 */
type TaobaoAilabAicloudTopHotwordsUpdateAPIRequest struct {
	model.Params
	// schemaKey
	_schema string
	// 三方用户id
	_userId string
	// 业务类型
	_bizClass string
	// 操作类型
	_opType string
	// 热词内容
	_content *HotWordsContent
}

// New
