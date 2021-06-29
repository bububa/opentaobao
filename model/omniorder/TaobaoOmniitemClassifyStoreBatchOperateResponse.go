package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量添加/删除门店和分类的关联关系 APIResponse
taobao.omniitem.classify.store.batch.operate

批量添加/删除门店和分类的关联关系
*/
type TaobaoOmniitemClassifyStoreBatchOperateAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemClassifyStoreBatchOperateResponse
}

type TaobaoOmniitemClassifyStoreBatchOperateResponse struct {
    XMLName xml.Name `xml:"omniitem_classify_store_batch_operate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TaobaoOmniitemClassifyStoreBatchOperateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
