package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除商品和分类的关联关系 API返回值 
taobao.omniitem.classify.item.batch.operate

批量添加/删除商品和分类的关联关系
*/
type TaobaoOmniitemClassifyItemBatchOperateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyItemBatchOperateResponse
}

// 批量添加/删除商品和分类的关联关系 成功返回结果
type TaobaoOmniitemClassifyItemBatchOperateResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_item_batch_operate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TaobaoOmniitemClassifyItemBatchOperateResult `json:"result,omitempty" xml:"result,omitempty"`
}
