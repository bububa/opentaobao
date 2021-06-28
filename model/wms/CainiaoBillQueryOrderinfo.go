package wms

// CainiaoBillQueryOrderinfo 
/* model for simplify = false
type CainiaoBillQueryOrderinfo struct {

    // 菜鸟订单编码
    
    CnOrderCode   string `json:"cn_order_code,omitempty"`
    

    // ERP订单号
    
    OrderCode   string `json:"order_code,omitempty"`
    

    // 单据类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库 904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
    
    OrderSource   string `json:"order_source,omitempty"`
    

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 单据状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败 WMS_CONFIRMED 仓库生产完成，注：此状态表示单据已经在WMS处理完成，可能通过获取单据详情接口获取单据详细信息。 WMS_CANCEL 取消仓库发货  -- WMS_FAILED 订单发货失败 TMS_SIGN 买家签收 TMS_REJECT 买家拒签 TMS_CANCEL 拦截派送
    
    Status   string `json:"status,omitempty"`
    

    // 订单最后修改时间
    
    ModifiedTime   string `json:"modified_time,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 交易订单号
    
    OrderSourceCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"order_source_codes,omitempty"`
    

}
*/

// CainiaoBillQueryOrderinfo 
type CainiaoBillQueryOrderinfo struct {

    // 菜鸟订单编码
    CnOrderCode   string `json:"cn_order_code,omitempty"`

    // ERP订单号
    OrderCode   string `json:"order_code,omitempty"`

    // 单据类型 201 销售出库 501 退货入库 502 换货出库 503 补发出库 904 普通入库 903 普通出库单 306 B2B入库单 305 B2B出库单 601 采购入库 901 退供出库单 701 盘点出库 702 盘点入库 711 库存异动单
    OrderType   int64 `json:"order_type,omitempty"`

    // 订单来源（213 天猫，201 淘宝，214 京东，202 1688 阿里中文站 ，203 苏宁在线，204 亚马逊中国，205 当当，208 1号店，207 唯品会，209 国美在线，210 拍拍，206 易贝ebay，211 聚美优品，212 乐蜂网，215 邮乐，216 凡客，217 优购，218 银泰，219 易讯，221 聚尚网，222 蘑菇街，223 POS门店，301 其他）
    OrderSource   string `json:"order_source,omitempty"`

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 单据状态 WMS_ACCEPT 接单成功 WMS_REJECT 接单失败 WMS_CONFIRMED 仓库生产完成，注：此状态表示单据已经在WMS处理完成，可能通过获取单据详情接口获取单据详细信息。 WMS_CANCEL 取消仓库发货  -- WMS_FAILED 订单发货失败 TMS_SIGN 买家签收 TMS_REJECT 买家拒签 TMS_CANCEL 拦截派送
    Status   string `json:"status,omitempty"`

    // 订单最后修改时间
    ModifiedTime   string `json:"modified_time,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 交易订单号
    OrderSourceCodes   []string `json:"order_source_codes,omitempty"`

}
