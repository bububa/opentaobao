package security

// RpUserResult 
/* model for simplify = false
type RpUserResult struct {

    // users
    
    Users  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"users,omitempty"`
    

}
*/

// RpUserResult 
type RpUserResult struct {

    // users
    Users   []string `json:"users,omitempty"`

}
