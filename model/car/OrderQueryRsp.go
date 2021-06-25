package car

// OrderQueryRsp 
type OrderQueryRsp struct {

    // 订单id
    OrderId   string `json:"order_id,omitempty"`

    // 订单状态。21-等待商家确认接单，22-商家已确认接单（未派司机），23-商家已确认接单（已派司机），24-司机服务已完成，25-司机已出发，26-司机已到达目的地，27-司机开始服务，60-订单已关闭，70-订单已完成。其他状态可不必关心。
    Status   int64 `json:"status,omitempty"`

}
