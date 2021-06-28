package security

// ShieldChannel 
/* model for simplify = false
type ShieldChannel struct {

    // 渠道名称列表
    
    Values  struct {
        String  []string `json:"string,omitempty"`
    } `json:"values,omitempty"`
    

    // AndroidManifest.xml 中用于表示渠道信息的meta-data标签的android:name
    
    MetaName   string `json:"meta_name,omitempty"`
    

}
*/

// ShieldChannel 
type ShieldChannel struct {

    // 渠道名称列表
    Values   []string `json:"values,omitempty"`

    // AndroidManifest.xml 中用于表示渠道信息的meta-data标签的android:name
    MetaName   string `json:"meta_name,omitempty"`

}
