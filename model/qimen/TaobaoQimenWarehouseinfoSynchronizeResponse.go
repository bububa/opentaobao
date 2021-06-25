package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓库同步接口 APIResponse
taobao.qimen.warehouseinfo.synchronize

仓库同步接口
*/
type TaobaoQimenWarehouseinfoSynchronizeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQimenWarehouseinfoSynchronizeResponse `json:"taobao_qimen_warehouseinfo_synchronize_response,omitempty"`
}

type TaobaoQimenWarehouseinfoSynchronizeResponse struct {

    // 响应报文
    Response   *Response `json:"response,omitempty"`

}
