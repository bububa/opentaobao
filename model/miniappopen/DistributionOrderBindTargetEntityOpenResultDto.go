package miniappopen

// DistributionOrderBindTargetEntityOpenResultDto 
type DistributionOrderBindTargetEntityOpenResultDto struct {
    // 绑定的目标url
    Url   string `json:"url,omitempty" xml:"url,omitempty"`
    // 绑定投放的场景
    SceneName   string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
    // 商品列表的绑定结果
    BindResultList   []DistributionOrderBindBaseDto `json:"bind_result_list,omitempty" xml:"bind_result_list>distribution_order_bind_base_dto,omitempty"`
}
