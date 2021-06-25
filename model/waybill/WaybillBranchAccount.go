package waybill

// WaybillBranchAccount 
type WaybillBranchAccount struct {

    // 已用面单数量
    AllocatedQuantity   int64 `json:"allocated_quantity,omitempty"`

    // 网点Code
    BranchCode   string `json:"branch_code,omitempty"`

    // 网点名称
    BranchName   string `json:"branch_name,omitempty"`

    // 网点状态
    BranchStatus   int64 `json:"branch_status,omitempty"`

    // 取消的面单总数
    CancelQuantity   int64 `json:"cancel_quantity,omitempty"`

    // 已经打印的面单总数
    PrintQuantity   int64 `json:"print_quantity,omitempty"`

    // 电子面单余额数量
    Quantity   int64 `json:"quantity,omitempty"`

    // 当前网点下的发货地址
    ShippAddressCols   []AddressDto `json:"shipp_address_cols,omitempty"`

    // 可用的服务信息列表
    ServiceInfoCols   []ServiceInfoDto `json:"service_info_cols,omitempty"`

    // 号段信息
    SegmentCode   string `json:"segment_code,omitempty"`

    // 商家ID
    SellerId   int64 `json:"seller_id,omitempty"`

}
