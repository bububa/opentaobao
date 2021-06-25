package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
产品图片删除 APIResponse
taobao.fenxiao.product.image.delete

产品图片删除，只删除图片信息，不真正删除图片
*/
type TaobaoFenxiaoProductImageDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoProductImageDeleteResponse `json:"taobao_fenxiao_product_image_delete_response,omitempty"`
}

type TaobaoFenxiaoProductImageDeleteResponse struct {

    // 操作结果
    Result   bool `json:"result,omitempty"`

    // 操作时间
    Created   string `json:"created,omitempty"`

}
