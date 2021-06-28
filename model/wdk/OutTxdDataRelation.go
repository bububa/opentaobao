package wdk

// OutTxdDataRelation 
/* model for simplify = false
type OutTxdDataRelation struct {

    // 子数据类型
    
    SubBizCode   string `json:"sub_biz_code,omitempty"`
    

    // 数据类型
    
    BizCode   string `json:"biz_code,omitempty"`
    

    // 外部数据Id
    
    OutDataId   string `json:"out_data_id,omitempty"`
    

    // 淘鲜达数据Id
    
    TxdDataId   string `json:"txd_data_id,omitempty"`
    

}
*/

// OutTxdDataRelation 
type OutTxdDataRelation struct {

    // 子数据类型
    SubBizCode   string `json:"sub_biz_code,omitempty"`

    // 数据类型
    BizCode   string `json:"biz_code,omitempty"`

    // 外部数据Id
    OutDataId   string `json:"out_data_id,omitempty"`

    // 淘鲜达数据Id
    TxdDataId   string `json:"txd_data_id,omitempty"`

}
