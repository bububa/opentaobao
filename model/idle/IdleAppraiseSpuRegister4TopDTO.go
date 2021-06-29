package idle

// IdleAppraiseSpuRegister4TopDTO 
type IdleAppraiseSpuRegister4TopDTO struct {
    // 操作类型，0新增，-1删除。当spu第一次挂载时，会进入1测试中状态。服务商联调通过后，需要再次挂载，actionType还传0，挂载信息状态会变成0已上线。
    ActionType   int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
    // 验货类型，1新品，2二手
    AppraiseScene   int64 `json:"appraise_scene,omitempty" xml:"appraise_scene,omitempty"`
    // brandId
    BrandId   int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 类目Id
    CateId   string `json:"cate_id,omitempty" xml:"cate_id,omitempty"`
    // spuId
    SpuId   int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}
