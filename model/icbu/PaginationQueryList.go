package icbu

// PaginationQueryList 
/* model for simplify = false
type PaginationQueryList struct {

    // list
    
    List  struct {
        PhotobankImageDo  []PhotobankImageDo `json:"photobank_image_do,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// PaginationQueryList 
type PaginationQueryList struct {

    // list
    List   []PhotobankImageDo `json:"list,omitempty"`

}
