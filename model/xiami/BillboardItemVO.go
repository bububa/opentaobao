package xiami

// BillboardItemVO 
/* model for simplify = false
type BillboardItemVO struct {

    // billboardId
    
    BillboardId   int64 `json:"billboard_id,omitempty"`
    

    // songs
    
    Songs  struct {
        Songs  []Songs `json:"songs,omitempty"`
    } `json:"songs,omitempty"`
    

}
*/

// BillboardItemVO 
type BillboardItemVO struct {

    // billboardId
    BillboardId   int64 `json:"billboard_id,omitempty"`

    // songs
    Songs   []Songs `json:"songs,omitempty"`

}
