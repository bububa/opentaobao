package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvWxrobotCallbackAPIRequest
isv机器人回调接口 API请求
taobao.sebp.isv.wxrobot.callback

机器人入群回调，进行校验、功能开通等操作 */
type TaobaoSebpIsvWxrobotCallbackAPIRequest struct {
	model.Params
	// 操作类型
	_nType string
	// 调用签名
	_strSign string
	// 参数
	_strContext string
}

// New
