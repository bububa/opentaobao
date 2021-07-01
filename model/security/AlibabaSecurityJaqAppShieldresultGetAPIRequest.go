package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqAppShieldresultGetAPIRequest
用户查询加固结果 API请求
alibaba.security.jaq.app.shieldresult.get

用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包 */
type AlibabaSecurityJaqAppShieldresultGetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// New
