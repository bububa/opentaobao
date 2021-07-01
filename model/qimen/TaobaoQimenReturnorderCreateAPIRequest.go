package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenReturnorderCreateAPIRequest
退货入库单创建接口 API请求
taobao.qimen.returnorder.create

ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作 */
type TaobaoQimenReturnorderCreateAPIRequest struct {
	model.Params
	//
	_request *ReturnOrderCreateRequest
}

// New
