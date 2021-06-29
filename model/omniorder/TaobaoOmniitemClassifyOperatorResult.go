package omniorder

// TaobaoOmniitemClassifyOperatorResult 
type TaobaoOmniitemClassifyOperatorResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 状态码：0 成功，CATEGORY_NOT_FOUND 类目找不到，编辑失败，CATEGORY_INNER_NAME_USED 类目内部名称已被使用，CATEGORY_MAX_EXCEEDED 类目数量已达到上限100，OPERATOR_ITEM_LINK_FORBIDDEN 商品不属于当前卖家，不允许操作
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
