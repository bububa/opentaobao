package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoLotterytypesGetAPIRequest
获取可用的彩种列表 API请求
taobao.caipiao.lotterytypes.get

获取彩票系统支持的可用于赠送的彩种列表 */
type TaobaoCaipiaoLotterytypesGetAPIRequest struct {
	model.Params
}

// New
