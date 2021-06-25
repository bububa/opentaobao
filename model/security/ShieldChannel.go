package security

// ShieldChannel 
type ShieldChannel struct {

    // 渠道名称列表
    Values   []String `json:"values,omitempty"`

    // AndroidManifest.xml 中用于表示渠道信息的meta-data标签的android:name
    MetaName   string `json:"meta_name,omitempty"`

}
