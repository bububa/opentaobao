package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 APIResponse
alibaba.item.operate.downshelf

商品下架
*/
type AlibabaItemOperateDownshelfAPIResponse struct {
    model.CommonResponse
    Response *AlibabaItemOperateDownshelfResponse `json:"alibaba_item_operate_downshelf_response,omitempty"`
}

type AlibabaItemOperateDownshelfResponse struct {

    // 商品下架是否成功
    Result   string `json:"result,omitempty"`

}
