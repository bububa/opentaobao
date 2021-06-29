package omniorder

// TaobaoOmniitemClassifyItemBatchOperateResult 
type TaobaoOmniitemClassifyItemBatchOperateResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 状态码：0 成功,CATEGORY_NOT_FOUND 分类不属于当前卖家或找不到分类，OPERATOR_ITEM_LINK_FORBIDDEN 商品不属于当前卖家，不允许操作，EXCEPTION 程序异常
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
