package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpToolsGetAPIRequest
查询工具列表 API请求
taobao.ump.tools.get

查询工具列表 */
type TaobaoUmpToolsGetAPIRequest struct {
	model.Params
	// 工具编码
	_toolCode string
}

// New
