package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpToolGetAPIRequest
查询工具 API请求
taobao.ump.tool.get

根据工具id获取一个工具对象 */
type TaobaoUmpToolGetAPIRequest struct {
	model.Params
	// 工具的id
	_toolId int64
}

// New
