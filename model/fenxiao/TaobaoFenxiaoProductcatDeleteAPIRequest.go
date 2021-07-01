package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductcatDeleteAPIRequest
删除产品线 API请求
taobao.fenxiao.productcat.delete

删除产品线 */
type TaobaoFenxiaoProductcatDeleteAPIRequest struct {
	model.Params
	// 产品线ID
	_productLineId int64
}

// New
