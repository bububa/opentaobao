package omniorder

// TaobaoOmniitemClassifyDeleteResult 
type TaobaoOmniitemClassifyDeleteResult struct {

    // 提示信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // 状态码： 0 成功，CATEGORY_NOT_FOUND 删除失败， 没有发现属于当前卖家的，可被删除的分类，请确认商家是否拥有该分类
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

}
