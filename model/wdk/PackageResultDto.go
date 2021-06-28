package wdk

// PackageResultDto 
/* model for simplify = false
type PackageResultDto struct {

    // 路由节点
    
    RouteNodes  struct {
        RouteNodeDto  []RouteNodeDto `json:"route_node_dto,omitempty"`
    } `json:"route_nodes,omitempty"`
    

    // 令牌列表
    
    TokenCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"token_codes,omitempty"`
    

    // 作业要求
    
    OperationalRequirement   string `json:"operational_requirement,omitempty"`
    

    // 合单标
    
    MergeOrderTag   string `json:"merge_order_tag,omitempty"`
    

    // 收件人地址
    
    BuyerAddress   string `json:"buyer_address,omitempty"`
    

    // 收件人电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty"`
    

    // 收件人名称
    
    BuyerName   string `json:"buyer_name,omitempty"`
    

    // 末端合流标示
    
    TimeOrderTag   string `json:"time_order_tag,omitempty"`
    

    // AOI名称
    
    AoiName   string `json:"aoi_name,omitempty"`
    

    // 区块ID
    
    RouteArea   string `json:"route_area,omitempty"`
    

    // 波次时间
    
    WaveTime   string `json:"wave_time,omitempty"`
    

    // 应履约日期
    
    FulfillDate   string `json:"fulfill_date,omitempty"`
    

}
*/

// PackageResultDto 
type PackageResultDto struct {

    // 路由节点
    RouteNodes   []RouteNodeDto `json:"route_nodes,omitempty"`

    // 令牌列表
    TokenCodes   []string `json:"token_codes,omitempty"`

    // 作业要求
    OperationalRequirement   string `json:"operational_requirement,omitempty"`

    // 合单标
    MergeOrderTag   string `json:"merge_order_tag,omitempty"`

    // 收件人地址
    BuyerAddress   string `json:"buyer_address,omitempty"`

    // 收件人电话
    BuyerPhone   string `json:"buyer_phone,omitempty"`

    // 收件人名称
    BuyerName   string `json:"buyer_name,omitempty"`

    // 末端合流标示
    TimeOrderTag   string `json:"time_order_tag,omitempty"`

    // AOI名称
    AoiName   string `json:"aoi_name,omitempty"`

    // 区块ID
    RouteArea   string `json:"route_area,omitempty"`

    // 波次时间
    WaveTime   string `json:"wave_time,omitempty"`

    // 应履约日期
    FulfillDate   string `json:"fulfill_date,omitempty"`

}
