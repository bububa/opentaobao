package logistic

// YxWarehouseResourceRelationQueryRequest 
/* model for simplify = false
type YxWarehouseResourceRelationQueryRequest struct {

    // 网格仓外部编码
    
    ToOrgResourceCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"to_org_resource_codes,omitempty"`
    

}
*/

// YxWarehouseResourceRelationQueryRequest 
type YxWarehouseResourceRelationQueryRequest struct {

    // 网格仓外部编码
    ToOrgResourceCodes   []string `json:"to_org_resource_codes,omitempty"`

}
