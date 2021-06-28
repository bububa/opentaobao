package hotel

// PicStringArrayDo 
/* model for simplify = false
type PicStringArrayDo struct {

    // prefix
    
    Prefix   string `json:"prefix,omitempty"`
    

    // sources
    
    Sources  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sources,omitempty"`
    

    // suffixs
    
    Suffixs  struct {
        String  []string `json:"string,omitempty"`
    } `json:"suffixs,omitempty"`
    

}
*/

// PicStringArrayDo 
type PicStringArrayDo struct {

    // prefix
    Prefix   string `json:"prefix,omitempty"`

    // sources
    Sources   []string `json:"sources,omitempty"`

    // suffixs
    Suffixs   []string `json:"suffixs,omitempty"`

}
