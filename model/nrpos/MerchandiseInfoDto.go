package nrpos

// MerchandiseInfoDto 
type MerchandiseInfoDto struct {

    // 门店号
    
    Storeno   string `json:"storeno,omitempty" xml:"storeno,omitempty"`
    

    // VIP活动折扣规则
    
    A0551   string `json:"a0551,omitempty" xml:"a0551,omitempty"`
    

    // VIP普通折扣规则
    
    A0550   string `json:"a0550,omitempty" xml:"a0550,omitempty"`
    

    // 活动套号
    
    A0545   string `json:"a0545,omitempty" xml:"a0545,omitempty"`
    

    // 保留字段
    
    A0544   string `json:"a0544,omitempty" xml:"a0544,omitempty"`
    

    // 实时积分基数
    
    A0543   string `json:"a0543,omitempty" xml:"a0543,omitempty"`
    

    // 保留字段
    
    A0542   string `json:"a0542,omitempty" xml:"a0542,omitempty"`
    

    // 会员价
    
    A0541   string `json:"a0541,omitempty" xml:"a0541,omitempty"`
    

    // 是否接受VIP 0=不接受 1=接受
    
    A0540   string `json:"a0540,omitempty" xml:"a0540,omitempty"`
    

    // 最低限制折扣，该商品所有折扣累加不能小于此折扣
    
    A0531   string `json:"a0531,omitempty" xml:"a0531,omitempty"`
    

    // 用于文化广场最高折扣，跟最低折扣对应
    
    A0530   string `json:"a0530,omitempty" xml:"a0530,omitempty"`
    

    // 库存数量，此功能未启用
    
    A0526   string `json:"a0526,omitempty" xml:"a0526,omitempty"`
    

    // 默认折扣率
    
    A0525   string `json:"a0525,omitempty" xml:"a0525,omitempty"`
    

    // 默认优惠价格
    
    A0524   string `json:"a0524,omitempty" xml:"a0524,omitempty"`
    

    // 折扣类别   0=不允许 1=允许
    
    A0523   string `json:"a0523,omitempty" xml:"a0523,omitempty"`
    

    // 价格下限，此功能已取消
    
    A0522   string `json:"a0522,omitempty" xml:"a0522,omitempty"`
    

    // 价格上限，此功能已取消
    
    A0521   string `json:"a0521,omitempty" xml:"a0521,omitempty"`
    

    // 定价方式   0=定价 1=开价
    
    A0520   string `json:"a0520,omitempty" xml:"a0520,omitempty"`
    

    // 零售价
    
    A0517   string `json:"a0517,omitempty" xml:"a0517,omitempty"`
    

    // 保留字段
    
    A0516   string `json:"a0516,omitempty" xml:"a0516,omitempty"`
    

    // 部门
    
    A0508   string `json:"a0508,omitempty" xml:"a0508,omitempty"`
    

    // 经营方式
    
    A0506   string `json:"a0506,omitempty" xml:"a0506,omitempty"`
    

    // 专柜
    
    A0504   string `json:"a0504,omitempty" xml:"a0504,omitempty"`
    

    // 商品名称
    
    A0503   string `json:"a0503,omitempty" xml:"a0503,omitempty"`
    

    // 商品条码
    
    A0502   string `json:"a0502,omitempty" xml:"a0502,omitempty"`
    

    // 商品编码
    
    A0501   string `json:"a0501,omitempty" xml:"a0501,omitempty"`
    

}
