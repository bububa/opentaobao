package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品删除 APIResponse
alibaba.item.operate.delete

商品删除
*/
type AlibabaItemOperateDeleteAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemOperateDeleteResponse `json:"alibaba_item_operate_delete_response,omitempty"`
}

type AlibabaItemOperateDeleteResponse struct {

    // 商品删除是否成功
    Result   string `json:"result,omitempty"`

}
