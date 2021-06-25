package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单创建接口 APIResponse
taobao.qimen.returnorder.create

ERP调用奇门的接口,创建退货单信息;该接口和入库单的区别就是该接口是从入库单接口中单独剥离出来的，专门处理退货引起的入 库操作
*/
type TaobaoQimenReturnorderCreateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenReturnorderCreateResponse `json:"taobao_qimen_returnorder_create_response,omitempty"`
}

type TaobaoQimenReturnorderCreateResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
