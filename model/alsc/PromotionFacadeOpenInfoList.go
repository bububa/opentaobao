package alsc

// PromotionFacadeOpenInfoList 
type PromotionFacadeOpenInfoList struct {

    // 可用时段  0000101（星期五，星期天）       [{"days":"0000101","startTime":"08:00:00","endTime":"11:59:59"},       {"days":"0010101","startTime":"08:00:00","endTime":"11:59:59"}]
    AvailableTime   string `json:"available_time,omitempty"`

    // 是否已经删除
    Deleted   bool `json:"deleted,omitempty"`

    // 规则描述
    Description   string `json:"description,omitempty"`

    // 促销周期结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 扩展字段  isAdded:是否叠加,isVoucherShared:是否与优惠券共享,giftGoodsIdList:赠品类似、活动商品,privilegeCondition:权益条件(类型type:         满量 "type":"FULL_AMOUNT","name":"num", "value":"3"         满额 "type":"FULL_CAPACITY", "name":"money","value":"300"         下一份 "type":"NEXT", "name":"num","value":"3"         加价购),privilegeType:权益类型(一口价"type":"FIXPRICE","name":"money","value":"3000"         减免"type":"DECREASE","name":"money","value":"10"         减低价"type":"REDUCE_LOW_PRICE","name":"","value":""         折扣："type":"DISCOUNT","name":"discount","value":"80"         赠品 "type":"GIFT", "name":"num", "value":"2")
    ExtInfo   string `json:"ext_info,omitempty"`

    // 适用商品范围  值：ALL，PART_AVAILABLE，PART_UNAVAILABLE      * 说明：全部商品可用，部分商品可用，部分商品不可用
    ItemCoverage   string `json:"item_coverage,omitempty"`

    // 圈选商品
    ItemSelectedOpenInfoList   []ItemSelectedOpenInfoList `json:"item_selected_open_info_list,omitempty"`

    // 促销规则名称
    Name   string `json:"name,omitempty"`

    // 促销规则Id
    PromotionId   string `json:"promotion_id,omitempty"`

    // 圈选门店
    ShopSelectedOpenInfoList   []ShopSelectedOpenInfoList `json:"shop_selected_open_info_list,omitempty"`

    // 促销周期开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 状态 促销活动状态 值：UNUSED,USED,NO_INVENTORY,INVALID 说明：未使用,使用中,使用中,使用中
    Status   string `json:"status,omitempty"`

    // 促销活动适用人群 值：MEMBER，CUSTOMER，ALL      * 说明：会员,非会员，不限
    SuitablePeople   string `json:"suitable_people,omitempty"`

    // /**      * 满量促销      */     TYPE_FULL_AMOUNT,      /**      * 满额促销      */     TYPE_FULL_CAPACITY,      /**      * 买赠活动      */     TYPE_BOUGHT_GIFT;
    Type   string `json:"type,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 更新人
    UpdateBy   string `json:"update_by,omitempty"`

    // 更新人名称
    UpdateByName   string `json:"update_by_name,omitempty"`

}
