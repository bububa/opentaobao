package omniorder

// TaobaoOmniitemClassifyStoreBatchOperateResult 
type TaobaoOmniitemClassifyStoreBatchOperateResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 状态码：0 成功，OPERATOR_STORE_LINK_FORBIDDEN 门店不属于当前卖家，不允许操作，CATEGORY_NOT_FOUND 分类不属于当前卖家，或者找不到分类，EXCEPTION 程序异常
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
