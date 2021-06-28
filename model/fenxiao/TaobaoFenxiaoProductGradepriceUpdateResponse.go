package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据sku设置折扣价 APIResponse
taobao.fenxiao.product.gradeprice.update

供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格
*/
type TaobaoFenxiaoProductGradepriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_gradeprice_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回操作结果：成功或失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"