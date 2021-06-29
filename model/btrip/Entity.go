package btrip

// Entity 
type Entity struct {

    // 实体id，all_employe为false时，entities里元素的id必传
    
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    

    // 实体名，all_employe为false时，entities里元素的name必传
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 1：员工，all_employe为false时，entities里元素的type必传
    
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    

}
