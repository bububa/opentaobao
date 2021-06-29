package logistic

// WarehouseResourceRelationDTO 
type WarehouseResourceRelationDTO struct {
    // from资源外部编码
    FromOrgResourceCode   string `json:"from_org_resource_code,omitempty" xml:"from_org_resource_code,omitempty"`
    // from资源来源
    FromOrgSource   string `json:"from_org_source,omitempty" xml:"from_org_source,omitempty"`
    // from资源编码
    FromResourceCode   string `json:"from_resource_code,omitempty" xml:"from_resource_code,omitempty"`
    // from资源名称
    FromResourceName   string `json:"from_resource_name,omitempty" xml:"from_resource_name,omitempty"`
    // from资源类型
    FromResourceType   string `json:"from_resource_type,omitempty" xml:"from_resource_type,omitempty"`
    // 商家code
    MerchantCode   string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
    // 网络名称
    NetworkCode   string `json:"network_code,omitempty" xml:"network_code,omitempty"`
    // 关系类型
    RelationType   string `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
    // to资源外部编码
    ToOrgResourceCode   string `json:"to_org_resource_code,omitempty" xml:"to_org_resource_code,omitempty"`
    // to资源来源
    ToOrgSource   string `json:"to_org_source,omitempty" xml:"to_org_source,omitempty"`
    // to资源编码
    ToResourceCode   string `json:"to_resource_code,omitempty" xml:"to_resource_code,omitempty"`
    // to资源名称
    ToResourceName   string `json:"to_resource_name,omitempty" xml:"to_resource_name,omitempty"`
    // to资源类型
    ToResourceType   string `json:"to_resource_type,omitempty" xml:"to_resource_type,omitempty"`
}
