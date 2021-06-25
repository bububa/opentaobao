package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品同步接口 APIResponse
taobao.qimen.singleitem.synchronize

ERP调用奇门的接口,同步商品信息给WMS
*/
type TaobaoQimenSingleitemSynchronizeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenSingleitemSynchronizeResponse `json:"taobao_qimen_singleitem_synchronize_response,omitempty"`
}

type TaobaoQimenSingleitemSynchronizeResponse struct {

    // 
    Response   *Response `json:"response,omitempty"`

}
