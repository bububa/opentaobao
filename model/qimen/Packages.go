package qimen

// Packages 
/* model for simplify = false
type Packages struct {

    // 包裹详情
    
    PackageValue  *struct {
        Package  *Package `json:"package,omitempty"`
    } `json:"packageValue,omitempty"`
    

}
*/

// Packages 
type Packages struct {

    // 包裹详情
    PackageValue   *Package `json:"packageValue,omitempty"`

}
