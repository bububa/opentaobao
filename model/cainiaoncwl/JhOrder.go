package cainiaoncwl

// JhOrder 
type JhOrder struct {
    // 此集货单集货开始时间
    BeginTime   string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
    // 此集货单集货停止时间
    FinishTime   string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
    // 此集货单商品列表
    ItemInfoList   []JhItemInfo `json:"item_info_list,omitempty" xml:"item_info_list>jh_item_info,omitempty"`
    // 收货信息
    JhReceiverInfo   *JhReceiverInfo `json:"jh_receiver_info,omitempty" xml:"jh_receiver_info,omitempty"`
    // 集货单号
    OrderCode   string `json:"order_code,omitempty" xml:"order_code,omitempty"`
    // 补货原主单号
    OriginalOrder   string `json:"original_order,omitempty" xml:"original_order,omitempty"`
    // 集货单状态，CONSOLIDATED 已集货;CONSIGNED 已发货;INBOUND 县仓入库;LACK 缺货LACK
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 预留字段，本集包费用，单位分
    JhFee   *JhFee `json:"jh_fee,omitempty" xml:"jh_fee,omitempty"`
    // 是否是补货
    ReplenishFlag   bool `json:"replenish_flag,omitempty" xml:"replenish_flag,omitempty"`
}
