package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRdcAligeniusOrdermsgUpdateAPIRequest
订单消息状态回传 API请求
taobao.rdc.aligenius.ordermsg.update

用于订单消息处理状态回传 */
type TaobaoRdcAligeniusOrdermsgUpdateAPIRequest struct {
	model.Params
	// 子订单（消息中传的子订单）
	_oid int64
	// 处理状态，1=成功，2=处理失败
	_status int64
	// 主订单（消息中传的主订单）
	_tid int64
}

// New
