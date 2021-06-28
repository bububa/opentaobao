package wdk

// PdStockDetailDto 
type PdStockDetailDto struct {

    // 备注
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 盘点时间，盘点结果提交的时间
    
    EffectiveTime   string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
    

    // 盘点人员
    
    OperatorName   string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
    

    // 差异数量，盘点结果数量与快照数量的差异值
    
    DiffCount   string `json:"diff_count,omitempty" xml:"diff_count,omitempty"`
    

    // 实盘数量，盘点结果数量
    
    RealCount   string `json:"real_count,omitempty" xml:"real_count,omitempty"`
    

    // 快照数量，盘点任务单下发时的商品数量
    
    SnapshotCount   string `json:"snapshot_count,omitempty" xml:"snapshot_count,omitempty"`
    

    // 库位，盒马系统中的库位号
    
    CabinetCode   string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
    

    // 商品code
    
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
    

}
