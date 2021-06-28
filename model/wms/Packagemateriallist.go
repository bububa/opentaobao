package wms

// Packagemateriallist 
/* model for simplify = false
type Packagemateriallist struct {

    // 包裹包材信息
    
    PackageMaterial  *struct {
        Packagematerial  *Packagematerial `json:"packagematerial,omitempty"`
    } `json:"package_material,omitempty"`
    

}
*/

// Packagemateriallist 
type Packagemateriallist struct {

    // 包裹包材信息
    PackageMaterial   *Packagematerial `json:"package_material,omitempty"`

}
