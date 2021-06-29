package smartstore

// DeviceBizDataDo 
type DeviceBizDataDo struct {
    // "ACTION枚举值： ITEM_CLICK（商品点击时必须设置ITEM_ID） ITEM_SENSOR（商品感应时必须设置ITEM_ID，如RFID、蓝牙、检测仪） RECEIVE_COUPONS（领取优惠券时必须设置COUPON_ID） BUY_CLICK（点击购买） SHARE_CLICK（点击分享）"
    Action   string `json:"action,omitempty" xml:"action,omitempty"`
    // 字段废弃，考虑兼容，等同于op_time，两个必须传一个
    StartTime   string `json:"start_time,omitempty" xml:"start_time,omitempty"`
    // 字段废弃
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 操作时间，后续统一使用该字段，考虑兼容，start_time跟该字段含义一致
    OpTime   string `json:"op_time,omitempty" xml:"op_time,omitempty"`
    // 数据外部编码，保证数据唯一性
    OuterBizId   string `json:"outer_biz_id,omitempty" xml:"outer_biz_id,omitempty"`
    // 业务数据信息
    BizData   *Bizdatamap `json:"biz_data,omitempty" xml:"biz_data,omitempty"`
}
