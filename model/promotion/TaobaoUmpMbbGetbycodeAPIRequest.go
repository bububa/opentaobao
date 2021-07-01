package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpMbbGetbycodeAPIRequest
根据营销积木块代码获取积木块 API请求
taobao.ump.mbb.getbycode

根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。 */
type TaobaoUmpMbbGetbycodeAPIRequest struct {
	model.Params
	// 营销积木块code
	_code string
}

// New
