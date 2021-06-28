package wms

// Tmsorder 
/* model for simplify = false
type Tmsorder struct {

    // 包材信息
    
    PackageMaterialList  struct {
        Packagemateriallist  []Packagemateriallist `json:"packagemateriallist,omitempty"`
    } `json:"package_material_list,omitempty"`
    

    // 快递公司服务编码
    
    TmsCode   string `json:"tms_code,omitempty"`
    

    // 运单编码
    
    TmsOrderCode   string `json:"tms_order_code,omitempty"`
    

    // 包裹号
    
    PackageCode   string `json:"package_code,omitempty"`
    

    // 包裹重量，单位：克
    
    PackageWeight   int64 `json:"package_weight,omitempty"`
    

    // 包裹长度，单位：毫米
    
    PackageLength   int64 `json:"package_length,omitempty"`
    

    // 包裹宽度，单位：毫米
    
    PackageWidth   int64 `json:"package_width,omitempty"`
    

    // 包裹高度，单位：毫米
    
    PackageHeight   int64 `json:"package_height,omitempty"`
    

    // 包裹里面的商品信息列表
    
    TmsItemList  struct {
        Tmsitemlist  []Tmsitemlist `json:"tmsitemlist,omitempty"`
    } `json:"tms_item_list,omitempty"`
    

}
*/

// Tmsorder 
type Tmsorder struct {

    // 包材信息
    PackageMaterialList   []Packagemateriallist `json:"package_material_list,omitempty"`

    // 快递公司服务编码
    TmsCode   string `json:"tms_code,omitempty"`

    // 运单编码
    TmsOrderCode   string `json:"tms_order_code,omitempty"`

    // 包裹号
    PackageCode   string `json:"package_code,omitempty"`

    // 包裹重量，单位：克
    PackageWeight   int64 `json:"package_weight,omitempty"`

    // 包裹长度，单位：毫米
    PackageLength   int64 `json:"package_length,omitempty"`

    // 包裹宽度，单位：毫米
    PackageWidth   int64 `json:"package_width,omitempty"`

    // 包裹高度，单位：毫米
    PackageHeight   int64 `json:"package_height,omitempty"`

    // 包裹里面的商品信息列表
    TmsItemList   []Tmsitemlist `json:"tms_item_list,omitempty"`

}
