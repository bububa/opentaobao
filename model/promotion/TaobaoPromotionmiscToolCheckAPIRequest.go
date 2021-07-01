package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionmiscToolCheckAPIRequest
UMP工具检测 API请求
taobao.promotionmisc.tool.check

UMP工具检测。ISV通过该接口检测（通过taobao.ump.tool.add）创建的UMP工具（tool）是否符合规范，如果不符合，则返回错误信息和对应的解决方案的；工具检测通过后才可以提交工具审核邮件，提交工具审核时，需提供该接口的返回值。 */
type TaobaoPromotionmiscToolCheckAPIRequest struct {
	model.Params
	// 工具ID, taobao.ump.tool.add成功后返回的id。
	_toolId int64
	// 可使用的元数据。PRD审核后，会告诉isv可使用的元数据。
	_metaAllow string
}

// New
